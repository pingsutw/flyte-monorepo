import typing
from collections import OrderedDict

import pytest

from flytekit.common.exceptions.user import FlyteValidationException
from flytekit.common.translator import get_serializable
from flytekit.core import context_manager
from flytekit.core.context_manager import Image, ImageConfig
from flytekit.core.launch_plan import LaunchPlan
from flytekit.core.task import task
from flytekit.core.workflow import ImperativeWorkflow, get_promise, workflow
from flytekit.models import literals as literal_models

default_img = Image(name="default", fqn="test", tag="tag")
serialization_settings = context_manager.SerializationSettings(
    project="project",
    domain="domain",
    version="version",
    env=None,
    image_config=ImageConfig(default_image=default_img, images=[default_img]),
)


def test_imperative():
    @task
    def t1(a: str) -> str:
        return a + " world"

    @task
    def t2():
        print("side effect")

    wb = ImperativeWorkflow(name="my.workflow")
    wb.add_workflow_input("in1", str)
    node = wb.add_entity(t1, a=wb.inputs["in1"])
    wb.add_entity(t2)
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb(in1="hello") == "hello world"

    wf_spec = get_serializable(OrderedDict(), serialization_settings, wb)
    assert len(wf_spec.template.nodes) == 2
    assert wf_spec.template.nodes[0].task_node is not None
    assert len(wf_spec.template.outputs) == 1
    assert wf_spec.template.outputs[0].var == "from_n0t1"
    assert len(wf_spec.template.interface.inputs) == 1
    assert len(wf_spec.template.interface.outputs) == 1

    # Create launch plan from wf, that can also be serialized.
    lp = LaunchPlan.create("test_wb", wb)
    lp_model = get_serializable(OrderedDict(), serialization_settings, lp)
    assert lp_model.spec.workflow_id.name == "my.workflow"


def test_imperative_list_bound():
    @task
    def t1(a: typing.List[int]) -> int:
        return sum(a)

    wb = ImperativeWorkflow(name="my.workflow.a")
    wb.add_workflow_input("in1", int)
    wb.add_workflow_input("in2", int)
    node = wb.add_entity(t1, a=[wb.inputs["in1"], wb.inputs["in2"]])
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb(in1=3, in2=4) == 7


def test_imperative_map_bound():
    @task
    def t1(a: typing.Dict[str, typing.List[int]]) -> typing.Dict[str, int]:
        return {k: sum(v) for k, v in a.items()}

    wb = ImperativeWorkflow(name="my.workflow.a")
    in1 = wb.add_workflow_input("in1", int)
    wb.add_workflow_input("in2", int)
    in3 = wb.add_workflow_input("in3", int)
    node = wb.add_entity(t1, a={"a": [in1, wb.inputs["in2"]], "b": [wb.inputs["in2"], in3]})
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb(in1=3, in2=4, in3=5) == {"a": 7, "b": 9}


def test_imperative_with_list_io():
    @task
    def t1(a: int) -> typing.List[int]:
        return [1, a, 3]

    @task
    def t2(a: typing.List[int]) -> int:
        return sum(a)

    wb = ImperativeWorkflow(name="my.workflow.a")
    t1_node = wb.add_entity(t1, a=2)
    t2_node = wb.add_entity(t2, a=t1_node.outputs["o0"])
    wb.add_workflow_output("from_n0t2", t2_node.outputs["o0"])

    assert wb() == 6


def test_imperative_wf_list_input():
    @task
    def t1(a: int) -> typing.List[int]:
        return [1, a, 3]

    @task
    def t2(a: typing.List[int], b: typing.List[int]) -> int:
        return sum(a) + sum(b)

    wb = ImperativeWorkflow(name="my.workflow.a")
    wf_in1 = wb.add_workflow_input("in1", typing.List[int])
    t1_node = wb.add_entity(t1, a=2)
    t2_node = wb.add_entity(t2, a=t1_node.outputs["o0"], b=wf_in1)
    wb.add_workflow_output("from_n0t2", t2_node.outputs["o0"])

    assert wb(in1=[5, 6, 7]) == 24

    wf_spec = get_serializable(OrderedDict(), serialization_settings, wb)
    assert len(wf_spec.template.nodes) == 2
    assert wf_spec.template.nodes[0].task_node is not None


def test_imperative_scalar_bindings():
    @task
    def t1(a: typing.Dict[str, typing.List[int]]) -> typing.Dict[str, int]:
        return {k: sum(v) for k, v in a.items()}

    wb = ImperativeWorkflow(name="my.workflow.a")
    node = wb.add_entity(t1, a={"a": [3, 4], "b": [5, 6]})
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb() == {"a": 7, "b": 11}

    wf_spec = get_serializable(OrderedDict(), serialization_settings, wb)
    assert len(wf_spec.template.nodes) == 1
    assert wf_spec.template.nodes[0].task_node is not None


def test_imperative_list_bound_output():
    @task
    def t1() -> int:
        return 3

    @task
    def t2(a: typing.List[int]) -> int:
        return sum(a)

    wb = ImperativeWorkflow(name="my.workflow.a")
    t1_node = wb.add_entity(t1)
    t2_node = wb.add_entity(t2, a=[1, 2, 3])
    wb.add_workflow_output("wf0", [t1_node.outputs["o0"], t2_node.outputs["o0"]], python_type=typing.List[int])

    assert wb() == [3, 6]


def test_call_normal():
    @task
    def t1(a: int) -> (int, str):
        return a + 2, "world"

    @workflow
    def my_functional_wf(a: int) -> (int, str):
        return t1(a=a)

    my_functional_lp = LaunchPlan.create("my_functional_wf.lp0", my_functional_wf, default_inputs={"a": 3})

    wb = ImperativeWorkflow(name="imperio")
    node = wb.add_entity(my_functional_wf, a=3)
    wb.add_workflow_output("from_n0_1", node.outputs["o0"])
    wb.add_workflow_output("from_n0_2", node.outputs["o1"])

    assert wb() == (5, "world")

    wb_lp = ImperativeWorkflow(name="imperio")
    node = wb_lp.add_entity(my_functional_lp)
    wb_lp.add_workflow_output("from_n0_1", node.outputs["o0"])
    wb_lp.add_workflow_output("from_n0_2", node.outputs["o1"])

    assert wb_lp() == (5, "world")


def test_imperative_call_from_normal():
    @task
    def t1(a: str) -> str:
        return a + " world"

    wb = ImperativeWorkflow(name="my.workflow")
    wb.add_workflow_input("in1", str)
    node = wb.add_entity(t1, a=wb.inputs["in1"])
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb(in1="hello") == "hello world"

    @workflow
    def my_functional_wf(a: str) -> str:
        x = wb(in1=a)
        return x

    assert my_functional_wf(a="hello") == "hello world"

    # Create launch plan from wf
    lp = LaunchPlan.create("test_wb_2", wb, fixed_inputs={"in1": "hello"})

    @workflow
    def my_functional_wf_lp() -> str:
        x = lp()
        return x

    assert my_functional_wf_lp() == "hello world"


def test_codecov():
    with pytest.raises(FlyteValidationException):
        get_promise(literal_models.BindingData(), {})

    with pytest.raises(FlyteValidationException):
        get_promise(literal_models.BindingData(promise=3), {})

    @task
    def t1(a: str) -> str:
        return a + " world"

    wb = ImperativeWorkflow(name="my.workflow")
    wb.add_workflow_input("in1", str)
    node = wb.add_entity(t1, a=wb.inputs["in1"])
    wb.add_workflow_output("from_n0t1", node.outputs["o0"])

    assert wb(in1="hello") == "hello world"

    with pytest.raises(AssertionError):
        wb(3)

    with pytest.raises(ValueError):
        wb(in2="hello")

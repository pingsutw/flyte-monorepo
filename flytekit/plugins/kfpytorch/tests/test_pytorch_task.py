from flytekitplugins.kfpytorch.task import PyTorch

from flytekit import task
from flytekit.annotated.context_manager import Image, ImageConfig, SerializationSettings
from flytekit.annotated.resources import Resources


def test_pytorch_task():
    @task(task_config=PyTorch(num_workers=10, per_replica_requests=Resources(cpu="1")), cache=True, cache_version="1")
    def my_pytorch_task(x: int, y: str) -> int:
        return x

    assert my_pytorch_task(x=10, y="hello") == 10

    assert my_pytorch_task.task_config is not None

    default_img = Image(name="default", fqn="test", tag="tag")
    settings = SerializationSettings(
        project="project",
        domain="domain",
        version="version",
        env={"FOO": "baz"},
        image_config=ImageConfig(default_image=default_img, images=[default_img]),
    )

    assert my_pytorch_task.get_custom(settings) == {"workers": 10}
    assert my_pytorch_task.resources.limits == Resources()
    assert my_pytorch_task.resources.requests == Resources(cpu="1")
    assert my_pytorch_task.task_type == "pytorch"

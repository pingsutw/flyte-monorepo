# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/condition.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1d\x66lyteidl/core/condition.proto\x12\rflyteidl.core\x1a\x1c\x66lyteidl/core/literals.proto\"\x8f\x02\n\x14\x43omparisonExpression\x12H\n\x08operator\x18\x01 \x01(\x0e\x32,.flyteidl.core.ComparisonExpression.OperatorR\x08operator\x12\x35\n\nleft_value\x18\x02 \x01(\x0b\x32\x16.flyteidl.core.OperandR\tleftValue\x12\x37\n\x0bright_value\x18\x03 \x01(\x0b\x32\x16.flyteidl.core.OperandR\nrightValue\"=\n\x08Operator\x12\x06\n\x02\x45Q\x10\x00\x12\x07\n\x03NEQ\x10\x01\x12\x06\n\x02GT\x10\x02\x12\x07\n\x03GTE\x10\x03\x12\x06\n\x02LT\x10\x04\x12\x07\n\x03LTE\x10\x05\"^\n\x07Operand\x12\x38\n\tprimitive\x18\x01 \x01(\x0b\x32\x18.flyteidl.core.PrimitiveH\x00R\tprimitive\x12\x12\n\x03var\x18\x02 \x01(\tH\x00R\x03varB\x05\n\x03val\"\xac\x01\n\x11\x42ooleanExpression\x12H\n\x0b\x63onjunction\x18\x01 \x01(\x0b\x32$.flyteidl.core.ConjunctionExpressionH\x00R\x0b\x63onjunction\x12\x45\n\ncomparison\x18\x02 \x01(\x0b\x32#.flyteidl.core.ComparisonExpressionH\x00R\ncomparisonB\x06\n\x04\x65xpr\"\xa5\x02\n\x15\x43onjunctionExpression\x12P\n\x08operator\x18\x01 \x01(\x0e\x32\x34.flyteidl.core.ConjunctionExpression.LogicalOperatorR\x08operator\x12I\n\x0fleft_expression\x18\x02 \x01(\x0b\x32 .flyteidl.core.BooleanExpressionR\x0eleftExpression\x12K\n\x10right_expression\x18\x03 \x01(\x0b\x32 .flyteidl.core.BooleanExpressionR\x0frightExpression\"\"\n\x0fLogicalOperator\x12\x07\n\x03\x41ND\x10\x00\x12\x06\n\x02OR\x10\x01\x42\xae\x01\n\x11\x63om.flyteidl.coreB\x0e\x43onditionProtoP\x01Z4github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core\xa2\x02\x03\x46\x43X\xaa\x02\rFlyteidl.Core\xca\x02\rFlyteidl\\Core\xe2\x02\x19\x46lyteidl\\Core\\GPBMetadata\xea\x02\x0e\x46lyteidl::Coreb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.core.condition_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.flyteidl.coreB\016ConditionProtoP\001Z4github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core\242\002\003FCX\252\002\rFlyteidl.Core\312\002\rFlyteidl\\Core\342\002\031Flyteidl\\Core\\GPBMetadata\352\002\016Flyteidl::Core'
  _COMPARISONEXPRESSION._serialized_start=79
  _COMPARISONEXPRESSION._serialized_end=350
  _COMPARISONEXPRESSION_OPERATOR._serialized_start=289
  _COMPARISONEXPRESSION_OPERATOR._serialized_end=350
  _OPERAND._serialized_start=352
  _OPERAND._serialized_end=446
  _BOOLEANEXPRESSION._serialized_start=449
  _BOOLEANEXPRESSION._serialized_end=621
  _CONJUNCTIONEXPRESSION._serialized_start=624
  _CONJUNCTIONEXPRESSION._serialized_end=917
  _CONJUNCTIONEXPRESSION_LOGICALOPERATOR._serialized_start=883
  _CONJUNCTIONEXPRESSION_LOGICALOPERATOR._serialized_end=917
# @@protoc_insertion_point(module_scope)

//
//  Generated code. Do not modify.
//  source: sudoku/entities/v1/board.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use boardDescriptor instead')
const Board$json = {
  '1': 'Board',
  '2': [
    {'1': 'rows', '3': 1, '4': 3, '5': 11, '6': '.sudoku.entities.v1.Row', '10': 'rows'},
  ],
};

/// Descriptor for `Board`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List boardDescriptor = $convert.base64Decode(
    'CgVCb2FyZBIrCgRyb3dzGAEgAygLMhcuc3Vkb2t1LmVudGl0aWVzLnYxLlJvd1IEcm93cw==');

@$core.Deprecated('Use rowDescriptor instead')
const Row$json = {
  '1': 'Row',
  '2': [
    {'1': 'cells', '3': 1, '4': 3, '5': 11, '6': '.sudoku.entities.v1.Cell', '10': 'cells'},
  ],
};

/// Descriptor for `Row`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rowDescriptor = $convert.base64Decode(
    'CgNSb3cSLgoFY2VsbHMYASADKAsyGC5zdWRva3UuZW50aXRpZXMudjEuQ2VsbFIFY2VsbHM=');

@$core.Deprecated('Use cellDescriptor instead')
const Cell$json = {
  '1': 'Cell',
  '2': [
    {'1': 'value', '3': 1, '4': 1, '5': 13, '10': 'value'},
    {'1': 'is_const', '3': 2, '4': 1, '5': 8, '10': 'isConst'},
  ],
};

/// Descriptor for `Cell`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List cellDescriptor = $convert.base64Decode(
    'CgRDZWxsEhQKBXZhbHVlGAEgASgNUgV2YWx1ZRIZCghpc19jb25zdBgCIAEoCFIHaXNDb25zdA'
    '==');


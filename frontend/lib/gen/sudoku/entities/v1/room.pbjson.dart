//
//  Generated code. Do not modify.
//  source: sudoku/entities/v1/room.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roomDescriptor instead')
const Room$json = {
  '1': 'Room',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'players', '3': 3, '4': 3, '5': 11, '6': '.sudoku.entities.v1.User', '10': 'players'},
    {'1': 'board', '3': 4, '4': 1, '5': 11, '6': '.sudoku.entities.v1.Board', '10': 'board'},
  ],
};

/// Descriptor for `Room`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roomDescriptor = $convert.base64Decode(
    'CgRSb29tEg4KAmlkGAEgASgNUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEjIKB3BsYXllcnMYAy'
    'ADKAsyGC5zdWRva3UuZW50aXRpZXMudjEuVXNlclIHcGxheWVycxIvCgVib2FyZBgEIAEoCzIZ'
    'LnN1ZG9rdS5lbnRpdGllcy52MS5Cb2FyZFIFYm9hcmQ=');


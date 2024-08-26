//
//  Generated code. Do not modify.
//  source: sudoku/auth/v1/auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use oAuthProviderDescriptor instead')
const OAuthProvider$json = {
  '1': 'OAuthProvider',
  '2': [
    {'1': 'OAUTH_PROVIDER_UNSPECIFIED', '2': 0},
    {'1': 'OAUTH_PROVIDER_GITHUB', '2': 1},
  ],
};

/// Descriptor for `OAuthProvider`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List oAuthProviderDescriptor = $convert.base64Decode(
    'Cg1PQXV0aFByb3ZpZGVyEh4KGk9BVVRIX1BST1ZJREVSX1VOU1BFQ0lGSUVEEAASGQoVT0FVVE'
    'hfUFJPVklERVJfR0lUSFVCEAE=');

@$core.Deprecated('Use signInStatusDescriptor instead')
const SignInStatus$json = {
  '1': 'SignInStatus',
  '2': [
    {'1': 'SIGN_IN_STATUS_UNSPECIFIED', '2': 0},
    {'1': 'SIGN_IN_STATUS_ALREADY_SIGNED_IN', '2': 1},
    {'1': 'SIGN_IN_STATUS_REQUIRES_SIGNING_IN', '2': 2},
  ],
};

/// Descriptor for `SignInStatus`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List signInStatusDescriptor = $convert.base64Decode(
    'CgxTaWduSW5TdGF0dXMSHgoaU0lHTl9JTl9TVEFUVVNfVU5TUEVDSUZJRUQQABIkCiBTSUdOX0'
    'lOX1NUQVRVU19BTFJFQURZX1NJR05FRF9JThABEiYKIlNJR05fSU5fU1RBVFVTX1JFUVVJUkVT'
    'X1NJR05JTkdfSU4QAg==');

@$core.Deprecated('Use signInRequestDescriptor instead')
const SignInRequest$json = {
  '1': 'SignInRequest',
  '2': [
    {'1': 'provider', '3': 1, '4': 1, '5': 14, '6': '.sudoku.auth.v1.OAuthProvider', '10': 'provider'},
  ],
};

/// Descriptor for `SignInRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signInRequestDescriptor = $convert.base64Decode(
    'Cg1TaWduSW5SZXF1ZXN0EjkKCHByb3ZpZGVyGAEgASgOMh0uc3Vkb2t1LmF1dGgudjEuT0F1dG'
    'hQcm92aWRlclIIcHJvdmlkZXI=');

@$core.Deprecated('Use signInResponseDescriptor instead')
const SignInResponse$json = {
  '1': 'SignInResponse',
  '2': [
    {'1': 'authorization_url', '3': 1, '4': 1, '5': 9, '10': 'authorizationUrl'},
    {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.sudoku.auth.v1.SignInStatus', '10': 'status'},
  ],
};

/// Descriptor for `SignInResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signInResponseDescriptor = $convert.base64Decode(
    'Cg5TaWduSW5SZXNwb25zZRIrChFhdXRob3JpemF0aW9uX3VybBgBIAEoCVIQYXV0aG9yaXphdG'
    'lvblVybBI0CgZzdGF0dXMYAiABKA4yHC5zdWRva3UuYXV0aC52MS5TaWduSW5TdGF0dXNSBnN0'
    'YXR1cw==');

@$core.Deprecated('Use signOutRequestDescriptor instead')
const SignOutRequest$json = {
  '1': 'SignOutRequest',
};

/// Descriptor for `SignOutRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signOutRequestDescriptor = $convert.base64Decode(
    'Cg5TaWduT3V0UmVxdWVzdA==');

@$core.Deprecated('Use signOutResponseDescriptor instead')
const SignOutResponse$json = {
  '1': 'SignOutResponse',
};

/// Descriptor for `SignOutResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signOutResponseDescriptor = $convert.base64Decode(
    'Cg9TaWduT3V0UmVzcG9uc2U=');


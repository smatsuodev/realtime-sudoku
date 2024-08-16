//
//  Generated code. Do not modify.
//  source: sudoku/auth/v1/auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'auth.pb.dart' as $0;

export 'auth.pb.dart';

@$pb.GrpcServiceName('sudoku.auth.v1.AuthService')
class AuthServiceClient extends $grpc.Client {
  static final _$signIn = $grpc.ClientMethod<$0.SignInRequest, $0.SignInResponse>(
      '/sudoku.auth.v1.AuthService/SignIn',
      ($0.SignInRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.SignInResponse.fromBuffer(value));
  static final _$signOut = $grpc.ClientMethod<$0.SignOutRequest, $0.SignOutResponse>(
      '/sudoku.auth.v1.AuthService/SignOut',
      ($0.SignOutRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.SignOutResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.SignInResponse> signIn($0.SignInRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$signIn, request, options: options);
  }

  $grpc.ResponseFuture<$0.SignOutResponse> signOut($0.SignOutRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$signOut, request, options: options);
  }
}

@$pb.GrpcServiceName('sudoku.auth.v1.AuthService')
abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'sudoku.auth.v1.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.SignInRequest, $0.SignInResponse>(
        'SignIn',
        signIn_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SignInRequest.fromBuffer(value),
        ($0.SignInResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SignOutRequest, $0.SignOutResponse>(
        'SignOut',
        signOut_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SignOutRequest.fromBuffer(value),
        ($0.SignOutResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.SignInResponse> signIn_Pre($grpc.ServiceCall call, $async.Future<$0.SignInRequest> request) async {
    return signIn(call, await request);
  }

  $async.Future<$0.SignOutResponse> signOut_Pre($grpc.ServiceCall call, $async.Future<$0.SignOutRequest> request) async {
    return signOut(call, await request);
  }

  $async.Future<$0.SignInResponse> signIn($grpc.ServiceCall call, $0.SignInRequest request);
  $async.Future<$0.SignOutResponse> signOut($grpc.ServiceCall call, $0.SignOutRequest request);
}

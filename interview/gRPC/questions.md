1. grpc有哪四种服务类型？
   1. 简单 RPC（Unary RPC）一般的rpc调用，传入一个请求对象，返回一个返回对象
   2. 服务端流式 RPC （Server streaming RPC）传入一个请求对象，服务端可以返回多个结果对象
   3. 客户端流式 RPC （Client streaming RPC）客户端传入多个请求对象，服务端返回一个结果对象
   4. 双向流式 RPC（Bi-directional streaming RPC）结合客户端流式RPC和服务端流式RPC，可以传入多个请求对象，返回多个结果对象
   gRPC 是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计。
2. 
Title: Protobuf in Go
Id: 9729
Versions:
[{"Name":"1.4.0","GroupName":null},{"Name":"1.5.0","GroupName":null},{"Name":"1.6.0","GroupName":null},{"Name":"1.7.0","GroupName":null},{"Name":"1.7.3","GroupName":null},{"Name":"1.8.0","GroupName":null}]
|======|
Introduction:
**Protobuf** or Protocol Buffer encodes and decodes data so that different applications or modules written in unlike languages can exchange the large number of messages quickly and reliably without overloading the communication channel. With protobuf, the performance is directly proportional to the number of message you tend to send. It compress the message to send in a serialized binary format by providing your the tools to encode the message at source and decode it at the destination.
|======|
Remarks:
There are two steps of using **protobuf**. 

 1. First you must compile the protocol buffer definitions
 2. Import the above definitions, with the support library into your program.

**gRPC Support**

If a proto file specifies RPC services, protoc-gen-go can be instructed to generate code compatible with gRPC (http://www.grpc.io/). To do this, pass the `plugins` parameter to protoc-gen-go; the usual way is to insert it into the --go_out argument to protoc:

    protoc --go_out=plugins=grpc:. *.proto



|======|

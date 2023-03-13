# TomaRPC

## What is RPC?
RPC (Remote Procedure Call) is a computer communication protocol that allows programs in different process spaces to be called. The RPC client and server can be on the same machine or on different machines. When programmers use it, it is like calling a local program without paying attention to internal implementation details.

There are many communication methods between different applications, such as the HTTP protocol-based Restful API widely used between browsers and servers. Compared with RPC, Restful API has a relatively unified standard, so it is more general, better compatible, and supports different languages. The HTTP protocol is text-based, but has the following disadvantages:

1. The Restful interface requires additional definitions, whether it is the client or the server, requires additional code to process, while the RPC call is closer to a direct call.
2.Restful message redundancy based on the HTTP protocol carries too much invalid information, while RPC usually uses a custom protocol format to reduce redundant messages.
3. RPC can adopt a more efficient serialization protocol to convert text into binary transmission to obtain higher performance.

Due to the flexibility of RPC, it is easier to expand and integrate functions such as registry and load balancing.

## What problems does the RPC framework need to solve

1. Solving transport protocols and message encodings
First of all, the application programs between the two machines need to communicate, and the transmission protocol needs to be determined first. If they are different machines, the TCP protocol or HTTP protocol is generally used, and if they are the same, the Unix Socket protocol can be used. After the protocol is determined, determine the encoding format of the message, such as JSON and XML. If it is relatively large, the protobuf method can be used. Therefore, codec is required, and even compression and decompression operations are needed.

2. Solve usability issues, such as connection timeout , support for asynchronous requests and concurrency

If there are many instances on the server side, the client does not care about the addresses and deployment locations of these instances, but only cares about whether it can obtain the expected results, which leads to the problems of registry and load balance. To put it simply, the client and the server do not perceive each other's existence. When the server starts, it registers itself with the registration center. When the client calls, it obtains all available instances from the registration center and selects one to call. In this way, the server and the client only need to perceive the existence of the registration center. The registration center usually also needs to implement functions such as dynamic addition and deletion of services, and use of heartbeat to ensure that services are available.

In addition, if the server is provided by different teams, if there is no unified RPC framework, then the servers of each team need to implement a set of message encoding and decoding, connection pool, sending and receiving threads, and timeout processing, which increases the workload and complexity.

## About
The RPC framework in this repository is mainly used for self-learning and use, for a deeper understanding of the essence of the RPC framework and to improve one's understanding of the go language. This framework implements the core part of the Go language standard library net/rpc from scratch, and adds protocol exchange, registry center, service discovery, load balance, timeout processing and other features.
# Week 7 - Service Integration

## What is Service Integration?

You can imagine services similar to a set of organs in the body all contributing
to solve a single, larger objective of living. Then **service integration** would
be the process of connecting those services together through blood vessels.

> SIAM (short for Service Integration And Management) is the process of managing
> multiple independent services together, from one or many different tools, in
> order to unify their workflow or process heading towards a business objective.
>
> - 🤓

## Integration Methods

### Synchronous Integration

Similar to us having a chat, after we speak (a request), we wait for the opponent's
reply (response). Because we had to wait (blocking), it is a Synchronous operation.

Example: REST API. When our app makes a request to the system, we have to wait for
its response before continuing with the application's cycle. If we are logging
in, we need to wait for the system to check if our credentials are correct,
before we can continue the session with logged in details.

**The Weakness**: It requires both sides to be online or available to actually work,
for example, if the API is down, then the app has no data to go forward with its
intended functions.

### Asynchronous Integration

Similar to sending an email. We can send the email (called a message), and we don't
have to wait for an instant reply to continue our work. When the opponent has some
time, they can "acknowledge" that they received our message later. We didn't have
to wait (non-blocking), therefore, it was asynchronous.

Example: Message Brokers or Data Streaming applications. We can send an information
piece, and wait for it to be processed, and when it is eventually processed, we will
get its results later.

**The Weakness**: It gets increasingly complex to manage such an asynchronous system,
in the case of an error, it's hard to track down whether that error came from the
publisher, or the consumer, or the broker itself. (_It also creates students that
claim mastery over Kafka, while knowing only how to send and receive messages_)

## gRPC - Synchronous Communication

### What is gRPC?

gRPC is a protocol for synchronous communication, based on Remote Procedure Calls,
developed by Google. _It also has a mascot named Golden Retriever PanCakes_.

Overall, gRPC allows a system to call a method on another system as if it's a local
method, allows for smooth integration between distributed applications. When you
look at it in an implementation view, the server will run a layer called gRPC server,
that listens and handles gRPC requests, whereas the client has a "stub" that packages
information and sends to the server based on gRPC's protocol.

gRPC has officially supported libraries and languages:

- C# / .NET (ew)
- C++ (peak)
- Dart
- Go
- Java
- Kotlin
- Node
- ObjectiveC
- PHP
- Python
- Ruby
- Swift

### The gRPC Architecture

![gRPC Architecture](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.koyeb.com%2Fstatic%2Fimages%2Fblog%2Fwhat-is-grpc%2Fgrpc-architecture.png&f=1&nofb=1&ipt=2bc4499392812323c5dde0d72a8eb9c3a606bfee67b958fcfb214b40aa9588bf)

Each service, requests and responses are defined under the file `.proto` as an
intermediate layer, similar to how LLVM backends work, and then it gets
transformed into the language's binding.

gRPC uses protocol buffers (a serialization method, discussed below) to transfer
data.

### Protocol Buffers

Protocol Buffer is a method to serialize structured data in a way that it is
extensible, adaptable to many languages or platforms. Google claims that it is
JSON, but much smaller and faster.

Google also claims that this form of data transmission is used the most inside
the company, and also used for all communication needs between servers.

```proto
edition = "2023";

message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;
}
```

### Types of gRPC calls

- **Unary RPC**: as simple as it gets, the client sends a request, and expects
  a single response. Similar to REST API.
- **Server Streaming RPC**: client sends a request, and the server can send
  back a stream of responses. Similar to Ollama's API, when you send a single
  request, it sends back a stream of tokens instead of the final response once.
- **Client Streaming RPC**: The client sends a stream of requests, and the
  server can response after it has read all of the streams. Similar to sending
  your professor a bunch of emails explaining and he replies with "ok".
- **Bidrectional Streaming RPC**: Both sides send requests and receive responses
  in real-time, but how detailed and when to respond, is up to the developer.
  Similar to websockets.

### gRPC Merits

- High performance and efficiency: Protocol Buffers are binary data, which is
  much faster to deserialize and transfer compared to text-based data formats like
  XML or JSON.
- Allows bidirectional streaming: both the client and the server can communicate
  at real-time using this feature.
- Supports many languages: many languages have official libraries supported by
  Google to implement the gRPC protocols.
- Type-checking: Protocol Buffers allow gRPC to typecheck each message, so that
  transferred messages are not mismatched when dealing with garbage languages like
  JavaScript.

### gRPC Demerits

- Applicable to any new technologies, it requires time and energy investment to
  learn a new thing, or get accustomed to concepts like Protocol Buffers.
- Not applicable to browsers directly (but there's a library for it).
- Hard to debug, as packets are transferred in binary format, unless you're a robot.
- Works well with data types that can fit snugly in a few megabytes, for image data,
  or video data it does not fit well.

## Message Brokers - Asynchronous Communication

### What is Message Broker?

Message Broker is an intermediate middleware component, that does the job of
forwarding messages from one sender, through protocol A to one receiver's
protocol B. Essentially, it's an architectural pattern that allows the process
of validating, transferring, and routing messages, allowing for inter-service
communication without them having to be aware of each other's existence.

### Basic Concepts of a Message Broker

- **Producer**: The one who produces the message. Or in a pub-sub pattern, also
  called the **publisher**.
- **Consumer**: The one who consumes and processes the message, also called in
  the **subscriber**.
- **Queue**: the place where messages are put, before getting consumed.
- **Topic**: The address or the subscription name for messages to be sent to.
  Similar to RSS feeds, messages sent to a topic are sent to all subscribers for
  that topic.
- **Exchange** (RabbitMQ specific): RabbitMQ doesn't directly send messages
  like an endpoint-to-endpoint system, but messages are sent to an exchange, which
  conducts how messages are routed and where to. There are multiple types of
  Exchanges as below in the RabbitMQ section.
- **Partition** (Kafka specific): A way to split a topic into multiple streams
  to scale. Similar to a highway split into multiple lanes, where each lane has
  their own stop for processing the messages. The order of messages is only
  guaranteed inside a lane, not across multiple lanes.

## RabbitMQ

### RabbitMQ Overview

RabbitMQ is a free and open source software, enterprise-grade for sending messages
and streaming data, easily deployable in various cloud environments. Example use
cases that RabbitMQ would be a good option for:

- **Decoupling services**: For example, if you want to send a notification across
  multiple channels, such as an email or a push notification, when needed, the
  backend just needs to publish a message, and both services (Email Sender Service
  and Push Notification Service) listening can catch and execute their job.
- **Streaming**: When the user uploads a video, there are multiple things that
  need running, so a message published to all can enable post-upload analysis
  service, compression service, transcoding service and push notifications to all
  subscribers, etc. Each service can independently listen and execute the tasks.

### RabbitMQ Architecture

Each message in RabbitMQ has a set of _headers_ and a _routing_ key, as metadata
for where to send to. Imagine what you have to put on an envelope to get it
delivered. Make it factual, and as detailed as possible to be able to send to
the one you actually want: "Mr. Harry Potter, Under the Stairs, 4th Privet Drive".

RabbitMQ works as a broker, but uses exchanges instead. Exchanges are like
buffers for publishers to send messages into, and then RabbitMQ would calculate
where to send the messages to. RabbitMQ has many types of exchanges, which mainly
just conduct the way that messages are delivered:

- **Fanout**: Indiscriminately send to all queues, streams or other exchanges
  bound to this exchange. This completely ignores the message's routing key.
  Similar to a broadcast in networking.
- **Topic**: Uses the message's routing key to decide which destination to
  send to, similar to an address. Usually keys have a format like `regions.na.cities`,
  and routing keys can use wildcard patterns (`*`, `#`) to match where to send.
  The special key `#` acts as a fanout, sending to everything. Similar to a
  multicast in networking.
- **Direct**: Binds to one or many queues. The routing key `abc` will only match
  those with the exact same key.
- **Header**: Instead of using routing keys, use the message's headers instead.
  This is more useful when destinations can't be specified cleanly in a key, but
  needs many things among headers. Header destinations can match _any_ (as long as
  one header matches, it will receive) or _all_ (all headers need to match).

### RabbitMQ Patterns

#### Work Queues

The main idea behind Work Queues (or Task Queues) is to avoid doing a
resource-intensive task immediately and having to wait for it to complete. So we
schedule the task to be run later. We send the task as a message to a queue,
and a worker service will pop the tasks and sequentially execute the job. And
we can have as many workers as we need it to.

#### Pub-Sub

The main idea is to have one producer produces a message that is delivered to
multiple consumers. Different from work queues, each consumer can consume the
message independently of each other, and one consuming doesn't mean the others
can no longer do so.

Similar to using Fluentbit or Logstash, we can have them as publishers to publish
log messages to various channels, either to the console, or to a dedicated system
like OpenSearch or Elasticsearch.

#### Routing

RabbitMQ has a binding, which is a relationship between an exchange and a queue.
When we bind a queue to an exchange, this just means that the queue is interested
in messages from the exchange. The _routing key_ now depends on the exchange type.

**Direct Exchange**: A simple pub-sub system above broadcasts messages to all
consumers. But now we want to filter it based on something else.

#### RPC

What if we need to run a function and wait for the result? The pattern is commonly
known as Remote Procedure Call or RPC.

> RPC has a bad reputation that it is difficult, unpredictable, and complex to
> debug if the programmer doesn't know if the function is local or it's a slow
> RPC. Misused RPC can result in unmaintainable spaghetti code.
>
> When in doubt, avoid RPC.

How an RPC works:

1. When the client starts, it creates a callback queue exclusive to it.
2. For each RPC, the Client sends a message with `reply_to` (to the callback
   queue) and `correlation_id` (unique for each request).
3. Request is sent to an `rpc_queue` queue.
4. The RPC worker sees that request, does the job and sends a message back to
   the callback queue with the `correlation_id`.
5. The client checks the `correlation_id`. If it matches, returns the response
   to the application.

### Merits and Demerits

Merits:

- Independent of platforms. Similar to gRPC, a lot of languages and platforms
  are supported.
- Because of protocols like AMQP, even if the broker is down, the messages are
  still persistent, so when it's up again, messages can continue be delivered.
- Brokers have high availability, so a few of them making up a clustering would
  be amazing for scaling architecture.
- You can believe that the broker will deliver the message properly to the final
  destination.
- Messages can be routed via multiple rules, routing keys flexibly.

Demerits:

- To try to implement the protocol itself (AMQP) would be very daunting, difficult,
  and need time to brush up on. Not just ask GPT and you now put "Kafka" on your
  resume.
- The Protocol itself is heavy, using TCP/IP, it takes time to get carried over.
- There's a limit to a message's size. Messages are not meant to hold much data,
  so sending large files would need splitting into multiple messages, which is not
  optimal.

## Apache Kafka

### Kafka Overview

Apache Kafka is an open-source distributed event streaming platform used by
companies for high-performance data pipelines and streaming analytics. It is meant
for the online world, where the user of software is more software. It specializes
in scalability, availability and fault-tolerance. If any of its servers fails,
the other servers will take over the work, even in the case of network problems or
machine failures.

It is trusted by thousands of organizations, because it supports mission-critical
use cases that guarantee ordering, no message loss, and very efficient processing.

### Components of Kafka

- **Producers**: client applications that write events to Kafka.
- **Consumers**: read and process those events.
- **Topic**: similar to a folder in a filesystem, and the events are the files
  in that folder. Topics in Kafka are always multi-producer and multi-consumer.
- **Partition**: A topic is spread over a number of buckets located on different
  Kafka brokers. It allows clients to both read and write the data from and to many
  brokers at the same time. When a new event is published, it is written to one of
  the topic's partitions.
- **Broker**: A single machine or a single moving part that allows information to
  go between producers and consumers.
- **ZooKeeper/KRaft**: KRaft, based on Raft, relies on a quorom of participants
  to vote on the validity of a transaction or a batch of data. It defines the
  **replicated state machine**, where a number of node process the same sequence
  of inputs in the same order to stay synchronized.

More on KRaft:

- Each node is in one of three states: follower, candidate, leader.
- The Election: All nodes start as followers, only one can be the leader at any
  given time. If the leader isn't responsive, the followers assume it died, and
  requests all other nodes to `Vote`. Each follower votes for the first valid
  candidate (with an up-to-date log and valid term).
- Log Replication: The leader is responsible for replicating all changes
  across the cluster, stored in a list called the _log_, then ask all followers to
  append entries to their logs. Once a majority has replicated the entry, the
  leader commits, and notifies the followers that the entry is committed. Then all
  followers commit also, sync up the system.

## gRPC vs Message Broker

| Features            | gRPC                                           | Message Broker                                |
| ------------------- | ---------------------------------------------- | --------------------------------------------- |
| Communication Model | Sync                                           | Async                                         |
| Coupling            | Moderate                                       | Decoupled                                     |
| Response Time       | Instant                                        | Unknown                                       |
| Latency             | Low                                            | Higher, due to broker's logic and persistence |
| Fault Handling      | Not provided                                   | Built-in, retries, node takeover              |
| Scalability         | Scale by adding servers behind a load balancer | Easily, independently scalable                |

## Kafka vs RabbitMQ

| Feature            | RabbitMQ                            | Kafka                                         |
| ------------------ | ----------------------------------- | --------------------------------------------- |
| Model              | Traditional Queue                   | Distributed Streaming                         |
| Delivery Mechanism | Push to consumers, consume = delete | Consumers pull messages                       |
| Throughput         | Moderate                            | Very High                                     |
| Latency            | Low, time-sensitive                 | Moderate (batching and disk persistence)      |
| Retention          | Short-term                          | Long-term to disk                             |
| Ordering           | Guaranteed per queue                | Guaranteed per partition                      |
| Use Case           | RPC, Work Queues                    | Event Sourcing, Log, Real-time Data Pipelines |
| Complexity         | Lower                               | Higher                                        |
| Replay             | Consume = Gone                      | Can replay at any time if saved               |

## When to use what?

When to use gRPC?

- Durect, Real-time calls.
- Needs for synchronous communication.
- Low latency inter-service communication.

When to use RabbitMQ?

- Asynchronous Message Queuing
- Simple delivery and routing that is guaranteed

When to use Kafka?

- High-Throughput Event Streaming
- Large distributed system
- Replayability for events

## Hybrid Architecture

### What is it?

It is a system design approach that combines two or more distinct architectural
styles to get the benefits of both.

It means intentionally using synchronous communication (like gRPC) for real-time
commands and asynchronous communication (like Message Brokers) for decoupled
events and background tasks in a microservices architecture.

### Design Principles

- Utilize Domain-Driven Design to define where the services' boundaries lie,
  what should be designed as a command (gRPC) and what should be designed as an
  event (broker).
- Avoid shared services to allow for each service to not depend on each other.
  Use events to synchronize data between them.
- Don't couple temporally. Services communication asynchronously do not need to
  wait for others to respond, or to be online to continue its work.
- Design for data consistency and atomicity: each event should only be treated
  as a single, atomic transaction.

## References

I forgot to list some lol hihi have a kero for fun:

![kero](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%2Fid%2FOIP.2voCm1Uc_K94nUl-DUQ2WgHaGq%3Fcb%3Ducfimg2%26pid%3DApi%26ucfimg%3D1&f=1&ipt=35ee29ee71e1664c11ee079cd5ba13633200fefa42f80adc9548e7bb9717073f)

- [Kafka](https://kafka.apache.org/documentation/#gettingStarted)
- [KRaft](https://developer.mamezou-tech.com/blogs/2024/01/22/kraft-kafka-without-zk/)
- [A deep dive into KRaft Protocol](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%2Fid%2FOIP.2voCm1Uc_K94nUl-DUQ2WgHaGq%3Fcb%3Ducfimg2%26pid%3DApi%26ucfimg%3D1&f=1&ipt=35ee29ee71e1664c11ee079cd5ba13633200fefa42f80adc9548e7bb9717073f)

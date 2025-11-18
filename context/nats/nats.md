### Install nats.go Client

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Instructions for installing the nats.go client using the Go package manager. It covers fetching the latest release or a specific version, and notes the NATS Server version compatibility.

```bash
# To get the latest released Go client:
go get github.com/nats-io/nats.go@latest

# To get a specific version:
go get github.com/nats-io/nats.go@v1.44.0

# Note that the latest major version for NATS Server is v2:
go get github.com/nats-io/nats-server/v2@latest
```

---

### NATS Client Request Fanout Example

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Illustrates how a NATS client can send a request to a subject and asynchronously receive responses from multiple services. This example shows subscribing to a reply subject and processing incoming messages within a time limit.

```go
import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// Assume nc is an initialized *nats.Conn

func sendFanoutRequest(nc *nats.Conn) {
	sub, _ := nc.SubscribeSync("rply")
	defer sub.Unsubscribe()

	// Publish request to the subject that services are listening on
	_ = nc.PublishRequest("svc.echo", "rply", nil)

	// Wait for responses for up to 5 seconds
	for start := time.Now(); time.Since(start) < 5*time.Second; {
		msg, err := sub.NextMsg(1 * time.Second)
		if err != nil {
			// Handle timeout or other errors
			break
		}
		fmt.Println("Received ", string(msg.Data))
	}
}
```

---

### NATS Microservice Discovery and Status Commands

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Provides examples of using the 'nats' CLI tool to discover and retrieve information about registered microservices. It shows how to ping services to get their IDs, retrieve detailed info, and fetch statistics, using specific subject patterns.

```APIDOC
Service Discovery and Monitoring Subjects:

- `$SRV.<operation>`: Targets all services.
- `$SRV.<operation>.<service_name>`: Targets services by name.
- `$SRV.<operation>.<service_name>.<service_id>`: Targets a specific service instance.

Available Operations:
- PING: Used for service discovery and Round-Trip Time (RTT) calculation.
- INFO: Returns service configuration details (subjects, metadata, etc.).
- STATS: Provides service statistics (started time, endpoint requests, errors).

Examples:

1. Ping all instances of 'EchoService' to discover their IDs:
nats req '$SRV.PING.EchoService' '' --replies=3

   Example Response:
   {"name":"EchoService","id":"x3Yuiq7g7MoxhXdxk7i4K7","version":"1.0.0","metadata":{},"type":"io.nats.micro.v1.ping_response"}

2. Retrieve detailed information for a specific service instance:
nats req '$SRV.INFO.EchoService.x3Yuiq7g7MoxhXdxk7i4K7' '' | jq

   Example Response:
   {
     "name": "EchoService",
     "id": "x3Yuiq7g7MoxhXdxk7i4K7",
     "version": "1.0.0",
     "metadata": {},
     "type": "io.nats.micro.v1.info_response",
     "description": "",
     "endpoints": [
       {
         "name": "default",
         "subject": "svc.echo",
         "queue_group": "q",
         "metadata": null
       }
     ]
   }

3. Get statistics for a specific service instance:
nats req '$SRV.STATS.EchoService.x3Yuiq7g7MoxhXdxk7i4K7' '' | jq

   Example Response:
   {
     "name": "EchoService",
     "id": "x3Yuiq7g7MoxhXdxk7i4K7",
     "version": "1.0.0",
     "metadata": {},
     "type": "io.nats.micro.v1.stats_response",
     "started": "2024-09-24T11:02:55.564771Z",
     "endpoints": [
       {
         "name": "default",
         "subject": "svc.echo",
         "queue_group": "q",
         "num_requests": 0,
         "num_errors": 0,
         "last_error": "",
         "processing_time": 0,
         "average_processing_time": 0
       }
     ]
   }
```

---

### Import NATS Micro Package

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Import the NATS micro package into your Go application to start building microservices.

```go
import "github.com/nats-io/nats.go/micro"
```

---

### Create and Manage KeyValue Bucket

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Shows how to create, retrieve, and delete JetStream KeyValue buckets. Covers basic CRUD operations like putting, getting, updating, and deleting entries within a bucket.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()

// Create a new bucket. Bucket name is required and has to be unique within a JetStream account.
kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})

// Set a value for a given key
// Put will either create or update a value for a given key
kv.Put(ctx, "sue.color", []byte("blue"))

// Get an entry for a given key
// Entry contains key/value, but also metadata (revision, timestamp, etc.))
entry, _ := kv.Get(ctx, "sue.color")

// Prints `sue.color @ 1 -> "blue"`
fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

// Update a value for a given key
// Update will fail if the key does not exist or the revision has changed
kv.Update(ctx, "sue.color", []byte("red"), 1)

// Create will fail if the key already exists
_, err := kv.Create(ctx, "sue.color", []byte("purple"))
fmt.Println(err) // prints `nats: key exists`

// Delete a value for a given key.
// Delete is not destructive, it will add a delete marker for a given key
// and all previous revisions will still be available
kv.Delete(ctx, "sue.color")

// getting a deleted key will return an error
_, err = kv.Get(ctx, "sue.color")
fmt.Println(err) // prints `nats: key not found`

// A bucket can be deleted once it is no longer needed
js.DeleteKeyValue(ctx, "profiles")
```

---

### Watch KeyValue Bucket Changes in Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates creating a watcher for a NATS KeyValue bucket to monitor changes. It shows how to initialize a watcher, process initial values, and handle subsequent updates or delete markers. The example includes using context and configuring watchers with options like `UpdatesOnly`.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})

kv.Put(ctx, "sue.color", []byte("blue"))

// A watcher can be created to watch for changes on a given key or the whole bucket
// By default, watcher will return most recent values for all matching keys.
// Watcher can be configured to only return updates by using jetstream.UpdatesOnly() option.
watcher, _ := kv.Watch(ctx, "sue.*")
defer watcher.Stop()

kv.Put(ctx, "sue.age", []byte("43"))
kv.Put(ctx, "sue.color", []byte("red"))

// First, the watcher sends most recent values for all matching keys.
// In this case, it will send a single entry for `sue.color`.
entry := <-watcher.Updates()
// Prints `sue.color @ 1 -> "blue"`
fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

// After all current values have been sent, watcher will send nil on the channel.
entry = <-watcher.Updates()
if entry != nil {
    fmt.Println("Unexpected entry received")
}

// After that, watcher will send updates when changes occur
// In this case, it will send an entry for `sue.color` and `sue.age`.

entry = <-watcher.Updates()
// Prints `sue.age @ 2 -> "43"`
fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

entry = <-watcher.Updates()
// Prints `sue.color @ 3 -> "red"`
fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))
```

---

### Get NATS KeyValue Bucket Status (Go)

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Shows how to retrieve metadata and statistics for a NATS KeyValue bucket. The code initializes a KeyValue instance and calls `Status` to get information such as the bucket name, number of values, and total size.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})

kv.Put(ctx, "sue.color", []byte("blue"))
kv.Put(ctx, "sue.age", []byte("43"))
kv.Put(ctx, "bucket", []byte("profiles"))

status, _ := kv.Status(ctx)

fmt.Println(status.Bucket()) // prints `profiles`
fmt.Println(status.Values()) // prints `3`
fmt.Println(status.Bytes()) // prints the size of all values in bytes
```

---

### KeyValue Store (KV) Basic Usage

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Provides an interface for using JetStream as a Key-Value store, covering basic operations like putting, getting, and deleting entries.

```APIDOC
KeyValue Store (KV):
  - CreateKeyValue(cfg KeyValueConfig) (KeyValue, error)
    - Creates a new Key-Value bucket or retrieves an existing one.
    - Parameters:
      - cfg: KeyValueConfig struct defining bucket name and storage options.
    - Returns:
      - KeyValue: The KeyValue interface for interacting with the bucket.
      - error: If bucket creation/retrieval fails.

  - KeyValue(bucketName string) (KeyValue, error)
    - Retrieves an existing KeyValue bucket.
    - Parameters:
      - bucketName: The name of the KV bucket.
    - Returns:
      - KeyValue: The KeyValue interface.
      - error: If the bucket is not found or inaccessible.

  - Put(key string, value []byte) (uint64, error)
    - Puts a key-value pair into the bucket.
    - Parameters:
      - key: The key string.
      - value: The value as a byte slice.
    - Returns:
      - uint64: The revision number of the entry.
      - error: If the put operation fails.

  - Get(key string) (KeyValueEntry, error)
    - Retrieves the latest value for a given key.
    - Parameters:
      - key: The key string.
    - Returns:
      - KeyValueEntry: An object containing the value, metadata, and revision.
      - error: If the get operation fails or key not found.

  - Delete(key string) (uint64, error)
    - Deletes a key from the bucket.
    - Parameters:
      - key: The key string.
    - Returns:
      - uint64: The revision number after deletion.
      - error: If the delete operation fails.
```

---

### Get NATS JetStream Bucket Status Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Illustrates how to retrieve the current status of a NATS JetStream Object Store bucket. This includes information like the bucket name and its total size in bytes.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
os, _ := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})

os.PutString(ctx, "config-1", "cfg1")
os.PutString(ctx, "config-2", "cfg1")
os.PutString(ctx, "config-3", "cfg1")

status, _ := os.Status(ctx)

fmt.Println(status.Bucket()) // prints `configs`
fmt.Println(status.Size()) // prints the size of the bucket in bytes
```

---

### Purge KeyValue Bucket in NATS (Go)

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Illustrates how to clear a NATS KeyValue bucket. The example shows using `Purge` to remove all keys, leaving delete markers for the latest revision, and `PurgeDeletes` to remove only keys that have delete markers.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})

kv.Put(ctx, "sue.color", []byte("blue"))
kv.Put(ctx, "sue.age", []byte("43"))
kv.Put(ctx, "bucket", []byte("profiles"))

// Purge will remove all keys from a bucket.
// The latest revision of each key will be kept
// with a delete marker, all previous revisions will be removed
// permanently.
kv.Purge(ctx)

// PurgeDeletes will remove all keys from a bucket
// with a delete marker.
kv.PurgeDeletes(ctx)
```

---

### Create NATS Micro Service and Add Endpoint

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Demonstrates creating a NATS microservice with a name, version, and a basic echo endpoint handler. It also shows how to connect to NATS.

```go
nc, _ := nats.Connect(nats.DefaultURL)

// request handler
echoHandler := func(req micro.Request) {
    req.Respond(req.Data())
}

srv, err := micro.AddService(nc, micro.Config{
    Name:        "EchoService",
    Version:     "1.0.0",
    // base handler
    Endpoint: &micro.EndpointConfig{
        Subject: "svc.echo",
        Handler: micro.HandlerFunc(echoHandler),
    },
})
```

---

### Basic NATS JetStream Usage

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates the fundamental workflow of NATS JetStream, including setting up context, connecting to NATS, creating a JetStream context, managing streams, publishing messages, and consuming messages using fetch, callbacks, and iterators.

```go
package main

import (
    "context"
    "fmt"
    "strconv"
    "time"

    "github.com/nats-io/nats.go"
    "github.com/nats-io/nats.go/jetstream"
)

func main() {
    // In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    nc, _ := nats.Connect(nats.DefaultURL)

    // Create a JetStream management interface
    js, _ := jetstream.New(nc)

    // Create a stream
    s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
        Name:     "ORDERS",
        Subjects: []string{"ORDERS.*"},
    })

    // Publish some messages
    for i := 0; i < 100; i++ {
        js.Publish(ctx, "ORDERS.new", []byte("hello message "+strconv.Itoa(i)))
        fmt.Printf("Published hello message %d\n", i)
    }

    // Create durable consumer
    c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
        Durable:   "CONS",
        AckPolicy: jetstream.AckExplicitPolicy,
    })

    // Get 10 messages from the consumer
    messageCounter := 0
    msgs, err := c.Fetch(10)
    if err != nil {
        // handle error
    }

    for msg := range msgs.Messages() {
        msg.Ack()
        fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
        messageCounter++
    }

    fmt.Printf("received %d messages\n", messageCounter)

    if msgs.Error() != nil {
        fmt.Println("Error during Fetch(): ", msgs.Error())
    }

    // Receive messages continuously in a callback
    cons, _ := c.Consume(func(msg jetstream.Msg) {
        msg.Ack()
        fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
        messageCounter++
    })
    defer cons.Stop()

    // Iterate over messages continuously
    it, _ := c.Messages()
    for i := 0; i < 10; i++ {
        msg, _ := it.Next()
        msg.Ack()
        fmt.Printf("Received a JetStream message via iterator: %s\n", string(msg.Data()))
        messageCounter++
    }
    it.Stop()

    // block until all 100 published messages have been processed
    for messageCounter < 100 {
        time.Sleep(10 * time.Millisecond)
    }
}
```

---

### Basic NATS Go Client Usage

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Demonstrates fundamental operations of the nats.go client, including connecting to a server, publishing messages, subscribing asynchronously and synchronously, handling requests and replies, and managing subscriptions.

```go
import "github.com/nats-io/nats.go"

// Connect to a server
nc, _ := nats.Connect(nats.DefaultURL)

// Simple Publisher
nc.Publish("foo", []byte("Hello World"))

// Simple Async Subscriber
nc.Subscribe("foo", func(m *nats.Msg) {
    fmt.Printf("Received a message: %s\n", string(m.Data))
})

// Responding to a request message
nc.Subscribe("request", func(m *nats.Msg) {
    m.Respond([]byte("answer is 42"))
})

// Simple Sync Subscriber
sub, err := nc.SubscribeSync("foo")
m, err := sub.NextMsg(timeout)

// Channel Subscriber
ch := make(chan *nats.Msg, 64)
sub, err := nc.ChanSubscribe("foo", ch)
msg := <- ch

// Unsubscribe
sub.Unsubscribe()

// Drain
sub.Drain()

// Requests
msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)

// Replies
nc.Subscribe("help", func(m *nats.Msg) {
    nc.Publish(m.Reply, []byte("I can help!"))
})

// Drain connection (Preferred for responders)
// Close() not needed if this is called.
nc.Drain()

// Close connection
nc.Close()
```

---

### Add NATS Microservice with nats.go

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Demonstrates how to connect to a NATS server and register a new microservice instance using the nats.go micro library. It includes setting up service configuration like name, version, and endpoints, and handling potential errors during service addition.

```go
nc, _ := nats.Connect("nats://localhost:4222")
echoHandler := func(req micro.Request) {
    req.Respond(req.Data())
}

config := micro.Config{
    Name:    "EchoService",
    Version: "1.0.0",
    Endpoint: &micro.EndpointConfig{
        Subject: "svc.echo",
        Handler: micro.HandlerFunc(echoHandler),
    },
}
for i := 0; i < 3; i++ {
    srv, err := micro.AddService(nc, config)
    if err != nil {
        log.Fatal(err)
    }
    defer srv.Stop()
}
```

---

### Context Support for NATS Operations

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Illustrates how to use Go's context package with NATS operations. This includes setting timeouts for requests and using context with synchronous subscribers to manage operation lifecycles and cancellations.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

nc, err := nats.Connect(nats.DefaultURL)

// Request with context
msg, err := nc.RequestWithContext(ctx, "foo", []byte("bar"))

// Synchronous subscriber with context
sub, err := nc.SubscribeSync("foo")
msg, err := sub.NextMsgWithContext(ctx)
```

---

### NATS JetStream Basic Management: Streams and Consumers

Source: https://github.com/nats-io/nats.go/blob/main/legacy_jetstream.md

Shows how to manage JetStream resources by connecting to NATS, creating a JetStream context, adding/updating streams, adding/deleting consumers, and deleting streams. Essential for setting up JetStream environments.

```go
import "github.com/nats-io/nats.go"

// Connect to NATS
nc, _ := nats.Connect(nats.DefaultURL)

// Create JetStream Context
js, _ := nc.JetStream()

// Create a Stream
js.AddStream(&nats.StreamConfig{
    Name:     "ORDERS",
    Subjects: []string{"ORDERS.*"},
})

// Update a Stream
js.UpdateStream(&nats.StreamConfig{
    Name:     "ORDERS",
    MaxBytes: 8,
})

// Create a Consumer
js.AddConsumer("ORDERS", &nats.ConsumerConfig{
    Durable: "MONITOR",
})

// Delete Consumer
js.DeleteConsumer("ORDERS", "MONITOR")

// Delete Stream
js.DeleteStream("ORDERS")
```

---

### NATS.go Object Store: Basic CRUD Operations

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates fundamental Create, Read, Update, and Delete (CRUD) operations for NATS.go Object Stores. This includes creating buckets, putting objects using various data types (buffer, string, bytes, file), retrieving objects, and deleting them. It requires a NATS connection and JetStream context.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()

// Create a new bucket. Bucket name is required and has to be unique within a JetStream account.
os, _ := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})

config1 := bytes.NewBufferString("first config")
// Put an object in a bucket. Put expects an object metadata and a reader
// to read the object data from.
os.Put(ctx, jetstream.ObjectMeta{Name: "config-1"}, config1)

// Objects can also be created using various helper methods

// 1. As raw strings
os.PutString(ctx, "config-2", "second config")

// 2. As raw bytes
os.PutBytes(ctx, "config-3", []byte("third config"))

// 3. As a file
os.PutFile(ctx, "config-4.txt")

// Get an object
// Get returns a reader and object info
// Similar to Put, Get can also be used with helper methods
// to retrieve object data as a string, bytes or to save it to a file
object, _ := os.Get(ctx, "config-1")
data, _ := io.ReadAll(object)
info, _ := object.Info()

// Prints `configs.config-1 -> "first config"`
fmt.Printf("%s.%s -> %q\n", info.Bucket, info.Name, string(data))

// Delete an object.
// Delete will remove object data from stream, but object metadata will be kept
// with a delete marker.
os.Delete(ctx, "config-1")

// getting a deleted object will return an error
_, err := os.Get(ctx, "config-1")
fmt.Println(err) // prints `nats: object not found`

// A bucket can be deleted once it is no longer needed
js.DeleteObjectStore(ctx, "configs")
```

---

### NATS JetStream Basic Usage: Publish, Subscribe, Consume

Source: https://github.com/nats-io/nats.go/blob/main/legacy_jetstream.md

Demonstrates connecting to NATS, creating a JetStream context, publishing messages asynchronously and synchronously, subscribing to subjects, and consuming messages using both sync and pull consumers. Covers basic message flow.

```go
import "github.com/nats-io/nats.go"
import "time"
import "fmt"

// Connect to NATS
nc, _ := nats.Connect(nats.DefaultURL)

// Create JetStream Context
js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

// Simple Stream Publisher
js.Publish("ORDERS.scratch", []byte("hello"))

// Simple Async Stream Publisher
for i := 0; i < 500; i++ {
    js.PublishAsync("ORDERS.scratch", []byte("hello"))
}
select {
case <-js.PublishAsyncComplete():
case <-time.After(5 * time.Second):
    fmt.Println("Did not resolve in time")
}

// Simple Async Ephemeral Consumer
js.Subscribe("ORDERS.*", func(m *nats.Msg) {
    fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
})

// Simple Sync Durable Consumer (optional SubOpts at the end)
sub, err := js.SubscribeSync("ORDERS.*", nats.Durable("MONITOR"), nats.MaxDeliver(3))
m, err := sub.NextMsg(timeout)

// Simple Pull Consumer
sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
msgs, err := sub.Fetch(10)

// Unsubscribe
sub.Unsubscribe()

// Drain
sub.Drain()
```

---

### List Keys in NATS KeyValue Bucket (Go)

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates how to retrieve all keys stored within a NATS KeyValue bucket. The code initializes a KeyValue instance and then uses the `ListKeys` method to obtain an iterator for all keys, which are then printed.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: "profiles"})

kv.Put(ctx, "sue.color", []byte("blue"))
kv.Put(ctx, "sue.age", []byte("43"))
kv.Put(ctx, "bucket", []byte("profiles"))

keys, _ := kv.ListKeys(ctx)

// Prints all 3 keys
for key := range keys.Keys() {
    fmt.Println(key)
}
```

---

### NATS JetStream Basic Usage

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Illustrates the basic usage of the JetStream API within nats.go. It shows how to connect to NATS, create a JetStream context, retrieve stream and consumer handles, and consume messages with acknowledgments.

```go
// connect to nats server
nc, _ := nats.Connect(nats.DefaultURL)

// create jetstream context from nats connection
js, _ := jetstream.New(nc)

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
deffer cancel()

// get existing stream handle
stream, _ := js.Stream(ctx, "foo")

// retrieve consumer handle from a stream
cons, _ := stream.Consumer(ctx, "cons")

// consume messages from the consumer in callback
cc, _ := cons.Consume(func(msg jetstream.Msg) {
    fmt.Println("Received jetstream message: ", string(msg.Data()))
    msg.Ack()
})
defer cc.Stop()
```

---

### Nkey Authentication

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Demonstrates how to authenticate with a NATS server using Nkey seeds. It shows loading a seed file and directly providing public key and signature callback.

```go
import (
	"github.com/nats-io/nats.go"
)

// Load Nkey option from a seed file
opt, err := nats.NkeyOptionFromSeed("seed.txt")
// Connect using the Nkey option
// nc, err := nats.Connect(serverUrl, opt)

// Direct authentication using public key and signature callback
// nc, err := nats.Connect(serverUrl, nats.Nkey(pubNkey, sigCB))
```

---

### NATS User Credentials Authentication

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Demonstrates how to authenticate with NATS servers using user credentials files (JWT and NKey). It shows the helper method `UserCredentials` and how to directly set callback handlers for JWT and signature.

```go
nc, err := nats.Connect(url, nats.UserCredentials("user.creds"))
```

```go
nc, err := nats.Connect(url, nats.UserCredentials("user.jwt", "user.nk"))
```

```go
nc, err := nats.Connect(url, nats.UserJWT(jwtCB, sigCB))
```

---

### Connect to NATS Cluster with Options

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Demonstrates connecting to a NATS cluster using various configuration options. This includes setting maximum reconnect attempts, custom reconnect delays, disabling server randomization, and handling disconnect/reconnect/closed events.

```go
var servers = "nats://localhost:1222, nats://localhost:1223, nats://localhost:1224"

nc, err := nats.Connect(servers)

// Optionally set ReconnectWait and MaxReconnect attempts.
// This example means 10 seconds total per backend.
nc, err = nats.Connect(servers, nats.MaxReconnects(5), nats.ReconnectWait(2 * time.Second))

// You can also add some jitter for the reconnection.
// This call will add up to 500 milliseconds for non TLS connections and 2 seconds for TLS connections.
// If not specified, the library defaults to 100 milliseconds and 1 second, respectively.
nc, err = nats.Connect(servers, nats.ReconnectJitter(500*time.Millisecond, 2*time.Second))

// You can also specify a custom reconnect delay handler. If set, the library will invoke it when it has tried
// all URLs in its list. The value returned will be used as the total sleep time, so add your own jitter.
// The library will pass the number of times it went through the whole list.
nc, err = nats.Connect(servers, nats.CustomReconnectDelay(func(attempts int) time.Duration {
    return someBackoffFunction(attempts)
}))

// Optionally disable randomization of the server pool
nc, err = nats.Connect(servers, nats.DontRandomize())

// Setup callbacks to be notified on disconnects, reconnects and connection closed.
nc, err = nats.Connect(servers,
	nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		fmt.Printf("Got disconnected! Reason: %q\n", err)
	}),
	nats.ReconnectHandler(func(nc *nats.Conn) {
		fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
	}),
	nats.ClosedHandler(func(nc *nats.Conn) {
		fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
	})
)
```

---

### Create NATS Services with Custom Queue Groups

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Demonstrates registering multiple NATS services with unique queue group names to facilitate fanout request patterns. Each service instance can be assigned a distinct queue group, allowing a single request to be processed by multiple services.

```go
import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// Assume nc is an initialized *nats.Conn and echoHandler is defined

func setupServices(nc *nats.Conn) {
	for i := 0; i < 5; i++ {
		srv, _ := micro.AddService(nc, micro.Config{
			Name:        "EchoService",
			Version:     "1.0.0",
			QueueGroup:  fmt.Sprintf("q-%d", i),
			// base handler
			Endpoint: &micro.EndpointConfig{
				Subject: "svc.echo",
				Handler: micro.HandlerFunc(echoHandler),
			},
		})
		_ = srv
	}
}
```

---

### List Objects in NATS JetStream Bucket Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Shows how to list all objects stored within a NATS JetStream Object Store bucket. It returns a slice of object information, allowing iteration over object names and other details.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
os, _ := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})

os.PutString(ctx, "config-1", "cfg1")
os.PutString(ctx, "config-2", "cfg1")
os.PutString(ctx, "config-3", "cfg1")

// List will return information about all objects in a bucket
objects, _ := os.List(ctx)

// Prints all 3 objects
for _, object := range objects {
    fmt.Println(object.Name)
}
```

---

### NATS Connection with Authentication

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Shows how to authenticate NATS connections, especially in clustered environments requiring credentials. It covers username/password authentication and token-based authentication, including scenarios where credentials might be specified in URLs or via options.

```go
// When connecting to a mesh of servers with auto-discovery capabilities,
// you may need to provide a username/password or token in order to connect
// to any server in that mesh when authentication is required.
// Instead of providing the credentials in the initial URL, you will use
// new option setters:
nc, err = nats.Connect("nats://localhost:4222", nats.UserInfo("foo", "bar"))

// For token based authentication:
nc, err = nats.Connect("nats://localhost:4222", nats.Token("S3cretT0ken"))

// You can even pass the two at the same time in case one of the server
// in the mesh requires token instead of user name and password.
nc, err = nats.Connect("nats://localhost:4222",
    nats.UserInfo("foo", "bar"),
    nats.Token("S3cretT0ken"))

// Note that if credentials are specified in the initial URLs, they take
// precedence on the credentials specified through the options.
// For instance, in the connect call below, the client library will use
// the user "my" and password "pwd" to connect to localhost:4222, however,
// it will use username "foo" and password "bar" when (re)connecting to
// a different server URL that it got as part of the auto-discovery.
nc, err = nats.Connect("nats://my:pwd@localhost:4222", nats.UserInfo("foo", "bar"))
```

---

### Manage Pull Consumers (JetStream Interface) in NATS Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates CRUD operations for JetStream pull consumers using the `JetStream` interface. Covers creating, creating or updating, updating, fetching, and deleting consumers. Requires a `jetstream.JetStream` instance and context.

```go
js, _ := jetstream.New(nc)

// create a consumer (this is an idempotent operation)
// an error will be returned if consumer already exists and has different configuration.
cons, _ := js.CreateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    Durable: "foo",
    AckPolicy: jetstream.AckExplicitPolicy,
})

// create an ephemeral pull consumer by not providing `Durable`
ephemeral, _ := js.CreateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    AckPolicy: jetstream.AckExplicitPolicy,
})


// consumer can also be created using CreateOrUpdateConsumer
// this method will either create a consumer if it does not exist
// or update existing consumer (if possible)
cons2 := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    Name: "bar",
})

// consumers can be updated
// an error will be returned if consumer with given name does not exist
// or an illegal property is to be updated (e.g. AckPolicy)
updated, _ := js.UpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    AckPolicy: jetstream.AckExplicitPolicy,
    Description: "updated consumer"
})

// get consumer handle
cons, _ = js.Consumer(ctx, "ORDERS", "foo")

// delete a consumer
js.DeleteConsumer(ctx, "ORDERS", "foo")
```

---

### Run Go Tests with Specific Modfile

Source: https://github.com/nats-io/nats.go/blob/main/CONTRIBUTING.md

This command executes all Go tests within the project, explicitly instructing the Go toolchain to use `go_test.mod` for resolving test-specific dependencies. This ensures that tests run against the correct set of dependencies without affecting the main application's `go.mod`.

```Shell
go test ./... -modfile=go_test.mod
```

---

### TLS Connection Options

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Illustrates various ways to establish secure TLS connections with NATS servers. This includes default TLS, handling self-signed certificates with RootCAs, client certificate authentication, and providing a custom tls.Config.

```go
import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"github.com/nats-io/nats.go"
)

// Connect using TLS scheme (verifies server name by default)
// nc, err := nats.Connect("tls://nats.demo.io:4443")

// Connect with self-signed certificates using RootCAs
// nc, err = nats.Connect("tls://localhost:4443", nats.RootCAs("./configs/certs/ca.pem"))

// Connect requiring client certificate
// certOption := nats.ClientCert("./configs/certs/client-cert.pem", "./configs/certs/client-key.pem")
// nc, err = nats.Connect("tls://localhost:4443", certOption)

// Connect with a complete custom tls.Config
// certFile := "./configs/certs/client-cert.pem"
// keyFile := "./configs/certs/client-key.pem"
// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
// pool := x509.NewCertPool()
// caCert, err := ioutil.ReadFile("./configs/certs/ca.pem")
// pool.AppendCertsFromPEM(caCert)
// config := &tls.Config{
// 	ServerName: 	opts.Host,
// 	Certificates: 	[]tls.Certificate{cert},
// 	RootCAs:    	pool,
// 	MinVersion: 	tls.VersionTLS12,
// }
// nc, err = nats.Connect("nats://localhost:4443", nats.Secure(config))

```

---

### Synchronous Publish Message

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates synchronous message publishing using JetStream in nats.go. Covers publishing with a full nats.Msg struct or a simplified subject/data pair. Supports options like message ID and headers.

```go
js, _ := jetstream.New(nc)

// Publish message on subject ORDERS.new
// Given subject has to belong to a stream
ack, err := js.PublishMsg(ctx, &nats.Msg{
    Data:    []byte("hello"),
    Subject: "ORDERS.new",
})
fmt.Printf("Published msg with sequence number %d on stream %q", ack.Sequence, ack.Stream)

// A helper method accepting subject and data as parameters
ack, err = js.Publish(ctx, "ORDERS.new", []byte("hello"))
```

```go
// All 3 implementations are work identically
ack, err := js.PublishMsg(ctx, &nats.Msg{
    Data:    []byte("hello"),
    Subject: "ORDERS.new",
    Header: nats.Header{
        "Nats-Msg-Id": []string{"id"},
    },
})

ack, err = js.PublishMsg(ctx, &nats.Msg{
    Data:    []byte("hello"),
    Subject: "ORDERS.new",
}, jetstream.WithMsgID("id"))

ack, err = js.Publish(ctx, "ORDERS.new", []byte("hello"), jetstream.WithMsgID("id"))
```

---

### Publishing Messages on a Stream

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

APIs for publishing messages to a JetStream stream, supporting both synchronous and asynchronous modes.

```APIDOC
Publishing Messages:
  - Publish(subject string, data []byte, opts ...PublishOption) (*PubAck, error)
    - Publishes a message to a stream synchronously.
    - Waits for acknowledgement from JetStream.
    - Parameters:
      - subject: The subject to publish to.
      - data: The message payload as a byte slice.
      - opts: Optional PublishOption for headers, sequence, etc.
    - Returns:
      - *PubAck: The publish acknowledgement containing sequence information.
      - error: If the publish operation fails.

  - PublishAsync(subject string, data []byte, ackHandler AsyncAckHandler, opts ...PublishOption) error
    - Publishes a message to a stream asynchronously.
    - Provides a callback for acknowledgement.
    - Parameters:
      - subject: The subject to publish to.
      - data: The message payload as a byte slice.
      - ackHandler: A callback function (AsyncAckHandler) to process the acknowledgement or error.
      - opts: Optional PublishOption for headers, sequence, etc.
    - Returns:
      - error: If initiating the asynchronous publish fails.
```

---

### Configure Message Buffer with Options

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Illustrates configuring the `Messages()` iterator with options like `PullMaxMessages` and `PullMaxBytes` to control the client-side buffer size. This allows for pre-buffering a specified number of messages or bytes, improving efficiency for high-throughput scenarios.

```go
// a maximum of 10 messages or 1024 bytes will be stored in memory (whichever is encountered first)
iter, _ := cons.Messages(jetstream.PullMaxMessages(10), jetstream.PullMaxBytes(1024))
```

---

### Go Testing Dependency Management

Source: https://github.com/nats-io/nats.go/blob/main/CONTRIBUTING.md

Commands for managing testing dependencies and running tests using a separate go_test.mod file in the nats.go project. This ensures testing dependencies are isolated from the main project dependencies.

```shell
go mod tidy -modfile=go_test.mod
```

```shell
go test ./... -modfile=go_test.mod
```

---

### Queue Groups

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Demonstrates how to implement message queuing semantics using queue groups. Subscriptions within the same queue group receive messages in a load-balanced fashion, ensuring each message is processed by only one subscriber in the group.

```go
import (
	"github.com/nats-io/nats.go"
)

// Subscribe to a subject with a queue group name
// All subscribers with the same queue name form a group.
// Each message is delivered to only one subscriber per group.
// var received int
// nc.QueueSubscribe("foo", "job_workers", func(_ *nats.Msg) {
//   received += 1
// })

```

---

### Discover NATS Service IDs using nats req

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

This shell command illustrates how to use the `nats req` utility to send a PING request to a service name. It queries all instances of 'EchoService' and returns their unique IDs, along with other basic information, demonstrating service discovery.

```sh
nats req '$SRV.PING.EchoService' '' --replies=3
```

---

### NATS JetStream Stream Listing

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

API methods for listing all available streams and retrieving just their names from the JetStream context.

```APIDOC
// List all streams and their configurations
streams := js.ListStreams(ctx)
for s := range streams.Info() {
    fmt.Println(s.Config.Name)
}
if streams.Err() != nil {
    fmt.Println("Unexpected error occurred listing streams")
}

// List only the names of all streams
names := js.StreamNames(ctx)
for name := range names.Name() {
    fmt.Println(name)
}
if names.Err() != nil {
    fmt.Println("Unexpected error occurred listing stream names")
}
```

---

### Iterate Over Incoming Messages

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates how to use the `Messages()` method to obtain an iterator for incoming JetStream messages. It shows a basic loop for fetching messages, handling potential errors from the iterator, printing message data, and acknowledging received messages.

```go
iter, _ := cons.Messages()
for {
    msg, err := iter.Next()
    // Next can return error, e.g. when iterator is closed or no heartbeats were received
    if err != nil {
        //handle error
    }
    fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
    msg.Ack()
}
iter.Stop()
```

---

### Listing Streams and Stream Names

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Retrieves a list of all available streams or just their names within the JetStream context.

```APIDOC
Listing Streams:
  - Streams() ([]*StreamInfo, error)
    - Retrieves information for all streams.
    - Returns:
      - []*StreamInfo: A slice of StreamInfo structs for all streams.
      - error: If listing streams fails.

  - StreamNames() ([]string, error)
    - Retrieves only the names of all available streams.
    - Returns:
      - []string: A slice of stream names.
      - error: If retrieving stream names fails.
```

---

### KeyValue Store (KV) Watching for Changes

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Allows clients to watch for changes (updates, deletes) on specific keys or prefixes within a Key-Value bucket.

```APIDOC
KV Watching:
  - Watch(key string, opts ...WatchOption) (<-chan KeyValueEntry, error)
    - Creates a watch for changes on a specific key.
    - Returns a channel that yields entries when the key is updated or deleted.
    - Parameters:
      - key: The key to watch.
      - opts: Optional WatchOption to specify starting revision or filter.
    - Returns:
      - <-chan KeyValueEntry: A channel for receiving updates.
      - error: If setting up the watch fails.

  - WatchPrefix(prefix string, opts ...WatchOption) (<-chan KeyValueEntry, error)
    - Creates a watch for changes on all keys with a given prefix.
    - Parameters:
      - prefix: The prefix string.
      - opts: Optional WatchOption.
    - Returns:
      - <-chan KeyValueEntry: A channel for receiving updates.
      - error: If setting up the watch fails.
```

---

### Manage Consumers (Stream Interface) in NATS Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Illustrates how to manage JetStream consumers through the `Stream` interface. This includes creating, fetching, and deleting consumers associated with a specific stream. Requires a `jetstream.Stream` handle.

```go
// Create a JetStream management interface
js, _ := jetstream.New(nc)

// get stream handle
stream, _ := js.Stream(ctx, "ORDERS")

// create consumer
cons, _ := stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
    Durable:   "foo",
    AckPolicy: jetstream.AckExplicitPolicy,
})

// get consumer handle
cons, _ = stream.Consumer(ctx, "ORDERS", "foo")

// delete a consumer
stream.DeleteConsumer(ctx, "foo")
```

---

### Object Store Basic Usage

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Provides an interface for using JetStream as an Object Store for storing and retrieving larger data blobs.

```APIDOC
Object Store (ObjectStore):
  - CreateObjectStore(cfg ObjectStoreConfig) (ObjectStore, error)
    - Creates a new Object Store bucket or retrieves an existing one.
    - Parameters:
      - cfg: ObjectStoreConfig struct defining bucket name and storage options.
    - Returns:
      - ObjectStore: The ObjectStore interface for interacting with the bucket.
      - error: If bucket creation/retrieval fails.

  - ObjectStore(bucketName string) (ObjectStore, error)
    - Retrieves an existing Object Store bucket.
    - Parameters:
      - bucketName: The name of the Object Store bucket.
    - Returns:
      - ObjectStore: The ObjectStore interface.
      - error: If the bucket is not found or inaccessible.

  - Put(info ObjectInfo, data io.Reader) (string, error)
    - Puts an object into the store.
    - Parameters:
      - info: ObjectInfo struct containing metadata like name, type, etc.
      - data: An io.Reader for the object's content.
    - Returns:
      - string: The unique ID of the stored object.
      - error: If the put operation fails.

  - Get(name string) (ObjectStoreEntry, error)
    - Retrieves an object by its name.
    - Parameters:
      - name: The name of the object.
    - Returns:
      - ObjectStoreEntry: An object containing metadata and an io.Reader for the data.
      - error: If the get operation fails or object not found.

  - Delete(name string) error
    - Deletes an object from the store.
    - Parameters:
      - name: The name of the object to delete.
    - Returns:
      - error: If the delete operation fails.
```

---

### Fetch Single Messages for Work Queues

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Shows how to use `Messages()` with `jetstream.PullMaxMessages(1)` to fetch messages one by one. This pattern is suitable for implementing work queues where messages are processed individually by multiple workers, preventing redeliveries when processing is slow.

```go
// PullMaxMessages determines how many messages will be sent to the client in a single pull request
iter, _ := cons.Messages(jetstream.PullMaxMessages(1))
numWorkers := 5
sem := make(chan struct{}, numWorkers)
for {
    sem <- struct{}{}
    go func() {
        defer func() {
            <-sem
        }()
        msg, err := iter.Next()
        if err != nil {
            // handle err
        }
        fmt.Printf("Processing msg: %s\n", string(msg.Data()))
        doWork()
        msg.Ack()
    }()
}
```

---

### Listing Consumers and Consumer Names

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Fetches lists of consumers associated with a stream, either full details or just their names.

```APIDOC
Listing Consumers:
  - Consumers(streamName string) ([]ConsumerInfo, error)
    - Retrieves information for all consumers of a specific stream.
    - Parameters:
      - streamName: The name of the stream.
    - Returns:
      - []ConsumerInfo: A slice of ConsumerInfo structs.
      - error: If listing consumers fails.

  - ConsumerNames(streamName string) ([]string, error)
    - Retrieves only the names of consumers for a specific stream.
    - Parameters:
      - streamName: The name of the stream.
    - Returns:
      - []string: A slice of consumer names.
      - error: If retrieving consumer names fails.
```

---

### List JetStream Consumers and Names in NATS Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Provides methods to list all consumers associated with a stream and to retrieve just their names. Handles potential errors during the listing process. Requires a `jetstream.Stream` interface.

```go
// list consumers
consumers := s.ListConsumers(ctx)
for cons := range consumers.Info() {
    fmt.Println(cons.Name)
}
if consumers.Err() != nil {
    fmt.Println("Unexpected error occurred")
}

// list consumer names
names := s.ConsumerNames(ctx)
for name := range names.Name() {
    fmt.Println(name)
}
if names.Err() != nil {
    fmt.Println("Unexpected error occurred")
}
```

---

### Consumer Management

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Handles the creation, retrieval, updating, and deletion of JetStream consumers for specific streams.

```APIDOC
Consumer Management:
  - CreateConsumer(streamName string, cfg ConsumerConfig) (Consumer, error)
    - Creates a new consumer for a given stream.
    - Parameters:
      - streamName: The name of the stream to attach the consumer to.
      - cfg: ConsumerConfig struct defining consumer name, delivery policy, etc.
    - Returns:
      - Consumer: The created Consumer interface.
      - error: If consumer creation fails.

  - ConsumerInfo(streamName string, consumerName string) (ConsumerInfo, error)
    - Retrieves information about a specific consumer.
    - Parameters:
      - streamName: The name of the stream.
      - consumerName: The name of the consumer.
    - Returns:
      - ConsumerInfo: Details of the consumer.
      - error: If consumer information retrieval fails.

  - UpdateConsumer(streamName string, cfg ConsumerConfig) (Consumer, error)
    - Updates an existing consumer.
    - Parameters:
      - streamName: The name of the stream.
      - cfg: ConsumerConfig struct with updated consumer settings.
    - Returns:
      - Consumer: The updated Consumer interface.
      - error: If consumer update fails.

  - DeleteConsumer(streamName string, consumerName string) error
    - Deletes a specified consumer.
    - Parameters:
      - streamName: The name of the stream.
      - consumerName: The name of the consumer to delete.
    - Returns:
      - error: If consumer deletion fails.
```

---

### Fetch Consumer Information in NATS Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Retrieves information about a JetStream consumer. Supports fetching the latest data from the server or using cached information without an API call. Requires a `jetstream.Consumer` interface.

```go
// Fetches latest consumer info from server
info, _ := cons.Info(ctx)
fmt.Println(info.Config.Durable)

// Returns the most recently fetched ConsumerInfo, without making an API call to the server
cachedInfo := cons.CachedInfo()
fmt.Println(cachedInfo.Config.Durable)
```

---

### Receive Messages with Push Consumers

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Details the process of receiving messages using a push consumer via the `Consume()` method. It highlights the need for explicit acknowledgments (`AckPolicy: jetstream.AckExplicitPolicy`), setting `IdleHeartbeat` for push consumers, and optionally filtering messages with `FilterSubject`. It also mentions the `ConsumeErrHandler` for custom error management.

```go
cons, _ := js.CreateOrUpdatePushConsumer("ORDERS", jetstream.ConsumerConfig{
    DeliverSubject: nats.NewInbox()
    AckPolicy: jetstream.AckExplicitPolicy,
    // receive messages from ORDERS.A subject only
    FilterSubject: "ORDERS.A",
    // unlike pull consumers, idle heartbeat is configured on the consumer level
    IdleHeartbeat: 30 * time.Second
})

consContext, _ := c.Consume(func(msg jetstream.Msg) {
    fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
    // messages are not acknowledged automatically
    msg.Ack()
})
defer consContext.Stop()
```

---

### Fetch messages using Fetch() in Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Fetches a specified number of messages from a JetStream pull consumer. The Fetch() method waits up to 30 seconds by default for messages, which can be configured. It returns a batch of messages and an error if any occurs during the fetch operation.

```go
import (
    "fmt"
    "github.com/nats-io/nats.go"
)

// Assuming 'c' is an initialized nats.Consumer interface
// msgs, err := c.Fetch(10)
// if err != nil {
//     // handle error
// }

// for msg := range msgs.Messages() {
//     fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
// }

// if msgs.Error() != nil {
//     // handle error
// }
```

---

### Receiving Messages from Push Consumers

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Methods for receiving messages from a push-based JetStream consumer, typically using a callback function.

```APIDOC
Push Consumer Message Reception:
  - Consume(handler MsgHandler, opts ...ConsumeOption) error
    - Registers a message handler callback function to process messages pushed by JetStream.
    - Messages are delivered asynchronously to the handler.
    - Parameters:
      - handler: A function of type MsgHandler (func(Msg) error) that processes each message.
      - opts: Optional ConsumeOption to configure batching, concurrency, etc.
    - Returns:
      - error: If starting the consumption process fails.

  - Channel(opts ...ConsumeOption) (<-chan Msg, error)
    - Returns a channel that receives messages pushed by JetStream.
    - This allows iterating over incoming messages in a loop.
    - Parameters:
      - opts: Optional ConsumeOption to configure batching, concurrency, etc.
    - Returns:
      - <-chan Msg: A read-only channel for receiving messages.
      - error: If setting up the message channel fails.
```

---

### Asynchronous Publish Message

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates asynchronous message publishing using JetStream in nats.go. Publishing does not wait for acknowledgment, returning a future. The future can be used to check for success or errors.

```go
js, _ := jetstream.New(nc)

// publish message and do not wait for ack
ackF, err := js.PublishMsgAsync(ctx, &nats.Msg{
    Data:    []byte("hello"),
    Subject: "ORDERS.new",
})

// block and wait for ack
select {
case ack := <-ackF.Ok():
    fmt.Printf("Published msg with sequence number %d on stream %q", ack.Sequence, ack.Stream)
case err := <-ackF.Err():
    fmt.Println(err)
}

// similarly to synchronous publish, there is a helper method accepting subject and data
ackF, err = js.PublishAsync("ORDERS.new", []byte("hello"))
```

---

### Consume messages with callback using Consume() in Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Enables push-like message receiving by providing a callback function to the Consume() method. This method performs optimizations like pre-buffering and allows for finer control over fetching single messages on demand, making it suitable for work-queue scenarios. It returns a ConsumeContext that must be stopped to avoid goroutine leaks.

```go
import (
    "fmt"
    "time"
    "github.com/nats-io/nats.go"
)

// Assuming 'js' is an initialized nats.Conn and 'c' is a nats.Consumer interface

// Example of creating a consumer with subject filtering:
// cons, _ := js.CreateOrUpdateConsumer("ORDERS", jetstream.ConsumerConfig{
//     AckPolicy:     jetstream.AckExplicitPolicy,
//     FilterSubject: "ORDERS.A"
// })

// Example of consuming messages with a callback:
// consContext, _ := c.Consume(func(msg jetstream.Msg) {
//     fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
//     // messages are not acknowledged automatically, must be acked manually
//     msg.Ack()
// })
// defer consContext.Stop() // Important: Stop the context to avoid leaks

// Consume() can also be supplied with options like:
// - PullMaxMessages(int): Max messages to buffer.
// - PullMaxBytes(int): Max bytes to buffer (mutually exclusive with PullMaxMessages).
// - PullExpiry(time.Duration): Timeout for a single pull request.
// - PullThresholdMessages(int): Amount of messages triggering buffer refill.
// - PullThresholdBytes(int): Amount of bytes triggering buffer refill.
// - PullHeartbeat(time.Duration): Idle heartbeat duration.
// - ConsumeErrHandler(func (ConsumeContext, error)): Custom error handler.
// - PullMaxMessagesWithBytesLimit: Advanced option for buffering and limiting fetch size.
```

---

### Wildcard Subscriptions

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Explains how to use wildcard characters in NATS subjects for flexible message routing. It covers the '\*' wildcard for matching single tokens and the '>' wildcard for matching the tail of a subject.

```go
import (
	"fmt"
	"github.com/nats-io/nats.go"
)

// Subscribe to subjects with a wildcard in the middle
// nc.Subscribe("foo.*.baz", func(m *nats.Msg) {
//     fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data))
// })

// Subscribe to subjects with a wildcard at the end
// nc.Subscribe("foo.bar.*", func(m *nats.Msg) {
//     fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data))
// })

// Subscribe to subjects with a wildcard matching the tail
// nc.Subscribe("foo.>", func(m *nats.Msg) {
//     fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data))
// })

// Publish a message that matches multiple wildcard subscriptions
// nc.Publish("foo.bar.baz", []byte("Hello World"))
```

---

### Update Go Test Dependencies

Source: https://github.com/nats-io/nats.go/blob/main/CONTRIBUTING.md

This command updates the testing dependencies for the Go project, ensuring that changes are managed within `go_test.mod` rather than altering the main `go.mod` file. This practice helps maintain a clean and stable primary dependency list.

```Shell
go mod tidy -modfile=go_test.mod
```

---

### Stream Management (CRUD)

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Operations for creating, retrieving, updating, and deleting JetStream streams. This includes defining stream configurations and managing their lifecycle.

```APIDOC
Stream Management (CRUD):
  - CreateStream(cfg StreamConfig) (*StreamInfo, error)
    - Creates a new stream with the specified configuration.
    - Parameters:
      - cfg: StreamConfig struct defining stream name, retention policy, limits, etc.
    - Returns:
      - StreamInfo: Information about the created stream.
      - error: If stream creation fails.

  - StreamInfo(streamName string) (*StreamInfo, error)
    - Retrieves information about a specific stream.
    - Parameters:
      - streamName: The name of the stream.
    - Returns:
      - StreamInfo: Details of the stream.
      - error: If stream information retrieval fails.

  - UpdateStream(cfg StreamConfig) (*StreamInfo, error)
    - Updates an existing stream with a new configuration.
    - Parameters:
      - cfg: StreamConfig struct with updated stream settings.
    - Returns:
      - StreamInfo: Information about the updated stream.
      - error: If stream update fails.

  - DeleteStream(streamName string) error
    - Deletes a specified stream.
    - Parameters:
      - streamName: The name of the stream to delete.
    - Returns:
      - error: If stream deletion fails.

  - PurgeStream(streamName string, opts ...PurgeOption) (*StreamInfo, error)
    - Purges messages from a stream based on criteria.
    - Parameters:
      - streamName: The name of the stream to purge.
      - opts: Optional PurgeOption to specify criteria (e.g., sequence, timestamp).
    - Returns:
      - StreamInfo: Information about the stream after purging.
      - error: If purging fails.
```

---

### NATS.go Object Store: Watching for Changes

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Explains how to use Watchers to monitor changes in NATS.go Object Stores. Watchers receive notifications for object updates and deletes, providing metadata but not the full object data. Configuration options like `UpdatesOnly` and `IncludeHistory` are available.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
os, _ := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})

os.PutString(ctx, "config-1", "first config")

// By default, watcher will return most recent values for all objects in a bucket.
// Watcher can be configured to only return updates by using jetstream.UpdatesOnly() option.
watcher, _ := os.Watch(ctx)
defer watcher.Stop()

// create a second object
os.PutString(ctx, "config-2", "second config")

// update metadata of the first object
os.UpdateMeta(ctx, "config-1", jetstream.ObjectMeta{Name: "config-1", Description: "updated config"})

// First, the watcher sends most recent values for all matching objects.
// In this case, it will send a single entry for `config-1`.
object := <-watcher.Updates()
// Prints `configs.config-1 -> ""`
fmt.Printf("%s.%s -> %q\n", object.Bucket, object.Name, object.Description)

// After all current values have been sent, watcher will send nil on the channel.
object = <-watcher.Updates()
if object != nil {
    fmt.Println("Unexpected object received")
}

// After that, watcher will send updates when changes occur
// In this case, it will send an entry for `config-2` and `config-1`.
object = <-watcher.Updates()
// Prints `configs.config-2 -> ""`
fmt.Printf("%s.%s -> %q\n", object.Bucket, object.Name, object.Description)

object = <-watcher.Updates()
// Prints `configs.config-1 -> "updated config"`
fmt.Printf("%s.%s -> %q\n", object.Bucket, object.Name, object.Description)
```

---

### Object Store Watching for Changes

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Enables clients to monitor changes within an Object Store bucket, such as new objects being added or existing ones updated.

```APIDOC
ObjectStore Watching:
  - Watch(opts ...WatchOption) (<-chan ObjectStoreEntry, error)
    - Creates a watch for changes on all objects within the store.
    - Returns a channel that yields entries when objects are added or updated.
    - Parameters:
      - opts: Optional WatchOption to specify starting revision or filter.
    - Returns:
      - <-chan ObjectStoreEntry: A channel for receiving updates.
      - error: If setting up the watch fails.
```

---

### JetStream Client Interfaces

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

The jetstream package in nats.go provides several core interfaces for interacting with NATS JetStream. These interfaces abstract common operations for managing JetStream resources and consuming messages.

```APIDOC
JetStream Interface:
  - Top-level interface for creating and managing streams, consumers, and publishing messages.

Stream Interface:
  - Manages consumers for a specific stream.
  - Performs stream-specific operations like purging, fetching, and deleting messages by sequence number.
  - Fetches stream information.

Consumer Interface:
  - Retrieves information about a consumer.
  - Consumes messages from a stream.

Msg Interface:
  - Provides methods for message-specific operations.
  - Reading data, headers, and metadata.
  - Performing various types of acknowledgements.
```

---

### Retrieve Specific NATS Service Info using nats req

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

This command demonstrates how to fetch detailed configuration information for a specific NATS microservice instance. By targeting the service name and its unique ID with an INFO request, it retrieves metadata, endpoint configurations, and other service details.

```sh
nats req '$SRV.INFO.EchoService.x3Yuiq7g7MoxhXdxk7i4K7' '' | jq
```

---

### NATS JetStream Pull Consumer Message Iterator Options

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Provides a comprehensive list and explanation of configuration options available for the `Messages()` iterator used with NATS JetStream pull consumers. These options control buffering, fetching behavior, timeouts, and error handling for message retrieval.

```APIDOC
Messages() Options:

- PullMaxMessages(int): Specifies the maximum number of messages to buffer in the client's memory. This setting controls how many messages are prefetched from the server.

- PullMaxBytes(int): Sets the maximum total size in bytes for messages to be buffered in the client's memory. This and `PullMaxMessages` are mutually exclusive. The value should be sufficient for the largest expected message plus overhead.

- PullExpiry(time.Duration): Configures the timeout duration for a single pull request to the JetStream server. If the server does not respond within this duration, an error may occur.

- PullThresholdMessages(int): Defines the number of messages that, when reached in the buffer, triggers a new pull request to refill the buffer.

- PullThresholdBytes(int): Defines the total byte size that, when reached in the buffer, triggers a new pull request to refill the buffer.

- PullHeartbeat(time.Duration): Sets the idle heartbeat duration for a single pull request. If the client misses two heartbeats (by default), an error is triggered, unless `WithMessagesErrOnMissingHeartbeat(false)` is used.

- PullMaxMessagesWithBytesLimit: An advanced option that buffers a specified number of messages and limits the size of a single fetch request. Use with caution, ensuring the byte limit is not lower than the maximum expected message size to prevent consumer stalling.
```

---

### Receiving Messages from Pull Consumers

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Methods for fetching messages from a pull-based JetStream consumer, supporting single fetches and continuous polling.

```APIDOC
Pull Consumer Message Retrieval:
  - Fetch(count int, opts ...PullOption) ([]Msg, error)
    - Fetches a specified number of messages from the consumer.
    - Parameters:
      - count: The number of messages to fetch.
      - opts: Optional PullOption to set timeouts or other fetch parameters.
    - Returns:
      - []Msg: A slice of fetched messages.
      - error: If fetching fails.

  - Messages(opts ...PullOption) (<-chan Msg, error)
    - Returns a channel that yields messages as they become available.
    - This method supports continuous polling.
    - Parameters:
      - opts: Optional PullOption to configure the polling behavior.
    - Returns:
      - <-chan Msg: A read-only channel for receiving messages.
      - error: If setting up the message channel fails.

  - FetchNext(opts ...PullOption) (Msg, error)
    - Fetches a single message from the consumer.
    - Parameters:
      - opts: Optional PullOption to set timeouts.
    - Returns:
      - Msg: The fetched message.
      - error: If fetching fails.
```

---

### Create Ordered Consumer in NATS Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Enables the creation of an ordered consumer for strict, deterministic message processing based on storage order. Supports filtering messages by subject. Note: Ordered consumers are not supported for push consumers.

```go
js, _ := jetstream.New(nc)

// create a consumer (this is an idempotent operation)
cons, _ := js.OrderedConsumer(ctx, "ORDERS", jetstream.OrderedConsumerConfig{
    // Filter results from "ORDERS" stream by specific subject
    FilterSubjects: []string{"ORDERS.A"},
})
```

---

### Update Object Metadata in NATS JetStream Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Demonstrates how to update metadata for an object in a NATS JetStream Object Store, including changing its name or adding a description. It requires a JetStream context and an Object Store instance.

```go
js, _ := jetstream.New(nc)
ctx := context.Background()
os, _ := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})

os.PutString(ctx, "config", "data")

// update metadata of the object to e.g. add a description
os.UpdateMeta(ctx, "config", jetstream.ObjectMeta{Name: "config", Description: "this is a config"})

// object can be moved under a new name (unless it already exists)
os.UpdateMeta(ctx, "config", jetstream.ObjectMeta{Name: "config-1", Description: "updated config"})
```

---

### NATS JetStream Stream Info

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Methods to retrieve information about a JetStream stream, either by fetching the latest data from the server or by accessing cached information.

```APIDOC
// Fetch the latest stream information from the server
info, _ := s.Info(ctx)
fmt.Println(info.Config.Name)

// Get the most recently fetched stream information without a server call
cachedInfo := s.CachedInfo()
fmt.Println(cachedInfo.Config.Name)
```

---

### Aggregate Endpoints Using Groups

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Organizes endpoints into groups, where each group shares a common subject prefix. This simplifies managing related endpoints.

```go
srv, _ := micro.AddService(nc, config)

numbersGroup := srv.AddGroup("numbers")

// endpoint will be registered under "numbers.add" subject
_ = numbersGroup.AddEndpoint("add", micro.HandlerFunc(addHandler))
// endpoint will be registered under "numbers.multiply" subject
_ = numbersGroup.AddEndpoint("multiply", micro.HandlerFunc(multiplyHandler))
```

---

### Fetch messages by bytes using FetchBytes() in Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Fetches messages from a JetStream pull consumer, limiting the total data size to a specified number of bytes. Similar to Fetch(), it waits for messages and returns them in batches. This method is useful when message volume is less predictable than message count.

```go
import (
    "fmt"
    "github.com/nats-io/nats.go"
)

// Assuming 'c' is an initialized nats.Consumer interface
// msgs, err := c.FetchBytes(1024) // Fetch up to 1024 bytes of data
// if err != nil {
//     // handle error
// }

// for msg := range msgs.Messages() {
//     fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
// }

// if msgs.Error() != nil {
//     // handle error
// }
```

---

### Fetch messages without waiting using FetchNoWait() in Go

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Fetches messages from a JetStream pull consumer without waiting for the full batch to be available. If not enough messages are ready, it returns immediately with the available messages. This is suitable for scenarios where immediate processing of available data is preferred over waiting for a complete batch.

```go
import (
    "fmt"
    "github.com/nats-io/nats.go"
)

// Assuming 'c' is an initialized nats.Consumer interface
// msgs, err := c.FetchNoWait(10) // Fetch up to 10 messages, without waiting
// if err != nil {
//     // handle error
// }

// for msg := range msgs.Messages() {
//     fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
// }

// if msgs.Error() != nil {
//     // handle error
// }
```

---

### Advanced Connection and Messaging

Source: https://github.com/nats-io/nats.go/blob/main/README.md

Covers advanced NATS client features including automatic retries on connection failure, setting maximum reconnect attempts and wait times, handling reconnect events, flushing messages, setting timeouts for flush operations, auto-unsubscribing after a certain number of messages, and managing multiple NATS connections.

```go
import (
	"fmt"
	"time"
	"github.com/nats-io/nats.go"
)

// Retry on failed connect, set max reconnects and wait time
// nc, err := nats.Connect(nats.DefaultURL,
//     nats.RetryOnFailedConnect(true),
//     nats.MaxReconnects(10),
//     nats.ReconnectWait(time.Second),
//     nats.ReconnectHandler(func(_ *nats.Conn) {
//         // Handle reconnect events
//     }))

// Flush connection to ensure all messages are processed
// nc.Flush()
// fmt.Println("All clear!")

// Flush with a timeout
// err := nc.FlushTimeout(1*time.Second)
// if err != nil {
//     fmt.Println("Flushed timed out!")
// } else {
//     fmt.Println("All clear!")
// }

// Auto-unsubscribe after receiving a specific number of messages
// const MAX_WANTED = 10
// sub, err := nc.Subscribe("foo")
// sub.AutoUnsubscribe(MAX_WANTED)

// Managing multiple connections
// nc1, _ := nats.Connect("nats://host1:4222")
// nc2, _ := nats.Connect("nats://host2:4222")
// nc1.Subscribe("foo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })
// nc2.Publish("foo", []byte("Hello World!"));

```

---

### Override NATS Queue Groups on Groups and Endpoints

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Shows how to set a default queue group for a service and then override it for specific groups and endpoints within that service. This allows for fine-grained control over message distribution, inheriting queue groups from parent configurations.

```go
import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// Assume nc is an initialized *nats.Conn

func overrideQueueGroups(nc *nats.Conn) {
	// Service with default queue group 'q1'
	srv, _ := micro.AddService(nc, micro.Config{
		Name:        "EchoService",
		Version:     "1.0.0",
		QueueGroup:  "q1",
	})

	// Add a group with queue group 'q2' inherited from the service
	g := srv.AddGroup("g", micro.WithGroupQueueGroup("q2"))

	// Add an endpoint to the group, inheriting queue group 'q2'
	g.AddEndpoint("bar", micro.HandlerFunc(func(r micro.Request) { /* handler logic */ }))

	// Add another endpoint to the group, explicitly setting queue group 'q3'
	g.AddEndpoint("bar", micro.HandlerFunc(func(r micro.Request) { /* handler logic */ }), micro.WithEndpointQueueGroup("q3"))
}
```

---

### NATS JetStream Stream Management (CRUD)

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Provides API methods for managing JetStream streams, including creating, updating, retrieving a specific stream by name, and deleting streams.

```APIDOC
js, _ := jetstream.New(nc)

// Create a stream (idempotent operation)
s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
    Name:     "ORDERS",
    Subjects: []string{"ORDERS.*"},
})

// Update an existing stream
s, _ = js.UpdateStream(ctx, jetstream.StreamConfig{
    Name:        "ORDERS",
    Subjects:    []string{"ORDERS.*"},
    Description: "updated stream",
})

// Get a handle to a specific stream
s, _ = js.Stream(ctx, "ORDERS")

// Delete a stream
js.DeleteStream(ctx, "ORDERS")
```

---

### Add Endpoint with Custom Subject

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Adds an endpoint to a NATS microservice, allowing specification of a custom subject for the endpoint using `micro.WithEndpointSubject()`.

```go
// endpoint will be registered under "svc.add" subject
err = srv.AddEndpoint("Adder", micro.HandlerFunc(echoHandler), micro.WithEndpointSubject("svc.add"))
```

---

### Add Endpoint to Existing Service

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Adds a new endpoint to an already created NATS microservice. The endpoint subject can be explicitly defined.

```go
srv, _ := micro.AddService(nc, config)

// endpoint will be registered under "svc.add" subject
err = srv.AddEndpoint("svc.add", micro.HandlerFunc(add))
```

---

### Retrieve NATS Service Statistics using nats req

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

This shell command shows how to obtain real-time performance statistics for a particular NATS microservice instance. It sends a STATS request to a specific service ID, returning metrics such as the number of requests, errors, and processing times for its endpoints.

```sh
nats req '$SRV.STATS.EchoService.x3Yuiq7g7MoxhXdxk7i4K7' '' | jq
```

---

### Disable NATS Queue Groups

Source: https://github.com/nats-io/nats.go/blob/main/micro/README.md

Illustrates how to disable queue group functionality at the service, group, or endpoint level. When queue groups are disabled, NATS creates standard subscriptions instead of using queue groups for message distribution.

```go
import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// Assume nc is an initialized *nats.Conn

func disableQueueGroups(nc *nats.Conn) {
	// Disable queue group for the entire service
	srv, _ := micro.AddService(nc, micro.Config{
		Name:              "EchoService",
		Version:           "1.0.0",
		QueueGroupDisabled: true,
	})

	// Create a group with queue group disabled
	srv.AddGroup("g", micro.WithEndpointQueueGroupDisabled())

	// Create an endpoint with queue group disabled
	srv.AddEndpoint("bar", micro.HandlerFunc(func(r micro.Request) { /* handler logic */ }), micro.WithEndpointQueueGroupDisabled())
}
```

---

### NATS JetStream Stream Message Operations

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

API methods for retrieving and deleting individual messages from a JetStream stream based on sequence number or subject.

```APIDOC
// Get a message from the stream by its sequence number
msg, _ := s.GetMsg(ctx, 100)

// Get the last message for a specific subject
msg, _ = s.GetLastMsgForSubject(ctx, "ORDERS.new")

// Delete a message from the stream by its sequence number
_ = s.DeleteMsg(ctx, 100)
```

---

### NATS JetStream Stream Purge Operations

Source: https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

Methods to purge messages from a JetStream stream. Supports purging all messages, messages matching a subject, messages up to a sequence number, or keeping a specified number of the newest messages.

```APIDOC
// Purge all messages from a stream
_ = s.Purge(ctx)

// Purge messages stored on a specific subject
_ = s.Purge(ctx, jetstream.WithPurgeSubject("ORDERS.new"))

// Purge all messages up to a specified sequence number
_ = s.Purge(ctx, jetstream.WithPurgeSequence(100))

// Purge messages, keeping the 10 newest messages
_ = s.Purge(ctx, jetstream.WithPurgeKeep(10))
```

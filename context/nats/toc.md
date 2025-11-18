# Table of Contents for nats.md

This file is auto-generated. Run `task context:toc` to regenerate.

## Contents

L0001-L0007 | Install nats.go Client
L0008-L0010 | To get the latest released Go client:
L0011-L0013 | To get a specific version:
L0014-L0019 | Note that the latest major version for NATS Server is v2:
L0020-L0056 | NATS Client Request Fanout Example
L0057-L0131 | NATS Microservice Discovery and Status Commands
L0132-L0143 | Import NATS Micro Package
L0144-L0190 | Create and Manage KeyValue Bucket
L0191-L0238 | Watch KeyValue Bucket Changes in Go
L0239-L0262 | Get NATS KeyValue Bucket Status (Go)
L0263-L0314 | KeyValue Store (KV) Basic Usage
L0315-L0337 | Get NATS JetStream Bucket Status Go
L0338-L0365 | Purge KeyValue Bucket in NATS (Go)
L0366-L0392 | Create NATS Micro Service and Add Endpoint
L0393-L0484 | Basic NATS JetStream Usage
L0485-L0542 | Basic NATS Go Client Usage
L0543-L0573 | Add NATS Microservice with nats.go
L0574-L0595 | Context Support for NATS Operations
L0596-L0636 | NATS JetStream Basic Management: Streams and Consumers
L0637-L0691 | NATS.go Object Store: Basic CRUD Operations
L0692-L0743 | NATS JetStream Basic Usage: Publish, Subscribe, Consume
L0744-L0768 | List Keys in NATS KeyValue Bucket (Go)
L0769-L0800 | NATS JetStream Basic Usage
L0801-L0822 | Nkey Authentication
L0823-L0842 | NATS User Credentials Authentication
L0843-L0888 | Connect to NATS Cluster with Options
L0889-L0924 | Create NATS Services with Custom Queue Groups
L0925-L0950 | List Objects in NATS JetStream Bucket Go
L0951-L0984 | NATS Connection with Authentication
L0985-L1030 | Manage Pull Consumers (JetStream Interface) in NATS Go
L1031-L1042 | Run Go Tests with Specific Modfile
L1043-L1085 | TLS Connection Options
L1086-L1126 | Synchronous Publish Message
L1127-L1159 | Publishing Messages on a Stream
L1160-L1172 | Configure Message Buffer with Options
L1173-L1188 | Go Testing Dependency Management
L1189-L1211 | Queue Groups
L1212-L1223 | Discover NATS Service IDs using nats req
L1224-L1251 | NATS JetStream Stream Listing
L1252-L1273 | Iterate Over Incoming Messages
L1274-L1296 | Listing Streams and Stream Names
L1297-L1326 | KeyValue Store (KV) Watching for Changes
L1327-L1354 | Manage Consumers (Stream Interface) in NATS Go
L1355-L1405 | Object Store Basic Usage
L1406-L1435 | Fetch Single Messages for Work Queues
L1436-L1462 | Listing Consumers and Consumer Names
L1463-L1490 | List JetStream Consumers and Names in NATS Go
L1491-L1536 | Consumer Management
L1537-L1554 | Fetch Consumer Information in NATS Go
L1555-L1580 | Receive Messages with Push Consumers
L1581-L1609 | Fetch messages using Fetch() in Go
L1610-L1638 | Receiving Messages from Push Consumers
L1639-L1667 | Asynchronous Publish Message
L1668-L1709 | Consume messages with callback using Consume() in Go
L1710-L1742 | Wildcard Subscriptions
L1743-L1754 | Update Go Test Dependencies
L1755-L1805 | Stream Management (CRUD)
L1806-L1854 | NATS.go Object Store: Watching for Changes
L1855-L1874 | Object Store Watching for Changes
L1875-L1901 | JetStream Client Interfaces
L1902-L1913 | Retrieve Specific NATS Service Info using nats req
L1914-L1939 | NATS JetStream Pull Consumer Message Iterator Options
L1940-L1976 | Receiving Messages from Pull Consumers
L1977-L1994 | Create Ordered Consumer in NATS Go
L1995-L2016 | Update Object Metadata in NATS JetStream Go
L2017-L2034 | NATS JetStream Stream Info
L2035-L2053 | Aggregate Endpoints Using Groups
L2054-L2082 | Fetch messages by bytes using FetchBytes() in Go
L2083-L2111 | Fetch messages without waiting using FetchNoWait() in Go
L2112-L2162 | Advanced Connection and Messaging
L2163-L2197 | Override NATS Queue Groups on Groups and Endpoints
L2198-L2228 | NATS JetStream Stream Management (CRUD)
L2229-L2241 | Add Endpoint with Custom Subject
L2242-L2256 | Add Endpoint to Existing Service
L2257-L2268 | Retrieve NATS Service Statistics using nats req
L2269-L2300 | Disable NATS Queue Groups
L2301-L2319 | NATS JetStream Stream Message Operations
L2320-L2338 | NATS JetStream Stream Purge Operations

## Usage Examples

```bash
# Read a specific section by line range
sed -n '54,102p' context/nats/nats.md

# Search for specific topics
grep -n -i 'install' context/nats/nats.md

# Show context around a match (5 lines before and after)
grep -n -C 5 'example' context/nats/nats.md
```

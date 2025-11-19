### Install nats-top CLI Tool

Source: https://docs.nats.io/using-nats/nats-tools/nats_top/nats-top-tutorial

Installs the nats-top command-line tool using Go modules. This requires a Go environment setup. The command downloads and installs the latest version of the nats-top binary.

```bash
go install github.com/nats-io/nats-top@latest
```

```bash
sudo -E go install github.com/nats-io/nats-top
```

--------------------------------

### NATS Server Authentication Example

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol/nats-client-dev

Demonstrates how to start a NATS server with user and password authentication using command-line arguments. This requires a client to provide credentials for connection.

```shell
nats-server -DV -m 8222 -user foo -pass bar
```

--------------------------------

### Start Local NATS Server

Source: https://docs.nats.io/nats-concepts/what-is-nats/walkthrough_setup

Command to start a simple demonstration NATS server locally. An optional flag can enable HTTP monitoring. The server listens on TCP Port 4222 by default.

```bash
nats-server

```

```bash
nats-server -m 8222

```

--------------------------------

### Start Nex Node with Configuration

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

This command initiates a new Nex node process, connecting to the NATS server and loading the specified configuration file. It requires superuser privileges due to underlying kernel operations for network setup and process execution. The output provides crucial information about the node's status, connections, and security keys.

```bash
$ sudo ./nex node up --config=/home/kevin/simple.json
INFO[0000] Established node NATS connection to: nats://127.0.0.1:4222 
INFO[0000] Loaded node configuration from '/home/kevin/simple.json' 
INFO[0000] Virtual machine manager starting             
INFO[0000] Internal NATS server started                  client_url="nats://0.0.0.0:41339" 
INFO[0000] Use this key as the recipient for encrypted run requests  public_xkey=XDJJVJLRTWBIOHEEUPSNAAUACO6ZRW4WP65MXMGOX2WBGNCELLST5TWI
INFO[0000] NATS execution engine awaiting commands       id=NCOBPU3MCEA7LF6XADFD4P74CHW2OL6GQZYPPRRNPDSBNQ5BJPFHHQB5 version=0.0.1
INFO[0000] Called startVMM(), setting up a VMM on /tmp/.firecracker.sock-370707-cmjg61n52omq8dovolmg 
INFO[0000] VMM metrics disabled.                        
INFO[0000] refreshMachineConfiguration: [GET /machine-config][200] getMachineConfigurationOK  &{CPUTemplate: MemSizeMib:0xc0004ac108 Smt:0xc0004ac113 TrackDirtyPages:false VcpuCount:0xc0004ac100} 
INFO[0000] PutGuestBootSource: [PUT /boot-source][204] putGuestBootSourceNoContent  
INFO[0000] Attaching drive /tmp/rootfs-cmjg61n52omq8dovolmg.ext4, slot 1, root true. 
INFO[0000] Attached drive /tmp/rootfs-cmjg61n52omq8dovolmg.ext4: [PUT /drives/{drive_id}][204] putGuestDriveByIdNoContent  
INFO[0000] Attaching NIC tap0 (hwaddr 5a:65:8e:fa:7f:25) at index 1 
INFO[0000] startInstance successful: [PUT /actions][204] createSyncActionNoContent  
INFO[0000] SetMetadata successful                       
INFO[0000] Machine started                               gateway=192.168.127.1 hosttap=tap0 ip=192.168.127.6 nats_host=192.168.127.1 nats_port=41339 netmask=ffffff00 vmid=cmjg61n52omq8dovolmg
INFO[0000] Adding new VM to warm pool                    ip=192.168.127.6 vmid=cmjg61n52omq8dovolmg
INFO[0000] Received agent handshake                      message="Host-supplied metadata" vmid=cmjg61n52omq8dovolmg
```

--------------------------------

### Install and Start NATS Windows Service using sc.exe

Source: https://docs.nats.io/running-a-nats-service/introduction/windows_srv

This snippet demonstrates how to install and start the NATS server as a Windows service using the `sc.exe` command. It includes specifying the executable path and any necessary NATS server flags. Ensure the `NATS_PATH` environment variable is set correctly.

```shell
sc.exe create nats-server binPath= "%NATS_PATH%\nats-server.exe [nats-server flags]"
sc.exe start nats-server
```

--------------------------------

### Start nats-top Monitoring Tool

Source: https://docs.nats.io/using-nats/nats-tools/nats_top/nats-top-tutorial

Launches the nats-top real-time monitoring tool. This command assumes nats-top has been installed and the NATS server is running. It connects to the default NATS server address.

```bash
nats-top
```

--------------------------------

### Example Nex Node Configuration JSON

Source: https://docs.nats.io/using-nats/nex/getting-started/installing

A sample JSON file for configuring a Nex node. This configuration specifies resource directories, machine pool size, CNI settings, machine templates, and tags. It assumes default file locations for the kernel and root filesystem.

```json
{
    "default_resource_dir":"/tmp/wd",
    "machine_pool_size": 1,
    "cni": {
        "network_name": "fcnet",
        "interface_name": "veth0"
    },
    "machine_template": {
        "vcpu_count": 1,
        "memsize_mib": 256
    },
    "tags": {
        "simple": "true"
    }
}
```

--------------------------------

### C# NKey Authentication

Source: https://docs.nats.io/using-nats/developer/connecting/nkey

Demonstrates NKey authentication setup in C# using the NATS.Net library. This example configures the client with the URL and specifies the NKey file for authentication. Make sure to install the NATS.Net package.

```C#
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.Core;

await using var client = new NatsClient(new NatsOpts
{
    Url = "127.0.0.1",
    Name = "API NKey Example",
    AuthOpts = new NatsAuthOpts
    {
        NKeyFile = "/path/to/nkeys/user.nk"
    }
});
```

--------------------------------

### NATS Server TLS Configuration Example

Source: https://docs.nats.io/using-nats/developer/connecting/tls

This bash command demonstrates how to start a NATS server with TLS enabled, specifying certificate, key, and CA certificate files. It also enables client verification.

```bash
nats-server --tls --tlscert=server-cert.pem --tlskey=server-key.pem --tlscacert rootCA.pem --tlsverify
```

--------------------------------

### NATS Three Server Cluster Example with Debug (Bash)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This example shows how to start a three-server NATS cluster on the same host, including the debug flag '-D' for verbose output. It demonstrates starting the seed server with specific client and cluster ports. An alternative configuration file approach is also mentioned.

```bash
nats-server -p 4222 -cluster nats://localhost:4248 -D
```

--------------------------------

### Install NATS Server Binary from Source (Go)

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

Installs the latest development build of the NATS server directly from the main branch using the Go toolchain. This is suitable for developers who want to test the latest features. Ensure Go is installed and the $GOPATH/bin directory is in your system's PATH.

```shell
go install github.com/nats-io/nats-server/v2@main
```

--------------------------------

### Start NATS Server

Source: https://docs.nats.io/running-a-nats-service/clients

Command to start a NATS server. This is the initial step for setting up a NATS service. No specific inputs are required, and the output shows server status and listening ports.

```shell
nats-server
```

--------------------------------

### Install NATS CLI Tools (Go)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Installs the latest versions of nats-server, nats, nk, and nsc using Go modules. Ensure Go is installed and GOPATH is set up.

```shell
export GO111MODULE=on
go install github.com/nats-io/nats-server/v2@latest
go install github.com/nats-io/natscli/nats@latest
go install github.com/nats-io/nkeys/nk@latest
go install github.com/nats-io/nsc/v2@latest
```

--------------------------------

### Install nk Command-line Tool

Source: https://docs.nats.io/using-nats/nats-tools/nk

This command installs the 'nk' tool, a command-line utility for generating NKeys, using Go modules. Ensure you have Go installed and configured.

```bash
go install github.com/nats-io/nkeys/nk@latest
```

--------------------------------

### Install Go NATS Client SDK

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This command installs the Go NATS client SDK, which is required to build NATS services in Go. It ensures that the necessary libraries are available for your project.

```bash
$ go get github.com/nats-io/nats.go
```

--------------------------------

### Start NATS Server with Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This shell command demonstrates how to start the NATS server using a specified configuration file. Ensure the path to the configuration file is correct.

```shell
nats-server -c server.conf
```

--------------------------------

### Install NATS Server via Package Managers (Windows, Mac, Arch)

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

Provides commands for installing the nats-server using different package managers. Includes instructions for Windows (scoop), macOS (brew), and Arch Linux (AUR with yay). Ensure the respective package managers are installed.

```shell
scoop install main/nats-server
```

```shell
brew install nats-server
```

```shell
yay -S nats-server
```

--------------------------------

### Install NATS Server and Generate Configuration

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

Installs the NATS server using go get and generates a server configuration file using the nsc tool. The configuration includes settings for the operator JWT and account resolver, with an option to include a resolver configuration file.

```shell
go get github.com/nats-io/nats-server
nsc generate config --nats-resolver > resolver.conf

server_name: servertest
listen: 127.0.0.1:4222
http: 8222

jetstream: enabled

include resolver.conf
```

--------------------------------

### Install Nex CLI using curl

Source: https://docs.nats.io/using-nats/nex/getting-started/installing

This command downloads and executes the Nex installation script from Synadia. It might require sudo privileges depending on your system's user permissions.

```shell
curl -sSf https://nex.synadia.com/install.sh | sh
```

--------------------------------

### Running a Simple NATS Cluster (Bash)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This example demonstrates how to start three NATS servers on the same machine to form a cluster. Server A acts as the seed server, and Servers B and C establish routes to Server A. The output shows the command-line arguments for each server, including client ports, cluster URLs, and route configurations.

```bash
nats-server -p 4222 -cluster nats://localhost:4248 --cluster_name test-cluster
nats-server -p 5222 -cluster nats://localhost:5248 -routes nats://localhost:4248 --cluster_name test-cluster
nats-server -p 6222 -cluster nats://localhost:6248 -routes nats://localhost:4248 --cluster_name test-cluster
```

--------------------------------

### Start NATS Account Server

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

This command initiates the NATS Account Server, which is responsible for serving JWT credentials to NATS servers. It requires the path to the NATS configuration directory (`~/.nsc/nats/SAOP` in this example). The server will then vend JWT configurations on a specified endpoint.

```shell
nats-account-server -nsc ~/.nsc/nats/SAOP
```

--------------------------------

### Download NATS Server Binary (Shell)

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

This command-line snippet downloads a specific version or the latest release of the nats-server binary using curl. It's a quick way to get the executable for direct use. Ensure curl is installed and the path is correct for execution.

```shell
curl -fsSL https://binaries.nats.dev/nats-io/nats-server/v2@v2.11.6 | sh
```

```shell
curl -fsSL https://binaries.nats.dev/nats-io/nats-server/v2@latest | sh
```

--------------------------------

### Basic NATS Connection and Dispatcher Setup in Java

Source: https://docs.nats.io/using-nats/developer/receiving/drain

This Java snippet shows the initial setup for connecting to a NATS server and creating a dispatcher with an inline message handler. It uses `Nats.connect` to establish the connection and `CountDownLatch` for synchronization, allowing the program to wait for messages. This is a foundational example for building more complex NATS applications in Java.

```java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

// Use a latch to wait for a message to arrive
CountDownLatch latch = new CountDownLatch(1);

// Create a dispatcher and inline message handler

```

--------------------------------

### NATS Client Connection with Authentication

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol/nats-client-dev

Illustrates how a NATS client can establish a connection to a server that requires authentication, by including username and password in the connection URL. This example shows a basic connection string format.

```shell
nats.Connect("nats://foo:bar@localhost:4222")
```

--------------------------------

### Install and Run NATS Server with Docker

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

This snippet demonstrates installing the NATS server using Docker. It first pulls the latest NATS image and then runs it as a container, exposing the default NATS port. Docker must be installed prior to executing these commands.

```docker
docker pull nats:latest
```

```docker
docker run -p 4222:4222 -ti nats:latest
```

--------------------------------

### Install nats-top via Go

Source: https://docs.nats.io/using-nats/nats-tools/nats_top

Installs the nats-top tool using the Go build tools. This command fetches the latest version of the nats-top package and compiles it. Depending on the Go version and system permissions, 'sudo' might be required, and the '-E' flag is recommended to preserve environment variables.

```bash
go install github.com/nats-io/nats-top
```

```bash
go install github.com/nats-io/nats-top@latest
```

```bash
sudo -E go get github.com/nats-io/nats-top
```

--------------------------------

### Download NATS Python Examples

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-docker-python

Fetches the nats-pub.py and nats-sub.py example scripts from the official nats.py GitHub repository using curl. These scripts demonstrate basic NATS publishing and subscribing functionalities.

```shell
curl -o nats-pub.py -O -L https://raw.githubusercontent.com/nats-io/nats.py/master/examples/nats-pub/__main__.py
curl -o nats-sub.py -O -L https://raw.githubusercontent.com/nats-io/nats.py/master/examples/nats-sub/__main__.py
```

--------------------------------

### C NKey Authentication Setup

Source: https://docs.nats.io/using-nats/developer/connecting/nkey

Shows the setup for NKey authentication in C using the NATS client library. This involves creating NATS options, setting the NKey with a public key and a signature handler function, and then connecting to the NATS server. Remember to destroy the created objects.

```C
static natsStatus
sigHandler(
    char            **customErrTxt,
    unsigned char   **signature,
    int             *signatureLength,
    const char      *nonce,
    void            *closure)
{
    // Sign the given `nonce` and return the signature as `signature`.
    // This needs to allocate memory. The length of the signature is
    // returned as `signatureLength`.
    // If an error occurs the user can return specific error text through
    // `customErrTxt`. The library will free this pointer.

    return NATS_OK;
}

(...)

natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;
const char          *pubKey    = "my public key......";

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_SetNKey(opts, pubKey, sigHandler, NULL);
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### Starting NATS Cluster with Docker Compose

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This command initiates the NATS cluster defined in the 'nats-cluster.yaml' file. It uses Docker Compose to build and start the services, bringing the NATS cluster online.

```bash
docker-compose -f nats-cluster.yaml up
```

--------------------------------

### Start NATS Server with User/Password

Source: https://docs.nats.io/using-nats/developer/connecting/userpass

Starts the NATS server with specified username and password. This is a command-line operation.

```bash
nats-server --user myname --pass password
```

--------------------------------

### List NATS Microservices

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This command lists all active NATS microservices, including the 'EchoService' that was started. It helps verify that the service is registered and running correctly within the NATS ecosystem.

```bash
$ nats micro ls
```

--------------------------------

### NATS INFO Message Example

Source: https://docs.nats.io/reference/reference-protocols/nats-server-protocol

An example of an INFO message received by a NATS server, including server details and connection parameters.

```json
INFO {"server_id":"KP19vTlB417XElnv8kKaC5","version":"2.0.0","go":"","host":"localhost","port":5222,"auth_required":false,"tls_required":false,"tls_verify":false,"max_payload":1048576,"ip":"nats-route://127.0.0.1:5222/","connect_urls":["localhost:4222"]}
```

--------------------------------

### Install nsc Tool

Source: https://docs.nats.io/using-nats/nats-tools/nsc

Installs the latest version of the nsc tool by downloading and executing a Python installation script from GitHub. This is the primary method for setting up the NATS account configuration tool.

```shell
curl -L https://raw.githubusercontent.com/nats-io/nsc/master/install.py | python
```

--------------------------------

### NATS Context JSON Configuration Example

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli

An example of a NATS context configuration stored in JSON format. This file defines various settings such as server URL, authentication credentials, and JetStream domain.

```json
{
  "description": "",
  "url": "nats://127.0.0.1:4222",
  "token": "",
  "user": "",
  "password": "",
  "creds": "",
  "nkey": "",
  "cert": "",
  "key": "",
  "ca": "",
  "nsc": "",
  "jetstream_domain": "",
  "jetstream_api_prefix": "",
  "jetstream_event_prefix": "",
  "inbox_prefix": "",
  "user_jwt": ""
}
```

--------------------------------

### NATS Server Configuration with JWT Authentication

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This example shows a NATS server configuration file for a node using JWT-based authentication. It specifies the server port, cluster configuration with routes to other nodes, and potentially other settings for a decentralized auth setup.

```hocon
port = 4223

cluster {
  port = 6223
  routes [ nats://127.0.0.1:6222 ]
}

# debug = true
# trace = true
```

--------------------------------

### Start NATS Server with Single User Credentials (CLI)

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/username_password

This command demonstrates how to start the NATS server and provide a single user's credentials (username and password) directly via command-line arguments. This is an alternative to configuring it in a file.

```bash
> nats-server --user a --pass b
```

--------------------------------

### Perform Nex Node Preflight Check

Source: https://docs.nats.io/using-nats/nex/getting-started/installing

This command initiates a preflight check for the Nex node, validating dependencies such as CNI plugins, required binaries, CNI configuration, and user-provided files. It uses a specified configuration file and reports the status of each check.

```shell
$ nex node preflight --config=../examples/nodeconfigs/simple.json
Validating - Required CNI Plugins [/opt/cni/bin]
	✅ Dependency Satisfied - /opt/cni/bin/host-local [host-local CNI plugin]
	✅ Dependency Satisfied - /opt/cni/bin/ptp [ptp CNI plugin]
	✅ Dependency Satisfied - /opt/cni/bin/tc-redirect-tap [tc-redirect-tap CNI plugin]

Validating - Required binaries [/usr/local/bin]
	✅ Dependency Satisfied - /usr/local/bin/firecracker [Firecracker VM binary]

Validating - CNI configuration requirements [/etc/cni/conf.d]
	✅ Dependency Satisfied - /etc/cni/conf.d/fcnet.conflist [CNI Configuration]

Validating - User provided files []
	✅ Dependency Satisfied - /tmp/wd/vmlinux [VMLinux Kernel]
	✅ Dependency Satisfied - /tmp/wd/rootfs.ext4 [Root Filesystem Template]
```

--------------------------------

### Configure NATS Ping/Pong Settings (Python)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Provides a Python example for configuring the NATS client with a ping interval of 20 seconds and a maximum of 5 outstanding pings. The code uses an async connection setup and includes comments indicating where to perform connection-related actions.

```Python
nc = NATS()

await nc.connect(
   servers=["nats://demo.nats.io:4222"],
   # Set Ping Interval to 20 seconds and Max Pings Outstanding to 5
   ping_interval=20,
   max_outstanding_pings=5,
   )

# Do something with the connection.

```

--------------------------------

### Start NATS Seed Node Server

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Command to start the NATS server with a specified configuration file and enable debug logging. This is the initial step in setting up a NATS cluster.

```bash
nats-server -config ./seed.conf -D
```

--------------------------------

### Listing NATS Accounts with NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This command uses the NSC tool to list all accounts within the NATS setup. It displays the name and public key of each account, which is useful for verifying the setup and for referencing account keys in server configurations.

```bash
nsc list accounts
```

--------------------------------

### NATS Account and User Creation using NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This snippet demonstrates how to use the NATS Server Command Line Interface (NSC) to add accounts and users. These commands are foundational for setting up authentication and authorization within NATS, particularly for initializing a trusted operator setup.

```bash
nsc add account --name SYS
nsc add user    --name sys
nsc add account --name A
nsc add user -a A --name test
nsc add account --name B
nsc add user -a B --name test
```

--------------------------------

### Go Client Example Connection Message

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol

An example of the CONNECT message format used by the Go NATS client. This message is sent during the initial connection to the NATS server, containing client capabilities and identification.

```go
CONNECT {"verbose":false,"pedantic":false,"tls_required":false,"name":"","lang":"go","version":"1.2.2","protocol":1}​

```

--------------------------------

### NATS JWT Generation in C# using NATS.Jwt Package

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This snippet illustrates how to generate a NATS JWT using the NATS.Jwt package in C#. It requires the NATS.Jwt package. The example shows the setup for loading an account signing key and account identity, which are prerequisites for issuing JWTs.

```csharp
// dotnet add package NATS.Jwt --prerelease
using NATS.Jwt;
using NATS.Jwt.Models;
using NATS.NKeys;

const string accSeed = "SAANWFZ3JINNPERWT3ALE45U7GYT2ZDW6GJUIVPDKUF6GKAX6AISZJMAS4";
const string accId = "ACV63DGCZGOIT3P5ZA7PQT3KYJ6UDFFHZ7KETHYMDMZ4N44KYAQ2ZZ5F";

var jwt = new NatsJwt();

// Load account signing key

```

--------------------------------

### NATS Server Log Output - Seed Accepting Third Server Route

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Log output from the seed NATS server confirming it accepted the route connection from the third NATS server. This completes the cluster setup for this example.

```log
[83329] 2020/02/12 16:05:12.840111 [INF] 127.0.0.1:62945 - rid:2 - Route connection created
[83329] 2020/02/12 16:05:12.840350 [DBG] 127.0.0.1:62945 - rid:2 - Registering remote route "NBE7SLUDLFIMHS2U6347N3DQEJ"
[83329] 2020/02/12 16:05:12.840363 [DBG] 127.0.0.1:62945 - rid:2 - Sent local subscriptions to route
```

--------------------------------

### Start and Reload NATS Server Configuration

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Starts the NATS server with a specified configuration file and demonstrates how to reload the configuration without restarting the server.

```shell
nats-server -c server.conf
nats-server --signal reload
```

--------------------------------

### Build NATS Server Release from Source (Go & Goreleaser)

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

Builds a reproducible release version of the NATS server from source using goreleaser. This requires Go and goreleaser to be installed. The process involves cloning the repository, checking out a specific tag, and running the goreleaser build command. It is recommended to have a clean Git repository state before building.

```shell
go install github.com/goreleaser/goreleaser/v2@latest

git clone git@github.com:nats-io/nats-server.git
cd nats-server
git checkout v2.12.0 
[[ `git status --porcelain` ]] && echo "Must have repo in clean state before building"

goreleaser release --skip=announce,publish,validate --clean -f .goreleaser.yml
```

--------------------------------

### NATS Configuration: Include Directive Example

Source: https://docs.nats.io/running-a-nats-service/configuration

Illustrates the use of the `include` directive to split NATS server configurations into multiple files. It shows how to include external configuration files using relative paths, promoting modularity.

```shell
server.conf:
listen: 127.0.0.1:4222
include ./auth.conf

auth.conf:
authorization {
    token: "f0oBar"
}

> nats-server -c server.conf
```

--------------------------------

### NATS CONNECT Message Example

Source: https://docs.nats.io/reference/reference-protocols/nats-server-protocol

An example of a CONNECT message sent by a client to a NATS server, providing connection options such as TLS requirements and authentication details.

```json
CONNECT {"tls_required":false,"name":"wt0vffeQyoDGMVBC2aKX0b"}\r
```

--------------------------------

### Start NATS Server with Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/gateways/gateway

Command to start the NATS server using a specified configuration file. This is a common shell command for initiating the NATS service with custom settings.

```shell
nats-server -c A.conf
```

--------------------------------

### Run NATS Service Locally

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This command runs the Go NATS service from the main.go file. It starts the 'EchoService' locally, making it available to receive requests on the 'svc.echo' subject.

```bash
$ go run main.go
```

--------------------------------

### Start NATS Server with Token Authentication

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/tokens

This command demonstrates starting the NATS server with token authentication enabled via a command-line flag. The `--auth` flag takes the token string as an argument, which is then used by the server for authentication. This is an alternative to configuring the token in a file.

```shell
nats-server --auth s3cr3t
```

--------------------------------

### NATS Configuration File Syntax Examples

Source: https://docs.nats.io/running-a-nats-service/configuration

Illustrates various syntax elements supported in NATS configuration files, including comments, value assignments, arrays, maps, and optional terminators. It highlights the flexibility in defining configuration properties.

```shell
# Lines can be commented with `#` and `//`
# Values can be assigned to properties with delimiters:
# Equals sign: foo = 2
# Colon: foo: 2
# Whitespace: foo 2
# Arrays are enclosed in brackets: ["a", "b", "c"]
# Maps are enclosed in braces: {foo: 2}
# Maps can be assigned with no delimiter accounts { SYS {...}, cloud-user {...} }
# Semicolons can be optionally used as terminators host: 127.0.0.1; port: 4222;
```

--------------------------------

### NATS Authorization Configuration with Variables

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/authorization

An example of NATS authorization configuration using variables to define user permissions. It illustrates default permissions, and specific configurations for ADMIN, REQUESTOR, and RESPONDER roles, assigning them to different users. This setup allows for flexible and reusable permission definitions.

```hcl
authorization {
  default_permissions = {
    publish = "SANDBOX.*"
    subscribe = ["PUBLIC. திராக", "_INBOX. திராக"]
  }
  ADMIN = {
    publish = ">"
    subscribe = ">"
  }
  REQUESTOR = {
    publish = ["req.a", "req.b"]
    subscribe = "_INBOX. திராக"
  }
  RESPONDER = {
    subscribe = ["req.a", "req.b"]
    publish = "_INBOX. திராக"
  }
  users = [
    {user: admin,   password: $ADMIN_PASS, permissions: $ADMIN}
    {user: client,  password: $CLIENT_PASS, permissions: $REQUESTOR}
    {user: service,  password: $SERVICE_PASS, permissions: $RESPONDER}
    {user: other, password: $OTHER_PASS}
  ]
}
```

--------------------------------

### Java NATS JetStream Publisher and Subscriber Setup

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Sets up NATS JetStream subscribers and a publisher. It configures durable subscriptions, uses an AtomicInteger for tracking received messages, and manages subscriber and publisher threads. The code flushes the connection, starts threads, waits for completion, reports results, and deletes the stream.

```java
            // Setup the subscribers
            // - the PushSubscribeOptions can be re-used since all the subscribers are the same
            // - use a concurrent integer to track all the messages received
            // - have a list of subscribers and threads so I can track them
            PushSubscribeOptions pso = PushSubscribeOptions.builder().durable(exArgs.durable).build();
            AtomicInteger allReceived = new AtomicInteger();
            List<JsQueueSubscriber> subscribers = new ArrayList<>();
            List<Thread> subThreads = new ArrayList<>();
            for (int id = 1; id <= exArgs.subCount; id++) {
                // setup the subscription
                JetStreamSubscription sub = js.subscribe(exArgs.subject, exArgs.queue, pso);
                // create and track the runnable
                JsQueueSubscriber qs = new JsQueueSubscriber(id, exArgs, js, sub, allReceived);
                subscribers.add(qs);
                // create, track and start the thread
                Thread t = new Thread(qs);
                subThreads.add(t);
                t.start();
            }
            nc.flush(Duration.ofSeconds(1)); // flush outgoing communication with/to the server

            // create and start the publishing
            Thread pubThread = new Thread(new JsPublisher(js, exArgs));
            pubThread.start();

            // wait for all threads to finish
            pubThread.join();
            for (Thread t : subThreads) {
                t.join();
            }

            // report
            for (JsQueueSubscriber qs : subscribers) {
                qs.report();
            }

            System.out.println();

            // delete the stream since we are done with it.
            jsm.deleteStream(exArgs.stream);
        }
        catch (Exception e) {
            e.printStackTrace();
        }
    }

    static class JsPublisher implements Runnable {
        JetStream js;
        ExampleArgs exArgs;

        public JsPublisher(JetStream js, ExampleArgs exArgs) {
            this.js = js;
            this.exArgs = exArgs;
        }

        @Override
        public void run() {
            for (int x = 1; x <= exArgs.msgCount; x++) {
                try {
                    PublishAck pa = js.publish(exArgs.subject, ("Data # " + x).getBytes(StandardCharsets.US_ASCII));
                } catch (IOException | JetStreamApiException e) {
                    // something pretty wrong here
                    e.printStackTrace();
                    System.exit(-1);
                }
            }
        }
    }

    static class JsQueueSubscriber implements Runnable {
        int id;
        int thisReceived;
        List<String> datas;

        ExampleArgs exArgs;
        JetStream js;
        JetStreamSubscription sub;
        AtomicInteger allReceived;

        public JsQueueSubscriber(int id, ExampleArgs exArgs, JetStream js, JetStreamSubscription sub, AtomicInteger allReceived) {
            this.id = id;
            thisReceived = 0;
            datas = new ArrayList<>();
            this.exArgs = exArgs;
            this.js = js;
            this.sub = sub;
            this.allReceived = allReceived;
        }

        public void report() {
            System.out.printf("Sub # %d handled %d messages.\n", id, thisReceived);
        }

        @Override
        public void run() {
            while (allReceived.get() < exArgs.msgCount) {
                try {
                    Message msg = sub.nextMessage(Duration.ofMillis(500));
                    while (msg != null) {
                        thisReceived++;
                        allReceived.incrementAndGet();
                        String data = new String(msg.getData(), StandardCharsets.US_ASCII);
                        datas.add(data);
                        System.out.printf("QS # %d message # %d %s\n", id, thisReceived, data);
                        msg.ack();

                        msg = sub.nextMessage(Duration.ofMillis(500));
                    }
                } catch (InterruptedException e) {
                    // just try again
                }
            }
            System.out.printf("QS # %d completed.\n", id);
        }
    }
}
```

--------------------------------

### Publish and Subscribe with NATS in C

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This C example demonstrates publishing and subscribing to NATS messages using the nats-client C library. It shows how to connect to the NATS server, set up a subscription callback, and publish messages.

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "nats/nats.h"

// Callback function for received messages
void
onMsg(natsConnection *conn, natsSubscription *sub, natsMsg *msg, void *closure)
{
    printf("Received msg: %s - %.*s\n",
           natsMsg_GetSubject(msg),
           natsMsg_GetDataLength(msg),
           natsMsg_GetData(msg));

    // Need to destroy the message!
    natsMsg_Destroy(msg);
}

int main(int argc, char **argv)
{
    natsConnection      *conn = NULL;
    natsSubscription    *sub  = NULL;
    natsStatus          s;

    // Connect to NATS server
    s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);
    if (s != NATS_OK) {
        fprintf(stderr, "Error connecting to NATS: %s\n", natsStatus_GetError(s));
        return 1;
    }
    printf("Connected to NATS at %s\n", NATS_DEFAULT_URL);

    // Subscribe to messages
    s = natsConnection_Subscribe(&sub, conn, "time.வுகளை", onMsg, NULL);
    if (s != NATS_OK) {
        fprintf(stderr, "Error subscribing to 'time.வுகளை': %s\n", natsStatus_GetError(s));
        natsConnection_Destroy(conn);
        return 1;
    }

    // Publish messages
    natsConnection_Publish(conn, "time.A.east", (const void*)"A", 1);
    natsConnection_Publish(conn, "time.B.east", (const void*)"B", 1);
    natsConnection_Publish(conn, "time.C.west", (const void*)"C", 1);
    natsConnection_Publish(conn, "time.D.west", (const void*)"D", 1);
    printf("Published 4 messages.\n");

    // Keep the connection open to receive messages for a short period
    nats_Sleep(2000); // Sleep for 2 seconds

    // Destroy objects that were created
    natsSubscription_Destroy(sub);
    natsConnection_Destroy(conn);

    return 0;
}

```

--------------------------------

### NATS Server Configuration File Example

Source: https://docs.nats.io/nats-server/configuration

An example of a NATS server configuration file showing the structure for general settings, JetStream, TLS, gateway, leafnodes, MQTT, websocket, accounts, authorization, mappings, and resolver settings. All blocks and properties are optional, except for host and port.

```hcl
#General settings
host: 0.0.0.0
port: 4222

# Various server level options
# ...

# The following sections are maps with a set of (nested) properties

jetstream {
    # JetStream storage location, limits and encryption
	store_dir: nats
}

tls {
    # Configuration map for tls parameters used for client connections, 
    # routes and https monitoring connections.
}

gateway {
    # Configuration map for gateway. Gateways are used to connected clusters.
}

leafnodes {
    # Configuration map for leafnodes. LeafNodes are lightweight clusters.
}

mqtt {
    # Configuration map for mqtt. Allow clients to connect via mqtt protocol.
}

websocket {
    # Configuration map for websocket. Allow clients to connect via websockets.
}

accounts {
    # List of accounts and user within accounts
    # User may have an authorization and authentication section
}

authorization {
    # User may have an authorization and authentication section
    # This section is only useful when no accounts are defined
}

mappings {
    # Subject mappings for default account
    # When accounts are defined this section must be in the account map
}

resolver {
    # Pointer to external Authentication/Authorization resolver
    # There are multiple possible resolver type explained in their own chapters of this docuemntaion
    # memory, nats-base, url ... more may be added in the future
    # This parameter can be a value `MEMORY` for simple configuration
    # or a map of properties for connecting to the resolver
}

resolver_tls {
    # TLS configuration for an URL based resolver
}

resolver_preload {
    # List of JWT tokens to be loaded at server start.
}

```

--------------------------------

### Start NATS Server with TLS Verification via Command Line

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/tls_mutual_auth

This command line demonstrates how to start the NATS server with TLS client certificate verification enabled. It specifies the server's certificate, key, and the CA certificate for verification. This method achieves the same outcome as the `tls` configuration block.

```bash
nats-server --tlsverify --tlscert=server-cert.pem --tlskey=server-key.pem --tlscacert=rootCA.pem
```

--------------------------------

### NATS Server Output Indicating Leaf Node Connection

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This is example output from the NATS server logs, showing that it has started listening for client connections on the default port and has successfully connected a leaf node to the specified remote server ('connect.ngs.global').

```log
...
[4985] 2023/03/03 10:55:51.577569 [INF] Listening for client connections on 0.0.0.0:4222
...
[4985] 2023/03/03 10:55:51.918781 [INF] Connected leafnode to "connect.ngs.global"

```

--------------------------------

### Configure NATS Stream with Sources using JSON

Source: https://docs.nats.io/nats-concepts/jetstream/source_and_mirror/source_and_mirror_example

This JSON configuration defines a NATS JetStream stream with multiple sources. It specifies subjects, retention policies, and source-specific configurations like filter subjects and starting sequence numbers. It is used with the NATS CLI for stream management.

```json
{
  "name": "SOURCE_TARGET",
  "subjects": [
    "foo1.ext.*",
    "foo2.ext.*"
  ],
  "discard": "old",
  "duplicate_window": 120000000000,
  "sources": [
    {
      "name": "SOURCE1_ORIGIN"
    }
  ],
  "deny_delete": false,
  "sealed": false,
  "max_msg_size": -1,
  "allow_rollup_hdrs": false,
  "max_bytes": -1,
  "storage": "file",
  "allow_direct": false,
  "max_age": 0,
  "max_consumers": -1,
  "max_msgs_per_subject": -1,
  "num_replicas": 1,
  "name": "SOURCE_TARGET",
  "deny_purge": false,
  "compression": "none",
  "max_msgs": -1,
  "retention": "limits",
  "mirror_direct": false
}
```

```json
{
  "name": "SOURCE_TARGET",
  "subjects": [
    "foo1.ext.*",
    "foo2.ext.*"
  ],
  "discard": "old",
  "duplicate_window": 120000000000,
  "sources": [
    {
      "name": "SOURCE1_ORIGIN",
      "filter_subject": "foo1.bar",
      "opt_start_seq": 42,
      "external": {
        "deliver": "",
        "api": "$JS.domainA.API"
      }
    },
    {
      "name": "SOURCE2_ORIGIN",
      "filter_subject": "foo2.bar"
    }
  ],
  "consumer_limits": {
    
  },
  "deny_delete": false,
  "sealed": false,
  "max_msg_size": -1,
  "allow_rollup_hdrs": false,
  "max_bytes": -1,
  "storage": "file",
  "allow_direct": false,
  "max_age": 0,
  "max_consumers": -1,
  "max_msgs_per_subject": -1,
  "num_replicas": 1,
  "name": "SOURCE_TARGET",
  "deny_purge": false,
  "compression": "none",
  "max_msgs": -1,
  "retention": "limits",
  "mirror_direct": false
}
```

--------------------------------

### Configure NATS Ping/Pong Settings (Go)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Demonstrates how to configure the NATS client in Go to set a ping interval of 20 seconds and a maximum of 5 outstanding pings. This ensures timely detection of connection issues. The example shows the connection setup and error handling.

```Go
// Set Ping Interval to 20 seconds and Max Pings Outstanding to 5
nc, err := nats.Connect("demo.nats.io", nats.Name("API Ping Example"), nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5))
if err != nil {
    log.Fatal(err)
}
deferr nc.Close()

// Do something with the connection

```

--------------------------------

### Install NATS Helm Chart

Source: https://docs.nats.io/running-a-nats-service/nats-kubernetes

Installs the NATS chart using Helm with the specified release name and repository. Assumes the repository has already been added.

```shell
helm install nats nats/nats
```

--------------------------------

### NATS Server Event JSON Output

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

This is an example of the JSON output received when subscribing to NATS server events. It details server information, account activity (connections, total connections), and client connection details (start time, host, user, language, version, stop time, reason for closure).

```json
{
  "server": {
    "host": "0.0.0.0",
    "id": "NBTGVY3OKDKEAJPUXRHZLKBCRH3LWCKZ6ZXTAJRS2RMYN3PMDRMUZWPR",
    "ver": "2.0.0-RC5",
    "seq": 32,
    "time": "2019-05-03T14:53:15.455266-05:00"
  },
  "acc": "ADWJVSUSEVC2GHL5GRATN2LOEOQOY2E6Z2VXNU3JEIK6BDGPWNIW3AXF",
  "conns": 1,
  "total_conns": 1
}
{
  "server": {
    "host": "0.0.0.0",
    "id": "NBTGVY3OKDKEAJPUXRHZLKBCRH3LWCKZ6ZXTAJRS2RMYN3PMDRMUZWPR",
    "ver": "2.0.0-RC5",
    "seq": 33,
    "time": "2019-05-03T14:53:15.455304-05:00"
  },
  "client": {
    "start": "2019-05-03T14:53:15.453824-05:00",
    "host": "127.0.0.1",
    "id": 6,
    "acc": "ADWJVSUSEVC2GHL5GRATN2LOEOQOY2E6Z2VXNU3JEIK6BDGPWNIW3AXF",
    "user": "UACPEXCAZEYWZK4O52MEGWGK4BH3OSGYM3P3C3F3LF2NGNZUS24IVG36",
    "name": "NATS Sample Publisher",
    "lang": "go",
    "ver": "1.7.0",
    "stop": "2019-05-03T14:53:15.45526-05:00"
  },
  "sent": {
    "msgs": 1,
    "bytes": 3
  },
  "received": {
    "msgs": 0,
    "bytes": 0
  },
  "reason": "Client Closed"
}
```

--------------------------------

### Start NATS Server with Monitoring

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli/natsbench

Starts the NATS server with JetStream enabled and the HTTP monitoring interface on port 8222. This is a prerequisite for running benchmarks and observing server status.

```bash
nats-server -m 8222 -js
```

--------------------------------

### NATS Server Mapping Configuration Examples

Source: https://docs.nats.io/nats-concepts/subject_mapping

Configuration snippets for NATS server demonstrating subject mapping rules. These examples show how to define transformations for subjects, including wildcard usage, for both central and leaf clusters.

```yaml
server_name: "hub"
cluster: { name: "hub" }
mappings: {
	orders.* orders.central.{{wildcard(1)}}
}
```

```yaml
server_name: "hub"
cluster: { name: "hub" }
mappings: {
	orders.> orders.central.>
}
```

```yaml
server_name: "store1"
cluster: { name: "store1" }
mappings: {
	orders.central.* orders.local.{{wildcard(1)}}
}
```

--------------------------------

### NATS Server Configuration with NKEYS Authentication

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This is a NATS server configuration file example that utilizes NKEYS for authentication. It defines the server port, cluster settings, system account, and explicitly lists accounts with their associated NKEYS users. This configuration represents a centralized auth model.

```hocon
port = 4222

cluster {
  port = 6222

  # We will bridge two different servers with different auth models via routes
  # routes [ nats://127.0.0.1:6223 ]
}

system_account = ABKOWIYVTHNEK5HELPWLAT2CF2CUPELIK4SZH2VCJHLFU22B5U2IIZUO

accounts {
  # Account A
  ADFB2JXYTXOJEL6LNAXDREUGRX35BOLZI3B4PFFAC7IRPR3OA4QNKBN2 {
    nkey: ADFB2JXYTXOJEL6LNAXDREUGRX35BOLZI3B4PFFAC7IRPR3OA4QNKBN2
    users = [
      {nkey: "UAPOK2P7EN3UFBL7SBJPQK3M3JMLALYRYKX5XWSVMVYK63ZMBHTOHVJR" }
    ]
  }

  # Account B
  ACWOMQA7PZTKJSBTR7BF6TBK3D776734PWHWDKO7HFMQOM5BIOYPSYZZ {
  }

  # Account SYS
  ABKOWIYVTHNEK5HELPWLAT2CF2CUPELIK4SZH2VCJHLFU22B5U2IIZUO {
  }
}
```

--------------------------------

### Install NATS.py and Dependencies

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-docker-python

Updates package lists and installs necessary build tools and curl, followed by the installation of the nats.py Python client library with NKeys support. This prepares the Python environment for NATS communication.

```shell
apt-get update && apt-get install -y build-essential curl
pip install asyncio-nats-client[nkeys]
```

--------------------------------

### List NATS Nodes using 'nex'

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

The 'nex node ls' command is used to display summary information for all discovered NATS execution nodes. This includes the node's ID, version, uptime, and the number of workloads it is managing. This command is useful for monitoring the status of your NATS infrastructure.

```bash
$ nex node ls
╭─────────────────────────────────────────────────────────────────────────────────────────╮
│                                   NATS Execution Nodes                                  │
├──────────────────────────────────────────────────────────┬─────────┬────────┬───────────┤
│ ID                                                       │ Version │ Uptime │ Workloads │
├──────────────────────────────────────────────────────────┼─────────┼────────┼───────────┤
│ NCOBPU3MCEA7LF6XADFD4P74CHW2OL6GQZYPPRRNPDSBNQ5BJPFHHQB5 │ 0.0.1   │ 10m15s │         0 │
╰──────────────────────────────────────────────────────────┴─────────┴────────┴───────────╯
```

--------------------------------

### NATS Server Syslog Configuration Examples

Source: https://docs.nats.io/running-a-nats-service/configuration/logging

Provides examples for configuring the NATS server to send logs to a remote syslog server using UDP or a syslog URI format.

```bash
nats-server -r udp://localhost:514
```

```bash
nats-server -r syslog://<hostname>:<port>
```

```bash
syslog://logs.papertrailapp.com:26900
```

--------------------------------

### Generate Config and Start NATS Server

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This snippet demonstrates how to generate a NATS server configuration file and start the NATS server in the background. It uses the 'nsc generate config' command to create the configuration and then launches 'nats-server' with the generated configuration file.

```shell
nsc generate config --nats-resolver > nats-res.cfg
nats-server -c nats-res.cfg --addr localhost --port 4222 &
```

--------------------------------

### NATS Subject Mappings Configuration Example

Source: https://docs.nats.io/nats-server/configuration

An example demonstrating subject mappings in NATS configuration, both at the server level for the default account and within specific accounts. This allows for subject aliasing and pattern-based translation.

```hcl
host: 0.0.0.0
port:4222

mappings: {
	foo: bar
}

accounts: {
    accountA: { 
	mappings: {
	    orders.acme.*: orders.$1
	}
        users: [
            {user: admin, password: admin},
            {user: user, password: user}
           ]
    },
}

```

--------------------------------

### NATS Server Configuration Example

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

This is an example of a NATS server configuration file (`server.conf`). It specifies the path to the operator JWT, the public key of the system account, and the URL of the account resolver (NATS Account Server). The `URL()` wrapper is important for specifying the resolver address.

```HCL
operator: /Users/synadia/.nsc/nats/SAOP/SAOP.jwt
system_account: ADWJVSUSEVC2GHL5GRATN2LOEOQOY2E6Z2VXNU3JEIK6BDGPWNIW3AXF
resolver: URL(http://localhost:9090/jwt/v1/accounts/)
```

--------------------------------

### NATS Cluster Authorization Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering/cluster_config

Configures authorization settings for NATS cluster route connections. It specifies user credentials, password, and timeout for authentication between NATS servers. This setup ensures secure communication and proper routing within the cluster.

```hcl
cluster {
  name: example

  # host/port for inbound route connections from other server
  listen: localhost:4244

  # Authorization for route connections
  # Other server can connect if they supply the credentials listed here
  # This server will connect to discovered routes using this user
  authorization {
    user: route_user
    password: pwd
    timeout: 0.5
  }

  # This server establishes routes with these server.
  # This server solicits new routes and Routes are actively solicited and connected to from this server.
  # Other servers can connect to us if they supply the correct credentials
  # in their routes definitions from above.
  routes = [
    nats://route_user:pwd@127.0.0.1:4245
    nats://route_user:pwd@127.0.0.1:4246
  ]
}
```

--------------------------------

### Run NATS Server Docker Image

Source: https://docs.nats.io/running-a-nats-service/nats_docker/nats-docker-tutorial

Command to run the NATS server Docker image. It exposes the necessary ports for client connections (4222), monitoring (8222), and routing (6222). Ensure Docker is installed and running.

```bash
docker run -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats-server -ti nats:latest
```

--------------------------------

### Go: User Provisioning Service and Process Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This Go code sets up a distributed user provisioning service using channels and goroutines. It includes functions for authorization token handling, JWT generation, and initiating user provisioning processes. The service allows for different authorization strategies and securely generates user JWTs without exposing private keys.

```go
func ObtainAuthorizationToken() interface{} {
    // whatever you want, 3rd party token/username&password
    return ""
}

func IsTokenAuthorized(token interface{}) bool {
    // whatever logic to determine if the input authorizes the requester to obtain a user jwt
    return token.(string) == ""
}

// request struct to exchange data
type userRequest struct {
    UserJWTResponseChan chan string
    UserPublicKey       string
    AuthInfo            interface{}
}

func startUserProvisioningService(isAuthorizedCb func(token interface{}) bool) chan userRequest {
    userRequestChan := make(chan userRequest) // channel to send requests for jwt to
    go func() {
        accountSigningKey := GetAccountSigningKey() // Setup, obtain account signing key
        for {
            req := <-userRequestChan // receive request
            if !isAuthorizedCb(req.AuthInfo) {
                fmt.Printf("Request is not authorized to receive a JWT, timeout on purpose")
            } else if userJWT := generateUserJWT(req.UserPublicKey, accountSigningKey); userJWT != "" {
                req.UserJWTResponseChan <- userJWT // respond with jwt
            }
        }
    }()
    return userRequestChan
}

func startUserProcess(userRequestChan chan userRequest, obtainAuthorizationCb func() interface{}) {
    requestUser := func(userRequestChan chan userRequest, authInfo interface{}) (jwtAuthOption nats.Option) {
        userPublicKey, _, userKeyPair := generateUserKey()
        respChan := make(chan string)
        // request jwt
        userRequestChan <- userRequest{
            respChan,
            userPublicKey,
            authInfo,
        }
        userJWT := <-respChan // wait for response
        // userJWT and userKeyPair can be used in conjunction with this nats.Option
        jwtAuthOption = nats.UserJWT(func() (string, error) {
            return userJWT, nil
        },
            func(bytes []byte) ([]byte, error) {
                return userKeyPair.Sign(bytes)
            },
        )
        // Alternatively you can create a creds file and use it as nats.Option
        return
    }
    go func() {
        jwtAuthOption := requestUser(userRequestChan, obtainAuthorizationCb())
        nc, err := nats.Connect("nats://localhost:4222", jwtAuthOption)
        if err != nil {
            return
        }
        defer nc.Close()
        time.Sleep(time.Second) // simulate work one would want to do
    }()
}

func RequestUserDistributed() {
    reqChan := startUserProvisioningService(IsTokenAuthorized)
    defer close(reqChan)
    // start multiple user processes
    for i := 0; i < 4; i++ {
        startUserProcess(reqChan, ObtainAuthorizationToken)
    }
    time.Sleep(5 * time.Second)
}
```

--------------------------------

### NATS PUB Message Syntax and Example

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol

The PUB message publishes a payload to a subject. It requires the subject, payload size, and optionally a reply-to subject. The example demonstrates publishing a string message and a request message.

```nats
PUB FOO 11​
Hello NATS!​

```

```nats
PUB FRONT.DOOR JOKE.22 11​
Knock Knock​

```

```nats
PUB NOTIFY 0​
​

```

--------------------------------

### Python NATS Subscribe, Publish, and Drain Example

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Provides an asynchronous example in Python for connecting to NATS, subscribing with a callback handler, publishing messages, and gracefully draining the subscription. This example uses the nats.aio.client library.

```python
import asyncio
from nats.aio.client import Client as NATS

async def example(loop):
    nc = NATS()

    await nc.connect("nats://127.0.0.1:4222", loop=loop)

    async def handler(msg):
        print("[Received] ", msg)
        await nc.publish(msg.reply, b'I can help')

        # Can check whether client is in draining state
        if nc.is_draining:
            print("Connection is draining")

    sid = await nc.subscribe("help", "workers", cb=handler)
    await nc.flush()

    # Gracefully unsubscribe the subscription
    await nc.drain(sid)
```

--------------------------------

### Manage NATS JetStream Streams and Consumers in JavaScript

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

This JavaScript code demonstrates managing NATS JetStream streams and consumers. It covers listing, adding, updating, finding, and retrieving information about streams. It also shows how to get messages, delete messages, and purge streams. This example uses the NATS client for JavaScript.

```javascript
import { AckPolicy, connect, Empty } from "../../src/mod.ts";

const nc = await connect();
const jsm = await nc.jetstreamManager();

// list all the streams, the `next()` function
// retrieves a paged result.
const streams = await jsm.streams.list().next();
streams.forEach((si) => {
    console.log(si);
});

// add a stream
const stream = "mystream";
const subj = `mystream.*`;
await jsm.streams.add({ name: stream, subjects: [subj] });

// publish a reg nats message directly to the stream
for (let i = 0; i < 10; i++) {
    nc.publish(`${subj}.a`, Empty);
}

// find a stream that stores a specific subject:
const name = await jsm.streams.find("mystream.A");

// retrieve info about the stream by its name
const si = await jsm.streams.info(name);

// update a stream configuration
si.config.subjects?.push("a.b");
await jsm.streams.update(name, si.config);

// get a particular stored message in the stream by sequence
// this is not associated with a consumer
const sm = await jsm.streams.getMessage(stream, { seq: 1 });
console.log(sm.seq);

// delete the 5th message in the stream, securely erasing it
await jsm.streams.deleteMessage(stream, 5);

// purge all messages in the stream, the stream itself
// remains.
await jsm.streams.purge(stream);


```

--------------------------------

### NATS Server Startup Log Output

Source: https://docs.nats.io/running-a-nats-service/configuration/gateways/gateway

Example log output from a NATS server upon startup, showing version information, gateway name, listening addresses for clients and gateways, and connection statuses. This helps in verifying server initialization and gateway presence.

```log
[85803] 2019/05/07 10:50:55.902474 [INF] Starting nats-server version 2.0.0
[85803] 2019/05/07 10:50:55.903669 [INF] Gateway name is A
[85803] 2019/05/07 10:50:55.903684 [INF] Listening for gateways connections on localhost:7222
[85803] 2019/05/07 10:50:55.903696 [INF] Address for gateway "A" is localhost:7222
[85803] 2019/05/07 10:50:55.903909 [INF] Listening for client connections on 0.0.0.0:4222
[85803] 2019/05/07 10:50:55.903914 [INF] Server id is NBHUDBF3TVJSWCDPG2HSKI4I2SBSPDTNYEXEMOFAZUZYXVA2IYRUGPZU
[85803] 2019/05/07 10:50:55.903917 [INF] Server is ready
[85803] 2019/05/07 10:50:56.830669 [INF] 127.0.0.1:50892 - gid:2 - Processing inbound gateway connection
[85803] 2019/05/07 10:50:56.830673 [INF] 127.0.0.1:50891 - gid:1 - Processing inbound gateway connection
[85803] 2019/05/07 10:50:56.831079 [INF] 127.0.0.1:50892 - gid:2 - Inbound gateway connection from "C" (NBHWDFO3KHANNI6UCEUL27VNWL7NWD2MC4BI4L2C7VVLFBSMZ3CRD7HE) registered
[85803] 2019/05/07 10:50:56.831211 [INF] 127.0.0.1:50891 - gid:1 - Inbound gateway connection from "B" (ND2UJB3GFUHXOQ2UUMZQGOCL4QVR2LRJODPZH7MIPGLWCQRARJBU27C3) registered
[85803] 2019/05/07 10:50:56.906103 [INF] Connecting to explicit gateway "B" (localhost:7333) at 127.0.0.1:7333
[85803] 2019/05/07 10:50:56.906104 [INF] Connecting to explicit gateway "C" (localhost:7444) at 127.0.0.1:7444
[85803] 2019/05/07 10:50:56.906404 [INF] 127.0.0.1:7333 - gid:3 - Creating outbound gateway connection to "B"
[85803] 2019/05/07 10:50:56.906444 [INF] 127.0.0.1:7444 - gid:4 - Creating outbound gateway connection to "C"
[85803] 2019/05/07 10:50:56.906647 [INF] 127.0.0.1:7444 - gid:4 - Outbound gateway connection to "C" (NBHWDFO3KHANNI6UCEUL27VNWL7NWD2MC4BI4L2C7VVLFBSMZ3CRD7HE) registered
[85803] 2019/05/07 10:50:56.906772 [INF] 127.0.0.1:7333 - gid:3 - Outbound gateway connection to "B" (ND2UJB3GFUHXOQ2UUMZQGOCL4QVR2LRJODPZH7MIPGLWCQRARJBU27C3) registered
```

--------------------------------

### Account JWT Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

An example of a decoded account JWT, detailing its issuance information, name, and specific NATS service limits such as connections, data transfer, and wildcard subscriptions.

```json
{
 "iat": 1603474600,
 "iss": "OBU5O5FJ324UDPRBIVRGF7CNEOHGLPS7EYPBTVQZKSBHIIZIB6HD66JF",
 "jti": "CZDE4PM7MGFNYHRZSE6INTP6QDU4DSLACVHPQFA7XEYNJT6R6LLQ",
 "name": "demo-test",
 "nats": {
  "limits": {
   "conn": -1,
   "data": -1,
   "exports": -1,
   "imports": -1,
   "leaf": -1,
   "payload": -1,
   "subs": -1,
   "wildcards": true
  }
 },
 "sub": "ADKGAJU55CHYOIF5H432K2Z2ME3NPSJ5S3VY5Q42Q3OTYOCYRRG7WOWV",
 "type": "account"
}
```

--------------------------------

### Running NATS Server with Configuration File

Source: https://docs.nats.io/nats-server/configuration

This command starts the NATS server using a specified configuration file. Ensure the configuration file is correctly formatted.

```shell
nats-server -config my-server.conf
```

--------------------------------

### NKEY Account Configuration Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Illustrates how to use NKEYs in the NATS account configuration file. It shows users associated with specific NKEYs for authentication.

```yaml
accounts: {
    A: {
        users: [{nkey:UC435ZYS52HF72E2VMQF4GO6CUJOCHDUUPEBU7XDXW5AQLIC6JZ46PO5}]
        imports: [{stream: {account: B, subject: "foo"}}]
    },
    B: {
        users: [{nkey:UARZVI6JAV7YMJTPRANXANOOW4K3ZCD45NYP6S7C7XKCBHPVN2TFZ7ZC}]
        exports: [{stream: "foo"}]
    },
}
```

--------------------------------

### C NATS Client TLS Connection Setup

Source: https://docs.nats.io/using-nats/developer/connecting/tls

Initializes a TLS connection to a NATS server using the C NATS client library. It demonstrates the creation of NATS options, loading of certificates (client cert/key and CA), and establishing the connection. Proper cleanup of created objects is also shown.

```c
natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_LoadCertificatesChain(opts, "client-cert.pem", "client-key.pem");
if (s == NATS_OK)
    s = natsOptions_LoadCATrustedCertificates(opts, "rootCA.pem");
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### Start NATS Server with Leaf Node Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This command starts the NATS server using a specified configuration file that includes the leaf node settings. The `-c` flag points to the configuration file. The server will then listen for client connections and establish leaf node connections as defined in the configuration.

```bash
nats-server -c /tmp/ngs_leaf.conf
```

--------------------------------

### Verify Statically Linked Binary

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This command uses the 'file' utility to inspect the compiled executable and confirm that it is a statically linked binary. This is a crucial step before deploying the service.

```bash
$ file echoservice
```

--------------------------------

### Start NATS Server with Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This command starts the NATS server using a specified configuration file. Ensure the configuration file (e.g., /tmp/server.conf) is correctly set up to enable leaf node functionalities.

```bash
nats-server -c /tmp/server.conf
```

--------------------------------

### User JWT Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

An example of a decoded user JWT, presenting its issuance timestamp, issuer, unique identifier, name, and NATS permissions for public and subscription operations.

```json
{
 "iat": 1603475001,
 "iss": "ADKGAJU55CHYOIF5H432K2Z2ME3NPSJ5S3VY5Q42Q3OTYOCYRRG7WOWV",
 "jti": "GOOPXCFDWVMEU3U6I6MT344Z56MGBYIS42GDXMUXDFA3NYDR2RUQ",
 "name": "alpha",
 "nats": {
  "pub": {},
  "sub": {}
 },
 "sub": "UC56LV5NNMP5FURQZ7HZTGWCRRTWSMHZNNELQMHDLH3DCYNGX57B2TN6",
 "type": "user"
}
>
```

--------------------------------

### Create NATS Subscriber using CLI

Source: https://docs.nats.io/nats-concepts/core-nats/pubsub/pubsub_walkthrough

Starts a NATS client that listens for messages on a specified subject. The `<subject>` parameter determines which messages the subscriber will receive. This command is fundamental for setting up message consumers.

```bash
nats sub <subject>
```

```bash
nats sub msg.test
```

```bash
nats sub msg.test.new
```

```bash
nats sub msg.*
```

--------------------------------

### JavaScript JetStream Pull Subscription Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This JavaScript example demonstrates setting up a JetStream pull subscription. It publishes messages, creates a pull subscriber with explicit ACKs and a short ack wait to trigger redeliveries, and periodically pulls messages. It uses async/await syntax and the nats.deno module.

```javascript
import { AckPolicy, connect, nanos } from "../../src/mod.ts";
import { nuid } from "../../nats-base-client/nuid.ts";

const nc = await connect();

const stream = nuid.next();
const subj = nuid.next();
const durable = nuid.next();

const jsm = await nc.jetstreamManager();
await jsm.streams.add({ name: stream, subjects: [subj] });

const js = nc.jetstream();
await js.publish(subj);
await js.publish(subj);
await js.publish(subj);
await js.publish(subj);

const psub = await js.pullSubscribe(subj, {
  mack: true,
  // artificially low ack_wait, to show some messages
  // not getting acked being redelivered
  config: {
    durable_name: durable,
    ack_policy: AckPolicy.Explicit,
    ack_wait: nanos(4000),
  },
});

(async () => {
  for await (const m of psub) {
    console.log(
      `[${m.seq}] ${ 
        m.redelivered ? `- redelivery ${m.info.redeliveryCount}` : ""
      }`
    );
    if (m.seq % 2 === 0) {
      m.ack();
    }
  }
})();

const fn = () => {
  console.log("[PULL]");
  psub.pull({ batch: 1000, expires: 10000 });
};

// do the initial pull
fn();
// and now schedule a pull every so often
const interval = setInterval(fn, 10000); // and repeat every 2s

setTimeout(() => {
  clearInterval(interval);
  nc.drain();
}, 20000);
```

--------------------------------

### Get nsc Tool Help

Source: https://docs.nats.io/using-nats/nats-tools/nsc

Displays the help documentation for the nsc tool, providing an overview of its commands and syntax. This is useful for understanding the full capabilities and options available within the tool.

```shell
nsc help
```

--------------------------------

### Run NATS Box Container

Source: https://docs.nats.io/running-a-nats-service/nats_docker

Starts a NATS Box container with the 'nats' network. This container provides tools to interact with the NATS cluster.

```bash
docker run --network nats --rm -it natsio/nats-box
```

--------------------------------

### Nex Node Firecracker VM Startup

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

This log confirms the successful startup of a Firecracker virtual machine by the Nex node. It includes essential details for debugging and monitoring, such as the gateway IP, tap device name, internal NATS host and port, and the virtual machine's unique ID.

```log
INFO[0000] Machine started                               gateway=192.168.127.1 hosttap=tap0 ip=192.168.127.6 nats_host=192.168.127.1 nats_port=41339 netmask=ffffff00 vmid=cmjg61n52omq8dovolmg
```

--------------------------------

### Configure NATS Stream and Consumers for ORDERS Scenario

Source: https://docs.nats.io/nats-concepts/jetstream/consumers/example_configuration

This configuration uses the `nats` CLI to set up the ORDERS stream and its associated consumers. It defines stream retention policies and consumer delivery modes. The `--storage file` option indicates that messages will be stored on disk. The consumers `NEW` and `DISPATCH` are configured for pull-based delivery with explicit acknowledgments, while `MONITOR` uses none for pub/sub semantics.

```bash
nats stream add ORDERS --subjects "ORDERS.*" --ack --max-msgs=-1 --max-bytes=-1 --max-age=1y --storage file --retention limits --max-msg-size=-1 --discard=old
nats consumer add ORDERS NEW --filter ORDERS.received --ack explicit --pull --deliver all --max-deliver=-1 --sample 100
nats consumer add ORDERS DISPATCH --filter ORDERS.processed --ack explicit --pull --deliver all --max-deliver=-1 --sample 100
nats consumer add ORDERS MONITOR --filter '' --ack none --target monitor.ORDERS --deliver last --replay instant
```

--------------------------------

### C# JetStream Publish and Consume Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This C# example demonstrates creating a JetStream context, creating a stream, publishing messages (both individually and concurrently), and then creating a consumer. The consumer is configured with a maximum number of inflight messages and fetches messages in batches, acknowledging them.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.JetStream;
using NATS.Client.JetStream.Models;

await using var client = new NatsClient();

INatsJSContext js = client.CreateJetStreamContext();

// Create a stream
var streamConfig = new StreamConfig(name: "FOO", subjects: ["foo"]);
await js.CreateStreamAsync(streamConfig);

// Publish a message
{
    PubAckResponse ack = await js.PublishAsync("foo", "Hello, JetStream!");
    ack.EnsureSuccess();
}

// Publish messages concurrently
List<NatsJSPublishConcurrentFuture> futures = new();
for (var i = 0; i < 500; i++)
{
    NatsJSPublishConcurrentFuture future
        = await js.PublishConcurrentAsync("foo", "Hello, JetStream 1!");
    futures.Add(future);
}

foreach (var future in futures)
{
    await using (future)
    {
        PubAckResponse ack = await future.GetResponseAsync();
        ack.EnsureSuccess();
    }
}


// Create a consumer with a maximum 128 inflight messages
INatsJSConsumer consumer = await js.CreateConsumerAsync("FOO", new ConsumerConfig(name: "foo")
{
    MaxWaiting = 128,
});

using CancellationTokenSource cts = new(TimeSpan.FromSeconds(10));

while (cts.IsCancellationRequested == false)
{
    var opts = new NatsJSFetchOpts { MaxMsgs = 10 };
    await foreach (NatsJSMsg<string> msg in consumer.FetchAsync<string>(opts, cancellationToken: cts.Token))
    {
        await msg.AckAsync(cancellationToken: cts.Token);
    }
}
```

--------------------------------

### Configure JetStream with a File

Source: https://docs.nats.io/running-a-nats-service/introduction/running

This command starts the NATS server using a configuration file, enabling JetStream and specifying its data storage directory. The `js.conf` file defines `store_dir` for JetStream data, recommended to be a local directory for performance.

```shell
nats-server -c js.conf
```

```conf
# js.conf
jetstream {
   store_dir=nats
}
```

--------------------------------

### NATS Server Configuration File Example

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This snippet shows the structure of a NATS server configuration file. It includes the operator JWT path, resolver type, and preloaded account JWTs. This is useful for understanding the manual configuration process.

```NATS Server Config
operator: <path to the operator jwt or jwt itself>
resolver: MEMORY
resolver_preload: {
    <public key for an account>: <contents of the account jwt>
    ### add as many accounts as you want
    ...
}
```

```NATS Server Config
operator: /Users/synadia/.nsc/nats/memory/memory.jwt
resolver: MEMORY
resolver_preload: {
ACSU3Q6LTLBVLGAQUONAGXJHVNWGSKKAUA7IY5TB4Z7PLEKSR5O6JTGR: eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJPRFhJSVI2Wlg1Q1AzMlFJTFczWFBENEtTSDYzUFNNSEZHUkpaT05DR1RLVVBISlRLQ0JBIiwiaWF0IjoxNTU2NjU1Njk0LCJpc3MiOiJPRFdaSjJLQVBGNzZXT1dNUENKRjZCWTRRSVBMVFVJWTRKSUJMVTRLM1lER0NHSElXQlZXQkhVWiIsIm5hbWUiOiJBIiwic3ViIjoiQUNTVTNRNkxUTEJWTEdBUVVPTkFHWEpIVk5XR1NLS0FVQTdJWTVUQjRaN1BMRUtTUjVPNkpUR1IiLCJ0eXBlIjoiYWNjb3VudCIsIm5hdHMiOnsibGltaXRzIjp7InN1YnMiOi0xLCJjb25uIjotMSwibGVhZiI6LTEsImltcG9ydHMiOi0xLCJleHBvcnRzIjotMSwiZGF0YSI6LTEsInBheWxvYWQiOi0xLCJ3aWxkY2FyZHMiOnRydWV9fX0._WW5C1triCh8a4jhyBxEZZP8RJ17pINS8qLzz-01o6zbz1uZfTOJGvwSTS6Yv2_849B9iUXSd-8kp1iMXHdoBA
}
```

--------------------------------

### Go NATS Request-Reply Example

Source: https://docs.nats.io/using-nats/developer/sending/replyto

Shows how to implement request-reply semantics in Go using the NATS client library. It covers connecting to the NATS server, publishing a request with a reply subject, and handling the response.

```c
natsConnection      *conn      = NULL;
natsStatus          s          = NATS_OK;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);
// Publish a message and provide a reply subject
if (s == NATS_OK)
    s = natsConnection_PublishRequestString(conn, "request", "reply", "this is the request");

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
```

--------------------------------

### NATS Export Configuration Example (DSL)

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/accounts

An example demonstrating the NATS export configuration using its Domain Specific Language (DSL). This configuration defines accounts and their exported streams and services, including access restrictions to specific accounts.

```plaintext
accounts: {
    A: {
        users: [
            {user: a, password: a}
        ]
        exports: [
            {stream: puba.>}
            {service: pubq.>}
            {stream: b.>, accounts: [B]}
            {service: q.b, accounts: [B]}
        ]
    }
    ...
}
```

--------------------------------

### Operator JWT Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

An example of a decoded operator JWT, showing its structure and key-value pairs. It includes timestamps, identifiers, the operator's name, and configuration details for NATS services.

```json
{
 "iat": 1603473819,
 "iss": "OBU5O5FJ324UDPRBIVRGF7CNEOHGLPS7EYPBTVQZKSBHIIZIB6HD66JF",
 "jti": "57BWRLW67I6JTVYMQAZQF54G2G37DJB5WG5IFIPVYI4PEYNX57ZQ",
 "name": "DEMO",
 "nats": {
  "account_server_url": "nats://localhost:4222",
  "system_account": "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5"
 },
 "sub": "OBU5O5FJ324UDPRBIVRGF7CNEOHGLPS7EYPBTVQZKSBHIIZIB6HD66JF",
 "type": "operator"
}
```

--------------------------------

### NATS CLI Subscriber and Publisher Example

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

Demonstrates how to create a NATS subscriber that listens to all messages and publish a message 'NATS' to the subject 'hello' using the NATS CLI tool. Requires a user credential file for authentication.

```shell
nats sub --creds ~/.nkeys/creds/MyOperator/MyAccount/MyUser.creds ">"
nats pub --creds ~/.nkeys/creds/MyOperator/MyAccount/MyUser.creds hello NATS
```

--------------------------------

### C# NATS Request-Reply Example

Source: https://docs.nats.io/using-nats/developer/sending/replyto

Illustrates the request-reply pattern in C# using the NATS.Net library. It shows how to connect, subscribe, publish a request with a reply subject, and receive the response.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.Core;

await using var client = new NatsClient();

await client.ConnectAsync();

// Create a new inbox for the subscription subject
string inbox = client.Connection.NewInbox();

// Use core API to subscribe to have a more fine-grained control over
// the subscriptions. We use <string> as the type, but we are not
// really interested in the message payload.
await using INatsSub<string> timeSub
    = await client.Connection.SubscribeCoreAsync<string>("time");

Task responderTask = Task.Run(async () =>
{
    await foreach (var msg in timeSub.Msgs.ReadAllAsync())
    {
        // The default serializer uses StandardFormat with Utf8Formatter
        // when formatting DateTimeOffset types.
        await msg.ReplyAsync<DateTimeOffset>(DateTimeOffset.UtcNow);
    }
});

// Subscribe to the inbox with the expected type of the response
await using INatsSub<DateTimeOffset> inboxSub
    = await client.Connection.SubscribeCoreAsync<DateTimeOffset>(inbox);

// The default serializer uses UTF-8 encoding for strings
await client.PublishAsync(subject: "time", replyTo: inbox);

// Read the response from subscription message channel reader
NatsMsg<DateTimeOffset> reply = await inboxSub.Msgs.ReadAsync();

// Print the current time in RFC1123 format taking advantage of the
// DateTimeOffset's formatting capabilities.
Console.WriteLine($"The current date and time is: {reply.Data:R}");

await inboxSub.UnsubscribeAsync();
await timeSub.UnsubscribeAsync();

// make sure the responder task is completed cleanly
await responderTask;

// Output:
// The current date and time is: Tue, 22 Oct 2024 12:21:09 GMT
```

--------------------------------

### NATS Cluster Configuration File Example (HOCON)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This snippet provides an example of a NATS cluster configuration file, specifically for the seed server. It demonstrates how to define cluster settings within a HOCON (Human-Optimized Config Object Notation) format, which is an alternative to command-line arguments.

```hocon
{
  "server": {
    "port": 4222,
    "cluster": {
      "port": 4248
    },
    "name": "test-cluster"
  }
}
```

--------------------------------

### Start Service on Red Leaf Node (NATS CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-leafnodes-docker

Starts a simple NATS service on the 'red' leaf node, listening on localhost:4222. This service acts as a replier to requests on the 'docker-leaf-test' subject, echoing back the received message with a timestamp.

```shell
nats -s localhost:4222 reply docker-leaf-test "At {{Time}}, I received your request: {{Request}}"

```

--------------------------------

### Example: Requesting Stream Info with NATS CLI

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Demonstrates how to use the `nats req` command to request information about a specific stream. This example shows the command to publish a request and the expected JSON response, including error handling for non-existent streams and successful retrieval of stream configuration.

```shell
nats req '$JS.API.STREAM.INFO.nonexisting' ''

```

```json
{
  "type": "io.nats.jetstream.api.v1.stream_info_response",
  "error": {
    "code": 404,
    "description": "stream not found"
  }
}

```

```shell
nats req '$JS.STREAM.INFO.ORDERS' ''

```

```json
{
  "type": "io.nats.jetstream.api.v1.stream_info_response",
  "config": {
    "name": "ORDERS",
  ...
}

```

--------------------------------

### List All Microservices using NATS CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-services

This command lists all deployed microservices, displaying their name, version, ID, and description. It is used to verify that the 'echoservice' has been successfully deployed and is running.

```bash
$ nats micro ls
╭──────────────────────────────────────────────────────────────╮
│                      All Micro Services                      │
├─────────────┬─────────┬────────────────────────┬─────────────┤
│ Name        │ Version │ ID                     │ Description │
├─────────────┼─────────┼────────────────────────┼─────────────┤
│ EchoService │ 1.0.0   │ NsMaTbN7u5ZPUNN47bSEI6 │             │
╰─────────────┴─────────┴────────────────────────┴─────────────╯
```

--------------------------------

### Start NATS Server with Encryption Key from Environment

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/encryption_at_rest

This shell command shows how to start the NATS server with encryption enabled, providing the encryption key via an environment variable. The `JS_KEY` environment variable is set to `"mykey"` and passed to the `nats-server` executable along with the configuration file path.

```shell
JS_KEY="mykey" nats-server -c js.conf
```

--------------------------------

### Start NATS Queue Group Member (Bash)

Source: https://docs.nats.io/nats-concepts/core-nats/queue/queues_walkthrough

Starts a NATS subscriber that automatically joins a queue group. This command utilizes the 'nats reply' command to listen for messages on the 'foo' subject and provide a response, acting as a member of the queue group 'NATS-RPLY-22' by default.

```bash
nats reply foo "service instance A Reply# {{Count}}"
```

```bash
nats reply foo "service instance B Reply# {{Count}}"
```

```bash
nats reply foo "service instance C Reply# {{Count}}"
```

--------------------------------

### Enable JetStream (Shell)

Source: https://docs.nats.io/nats-concepts/jetstream/key-value-store/kv_walkthrough

Restarts the NATS server with JetStream enabled. This is a prerequisite for using the Key/Value store functionality. Assumes a local nats-server installation.

```shell
nats-server -js
```

--------------------------------

### Flush and Ping Example in Python

Source: https://docs.nats.io/using-nats/developer/sending/caches

Demonstrates publishing a message and then calling `flush` with a timeout to ensure the message is processed by the server. This confirms delivery using the PING/PONG mechanism. Requires the NATS Python client library.

```python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

await nc.publish("updates", b'All is Well')

# Sends a PING and wait for a PONG from the server, up to the given timeout.
# This gives guarantee that the server has processed above message.
await nc.flush(timeout=1)
```

--------------------------------

### Python JetStream Publish and Pull Subscribe Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This Python example demonstrates publishing messages to a JetStream stream and then setting up a pull subscriber to fetch and process them. It uses the asyncio library for asynchronous operations and the nats-client library. Messages are published and then fetched in batches.

```python
import asyncio

import nats
from nats.errors import TimeoutError

async def main():
    nc = await nats.connect("localhost")

    # Create JetStream context.
    js = nc.jetstream()

    # Persist messages on 'foo's subject.
    await js.add_stream(name="sample-stream", subjects=["foo"])

    for i in range(0, 10):
        ack = await js.publish("foo", f"hello world: {i}".encode())
        print(ack)

    # Create pull based consumer on 'foo'.
    psub = await js.pull_subscribe("foo", "psub")

    # Fetch and ack messagess from consumer.
    for i in range(0, 10):
        msgs = await psub.fetch(1)
        for msg in msgs:
            print(msg)

    await nc.close()

if __name__ == '__main__':
    asyncio.run(main())
```

--------------------------------

### Flush and Ping Example in Go

Source: https://docs.nats.io/using-nats/developer/sending/caches

Demonstrates publishing a message and then using `FlushTimeout` to ensure the message has been sent to the NATS server. This uses the PING/PONG mechanism for confirmation. Requires the NATS Go client library.

```go
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Just to not collide using the demo server with other users.
subject := nats.NewInbox()

if err := nc.Publish(subject, []byte("All is Well")); err != nil {
    log.Fatal(err)
}
// Sends a PING and wait for a PONG from the server, up to the given timeout.
// This gives guarantee that the server has processed the above message.
if err := nc.FlushTimeout(time.Second); err != nil {
    log.Fatal(err)
}
```

--------------------------------

### NATS Server Log File Redirection Example

Source: https://docs.nats.io/running-a-nats-service/configuration/logging

Shows how to redirect NATS server's debug and trace output to a specified log file named 'nats.log'.

```bash
nats-server -DV -m 8222 -l nats.log
```

--------------------------------

### Create NATS Echo Service in Go

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This Go code defines a simple NATS microservice named 'EchoService'. It connects to a NATS server, sets up an 'echo' endpoint that responds to requests with the same data it receives, and runs until the context is done. It requires the Go NATS client SDK.

```go
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
	services "github.com/nats-io/nats.go/micro"
)

func main() {
	ctx := context.Background()

	natsUrl := os.Getenv("NATS_URL")
	if len(strings.TrimSpace(natsUrl)) == 0 {
		natsUrl = nats.DefaultURL
	}
	fmt.Printf("Echo service using NATS url '%s'\n", natsUrl)
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		panic(err)
	}

	// request handler
	echoHandler := func(req services.Request) {
		req.Respond(req.Data())
	}

	fmt.Println("Starting echo service")

	_, err = services.AddService(nc, services.Config{
		Name:    "EchoService",
		Version: "1.0.0",
		// base handler
		Endpoint: &services.EndpointConfig{
			Subject: "svc.echo",
			Handler: services.HandlerFunc(echoHandler),
		},
	})

	if err != nil {
		panic(err)
	}

	<-ctx.Done()
}
```

--------------------------------

### NATS JetStream Stream and Consumer Management (C#)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

This C# example demonstrates creating, updating, listing, and deleting streams and consumers within NATS JetStream. It includes examples for creating durable consumers and retrieving information about existing streams and consumers. Requires the 'NATS.Net' NuGet package.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.JetStream;
using NATS.Client.JetStream.Models;

await using var client = new NatsClient();

INatsJSContext js = client.CreateJetStreamContext();

// Create a stream
var streamConfig = new StreamConfig(name: "example-stream", subjects: ["example-subject"]) 
{
    MaxBytes = 1024,
};
await js.CreateStreamAsync(streamConfig);

// Update the stream
var streamConfigUpdated = streamConfig with { MaxBytes = 2048 };
await js.UpdateStreamAsync(streamConfigUpdated);

// Create a durable consumer
await js.CreateConsumerAsync("example-stream", new ConsumerConfig("example-consumer-name"));

// Get information about all streams
await foreach (var stream in js.ListStreamsAsync())
{
    Console.WriteLine($"stream name: {stream.Info.Config.Name}");
}

// Get information about all consumers in a stream
await foreach (var consumer in js.ListConsumersAsync("example-stream"))
{
    Console.WriteLine($"consumer name: {consumer.Info.Config.Name}");
}

// Delete a consumer
await js.DeleteConsumerAsync("example-stream", "example-consumer-name");

// Delete a stream
await js.DeleteStreamAsync("example-stream");

// Output:
// stream name: example-stream
// consumer name: example-consumer-name
```

--------------------------------

### NATS Server Subscription and Reply

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This Go code snippet demonstrates how to subscribe to a NATS topic ('test') and respond to incoming messages. It uses the 'nats.io/nats.go' library and logs received messages. The server continuously listens for messages and replies with 'pong from nkeys based server'.

```go
log.Fatal(err)
    }
    nc.Subscribe("test", func(m *nats.Msg){
        log.Printf("[Received] %q, replying... \n", string(m.Data))
        m.Respond([]byte("pong from nkeys based server"))
    })

    select {}
```

--------------------------------

### Create and Save NATS Contexts

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli

Demonstrates how to create a new NATS context with specific server details and descriptions, or save an existing configuration. It also shows how to select a default context for easier management.

```shell
nats context create my_context_name
nats context edit my_context_name
nats context save example --server nats://nats.example.net:4222 --description 'Example.Net Server'
nats context save local --server nats://localhost:4222 --description 'Local Host' --select
```

--------------------------------

### nats-top Command-Line Usage

Source: https://docs.nats.io/using-nats/nats-tools/nats_top

Demonstrates the command-line syntax for running nats-top with various optional arguments. These arguments allow customization of the server to monitor, the HTTP monitor port, the number of connections to display, the refresh delay, and the sorting criteria.

```bash
nats-top [-s server] [-m monitor] [-n num_connections] [-d delay_in_secs] [-sort by]
```

--------------------------------

### Deploy Service with Environment Variable using Nex CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-services

This command deploys the 'echoservice' to the first available Nex node and sets the 'NATS_URL' environment variable. It requires a statically linked binary and may create publisher and encryption keys if they don't exist. The output includes the accepted workload ID and node information.

```bash
$ nex devrun ../examples/echoservice/echoservice nats_url=nats://192.168.127.1:4222
Reusing existing issuer account key: /home/kevin/.nex/issuer.nk
Reusing existing publisher xkey: /home/kevin/.nex/publisher.xk
🚀 Workload 'echoservice' accepted. You can now refer to this workload with ID: cmji29n52omrb71g07a0 on node NBS3Y3NWXLTFNC73XMVD6USFJF2H5QXTLEJQNOPEBPYDUDVB5YYYZOGI
```

--------------------------------

### Get Value from KV Bucket (Shell)

Source: https://docs.nats.io/nats-concepts/jetstream/key-value-store/kv_walkthrough

Retrieves the value associated with a given key from a KV bucket. This operation is referred to as 'getting' a value. The command requires the bucket name and the key.

```shell
nats kv get my-kv Key1
```

--------------------------------

### JavaScript NATS Subscribe, Publish, and Drain Example

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Shows how to subscribe to a NATS subject, publish messages, and then drain the subscription to process any remaining messages. This example uses the NATS JavaScript client.

```javascript
const sub = nc.subscribe(subj, { callback: (_err, _msg) => {} });
nc.publish(subj);
nc.publish(subj);
nc.publish(subj);
await sub.drain();
```

--------------------------------

### Start Third NATS Server in Cluster

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Command to start a third NATS server, which also joins the cluster. It uses a different client port and cluster address but connects to the same seed server's cluster address.

```bash
nats-server -p 6222 -cluster nats://localhost:6248 -routes nats://localhost:4248 -D
```

--------------------------------

### Python NKey Authentication

Source: https://docs.nats.io/using-nats/developer/connecting/nkey

Provides a Python example for connecting to NATS using NKey authentication. It specifies the path to the NKey seed file and includes an error callback function. Ensure the correct path to your NKey seed is provided.

```Python
nc = NATS()

async def error_cb(e):
    print("Error:", e)

await nc.connect("nats://localhost:4222",
                 nkeys_seed="./path/to/nkeys/user.nk",
                 error_cb=error_cb,
                 )

# Do something with the connection

await nc.close()
```

--------------------------------

### Nex Node Execution Engine Status

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

This log indicates that the NATS execution engine within the Nex node is ready to receive commands. It provides the unique ID and version of the engine, crucial for identifying and managing the node's operational state.

```log
INFO[0000] NATS execution engine awaiting commands       id=NCOBPU3MCEA7LF6XADFD4P74CHW2OL6GQZYPPRRNPDSBNQ5BJPFHHQB5 version=0.0.1
```

--------------------------------

### Create NATS Windows Service with Explicit Log File

Source: https://docs.nats.io/running-a-nats-service/introduction/windows_srv

This `sc.exe create` command demonstrates how to install the NATS server service while specifying an explicit log file path. This is recommended for easier debugging and log management, especially when running under restricted user accounts.

```shell
sc.exe create nats-server binPath= "%NATS_PATH%\nats-server.exe --log C:\temp\nats-server.log [other flags]"
```

--------------------------------

### C#: Put Data in NATS KV Store

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Provides a C# example for putting data into the NATS KV store. It supports generic types for values and allows for custom serializers and cancellation tokens.

```C#
// dotnet add package NATS.Net

// Put a value into the bucket using the key
// returns revision number
ValueTask<ulong> PutAsync<T>(string key, T value, INatsSerialize<T>? serializer = default, CancellationToken cancellationToken = default);
```

--------------------------------

### NATS Server Startup Log (Basic TLS)

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/tls

This log output indicates that the NATS server has started successfully and that TLS is now required for all client connections. This is a confirmation of the TLS configuration being active.

```log
[21417] 2019/05/16 11:21:19.801539 [INF] Starting nats-server version 2.0.0
[21417] 2019/05/16 11:21:19.801621 [INF] Git commit [not set]
[21417] 2019/05/16 11:21:19.801777 [INF] Listening for client connections on 0.0.0.0:4222
[21417] 2019/05/16 11:21:19.801782 [INF] TLS required for client connections
[21417] 2019/05/16 11:21:19.801785 [INF] Server id is ND6ZZDQQDGKYQGDD6QN2Y26YEGLTH6BMMOJZ2XJB2VASPVII3XD6RFOQ
[21417] 2019/05/16 11:21:19.801787 [INF] Server is ready
```

--------------------------------

### NATS Cluster Server Startup Command

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Command to start a NATS server that joins an existing cluster. It specifies the client port, the cluster listen address, and the routes to connect to (in this case, the seed server). The '-D' flag enables debug logging.

```bash
nats-server -p 5222 -cluster nats://localhost:5248 -routes nats://localhost:4248 -D
```

--------------------------------

### Subject Naming - Pragmatic Example

Source: https://docs.nats.io/nats-concepts/subjects

A practical example of a NATS subject name that encodes business semantics effectively. It includes elements like order type, store, and specific order ID, suitable for real-world applications.

```shell
orders.online.store123.order171711
```

--------------------------------

### Fetch Message by Stream Sequence - Response Example

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

An example of the response received when fetching a message by its stream sequence. The response includes message details such as subject, sequence number, base64 encoded data, and timestamp.

```json
{
  "type": "io.nats.jetstream.api.v1.stream_msg_get_response",
  "message": {
    "subject": "x",
    "seq": 1,
    "data": "aGVsbG8=",
    "time": "2020-05-06T13:18:58.115424+02:00"
  }
}
```

--------------------------------

### Flush and Ping Example in JavaScript

Source: https://docs.nats.io/using-nats/developer/sending/caches

Shows how to publish a message and use the `flush` method to confirm its round trip completion. The `flush` operation internally uses PING/PONG to verify server acknowledgement. Requires the NATS JavaScript client library.

```javascript
const start = Date.now();
nc.flush().then(() => {
  t.log("round trip completed in", Date.now() - start, "ms");
});
```

--------------------------------

### Connect to Default NATS Server in Python

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Connects to the default NATS server using the Python NATS client library. This asynchronous example includes connecting, a placeholder for performing actions, and closing the connection.

```python
nc = NATS()
await nc.connect()

# Do something with the connection

await nc.close()
```

--------------------------------

### C NATS Subscription Callback Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This C example defines a callback function `onMsg` for NATS message delivery. It prints the received message subject and data, tracks the count of received messages, and calculates the elapsed time for processing. This callback is designed for asynchronous message handling in C NATS applications.

```c
#include "examples.h"

static const char *usage = "\
-gd            use global message delivery thread pool\
-sync          receive synchronously (default is asynchronous)\
-pull          use pull subscription\
-fc            enable flow control\
-count         number of expected messages\
";

static void
onMsg(natsConnection *nc, natsSubscription *sub, natsMsg *msg, void *closure)
{
    if (print)
        printf("Received msg: %s - %.*s\n",
               natsMsg_GetSubject(msg),
               natsMsg_GetDataLength(msg),
               natsMsg_GetData(msg));

    if (start == 0)
        start = nats_Now();

    // We should be using a mutex to protect those variables since
    // they are used from the subscription's delivery and the main
    // threads. For demo purposes, this is fine.
    if (++count == total)
        elapsed = nats_Now() - start;

    // Since this is auto-ack callback, we don't need to ack here.
    natsMsg_Destroy(msg);
}

static void
asyncCb(natsConnection *nc, natsSubscription *sub, natsStatus err, void *closure)
{
```

--------------------------------

### Initialize NATS Operator and System Account

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This command initializes a new NATS operator named 'DEMO' and creates a system account and user within it. It generates necessary keys and stores user credentials in a local file.

```shell
nsc add operator -n DEMO --sys

```

--------------------------------

### Flush and Ping Example in C#

Source: https://docs.nats.io/using-nats/developer/sending/caches

Shows how to publish a message and then use `PingAsync` to ensure the message has been processed by the server. The `PingAsync` method verifies that the server has acknowledged the outstanding messages. Requires the NATS.Net NuGet package.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient();

await client.PublishAsync("updates", "All is well");

// Sends a PING and wait for a PONG from the server.
// This gives a guarantee that the server has processed the above message
// since the underlining TCP connection sends and receives messages in order.
await client.PingAsync();
```

--------------------------------

### Flow Control Subject Example in NATS JetStream

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Provides an example of a flow control subject in NATS JetStream, used for managing message delivery rates between JetStream and consumers.

```text
$JS.FC.orders.6i5h0GiQ.ep3Y
```

--------------------------------

### NATS Server CONNECT Message with NKEY

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

An example of a NATS server trace showing a client's CONNECT message, including the `nkey` and `sig` (signature) fields, indicating NKEY authentication.

```text
[95184] 2020/10/26 12:15:44.350577 [TRC] [::1]:55551 - cid:2 - <<- [CONNECT {
    "echo": true,
    "headers": true,
    "lang": "go",
    "name": "NATS CLI",
    "nkey": "UC435ZYS52HF72E2VMQF4GO6CUJOCHDUUPEBU7XDXW5AQLIC6JZ46PO5",
    "no_responders": true,
    "pedantic": false,
    "protocol": 1,
    "sig": "lopzgs98JBQYyRdw1zT_BoBpSFRDCfTvT4le5MYSKrt0IqGWZ2OXhPW1J_zo2_sBod8XaWgQc9oWohWBN0NdDg",
    "tls_required": false,
    "verbose": false,
    "version": "1.11.0"
}]
```

--------------------------------

### JetStream Subscriptions in C (NATS)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Demonstrates JetStream subscription functionalities using the NATS C client. It covers setting up different subscription types (pull, asynchronous, synchronous), enabling flow control, and handling messages with callbacks. This example requires the NATS C library and includes argument parsing for configuration.

```c
#include "examples.h"

static const char *usage = "\"
-gd            use global message delivery thread pool\n"
"-sync          receive synchronously (default is asynchronous)\n"
"-pull          use pull subscription\n"
"-fc            enable flow control\n"
"-count         number of expected messages\n";

static void
onMsg(natsConnection *nc, natsSubscription *sub, natsMsg *msg, void *closure)
{
    if (print)
        printf("Received msg: %s - %.*s\n",
               natsMsg_GetSubject(msg),
               natsMsg_GetDataLength(msg),
               natsMsg_GetData(msg));

    if (start == 0)
        start = nats_Now();

    // We should be using a mutex to protect those variables since
    // they are used from the subscription's delivery and the main
    // threads. For demo purposes, this is fine.
    if (++count == total)
        elapsed = nats_Now() - start;

    // Since this is auto-ack callback, we don't need to ack here.
    natsMsg_Destroy(msg);
}

static void
asyncCb(natsConnection *nc, natsSubscription *sub, natsStatus err, void *closure)
{
    printf("Async error: %u - %s\n", err, natsStatus_GetText(err));

    natsSubscription_GetDropped(sub, (int64_t*) &dropped);
}

int main(int argc, char **argv)
{
    natsConnection      *conn  = NULL;
    natsStatistics      *stats = NULL;
    natsOptions         *opts  = NULL;
    natsSubscription    *sub   = NULL;
    natsMsg             *msg   = NULL;
    jsCtx               *js    = NULL;
    jsErrCode           jerr   = 0;
    jsOptions           jsOpts;
    jsSubOptions        so;
    natsStatus          s;
    bool                delStream = false;

    opts = parseArgs(argc, argv, usage);

    printf("Created %s subscription on '%s'.\n",
        (pull ? "pull" : (async ? "asynchronous" : "synchronous")), subj);

    s = natsOptions_SetErrorHandler(opts, asyncCb, NULL);

    if (s == NATS_OK)
        s = natsConnection_Connect(&conn, opts);

    if (s == NATS_OK)
        s = jsOptions_Init(&jsOpts);

    if (s == NATS_OK)
        s = jsSubOptions_Init(&so);
    if (s == NATS_OK)
    {
        so.Stream = stream;
        so.Consumer = durable;
        if (flowctrl)
        {
            so.Config.FlowControl = true;
            so.Config.Heartbeat = (int64_t)1E9;
        }
    }

    if (s == NATS_OK)
        s = natsConnection_JetStream(&js, conn, &jsOpts);

    if (s == NATS_OK)
    {
        jsStreamInfo    *si = NULL;

        // First check if the stream already exists.
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_NOT_FOUND)
        {
            jsStreamConfig  cfg;

            // Since we are the one creating this stream, we can delete at the end.
            delStream = true;

            // Initialize the configuration structure.
            jsStreamConfig_Init(&cfg);
            cfg.Name = stream;
            // Set the subject
            cfg.Subjects = (const char*[1]){subj};
            cfg.SubjectsLen = 1;
            // Make it a memory stream.

```

--------------------------------

### Start NATS with WebSocket on Docker - Bash Command

Source: https://docs.nats.io/running-a-nats-service/configuration/websocket/websocket_conf

Illustrates the Docker command to run a NATS server with a custom configuration file, enabling WebSocket support. The command mounts a local configuration file into the container and maps the WebSocket port.

```bash
docker run -it --rm  -v /tmp:/container -p 8080:8080 nats -c /container/nats.conf

```

--------------------------------

### Subject Naming - Less Useful Example

Source: https://docs.nats.io/nats-concepts/subjects

An example of a NATS subject name that is overly complex and encodes too many technical details, making it less useful and harder to manage. It highlights the importance of concise and semantic naming.

```shell
orders.online.us.server42.ccpayment.premium.store123.electronics.deliver-dhl.order171711.create
```

--------------------------------

### Subscribe to Remote NATS Server

Source: https://docs.nats.io/running-a-nats-service/clients

Command to subscribe to messages from a remote NATS server. This example demonstrates subscribing to a demo NATS server. It requires specifying the NATS URL of the remote server.

```shell
nats sub -s nats://demo.nats.io ">"
```

--------------------------------

### Enable NATS Monitoring via Command Line

Source: https://docs.nats.io/running-a-nats-service/configuration/monitoring

Starts the NATS server with the monitoring HTTP server enabled on a specified port. This is a quick way to activate monitoring for basic statistics.

```bash
nats-server -m 8222
```

--------------------------------

### Flush and Ping Example in C

Source: https://docs.nats.io/using-nats/developer/sending/caches

Demonstrates publishing a string message and using `natsConnection_FlushTimeout` to ensure the message is sent and acknowledged by the server. This confirms delivery via the PING/PONG mechanism. Requires the NATS C client library.

```c
natsConnection      *conn      = NULL;
natsStatus          s          = NATS_OK;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);

// Send a request and wait for up to 1 second
if (s == NATS_OK)
    s = natsConnection_PublishString(conn, "foo", "All is Well");

// Sends a PING and wait for a PONG from the server, up to the given timeout.
// This gives guarantee that the server has processed the above message.
if (s == NATS_OK)
    s = natsConnection_FlushTimeout(conn, 1000);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
```

--------------------------------

### Nex Node Encrypted Run Request Key

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

This log output shows the public Xkey generated by the Nex node, which serves as the recipient for encrypted data, specifically environment variables for workload deployment requests. The 'nex' CLI automatically handles the encryption process using this key.

```log
INFO[0000] Use this key as the recipient for encrypted run requests  public_xkey=XDJJVJLRTWBIOHEEUPSNAAUACO6ZRW4WP65MXMGOX2WBGNCELLST5TWI
```

--------------------------------

### Basic NATS TLS Configuration (Command Line)

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/tls

This snippet demonstrates how to enable and configure TLS for a NATS server directly from the command line using flags. This is an alternative to using a configuration file for basic TLS setup.

```shell
nats-server --tls --tlscert=./server-cert.pem --tlskey=./server-key.pem
```

--------------------------------

### Start NATS Server with JetStream in Docker

Source: https://docs.nats.io/running-a-nats-service/nats_docker/jetstream_docker

This command starts a NATS server with JetStream enabled in a Docker container. It forwards the local port 4222 to the container's default client connection port.

```shell
docker run -p 4222:4222 nats -js
```

--------------------------------

### Flush and Ping Example in Ruby

Source: https://docs.nats.io/using-nats/developer/sending/caches

Illustrates publishing a message and using `flush` to confirm that it has been sent and processed. The `flush` callback ensures that the PING/PONG interaction has completed. Requires the NATS Ruby client gem.

```ruby
require 'nats/client'
require 'fiber'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  nc.subscribe("updates") do |msg|
    puts msg
  end

  nc.publish("updates", "All is Well")

  nc.flush do
    # Sends a PING and wait for a PONG from the server, up to the given timeout.
    # This gives guarantee that the server has processed above message at this point.
  end
end
```

--------------------------------

### NATS Subscription Callback (C)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

A C example demonstrating how to subscribe to NATS messages and define a callback function (`onMsg`) to process incoming messages. It includes proper resource management for connections and subscriptions.

```c
static void
onMsg(natsConnection *conn, natsSubscription *sub, natsMsg *msg, void *closure)
{
    printf("Received msg: %s - %.*s\n",
           natsMsg_GetSubject(msg),
           natsMsg_GetDataLength(msg),
           natsMsg_GetData(msg));

    // Need to destroy the message!
    natsMsg_Destroy(msg);
}


(...)

natsConnection      *conn = NULL;
natsSubscription    *sub  = NULL;
natsStatus          s;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);
if (s == NATS_OK)
    s = natsConnection_Subscribe(&sub, conn, "time.*.east", onMsg, NULL);

(...)


// Destroy objects that were created
natsSubscription_Destroy(sub);
natsConnection_Destroy(conn);
```

--------------------------------

### Run First NATS Leaf Node (Docker)

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-leafnodes-docker

Launches the first NATS leaf node in a Docker container. It maps ports, mounts the configuration and credentials files, and specifies the NATS server entry point. This command is used for the 'red' user's leaf node.

```shell
docker run  -p 4222:4222 -v leafnode.conf:/leafnode.conf -v /etc/ssl/cert.pem:/etc/ssl/cert.pem -v default-red.creds:/ngs.creds  nats:latest -c /leafnode.conf

```

--------------------------------

### NATS Configuration: String and Number Handling

Source: https://docs.nats.io/running-a-nats-service/configuration

Shows how NATS configuration handles strings and numbers, including support for units like 'K' and 'KB'. It specifically addresses potential issues with string values starting with digits and how to resolve them by quoting.

```shell
listen: 127.0.0.1:4222
authorization {
    # Bad - Number parsing error
    token: 3secret
}

# Fixed Config:
listen: 127.0.0.1:4222
authorization {
    # Good
    token: "3secret"
}
```

--------------------------------

### View NATS JetStream Stream Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/replication

Retrieve detailed information about a specific NATS JetStream stream, including its configuration, state, and source or mirror information. This command is useful for debugging and understanding stream setup.

```shell
nats stream info ARCHIVE
```

```shell
$ nats stream info REPORT
```

--------------------------------

### Statically Compile Go Service

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This Go command builds a statically linked executable of the service. Static compilation is necessary for deployment in environments like Nex, ensuring that all dependencies are included in the binary. The flags disable CGO and embed the Go build ID.

```bash
$ CGO_ENABLED=0 go build -tags netgo -ldflags '-extldflags "-static"'
```

--------------------------------

### Drain NATS Connection (Python)

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Provides an asynchronous example of draining a NATS connection in Python. It demonstrates setting up a subscriber with a handler that can check if the connection is draining. The example also shows making requests and then gracefully draining the connection using `await nc.drain()`.

```Python
import asyncio
from nats.aio.client import Client as NATS

async def example(loop):
    nc = NATS()

    await nc.connect("nats://127.0.0.1:4222", loop=loop)

    async def handler(msg):
        print("[Received] ", msg)
        await nc.publish(msg.reply, b'I can help')

        # Can check whether client is in draining state
        if nc.is_draining:
            print("Connection is draining")

    await nc.subscribe("help", "workers", cb=handler)
    await nc.flush()

    requests = []
    for i in range(0, 10):
        request = nc.request("help", b'help!', timeout=1)
        requests.append(request)

    # Wait for all the responses
    responses = []
    responses = await asyncio.gather(*requests)

    # Gracefully close the connection.
    await nc.drain()

    print("Received {} responses".format(len(responses)))

```

--------------------------------

### Start Lightweight Docker Container

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-docker-python

Initiates a basic Docker container using the python:3.8-slim-buster image with bash as the entrypoint. This is a foundational step for running Python applications in an isolated environment.

```shell
docker run --entrypoint /bin/bash -it python:3.8-slim-buster
```

--------------------------------

### Publish Messages with NATS in JavaScript

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This JavaScript example demonstrates publishing messages to NATS subjects. It shows how to connect to a NATS server and publish messages with specific subject names.

```javascript
import { connect } from 'nats';

async function publishMessages() {
    const nc = await connect({
        servers: ["nats://demo.nats.io:4222"]
    });
    console.log(`Connected to NATS at ${nc.url.host}...`);

    // Publish messages
    nc.publish('time.us.east');
    nc.publish('time.us.central');
    nc.publish('time.us.mountain');
    nc.publish('time.us.west');

    console.log('Published 4 messages.');

    await nc.close();
}

publishMessages();
```

--------------------------------

### Create and Configure NATS User with JWT and Credentials (Go)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Demonstrates the process of generating a user NKEY pair, creating a user JWT signed by an account key, and configuring NATS client options for authentication using either JWT directly or a generated credentials file. This snippet is useful for dynamically provisioning users in Go applications.

```go
func GetAccountSigningKey() nkeys.KeyPair {
    // Content of the account signing key seed can come from a file or an environment variable as well
    accSeed := []byte("SAAJGCAHPHHM6AVJJWQ2YAS3I4NETXMWVQSTCQMJ7VVTGAJF5UCN3IX7J4")
    accountSigningKey, err := nkeys.ParseDecoratedNKey(accSeed)
    if err != nil {
        panic(err)
    }
    return accountSigningKey
}

func RequestUser() {
    // Setup! Obtain the account signing key!
    accountPublicKey := GetAccountPublicKey()
    accountSigningKey := GetAccountSigningKey()
    userPublicKey, userSeed, userKeyPair := generateUserKey()
    userJWT := generateUserJWT(userPublicKey, accountPublicKey, accountSigningKey)
    // userJWT and userKeyPair can be used in conjunction with this nats.Option
    var jwtAuthOption nats.Option
    jwtAuthOption = nats.UserJWT(func() (string, error) {
            return userJWT, nil
        },
        func(bytes []byte) ([]byte, error) {
            return userKeyPair.Sign(bytes)
        },
    )
    // Alternatively you can create a creds file and use it as nats.Option
    credsContent, err := jwt.FormatUserConfig(userJWT, userSeed);
    if err != nil {
        panic(err)
    }
    ioutil.WriteFile("my.creds", credsContent, 0644)
    jwtAuthOption = nats.UserCredentials("my.creds")
    // use in a connection as desired
    nc, err := nats.Connect("nats://localhost:4222", jwtAuthOption)
    // ...
}
```

--------------------------------

### Get Nex Node Information using Nex CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-services

This command retrieves detailed information about a specific Nex node, including its ID, version, uptime, tags, memory statistics, and running workloads. This is useful for monitoring and debugging.

```bash
$ nex node info NBS3Y3NWXLTFNC73XMVD6USFJF2H5QXTLEJQNOPEBPYDUDVB5YYYZOGI
NEX Node Information

         Node: NBS3Y3NWXLTFNC73XMVD6USFJF2H5QXTLEJQNOPEBPYDUDVB5YYYZOGI
         Xkey: XASQSWNSIKHM5MDKDOGPSPGBA3V6JMETMIJK2YTXKAJZNMAFKXER5RUK
      Version: 0.0.1
       Uptime: 8m47s
         Tags: nex.arch=amd64, nex.cpucount=8, nex.os=linux, simple=true

Memory in kB:

           Free: 33,545,884
      Available: 56,529,644
          Total: 63,883,232

Workloads:

             Id: cmji29n52omrb71g07a0
        Healthy: true
        Runtime: 8m47s
           Name: echoservice
    Description: Workload published in devmode
```

--------------------------------

### Create a NATS Object Store Bucket

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This command creates a new object store bucket named 'myobjbucket'. Buckets are fundamental to organizing objects in NATS Object Store. The output confirms the bucket creation and its initial status.

```shell
nats object add myobjbucket
```

--------------------------------

### Add NATS Helm Repository

Source: https://docs.nats.io/running-a-nats-service/nats-kubernetes

Registers the official NATS Helm chart repository. This command is a prerequisite for installing the NATS chart on Kubernetes.

```shell
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
```

--------------------------------

### Generate Replication Graphviz Report using NATS CLI

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/replication

This command generates a GraphViz formatted report detailing the replication setup between NATS streams. The output can be used to visualize stream dependencies and replication flow.

```shell
nats s report --dot replication.dot
```

--------------------------------

### Connect to Default NATS Server in JavaScript

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Demonstrates connecting to the default NATS server asynchronously using the JavaScript NATS client. It shows the connection initiation, a placeholder for usage, and the asynchronous closing of the connection.

```javascript
const nc = await connect();
// Do something with the connection
doSomething();
// When done close it
await nc.close();
```

--------------------------------

### Server Configuration: Basic Mappings

Source: https://docs.nats.io/nats-concepts/subject_mapping

An example of NATS server configuration file snippet defining basic subject mappings. These mappings apply to the default account and are crucial for message routing.

```yaml
server_name: "hub"
cluster: { name: "hub" }
mappings: {
    orders.flush  orders.central.flush 
	ors.* orders.central.{{wildcard(1)}}
}
```

--------------------------------

### View a Private Key File

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

This command displays the content of a private key file (`.nk`). The content is a string representing the encoded private key, starting with 'S' for seed.

```bash
cat ~/.nkeys/keys/U/DB/UDBD5FNQPSLIO6CDMIS5D4EBNFKYWVDNULQTFTUZJXWFNYLGFF52VZN7.nk
```

--------------------------------

### Go JetStream Push Subscribe Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Demonstrates producing and consuming messages using NATS JetStream in Go. It covers asynchronous publishing, various subscription types (async, sync, queue, manual ack, ChanSubscribe), and metadata retrieval. Requires a running NATS server.

```go
func ExampleJetStream() {
	nc, err := nats.Connect("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
	})

	js.Publish("foo", []byte("Hello JS!"))

	// Publish messages asynchronously.
	for i := 0; i < 500; i++ {
		js.PublishAsync("foo", []byte("Hello JS Async!"))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Create async consumer on subject 'foo'. Async subscribers
	// ack a message once exiting the callback.
	js.Subscribe("foo", func(msg *nats.Msg) {
		meta, _ := msg.Metadata()
		fmt.Printf("Stream Sequence  : %v\n", meta.Sequence.Stream)
		fmt.Printf("Consumer Sequence: %v\n", meta.Sequence.Consumer)
	})

	// Async subscriber with manual acks.
	js.Subscribe("foo", func(msg *nats.Msg) {
		msg.Ack()
	}, nats.ManualAck())

	// Async queue subscription where members load balance the
	// received messages together.
	// If no consumer name is specified, either with nats.Bind() 
	// or nats.Durable() options, the queue name is used as the
	// durable name (that is, as if you were passing the 
	// nats.Durable(<queue group name>) option.
	// It is recommended to use nats.Bind() or nats.Durable() 
	// and preferably create the JetStream consumer beforehand 
	// (using js.AddConsumer) so that the JS consumer is not 
	// deleted on an Unsubscribe() or Drain() when the member 
	// that created the consumer goes away first.
	// Check Godoc for the QueueSubscribe() API for more details.
	js.QueueSubscribe("foo", "group", func(msg *nats.Msg) {
		msg.Ack()
	}, nats.ManualAck())

	// Subscriber to consume messages synchronously.
	sub, _ := js.SubscribeSync("foo")
	msg, _ := sub.NextMsg(2 * time.Second)
	msg.Ack()

	// We can add a member to the group, with this member using
	// the synchronous version of the QueueSubscribe.
	sub, _ = js.QueueSubscribeSync("foo", "group")
	msg, _ = sub.NextMsg(2 * time.Second)
	msg.Ack()

	// ChanSubscribe
	msgCh := make(chan *nats.Msg, 8192)
	sub, _ = js.ChanSubscribe("foo", msgCh)

	select {
	case msg := <-msgCh:
		fmt.Println("[Received]", msg)
	case <-time.After(1 * time.Second):
	}
}

```

--------------------------------

### Flush and Ping Example in Java

Source: https://docs.nats.io/using-nats/developer/sending/caches

Illustrates sending a message and then flushing the connection to ensure delivery. The `flush` method with a timeout confirms that the server has processed the message. Requires the NATS Java client library.

```java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

nc.publish("updates", "All is Well".getBytes(StandardCharsets.UTF_8));
nc.flush(Duration.ofSeconds(1)); // Flush the message queue

nc.close();
```

--------------------------------

### Connect to Default NATS Server in Ruby

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Connects to the default NATS server using the Ruby NATS client library. The example uses a block-based approach to manage the connection and includes placeholders for operations and closing the connection.

```ruby
require 'nats/client'

NATS.start do |nc|
   # Do something with the connection

   # Close the connection
   nc.close
end
```

--------------------------------

### Stream Info Endpoint Example

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Demonstrates how to retrieve information about a specific stream, including potential error responses.

```APIDOC
## GET $JS.API.STREAM.INFO.<stream_name>

### Description
Retrieves information about a specific stream.

### Method
GET

### Endpoint
$JS.API.STREAM.INFO.<stream_name>

### Parameters
#### Path Parameters
- **stream_name** (string) - Required - The name of the stream to retrieve information for.

#### Query Parameters
None

#### Request Body
None

### Request Example
```shell
nats req '$JS.API.STREAM.INFO.my_stream' ''
```

### Response
#### Success Response (200)
- **type** (string) - The type of the response payload, e.g., "io.nats.jetstream.api.v1.stream_info_response".
- **config** (object) - The configuration of the stream.
- **state** (object) - The current state of the stream.

#### Error Response (404)
- **type** (string) - The type of the error response, e.g., "io.nats.jetstream.api.v1.stream_info_response".
- **error** (object) - An object containing error details.
  - **code** (integer) - The error code (e.g., 404).
  - **description** (string) - A description of the error (e.g., "stream not found").

#### Response Example (Success)
```json
{
  "type": "io.nats.jetstream.api.v1.stream_info_response",
  "config": {
    "name": "ORDERS",
    "subjects": ["ORDERS.*"],
    "retention": "LATEST",
    "storage": "MEMORY",
    "max_age": 0,
    "max_bytes": 0,
    "replicas": 1,
    "template_limits": {
      "max_consumers": 0
    }
  },
  "state": {
    "name": "ORDERS",
    "domain": "",
    "work_queue": "",
    "messages": 0,
    "bytes": 0,
    "first_seq": 0,
    "first_ts": "0001-01-01T00:00:00Z",
    "last_seq": 0,
    "last_ts": "0001-01-01T00:00:00Z",
    "consumer_seq": 0,
    "consumers": [
      {
        "name": "consumer1",
        "delivered": {"messages": 0, "bytes": 0},
        "ack_floor": 0
      }
    ],
    "num_redelivered": 0,
    "info": {
      "name": "ORDERS",
      "domain": "",
      "work_queue": "",
      "messages": 0,
      "bytes": 0,
      "first_seq": 0,
      "first_ts": "0001-01-01T00:00:00Z",
      "last_seq": 0,
      "last_ts": "0001-01-01T00:00:00Z",
      "consumer_seq": 0,
      "consumers": [],
      "num_redelivered": 0,
      "acknowledged_by_self": 0
    }
  }
}
```

#### Response Example (Not Found)
```json
{
  "type": "io.nats.jetstream.api.v1.stream_info_response",
  "error": {
    "code": 404,
    "description": "stream not found"
  }
}
```
```

--------------------------------

### Run NATS Publisher Example

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-docker-python

Executes the nats-pub.py script to publish a message with the payload 'world' to the 'hello' subject on the global NATS server using TLS and NGS credentials.

```shell
python nats-pub.py --creds /creds/NGS.creds  -s tls://connect.ngs.global:4222 hello -d world
```

--------------------------------

### Restart nats-top with Query and Sort Options

Source: https://docs.nats.io/using-nats/nats-tools/nats_top/nats-top-tutorial

Restarts the nats-top tool, limiting the output to the top 1 connection and sorting the results by the number of subscriptions in ascending order. This is useful for identifying clients with the most activity.

```bash
nats-top -n 1 -sort subs
```

--------------------------------

### Get Keys List in NATS KV Store (C)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieves all keys in the NATS KV bucket and populates a kvKeysList structure. This function requires a pre-allocated kvKeysList and watch options.

```c
NATS_EXTERN natsStatus 	ki_Keys (kvKeysList *list, kvStore *kv, kvWatchOptions *opts)
 	Returns all keys in the bucket.
 
```

--------------------------------

### Example JWT Header and Payload (JSON)

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

Illustrates the structure of a NATS JWT, showing a typical header containing the algorithm and type, and a payload with standard JWT claims (jti, iat, iss, sub, type) and NATS-specific configuration.

```json
{
  "typ": "jwt",
  "alg": "ed25519"
}
{
  "jti": "ZP2X3T2R57SLXD2U5J3OLLYIVW2LFBMTXRPMMGISQ5OF7LANUQPQ",
  "iat": 1575468772,
  "iss": "OAFEEYZSYYVI4FXLRXJTMM32PQEI3RGOWZJT7Y3YFM4HB7ACPE4RTJPG",
  "name": "O",
  "sub": "OAFEEYZSYYVI4FXLRXJTMM32PQEI3RGOWZJT7Y3YFM4HB7ACPE4RTJPG",
  "type": "operator",
  "nats": {
    "operator_service_urls": [
      "nats://localhost:4222"
    ]
  }
}
```

--------------------------------

### C# JetStream Stream and Consumer Management

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

Provides examples for creating, updating, listing, and deleting streams and consumers using the NATS.Net client library.

```APIDOC
## C# JetStream API

### Stream Management

- **Description**: Manages streams, including creation, updating, and listing.
- **Method**: N/A (Code example provided)
- **Endpoint**: N/A (NATS.Net JetStream client methods)

#### `js.CreateStreamAsync(streamConfig)`
- **Description**: Creates a new stream.
  - **Request Body**: `streamConfig` (StreamConfig) - Configuration object for the stream (e.g., `name`, `subjects`, `maxBytes`).

#### `js.UpdateStreamAsync(streamConfig)`
- **Description**: Updates an existing stream.
  - **Request Body**: `streamConfig` (StreamConfig) - Updated configuration object for the stream.

#### `js.ListStreamsAsync()`
- **Description**: Asynchronously iterates through all available streams.

#### `js.DeleteStreamAsync(streamName)`
- **Description**: Deletes a stream by its name.

### Consumer Management

- **Description**: Manages consumers within streams, including creation, listing, and deletion.
- **Method**: N/A (Code example provided)
- **Endpoint**: N/A (NATS.Net JetStream client methods)

#### `js.CreateConsumerAsync(streamName, consumerConfig)`
- **Description**: Creates a new consumer for a specified stream.
  - **Parameters**:
    - `streamName` (string) - The name of the stream.
    - `consumerConfig` (ConsumerConfig) - Configuration object for the consumer (e.g., `durable_name`).

#### `js.ListConsumersAsync(streamName)`
- **Description**: Asynchronously iterates through all consumers in a given stream.

#### `js.DeleteConsumerAsync(streamName, consumerName)`
- **Description**: Deletes a consumer from a stream by name.
```

--------------------------------

### Python NATS Request-Reply Example

Source: https://docs.nats.io/using-nats/developer/sending/replyto

Demonstrates sending a request and handling a reply using the NATS client library in Python. It involves subscribing to a subject, publishing a request with an inbox for the reply, and processing the response.

```python
require 'nats/client'
require 'fiber'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  Fiber.new do
    f = Fiber.current

    nc.subscribe("time") do |msg, reply|
      f.resume msg
    end

    nc.publish("time", 'example', NATS.create_inbox)

    # Use the response
    msg = Fiber.yield
    puts "Reply: #{msg}"

  end.resume
end
```

--------------------------------

### Generate User NKey Shell Command

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/nkey_auth

This command generates a new NKey pair for user authentication using the `nk` tool. It outputs both the private seed (starting with 'S') and the public key (starting with 'U'). The seed is kept secret, while the public key is shared for server configuration.

```shell
nk -gen user -pubout
```

--------------------------------

### Object Operations

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

APIs for getting, putting, and updating objects within an object store.

```APIDOC
## GET /websites/nats_io/INatsObjStore/GetBytesAsync

### Description
Retrieves an object's value as a byte array by its key.

### Method
GET

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/bytes/{key}`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.
- **key** (string) - Required - The key of the object.

#### Query Parameters
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **byte[]** - The object's value as a byte array.

#### Response Example
```json
"SGVsbG8gV29ybGQh"
```

## GET /websites/nats_io/INatsObjStore/GetAsync

### Description
Retrieves an object by its key and writes it to a stream.

### Method
GET

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/{key}`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.
- **key** (string) - Required - The key of the object.

#### Query Parameters
- **stream** (Stream) - Required - The stream to write the object value to.
- **leaveOpen** (bool) - Optional - Whether to leave the stream open after writing.
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **ObjectMetadata** - The object's metadata.

#### Response Example
```json
{
  "name": "example_object",
  "size": 1234,
  "digest": {
    "algorithm": "sha256",
    "value": "..."
  }
}
```

## POST /websites/nats_io/INatsObjStore/PutAsync (byte array)

### Description
Puts an object by its key using a byte array as the value.

### Method
POST

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/{key}`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.
- **key** (string) - Required - The key of the object.

#### Request Body
- **value** (byte[]) - Required - The object's value as a byte array.

#### Query Parameters
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **ObjectMetadata** - The metadata of the created object.

#### Response Example
```json
{
  "name": "example_object",
  "size": 1234,
  "digest": {
    "algorithm": "sha256",
    "value": "..."
  }
}
```

## POST /websites/nats_io/INatsObjStore/PutAsync (stream)

### Description
Puts an object by its key using a stream for the value.

### Method
POST

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/{key}`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.
- **key** (string) - Required - The key of the object.

#### Request Body
- **stream** (Stream) - Required - The stream to read the value from.

#### Query Parameters
- **leaveOpen** (bool) - Optional - Whether to leave the stream open after writing.
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **ObjectMetadata** - The metadata of the created object.

#### Response Example
```json
{
  "name": "example_object",
  "size": 1234,
  "digest": {
    "algorithm": "sha256",
    "value": "..."
  }
}
```

## POST /websites/nats_io/INatsObjStore/PutAsync (metadata and stream)

### Description
Puts an object using provided metadata and a stream for the value.

### Method
POST

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.

#### Request Body
- **meta** (ObjectMetadata) - Required - The metadata for the object.
- **stream** (Stream) - Required - The stream to read the value from.

#### Query Parameters
- **leaveOpen** (bool) - Optional - Whether to leave the stream open after writing.
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **ObjectMetadata** - The metadata of the created object.

#### Response Example
```json
{
  "name": "example_object",
  "size": 1234,
  "digest": {
    "algorithm": "sha256",
    "value": "..."
  }
}
```

## PUT /websites/nats_io/INatsObjStore/UpdateMetaAsync

### Description
Updates the metadata for an existing object.

### Method
PUT

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/{key}/meta`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.
- **key** (string) - Required - The key of the object.

#### Request Body
- **meta** (ObjectMetadata) - Required - The new metadata for the object.

#### Query Parameters
- **cancellationToken** (CancellationToken) - Optional - Used to cancel the API call.

### Response
#### Success Response (200)
- **ObjectMetadata** - The updated metadata of the object.

#### Response Example
```json
{
  "name": "example_object",
  "size": 1234,
  "digest": {
    "algorithm": "sha256",
    "value": "..."
  }
}
```

## POST /websites/nats_io/INatsObjStore/AddLink

### Description
Adds a link to another object.

### Method
POST

### Endpoint
`/websites/nats_io/INatsObjStore/{bucket}/link`

### Parameters
#### Path Parameters
- **bucket** (string) - Required - The name of the bucket.

#### Request Body
- **link** (string) - Required - The name of the link.
- **target** (string) - Required - The name of the target object.

### Response
#### Success Response (200)
- **OK** - Indicates the link was successfully added.

#### Response Example
```json
{
  "status": "success"
}
```
```

--------------------------------

### Initialize NATS JetStream Connection (Python)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

A basic Python example to connect to a NATS server, create a JetStream context, and add a stream. This serves as a foundation for further JetStream operations. Requires the 'nats-py' library.

```python
import asyncio

import nats
from nats.errors import TimeoutError
    
async def main():
    nc = await nats.connect("localhost")

    # Create JetStream context.
    js = nc.jetstream()
        
    # Persist messages on 'foo's subject.
    await js.add_stream(name="sample-stream", subjects=["foo"])

    await nc.close()

if __name__ == '__main__':
    asyncio.run(main())
```

--------------------------------

### Connect to NATS Cluster in Ruby

Source: https://docs.nats.io/using-nats/developer/connecting/cluster

An example of connecting to a NATS cluster using the Ruby NATS client. It uses the `NATS.start` method with a block, passing an array of server URLs. The block contains a placeholder for operations and explicitly closes the connection.

```Ruby
require 'nats/client'

NATS.start(servers: ["nats://127.0.0.1:1222", "nats://127.0.0.1:1223", "nats://127.0.0.1:1224"]) do |nc|
   # Do something with the connection

   # Close the connection
   nc.close
end
```

--------------------------------

### NATS Configuration Variables Example

Source: https://docs.nats.io/nats-server/configuration

Shows how to define and reference variables within a NATS configuration file. Variables are block-scoped and referenced with a '$' prefix. Undefined variables can fall back to environment variables.

```plaintext
# Define a variable in the config
TOKEN: "secret"

# Reference the variable
authorization {
    token: $TOKEN
}
```

```plaintext
# Define a variable in the config
# But TOKEN is never used resulting in a config parsing error
TOKEN: "secret"

# Reference the variable
authorization {
    token: "another secret"
}
```

```plaintext
unknown field "TOKEN"
```

```shell
export TOKEN="hello"
nats-server -c /config/file
```

```plaintext
# TOKEN is defined in the environment
authorization {
    token: $TOKEN
}
```

--------------------------------

### Receive JSON Data in Python

Source: https://docs.nats.io/using-nats/developer/receiving/structure

This Python example shows how to connect to a NATS server using asyncio, subscribe to the 'updates' subject, and handle incoming messages. It decodes the message data from bytes to a JSON string and then parses it into a Python dictionary. The example also demonstrates publishing a JSON message and unsubscribing.

```python
import asyncio
import json
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

async def run(loop):
    nc = NATS()

    await nc.connect(servers=["nats://127.0.0.1:4222"], loop=loop)

    async def message_handler(msg):
        data = json.loads(msg.data.decode())
        print(data)

    sid = await nc.subscribe("updates", cb=message_handler)
    await nc.flush()

    await nc.auto_unsubscribe(sid, 2)
    await nc.publish("updates", json.dumps({"symbol": "GOOG", "price": 1200 }).encode())
    await asyncio.sleep(1, loop=loop)
    await nc.close()
```

--------------------------------

### Rust Project Configuration for WebAssembly

Source: https://docs.nats.io/using-nats/nex/getting-started/building-function

Configuration for a Rust project to build a WebAssembly module for Nex using the WASI interface. This `Cargo.toml` specifies the package details and sets up a binary target for `src/main.rs`.

```toml
[package]
name = "echofunction"
version = "0.1.0"
edition = "2021"

[[bin]]
name = "echofunction"
path = "src/main.rs"
```

--------------------------------

### Add a NATS JetStream Consumer (CLI)

Source: https://docs.nats.io/nats-concepts/jetstream/js_walkthrough

Creates a new NATS JetStream consumer with specified configurations. This command-line interface allows interactive setup of consumer name, delivery policy, acknowledgment, replay policy, and stream association. It supports pull consumers and defines how messages are delivered and acknowledged.

```shell
nats consumer add
```

--------------------------------

### NATS Server Log Output - Seed Node

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Example log output from a NATS seed server upon startup. It shows version information, listening ports for clients and cluster, and the server ID.

```log
[83329] 2020/02/12 16:04:52.369039 [INF] Starting nats-server version 2.1.4
[83329] 2020/02/12 16:04:52.369130 [DBG] Go build version go1.13.6
[83329] 2020/02/12 16:04:52.369133 [INF] Git commit [not set]
[83329] 2020/02/12 16:04:52.369360 [INF] Starting http monitor on 127.0.0.1:8222
[83329] 2020/02/12 16:04:52.369436 [INF] Listening for client connections on 127.0.0.1:4222
[83329] 2020/02/12 16:04:52.369441 [INF] Server id is NDSGCS74MG5ZUMBOVWOUJ5S3HIOW
[83329] 2020/02/12 16:04:52.369443 [INF] Server is ready
[83329] 2020/02/12 16:04:52.369534 [INF] Listening for route connections on 127.0.0.1:4248
```

--------------------------------

### Get Nex Node Information

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-functions

This command retrieves detailed information about a specific Nex node, including its ID, version, uptime, tags, memory usage, and the workloads it is running. This is useful for verifying deployment status and resource allocation.

```bash
$ nex node info NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
NEX Node Information

         Node: NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
         Xkey: XDKZMOZKVBXSY3YXPIXEFKGPML75PLD7APFHZ474EOCILZDQGPZSXJNZ
      Version: 0.0.1
       Uptime: 2m26s
         Tags: nex.arch=amd64, nex.cpucount=8, nex.os=linux, simple=true

Memory in kB:

           Free: 32,354,208
      Available: 55,985,740
          Total: 63,883,232

Workloads:

             Id: cmjud7n52omhlsa377cg
        Healthy: true
        Runtime: 2m26s
           Name: echofunctionjs
    Description: Workload published in devmode
```

--------------------------------

### Get Key-Value Entries

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieve the latest value or a specific revision of a value for a given key.

```APIDOC
## Get Key-Value Entries

### Description

Allows retrieving data associated with a key. You can fetch the most recent value or a specific historical revision.

### Methods

- `Get(key: str)`
- `GetRevision(key: str, revision: uint64)`

### Parameters

#### `Get(key: str)`

- **key** (str) - Required - The key whose value needs to be retrieved.

#### `GetRevision(key: str, revision: uint64)`

- **key** (str) - Required - The key whose value needs to be retrieved.
- **revision** (uint64) - Required - The specific revision number of the value to retrieve.

### Responses

#### `Get` and `GetRevision`

- **Success Response (200)**: Returns a `KeyValueEntry` object containing the value and metadata. Returns `null` if the key is not found.

### Examples

#### Go Example
```go
entry, err := kv.Get("my_key")
revision_entry, err := kv.GetRevision("my_key", 10)
```

#### Java Example
```java
KeyValueEntry entry = kv.get("my_key");
KeyValueEntry revisionEntry = kv.get("my_key", 10);
```

#### JavaScript Example
```javascript
const entry = await kv.get("my_key");
```

#### Python Example
```python
entry = await kv.get("my_key")
```

#### C# Example
```csharp
var entry = await kvStore.GetEntryAsync<string>("my_key");
var revisionEntry = await kvStore.GetEntryAsync<string>("my_key", 10);
```

#### C Example
```c
kvEntry *entry;
kvStore_Get(entry, kv, "my_key");
kvEntry *revision_entry;
kvStore_GetRevision(revision_entry, kv, "my_key", 10);
// Remember to free entry and revision_entry if not NULL
```
```

--------------------------------

### Connect to NATS Cluster in Python

Source: https://docs.nats.io/using-nats/developer/connecting/cluster

Provides an example of connecting to a NATS cluster using the Python NATS client. It uses an `async` function and an array of server URLs. The code includes placeholders for performing actions with the connection and for closing it.

```Python
nc = NATS()
await nc.connect(servers=[
   "nats://127.0.0.1:1222",
   "nats://127.0.0.1:1223",
   "nats://127.0.0.1:1224"
   ])

# Do something with the connection

await nc.close()
```

--------------------------------

### Get Store Status

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

Retrieves the run-time status of the object store.

```APIDOC
## GET /object/store/status

### Description
Fetches and returns the current run-time status of the object store, including information about its backing storage.

### Method
GET

### Endpoint
`/object/store/status`

### Parameters
No parameters required.

### Request Example
`GET /object/store/status`

### Response
#### Success Response (200)
- **status** (NatsObjStatus) - An object containing the status information of the object store.

#### Response Example
```json
{
  "capacity": 1073741824,
  "used": 52428800,
  "age": "2023-10-27T14:00:00Z"
}
```
```

--------------------------------

### Get Object Info

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

Retrieves metadata for a specific object by its key.

```APIDOC
## GET /object/store/info

### Description
Retrieves the metadata for a specific object using its unique key. Optionally includes deleted objects.

### Method
GET

### Endpoint
`/object/store/info`

### Parameters
#### Query Parameters
- **key** (string) - Required - The unique key of the object.
- **showDeleted** (boolean) - Optional - If true, includes deleted objects in the result. Defaults to false.

### Request Example
`GET /object/store/info?key=myObjectKey&showDeleted=true`

### Response
#### Success Response (200)
- **metadata** (ObjectMetadata) - Metadata of the requested object.

#### Response Example
```json
{
  "name": "myObjectKey",
  "size": 512,
  "digest": "jkl012pqr",
  "timestamp": "2023-10-27T12:00:00Z"
}
```

### Errors
- **NatsObjException**: Object was not found.
```

--------------------------------

### GET /varz - General Server Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves general information about the NATS server's state and configuration. Supports JSONP and CORS.

```APIDOC
## GET /varz - General Server Information

### Description
The `/varz` endpoint returns general information about the server state and configuration.

### Method
GET

### Endpoint
`/varz`

### Parameters

#### Path Parameters
N/A

#### Query Parameters
N/A

#### Request Body
N/A

### Request Example
`https://demo.nats.io:8222/varz`

### Response
#### Success Response (200)
- **server_id** (string) - Unique identifier for the NATS server.
- **version** (string) - The version of the NATS server.
- **proto** (integer) - The protocol version supported by the server.
- **go** (string) - The Go version used to build the server.
- **host** (string) - The host address the server is listening on.
- **port** (integer) - The port the NATS server is listening on.
- **max_connections** (integer) - The maximum number of allowed connections.
- **ping_interval** (integer) - The interval for sending pings to clients (in nanoseconds).
- **ping_max** (integer) - The maximum number of missed pings before disconnecting a client.
- **http_host** (string) - The host address for the monitoring HTTP server.
- **http_port** (integer) - The port for the monitoring HTTP server.
- **https_port** (integer) - The port for the HTTPS monitoring server (0 if not enabled).
- **auth_timeout** (number) - The authentication timeout in seconds.
- **max_control_line** (integer) - The maximum size of the control line.
- **max_payload** (integer) - The maximum payload size in bytes.
- **max_pending** (integer) - The maximum pending bytes for a connection.
- **cluster** (object) - Cluster configuration details.
- **gateway** (object) - Gateway configuration details.
- **leaf** (object) - Leaf node configuration details.
- **tls_timeout** (number) - The TLS handshake timeout in seconds.
- **write_deadline** (integer) - The write deadline for connections (in nanoseconds).
- **start** (string) - The timestamp when the server started.
- **now** (string) - The current timestamp.
- **uptime** (string) - The duration the server has been running.
- **mem** (integer) - Memory usage in bytes.
- **cores** (integer) - Number of CPU cores available.
- **gomaxprocs** (integer) - The number of OS threads that can execute user-level Go code simultaneously.
- **cpu** (number) - Current CPU utilization percentage.
- **connections** (integer) - The number of active connections.
- **total_connections** (integer) - The total number of connections established since server start.
- **routes** (integer) - The number of active routes.
- **remotes** (integer) - The number of active remote connections.
- **leafnodes** (integer) - The number of active leaf nodes.
- **in_msgs** (integer) - Total incoming messages.
- **out_msgs** (integer) - Total outgoing messages.
- **in_bytes** (integer) - Total incoming bytes.
- **out_bytes** (integer) - Total outgoing bytes.
- **slow_consumers** (integer) - The number of slow consumers.
- **subscriptions** (integer) - The number of active subscriptions.
- **http_req_stats** (object) - Statistics for HTTP requests made to various endpoints.
- **config_load_time** (string) - Timestamp when the configuration was last loaded.
- **slow_consumer_stats** (object) - Statistics related to slow consumers.

#### Response Example
```json
{
  "server_id": "NACDVKFBUW4C4XA24OOT6L4MDP56MW76J5RJDFXG7HLABSB46DCMWCOW",
  "version": "2.0.0",
  "proto": 1,
  "go": "go1.12",
  "host": "0.0.0.0",
  "port": 4222,
  "max_connections": 65536,
  "ping_interval": 120000000000,
  "ping_max": 2,
  "http_host": "0.0.0.0",
  "http_port": 8222,
  "https_port": 0,
  "auth_timeout": 1,
  "max_control_line": 4096,
  "max_payload": 1048576,
  "max_pending": 67108864,
  "cluster": {},
  "gateway": {},
  "leaf": {},
  "tls_timeout": 0.5,
  "write_deadline": 2000000000,
  "start": "2019-06-24T14:24:43.928582-07:00",
  "now": "2019-06-24T14:24:46.894852-07:00",
  "uptime": "2s",
  "mem": 9617408,
  "cores": 4,
  "gomaxprocs": 4,
  "cpu": 0,
  "connections": 0,
  "total_connections": 0,
  "routes": 0,
  "remotes": 0,
  "leafnodes": 0,
  "in_msgs": 0,
  "out_msgs": 0,
  "in_bytes": 0,
  "out_bytes": 0,
  "slow_consumers": 2,
  "subscriptions": 0,
  "http_req_stats": {
    "/": 0,
    "/connz": 0,
    "/gatewayz": 0,
    "/routez": 0,
    "/subsz": 0,
    "/varz": 1
  },
  "config_load_time": "2019-06-24T14:24:43.928582-07:00",
  "slow_consumer_stats": {
    "clients": 1,
    "routes": 1,
    "gateways": 0,
    "leafs": 0
  }
}
```
```

--------------------------------

### Get Echo Service Information using NATS CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-services

This command displays detailed information about the EchoService, including its ID, version, endpoints, and statistics such as request counts and processing time. It confirms the service's operational status and performance metrics.

```bash
$ nats micro info EchoService
Service Information

          Service: EchoService (NsMaTbN7u5ZPUNN47bSEI6)
      Description: 
          Version: 1.0.0

Endpoints:

               Name: default
            Subject: svc.echo
        Queue Group: q

Statistics for 1 Endpoint(s):

  default Endpoint Statistics:

           Requests: 1 in group q
    Processing Time: 15µs (average 15µs)
            Started: 2024-01-16 19:40:09 (7m46s ago)
             Errors: 0
```

--------------------------------

### Connect to NATS with TLS in Java

Source: https://docs.nats.io/using-nats/developer/connecting/tls

This Java example demonstrates connecting to a NATS server with TLS. It involves creating an SSLContext with client certificates and CA certificates, often requiring conversion to JKS format using openssl and keytool.

```java
// This examples requires certificates to be in the java keystore format (.jks).
// To do so openssl is used to generate a pkcs12 file (.p12) from client-cert.pem and client-key.pem.
// The resulting file is then imported int a java keystore named keystore.jks using keytool which is part of java jdk.
// keytool is also used to import the CA certificate rootCA.pem into truststore.jks.  
// 
// openssl pkcs12 -export -out keystore.p12 -inkey client-key.pem -in client-cert.pem -password pass:password
// keytool -importkeystore -srcstoretype PKCS12 -srckeystore keystore.p12 -srcstorepass password -destkeystore keystore.jks -deststorepass password
//
// keytool -importcert -trustcacerts -file rootCA.pem -storepass password -noprompt -keystore truststore.jks
class SSLUtils {
    public static String KEYSTORE_PATH = "keystore.jks";
    public static String TRUSTSTORE_PATH = "truststore.jks";
    public static String STORE_PASSWORD = "password";
    public static String KEY_PASSWORD = "password";
    public static String ALGORITHM = "SunX509";

    public static KeyStore loadKeystore(String path) throws Exception {
        KeyStore store = KeyStore.getInstance("JKS");
        BufferedInputStream in = new BufferedInputStream(new FileInputStream(path));
        try {
            store.load(in, STORE_PASSWORD.toCharArray());
        } finally {
            in.close();
        }

        return store;
    }

    public static KeyManager[] createTestKeyManagers() throws Exception {
        KeyStore store = loadKeystore(KEYSTORE_PATH);
        KeyManagerFactory factory = KeyManagerFactory.getInstance(ALGORITHM);
        factory.init(store, KEY_PASSWORD.toCharArray());
        return factory.getKeyManagers();
    }

    public static TrustManager[] createTestTrustManagers() throws Exception {
        KeyStore store = loadKeystore(TRUSTSTORE_PATH);
        TrustManagerFactory factory = TrustManagerFactory.getInstance(ALGORITHM);
        factory.init(store);
        return factory.getTrustManagers();
    }

    public static SSLContext createSSLContext() throws Exception {
        SSLContext ctx = SSLContext.getInstance(Options.DEFAULT_SSL_PROTOCOL);
        ctx.init(createTestKeyManagers(), createTestTrustManagers(), new SecureRandom());
        return ctx;
    }
}

public class ConnectTLS {
    public static void main(String[] args) {

        try {
            SSLContext ctx = SSLUtils.createSSLContext();
            Options options = new Options.Builder()
                .server("nats://localhost:4222")
                .sslContext(ctx) // Set the SSL context
                .build();
            Connection nc = Nats.connect(options);

            // Do something with the connection

            nc.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

```

--------------------------------

### Full Wildcard '>' Mapping Example

Source: https://docs.nats.io/nats-concepts/subject_mapping

Illustrates prefixing a subject using the full wildcard '>'. This mapping redirects messages from any subject under a specific pattern to a new prefixed subject.

```shell
nats server mapping ">"  "baz.>" bar.a.b
> baz.bar.b.a
```

--------------------------------

### Create and Manage JetStream KeyValue Stores

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This section covers the creation, retrieval, and deletion of JetStream KeyValue stores. It includes methods for initializing configurations, creating new stores, accessing existing ones, and cleaning up deleted stores. Dependencies include the NATS client libraries for each respective language.

```python
async def key_value(self, bucket: str) -> KeyValue:

async def create_key_value(
    self,
    config: Optional[api.KeyValueConfig] = None,
    **params,
) -> KeyValue:
    """
    create_key_value takes an api.KeyValueConfig and creates a KV in JetStream.
    """
    
async def delete_key_value(self, bucket: str) -> bool:
    """
    delete_key_value deletes a JetStream KeyValue store by destroying
    the associated stream.
    """
```

```csharp
// dotnet add package NATS.Net

// Create a new Key Value Store or get an existing one
ValueTask<INatsKVStore> CreateStoreAsync(string bucket, CancellationToken cancellationToken = default);

// Get a list of bucket names
IAsyncEnumerable<string> GetBucketNamesAsync(CancellationToken cancellationToken = default);

// Gets the status for all buckets
IAsyncEnumerable<NatsKVStatus> GetStatusesAsync(CancellationToken cancellationToken = default);

// Delete a Key Value Store
ValueTask<bool> DeleteStoreAsync(string bucket, CancellationToken cancellationToken = default);


```

```c
NATS_EXTERN natsStatus 	kvConfig_Init (kvConfig *cfg)
 	Initializes a KeyValue configuration structure.
 
NATS_EXTERN natsStatus 	js_CreateKeyValue (kvStore **new_kv, jsCtx *js, kvConfig *cfg)
 	Creates a KeyValue store with a given configuration.
 
NATS_EXTERN natsStatus 	js_KeyValue (kvStore **new_kv, jsCtx *js, const char *bucket)
 	Looks-up and binds to an existing KeyValue store.
 
NATS_EXTERN natsStatus 	js_DeleteKeyValue (jsCtx *js, const char *bucket)
 	Deletes a KeyValue store.
 
NATS_EXTERN void 	kvStore_Destroy (kvStore *kv)
 	Destroys a KeyValue store object.
```

--------------------------------

### GET /connz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves information about active client connections to the NATS server.

```APIDOC
## GET /connz

### Description
Retrieves information about active client connections to the NATS server.

### Method
GET

### Endpoint
/connz

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **name** (string) - The name of the server region.
- **host** (string) - The hostname or IP address of the server.
- **port** (integer) - The port the server is listening on.
- **outbound_gateways** (object) - Information about outbound gateway connections.
  - **[region_name]** (object) - Details for each outbound gateway connection.
    - **configured** (boolean) - Indicates if the gateway is configured.
    - **connection** (object) - Details about the connection.
      - **cid** (integer) - Connection ID.
      - **ip** (string) - IP address of the connection.
      - **port** (integer) - Port of the connection.
      - **start** (string) - Timestamp when the connection started.
      - **last_activity** (string) - Timestamp of the last activity.
      - **uptime** (string) - Duration the connection has been up.
      - **idle** (string) - Duration the connection has been idle.
      - **pending_bytes** (integer) - Number of pending bytes.
      - **in_msgs** (integer) - Number of incoming messages.
      - **out_msgs** (integer) - Number of outgoing messages.
      - **in_bytes** (integer) - Number of incoming bytes.
      - **out_bytes** (integer) - Number of outgoing bytes.
      - **subscriptions** (integer) - Number of subscriptions.
      - **name** (string) - Name of the connection.
- **inbound_gateways** (object) - Information about inbound gateway connections.
  - **[region_name]** (array) - An array of objects, each representing an inbound gateway connection.
    - **configured** (boolean) - Indicates if the gateway is configured.
    - **connection** (object) - Details about the connection (same structure as outbound connection).

#### Response Example
```json
{
  "server_id": "NANVBOU62MDUWTXWRQ5KH3PSMYNCHCEUHQV3TW3YH7WZLS7FMJE6END6",
  "now": "2019-07-24T18:02:55.597398-06:00",
  "name": "region1",
  "host": "2601:283:4601:1350:1895:efda:2010:95a1",
  "port": 4501,
  "outbound_gateways": {
    "region2": {
      "configured": true,
      "connection": {
        "cid": 7,
        "ip": "127.0.0.1",
        "port": 5500,
        "start": "2019-07-24T18:02:48.765621-06:00",
        "last_activity": "2019-07-24T18:02:48.765621-06:00",
        "uptime": "6s",
        "idle": "6s",
        "pending_bytes": 0,
        "in_msgs": 0,
        "out_msgs": 0,
        "in_bytes": 0,
        "out_bytes": 0,
        "subscriptions": 0,
        "name": "NCXBIYWT7MV7OAQTCR4QTKBN3X3HDFGSFWTURTCQ22ZZB6NKKJPO7MN4"
      }
    },
    "region3": {
      "configured": true,
      "connection": {
        "cid": 5,
        "ip": "::1",
        "port": 6500,
        "start": "2019-07-24T18:02:48.764685-06:00",
        "last_activity": "2019-07-24T18:02:48.764685-06:00",
        "uptime": "6s",
        "idle": "6s",
        "pending_bytes": 0,
        "in_msgs": 0,
        "out_msgs": 0,
        "in_bytes": 0,
        "out_bytes": 0,
        "subscriptions": 0,
        "name": "NCVS7Q65WX3FGIL2YQRLI77CE6MQRWO2Y453HYVLNMBMTVLOKMPW7R6K"
      }
    }
  },
  "inbound_gateways": {
    "region2": [
      {
        "configured": false,
        "connection": {
          "cid": 9,
          "ip": "::1",
          "port": 52029,
          "start": "2019-07-24T18:02:48.76677-06:00",
          "last_activity": "2019-07-24T18:02:48.767096-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCXBIYWT7MV7OAQTCR4QTKBN3X3HDFGSFWTURTCQ22ZZB6NKKJPO7MN4"
        }
      }
    ],
    "region3": [
      {
        "configured": false,
        "connection": {
          "cid": 4,
          "ip": "::1",
          "port": 52025,
          "start": "2019-07-24T18:02:48.764577-06:00",
          "last_activity": "2019-07-24T18:02:48.764994-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCVS7Q65WX3FGIL2YQRLI77CE6MQRWO2Y453HYVLNMBMTVLOKMPW7R6K"
        }
      },
      {
        "configured": false,
        "connection": {
          "cid": 8,
          "ip": "127.0.0.1",
          "port": 52026,
          "start": "2019-07-24T18:02:48.766173-06:00",
          "last_activity": "2019-07-24T18:02:48.766999-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCKCYK5LE3VVGOJQ66F65KA27UFPCLBPX4N4YOPOXO3KHGMW24USPCKN"
        }
      }
    ]
  }
}
```
```

--------------------------------

### Confirm NATS Connection with CONNECT Command

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol-demo

Confirms the established connection to the NATS server by sending a CONNECT message. An empty payload is sufficient for this example, and a '+OK' response indicates success.

```text
CONNECT {}
```

--------------------------------

### Configure NATS Ping/Pong Settings (C#)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Demonstrates setting a 20-second ping interval and a maximum of 5 outstanding pings for the NATS client in C#. The example utilizes the `NatsClient` constructor with `NatsOpts` to specify these parameters and uses `await using` for resource management.

```C#
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.Core;

await using var client = new NatsClient(new NatsOpts
{
    Url = "nats://demo.nats.io:4222",
    
    // Set Ping Interval to 20 seconds and Max Pings Outstanding to 5
    PingInterval = TimeSpan.FromSeconds(20),
    MaxPingOut = 5,
});

```

--------------------------------

### Java JetStream Push Subscribe Durable Queue Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Demonstrates JetStream push subscribing with a durable consumer and queueing in Java. It sets up a stream, connects to NATS, and prepares for message consumption. Requires a running NATS server and client libraries.

```java
package io.nats.examples.jetstream;

import io.nats.client.*;
import io.nats.client.api.PublishAck;
import io.nats.examples.ExampleArgs;
import io.nats.examples.ExampleUtils;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

import static io.nats.examples.jetstream.NatsJsUtils.createStreamExitWhenExists;

/**
 * This example will demonstrate JetStream push subscribing using a durable consumer and a queue
 */
public class NatsJsPushSubQueueDurable {
    static final String usageString =
        "\nUsage: java -cp <classpath> NatsJsPushSubQueueDurable [-s server] [-strm stream] [-sub subject] [-q queue] [-dur durable] [-mcnt msgCount] [-scnt subCount]" 
            + "\n\nDefault Values:"
            + "\n   [-strm stream]   qdur-stream" 
            + "\n   [-sub subject]   qdur-subject" 
            + "\n   [-q queue]       qdur-queue" 
            + "\n   [-dur durable]   qdur-durable" 
            + "\n   [-mcnt msgCount] 100" 
            + "\n   [-scnt subCount] 5" 
            + "\n\nUse tls:// or opentls:// to require tls, via the Default SSLContext\n" 
            + "\nSet the environment variable NATS_NKEY to use challenge response authentication by setting a file containing your private key.\n" 
            + "\nSet the environment variable NATS_CREDS to use JWT/NKey authentication by setting a file containing your user creds.\n" 
            + "\nUse the URL in the -s server parameter for user/pass/token authentication.\n";

    public static void main(String[] args) {
        ExampleArgs exArgs = ExampleArgs.builder("Push Subscribe, Durable Consumer, Queue", args, usageString)
                .defaultStream("qdur-stream")
                .defaultSubject("qdur-subject")
                .defaultQueue("qdur-queue")
                .defaultDurable("qdur-durable")
                .defaultMsgCount(100)
                .defaultSubCount(5)
                .build();

        try (Connection nc = Nats.connect(ExampleUtils.createExampleOptions(exArgs.server, true))) {

            // Create a JetStreamManagement context.
            JetStreamManagement jsm = nc.jetStreamManagement();

            // Use the utility to create a stream stored in memory.
            createStreamExitWhenExists(jsm, exArgs.stream, exArgs.subject);

            // Create our JetStream context
            JetStream js = nc.jetStream();

```

--------------------------------

### Download an Object with a Specific Output Path

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This demonstrates downloading an object from a bucket and specifying a custom output path using the `--output` flag. This allows for precise control over where the downloaded file is saved on the local filesystem.

```shell
nats object get myobjbucket --output /temp/Movies/NATS-logo.mov /Movies/NATS-logo.mov
```

--------------------------------

### Start Leaf Node Server

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This command initiates the NATS server configured as a leaf node, using the specified configuration file (e.g., /tmp/leaf.conf). This allows local clients to connect and interact with the remote NATS system.

```bash
nats-server -c /tmp/leaf.conf
```

--------------------------------

### Download NATS Server Release Build (Shell)

Source: https://docs.nats.io/running-a-nats-service/introduction/installation

This snippet downloads a specific release version of the NATS server binary archive from GitHub using curl. It's followed by commands to extract the archive and copy the server executable to a system directory.

```shell
curl -L https://github.com/nats-io/nats-server/releases/download/vX.Y.Z/nats-server-vX.Y.Z-linux-amd64.tar.gz -o nats-server.tar.gz
```

```shell
tar -xvzf nats-server.tar.gz
```

```shell
sudo cp nats-server-vX.Y.Z-linux-amd64/nats-server /usr/bin
```

--------------------------------

### Display NATS CLI Help Information

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli

Shows the general help message for the NATS CLI, listing available commands and global options. This is useful for understanding the CLI's capabilities and structure. It can also be used to get help for specific commands by appending the command name.

```shell
nats --help
```

--------------------------------

### Request-Reply with NATS.io in Python

Source: https://docs.nats.io/using-nats/developer/sending/replyto

An asynchronous example in Python using the 'nats-python' library. It sets up a subscriber for the 'time' subject that will reply with the current time. A request is then sent, and the response is processed asynchronously via a callback.

```python
nc = NATS()

future = asyncio.Future()

async def sub(msg):
  nonlocal future
  future.set_result(msg)

await nc.connect(servers=["nats://demo.nats.io:4222"])
await nc.subscribe("time", cb=sub)

unique_reply_to = nc.new_inbox()
await nc.publish("time", b'', unique_reply_to)
```

--------------------------------

### NATS Dispatcher with CountDownLatch (Java)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

A Java example demonstrating NATS client usage. It connects to a NATS server, sets up a dispatcher to handle incoming messages on 'time.>', and uses a `CountDownLatch` to wait for a specific number of messages before closing the connection.

```java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

// Use a latch to wait for 4 messages to arrive
CountDownLatch latch = new CountDownLatch(4);

// Create a dispatcher and inline message handler
Dispatcher d = nc.createDispatcher((msg) -> {
    String subject = msg.getSubject();
    String str = new String(msg.getData(), StandardCharsets.UTF_8);
    System.out.println(subject + ": " + str);
    latch.countDown();
});

// Subscribe
d.subscribe("time.>");

// Wait for messages to come in
latch.await();

// Close the connection
nc.close();
```

--------------------------------

### NATS Messaging with Fibers (Ruby)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This Ruby example uses NATS client and Fibers to publish messages and then yield to receive them. It demonstrates asynchronous message handling within a single thread.

```ruby
require 'nats/client'
require 'fiber'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  Fiber.new do
    f = Fiber.current

    nc.subscribe("time.*.east") do |msg, reply|
      f.resume Time.now
    end

    nc.publish("time.A.east", "A")
    nc.publish("time.B.east", "B")

    # Use the response
    msg_A = Fiber.yield
    puts "Msg A: #{msg_A}"

    msg_B = Fiber.yield
    puts "Msg B: #{msg_B}"

  end.resume
end
```

--------------------------------

### Publish Time Data to NATS Subjects (Ruby)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

Provides a Ruby example for publishing messages to NATS subjects. It demonstrates connecting to NATS, publishing data, and gracefully draining connections.

```ruby
NATS.start do |nc|
   nc.publish("time.us.east", '...')
   nc.publish("time.us.east.atlanta", '...')

   nc.publish("time.eu.east", '...')
   nc.publish("time.eu.east.warsaw", '...')

   nc.drain
end
```

--------------------------------

### GET /connz - Connection Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves detailed information on current and recently closed connections. Supports paging.

```APIDOC
## GET /connz - Connection Information

### Description
The `/connz` endpoint reports more detailed information on current and recently closed connections. It uses a paging mechanism which defaults to 1024 connections.

### Method
GET

### Endpoint
`/connz`

### Parameters

#### Path Parameters
N/A

#### Query Parameters
- **limit** (integer) - Optional. The number of connections to return per page. Defaults to 1024.
- **offset** (integer) - Optional. The offset for pagination.

#### Request Body
N/A

### Request Example
`https://demo.nats.io:8222/connz?limit=100`

### Response
#### Success Response (200)
- **total** (integer) - The total number of connections.
- **bounded** (boolean) - Indicates if the results are bounded by the limit.
- **connections** (array) - An array of connection objects.
  - **id** (integer) - Connection ID.
  - **cid** (integer) - Client connection ID.
  - **name** (string) - Name of the client.
  - **ip** (string) - Client IP address.
  - **port** (integer) - Client port.
  - **start** (string) - Connection start time.
  - **last** (string) - Last activity time.
  - **idle** (string) - Idle duration.
  - **lang** (string) - Client language.
  - **version** (string) - Client version.
  - **subs** (integer) - Number of subscriptions.
  - **pending** (integer) - Number of pending messages.
  - **in_bytes** (integer) - Total bytes received.
  - **out_bytes** (integer) - Total bytes sent.
  - **in_msgs** (integer) - Total messages received.
  - **out_msgs** (integer) - Total messages sent.
  - **account** (string) - The account the client is connected to.

#### Response Example
```json
{
  "total": 1,
  "bounded": true,
  "connections": [
    {
      "id": 1,
      "cid": 1,
      "name": "",
      "ip": "127.0.0.1",
      "port": 54321,
      "start": "2019-06-24T14:24:46.896036-07:00",
      "last": "2019-06-24T14:24:46.896036-07:00",
      "idle": "0s",
      "lang": "go",
      "version": "2.0.0",
      "subs": 0,
      "pending": 0,
      "in_bytes": 0,
      "out_bytes": 0,
      "in_msgs": 0,
      "out_msgs": 0,
      "account": ""
    }
  ]
}
```
```

--------------------------------

### NATS Gateway Configuration Example

Source: https://docs.nats.io/running-a-nats-service/configuration/gateways/gateway

This snippet demonstrates the configuration of a NATS gateway, including its name, listen address, and a list of other gateways it can connect to. The `urls` property allows specifying multiple endpoints for a single gateway, enabling resilient connections.

```hcl
gateway {
    name: "DC-A"
    listen: "localhost:7222"

    gateways: [
        {name: "DC-A", urls: ["nats://localhost:7222", "nats://localhost:7223", "nats://localhost:7224"]},
        {name: "DC-B", urls: ["nats://localhost:7332", "nats://localhost:7333", "nats://localhost:7334"]},
        {name: "DC-C", urls: ["nats://localhost:7442", "nats://localhost:7333", "nats://localhost:7335"]}
    ]
}
```

--------------------------------

### Run NATS Subscriber Example

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-docker-python

Executes the nats-sub.py script to create a persistent NATS subscription to the 'hello' subject. It uses provided NGS credentials and connects to the global NATS server via TLS.

```shell
python nats-sub.py --creds /creds/NGS.creds  -s tls://connect.ngs.global:4222 hello &
```

--------------------------------

### NATS Request-Reply in Python

Source: https://docs.nats.io/using-nats/developer/sending/request_reply

Provides an asynchronous Python example for the request-reply pattern. It sets up a subscriber for 'time' requests and then sends a request with a one-second timeout, handling potential timeouts.

```python
nc = NATS()

async def sub(msg):
  await nc.publish(msg.reply, b'response')

await nc.connect(servers=["nats://demo.nats.io:4222"])
await nc.subscribe("time", cb=sub)

# Send the request
try:
  msg = await nc.request("time", b'', timeout=1)
  # Use the response
  print("Reply:", msg)
except asyncio.TimeoutError:
  print("Timed out waiting for response")
```

--------------------------------

### GET /subsz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves detailed information about the current subscriptions and the routing data structure. This endpoint is not normally used.

```APIDOC
## GET /subsz

### Description
Retrieves detailed information about the current subscriptions and the routing data structure. This endpoint is not normally used.

### Method
GET

### Endpoint
/subsz

### Response
#### Success Response (200)
(Response structure not detailed in the provided text)

#### Error Response (400)
(Error response details not detailed in the provided text)
```

--------------------------------

### Get List of All Accounts

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Fetches a list of all active accounts known to the NATS server. The response includes server details and a list of account names.

```json
{
  "server_id": "NAB2EEQ3DLS2BHU4K2YMXMPIOOOAOFOAQAC5NQRIEUI4BHZKFBI4ZU4A",
  "now": "2021-02-08T17:31:29.551146-05:00",
  "system_account": "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5",
  "accounts": ["AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5", "$G"]
}
```

--------------------------------

### Server Configuration: Account-Specific Mappings

Source: https://docs.nats.io/nats-concepts/subject_mapping

An example of NATS server configuration showing how to define subject mappings within a specific account. This allows for granular control over message routing for different services or tenants.

```yaml
server_name: "hub"
cluster: { name: "hub" }

accounts {
    accountA: {
        mappings: {
            orders.flush  orders.central.flush 
        	orders.* orders.central.{{wildcard(1)}}
        }
    }
}
```

--------------------------------

### GET /connz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves information about NATS server connections. Supports filtering by state, client ID, and pagination, as well as detailed subscription information.

```APIDOC

            "queue": "",
            "max": 0,
            "active": true,
            "in_msgs": 10,
            "in_bytes": 1024
          }
        ]
      }
    ]
  }
}
```

**Sort Options**

| Option      | Sort by                                              |
| ----------- | ---------------------------------------------------- |
| cid         | Connection ID                                        |
| start       | Connection start time, same as CID                   |
| subs        | Number of subscriptions                              |
| pending     | Amount of data in bytes waiting to be sent to client |
| msgs_to    | Number of messages sent                              |
| msgs_from  | Number of messages received                          |
| bytes_to   | Number of bytes sent                                 |
| bytes_from | Number of bytes received                             |
| last        | Last activity                                        |
| idle        | Amount of inactivity                                 |
| uptime      | Lifetime of the connection                           |
| stop        | Stop time for a closed connection                    |
| reason      | Reason for a closed connection                       |
| rtt         | Round trip time                                      |
```

--------------------------------

### Connect to NATS with TLS in JavaScript (Node.js)

Source: https://docs.nats.io/using-nats/developer/connecting/tls

This JavaScript example for the Node.js runtime shows how to connect to a NATS server using TLS. It specifies the paths for CA certificate, client key, and client certificate files within the connection options.

```javascript
// tls options available depend on the javascript
// runtime, please verify the readme for the
// client you are using for specific details
// this example showing the node library
const nc = await connect({
  port: ns.port,
  debug: true,
  tls: {
    caFile: caCertPath,
    keyFile: clientKeyPath,
    certFile: clientCertPath,
  },
});
```

--------------------------------

### GET /jsz - JetStream Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves detailed JetStream information, including account, stream, and consumer data. Supports pagination and filtering options.

```APIDOC
## GET /jsz

### Description
Retrieves detailed JetStream information, including account, stream, and consumer data. Supports pagination and filtering options. In clustered environments, it's recommended to query the leader for the most accurate data.

### Method
GET

### Endpoint
/jsz

### Parameters
#### Query Parameters
- **acc** (string) - Optional - Include metrics for the specified account. Default is unset.
- **accounts** (boolean) - Optional - Include account-specific JetStream information. Default is false.
- **streams** (boolean) - Optional - Include streams. When set, implies `accounts=true`. Default is false.
- **consumers** (boolean) - Optional - Include consumer information. When set, implies `streams=true`. Default is false.
- **config** (boolean) - Optional - When stream or consumer are requested, include their respective configuration. Default is false.
- **leader-only** (boolean) - Optional - Only the leader responds. Default is false.
- **offset** (number) - Optional - Pagination offset. Default is 0.
- **limit** (number) - Optional - Number of results to return. Default is 1024.
- **raft** (boolean) - Optional - Include information details about the Raft group. Default is false.

### Request Example
`GET /jsz`
`GET /jsz?accounts=true&limit=16&offset=128`
`GET /jsz?consumers=true`

### Response
#### Success Response (200)
- The response structure varies based on the query parameters. It can include information about accounts, streams, consumers, configurations, and Raft groups.

#### Response Example
(Response structure depends on query parameters, e.g., `accounts=true`, `consumers=true`)
```json
{
  "server_id": "...",
  "now": "...",
  "accounts": [
    {
      "acc": "account_name",
      "js_info": {
        "max_memory": ..., 
        "max_storage": ..., 
        "storage": ..., 
        "reserve_memory": ..., 
        "reserve_storage": ...
      },
      "streams": [
        {
          "name": "stream_name",
          "subjects": ["subject.*"],
          "count": 10,
          "highest": 50,
          "lost": 0,
          "consumers": 5,
          "is_file": true,
          "is_memory": false,
          "subjects_wildcard": true,
          "storage": 1024000,
          "bytes": 512000,
          "msgs": 1000,
          "lost_msgs": 0,
          "hdms": 0,
          "allow_delete": true,
          "allow_purge": true,
          "max_msgs": 0,
          "max_bytes": 0,
          "max_age": 0,
          "replicas": [
            {
              "name": "server_id",
              "url": "..."
            }
          ]
        }
      ]
    }
  ]
}
```
```

--------------------------------

### Get Keys in NATS KV Store (C#)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieves all keys from the NATS KV bucket asynchronously. Supports optional watch options and cancellation tokens. Requires the NATS.Net package. Overload available for filtered key retrieval.

```csharp
// Get all the keys in the bucket
IAsyncEnumerable<string> GetKeysAsync(NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);

// Get a filtered set of keys in the bucket
IAsyncEnumerable<string> GetKeysAsync(IEnumerable<string> filters, NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);

//
```

--------------------------------

### Object Operations

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

APIs for performing operations on objects within an object store, including putting, getting, updating metadata, deleting, and adding links.

```APIDOC
## Object Operations API

### Description

APIs for performing operations on objects within an object store, including putting, getting, updating metadata, deleting, and adding links.

### Methods

- `put(ObjectMeta meta, InputStream inputStream)`: Place the contents of the input stream into a new object.
- `put(String objectName, InputStream inputStream)`: Place the contents of the input stream into a new object with a specified name.
- `put(String objectName, byte[] input)`: Place the bytes into a new object.
- `put(File file)`: Place the contents of the file into a new object using the file name as the object name.
- `get(String objectName, OutputStream outputStream)`: Get an object by name from the store, reading it into the output stream.
- `getInfo(String objectName)`: Get the info for an object by name.
- `getInfo(String objectName, boolean includingDeleted)`: Get the info for an object, optionally including deleted objects.
- `updateMeta(String objectName, ObjectMeta meta)`: Update the metadata of an object (name, description, headers).
- `delete(String objectName)`: Delete the object by name.
- `addLink(String objectName, ObjectInfo toInfo)`: Add a link to another object.
- `addBucketLink(String objectName, ObjectStore toStore)`: Add a link to another object store (bucket).

### Parameters

#### Put Operations

- **meta** (ObjectMeta) - Metadata for the object.
- **inputStream** (InputStream) - The input stream containing the object data.
- **objectName** (String) - The name of the object.
- **input** (byte[]) - The byte array containing the object data.
- **file** (File) - The file containing the object data.

#### Get Operations

- **objectName** (String) - The name of the object to retrieve.
- **outputStream** (OutputStream) - The output stream to write the object data to.
- **includingDeleted** (boolean) - Whether to include deleted objects in the info retrieval.

#### Update/Delete/Link Operations

- **objectName** (String) - The name of the object to operate on.
- **meta** (ObjectMeta) - The updated metadata for the object.
- **toInfo** (ObjectInfo) - The ObjectInfo of the object to link to.
- **toStore** (ObjectStore) - The ObjectStore to link to.

### Request Example (Put)

```json
{
  "objectName": "my-object.txt",
  "meta": {
    "description": "This is my object",
    "headers": {
      "Content-Type": "text/plain"
    }
  },
  "inputStream": "...binary content..."
}
```

### Response Example (Get Info)

```json
{
  "name": "my-object.txt",
  "digest": {
    "algorithm": "SHA-256",
    "hash": "a1b2c3d4..."
  },
  "size": 1024,
  "created": "2023-10-27T10:00:00Z",
  "metadata": {
    "description": "This is my object",
    "headers": {
      "Content-Type": "text/plain"
    }
  }
}
```
```

--------------------------------

### GET /leafz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves detailed information about the leaf node connections to the NATS server. It can optionally include internal subscriptions.

```APIDOC
## GET /leafz

### Description
Retrieves detailed information about the leaf node connections to the NATS server. It can optionally include internal subscriptions.

### Method
GET

### Endpoint
/leafz

#### Query Parameters
- **subs** (boolean) - Optional - Include internal subscriptions. Default is false.

### Request Example
```json
{
  "example": "/leafz?subs=true"
}
```

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **leafnodes** (integer) - The total number of leaf nodes connected.
- **leafs** (array) - An array of objects, each representing a leaf node connection.
  - **account** (string) - The account associated with the leaf node.
  - **ip** (string) - The IP address of the leaf node.
  - **port** (integer) - The port of the leaf node.
  - **rtt** (string) - The round-trip time to the leaf node.
  - **in_msgs** (integer) - The number of messages received from the leaf node.
  - **out_msgs** (integer) - The number of messages sent to the leaf node.
  - **in_bytes** (integer) - The number of bytes received from the leaf node.
  - **out_bytes** (integer) - The number of bytes sent to the leaf node.
  - **subscriptions** (integer) - The number of subscriptions the leaf node has.
  - **subscriptions_list** (array) - An array of strings, listing the subscriptions.

#### Response Example
```json
{
  "server_id": "NC2FJCRMPBE5RI5OSRN7TKUCWQONCKNXHKJXCJIDVSAZ6727M7MQFVT3",
  "now": "2019-08-27T09:07:05.841132-06:00",
  "leafnodes": 1,
  "leafs": [
    {
      "account": "$G",
      "ip": "127.0.0.1",
      "port": 6223,
      "rtt": "200µs",
      "in_msgs": 0,
      "out_msgs": 10000,
      "in_bytes": 0,
      "out_bytes": 1280000,
      "subscriptions": 1,
      "subscriptions_list": ["foo"]
    }
  ]
}
```
```

--------------------------------

### Subject Hierarchy Example - General Namespace

Source: https://docs.nats.io/nats-concepts/subjects

Demonstrates using the initial tokens of a subject to establish a general namespace, followed by more specific identifiers. This pattern helps in organizing subjects and routing messages effectively.

```shell
factory1.tools.group42.unit17
```

--------------------------------

### Get Consumer Info

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Retrieves detailed information about a specific consumer by its name.

```APIDOC
## POST $JS.API.CONSUMER.INFO.<stream>.<consumer>

### Description
Retrieves information about a specific consumer by name.

### Method
POST

### Endpoint
`$JS.API.CONSUMER.INFO.<stream>.<consumer>`

### Parameters
#### Path Parameters
- **stream** (string) - Required - The name of the stream.
- **consumer** (string) - Required - The name of the consumer.

#### Request Body
- **data** (empty payload)

### Response
#### Success Response (200)
- **data** (api.JSApiConsumerInfoResponse) - Information about the specified consumer.
```

--------------------------------

### Java NKey Authentication

Source: https://docs.nats.io/using-nats/developer/connecting/nkey

Shows how to implement NKey authentication in Java for NATS clients. This example defines a custom AuthHandler to provide the public key ID and sign nonces using an NKey object. The NKey should be securely loaded.

```Java
NKey theNKey = NKey.createUser(null); // really should load from somewhere
Options options = new Options.Builder()
    .server("nats://localhost:4222")
    .authHandler(new AuthHandler(){
        public char[] getID() {
            try {
                return theNKey.getPublicKey();
            } catch (GeneralSecurityException|IOException|NullPointerException ex) {
                return null;
            }
        }

        public byte[] sign(byte[] nonce) {
            try {
                return theNKey.sign(nonce);
            } catch (GeneralSecurityException|IOException|NullPointerException ex) {
                return null;
            }
        }

        public char[] getJWT() {
            return null;
        }
    })
    .build();
Connection nc = Nats.connect(options);

// Do something with the connection

nc.close();
```

--------------------------------

### Get Information About an Object Store Bucket

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This command retrieves and displays detailed information about an object store bucket, including its name, replication status, TTL, size, and backing JetStream stream. This is useful for monitoring and managing bucket properties.

```shell
nats object info myobjbucket
```

--------------------------------

### Connect to NATS Server using URL (Ruby)

Source: https://docs.nats.io/using-nats/developer/connecting/specific_server

Starts a NATS client connection to the specified server URL. The connection is managed within a block, and `NATS.start` handles the lifecycle. The connection object `nc` is available within the block for NATS operations.

```ruby
require 'nats/client'

NATS.start(servers: ["nats://demo.nats.io:4222"]) do |nc|
   # Do something with the connection

   # Close the connection
   nc.close
end
```

--------------------------------

### Listing NATS Users for an Account with NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This NSC command lists all users associated with a specific account, identified by its name. It outputs the name and public key of each user, essential for configuring user access and permissions.

```bash
nsc list users -a A
```

--------------------------------

### Deploy JavaScript Function with Nex

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-functions

This command deploys a JavaScript function to Nex, specifying a trigger subject for activation. It outputs workload information upon successful acceptance. Ensure the path to your JavaScript file is correct.

```bash
$ nex devrun /home/kevin/echofunction.js --trigger_subject=js.echo
Reusing existing issuer account key: /home/kevin/.nex/issuer.nk
Reusing existing publisher xkey: /home/kevin/.nex/publisher.xk
🚀 Workload 'echofunctionjs' accepted. You can now refer to this workload with ID: cmjud7n52omhlsa377cg on node NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
```

--------------------------------

### Get Information about a NATS JetStream Stream

Source: https://docs.nats.io/nats-concepts/jetstream/js_walkthrough

Retrieves and displays detailed information about an existing JetStream stream, including its subjects, storage type, retention policy, and current state.

```shell
nats stream info my_stream
```

--------------------------------

### Publish and Subscribe with NATS in C#

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This C# example shows how to publish messages to NATS subjects and subscribe to them using the NATS.Net client library. It demonstrates setting time zone specific data for messages and handling subscriptions.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using System;
using System.Threading.Tasks;

public class NatsExample
{
    public static async Task Main(string[] args)
    {
        await using var client = new NatsClient();
        await client.ConnectAsync("nats://demo.nats.io:4222");
        Console.WriteLine("Connected to NATS.");

        // Publish messages with time zone information
        var easternZone = TimeZoneInfo.FindSystemTimeZoneById("America/New_York");
        var nowEastern = TimeZoneInfo.ConvertTimeFromUtc(DateTime.UtcNow, easternZone);
        string formattedEastern = nowEastern.ToString("yyyy-MM-ddTHH:mm:ss zzz");

        await client.PublishAsync("time.us.east", System.Text.Encoding.UTF8.GetBytes(formattedEastern));
        await client.PublishAsync("time.us.east.atlanta", System.Text.Encoding.UTF8.GetBytes(formattedEastern));

        var warsawZone = TimeZoneInfo.FindSystemTimeZoneById("Europe/Warsaw");
        var nowWarsaw = TimeZoneInfo.ConvertTimeFromUtc(DateTime.UtcNow, warsawZone);
        string formattedWarsaw = nowWarsaw.ToString("yyyy-MM-ddTHH:mm:ss zzz");

        await client.PublishAsync("time.eu.east", System.Text.Encoding.UTF8.GetBytes(formattedWarsaw));
        await client.PublishAsync("time.eu.east.warsaw", System.Text.Encoding.UTF8.GetBytes(formattedWarsaw));

        Console.WriteLine("Published messages.");

        // Subscribe to receive messages
        var subscription = await client.SubscribeAsync<string>("time.>");
        var count = 0;
        await foreach (var msg in subscription)
        {
            Console.WriteLine($"Received {{++count}}: {msg.Subject}: {msg.Data}");
            if (count == 4)
            {
                break;
            }
        }

        Console.WriteLine("Done");
        await client.DrainAsync();
    }
}
```

--------------------------------

### GET /gatewayz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Reports information about gateways used to create a NATS supercluster. This endpoint provides details on connected gateways, including their names and account information.

```APIDOC
## GET /gatewayz

### Description
Reports information about gateways used to create a NATS supercluster. This endpoint provides details on connected gateways, including their names and account information.

### Method
GET

### Endpoint
/gatewayz

### Parameters
#### Query Parameters
- **accs** (boolean, integer) - Optional - Include account information. Can be true, 1, false, or 0. Default is false.
- **gw_name** (string) - Optional - Return only remote gateways with this name.
- **acc_name** (string) - Optional - Limit the list of accounts to this account name.

### Request Example
```json
{
  "example": "GET /gatewayz?accs=true&gw_name=mygateway"
}
```

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **num_gateways** (integer) - The total number of active gateways.
- **gateways** (array) - An array of gateway objects, each containing:
    - **rid** (integer) - Gateway ID.
    - **remote_id** (string) - The ID of the remote gateway.
    - **did_solicit** (boolean) - Indicates if the gateway solicited the connection.
    - **ip** (string) - Remote gateway IP address.
    - **port** (integer) - Remote gateway port.
    - **pending_size** (integer) - Size of pending data.
    - **in_msgs** (integer) - Number of incoming messages.
    - **out_msgs** (integer) - Number of outgoing messages.
    - **in_bytes** (integer) - Number of incoming bytes.
    - **out_bytes** (integer) - Number of outgoing bytes.
    - **subscriptions** (integer) - Number of subscriptions.
    - **accounts** (array) - Optional array of account objects associated with the gateway.

#### Response Example
```json
{
  "server_id": "NACDVKFBUW4C4XA24OOT6L4MDP56MW76J5RJDFXG7HLABSB46DCMWCOW",
  "now": "2019-06-24T14:29:16.046656-07:00",
  "num_gateways": 1,
  "gateways": [
    {
      "rid": 1,
      "remote_id": "de475c0041418afc799bccf0fdd61b47",
      "did_solicit": true,
      "ip": "127.0.0.1",
      "port": 61791,
      "pending_size": 0,
      "in_msgs": 0,
      "out_msgs": 0,
      "in_bytes": 0,
      "out_bytes": 0,
      "subscriptions": 0
    }
  ]
}
```
```

--------------------------------

### Publish and Subscribe with NATS in Ruby

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This Ruby example uses the 'nats/client' gem to publish and subscribe to NATS messages. It demonstrates using Fibers for asynchronous operations and includes publishing messages with timestamps.

```ruby
require 'nats/client'
require 'fiber'

NATS.start(servers: ["nats://127.0.0.1:4222"]) do |nc|
  Fiber.new do
    f = Fiber.current

    nc.subscribe("time.>") do |msg|
      f.resume Time.now.to_f
    end

    # Publish messages
    nc.publish("time.A.east", "A")
    nc.publish("time.B.east", "B")
    nc.publish("time.C.west", "C")
    nc.publish("time.D.west", "D")

    # Use the response
    4.times do
      msg = Fiber.yield
      puts "Msg: #{msg}"
    end
  end.resume
end
```

--------------------------------

### Request Data from NATS Service

Source: https://docs.nats.io/using-nats/nex/getting-started/building-service

This command sends a request to the 'svc.echo' subject on the NATS server with the message 'this is a test'. The 'EchoService' should receive this request and respond with the same message.

```bash
$ nats req svc.echo 'this is a test'
```

--------------------------------

### NATS Server Clustering Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration

Enables NATS servers to form a cluster for load balancing and redundancy. This configuration block is not active by default and requires explicit setup.

```configuration
cluster {
  # Cluster configuration options go here
}
```

--------------------------------

### NATS Server INFO Message with Nonce

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

An example of a NATS server INFO message sent to a client upon connection, including a `nonce` field. This nonce is used by the client to generate a signature for authentication.

```text
INFO {
    "auth_required": true,
    "client_id": 3,
    "client_ip": "::1",
    "go": "go1.14.1",
    "headers": true,
    "host": "0.0.0.0",
    "max_payload": 1048576,
    "nonce": "-QPTE1Jsk8kI3rE",
    "port": 4222,
    "proto": 1,
    "server_id": "NBSHIXACRHUODC4FY2Z3OYXSZSRUBRH6VWIKQNGVPKOTA7H4YTXWJRTO",
    "server_name": "NBSHIXACRHUODC4FY2Z3OYXSZSRUBRH6VWIKQNGVPKOTA7H4YTXWJRTO",
    "version": "2.2.0-beta.26"
}
```

--------------------------------

### Leaf Node Connections - Outgoing (Non Operator Mode)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Configuration example for leaf node outgoing connections when the cluster is not in operator mode. This includes setting up credentials for user and system accounts.

```APIDOC
## Leaf Node Connections - Outgoing (Non Operator Mode)

### Description
Configuring outgoing leaf node connections when the cluster is not in Operator mode. The system account may differ from the user account.

### Configuration Example

```
leafnodes {
    remotes = [
        {
          url: "nats://localhost:4222"
          credentials: "./your-account.creds"
        },
        {
          url: "nats://localhost:4222"
          account: "$SYS"
          credentials: "./system-account.creds"
        },
    ]
}
```

### Notes
* Leaf nodes do not multiplex between accounts; each account needs to be explicitly listed.
* The system account (`$SYS`) must be explicitly connected if needed, even if ends use the same system account.
* System account credentials can be restricted to specific subjects.
* Default Account `$G` and default system account `$SYS` are referenced in the example.
```

--------------------------------

### Manage KV Buckets (Java)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This Java code provides methods for managing NATS Key/Value (KV) buckets. It includes functionality to create a new KV store with configuration, retrieve a list of all available bucket names, get detailed information about a specific bucket, and delete an existing bucket. Error handling for IO and JetStream API exceptions is included.

```java
import io.nats.client.JetStreamApiException;
import io.nats.client.KeyValue;
import io.nats.client.KeyValueConfiguration;
import io.nats.client.KeyValueStatus;
import java.io.IOException;
import java.util.List;

public class KeyValueManager {

    /**
     * Create a key value store.
     * @param config the key value configuration
     * @return bucket info
     * @throws IOException covers various communication issues with the NATS
     *         server such as timeout or interruption
     * @throws JetStreamApiException the request had an error related to the data
     * @throws IllegalArgumentException the server is not JetStream enabled
     */
    public static KeyValueStatus create(KeyValueConfiguration config) throws IOException, JetStreamApiException {
        // ... implementation details ...
        return null;
    }

    /**
    * Get the list of bucket names.
    * @return list of bucket names
    * @throws IOException covers various communication issues with the NATS
    *         server such as timeout or interruption
    * @throws JetStreamApiException the request had an error related to the data
    * @throws InterruptedException if the thread is interrupted
    */
    public static List<String> getBucketNames() throws IOException, JetStreamApiException, InterruptedException {
        // ... implementation details ...
        return null;
    }

    /**
    * Gets the info for an existing bucket.
    * @param bucketName the bucket name to use
    * @throws IOException covers various communication issues with the NATS
    *         server such as timeout or interruption
    * @throws JetStreamApiException the request had an error related to the data
    * @return the bucket status object
    */
    public static KeyValueStatus getBucketInfo(String bucketName) throws IOException, JetStreamApiException {
        // ... implementation details ...
        return null;
    }

    /**
    * Deletes an existing bucket. Will throw a JetStreamApiException if the delete fails.
    * @param bucketName the stream name to use.
    * @throws IOException covers various communication issues with the NATS
    *         server such as timeout or interruption
    * @throws JetStreamApiException the request had an error related to the data
    */
    public static void delete(String bucketName) throws IOException, JetStreamApiException {
        // ... implementation details ...
    }
}
```

--------------------------------

### Get Key History in Go, Java, JavaScript, C#, C

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieves all historical values for a specified key in the Key/Value store. Supports multiple programming languages, each with specific method signatures and error handling.

```Go
// History will return all historical values for the key.
History(key string, opts ...WatchOpt) ([]KeyValueEntry, error)
```

```Java
/**
 * Get the history (list of KeyValueEntry) for a key
 * @param key the key
 * @return List of KvEntry
 * @throws IOException covers various communication issues with the NATS
 *         server such as timeout or interruption
 * @throws JetStreamApiException the request had an error related to the data
 * @throws InterruptedException if the thread is interrupted
 */
List<KeyValueEntry> history(String key) throws IOException, JetStreamApiException, InterruptedException;
```

```JavaScript
async history(
    opts: { key?: string; headers_only?: boolean } = {},
  ): Promise<QueuedIterator<KvEntry>>
```

```C#
// dotnet add package NATS.Net

// Get the history of an entry by key
IAsyncEnumerable<NatsKVEntry<T>> HistoryAsync<T>(string key, INatsDeserialize<T>? serializer = default, NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);
```

```C
NATS_EXTERN natsStatus 	kvStore_History (kvEntryList *list, kvStore *kv, const char *key, kvWatchOptions *opts)
 	Returns all historical entries for the key.
 
NATS_EXTERN void 	kvEntryList_Destroy (kvEntryList *list)
 	Destroys this list of KeyValue store entries.
```

--------------------------------

### Create Pull-Based Consumer (CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/consumers

Demonstrates adding a pull-based consumer named 'NEW' to the 'ORDERS' stream. It specifies a sample rate, delivery target, start policy, subject filter, and maximum deliveries. The output shows the consumer's configuration and current state.

```shell
nats con ls ORDERS
nats con add --sample 100
nats con add ORDERS DISPATCH --filter ORDERS.processed --ack explicit --pull --deliver all --sample 100 --max-deliver 20
```

--------------------------------

### Go JWT Package for Auth Callout Data Structures

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_callout

Provides an example of using the nats-io/jwt Go package to work with the key data structures for auth callout: authorization request claims, authorization response claims, and user claims. This package is essential for implementing the auth service.

```go
import (
	"github.com/nats-io/jwt/v2"
)

// Example usage within a Go program (conceptual):
// var authRequestClaims jwt.AuthCalloutRequestClaims
// var authResponseClaims jwt.AuthCalloutResponseClaims
// var userClaims jwt.UserClaims
```

--------------------------------

### Create Push-Based Consumer (CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/consumers

Illustrates creating a push-based consumer named 'MONITOR' for the 'ORDERS' stream. This consumer has no acknowledgment policy, targets 'monitor.ORDERS', starts from the last message, and has an instant replay policy. The output displays the consumer's configuration.

```shell
nats con add
nats con add ORDERS MONITOR --ack none --target monitor.ORDERS --deliver last --replay instant --filter ''
```

--------------------------------

### GET /routez

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Reports information on active routes for a NATS cluster. This endpoint provides details about connected routes, including their IDs, remote addresses, and message statistics.

```APIDOC
## GET /routez

### Description
Reports information on active routes for a NATS cluster. This endpoint provides details about connected routes, including their IDs, remote addresses, and message statistics.

### Method
GET

### Endpoint
/routez

### Parameters
#### Query Parameters
- **subs** (boolean, integer, string) - Optional - Include subscriptions. Can be true, 1, false, 0, or 'detail'. Default is false. When set to 'detail', a list with more detailed subscription information will be returned.

### Request Example
```json
{
  "example": "GET /routez?subs=1"
}
```

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **num_routes** (integer) - The total number of active routes.
- **routes** (array) - An array of route objects, each containing:
    - **rid** (integer) - Route ID.
    - **remote_id** (string) - The ID of the remote server.
    - **did_solicit** (boolean) - Indicates if the route solicited the connection.
    - **ip** (string) - Remote server IP address.
    - **port** (integer) - Remote server port.
    - **pending_size** (integer) - Size of pending data.
    - **in_msgs** (integer) - Number of incoming messages.
    - **out_msgs** (integer) - Number of outgoing messages.
    - **in_bytes** (integer) - Number of incoming bytes.
    - **out_bytes** (integer) - Number of outgoing bytes.
    - **subscriptions** (integer) - Number of subscriptions.

#### Response Example
```json
{
  "server_id": "NACDVKFBUW4C4XA24OOT6L4MDP56MW76J5RJDFXG7HLABSB46DCMWCOW",
  "now": "2019-06-24T14:29:16.046656-07:00",
  "num_routes": 1,
  "routes": [
    {
      "rid": 1,
      "remote_id": "de475c0041418afc799bccf0fdd61b47",
      "did_solicit": true,
      "ip": "127.0.0.1",
      "port": 61791,
      "pending_size": 0,
      "in_msgs": 0,
      "out_msgs": 0,
      "in_bytes": 0,
      "out_bytes": 0,
      "subscriptions": 0
    }
  ]
}
```
```

--------------------------------

### GET /connz

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves information about active client connections to the NATS server. This endpoint provides details such as connection IDs, types, IP addresses, and message statistics.

```APIDOC
## GET /connz

### Description
Retrieves information about active client connections to the NATS server. This endpoint provides details such as connection IDs, types, IP addresses, and message statistics.

### Method
GET

### Endpoint
/connz

### Parameters
#### Query Parameters
- **subs** (boolean, integer, string) - Optional - Include subscriptions. Can be true, 1, false, 0, or 'detail'. Default is false. When set to 'detail', a list with more detailed subscription information will be returned.

### Request Example
```json
{
  "example": "GET /connz?subs=detail"
}
```

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **num_connections** (integer) - The total number of active connections.
- **total** (integer) - The total number of connections (can include inactive ones depending on server config).
- **offset** (integer) - Offset for pagination (if applicable).
- **limit** (integer) - Limit for pagination (if applicable).
- **connections** (array) - An array of connection objects, each containing:
    - **cid** (integer) - Connection ID.
    - **kind** (string) - Type of connection (e.g., 'Client').
    - **type** (string) - Specific connection type (e.g., 'nats', 'mqtt').
    - **ip** (string) - Client IP address.
    - **port** (integer) - Client port.
    - **start** (string) - Connection start time.
    - **last_activity** (string) - Last activity time.
    - **rtt** (string) - Round Trip Time.
    - **uptime** (string) - Connection uptime.
    - **idle** (string) - Time since last activity.
    - **pending_bytes** (integer) - Number of pending bytes.
    - **in_msgs** (integer) - Number of incoming messages.
    - **out_msgs** (integer) - Number of outgoing messages.
    - **in_bytes** (integer) - Number of incoming bytes.
    - **out_bytes** (integer) - Number of outgoing bytes.
    - **subscriptions** (integer) - Number of subscriptions.
    - **name** (string) - Optional client name.
    - **lang** (string) - Optional client language.
    - **version** (string) - Optional client version.
    - **mqtt_client** (string) - Optional MQTT client identifier.

#### Response Example
```json
{
  "server_id": "NACDVKFBUW4C4XA24OOT6L4MDP56MW76J5RJDFXG7HLABSB46DCMWCOW",
  "now": "2019-06-24T14:28:16.520365-07:00",
  "num_connections": 2,
  "total": 2,
  "offset": 0,
  "limit": 1024,
  "connections": [
    {
      "cid": 5,
      "kind": "Client",
      "type": "nats",
      "ip": "127.0.0.1",
      "port": 62714,
      "start": "2021-09-09T23:16:43.040862Z",
      "last_activity": "2021-09-09T23:16:43.042364Z",
      "rtt": "95µs",
      "uptime": "5s",
      "idle": "5s",
      "pending_bytes": 0,
      "in_msgs": 0,
      "out_msgs": 0,
      "in_bytes": 0,
      "out_bytes": 0,
      "subscriptions": 1,
      "name": "NATS Benchmark",
      "lang": "go",
      "version": "1.12.1"
    }
  ]
}
```
```

--------------------------------

### Get Keys in NATS KV Store (Go)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieves a list of all keys currently having a value associated with them in the NATS KV store. This function accepts optional watch options.

```go
// Keys will return all keys.
Keys(opts ...WatchOpt) ([]string, error)
```

--------------------------------

### Get Keys in NATS KV Store (Java)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Retrieves a list of all keys present in a NATS KV bucket. This method is asynchronous and may throw IOExceptions, JetStreamApiExceptions, or InterruptedException.

```java
/**
 * Get a list of the keys in a bucket.
 * @return List of keys
 * @throws IOException covers various communication issues with the NATS
 *         server such as timeout or interruption
 * @throws JetStreamApiException the request had an error related to the data
 * @throws InterruptedException if the thread is interrupted
 */
List<String> keys() throws IOException, JetStreamApiException, InterruptedException;
```

--------------------------------

### Go Subscriber Service Connection with NKEY Authentication

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This Go code snippet demonstrates how to connect a subscriber service to a NATS server using NKEY authentication. It includes options for loading the NKEY from a seed file and establishing a connection to the NATS server at a specified address.

```go
package main

import (
    "log"

    "github.com/nats-io/nats.go"
)

func main() {
    opts := make([]nats.Option, 0)

    // Extract public nkey from seed
    //
    // Public:  UAPOK2P7EN3UFBL7SBJPQK3M3JMLALYRYKX5XWSVMVYK63ZMBHTOHVJR
    // Private: SUANVBWRHHFMGHNIT6UJHPN2TGVBVIILE7VPVNEQ7DGCJ26ZD2V3KAHT4M
    // 
    nkey, err := nats.NkeyOptionFromSeed("path/to/seed.nkey")
    if err != nil {
        log.Fatal(err)
    }
    opts = append(opts, nkey)
    nc, err := nats.Connect("127.0.0.1:4222", opts...)
    if err != nil {

```

--------------------------------

### Add Users with Tags for Templated Permissions (NATS CLI)

Source: https://docs.nats.io/using-nats/nats-tools/nsc/signing_keys

These commands show how to add users to NATS using the CLI, assigning them specific tags that will be used by the templated signing key permissions. The `--tag` flag is crucial for providing the necessary context for template expansion.

```bash
nsc add user pam -K team-service --tag team:support
nsc add user joe -K team-service --tag team:leads
```

--------------------------------

### Publish to NATS JetStream Stream (C - Usage)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/publish

This C code snippet is a usage example for publishing messages to NATS JetStream. It defines command-line arguments for stream name, text content, message count, and publish mode (synchronous/asynchronous), but the actual publishing logic is not shown in this excerpt.

```c
#include "examples.h"

static const char *usage = "\
-stream        stream name (default is 'foo')\
-txt           text to send (default is 'hello')\
-count         number of messages to send\
-sync          publish synchronously (default is async)\
";

static void

```

--------------------------------

### NATS Server Log Output - Connecting Server

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Example log output from a NATS server connecting to a cluster. It shows the attempt to connect to a route, the successful creation of a route connection, and the registration of the remote route.

```log
[83330] 2020/02/12 16:05:09.661047 [INF] Starting nats-server version 2.1.4
[83330] 2020/02/12 16:05:09.661123 [DBG] Go build version go1.13.6
[83330] 2020/02/12 16:05:09.661125 [INF] Git commit [not set]
[83330] 2020/02/12 16:05:09.661341 [INF] Listening for client connections on 0.0.0.0:5222
[83330] 2020/02/12 16:05:09.661347 [INF] Server id is NAABC2CKRVPZBIECMLZZA6L3PK
[83330] 2020/02/12 16:05:09.661349 [INF] Server is ready
[83330] 2020/02/12 16:05:09.662429 [INF] Listening for route connections on localhost:5248
[83330] 2020/02/12 16:05:09.662676 [DBG] Trying to connect to route on localhost:4248
[83330] 2020/02/12 16:05:09.663308 [DBG] 127.0.0.1:4248 - rid:1 - Route connect msg sent
[83330] 2020/02/12 16:05:09.663370 [INF] 127.0.0.1:4248 - rid:1 - Route connection created
[83330] 2020/02/12 16:05:09.663537 [DBG] 127.0.0.1:4248 - rid:1 - Registering remote route "NDSGCS74MG5ZUMBOVWOUJ5S3HIOW"
[83330] 2020/02/12 16:05:09.663549 [DBG] 127.0.0.1:4248 - rid:1 - Sent local subscriptions to route
```

--------------------------------

### Get Keys in NATS KV Store (JavaScript)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Asynchronously retrieves all keys from the NATS KV store, optionally filtering with a provided key pattern. Returns a Promise that resolves to a QueuedIterator of strings.

```javascript
async keys(k = ">"): Promise<QueuedIterator<string>>
```

--------------------------------

### GET /accstatz - Account Statistics

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves account statistics, including connection counts and data transfer metrics. Supports an option to include accounts without active connections.

```APIDOC
## GET /accstatz

### Description
Retrieves account statistics, including connection counts and data transfer metrics. Supports an option to include accounts without active connections.

### Method
GET

### Endpoint
/accstatz

### Parameters
#### Query Parameters
- **unused** (boolean) - Optional - If true, include accounts that do not have any current connections. Default is false.

### Request Example
`GET /accstatz`
`GET /accstatz?unused=1`

### Response
#### Success Response (200)
- **server_id** (string) - The ID of the NATS server.
- **now** (string) - The current timestamp.
- **account_statz** (array) - An array of account statistics objects.
  - **acc** (string) - The account name.
  - **conns** (integer) - The number of active connections.
  - **leafnodes** (integer) - The number of leaf nodes.
  - **total_conns** (integer) - The total number of connections (including leaf nodes).
  - **num_subscriptions** (integer) - The number of subscriptions.
  - **sent** (object) - Information about sent messages and bytes.
    - **msgs** (integer) - Number of messages sent.
    - **bytes** (integer) - Number of bytes sent.
  - **received** (object) - Information about received messages and bytes.
    - **msgs** (integer) - Number of messages received.
    - **bytes** (integer) - Number of bytes received.
  - **slow_consumers** (integer) - The number of slow consumers.

#### Response Example
```json
{
  "server_id": "NDJ5M4F5WAIBUA26NJ3QMH532AQPN7QNTJP3Y4SBHSHL4Y7QUAKNJEAF",
  "now": "2022-10-19T17:16:20.881296749Z",
  "account_statz": [
    {
      "acc": "default",
      "conns": 31,
      "leafnodes": 2,
      "total_conns": 33,
      "num_subscriptions": 45,
      "sent": {
        "msgs": 1876970,
        "bytes": 246705616
      },
      "received": {
        "msgs": 1347454,
        "bytes": 219438308
      },
      "slow_consumers": 29
    },
    {
      "acc": "$G",
      "conns": 1,
      "leafnodes": 0,
      "total_conns": 1,
      "num_subscriptions": 3,
      "sent": {
        "msgs": 0,
        "bytes": 0
      },
      "received": {
        "msgs": 107,
        "bytes": 1094
      },
      "slow_consumers": 0
    }
  ]
}
```
```

--------------------------------

### Get Specific Account Details

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves detailed information for a specified NATS account, including its name, JWT, permissions, and subscription statistics. Requires the account name as a query parameter.

```json
{
  "server_id": "NAB2EEQ3DLS2BHU4K2YMXMPIOOOAOFOAQAC5NQRIEUI4BHZKFBI4ZU4A",
  "now": "2021-02-08T17:37:55.80856-05:00",
  "system_account": "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5",
  "account_detail": {
    "account_name": "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5",
    "update_time": "2021-02-08T17:31:22.390334-05:00",
    "is_system": true,
    "expired": false,
    "complete": true,
    "jetstream_enabled": false,
    "leafnode_connections": 0,
    "client_connections": 0,
    "subscriptions": 42,
    "exports": [
      {
        "subject": "$SYS.DEBUG.SUBSCRIBERS",
        "type": "service",
        "response_type": "Singleton"
      }
    ],
    "jwt": "eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJVVlU2VEpXRU8zS0hYWTZVMkgzM0RCVklET1A3U05DTkJPMlM0M1dPNUM2T1RTTDNVSUxBIiwiaWF0IjoxNjAzNDczNzg4LCJpc3MiOiJPQlU1T05GSjMyNFVEUFJCSVZSR0Y3Q05FT0hHTFBTN0VZUEJUVlFaS1NCSElJWklCNkhENjZKRiIsIm5hbWUiOiJTWVMiLCJzdWIiOiJBQUFYQVVWU0dLN1RDUkhGSVJBUzRTWVhWSjc2RVdETU5YWk02QVJGR1hQN0JBU05ER0xLVTdBNSIsInR5cGUiOiJhY2NvdW50IiwibmF0cyI6eyJsaW1pdHMiOnsic3VicyI6LTEsImNvbm4iOi0xLCJsZWFmIjotMSwiaW1wb3J0cyI6LTEsImV4cG9ydHMiOi0xLCJkYXRhIjotMSwicGF5bG9hZCI6LTEsIndpbGRjYXJkcyI6dHJ1ZX19fQ.CeGo16i5oD0b1uBJ8UdGmLH-l9dL8yNqXHggkAt2T5c88fM7k4G08wLguMAnlvzrdlYvdZvOx_5tHLuDZmGgCg",
    "issuer_key": "OBU5O5FJ324UDPRBIVRGF7CNEOHGLPS7EYPBTVQZKSBHIIZIB6HD66JF",
    "name_tag": "SYS",
    "decoded_jwt": {
      "jti": "UVU6TJWEO3KHXY6U2H33DBVIDOP7SNCNBO2S43WO5C6OTSL3UILA",
      "iat": 1603473788,
      "iss": "OBU5O5FJ324UDPRBIVRGF7CNEOHGLPS7EYPBTVQZKSBHIIZIB6HD66JF",
      "name": "SYS",
      "sub": "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5",
      "nats": {
        "limits": {
          "subs": -1,
          "data": -1,
          "payload": -1,
          "imports": -1,
          "exports": -1,
          "wildcards": true,
          "conn": -1,
          "leaf": -1
        },
        "default_permissions": {
          "pub": {},
          "sub": {}
        },
        "type": "account",
        "version": 1
      }
    },
    "sublist_stats": {
      "num_subscriptions": 42,
      "num_cache": 6,
      "num_inserts": 42,
      "num_removes": 0,
      "num_matches": 6,
      "cache_hit_rate": 0,
      "max_fanout": 1,
      "avg_fanout": 0.8333333333333334
    }
  }
}
```

--------------------------------

### Generate CA and Server Certificates with mkcert

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/tls

Generates a Certificate Authority (CA) and server certificates for localhost and IP address ::1 using mkcert. These certificates are then used to start a NATS server with TLS enabled. This is useful for local testing environments.

```bash
mkcert -install
mkcert -cert-file server-cert.pem -key-file server-key.pem localhost ::1
nats-server --tls --tlscert=server-cert.pem --tlskey=server-key.pem -ms 8222
```

--------------------------------

### Run a Single NATS Server with Docker

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This command runs a single NATS server instance using the Docker image. It creates a Docker network named 'nats', then starts the server, mapping ports 4222 and 8222, and setting the HTTP management port to 8222.

```bash
docker network create nats
docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

--------------------------------

### Generating NATS Configuration with Mem Resolver

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This command uses NSC to generate a NATS server configuration with an in-memory resolver. This is useful for creating temporary or dynamic configurations, especially when setting up trusted operator environments.

```bash
nsc generate config --mem-resolver --sys-account SYS
```

--------------------------------

### Run 1:N Throughput Test with NATS Bench

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli/natsbench

This command initiates a NATS benchmarking test with one publisher and five subscribers. It specifies sending 1,000,000 messages, each 16 bytes in size. The output includes overall and per-publisher/subscriber throughput statistics.

```bash
nats bench foo --pub 1 --sub 5 --size 16 --msgs 1000000
```

--------------------------------

### Get Subscription Routing Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Retrieves information about active subscriptions on the NATS server. Supports filtering by subscription count, offset, and limit. Can also test for the existence of a subscription.

```json
{
  "num_subscriptions": 2,
  "num_cache": 0,
  "num_inserts": 2,
  "num_removes": 0,
  "num_matches": 0,
  "cache_hit_rate": 0,
  "max_fanout": 0,
  "avg_fanout": 0
}
```

--------------------------------

### Asynchronous NATS Subscribe with Callback (Python)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This Python example uses the `asyncio` library to connect to NATS asynchronously. It subscribes to 'time.>' subjects and uses an `asyncio.Queue` to handle incoming messages via a callback function.

```python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

# Use queue to wait for 4 messages to arrive
queue = asyncio.Queue()
async def cb(msg):
  await queue.put(msg)

await nc.subscribe("time.>", cb=cb)

```

--------------------------------

### Manage KV Buckets (Go)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This Go code demonstrates how to interact with NATS Key/Value (KV) buckets. It shows functions to bind to an existing KV store, create a new one with specified configurations, and delete an existing KV store. These operations are crucial for managing KV data persistence.

```go
import (
	"github.com/nats-io/nats.go"
)

// KeyValue will lookup and bind to an existing KeyValue store.
func KeyValue(bucket string) (nats.KeyValue, error) {
	// ... implementation details ...
	return nil, nil
}

// CreateKeyValue will create a KeyValue store with the following configuration.
func CreateKeyValue(cfg *nats.KeyValueConfig) (nats.KeyValue, error) {
	// ... implementation details ...
	return nil, nil
}

// DeleteKeyValue will delete this KeyValue store (JetStream stream).
func DeleteKeyValue(bucket string) error {
	// ... implementation details ...
	return nil
}
```

--------------------------------

### Pull NATS Server Docker Image

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This command pulls the latest official NATS server Docker image from Docker Hub. Ensure Docker is installed and running before executing this command.

```bash
docker pull nats
```

--------------------------------

### Unsubscribe from NATS (Python)

Source: https://docs.nats.io/using-nats/developer/receiving/unsubscribing

Provides an example of unsubscribing from NATS in Python using asyncio. After subscribing and publishing a message, the unsubscribe method is called on the subscription object to remove interest in the subject.

```python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

future = asyncio.Future()

async def cb(msg):
  nonlocal future
  future.set_result(msg)

sub = await nc.subscribe("updates", cb=cb)
await nc.publish("updates", b'All is Well')

# Remove interest in subject
await sub.unsubscribe()

# Won't be received...
await nc.publish("updates", b'...')

```

--------------------------------

### C#: Create an Ordered Consumer

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Provides a C# example for creating an ordered consumer in NATS JetStream. It initializes a NATS client, creates a JetStream context, configures and creates a stream, publishes a message, and then creates an ordered consumer for that stream.

```C#
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.JetStream;
using NATS.Client.JetStream.Models;

await using var client = new NatsClient();

INatsJSContext js = client.CreateJetStreamContext();

var streamConfig = new StreamConfig(name: "FOO", subjects: ["foo"]);
await js.CreateStreamAsync(streamConfig);

PubAckResponse ack = await js.PublishAsync("foo", "Hello, JetStream!");
ack.EnsureSuccess();

INatsJSConsumer orderedConsumer = await js.CreateOrderedConsumerAsync("FOO");

using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(10));


```

--------------------------------

### NATS Server Connection Information (/connz) Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

This JSON response provides details about the NATS server's current connections, including server ID, host, port, and information about outbound and inbound gateways. It shows connection metrics like messages sent/received and uptime.

```json
{
  "server_id": "NANVBOU62MDUWTXWRQ5KH3PSMYNCHCEUHQV3TW3YH7WZLS7FMJE6END6",
  "now": "2019-07-24T18:02:55.597398-06:00",
  "name": "region1",
  "host": "2601:283:4601:1350:1895:efda:2010:95a1",
  "port": 4501,
  "outbound_gateways": {
    "region2": {
      "configured": true,
      "connection": {
        "cid": 7,
        "ip": "127.0.0.1",
        "port": 5500,
        "start": "2019-07-24T18:02:48.765621-06:00",
        "last_activity": "2019-07-24T18:02:48.765621-06:00",
        "uptime": "6s",
        "idle": "6s",
        "pending_bytes": 0,
        "in_msgs": 0,
        "out_msgs": 0,
        "in_bytes": 0,
        "out_bytes": 0,
        "subscriptions": 0,
        "name": "NCXBIYWT7MV7OAQTCR4QTKBN3X3HDFGSFWTURTCQ22ZZB6NKKJPO7MN4"
      }
    },
    "region3": {
      "configured": true,
      "connection": {
        "cid": 5,
        "ip": "::1",
        "port": 6500,
        "start": "2019-07-24T18:02:48.764685-06:00",
        "last_activity": "2019-07-24T18:02:48.764685-06:00",
        "uptime": "6s",
        "idle": "6s",
        "pending_bytes": 0,
        "in_msgs": 0,
        "out_msgs": 0,
        "in_bytes": 0,
        "out_bytes": 0,
        "subscriptions": 0,
        "name": "NCVS7Q65WX3FGIL2YQRLI77CE6MQRWO2Y453HYVLNMBMTVLOKMPW7R6K"
      }
    }
  },
  "inbound_gateways": {
    "region2": [
      {
        "configured": false,
        "connection": {
          "cid": 9,
          "ip": "::1",
          "port": 52029,
          "start": "2019-07-24T18:02:48.76677-06:00",
          "last_activity": "2019-07-24T18:02:48.767096-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCXBIYWT7MV7OAQTCR4QTKBN3X3HDFGSFWTURTCQ22ZZB6NKKJPO7MN4"
        }
      }
    ],
    "region3": [
      {
        "configured": false,
        "connection": {
          "cid": 4,
          "ip": "::1",
          "port": 52025,
          "start": "2019-07-24T18:02:48.764577-06:00",
          "last_activity": "2019-07-24T18:02:48.764994-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCVS7Q65WX3FGIL2YQRLI77CE6MQRWO2Y453HYVLNMBMTVLOKMPW7R6K"
        }
      },
      {
        "configured": false,
        "connection": {
          "cid": 8,
          "ip": "127.0.0.1",
          "port": 52026,
          "start": "2019-07-24T18:02:48.766173-06:00",
          "last_activity": "2019-07-24T18:02:48.766999-06:00",
          "uptime": "6s",
          "idle": "6s",
          "pending_bytes": 0,
          "in_msgs": 0,
          "out_msgs": 0,
          "in_bytes": 0,
          "out_bytes": 0,
          "subscriptions": 0,
          "name": "NCKCYK5LE3VVGOJQ66F65KA27UFPCLBPX4N4YOPOXO3KHGMW24USPCKN"
        }
      }
    ]
  }
}
```

--------------------------------

### JavaScript NATS JetStream Consumer Setup

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Sets up a NATS JetStream consumer in JavaScript. It connects to NATS, creates a regular subscription, and then configures a JetStream stream and consumer with explicit acknowledgment policy. The consumer is set up to deliver messages to the subscription.

```javascript
import { AckPolicy, connect } from "../../src/mod.ts";
import { nuid } from "../../nats-base-client/nuid.ts";

const nc = await connect();

// create a regular subscription - this is plain nats
const sub = nc.subscribe("my.messages", { max: 5 });
const done = (async () => {
  for await (const m of sub) {
    console.log(m.subject);
    m.respond();
  }
})();

const jsm = await nc.jetstreamManager();
const stream = nuid.next();
const subj = nuid.next();
await jsm.streams.add({ name: stream, subjects: [`${subj}.>`] });

// create a consumer that delivers to the subscription
await jsm.consumers.add(stream, {
  ack_policy: AckPolicy.Explicit,

```

--------------------------------

### Manage KV Buckets (JavaScript)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This JavaScript code snippet illustrates how to manage NATS Key/Value (KV) buckets. It includes static methods for creating a new KV store with optional configurations and binding to an existing one. It also provides an instance method to destroy (delete) a KV bucket. These functions facilitate dynamic KV store management in a Node.js environment.

```javascript
import { JetStreamClient, KV, KvOptions, KvCodecs } from 'nats';

class KeyValueManager {
  static async create(
    js: JetStreamClient,
    name: string,
    opts: Partial<KvOptions> = {},
  ): Promise<KV> {
    // ... implementation details ...
    return null;
  }

  static async bind(
    js: JetStreamClient,
    name: string,
    opts: Partial<{ codec: KvCodecs }> = {},
  ): Promise<KV> {
    // ... implementation details ...
    return null;
  }

  async destroy(): Promise<boolean> {
    // ... implementation details ...
    return false;
  }
}
```

--------------------------------

### Create NATS Operator with System Account using NSC

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This command creates a new NATS operator and automatically configures a system account for it. The system account is used by nats-server for system services. Ensure you have 'nsc' installed and configured.

```shell
nsc add operator -n <operator-name> --sys
```

--------------------------------

### Create NATS JetStream Stream (Interactive)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/streams

This snippet demonstrates how to add a new NATS JetStream Stream named 'ORDERS' interactively. The command prompts for various configuration details such as subjects, storage, retention, and message limits. It's useful for initial setup and understanding all available options.

```shell
nats str add ORDERS
```
? Subjects to consume ORDERS.*
? Storage backend file
? Retention Policy Limits
? Discard Policy Old
? Message count limit -1
? Message size limit -1
? Maximum message age limit 1y
? Maximum individual message size [? for help] (-1) -1
Stream ORDERS was created

Information for Stream ORDERS

Configuration:

             Subjects: ORDERS.*
     Acknowledgements: true
            Retention: File - Limits
             Replicas: 1
    Maximum Messages: -1
       Maximum Bytes: -1
         Maximum Age: 8760h0m0s
 Maximum Message Size: -1
  Maximum Consumers: -1

Statistics:

            Messages: 0
               Bytes: 0 B
            FirstSeq: 0
             LastSeq: 0
    Active Consumers: 0
```
```

--------------------------------

### Go NKey Authentication

Source: https://docs.nats.io/using-nats/developer/connecting/nkey

Demonstrates how to establish a NATS connection using NKey authentication in Go. It loads the NKey from a seed file and connects to the NATS server. Ensure 'seed.txt' contains your NKey seed.

```Go
opt, err := nats.NkeyOptionFromSeed("seed.txt")
if err != nil {
    log.Fatal(err)
}
nc, err := nats.Connect("127.0.0.1", opt)
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Do something with the connection
```

--------------------------------

### Add NATS JetStream Stream Sourced from Other Streams

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/replication

Create a new NATS JetStream stream named 'ARCHIVE' that sources messages from the 'ORDERS' and 'RETURNS' streams. This setup allows the ARCHIVE stream to capture messages published to its source streams.

```shell
nats s add ARCHIVE --source ORDERS --source RETURNS
```

--------------------------------

### Run NATS Server with Docker

Source: https://docs.nats.io/running-a-nats-service/introduction/running

This command pulls the latest NATS server image from Docker Hub and runs it as a container. It maps the host's port 4222 to the container's port 4222, allowing clients to connect. The output is similar to the standalone version, indicating server startup and listening ports.

```shell
docker run -p 4222:4222 -ti nats:latest
```

--------------------------------

### NATS Server Authorization Configuration with Bcrypt

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli

An example of how to configure the NATS server's authorization section using a bcrypt hashed password. The server uses this hash for verification when a client provides the plaintext password.

```plaintext
authorization {
    user: derek
    password: $2a$11$3kIDaCxw.Glsl1.u5nKa6eUnNDLV5HV9tIuUp7EHhMt6Nm9myW1aS
  }
```

--------------------------------

### Publish and Subscribe with Max Messages (Ruby)

Source: https://docs.nats.io/using-nats/developer/receiving/unsub_after

This Ruby example demonstrates publishing a message and subscribing to a subject with a maximum message limit. It uses NATS.start for connection and Fiber for asynchronous operations.

```ruby
require 'nats/client'
require 'fiber'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  Fiber.new do
    f = Fiber.current

    nc.subscribe("time", max: 1) do |msg, reply|
      f.resume Time.now
    end

    nc.publish("time", 'What is the time?', NATS.create_inbox)

    # Use the response
    msg = Fiber.yield
    puts "Reply: #{msg}"

    # Won't be received
    nc.publish("time", 'What is the time?', NATS.create_inbox)

  end.resume
end
```

--------------------------------

### Publish and Subscribe with NATS in Java

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This Java example demonstrates publishing messages to NATS subjects with time zone details and subscribing to receive them. It utilizes the NATS Java client and handles time zone conversions for message content.

```java
import io.nats.client.*;
import java.nio.charset.StandardCharsets;
import java.time.*;
import java.time.format.DateTimeFormatter;

public class NatsExample {
    public static void main(String[] args) {
        try {
            Connection nc = Nats.connect("nats://demo.nats.io:4222");
            System.out.println("Connected to NATS at " + nc.getConnectedUrl());

            // Publish messages with time zone information for Eastern US
            ZoneId zoneId = ZoneId.of("America/New_York");
            ZonedDateTime zonedDateTime = ZonedDateTime.ofInstant(Instant.now(), zoneId);
            String formatted = zonedDateTime.format(DateTimeFormatter.ISO_ZONED_DATE_TIME);

            nc.publish("time.us.east", formatted.getBytes(StandardCharsets.UTF_8));
            nc.publish("time.us.east.atlanta", formatted.getBytes(StandardCharsets.UTF_8));

            // Publish messages with time zone information for Europe
            zoneId = ZoneId.of("Europe/Warsaw");
            zonedDateTime = ZonedDateTime.ofInstant(Instant.now(), zoneId);
            formatted = zonedDateTime.format(DateTimeFormatter.ISO_ZONED_DATE_TIME);
            nc.publish("time.eu.east", formatted.getBytes(StandardCharsets.UTF_8));
            nc.publish("time.eu.east.warsaw", formatted.getBytes(StandardCharsets.UTF_8));

            System.out.println("Published messages.");

            // Subscribe to receive messages
            Subscription subscription = nc.subscribe("time.>");
            nc.flush(Duration.ZERO);

            int count = 0;
            while (count < 4) {
                Message msg = subscription.nextMessage(Duration.ofSeconds(5)); // Wait up to 5 seconds
                if (msg != null) {
                    System.out.printf("Received %d: %s: %s\n", ++count, msg.getSubject(), new String(msg.getData(), StandardCharsets.UTF_8));
                } else {
                    System.out.println("No more messages received within timeout.");
                    break;
                }
            }

            System.out.println("Done");
            nc.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

--------------------------------

### Allow Multiple Connection Types for MQTT User

Source: https://docs.nats.io/running-a-nats-service/configuration/mqtt/mqtt_config

This example shows how to configure a NATS user ('bar') to accept multiple connection types, specifically 'STANDARD' and 'MQTT', by listing them in the 'allowed_connection_types' field. This allows flexibility for different client types.

```hcl
authorization {
  users [
    {user: foo password: foopwd, permission: {...}}
    {user: bar password: barpwd, permission: {...}, allowed_connection_types: ["STANDARD", "MQTT"]}
  ]
}
```

--------------------------------

### Enable WebSocket on Docker - NATS Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/websocket/websocket_conf

Provides the minimal NATS configuration required to enable WebSocket connections when running NATS within a Docker container. This example specifies the port and disables TLS for simplicity.

```hcl
websocket 
{
     port: 8080
     no_tls: true
}

```

--------------------------------

### NATS JetStream: Consume Messages (C)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Provides examples for consuming messages from a NATS JetStream subscription using pull, asynchronous, and synchronous modes. Includes message acknowledgment and cleanup. Requires the NATS C client library.

```c
    if ((s == NATS_OK) && pull)
    {
        natsMsgList list;
        int         i;

        for (count = 0; (s == NATS_OK) && (count < total); )
        {
            s = natsSubscription_Fetch(&list, sub, 1024, 5000, &jerr);
            if (s != NATS_OK)
                break;

            if (start == 0)
                start = nats_Now();

            count += (int64_t) list.Count;
            for (i=0; (s == NATS_OK) && (i<list.Count); i++)
                s = natsMsg_Ack(list.Msgs[i], &jsOpts);

            natsMsgList_Destroy(&list);
        }
    }
    else if ((s == NATS_OK) && async)
    {
        while (s == NATS_OK)
        {
            if (count + dropped == total)
                break;

            nats_Sleep(1000);
        }
    }
    else if (s == NATS_OK)
    {
        for (count = 0; (s == NATS_OK) && (count < total); count++)
        {
            s = natsSubscription_NextMsg(&msg, sub, 5000);
            if (s != NATS_OK)
                break;

            if (start == 0)
                start = nats_Now();

            s = natsMsg_Ack(msg, &jsOpts);
            natsMsg_Destroy(msg);
        }
    }
```

--------------------------------

### Configure NATS Service User with sc config

Source: https://docs.nats.io/running-a-nats-service/introduction/windows_srv

This command shows how to configure the user account under which the NATS service runs. It's important for setting correct file permissions, especially when using JetStream. This example changes the user to 'NetworkService'.

```shell
sc config "nats-server" obj= "NT AUTHORITY\NetworkService" password= ""
```

--------------------------------

### NATS Server General Information Endpoint Response

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

Example JSON response from the `/varz` endpoint, providing comprehensive details about the NATS server's current state, configuration, and statistics. This includes server ID, version, network settings, connection counts, message throughput, and more.

```json
{
  "server_id": "NACDVKFBUW4C4XA24OOT6L4MDP56MW76J5RJDFXG7HLABSB46DCMWCOW",
  "version": "2.0.0",
  "proto": 1,
  "go": "go1.12",
  "host": "0.0.0.0",
  "port": 4222,
  "max_connections": 65536,
  "ping_interval": 120000000000,
  "ping_max": 2,
  "http_host": "0.0.0.0",
  "http_port": 8222,
  "https_port": 0,
  "auth_timeout": 1,
  "max_control_line": 4096,
  "max_payload": 1048576,
  "max_pending": 67108864,
  "cluster": {},
  "gateway": {},
  "leaf": {},
  "tls_timeout": 0.5,
  "write_deadline": 2000000000,
  "start": "2019-06-24T14:24:43.928582-07:00",
  "now": "2019-06-24T14:24:46.894852-07:00",
  "uptime": "2s",
  "mem": 9617408,
  "cores": 4,
  "gomaxprocs": 4,
  "cpu": 0,
  "connections": 0,
  "total_connections": 0,
  "routes": 0,
  "remotes": 0,
  "leafnodes": 0,
  "in_msgs": 0,
  "out_msgs": 0,
  "in_bytes": 0,
  "out_bytes": 0,
  "slow_consumers": 2,
  "subscriptions": 0,
  "http_req_stats": {
    "/": 0,
    "/connz": 0,
    "/gatewayz": 0,
    "/routez": 0,
    "/subsz": 0,
    "/varz": 1
  },
  "config_load_time": "2019-06-24T14:24:43.928582-07:00",
  "slow_consumer_stats": {
    "clients": 1,
    "routes": 1,
    "gateways": 0,
    "leafs": 0
  }
}
```

--------------------------------

### Publish Messages to US Time Subjects (JavaScript)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

A JavaScript snippet showing how to publish messages to various US time-related NATS subjects. This example assumes a pre-initialized NATS connection object `nc`.

```javascript
nc.publish('time.us.east');
nc.publish('time.us.central');
nc.publish('time.us.mountain');
nc.publish('time.us.west');
```

--------------------------------

### Upload a File to an Object Store Bucket

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This command uploads a local file to the 'myobjbucket'. By default, the file's path is used as its key within the bucket. The progress bar indicates the upload status and completion.

```shell
nats object put myobjbucket ~/Movies/NATS-logo.mov
```

--------------------------------

### Subject Hierarchy Example

Source: https://docs.nats.io/nats-concepts/subjects

Illustrates a hierarchical structure for NATS subjects, commonly used in applications like world clocks to group related messages semantically. The '.' character denotes levels in the hierarchy.

```markup
time.us
time.us.east
time.us.east.atlanta
time.eu.east
time.eu.east.warsaw
```

--------------------------------

### Configure NATS Ping/Pong Settings (Ruby)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Shows how to configure the NATS client in Ruby with a ping interval of 20 and a maximum of 5 outstanding pings. The example uses `NATS.start` and includes event handlers for reconnect and disconnect events.

```Ruby
require 'nats/client'
# Set Ping Interval to 20 seconds and Max Pings Outstanding to 5
NATS.start(ping_interval: 20, max_outstanding_pings: 5) do |nc|
   nc.on_reconnect do
    puts "Got reconnected to #{nc.connected_server}"
  end

  nc.on_disconnect do |reason|
    puts "Got disconnected! #{reason}"
  end

  # Do something with the connection
end

```

--------------------------------

### Send JSON Data with NATS in C#

Source: https://docs.nats.io/using-nats/developer/sending/structure

This C# example uses the `NATS.Net` library to publish a C# record as JSON. It leverages the default `System.Text.Json` serializer and requires the `NATS.Net` NuGet package.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient();

using var cts = new CancellationTokenSource();

Task process = Task.Run(async () =>
{
    // Let's deserialize the message as a UTF-8 string to see
    // the published serialized output in the console
    await foreach (var msg in client.SubscribeAsync<string>("updates", cancellationToken: cts.Token))
    {
        Console.WriteLine($"Received: {msg.Data}");
    }
});

// Wait for the subscription task to be ready
await Task.Delay(1000);

var stock = new Stock { Symbol = "MSFT", Price = 123.45 };

// The default serializer uses System.Text.Json to serialize the object
await client.PublishAsync<Stock>("updates", stock);

// Define the object
public record Stock {
    public string Symbol { get; set; }
    public double Price { get; set; }
}

// Output:
// Received: {"Symbol":"MSFT","Price":123.45}
```

--------------------------------

### Get a Specific Message from a NATS JetStream Stream

Source: https://docs.nats.io/nats-concepts/jetstream/js_walkthrough

Retrieves a specific message from a JetStream stream, typically by its sequence number or subject. This command allows direct access to stored messages.

```shell
nats stream get my_stream
```

--------------------------------

### Asynchronous Subscribe to 'updates' Subject

Source: https://docs.nats.io/using-nats/developer/receiving/async

Subscribes to the 'updates' subject and processes incoming messages asynchronously using callbacks. This pattern is common for real-time notifications. Ensure the NATS client library is installed.

```Go
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
deferr nc.Close()

// Use a WaitGroup to wait for a message to arrive
wg := sync.WaitGroup{}
wg.Add(1)

// Subscribe
if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
    wg.Done()
}); err != nil {
    log.Fatal(err)
}

// Wait for a message to come in
wg.Wait()
```

```Java
Connection nc = Nats.connect("nats://demo.nats.io:4222")

// Use a latch to wait for a message to arrive
CountDownLatch latch = new CountDownLatch(1)

// Create a dispatcher and inline message handler
Dispatcher d = nc.createDispatcher((msg) -> {
    String str = new String(msg.getData(), StandardCharsets.UTF_8)
    System.out.println(str)
    latch.countDown()
})

// Subscribe
d.subscribe("updates")

// Wait for a message to come in
latch.await()

// Close the connection
nc.close()
```

```JavaScript
const sc = StringCodec()
// this is an example of a callback subscription
// https://github.com/nats-io/nats.js/blob/master/README.md#async-vs-callbacks
nc.subscribe("updates", {
  callback: (err, msg) => {
    if (err) {
      t.error(err.message)
    } else {
      t.log(sc.decode(msg.data))
    }
  },
  max: 1,
})

// here's an iterator subscription - note the code in the
// for loop will block until the iterator completes
// either from a break/return from the iterator, an
// unsubscribe after the message arrives, or in this case
// an auto-unsubscribe after the first message is received
const sub = nc.subscribe("updates", { max: 1 })
for await (const m of sub) {
  t.log(sc.decode(m.data))
}

// subscriptions have notifications, simply wait
// the closed promise
sub.closed
  .then(() => {
    t.log("subscription closed")
  })
  .catch((err) => {
    t.err(`subscription closed with an error ${err.message}`)
  })
```

```Python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

future = asyncio.Future()

async def cb(msg):
  nonlocal future
  future.set_result(msg)

await nc.subscribe("updates", cb=cb)
await nc.publish("updates", b'All is Well')
await nc.flush()

# Wait for message to come in
msg = await asyncio.wait_for(future, 1)
```

```C#
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient();

// Subscribe to the "updates" subject and receive messages as <string> type.
// The default serializer understands all primitive types, strings,
// byte arrays, and uses JSON for complex types.
await foreach (var msg in client.SubscribeAsync<string>("updates"))
{
    Console.WriteLine($"Received: {msg.Data}")
    
    if (msg.Data == "exit")
    {
        // When we exit the loop, we unsubscribe from the subject
        // as a result of enumeration completion.
        break
    }
}
```

```Ruby
require 'nats/client'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  nc.subscribe("updates") do |msg|
    puts msg
    nc.close
  end

  nc.publish("updates", "All is Well")
end
```

```C
static void
onMsg(natsConnection *conn, natsSubscription *sub, natsMsg *msg, void *closure)
{
    printf("Received msg: %s - %.*s\n",
           natsMsg_GetSubject(msg),
           natsMsg_GetDataLength(msg),
           natsMsg_GetData(msg))

    // Need to destroy the message!
    natsMsg_Destroy(msg)
}

(...)

natsConnection      *conn = NULL
natsSubscription    *sub  = NULL
natsStatus          s

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL)
if (s == NATS_OK)
{
    // Creates an asynchronous subscription on subject "foo".
    // When a message is sent on subject "foo", the callback
    // onMsg() will be invoked by the client library.
    // You can pass a closure as the last argument.
    s = natsConnection_Subscribe(&sub, conn, "foo", onMsg, NULL)
}

(...)


// Destroy objects that were created
natsSubscription_Destroy(sub)
natsConnection_Destroy(conn)
```

--------------------------------

### Set Operator Service URLs and Publish Message

Source: https://docs.nats.io/using-nats/nats-tools/nsc/managed

This demonstrates setting service URLs for an operator, enabling 'nsc tool' commands like 'pub'. It then shows how to publish a simple 'hello world' message to the 'hello' subject.

```bash
nsc edit operator -n nats://localhost:4222
nsc tool pub hello world
```

--------------------------------

### Configure NATS Stream with Mirror using JSON

Source: https://docs.nats.io/nats-concepts/jetstream/source_and_mirror/source_and_mirror_example

This JSON configuration defines a NATS JetStream stream that mirrors another stream. It includes options for specifying the source mirror name and external API endpoints. This configuration is used with the NATS CLI.

```json
{
  "name": "MIRROR_TARGET"
  "discard": "old",
  "mirror": {
    "name": "MIRROR_ORIGIN"
  },
  "deny_delete": false,
  "sealed": false,
  "max_msg_size": -1,
  "allow_rollup_hdrs": false,
  "max_bytes": -1,
  "storage": "file",
  "allow_direct": false,
  "max_age": 0,
  "max_consumers": -1,
  "max_msgs_per_subject": -1,
  "num_replicas": 1,
  "name": "MIRROR_TARGET",
  "deny_purge": false,
  "compression": "none",
  "max_msgs": -1,
  "retention": "limits",
  "mirror_direct": false
}
```

```json
{
  "name": "MIRROR_TARGET"
  "discard": "old",
  "mirror": {
    "opt_start_time": "2024-07-11T08:57:20.4441646Z",
    "external": {
      "deliver": "",
      "api": "$JS.domainB.API"
    },
    "name": "MIRROR_ORIGIN"
  },
  "consumer_limits": {
    
  },
  "deny_delete": false,
  "sealed": false,
  "max_msg_size": -1,
  "allow_rollup_hdrs": false,
  "max_bytes": -1,
  "storage": "file",
  "allow_direct": false,
  "max_age": 0,
  "max_consumers": -1,
  "max_msgs_per_subject": -1,
  "num_replicas": 1,
  "name": "MIRROR_TARGET",
  "deny_purge": false,
  "compression": "none",
  "max_msgs": -1,
  "retention": "limits",
  "mirror_direct": false
}
```

--------------------------------

### Configure NATS Ping/Pong Settings (JavaScript)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Shows how to configure the NATS client in JavaScript with a ping interval of 20 seconds (20000 milliseconds) and a maximum of 5 outstanding pings. The example uses an async function to establish the connection.

```JavaScript
// Set Ping Interval to 20 seconds and Max Pings Outstanding to 5
const nc = await connect({
    pingInterval: 20 * 1000,
    maxPingOut: 5,
    servers: ["demo.nats.io:4222"],
});

```

--------------------------------

### Initialize Managed Operator with 'nsc init'

Source: https://docs.nats.io/using-nats/nats-tools/nsc/managed

This command initializes a new 'nsc' environment with a specified managed operator and a default account name. It's a convenient way to start using a managed operator like 'synadia'.

```bash
nsc init -o synadia -n MyFirstAccount
```

--------------------------------

### NATS JetStream Consumer Acknowledgement Functions (Python)

Source: https://docs.nats.io/using-nats/developer/anatomy

Provides examples of acknowledgment functions for NATS JetStream messages in Python. These functions enable robust message handling by communicating processing status back to the NATS server.

```Python
import asyncio
from nats.aio.client import Client as NATS

# Assuming 'msg' is a received Msg object from a JetStream subscription

async def handle_message(msg):
    # Positive acknowledgement
    await msg.ack()

    # Negative acknowledgement (with redelivery)
    await msg.nak()

    # Negative acknowledgement with delay (e.g., 1 minute)
    await msg.nak(delay=60)

    # Terminate message (for unrecoverable errors)
    await msg.term()

    # Indicate message is in progress (useful for long processing times)
    await msg.in_progress()

```

--------------------------------

### Python JetStream Initialization

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

Shows how to connect to NATS and initialize the JetStream context for stream operations.

```APIDOC
## Python JetStream API

### Initialize JetStream

- **Description**: Connects to a NATS server and creates a JetStream context for interacting with JetStream features.
- **Method**: N/A (Code example provided)
- **Endpoint**: N/A (NATS Python client methods)

#### `nats.connect(url)`
- **Description**: Establishes a connection to the NATS server.
- **Parameters**:
  - `url` (string) - The URL of the NATS server (e.g., "localhost").

#### `nc.jetstream()`
- **Description**: Returns a JetStream context object from the NATS connection.

#### `js.add_stream(stream_config)`
- **Description**: Adds a new stream to JetStream.
  - **Request Body**: `stream_config` (object) - Configuration for the stream. Example: `{ name: "sample-stream", subjects: ["foo"] }`.
```

--------------------------------

### NATS Account Configuration Example

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/accounts

This configuration defines two isolated accounts, 'A' and 'B', each with a single user. Messages published within one account are not visible to users in another, demonstrating fundamental multi-tenancy isolation in NATS.

```NATS Configuration
accounts: {
    A: {
        users: [
            {user: a, password: a}
        ]
    },
    B: {
        users: [
            {user: b, password: b}
        ]
    },
}
```

--------------------------------

### Handling String vs. Number Parsing in NATS Config

Source: https://docs.nats.io/nats-server/configuration

Demonstrates how NATS configuration handles string and number parsing, specifically addressing potential issues with string values that start with digits. Quoting is recommended to ensure values are treated as strings.

```plaintext
listen: 127.0.0.1:4222
authorization: {
    # Bad - Number parsing error
    token: 3secret
}
```

```plaintext
listen: 127.0.0.1:4222
authorization: {
    # Good
    token: "3secret"
}
```

--------------------------------

### Watch NATS Object Bucket Changes

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

Monitors a specified NATS object store bucket for any changes, such as puts or deletes. This command is useful for real-time tracking of bucket activity. It requires the NATS CLI to be installed and configured.

```shell
nats object watch myobjbucket
```

--------------------------------

### Verify NATS Server is Running via Telnet

Source: https://docs.nats.io/running-a-nats-service/nats_docker/nats-docker-tutorial

Tests the NATS server's client connection port (4222) using telnet. This confirms the server is accessible and provides initial server information upon successful connection.

```bash
telnet localhost 4222
```

--------------------------------

### NATS Cluster Docker Compose Output

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This output shows the successful creation and startup of the NATS cluster containers. It details the progress from container creation to server readiness and route connections being established between cluster nodes.

```text
[+] Running 3/3
 ⠿ Container xxx_nats_1    Created
 ⠿ Container xxx_nats-1_1  Created
 ⠿ Container xxx_nats-2_1  Created
Attaching to nats-1_1, nats-2_1, nats_1
nats_1    | [1] 2021/09/28 10:42:36.742844 [INF] Starting nats-server
nats_1    | [1] 2021/09/28 10:42:36.742898 [INF]   Version:  2.6.1
nats_1    | [1] 2021/09/28 10:42:36.742913 [INF]   Git:      [c91f0fe]
nats_1    | [1] 2021/09/28 10:42:36.742929 [INF]   Name:     NCZIIQ6QT4KT5K5WBP7H2RRBM4MSYD4C2TVSRZOZN57EHX6VTF4EWXAU
nats_1    | [1] 2021/09/28 10:42:36.742954 [INF]   ID:       NCZIIQ6QT4KT5K5WBP7H2RRBM4MSYD4C2TVSRZOZN57EHX6VTF4EWXAU
nats_1    | [1] 2021/09/28 10:42:36.745289 [INF] Starting http monitor on 0.0.0.0:8222
nats_1    | [1] 2021/09/28 10:42:36.745737 [INF] Listening for client connections on 0.0.0.0:4222
nats_1    | [1] 2021/09/28 10:42:36.750381 [INF] Server is ready
nats_1    | [1] 2021/09/28 10:42:36.750669 [INF] Cluster name is NATS
nats_1    | [1] 2021/09/28 10:42:36.751444 [INF] Listening for route connections on 0.0.0.0:6222
nats-1_1  | [1] 2021/09/28 10:42:37.709888 [INF] Starting nats-server
nats-1_1  | [1] 2021/09/28 10:42:37.709977 [INF]   Version:  2.6.1
nats-1_1  | [1] 2021/09/28 10:42:37.709999 [INF]   Git:      [c91f0fe]
nats-1_1  | [1] 2021/09/28 10:42:37.710023 [INF]   Name:     NBHTXXY3HYZVPXITYQ73BSDA5CQZINTKYRM23XFI46RWWTTUP5TAXQMB
nats-1_1  | [1] 2021/09/28 10:42:37.710042 [INF]   ID:       NBHTXXY3HYZVPXITYQ73BSDA5CQZINTKYRM23XFI46RWWTTUP5TAXQMB
nats-1_1  | [1] 2021/09/28 10:42:37.711646 [INF] Listening for client connections on 0.0.0.0:4222
nats-1_1  | [1] 2021/09/28 10:42:37.712197 [INF] Server is ready
nats-1_1  | [1] 2021/09/28 10:42:37.712376 [INF] Cluster name is NATS
nats-1_1  | [1] 2021/09/28 10:42:37.712469 [INF] Listening for route connections on 0.0.0.0:6222
nats_1    | [1] 2021/09/28 10:42:37.718918 [INF] 172.18.0.4:52950 - rid:4 - Route connection created
nats-1_1  | [1] 2021/09/28 10:42:37.719906 [INF] 172.18.0.3:6222 - rid:4 - Route connection created
nats-2_1  | [1] 2021/09/28 10:42:37.731357 [INF] Starting nats-server
nats-2_1  | [1] 2021/09/28 10:42:37.731518 [INF]   Version:  2.6.1
nats-2_1  | [1] 2021/09/28 10:42:37.731531 [INF]   Git:      [c91f0fe]
nats-2_1  | [1] 2021/09/28 10:42:37.731543 [INF]   Name:     NCG6UQ2N3IHE6OS76TL46RNZBAPHNUCQSA64FDFHG5US2LLJOQLD5ZK2
nats-2_1  | [1] 2021/09/28 10:42:37.731554 [INF]   ID:       NCG6UQ2N3IHE6OS76TL46RNZBAPHNUCQSA64FDFHG5US2LLJOQLD5ZK2
nats-2_1  | [1] 2021/09/28 10:42:37.732893 [INF] Listening for client connections on 0.0.0.0:4222
nats-2_1  | [1] 2021/09/28 10:42:37.733431 [INF] Server is ready
nats-2_1  | [1] 2021/09/28 10:42:37.733491 [INF] Cluster name is NATS
nats-2_1  | [1] 2021/09/28 10:42:37.733835 [INF] Listening for route connections on 0.0.0.0:6222
nats_1    | [1] 2021/09/28 10:42:37.740860 [INF] 172.18.0.4:52950 - rid:5 - Route connection created
nats-2_1  | [1] 2021/09/28 10:42:37.741557 [INF] 172.18.0.3:6222 - rid:4 - Route connection created
nats-1_1  | [1] 2021/09/28 10:42:37.743981 [INF] 172.18.0.5:6222 - rid:5 - Route connection created
nats-2_1  | [1] 2021/09/28 10:42:37.744332 [INF] 172.18.0.4:40250 - rid:5 - Route connection created
```

--------------------------------

### NATS Configuration with Encryption Enabled

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_callout

This configuration extends the basic account setup by incorporating an `xkey` for encrypting request payloads. This is a recommended security practice for NATS auth callout, ensuring sensitive data is protected during transmission.

```nats
accounts {
  AUTH: {
    users: [ { user: auth, password: auth } ]
  }
  APP: {}
  SYS: {}
}
system_account: SYS

authorization {
  auth_callout {
    issuer: ABJHLOVMPA4CI6R5KLNGOB4GSLNIY7IOUPAJC4YFNDLQVIOBYQGUWVLA
    auth_users: [ auth ]
    account: AUTH
    xkey: XAMHJVPKHHPYZQQM2IVWXKJH36KDDZZMSJ32QKSQBUODFX4I4HARO4GL
  }
}
```

--------------------------------

### NATS Client Request and Receive

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This Go code snippet shows a NATS client that connects to a NATS server and sends periodic requests to the 'test' topic. It uses 'nats.io/nats.go' and specifies user credentials. The client logs any errors during the request or prints the received response data.

```go
package main

import (
    "log"
    "time"

    "github.com/nats-io/nats.go"
)

func main() {
    nc, err := nats.Connect("127.0.0.1:4223", nats.UserCredentials("path/to/user.creds"))
    if err != nil {
        log.Fatal(err)
    }

    for range time.NewTicker(1 * time.Second).C {
        resp, err := nc.Request("test", []byte("test"), 1*time.Second)
        if err != nil {
            log.Println("[Error]", err)
            continue
        }
        log.Println("[Received]", string(resp.Data))
    }
}
```

--------------------------------

### NATS Server Log Output - Third Server Connecting

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Log output for the third NATS server starting up and connecting to the cluster. It shows client and route listening addresses, attempts to connect to the seed server, and successful route establishment.

```log
[83331] 2020/02/12 16:05:12.838022 [INF] Listening for client connections on 0.0.0.0:6222
[83331] 2020/02/12 16:05:12.838029 [INF] Server id is NBE7SLUDLFIMHS2U6347N3DQEJ
[83331] 2020/02/12 16:05:12.838031 [INF] Server is ready
...
[83331] 2020/02/12 16:05:12.839203 [INF] Listening for route connections on localhost:6248
[83331] 2020/02/12 16:05:12.839453 [DBG] Trying to connect to route on localhost:4248
[83331] 2020/02/12 16:05:12.840112 [DBG] 127.0.0.1:4248 - rid:1 - Route connect msg sent
[83331] 2020/02/12 16:05:12.840198 [INF] 127.0.0.1:4248 - rid:1 - Route connection created
[83331] 2020/02/12 16:05:12.840324 [DBG] 127.0.0.1:4248 - rid:1 - Registering remote route "NDSGCS74MG5ZUMBOVWOUJ5S3HIOW"
[83331] 2020/02/12 16:05:12.840342 [DBG] 127.0.0.1:4248 - rid:1 - Sent local subscriptions to route
[83331] 2020/02/12 16:05:12.840717 [INF] 127.0.0.1:62946 - rid:2 - Route connection created
[83331] 2020/02/12 16:05:12.840906 [DBG] 127.0.0.1:62946 - rid:2 - Registering remote route "NAABC2CKRVPZBIECMLZZA6L3PK"
[83331] 2020/02/12 16:05:12.840915 [DBG] 127.0.0.1:62946 - rid:2 - Sent local subscriptions to route
```

--------------------------------

### Send JSON Data with NATS in Python

Source: https://docs.nats.io/using-nats/developer/sending/structure

This Python example shows publishing a JSON-encoded dictionary to a NATS subject using the `nats` library. It requires the `nats-python` package and handles UTF-8 encoding.

```python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

await nc.publish("updates", json.dumps({"symbol": "GOOG", "price": 1200 }).encode())
```

--------------------------------

### NATS Server Log Output - Seed Accepting Route

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

Example log output from the NATS seed server showing that it accepted a route connection from another NATS server. It indicates the connection creation and the registration of the remote route.

```log
[83329] 2020/02/12 16:05:09.663386 [INF] 127.0.0.1:62941 - rid:1 - Route connection created
[83329] 2020/02/12 16:05:09.663665 [DBG] 127.0.0.1:62941 - rid:1 - Registering remote route "NAABC2CKRVPZBIECMLZZA6L3PK"
[83329] 2020/02/12 16:05:09.663681 [DBG] 127.0.0.1:62941 - rid:1 - Sent local subscriptions to route
```

--------------------------------

### Send Request to NATS System (Server Side)

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This command starts a NATS replier that listens for requests on the subject 'q' and responds with '42'. This is used on the main NATS server to demonstrate message handling when a leaf node connects.

```bash
nats reply -s nats://s3cr3t@localhost q 42
```

--------------------------------

### NATS Weighted Mapping for Traffic Shaping in Testing

Source: https://docs.nats.io/nats-concepts/subject_mapping

Demonstrates how to configure NATS weighted mappings to shape traffic for testing purposes. This example routes a percentage of traffic to a failure simulation endpoint based on wildcard subjects.

```yaml
myservice.requests.*: [{ destination: myservice.requests.{{wildcard(1)}}, weight: 80% }, { destination: myservice.requests.fail.{{wildcard(1)}}, weight: 20% }
```

--------------------------------

### Download an Object from a Bucket

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This command retrieves a specific object from 'myobjbucket' and saves it to the local filesystem. By default, it saves the object using its name relative to the current directory. The output indicates the download size and speed.

```shell
nats object get myobjbucket ~/Movies/NATS-logo.mov
```

--------------------------------

### Pull Consumer Fetch (Node.js)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This Node.js example illustrates how a pull consumer makes `fetch()` calls in a dispatch loop to explicitly request messages from the server. This provides the client with full control over message retrieval and flow control.

```javascript
const nc = await connect({
  servers: "nats://localhost:4222",
});
const jsm = await nats.js.streaming(nc);

const sub = await jsm.subscribe("updates", {
  ia: true, // Makes the subscription pull-based
});

// Fetch messages in a loop
(async () => {
  for await (const msg of sub) {
    // Process the message
    console.log(`Received message: ${msg.data.toString()}`);
    // Acknowledge the message
    msg.ack();
  }
})();

// To fetch messages manually:
// const messages = await sub.fetch({ max_messages: 10 });
// messages.forEach(msg => {
//   console.log(`Fetched message: ${msg.data.toString()}`);
//   msg.ack();
// });
```

--------------------------------

### NATS Configuration: Initial Canary Deployment (v1)

Source: https://docs.nats.io/running-a-nats-service/configuration/configuring_subject_mapping

This configuration directs 100% of traffic to the first version of a service, typically used at the start of a canary deployment. It assumes applications request `myservice.requests` and version 1 responders subscribe to `myservice.requests.v1`.

```NATS Configuration
myservice.requests: [
  { destination: myservice.requests.v1, weight: 100% }
]
```

--------------------------------

### NATS.IO CLI for System Event Publishing

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts

This command publishes a message to the 'foo' subject with the payload 'bar', using the 'a' account credentials. It connects to the NATS server running on localhost.

```bash
nats pub -s "nats://a:a@localhost:4222" foo bar
```

--------------------------------

### Connect to NATS Cluster in C

Source: https://docs.nats.io/using-nats/developer/connecting/cluster

Illustrates connecting to a NATS cluster using the C NATS client library. It involves creating `natsOptions`, setting the server list, and then establishing the connection. Error handling is included, and cleanup functions are shown for destroying the created objects.

```C
natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;
const char          *servers[] = {"nats://127.0.0.1:1222", "nats://127.0.0.1:1223", "nats://127.0.0.1:1224"};

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_SetServers(opts, servers, 3);
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### Define Well-Known Operator with Environment Variable

Source: https://docs.nats.io/using-nats/nats-tools/nsc/managed

This example demonstrates how to define a 'well-known' operator named 'zoom' by setting an environment variable. 'nsc' will use the provided URL to fetch the operator's JWT.

```bash
export nsc_zoom_operator=https://account-server-host/jwt/v1/operator
```

--------------------------------

### NATS JetStream: Add and Get Stream Info (C)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Demonstrates adding a stream with configuration and then retrieving its information. Handles potential errors during these operations and ensures proper cleanup of stream info objects. Requires the NATS C client library.

```c
            cfg.Storage = js_MemoryStorage;
            // Add the stream,
            s = js_AddStream(&si, js, &cfg, NULL, &jerr);
        }
        if (s == NATS_OK)
        {
            printf("Stream %s has %" PRIu64 " messages (%" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            // Need to destroy the returned stream object.
            jsStreamInfo_Destroy(si);
        }
    }

    if (s == NATS_OK)
    {
        jsStreamInfo *si = NULL;

        // Let's report some stats after the run
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_OK)
        {
            printf("\nStream %s has %" PRIu64 " messages (%" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            jsStreamInfo_Destroy(si);
        }
        if (delStream)
        {
            printf("\nDeleting stream %s: ", stream);
            s = js_DeleteStream(js, stream, NULL, &jerr);
            if (s == NATS_OK)
                printf("OK!");
            printf("\n");
        }
```

--------------------------------

### Discovering Servers + Stats

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

Discovers servers in the cluster and provides a small health summary for each.

```APIDOC
## GET $SYS.REQ.SERVER.PING

### Description
To discover servers in the cluster, and get a small health summary, publish a request to `$SYS.REQ.SERVER.PING`. Note that while the example below uses `nats-req`, only the first answer for the request will be printed. You can easily modify the example to wait until no additional responses are received for a specific amount of time, thus allowing for all responses to be collected.

### Method
REQUEST

### Endpoint
$SYS.REQ.SERVER.PING

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```shell
nats request --creds ~/.nkeys/SAOP/accounts/SYS/users/SYSU.creds $SYS.REQ.SERVER.PING ""
```

### Response
#### Success Response (200)
- **server** (object) - Server information.
  - **host** (string) - The host address of the server.
  - **id** (string) - The unique identifier of the server.
  - **ver** (string) - The NATS server version.
  - **seq** (integer) - Sequence number.
  - **time** (string) - Timestamp of the server information.
- **statsz** (object) - Server statistics.
  - **start** (string) - Server start time.
  - **mem** (integer) - Memory usage in bytes.
  - **cores** (integer) - Number of CPU cores.
  - **cpu** (integer) - CPU utilization percentage.
  - **connections** (integer) - Current number of connections.
  - **total_connections** (integer) - Total connections since start.
  - **active_accounts** (integer) - Number of active accounts.
  - **subscriptions** (integer) - Number of active subscriptions.
  - **sent** (object) - Sent message statistics.
    - **msgs** (integer) - Number of messages sent.
    - **bytes** (integer) - Number of bytes sent.
  - **received** (object) - Received message statistics.
    - **msgs** (integer) - Number of messages received.
    - **bytes** (integer) - Number of bytes received.
  - **slow_consumers** (integer) - Number of slow consumers.

#### Response Example
```json
{
  "server": {
    "host": "0.0.0.0",
    "id": "NCZQDUX77OSSTGN2ESEOCP4X7GISMARX3H4DBGZBY34VLAI4TQEPK6P6",
    "ver": "2.0.0-RC9",
    "seq": 47,
    "time": "2019-05-02T14:02:46.402166-05:00"
  },
  "statsz": {
    "start": "2019-05-02T13:41:01.113179-05:00",
    "mem": 12922880,
    "cores": 20,
    "cpu": 0,
    "connections": 2,
    "total_connections": 2,
    "active_accounts": 1,
    "subscriptions": 10,
    "sent": {
      "msgs": 7,
      "bytes": 2761
    },
    "received": {
      "msgs": 0,
      "bytes": 0
    },
    "slow_consumers": 0
  }
}
```
```

--------------------------------

### List Objects in a Bucket

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This command lists all objects currently stored within the specified bucket ('myobjbucket'). The output includes the object's name, size, and modification time, providing an overview of the bucket's contents.

```shell
nats object ls myobjbucket
```

--------------------------------

### NATS JetStream Stream Subject Transform Configuration

Source: https://docs.nats.io/nats-concepts/subject_mapping

A JSON configuration example for a NATS JetStream stream, illustrating the 'subject_transform' and 'republish' directives. This defines how subjects are transformed when imported into or republished from the stream.

```json
{
  "name": "orders",
  "subjects": [ "orders.local.*"],
  "subject_transform":{"src":"orders.local.*","dest":"orders.{{wildcard(1)}}"},
  "retention": "limits",
  ...
  "republish": [
    {
      "src": "orders.*",
      "dest": "orders.trace.{{wildcard(1)}}""
    }
  ]
,
```

--------------------------------

### JavaScript JQuery Example for NATS Monitoring

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

This JavaScript code snippet demonstrates how to use jQuery's `$.getJSON` method to fetch monitoring data from a NATS endpoint. It utilizes the `callback` parameter for JSONP requests, enabling cross-domain data retrieval for single-page web applications.

```javascript
$.getJSON("https://demo.nats.io:8222/connz?callback=?", function (data) {
  console.log(data);
});
```

--------------------------------

### Responding to NATS JetStream Flow Control Messages

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Illustrates how to respond to a flow control message from NATS JetStream, which is crucial for preventing consumer stalls. The example shows a generic way to respond, typically by sending nil.

```go
msg.Respond(nil) // Language equivalent for sending nil to the reply subject
```

--------------------------------

### Subject Hierarchy Example - Identifiers

Source: https://docs.nats.io/nats-concepts/subjects

Illustrates using the final tokens of a subject to represent specific identifiers, such as server or application IDs. This practice helps in uniquely addressing resources within a NATS messaging system.

```shell
service.deploy.server-acme.app123
```

--------------------------------

### NATS CLI Message Deduplication Example

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/model_deep_dive

Demonstrates how to use the NATS CLI to send messages with a `Nats-Msg-Id` header to test message deduplication. This ensures that duplicate messages, identified by the same `Nats-Msg-Id`, are only stored once in the stream.

```shell
nats req -H Nats-Msg-Id:1 ORDERS.new hello1
nats req -H Nats-Msg-Id:1 ORDERS.new hello2
nats req -H Nats-Msg-Id:1 ORDERS.new hello3
nats req -H Nats-Msg-Id:1 ORDERS.new hello4
```

--------------------------------

### Watch KV Bucket Changes (Shell)

Source: https://docs.nats.io/nats-concepts/jetstream/key-value-store/kv_walkthrough

Starts a watcher on a KV bucket to receive real-time updates on changes. By default, it only remembers the last change. The output shows the type of operation (DEL, PUT), bucket name, key, and value (if applicable).

```shell
nats kv watch my-kv
```

--------------------------------

### NATS Seed Node Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This snippet shows the basic configuration for a NATS seed node. It defines the ports for client connections and cluster communication. The configuration is typically loaded using the '-config' flag when starting the nats-server.

```plaintext
listen: 127.0.0.1:4222
http: 8222

cluster {
  listen: 127.0.0.1:4248
}
```

--------------------------------

### Create a NATS JetStream Stream

Source: https://docs.nats.io/nats-concepts/jetstream/js_walkthrough

Adds a new JetStream stream named 'my_stream' that will capture messages on the 'foo' subject. It prompts for configuration details and uses defaults for others.

```shell
nats stream add my_stream
? Subjects foo
? Storage file
? Replication 1
? Retention Policy Limits
? Discard Policy Old
? Stream Messages Limit -1
? Per Subject Messages Limit -1
? Total Stream Size -1
? Message TTL -1
? Max Message Size -1
? Duplicate tracking time window 2m0s
? Allow message Roll-ups No
? Allow message deletion Yes
? Allow purging subjects or the entire stream Yes
Stream my_stream was created
```

--------------------------------

### Java NATS Subscription and Drain Example

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Demonstrates creating a NATS dispatcher, subscribing to a subject, processing incoming messages, and draining the subscription to ensure all messages are received before closing the connection. Requires the NATS Java client library.

```java
Dispatcher d = nc.createDispatcher((msg) -> {
    String str = new String(msg.getData(), StandardCharsets.UTF_8);
    System.out.println(str);
    latch.countDown();
});

// Subscribe
d.subscribe("updates");

// Wait for a message to come in
latch.await();

// Messages that have arrived will be processed
CompletableFuture<Boolean> drained = d.drain(Duration.ofSeconds(10));

// Wait for the drain to complete
drained.get();

// Close the connection
nc.close();
```

--------------------------------

### Echo Function in Rust for WebAssembly (WASI)

Source: https://docs.nats.io/using-nats/nex/getting-started/building-function

A Rust WebAssembly function for Nex that concatenates the trigger subject with the incoming payload. It reads input from stdin and writes the combined output to stdout, following the WASI command pattern. This example requires a `wasm32-wasi` target.

```rust
use std::{env, io::{self, Read, Write}};

fn main() {
    let args: Vec<String> = env::args().collect();

    // When a WASI trigger executes:
    // argv[1] is the subject on which it was triggered
    // stdin bytes is the raw input payload
    // stdout bytes is the raw output payload

    let mut buf = Vec::new();
    io::stdin().read_to_end(&mut buf).unwrap();
    
    let mut subject = args[1].as_bytes().to_vec();
    buf.append(&mut subject);

    // This just returns the payload concatenated with the
    // subject
    
    io::stdout().write_all(&mut buf).unwrap();
}
```

--------------------------------

### Retrieve KeyValue Entries (Get)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This section demonstrates how to retrieve key-value entries from a JetStream store. It covers fetching the latest value for a given key and retrieving a specific revision of a value. Error handling for communication issues and API-specific exceptions is included.

```go
// Get returns the latest value for the key.
Get(key string) (entry KeyValueEntry, err error)
// GetRevision returns a specific revision value for the key.
GetRevision(key string, revision uint64) (entry KeyValueEntry, err error)
```

```java
/**
* Get the entry for a key
* @param key the key
* @return the KvEntry object or null if not found.
* @throws IOException covers various communication issues with the NATS
*         server such as timeout or interruption
* @throws JetStreamApiException the request had an error related to the data
* @throws IllegalArgumentException the server is not JetStream enabled
*/
KeyValueEntry get(String key) throws IOException, JetStreamApiException;

/**
* Get the specific revision of an entry for a key.
* @param key the key
* @param revision the revision
* @return the KvEntry object or null if not found.
* @throws IOException covers various communication issues with the NATS
*         server such as timeout or interruption
* @throws JetStreamApiException the request had an error related to the data
* @throws IllegalArgumentException the server is not JetStream enabled
*/
KeyValueEntry get(String key, long revision) throws IOException, JetStreamApiException;
```

```javascript
async get(k: string): Promise<KvEntry | null>
```

```python
async def get(self, key: str) -> Entry:
   """
   get returns the latest value for the key.
   """
```

```csharp
// dotnet add package NATS.Net

// Get an entry from the bucket using the key
ValueTask<NatsKVEntry<T>> GetEntryAsync<T>(string key, ulong revision = default, INatsDeserialize<T>? serializer = default, CancellationToken cancellationToken = default);


```

```c
NATS_EXTERN natsStatus 	kvStore_Get (kvEntry **new_entry, kvStore *kv, const char *key)
 	Returns the latest entry for the key.
 
NATS_EXTERN natsStatus 	kvStore_GetRevision (kvEntry **new_entry, kvStore *kv, const char *key, uint64_t revision)
 	Returns the entry at the specific revision for the key.
```

--------------------------------

### Generate NATS Server Config with Memory Resolver

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This command generates a NATS server configuration file that utilizes the memory resolver. The `--mem-resolver` flag enables the embedding of account JWTs directly into the server configuration, simplifying setups with a limited number of accounts. The output is saved to the specified `--config-file`.

```shell
nsc generate config --mem-resolver --config-file /tmp/server.conf
```

--------------------------------

### Publish Messages with Error Handling (C)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

This C example demonstrates publishing messages to a NATS JetStream stream, including error handling for asynchronous publishes. It showcases how to set up a custom error handler for publish acknowledgements. Requires the NATS C client library.

```c
#include "examples.h"

static const char *usage = "\"
-stream        stream name (default is 'foo')\n"
-txt           text to send (default is 'hello')\n"
-count         number of messages to send\n"
-sync          publish synchronously (default is async)\n";

static void
_jsPubErr(jsCtx *js, jsPubAckErr *pae, void *closure)
{
    int *errors = (int*) closure;

    printf("Error: %u - Code: %u - Text: %s\n", pae->Err, pae->ErrCode, pae->ErrText);
    printf("Original message: %.*s\n", natsMsg_GetDataLength(pae->Msg), natsMsg_GetData(pae->Msg));

    *errors = (*errors + 1);

    // If we wanted to resend the original message, we would do something like that:
    //
    // js_PublishMsgAsync(js, &(pae->Msg), NULL);
    //
    // Note that we use `&(pae->Msg)` so that the library set it to NULL if it takes
    // ownership, and the library will not destroy the message when this callback returns.

    // No need to destroy anything, everything is handled by the library.
}

int main(int argc, char **argv)
{
    natsConnection      *conn  = NULL;
    natsStatistics      *stats = NULL;
    natsOptions         *opts  = NULL;
    jsCtx               *js    = NULL;
    jsOptions           jsOpts;
    jsErrCode           jerr   = 0;
    natsStatus          s;
    int                 dataLen=0;
    volatile int        errors = 0;
    bool                delStream = false;

    opts = parseArgs(argc, argv, usage);
    dataLen = (int) strlen(payload);

    s = natsConnection_Connect(&conn, opts);

    if (s == NATS_OK)
        s = jsOptions_Init(&jsOpts);

    if (s == NATS_OK)
    {
        if (async)
        {
            jsOpts.PublishAsync.ErrHandler           = _jsPubErr;
            jsOpts.PublishAsync.ErrHandlerClosure    = (void*) &errors;
        }
        s = natsConnection_JetStream(&js, conn, &jsOpts);
    }

    if (s == NATS_OK)
    {
        jsStreamInfo    *si = NULL;

        // First check if the stream already exists.
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_NOT_FOUND)
        {
            jsStreamConfig  cfg;

            // Since we are the one creating this stream, we can delete at the end.
            delStream = true;

```

--------------------------------

### NATS Operator and Account Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This configuration sets up a NATS operator with specified accounts and resolvers. It includes JWTs for the operator, system account, and individual user accounts ('A', 'B', 'SYS'). The resolver is set to MEMORY and preloaded with account configurations.

```hcl
operator = "eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJQNDJBSkFTVVA0TUdGRU1EQzVCRVVGUkM1MlQ1M05OTzRIWkhRNEdETVk0S0RZTFVRV0JBIiwiaWF0IjoxNTc0Mzc1OTE2LCJpc3MiOiJPQ09KSk5aSUNINkNHUU5LM1NRMzVXTFpXWkpDUkRBTFJIWjZPVzQ0RFpZVVdNVVYzV1BSSEZSRCIsIm5hbWUiOiJLTyIsInN1YiI6Ik9DT0pKTlpJQ0g2Q0dRTkszU1EzNVdMWldaSkNSREFMUkhaNk9XNDREWllVV01VVjNXUFJIRlJEIiwidHlwZSI6Im9wZXJhdG9yIiwibmF0cyI6e319.pppa9-xhWXJLSCCtqj_dqlvXKR7WlVCh0cqoZ6lr8zg3WlWM8U0bNf6FHw_67-wRS7jj0n4PuA0P0MAJdE3pDA"

system_account = "ABKOWIYVTHNEK5HELPWLAT2CF2CUPELIK4SZH2VCJHLFU22B5U2IIZUO"

resolver = MEMORY

resolver_preload = {
  # Account "A"
  ADFB2JXYTXOJEL6LNAXDREUGRX35BOLZI3B4PFFAC7IRPR3OA4QNKBN2: "eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiIyQjNYNUJPQzQ3M05TU0hFWUVHTzJYRUhaTVNGRVBORFFBREJXWkJLRVdQVlg2TUlZU1JRIiwiaWF0IjoxNTc0Mzc1OTE2LCJpc3MiOiJPQ09KSk5aSUNINkNHUU5LM1NRMzVXTFpXWkpDUkRBTFJIWjZPVzQ0RFpZVVdNVVYzV1BSSEZSRCIsIm5hbWUiOiJBIiwic3ViIjoiQURGQjJKWFlUWE9KRUw2TE5BWERSRVVHUlgzNUJPTFpJM0I0UEZGQUM3SVJQUjNPQTRRTktCTjIiLCJ0eXBlIjoiYWNjb3VudCIsIm5hdHMiOnsiZXhwb3J0cyI6W3sibmFtZSI6InRlc3QiLCJzdWJqZWN0IjoidGVzdCIsInR5cGUiOiJzZXJ2aWNlIiwic2VydmljZV9sYXRlbmN5Ijp7InNhbXBsaW5nIjoxMDAsInJlc3VsdHMiOiJsYXRlbmN5Lm9uLnRlc3QifX1dLCJsaW1pdHMiOnsic3VicyI6LTEsImNvbm4iOi0xLCJsZWFmIjotMSwiaW1wb3J0cyI6LTEsImV4cG9ydHMiOi0xLCJkYXRhIjotMSwicGF5bG9hZCI6LTEsIndpbGRjYXJkcyI6dHJ1ZX19fQ.zZBetgDN6nCFDVpwzF_124BPkc8amGPDnrOmiKUa12xski5zskUI0Y0OeIa1vTo0bkHIKTgM2QDYpmXUQOHnAQ"

  # Account "B"
  ACWOMQA7PZTKJSBTR7BF6TBK3D776734PWHWDKO7HFMQOM5BIOYPSYZZ: "eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJBTFNFQkZGWDZMR0pQTlVMU1NXTDNTRTNISkk2WUZSWlVKSDNLV0E1VE41WUtWRE5MVTJRIiwiaWF0IjoxNTc0Mzc1OTE2LCJpc3MiOiJPQ09KSk5aSUNINkNHUU5LM1NRMzVXTFpXWkpDUkRBTFJIWjZPVzQ0RFpZVVdNVVYzV1BSSEZSRCIsIm5hbWUiOiJCIiwic3ViIjoiQUNXT01RQTdQWlRLSlNCVFI3QkY2VEJLM0Q3NzY3MzRQV0hXREtPN0hGTVFPTTVCSU9ZUFNZWloiLCJ0eXBlIjoiYWNjb3VudCIsIm5hdHMiOnsiaW1wb3J0cyI6W3sibmFtZSI6InRlc3QiLCJzdWJqZWN0IjoidGVzdCIsImFjY291bnQiOiJBREZCMkpYWVRYT0pFTDZMTkFYRFJFVUdSWDM1Qk9MWkkzQjRQRkZBQzdJUlBSM09BNFFOS0JOMiIsInRvIjoidGVzdCIsInR5cGUiOiJzZXJ2aWNlIn1dLCJsaW1pdHMiOnsic3VicyI6LTEsImNvbm4iOi0xLCJsZWFmIjotMSwiaW1wb3J0cyI6LTEsImV4cG9ydHMiOi0xLCJkYXRhIjotMSwicGF5bG9hZCI6LTEsIndpbGRjYXJkcyI6dHJ1ZX19fQ.AnzziBwt5Tnphc2prONUUOpMpkkAlJHvCPaag0GUtTYPCHKDphcJrwtAHi4v5NOI6npjoes0F0MlrfnHqidDAg"

  # Account "SYS"
  ABKOWIYVTHNEK5HELPWLAT2CF2CUPELIK4SZH2VCJHLFU22B5U2IIZUO: "eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiI1WVUyWkc1UkRTSU1TN1pGVE1MU0NZQUtLVkVFWUpPUlc0TDJPTlY3N1g1TlJZWkFGSkRRIiwiaWF0IjoxNTc0Mzc1OTE2LCJpc3MiOiJPQ09KSk5aSUNINkNHUU5LM1NRMzVXTFpXWkpDUkRBTFJIWjZPVzQ0RFpZVVdNVVYzV1BSSEZSRCIsIm5hbWUiOiJTWVMiLCJzdWIiOiJBQktPV0lZVlRITkVLNUhFTFBXTEFUMkNGMkNVUEVMSUs0U1pIMlZDSkhMRlUyMkI1VTJJSVpVTyIsInR5cGUiOiJhY2NvdW50IiwibmF0cyI6eyJsaW1pdHMiOnsic3VicyI6LTEsImNvbm4iOi0xLCJsZWFmIjotMSwiaW1wb3J0cyI6LTEsImV4cG9ydHMiOi0xLCJkYXRhIjotMSwicGF5bG9hZCI6LTEsIndpbGRjYXJkcyI6dHJ1ZX19fQ.5FrO4sZbWuFgRLuy7c1eQLUq_BQ4PNhIAN5A-sRLkYWmvlc4c_Y4VfTbgl5zhNzCxfvj9SxT7ySgphup2BiRAA
}

```

--------------------------------

### Disable NATS Message Echo (JavaScript)

Source: https://docs.nats.io/using-nats/developer/connecting/noecho

Illustrates disabling echo'd messages in a NATS connection using the `noEcho: true` option in JavaScript. This example also verifies that self-published messages are not received by the subscriber.

```JavaScript
const nc = await connect({
    servers: ["demo.nats.io"],
    noEcho: true,
});

const sub = nc.subscribe(subj, { callback: (_err, _msg) => {} });
nc.publish(subj);
await sub.drain();
// we won't get our own messages
t.is(sub.getProcessed(), 0);
```

--------------------------------

### Key/Value Bucket Management

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

APIs for creating, deleting, and retrieving information about Key/Value store buckets.

```APIDOC
## Key/Value Bucket Management

This section covers the management of Key/Value store buckets, which are independent instances of the Key/Value store built on JetStream persistence.

### Create Key/Value Bucket

**Description:** Creates a new Key/Value bucket with the specified configuration.

**Method:** N/A (API calls vary by language library)

**Endpoint:** N/A

**Parameters:**

*   **Go:** Requires a `KeyValueConfig` struct.
*   **Java:** Requires a `KeyValueConfiguration` object.
*   **JavaScript:** Requires a `name` (string) and optional `opts` object.

**Request Example (Conceptual):
```json
{
  "bucket_name": "my_bucket",
  "config": { ... }
}
```

**Response (Conceptual):
```json
{
  "bucket_info": { ... }
}
```

### Get Bucket Info

**Description:** Retrieves information about an existing Key/Value bucket.

**Method:** N/A (API calls vary by language library)

**Endpoint:** N/A

**Parameters:**

*   **Java:** Requires `bucketName` (string).

**Response Example (Java):
```java
// Returns KeyValueStatus object
```

### Get List of Bucket Names

**Description:** Retrieves a list of all existing Key/Value bucket names.

**Method:** N/A (API calls vary by language library)

**Endpoint:** N/A

**Response Example (Java):
```java
// Returns List<String> of bucket names
```

### Delete Key/Value Bucket

**Description:** Deletes an existing Key/Value bucket. This action is typically administrative.

**Method:** N/A (API calls vary by language library)

**Endpoint:** N/A

**Parameters:**

*   **Go:** Requires `bucket` (string).
*   **Java:** Requires `bucketName` (string).
*   **JavaScript:** Called on a KV object instance.

**Error Handling:** Throws `JetStreamApiException` if the delete fails.
```

--------------------------------

### NATS IO Agent Handshake Log

Source: https://docs.nats.io/using-nats/nex/getting-started/starting-node

This log message indicates a successful handshake between an agent within a virtual machine and the host node process. Successful handshakes are crucial for adding virtual machines to the pool; failures typically result in the node terminating.

```log
INFO[0001] Received agent handshake                      message="Host-supplied metadata" vmid=cmjg61n52omq8dovolmg
```

--------------------------------

### C# NATS Subscribe, Publish, and Drain Example

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Illustrates using the NATS .NET client to create a client, subscribe to a subject with a cancellation token, publish messages, and then cancel the subscription to initiate a drain. Requires the NATS.Net NuGet package.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient();

var subject = client.Connection.NewInbox();

// Make sure to use a cancellation token to end the subscription
using var cts = new CancellationTokenSource();

var sync = false;
var process = Task.Run(async () =>
{
    await foreach (var msg in client.SubscribeAsync<int>(subject, cancellationToken: cts.Token))
    {
        if (msg.Data == -1)
        {
            sync = true;
            continue;
        }
        Console.WriteLine($"Received: {msg.Data}");
        await Task.Delay(TimeSpan.FromMilliseconds(300));
    }

    Console.WriteLine("Subscription completed");
});

// Make sure the subscription is ready
while (sync == false)
{
    await Task.Delay(TimeSpan.FromMilliseconds(100));
    await client.PublishAsync(subject, -1);
}

for (var i = 0; i < 5; i++)
{
    await client.PublishAsync(subject, i);
}
Console.WriteLine("Published 5 messages");

// Cancelling the subscription will unsubscribe from the subject
// and messages that are already in the buffer will be processed
await cts.CancelAsync();
Console.WriteLine("Cancelled subscription");

Console.WriteLine("Waiting for subscription to complete");
await process;

Console.WriteLine("Done");
```

--------------------------------

### Build WebAssembly Module with Cargo

Source: https://docs.nats.io/using-nats/nex/getting-started/building-function

Command to build a Rust project into a WebAssembly module for the `wasm32-wasi` target. This command compiles the code in release mode, optimizing it for size and performance.

```bash
$ cargo build --target wasm32-wasi --release
```

--------------------------------

### NATS Operator Mode Configuration

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Configuration for NATS operator mode, specifying operator JWT, system account, and leaf node remotes with their respective URLs, accounts, and credential paths. This setup allows for distributed connections to NATS servers.

```hocon
operator: ./trustedOperator.jwt
system_account: AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5
leafnodes {
    remotes = [
        {
          url: "nats://localhost:4222"
          account: "ADKGAJU55CHYOIF5H432K2Z2ME3NPSJ5S3VY5Q42Q3OTYOCYRRG7WOWV"
          credentials: "./your-account.creds"
        },
        {
          url: "nats://localhost:4222"
          account: "AAAXAUVSGK7TCRHFIRAS4SYXVJ76EWDMNXZM6ARFGXP7BASNDGLKU7A5"
          credentials: "./system-account.creds"
        },
    ]
}
```

--------------------------------

### C NATS JetStream Subscription and Message Handling

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

This snippet demonstrates setting up NATS JetStream subscriptions (pull, async, sync) and handling messages. It covers initializing JetStream context, configuring stream and consumer options, fetching messages in pull mode, processing messages asynchronously, and receiving messages synchronously. Error handling and message acknowledgement are integrated.

```c
int main(int argc, char **argv)
{
    natsConnection      *conn  = NULL;
    natsStatistics      *stats = NULL;
    natsOptions         *opts  = NULL;
    natsSubscription    *sub   = NULL;
    natsMsg             *msg   = NULL;
    jsCtx               *js    = NULL;
    jsErrCode           jerr   = 0;
    jsOptions           jsOpts;
    jsSubOptions        so;
    natsStatus          s;
    bool                delStream = false;

    opts = parseArgs(argc, argv, usage);

    printf("Created %s subscription on '%s'.\n",
        (pull ? "pull" : (async ? "asynchronous" : "synchronous")), subj);

    s = natsOptions_SetErrorHandler(opts, asyncCb, NULL);

    if (s == NATS_OK)
        s = natsConnection_Connect(&conn, opts);

    if (s == NATS_OK)
        s = jsOptions_Init(&jsOpts);

    if (s == NATS_OK)
        s = jsSubOptions_Init(&so);
    if (s == NATS_OK)
    {
        so.Stream = stream;
        so.Consumer = durable;
        if (flowctrl)
        {
            so.Config.FlowControl = true;
            so.Config.Heartbeat = (int64_t)1E9;
        }
    }

    if (s == NATS_OK)
        s = natsConnection_JetStream(&js, conn, &jsOpts);

    if (s == NATS_OK)
    {
        jsStreamInfo    *si = NULL;

        // First check if the stream already exists.
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_NOT_FOUND)
        {
            jsStreamConfig  cfg;

            // Since we are the one creating this stream, we can delete at the end.
            delStream = true;

            // Initialize the configuration structure.
            jsStreamConfig_Init(&cfg);
            cfg.Name = stream;
            // Set the subject
            cfg.Subjects = (const char*[1]){subj};
            cfg.SubjectsLen = 1;
            // Make it a memory stream.
            cfg.Storage = js_MemoryStorage;
            // Add the stream,
            s = js_AddStream(&si, js, &cfg, NULL, &jerr);
        }
        if (s == NATS_OK)
        {
            printf("Stream %s has %" PRIu64 " messages ( %" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            // Need to destroy the returned stream object.
            jsStreamInfo_Destroy(si);
        }
    }

    if (s == NATS_OK)
    {
        if (pull)
            s = js_PullSubscribe(&sub, js, subj, durable, &jsOpts, &so, &jerr);
        else if (async)
            s = js_Subscribe(&sub, js, subj, onMsg, NULL, &jsOpts, &so, &jerr);
        else
            s = js_SubscribeSync(&sub, js, subj, &jsOpts, &so, &jerr);
    }
    if (s == NATS_OK)
        s = natsSubscription_SetPendingLimits(sub, -1, -1);

    if (s == NATS_OK)
        s = natsStatistics_Create(&stats);

    if ((s == NATS_OK) && pull)
    {
        natsMsgList list;
        int         i;

        for (count = 0; (s == NATS_OK) && (count < total); )
        {
            s = natsSubscription_Fetch(&list, sub, 1024, 5000, &jerr);
            if (s != NATS_OK)
                break;

            if (start == 0)
                start = nats_Now();

            count += (int64_t) list.Count;
            for (i=0; (s == NATS_OK) && (i<list.Count); i++)
                s = natsMsg_Ack(list.Msgs[i], &jsOpts);

            natsMsgList_Destroy(&list);
        }
    }
    else if ((s == NATS_OK) && async)
    {
        while (s == NATS_OK)
        {
            if (count + dropped == total)
                break;

            nats_Sleep(1000);
        }
    }
    else if (s == NATS_OK)
    {
        for (count = 0; (s == NATS_OK) && (count < total); count++)
        {
            s = natsSubscription_NextMsg(&msg, sub, 5000);
            if (s != NATS_OK)
                break;

            if (start == 0)
                start = nats_Now();

            s = natsMsg_Ack(msg, &jsOpts);
            natsMsg_Destroy(msg);
        }
    }

    if (s == NATS_OK)
    {
        printStats(STATS_IN|STATS_COUNT, conn, sub, stats);
        printPerf("Received");
    }
    if (s == NATS_OK)
    {
        jsStreamInfo *si = NULL;

        // Let's report some stats after the run
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_OK)
        {
            printf("\nStream %s has %" PRIu64 " messages ( %" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            jsStreamInfo_Destroy(si);
        }
        if (delStream)
        {
            printf("\nDeleting stream %s: ", stream);
            s = js_DeleteStream(js, stream, NULL, &jerr);
            if (s == NATS_OK)
                printf("OK!");
            printf("\n");
        }
    }
    else
    {
        printf("Error: %u - %s - jerr=%u\n", s, natsStatus_GetText(s), jerr);
        nats_PrintLastErrorStack(stderr);
    }

    // Destroy all our objects to avoid report of memory leak
    jsCtx_Destroy(js);
    return 0;
}

```

--------------------------------

### NATS Authorization Configuration with Explicit Allow/Deny

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/authorization

An example of NATS authorization configuration without variables, demonstrating the use of explicit 'allow' and 'deny' options for publish and subscribe permissions. This configuration sets full permissions for an 'admin' user and restricted permissions for a 'test' user, denying all publishes and allowing subscriptions to 'client. திராக'.

```hcl
authorization: {
    users = [
        {
            user: admin
            password: secret
            permissions: {
                publish: ">"
                subscribe: ">"
            }
        }
        {
            user: test
            password: test
            permissions: {
                publish: {
                    deny: ">"
                },
                subscribe: {
                    allow: "client. திராக"
                }
            }
        }
    ]
}
```

--------------------------------

### C: Put, PutString, and Create Data in NATS KV Store

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Offers C functions for interacting with the NATS KV store, including putting byte arrays ('kvStore_Put'), strings ('kvStore_PutString'), and creating new entries ('kvStore_Create').

```C
NATS_EXTERN natsStatus 	kvStore_Put (uint64_t *rev, kvStore *kv, const char *key, const void *data, int len) 
	 Places the new value for the key into the store.
 
NATS_EXTERN natsStatus 	kvStore_PutString (uint64_t *rev, kvStore *kv, const char *key, const char *data) 
	 Places the new value (as a string) for the key into the store.
 
NATS_EXTERN natsStatus 	kvStore_Create (uint64_t *rev, kvStore *kv, const char *key, const void *data, int len)
```

--------------------------------

### Test NATS Configuration with Publish

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This shell command tests the NATS server configuration by publishing a message to a subject using provided credentials. It verifies that the server is running and accessible with the specified configuration.

```shell
nats pub --creds ~/.nkeys/creds/memory/A/TA.creds hello world
```

--------------------------------

### Send Request from Blue Leaf Node (NATS CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-leafnodes-docker

Sends a request to the 'docker-leaf-test' subject using the 'blue' leaf node's connection (localhost:4333). It demonstrates sending a message and receiving a response from the service started on the 'red' leaf node.

```shell
$ nats -s localhost:4333 request docker-leaf-test "Hello World"

At 8:15PM, I received your request: Hello World

```

--------------------------------

### Requesting Next Message with Pull-Based Consumer (API)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/consumers

This example shows how to programmatically request the next message from a consumer using the NATS CLI's request feature, targeting the JetStream API. It illustrates publishing a request and receiving a message.

```shell
nats req '$JS.API.CONSUMER.MSG.NEXT.ORDERS.DISPATCH' ''
```

--------------------------------

### NATS Message Structure Definition

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol/nats-client-dev

Defines the 'Msg' structure commonly used in NATS clients to represent messages. This structure includes fields for the subject, reply-to subject, message data, and a reference to the subscription, essential for message handling.

```go
type Msg struct {
    Subject string
    Reply   string
    Data    []byte
    Sub     *Subscription
}
```

--------------------------------

### Create KV Bucket (Shell)

Source: https://docs.nats.io/nats-concepts/jetstream/key-value-store/kv_walkthrough

Creates a new Key/Value (KV) bucket, which is analogous to a stream in JetStream. This bucket will store key-value pairs. The command specifies the name of the bucket to be created.

```shell
nats kv add my-kv
```

--------------------------------

### C NATS Subscribe, Publish, and Drain Example

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Demonstrates subscribing to a NATS subject, publishing messages, and then draining the subscription using the C NATS client library. The drain operation ensures all pending messages are processed before unsubscribing.

```c
natsConnection      *conn      = NULL;
natsSubscription    *sub       = NULL;
natsStatus          s          = NATS_OK;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);

// Subscribe
if (s == NATS_OK)
    s = natsConnection_Subscribe(&sub, conn, "foo", onMsg, NULL);

// Publish 2 messages
if (s == NATS_OK)
{
    int i;
    for (i=0; (s == NATS_OK) && (i<2); i++)
    {
        s = natsConnection_PublishString(conn, "foo", "hello");
    }
}

// Call Drain on the subscription. It unsubscribes but
// wait for all pending messages to be processed.
if (s == NATS_OK)
    s = natsSubscription_Drain(sub);

(...)

// Destroy objects that were created
natsSubscription_Destroy(sub);
natsConnection_Destroy(conn);
```

--------------------------------

### Get CA Certificate Root Path with mkcert

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/tls

Retrieves the root directory path for the Certificate Authority (CA) certificates generated by mkcert. This is useful for locating the CA certificate to be used in NATS configurations or for copying to other directories.

```bash
mkcert -CAROOT
```

--------------------------------

### Configure NATS Client Authentication with NKey (Node.js)

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/nkey_auth

This JavaScript code demonstrates how to configure a NATS client to use NKey authentication. It requires the `nats` and `ts-nkeys` libraries. The client provides a `sigCB` (signature callback) function that securely loads the private key seed and signs connection challenges using it, proving control of the associated public key.

```javascript
const NATS = require('nats');
const nkeys = require('ts-nkeys');

const nkey_seed = ‘SUACSSL3UAHUDXKFSNVUZRF5UHPMWZ6BFDTJ7M6USDXIEDNPPQYYYCU3VY’;
const nc = NATS.connect({
  port: PORT,
  nkey: 'UDXU4RCSJNZOIQHZNWXHXORDPRTGNJAHAHFRGZNEEJCPQTT2M7NLCNF4',
  sigCB: function (nonce) {
    // client loads seed safely from a file
    // or some constant like `nkey_seed` defined in
    // the program
    const sk = nkeys.fromSeed(Buffer.from(nkey_seed));
    return sk.sign(nonce);
   }
});
...
```

--------------------------------

### Get NATS Connection Status (Go)

Source: https://docs.nats.io/using-nats/developer/connecting/events

This Go code snippet shows how to connect to a NATS server and retrieve the current connection status. It defines a helper function to translate the status code into a human-readable string. It also demonstrates closing the connection and checking the status afterward. Requires the 'nats' package.

```go
nc, err := nats.Connect("demo.nats.io", nats.Name("API Example"))
if err != nil {
    log.Fatal(err)
}
def newConn(nc *nats.Conn) string {
    switch nc.Status() {
    case nats.CONNECTED:
        return "Connected"
    case nats.CLOSED:
        return "Closed"
    default:
        return "Other"
    }
}
log.Printf("The connection is %v\n", newConn(nc))

nc.Close()

log.Printf("The connection is %v\n", newConn(nc))
```

--------------------------------

### Mimic Connect Timeout - Ruby NATS Client

Source: https://docs.nats.io/using-nats/developer/connecting/connect_timeout

Since the Ruby NATS client lacks a direct connect timeout API, this example demonstrates mimicking it using an `EM.add_timer`. A 10-second timer is set; if the connection is not established within this time, the timer can be canceled.

```Ruby
# There is currently no connect timeout as part of the Ruby NATS client API, but you can use a timer to mimic it.
require 'nats/client'

timer = EM.add_timer(10) do
  NATS.connect do |nc|
    # Do something with the connection

    # Close the connection
    nc.close
  end
end
EM.cancel_timer(timer)
```

--------------------------------

### Python Queue Subscription

Source: https://docs.nats.io/using-nats/developer/receiving/queues

Subscribes to the 'updates' subject with the queue group 'workers'. It uses an asyncio Future to await a message, publishes a test message, and then prints the received message. Includes connection setup and asynchronous callback handling.

```Python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

future = asyncio.Future()

async def cb(msg):
  nonlocal future
  future.set_result(msg)

await nc.subscribe("updates", queue="workers", cb=cb)
await nc.publish("updates", b'All is Well')

msg = await asyncio.wait_for(future, 1)
print("Msg", msg)
```

--------------------------------

### Describe User Permissions with nsc

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

This command displays the detailed authorization settings for a specific user, including their publish and subscribe allowances. This is useful for verifying configured permissions.

```shell
nsc describe user s
```

--------------------------------

### Unsubscribe from NATS (C#)

Source: https://docs.nats.io/using-nats/developer/receiving/unsubscribing

Demonstrates unsubscribing from NATS in C# using the NATS.Net library. Subscriptions can be cancelled using a CancellationTokenSource, or implicitly when exiting a loop processing messages. The example shows unsubscribing when a specific message 'exit' is received.

```csharp
// dotnet add package NATS.Net
using NATS.Net

await using var client = new NatsClient()

// Cancel the subscription after 10 seconds
using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(10))

// Subscribe to the "updates" subject
// We unsubscribe when we receive the message "exit"
// or when the cancellation token is triggered.
await foreach (var msg in client.SubscribeAsync<string>("updates").WithCancellation(cts.Token))
{
    Console.WriteLine($"Received: {msg.Data}")
    
    if (msg.Data == "exit")
    {
        // When we exit the loop, we unsubscribe from the subject
        break
    }
}

Console.WriteLine("Unsubscribed from updates")

```

--------------------------------

### Fetch Next Message (No Wait) - NATS CLI Example

Source: https://docs.nats.io/reference/reference-protocols/nats_api_reference

Demonstrates fetching messages from a pull-based consumer using the 'no_wait' option via the NATS CLI. The request includes a JSON payload specifying the batch size and the 'no_wait' flag. The output shows a '404 Status' indicating no messages available.

```shell
nats req '$JS.API.CONSUMER.MSG.NEXT.ORDERS.NEW' '{"no_wait": true, "batch": 10}'
```

--------------------------------

### Create NATS CLI Context

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

Creates a NATS CLI context named 'myuser' using specified credentials. This context simplifies future NATS CLI commands by avoiding the need to repeatedly provide connection and authentication arguments.

```shell
nats context add myuser --creds ~/.nkeys/creds/MyOperator/MyAccount/MyUser.creds
```

--------------------------------

### Publish Time Data to NATS Subjects (C# .NET)

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

Shows how to publish current time data to NATS subjects using C# and the NATS.Net library. It includes setting up the client, getting the current time in different time zones, and publishing.

```csharp
// dotnet add package NATS.Net
// dotnet add package NodaTime
using NATS.Net;
using NodaTime;

await using var client = new NatsClient();

Instant now = SystemClock.Instance.GetCurrentInstant();

{
    DateTimeZone zone = DateTimeZoneProviders.Tzdb["America/New_York"];
    string formatted = now.InZone(zone).ToString();
    await client.PublishAsync("time.us.east", formatted);
    await client.PublishAsync("time.us.east.atlanta", formatted);
}

{
    DateTimeZone zone = DateTimeZoneProviders.Tzdb["Europe/Warsaw"];
    string formatted = now.InZone(zone).ToString();
    await client.PublishAsync("time.eu.east", formatted);
    await client.PublishAsync("time.eu.east.warsaw", formatted);
}
```

--------------------------------

### Send JSON Data with NATS in Java

Source: https://docs.nats.io/using-nats/developer/sending/structure

This Java example demonstrates publishing a custom object serialized to JSON via NATS. It utilizes Gson for JSON conversion and requires the NATS Java client and Gson library.

```java
class StockForJsonPub {
    public String symbol;
    public float price;
}

public class PublishJSON {
    public static void main(String[] args) {
        try {
            Connection nc = Nats.connect("nats://demo.nats.io:4222");

            // Create the data object
            StockForJsonPub stk = new StockForJsonPub();
            stk.symbol="GOOG";
            stk.price=1200;

            // use Gson to encode the object to JSON
            GsonBuilder builder = new GsonBuilder();
            Gson gson = builder.create();
            String json = gson.toJson(stk);

            // Publish the message
            nc.publish("updates", json.getBytes(StandardCharsets.UTF_8));

            // Make sure the message goes through before we close
            nc.flush(Duration.ZERO);
            nc.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

--------------------------------

### NATS Leaf Node Information (/leafz) Example

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

This JSON response details the NATS server's leaf node connections, including server ID, current time, and a list of leaf nodes. Each leaf node entry contains its account, IP address, port, round-trip time (RTT), message and byte counts, and subscription information.

```json
{
  "server_id": "NC2FJCRMPBE5RI5OSRN7TKUCWQONCKNXHKJXCJIDVSAZ6727M7MQFVT3",
  "now": "2019-08-27T09:07:05.841132-06:00",
  "leafnodes": 1,
  "leafs": [
    {
      "account": "$G",
      "ip": "127.0.0.1",
      "port": 6223,
      "rtt": "200µs",
      "in_msgs": 0,
      "out_msgs": 10000,
      "in_bytes": 0,
      "out_bytes": 1280000,
      "subscriptions": 1,
      "subscriptions_list": ["foo"]
    }
  ]
}
```

--------------------------------

### Create NATS Overlay Network and Seed Server (Bash)

Source: https://docs.nats.io/running-a-nats-service/nats_docker/docker_swarm

This snippet creates an overlay network named 'nats-cluster-example' for the Docker Swarm cluster. It then instantiates the first NATS server as a service named 'nats-cluster-node-1', configuring it to listen for cluster joins on port 6222 and enabling debugging with '-DV'.

```bash
docker network create --driver overlay nats-cluster-example
docker service create --network nats-cluster-example --name nats-cluster-node-1 nats:1.0.0 -cluster nats://0.0.0.0:6222 -DV
```

--------------------------------

### Run NATS Server in a Cluster with Docker

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This command initiates the first NATS server for a cluster. It connects to the 'nats' Docker network, exposes client, management, and cluster ports, and configures the cluster name and routing. Subsequent servers can be added using similar commands, pointing to this initial server.

```bash
docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222 --cluster_name NATS --cluster nats://0.0.0.0:6222
```

--------------------------------

### NATS JetStream Server Configuration (Nats)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering/jetstream_clustering

Example NATS server configuration for a JetStream enabled node in a cluster. It includes server name, listen port, system account user credentials (with bcrypt password hash), JetStream storage directory, and cluster configuration with routing information.

```Nats
server_name=n1-c1
listen=4222

accounts {
  $SYS {
    users = [
      {
        user: "admin",
        pass: "$2a$11$DRh4C0KNbNnD8K/hb/buWe1zPxEHrLEiDmuq1Mi0rRJiH/W25Qidm"
      }
    ]
  }
}

jetstream {
   store_dir=/nats/storage
}

cluster {
  name: C1
  listen: 0.0.0.0:6222
  routes: [
    nats://host_b:6222
    nats://host_c:6222
  ]
}
```

--------------------------------

### Get NATS JetStream Server Report (Bash)

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes/jetstream_leafnodes

Retrieves a summary report of the JetStream server status, including stream and consumer counts, message volume, and resource usage. This command requires authentication to a server with system account privileges.

```bash
nats  --server nats://admin:admin@localhost:4222 server report jetstream
```

--------------------------------

### Publish to NATS JetStream Stream (Go)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/publish

This Go code snippet demonstrates how to connect to NATS, create a JetStream context, add a stream, and publish messages (both synchronous and asynchronous) to that stream. It includes error handling for connection and JetStream context creation.

```go
func ExampleJetStream() {
	nc, err := nats.Connect("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "example-stream",
		Subjects: []string{"example-subject"},
	})

	js.Publish("example-subject", []byte("Hello JS!"))

	// Publish messages asynchronously.
	for i := 0; i < 500; i++ {
		js.PublishAsync("example-subject", []byte("Hello JS Async!"))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
}
```

--------------------------------

### Connect to NATS using C

Source: https://docs.nats.io/using-nats/developer/connecting/userpass

Provides a C code snippet for connecting to a NATS server. It demonstrates the creation of NATS options, setting the URL with authentication, establishing the connection, and the subsequent cleanup of created objects.

```c
natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_SetURL(opts, "nats://myname:password@127.0.0.1:4222");
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### NATS Configuration: Artificial Message Loss

Source: https://docs.nats.io/running-a-nats-service/configuration/configuring_subject_mapping

This configuration introduces artificial message loss for chaos testing. By mapping a percentage of traffic to the same subject, that portion of messages will be dropped by the server. This example drops 50% of traffic published to `foo.loss.a`.

```NATS Configuration
foo.loss.>: [ { destination: foo.loss.>, weight: 50% } ]
```

--------------------------------

### Get NATS Stream Information

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/streams

Retrieves detailed information about a specific NATS stream, including message counts, byte size, sequence numbers, and active consumers. This command is essential for monitoring stream status and health.

```shell
nats str info ORDERS
```

--------------------------------

### C NATS Client Error Handling Setup

Source: https://docs.nats.io/using-nats/developer/connecting/events/slow

This C code snippet shows how to define and register a custom error callback function for the NATS connection. It includes initializing NATS options, setting the error handler, connecting to the NATS server, and ensuring proper cleanup of connection and options objects.

```c
static void
errorCB(natsConnection *conn, natsSubscription *sub, natsStatus s, void *closure)
{

    // Do something
    printf("Error: %d - %s", s, natsStatus_GetText(s));
}

(...)

natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_SetErrorHandler(opts, errorCB, NULL);
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### Connect to NATS Demo Instance using Telnet

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol-demo

Initiates a connection to the NATS demo instance at demo.nats.io on port 4222 using Telnet. This step is crucial for interacting with the NATS server via the text-based protocol.

```text
telnet demo.nats.io 4222
```

--------------------------------

### Persist JetStream Data with Docker Volume

Source: https://docs.nats.io/running-a-nats-service/nats_docker/jetstream_docker

This command starts a NATS server with JetStream enabled and persists its data to a Docker volume named 'nats'. The `-sd /data` flag specifies the directory for data storage within the container.

```shell
docker run -p 4222:4222 -v nats:/data nats -js -sd /data
```

--------------------------------

### Monitor NATS Connection Status (Ruby)

Source: https://docs.nats.io/using-nats/developer/connecting/events

This Ruby code snippet demonstrates how to start a NATS connection with reconnection attempts and monitor its status. It uses EventMachine for periodic checks and prints whether the connection is active, reconnecting, or closing. Requires the 'nats' gem and EventMachine.

```ruby
NATS.start(max_reconnect_attempts: 2) do |nc|
  puts "Connect is connected?: #{nc.connected?}"

  timer = EM.add_periodic_timer(1) do
    if nc.closing?
      puts "Connection closed..."
      EM.cancel_timer(timer)
      NATS.stop
    end

    if nc.reconnecting?
      puts "Reconnecting to NATS..."
      next
    end
  end
end
```

--------------------------------

### Connect to Default NATS Server in C

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Demonstrates connecting to the default NATS server URL (NATS_DEFAULT_URL) using the C NATS client library. It includes error handling for the connection attempt and proper destruction of the connection object.

```c
natsConnection      *conn = NULL;
natsStatus          s;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);
if (s != NATS_OK)
  // handle error

// Destroy connection, no-op if conn is NULL.
natsConnection_Destroy(conn);
```

--------------------------------

### Generating NKey for Account Issuer using nsc

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_callout

Illustrates how to use the nsc (NATS Security Configuration) tool to generate an NKey pair for the account issuer. This NKey is used for signing authorization payloads in the auth callout feature. The output shows example public and private key pairs.

```bash
$ nsc generate nkey --account
SAANDLKMXL6CUS3CP52WIXBEDN6YJ545GDKC65U5JZPPV6WH6ESWUA6YAI
ABJHLOVMPA4CI6R5KLNGOB4GSLNIY7IOUPAJC4YFNDLQVIOBYQGUWVLA
```

--------------------------------

### Java Object Store Interface for Key-Value Operations

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

The Java ObjectStore interface defines methods for managing key-value objects. It supports operations like putting, getting, updating metadata, deleting objects, and listing all objects in a bucket. It also includes functionality for creating watches and sealing the bucket.

```java
/**
 * Object Store Management context for creation and access to key value buckets.
 */
public interface ObjectStore {

    /**
     * Get the name of the object store's bucket.
     * @return the name
     */
    String getBucketName();

    /**
     * Place the contents of the input stream into a new object.
     */
    ObjectInfo put(ObjectMeta meta, InputStream inputStream) throws IOException, JetStreamApiException, NoSuchAlgorithmException;

    /**
     * Place the contents of the input stream into a new object.
     */
    ObjectInfo put(String objectName, InputStream inputStream) throws IOException, JetStreamApiException, NoSuchAlgorithmException;

    /**
     * Place the bytes into a new object.
     */
    ObjectInfo put(String objectName, byte[] input) throws IOException, JetStreamApiException, NoSuchAlgorithmException;

    /**
     * Place the contents of the file into a new object using the file name as the object name.
     */
    ObjectInfo put(File file) throws IOException, JetStreamApiException, NoSuchAlgorithmException;

    /**
     * Get an object by name from the store, reading it into the output stream, if the object exists.
     */
    ObjectInfo get(String objectName, OutputStream outputStream) throws IOException, JetStreamApiException, InterruptedException, NoSuchAlgorithmException;

    /**
     * Get the info for an object if the object exists / is not deleted.
     */
    ObjectInfo getInfo(String objectName) throws IOException, JetStreamApiException;

    /**
     * Get the info for an object if the object exists, optionally including deleted.
     */
    ObjectInfo getInfo(String objectName, boolean includingDeleted) throws IOException, JetStreamApiException;

    /**
     * Update the metadata of name, description or headers. All other changes are ignored.
     */
    ObjectInfo updateMeta(String objectName, ObjectMeta meta) throws IOException, JetStreamApiException;

    /**
     * Delete the object by name. A No-op if the object is already deleted.
     */
    ObjectInfo delete(String objectName) throws IOException, JetStreamApiException;

    /**
     * Add a link to another object. A link cannot be for another link.
     */
    ObjectInfo addLink(String objectName, ObjectInfo toInfo) throws IOException, JetStreamApiException;

    /**
     * Add a link to another object store (bucket).
     */
    ObjectInfo addBucketLink(String objectName, ObjectStore toStore) throws IOException, JetStreamApiException;

    /**
     * Close (seal) the bucket to changes. The store (bucket) will be read only.
     */
    ObjectStoreStatus seal() throws IOException, JetStreamApiException;

    /**
     * Get a list of all object [infos] in the store.
     */
    List<ObjectInfo> getList() throws IOException, JetStreamApiException, InterruptedException;

    /**
     * Create a watch on the store (bucket).
     */
    NatsObjectStoreWatchSubscription watch(ObjectStoreWatcher watcher, ObjectStoreWatchOption... watchOptions) throws IOException, JetStreamApiException, InterruptedException;

    /**
     * Get the ObjectStoreStatus object.
     */
    ObjectStoreStatus getStatus() throws IOException, JetStreamApiException;

```

--------------------------------

### Generate Data using NATS CLI

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/replication

These commands generate 100 new ORDER and RETURN messages using the NATS CLI. They are essential for populating streams and testing replication.

```shell
nats req ORDERS.new "ORDER {{Count}}" --count 100
nats req RETURNS.new "RETURN {{Count}}" --count 100
```

--------------------------------

### Get NATS Connection Status (Java)

Source: https://docs.nats.io/using-nats/developer/connecting/events

This Java code snippet demonstrates how to establish a connection to a NATS server and print its current status. It then closes the connection and prints the status again to show the change. Requires the NATS Java client library.

```java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

System.out.println("The Connection is: " + nc.getStatus()); // CONNECTED

nc.close();

System.out.println("The Connection is: " + nc.getStatus()); // CLOSED
```

--------------------------------

### C JetStream Publish Error Handling and Stream Creation

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

Demonstrates setting up asynchronous publish error handlers and creating streams in the NATS JetStream C client.

```APIDOC
## C JetStream API

### Asynchronous Publish Error Handling

- **Description**: Configures a callback function to handle errors that occur during asynchronous message publishing.
- **Method**: N/A (Code example provided)
- **Endpoint**: N/A (NATS JetStream C client options)

#### `jsOpts.PublishAsync.ErrHandler`
- **Description**: A function pointer to the error handler callback.
- **Callback Signature**: `_jsPubErr(jsCtx *js, jsPubAckErr *pae, void *closure)`
  - `js`: The JetStream context.
  - `pae`: Structure containing error details (error code, text, original message).
  - `closure`: User-defined data passed to the handler.

#### `jsOpts.PublishAsync.ErrHandlerClosure`
- **Description**: A pointer to user data that will be passed to the error handler callback.

### Stream Creation

- **Description**: Creates a stream if it does not already exist. Includes logic to check for existing streams and defines configuration options.
- **Method**: N/A (Code example provided)
- **Endpoint**: N/A (NATS JetStream C client functions)

#### `js_GetStreamInfo(jsStreamInfo **si, jsCtx *js, const char *stream, NATS_UNUSED(jsStream *) s, jsErrCode *jerr)`
- **Description**: Attempts to retrieve information about a stream. Returns `NATS_OK` if the stream exists, `NATS_NOT_FOUND` otherwise.

#### `js_CreateStream(jsStreamInfo **si, jsCtx *js, jsStreamConfig *cfg, jsErrCode *jerr)` (Implied)
- **Description**: Creates a new stream based on the provided configuration (`jsStreamConfig`). This function is called when `js_GetStreamInfo` returns `NATS_NOT_FOUND`.
```

--------------------------------

### Verify NATS Cluster Routes via HTTP

Source: https://docs.nats.io/running-a-nats-service/nats_docker

This command sends an HTTP GET request to the management port of a NATS server to retrieve information about active routes in the cluster. The response includes the number of routes and details about each connected server, useful for verifying cluster health and connectivity.

```bash
curl http://127.0.0.1:8222/routez
```

--------------------------------

### Upload a File with a Specific Name

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This demonstrates uploading a file while explicitly defining its name (key) within the object store bucket using the `--name` flag. This allows for custom key naming, independent of the file's original path.

```shell
nats object put --name /Movies/NATS-logo.mov myobjbucket ~/Movies/NATS-logo.mov
```

--------------------------------

### Configure NATS Server Token Authentication

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/tokens

This configuration snippet shows how to set up token authentication directly within the NATS server configuration file. It uses the `authorization` block with a `token` property. This is a basic setup for allowing clients to connect using a shared secret.

```nats-config
authorization {
    token: "s3cr3t"
}
```

--------------------------------

### Connect to NATS Cluster in JavaScript (Node.js)

Source: https://docs.nats.io/using-nats/developer/connecting/cluster

Shows how to establish a connection to a NATS cluster using the JavaScript client. It takes an array of server URLs and uses `async/await` for connection and closure. A placeholder function `doSomething()` is included to represent operations on the connection.

```JavaScript
const nc = await connect({
    servers: [
      "nats://demo.nats.io:4222",
      "nats://localhost:4222",
    ],
});
// Do something with the connection
doSomething();
// When done close it
await nc.close();
```

--------------------------------

### Get Maximum Payload Size in NATS Clients

Source: https://docs.nats.io/using-nats/developer/connecting/misc

Clients can obtain the configured maximum payload size after connection. This allows applications to chunk or limit data as needed to pass through the server. This functionality is demonstrated across various programming languages.

```Go
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

mp := nc.MaxPayload()
log.Printf("Maximum payload is %v bytes", mp)

// Do something with the max payload
```

```Java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

long mp = nc.getMaxPayload();
System.out.println("max payload for the server is " + mp + " bytes");
```

```JavaScript
t.log(`max payload for the server is ${nc.info.max_payload} bytes`);
```

```Python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

print("Maximum payload is %d bytes" % nc.max_payload)

# Do something with the max payload.
```

```C#
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient("nats://demo.nats.io:4222");

// Make sure we connect to a server to receive the server info,
// since connecting to servers is lazy in .NET client.
await client.ConnectAsync();

Console.WriteLine($"MaxPayload = {client.Connection.ServerInfo.MaxPayload}");
```

```Ruby
require 'nats/client'

NATS.start(max_outstanding_pings: 5) do |nc|
  nc.on_reconnect do
    puts "Got reconnected to #{nc.connected_server}"
  end

  nc.on_disconnect do |reason|
    puts "Got disconnected! #{reason}"
  end

  # Do something with the max_payload
  puts "Maximum Payload is #{nc.server_info[:max_payload]} bytes"
end
```

```C
natsConnection      *conn    = NULL;
natsStatus          s        = NATS_OK;

s = natsConnection_ConnectTo(&conn, NATS_DEFAULT_URL);
if (s == NATS_OK)
{
    int64_t mp = natsConnection_GetMaxPayload(conn);
    printf("Max payload: %d\n", (int) mp);
}

(...)

// Destroy objects that were created
natsConnection_Destroy(conn);
```

--------------------------------

### Create NATS JetStream Stream (Config File)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/streams

This snippet demonstrates creating a NATS JetStream Stream named 'ORDERS' using a configuration file. The `--config` flag points to a JSON file containing the stream's configuration, which follows the same format as the output of `nats str info ORDERS -j | jq .config`. This method is useful for managing complex configurations or version controlling stream settings.

```shell
nats str add ORDERS --config orders.json
```

--------------------------------

### Signal Specific NATS Server by PID

Source: https://docs.nats.io/running-a-nats-service/nats_admin/signals

Sends a signal to a specific NATS server process identified by its Process ID (PID). This is necessary when multiple NATS servers are running or `pgrep` is unavailable. Example: `nats-server --signal stop=<pid>`.

```shell
nats-server --signal stop=<pid>
```

--------------------------------

### NATS CLI Publish Command - Payload Violation Error

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This example shows an error scenario when publishing a message that exceeds the payload limit defined in the user's JWT. The NATS CLI returns an error indicating a 'Maximum Payload Violation', preventing the message from being sent.

```bash
> nats -s localhost:4222 "--creds=user.creds" pub "foo" "hello world"
nats: error: nats: Maximum Payload Violation, try --help
>
```

--------------------------------

### Disable NATS Message Echo (Python)

Source: https://docs.nats.io/using-nats/developer/connecting/noecho

Provides a Python example for disabling echo'd messages on a NATS connection using `no_echo=True`. It demonstrates two connections, one with echo off, to show that messages published by `ncA` are not received by `ncA` but are by `ncB`.

```Python
ncA = NATS()
ncB = NATS()

await ncA.connect(no_echo=True)
await ncB.connect()

async def handler(msg):
   # Messages sent by `ncA' will not be received.
   print("[Received] ", msg)

await ncA.subscribe("greetings", cb=handler)
await ncA.flush()
await ncA.publish("greetings", b'Hello World!')
await ncB.publish("greetings", b'Hello World!')

# Do something with the connection

await asyncio.sleep(1)
await ncA.drain()
await ncB.drain()
```

--------------------------------

### View Nex Node Information and Workloads

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-functions

This command retrieves detailed information about a specific Nex node, identified by its node ID. It displays node details such as version, uptime, tags, memory usage, and a list of all deployed workloads, including their IDs, health status, runtime, names, and descriptions. This is used to verify the deployment of both JavaScript and WebAssembly functions.

```bash
$ nex node info NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
NEX Node Information

         Node: NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
         Xkey: XDKZMOZKVBXSY3YXPIXEFKGPML75PLD7APFHZ474EOCILZDQGPZSXJNZ
      Version: 0.0.1
       Uptime: 7m31s
         Tags: nex.arch=amd64, nex.cpucount=8, nex.os=linux, simple=true

Memory in kB:

           Free: 32,280,180
      Available: 56,018,344
          Total: 63,883,232

Workloads:

             Id: cmjud7n52omhlsa377cg
        Healthy: true
        Runtime: 7m31s
           Name: echofunctionjs
    Description: Workload published in devmode
  
             Id: cmjudmn52omhlsa377d0
        Healthy: true
        Runtime: 6m31s
           Name: echofunctionwasm
    Description: Workload published in devmode
```

--------------------------------

### Send String Message to Subject - Go, Java, JavaScript, Python, C#, Ruby

Source: https://docs.nats.io/using-nats/developer/sending

Demonstrates sending a string message to a NATS subject ('updates') across multiple programming languages. These examples showcase the typical NATS client usage for publishing data, often involving encoding strings to byte arrays implicitly or explicitly. Dependencies include the respective NATS client libraries for each language. Input is a string and a subject; output is a published message.

```go
nc, err := nats.Connect("demo.nats.io", nats.Name("API PublishBytes Example"))
if err != nil {
    log.Fatal(err)
}
deferr nc.Close()

if err := nc.Publish("updates", []byte("All is Well")); err != nil {
    log.Fatal(err)
}
```

```java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

nc.publish("updates", "All is Well".getBytes(StandardCharsets.UTF_8));
```

```javascript
const sc = StringCodec();
nc.publish("updates", sc.encode("All is Well"));
```

```python
nc = NATS()

await nc.connect(servers=["nats://demo.nats.io:4222"])

await nc.publish("updates", b'All is Well')
```

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient(url: "demo.nats.io", name: "API Publish String Example");

// The default serializer uses UTF-8 encoding for strings
await client.PublishAsync<string>(subject: "updates", data: "All is Well");
```

```ruby
require 'nats/client'

NATS.start(servers:["nats://127.0.0.1:4222"]) do |nc|
  nc.publish("updates", "All is Well")
end
```

--------------------------------

### Configure NATS JetStream Encryption using Environment Variable

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/encryption_at_rest

This configuration demonstrates enabling encryption at rest for NATS JetStream by referencing an environment variable for the encryption key. This is a more secure approach as the key is not stored directly in the configuration file. The environment variable can be set before starting the NATS server.

```yaml
jetstream : {
  cipher: chachapoly
  key: $JS_KEY
}
```

--------------------------------

### Receive JSON Data in JavaScript

Source: https://docs.nats.io/using-nats/developer/receiving/structure

This JavaScript snippet demonstrates how to subscribe to a NATS subject and process incoming messages. It utilizes the `msg.json()` method to parse the message data as JSON and logs it to the console. The example shows a subscription with a callback and a maximum of one message.

```javascript
const sub = nc.subscribe(subj, {
  callback: (_err, msg) => {
    t.log(`${msg.json()}`);
  },
  max: 1,
});
```

--------------------------------

### Configure Leaf Node with TLS Remote

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This configuration is similar to the standard leaf node setup but specifies a TLS connection to the remote NATS system. This is important for securing the communication between the leaf node and the main cluster.

```hcl
listen: "127.0.0.1:4111"
leafnodes {
    remotes = [
        {
          url: "tls://s3cr3t@localhost"
        },
    ]
}
```

--------------------------------

### Receive JSON Data in Java

Source: https://docs.nats.io/using-nats/developer/receiving/structure

This Java example shows how to connect to a NATS server, subscribe to the 'updates' subject, and process incoming JSON messages. It uses Gson for JSON parsing and a CountDownLatch to wait for a specified number of messages. The received data is printed to the console.

```java
class StockForJsonSub {
    public String symbol;
    public float price;

    public String toString() {
        return symbol + " is at " + price;
    }
}

public class SubscribeJSON {
    public static void main(String[] args) {

        try {
            Connection nc = Nats.connect("nats://demo.nats.io:4222");

            // Use a latch to wait for 10 messages to arrive
            CountDownLatch latch = new CountDownLatch(10);

            // Create a dispatcher and inline message handler
            Dispatcher d = nc.createDispatcher((msg) -> {
                Gson gson = new Gson();

                String json = new String(msg.getData(), StandardCharsets.UTF_8);
                StockForJsonSub stk = gson.fromJson(json, StockForJsonSub.class);

                // Use the object
                System.out.println(stk);

                latch.countDown();
            });

            // Subscribe
            d.subscribe("updates");

            // Wait for a message to come in
            latch.await(); 

            // Close the connection
            nc.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

--------------------------------

### Set Up Queued Subscribers in Ruby

Source: https://docs.nats.io/using-nats/developer/receiving/drain

This Ruby code snippet sets up multiple subscribers for different subjects ('foo', 'bar', 'quux') using a queue group ('workers'). It demonstrates how NATS distributes messages across subscribers within the same queue group. The example also includes logic to gracefully drain the connection after a delay, ensuring all messages are processed before closing.

```ruby
NATS.start(drain_timeout: 1) do |nc|
  NATS.subscribe('foo', queue: "workers") do |msg, reply, sub|
    nc.publish(reply, "ACK:#{msg}")
  end

  NATS.subscribe('bar', queue: "workers") do |msg, reply, sub|
    nc.publish(reply, "ACK:#{msg}")
  end

  NATS.subscribe('quux', queue: "workers") do |msg, reply, sub|
    nc.publish(reply, "ACK:#{msg}")
  end

  EM.add_timer(2) do
    next if NATS.draining?

    # Drain gracefully closes the connection.
    NATS.drain do
      puts "Done draining. Connection is closed."
    end
  end
end
```

--------------------------------

### Deploy WebAssembly Function with Nex CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-functions

This command deploys a WebAssembly function named 'echofunction.wasm' to the Nex network, triggering it on the 'wasm.echo' subject. It requires the path to the WASM file and the trigger subject. The output shows the workload ID assigned by Nex.

```bash
$ nex devrun ../examples/wasm/echofunction/echofunction.wasm --trigger_subject=wasm.echo
Reusing existing issuer account key: /home/kevin/.nex/issuer.nk
Reusing existing publisher xkey: /home/kevin/.nex/publisher.xk
🚀 Workload 'echofunctionwasm' accepted. You can now refer to this workload with ID: cmjudmn52omhlsa377d0 on node NC7PXV2DLGXC4LTVM7W7MXYL3WVQFA345IFKJOMYA5ZDZMACLZ53NIIL
```

--------------------------------

### Publish Message to NATS Subject using PUB Command

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol-demo

Publishes a message with the payload 'hello' to the subject 'foo.bar'. The command requires the subject, the length of the payload (5 characters), and the payload on the next line. An '+OK' confirms publish, followed by a 'MSG' for the subscriber.

```text
PUB foo.bar 5
hello
```

--------------------------------

### Get NATS Client Info and Stats (JavaScript)

Source: https://docs.nats.io/using-nats/developer/connecting/events

This JavaScript snippet shows how to access information about a connected NATS client, including the server version and client statistics like messages sent and received. It assumes a `nc` object representing the NATS connection is available.

```javascript
// you can find out where you connected:
t.log(`connected to a nats server version ${nc.info.version}`);

// or information about the data in/out of the client:
const stats = nc.stats();
t.log(`client sent ${stats.outMsgs} messages and received ${stats.inMsgs}`);
```

--------------------------------

### Edit NATS Operator Settings with NSC

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Modify an existing NATS operator's configuration using the 'nsc edit operator' command. This example shows how to set the account server URL, which is crucial for resolvers. Note that operator JWT updates on running servers require manual intervention.

```shell
nsc edit operator --account-jwt-server-url "nats://localhost:4222"
```

--------------------------------

### Query NATS JetStream Stream Information (JSON Output)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/streams

This snippet demonstrates how to get NATS JetStream Stream information in JSON format using the `-j` flag. This is particularly useful for programmatic access and integration with other tools or scripts. The JSON output includes detailed configuration and state of the stream.

```shell
nats str info ORDERS -j
```json
{
  "config": {
    "name": "ORDERS",
    "subjects": [
      "ORDERS.*"
    ],
    "retention": "limits",
    "max_consumers": -1,
    "max_msgs": -1,
    "max_bytes": -1,
    "max_age": 31536000000000000,
    "max_msg_size": -1,
    "storage": "file",
    "discard": "old",
    "num_replicas": 1,
    "duplicate_window": 120000000000
  },
  "created": "2021-02-27T23:49:36.700424Z",
  "state": {
    "messages": 0,
    "bytes": 0,
    "first_seq": 0,
    "first_ts": "0001-01-01T00:00:00Z",
    "last_seq": 0,
    "last_ts": "0001-01-01T00:00:00Z",
    "consumer_count": 0
  }
}
```
```

--------------------------------

### NATS JetStream Pull Subscription in Go

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/consumers

Demonstrates setting up a NATS JetStream connection, creating a stream, publishing messages asynchronously, and subscribing to messages using a pull-based consumer. It fetches messages in batches and acknowledges them. This code requires a running NATS server.

```go
func ExampleJetStream() {
    nc, err := nats.Connect("localhost")
    if err != nil {
        log.Fatal(err)
    }

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
	})

	js.Publish("foo", []byte("Hello JS!"))

	// Publish messages asynchronously.
	for i := 0; i < 500; i++ {
		js.PublishAsync("foo", []byte("Hello JS Async!"))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Create Pull based consumer with maximum 128 inflight.
	sub, _ := js.PullSubscribe("foo", "wq", nats.PullMaxWaiting(128))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

        // Fetch will return as soon as any message is available rather than wait until the full batch size is available, using a batch size of more than 1 allows for higher throughput when needed.
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			msg.Ack()
		}
	}
}

```

--------------------------------

### Signal NATS Servers using Glob Expression

Source: https://docs.nats.io/running-a-nats-service/nats_admin/signals

Applies a signal to multiple NATS server processes that match a glob expression based on their PIDs. This feature, available since NATS v2.10.0, allows batch operations on servers. Example: `nats-server --signal ldm=12*`.

```shell
nats-server --signal ldm=12*
```

--------------------------------

### C# Object Store Management Methods

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

C# methods for managing NATS Object Stores, providing an interface to create and retrieve object stores. It supports creating stores with either a bucket name or a detailed configuration object.

```csharp
// dotnet add package NATS.Net

/// <summary>
/// NATS Object Store context.
/// </summary>
public interface INatsObjContext
{
    /// <summary>
    /// Provides access to the JetStream context associated with the Object Store operations.
    /// </summary>
    INatsJSContext JetStreamContext { get; }

    /// <summary>
    /// Create a new object store.
    /// </summary>
    /// <param name="bucket">Bucket name.</param>
    /// <param name="cancellationToken">A <see cref="CancellationToken"/> used to cancel the API call.</param>
    /// <returns>Object store object.</returns>
    ValueTask<INatsObjStore> CreateObjectStoreAsync(string bucket, CancellationToken cancellationToken = default);

    /// <summary>
    /// Create a new object store.
    /// </summary>
    /// <param name="config">Object store configuration.</param>
    /// <param name="cancellationToken">A <see cref="CancellationToken"/> used to cancel the API call.</param>
    /// <returns>Object store object.</returns>
    ValueTask<INatsObjStore> CreateObjectStoreAsync(NatsObjConfig config, CancellationToken cancellationToken = default);

    /// <summary>
    /// Get an existing object store.

```

--------------------------------

### Add Stream with Mirroring (Bash)

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes/jetstream_leafnodes

This command adds a new stream that mirrors an existing stream from a different JetStream domain. It's used to replicate stream data, ensuring availability even if the source stream's domain is temporarily offline. The configuration prompts guide the user through setting up retention, discard policies, and other stream properties.

```bash
nats  --server nats://acc:acc@localhost:4222 stream add --js-domain hub --mirror test
```
? Stream Name backup-test-leaf
? Storage backend file
? Retention Policy Limits
? Discard Policy Old
? Stream Messages Limit -1
? Message size limit -1
? Maximum message age limit -1
? Maximum individual message size -1
? Replicas 1
? Adjust mirror start No
? Import mirror from a different JetStream domain Yes
? Foreign JetStream domain name leaf
? Delivery prefix
Stream backup-test-leaf was created

Information for Stream backup-test-leaf created 2021-06-28T14:00:43-04:00

Configuration:

     Acknowledgements: true
            Retention: File - Limits
             Replicas: 1
       Discard Policy: Old
     Duplicate Window: 2m0s
    Maximum Messages: unlimited
       Maximum Bytes: unlimited
         Maximum Age: 0.00s
Maximum Message Size: unlimited
   Maximum Consumers: unlimited
               Mirror: test, API Prefix: $JS.leaf.API, Delivery Prefix:


State:

             Messages: 0
                Bytes: 0 B
             FirstSeq: 0
              LastSeq: 0
     Active Consumers: 0

```

--------------------------------

### nats-top In-App Commands

Source: https://docs.nats.io/using-nats/nats-tools/nats_top

Details the interactive commands available within the nats-top interface. These include 'o' for setting sort options, 'n' for limiting connections, 's' to toggle subscription display, '?' for help, and 'q' to quit.

```bash
o<option>
```

```bash
n<limit>
```

```bash
s
```

```bash
?
```

```bash
q
```

--------------------------------

### Configure NATS Server for Leaf Node Connections

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This configuration enables a NATS server to accept connections from leaf nodes on a specified port and sets up token-based authorization for these connections. It's a basic setup for allowing leaf nodes to join the NATS system.

```hcl
leafnodes {
    port: 7422
}
authorization {
    token: "s3cr3t"
}
```

--------------------------------

### Drain NATS Connection (Java)

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Shows how to drain a NATS connection in Java, ensuring all messages are processed before the connection is closed. It uses a CountDownLatch to wait for a message and a CompletableFuture to track the drain operation's completion. The example subscribes to a topic and then initiates the drain process.

```Java
Connection nc = Nats.connect("nats://demo.nats.io:4222");

// Use a latch to wait for a message to arrive
CountDownLatch latch = new CountDownLatch(1);

// Create a dispatcher and inline message handler
Dispatcher d = nc.createDispatcher((msg) -> {
    String str = new String(msg.getData(), StandardCharsets.UTF_8);
    System.out.println(str);
    latch.countDown();
});

// Subscribe
d.subscribe("updates");

// Wait for a message to come in
latch.await();

// Drain the connection, which will close it
CompletableFuture<Boolean> drained = nc.drain(Duration.ofSeconds(10));

// Wait for the drain to complete
drained.get();

```

--------------------------------

### Display NKEYS Directory Tree

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

This command displays the hierarchical structure of the ~/.nkeys directory, showing the organization of credential files and private keys.

```shell
tree ~/.nkeys
```

--------------------------------

### Signal Specific NATS Server by PID File

Source: https://docs.nats.io/running-a-nats-service/nats_admin/signals

Targets a specific NATS server process using the absolute path to its PID file. This method is an alternative to specifying a PID directly, especially in environments where PID files are managed. Example: `nats-server --signal stop=/path/to/pidfile`.

```shell
nats-server --signal stop=/path/to/pidfile
```

--------------------------------

### Drain Subscription and Verify Message Handling in Go

Source: https://docs.nats.io/using-nats/developer/receiving/drain

This Go example demonstrates using the `Drain` method on a NATS subscription. It subscribes to a subject, publishes two messages that are processed slowly, and then calls `Drain`. Draining unsubscribes and waits for pending messages to be processed. A third message published after initiating the drain should not be received. Error handling and synchronization using `sync.WaitGroup` are included.

```go
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

done := sync.WaitGroup{}
done.Add(1)

count := 0
errCh := make(chan error, 1)

msgAfterDrain := "not this one"

// Just to not collide using the demo server with other users.
subject := nats.NewInbox()

// This callback will process each message slowly
sub, err := nc.Subscribe(subject, func(m *nats.Msg) {
    if string(m.Data) == msgAfterDrain {
        errCh <- fmt.Errorf("Should not have received this message")
        return
    }
    time.Sleep(100 * time.Millisecond)
    count++
    if count == 2 {
        done.Done()
    }
})

// Send 2 messages
for i := 0; i < 2; i++ {
    nc.Publish(subject, []byte("hello"))
}

// Call Drain on the subscription. It unsubscribes but
// wait for all pending messages to be processed.
if err := sub.Drain(); err != nil {
    log.Fatal(err)
}

// Send one more message, this message should not be received
nc.Publish(subject, []byte(msgAfterDrain))

// Wait for the subscription to have processed the 2 messages.
done.Wait()

// Now check that the 3rd message was not received
select {
case e := <-errCh:
    log.Fatal(e)
case <-time.After(200 * time.Millisecond):
    // OK!
}

```

--------------------------------

### Drain NATS Connection (Go)

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Demonstrates draining a NATS connection in Go, processing pending messages before the connection is closed. This involves setting up a wait group to signal completion and an error handler. The example simulates message processing with a delay and ensures all messages are handled before the drain completes.

```Go
wg := sync.WaitGroup{}
wg.Add(1)

errCh := make(chan error, 1)

// To simulate a timeout, you would set the DrainTimeout()
// to a value less than the time spent in the message callback,
// so say: nats.DrainTimeout(10*time.Millisecond).

nc, err := nats.Connect("demo.nats.io",
    nats.DrainTimeout(10*time.Second),
    nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
        errCh <- err
    }),
    nats.ClosedHandler(func(_ *nats.Conn) {
        wg.Done()
    }))
if err != nil {
    log.Fatal(err)
}

// Just to not collide using the demo server with other users.
subject := nats.NewInbox()

// Subscribe, but add some delay while processing.
if _, err := nc.Subscribe(subject, func(_ *nats.Msg) {
    time.Sleep(200 * time.Millisecond)
}); err != nil {
    log.Fatal(err)
}

// Publish a message
if err := nc.Publish(subject, []byte("hello")); err != nil {
    log.Fatal(err)
}

// Drain the connection, which will close it when done.
if err := nc.Drain(); err != nil {
    log.Fatal(err)
}

// Wait for the connection to be closed.
wg.Wait()

// Check if there was an error
select {
case e := <-errCh:
    log.Fatal(e)
default:
}

```

--------------------------------

### Subscribe with NKEY Authentication

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Demonstrates how to subscribe to a NATS subject using the `nats` CLI with NKEY authentication. The NKEY file is specified using the `--nkey` flag.

```shell
nats -s nats://localhost:4222 sub --nkey=a.nk ">"
```

--------------------------------

### Key-Value Management API

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This section details the fundamental operations for managing Key-Value stores, including creation, retrieval, and deletion.

```APIDOC
## Key-Value Management API

### Description

Provides methods for interacting with Key-Value stores in JetStream.

### Methods

- `key_value(bucket: str)`
- `create_key_value(config: Optional[api.KeyValueConfig] = None, **params)`
- `delete_key_value(bucket: str)`

### Parameters

#### `key_value(bucket: str)`

- **bucket** (str) - Required - The name of the Key-Value bucket.

#### `create_key_value(config: Optional[api.KeyValueConfig] = None, **params)`

- **config** (api.KeyValueConfig) - Optional - Configuration for the new Key-Value store.
- **params** - Additional parameters for store creation.

#### `delete_key_value(bucket: str)`

- **bucket** (str) - Required - The name of the Key-Value bucket to delete.

### Responses

#### `key_value`

- **Success Response (200)**: Returns a `KeyValue` object representing the store.

#### `create_key_value`

- **Success Response (200)**: Returns a `KeyValue` object upon successful creation.

#### `delete_key_value`

- **Success Response (200)**: Returns `true` if the Key-Value store was deleted successfully, `false` otherwise.

### Examples

#### Python Example
```python
# Assuming 'kv' is an instance of the KeyValue interface
kv_store = await kv.key_value("my_bucket")
new_kv_store = await kv.create_key_value(api.KeyValueConfig(bucket="new_bucket"))
success = await kv.delete_key_value("bucket_to_delete")
```

#### C# Example
```csharp
// Assuming 'natsConnection' is an established NATS connection
var kvStore = await natsConnection.CreateStoreAsync("my_bucket");
var kvStores = natsConnection.GetBucketNamesAsync();
var statuses = natsConnection.GetStatusesAsync();
var deleted = await natsConnection.DeleteStoreAsync("bucket_to_delete");
```

#### C Example
```c
// Assuming 'js' is a jsCtx pointer
kvConfig cfg;
kvConfig_Init(&cfg);
kvStore *kv;
js_CreateKeyValue(&kv, js, &cfg);
kvStore *existing_kv;
js_KeyValue(&existing_kv, js, "my_bucket");
js_DeleteKeyValue(js, "bucket_to_delete");
kvStore_Destroy(existing_kv);
```
```

--------------------------------

### Slice Tokens with SliceFromLeft and SliceFromRight

Source: https://docs.nats.io/nats-concepts/subject_mapping

The `SliceFromLeft` and `SliceFromRight` functions allow you to split NATS subject tokens into smaller parts at a specified interval from either the start or end. This is useful for breaking down long tokens into more manageable segments. Usage involves specifying the wildcard index and the number of characters to slice.

```nats
nats server mapping "*" "{{slicefromleft(1,2)}}" 1234567
nats server mapping "*" "{{slicefromright(1,2)}}" 1234567
```

--------------------------------

### Configure NATS Ping/Pong Settings (Java)

Source: https://docs.nats.io/using-nats/developer/connecting/pingpong

Illustrates configuring the NATS client in Java with a ping interval of 20 seconds and a maximum of 5 outstanding pings. It uses the Options.Builder to set these parameters and demonstrates connection establishment within a try-with-resources block for automatic closing.

```Java
Options options = new Options.Builder()
    .server("nats://demo.nats.io")
    .pingInterval(Duration.ofSeconds(20)) // Set Ping Interval
    .maxPingsOut(5) // Set max pings in flight
    .build();

// Connection is AutoCloseable
try (Connection nc = Nats.connect(options)) {
    // Do something with the connection
}

```

--------------------------------

### Configure Dedicated Account Routes (NATS)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering/v2_routes

Specifies a list of accounts that will have a dedicated route connection between NATS servers. The system account always gets a dedicated route by default. This feature is only available when connection pooling is enabled (pool_size != -1).

```hcl
cluster {
  accounts: [acc1, acc2]
}
```

--------------------------------

### NATS User JWT and NKEY Seed Generation

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/jwt_nkey_auth

This section shows the structure of a NATS user JWT and its corresponding NKEY seed. The JWT contains authorization information for the user, while the NKEY seed is a private key used for signing and proving identity. NKEY seeds are sensitive and should be kept secret.

```text
-----BEGIN NATS USER JWT-----
...
------END NATS USER JWT------
************************* IMPORTANT *************************
NKEY Seed printed below can be used to sign and prove identity.
NKEYs are sensitive and should be treated as secrets.
-----BEGIN USER NKEY SEED-----
...
------END USER NKEY SEED------
*************************************************************

```

--------------------------------

### Telnet Connection to NATS Server

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol

This snippet demonstrates how to establish a telnet connection to a NATS server on the default port. It shows the initial connection attempt and the subsequent INFO message received from the server, which contains server details.

```bash
telnet demo.nats.io 4222

```

```text
Trying 107.170.221.32...
Connected to demo.nats.io.
Escape character is '^]'.
INFO {"server_id":"Zk0GQ3JBSrg3oyxCRRlE09","version":"1.2.0","proto":1,"go":"go1.10.3","host":"0.0.0.0","port":4222,"max_payload":1048576,"client_id":2392}

```

--------------------------------

### Receive JSON Data in C#

Source: https://docs.nats.io/using-nats/developer/receiving/structure

This C# example demonstrates subscribing to NATS subjects for different data types: int, string, and byte arrays. It utilizes the built-in JSON serialization of the NATS.Net library. The code iterates over incoming messages asynchronously and prints the received data to the console.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

// NATS .NET has a built-in serializer that does the 'unsurprising' thing
// for most types. Most primitive types are serialized as expected.
// For any other type, JSON serialization is used. You can also provide
// your own serializers by implementing the INatsSerializer and
// INasSerializerRegistry interfaces. See also for more information:
// https://nats-io.github.io/nats.net/documentation/advanced/serialization.html
await using var nc = new NatsClient();

CancellationTokenSource cts = new();

// Subscribe for int, string, bytes, json
List<Task> tasks =
[
    Task.Run(async () =>
    {
        await foreach (var msg in nc.SubscribeAsync<int>("x.int", cancellationToken: cts.Token))
        {
            Console.WriteLine($"Received int: {msg.Data}");
        }
    }),

    Task.Run(async () =>
    {
        await foreach (var msg in nc.SubscribeAsync<string>("x.string", cancellationToken: cts.Token))
        {
            Console.WriteLine($"Received string: {msg.Data}");
        }
    }),

    Task.Run(async () =>
    {
        await foreach (var msg in nc.SubscribeAsync<byte[]>("x.bytes", cancellationToken: cts.Token))
        {
            if (msg.Data != null)
            {
                Console.Write($"Received bytes: ");
                foreach (var b in msg.Data)
                {
                    Console.Write("0x{0:X2} ", b);
                }
                Console.WriteLine();
            }
        }
    }),
]
```

--------------------------------

### nats-top Sorting Options

Source: https://docs.nats.io/using-nats/nats-tools/nats_top

Illustrates how to sort the connection list in nats-top using the '-sort' command-line flag or the 'o<option>' in-app command. The sorting can be applied to various metrics like connection ID, subscriptions, message counts, byte counts, language, and version.

```bash
nats-top -sort bytes_to
```

```bash
o<option>
```

--------------------------------

### Handle Slow Consumer Errors in Python with Async

Source: https://docs.nats.io/using-nats/developer/connecting/events/slow

Configure a Python asynchronous NATS client to handle slow consumer errors. An `error_cb` function is defined to catch `nats.aio.errors.ErrSlowConsumer` and unsubscribe from the subject to prevent further message processing issues. This example demonstrates publishing and subscribing with a pending message limit.

```Python
   nc = NATS()

   async def error_cb(e):
     if type(e) is nats.aio.errors.ErrSlowConsumer:
       print("Slow consumer error, unsubscribing from handling further messages...")
       await nc.unsubscribe(e.sid)

   await nc.connect(
      servers=["nats://demo.nats.io:4222"],
      error_cb=error_cb,
      )

   msgs = []
   future = asyncio.Future()
   async def cb(msg):
       nonlocal msgs
       nonlocal future
       print(msg)
       msgs.append(msg)

       if len(msgs) == 3:
         # Head of line blocking on other messages caused
         # by single message processing taking too long...
         await asyncio.sleep(1)

   await nc.subscribe("updates", cb=cb, pending_msgs_limit=5)

   for i in range(0, 10):
     await nc.publish("updates", "msg #{}".format(i).encode())
     await asyncio.sleep(0)

   try:
     await asyncio.wait_for(future, 1)
   except asyncio.TimeoutError:
     pass

   for msg in msgs:
     print("[Received]", msg)

   await nc.close()
```

--------------------------------

### Publish to NATS JetStream Stream (C#)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/publish

This C# code snippet demonstrates connecting to NATS, creating a JetStream context, and publishing messages. It covers creating a stream, publishing a single message, and publishing multiple messages concurrently using `PublishConcurrentAsync`. It also includes error checking for publish acknowledgments.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.JetStream;
using NATS.Client.JetStream.Models;

await using var client = new NatsClient();

INatsJSContext js = client.CreateJetStreamContext();

// Create a stream
var streamConfig = new StreamConfig(name: "example-stream", subjects: ["example-subject"]);
await js.CreateStreamAsync(streamConfig);

// Publish a message
{
    PubAckResponse ack = await js.PublishAsync("example-subject", "Hello, JetStream!");
    ack.EnsureSuccess();
}

// Publish messages concurrently
List<NatsJSPublishConcurrentFuture> futures = new();
for (var i = 0; i < 500; i++)
{
    NatsJSPublishConcurrentFuture future
        = await js.PublishConcurrentAsync("example-subject", "Hello, JetStream 1!");
    futures.Add(future);
}

foreach (var future in futures)
{
    await using (future)
    {
        PubAckResponse ack = await future.GetResponseAsync();
        ack.EnsureSuccess();
    }
}
```

--------------------------------

### NATS Service Reply (Bash)

Source: https://docs.nats.io/using-nats/nats-tools/nsc/services

Sets up a process to reply to requests on a specific subject using the NATS CLI. This simulates a service listening for messages.

```bash
nats reply --creds ~/.nkeys/creds/O/A/U.creds help "I will help"
```

```bash
nsc reply --account A --user U help "I will help"
```

--------------------------------

### Drain NATS Subscription (C#)

Source: https://docs.nats.io/using-nats/developer/receiving/drain

Demonstrates draining a NATS subscription in C# using the NATS.Net library. This snippet shows how to subscribe to a subject and process incoming messages, with a mechanism to signal the subscription completion. The example uses a cancellation token and a loop to wait for messages before the subscription effectively ends.

```C#
// dotnet add package NATS.Net
using NATS.Net;

var client = new NatsClient();

var subject = client.Connection.NewInbox();

// Make sure to use a cancellation token to end all subscriptions
using var cts = new CancellationTokenSource();

var sync = false;
var process = Task.Run(async () =>
{
    await foreach (var msg in client.SubscribeAsync<int>(subject, cancellationToken: cts.Token))
    {
        if (msg.Data == -1)
        {
            sync = true;
            continue;
        }
        Console.WriteLine($"Received: {msg.Data}");
        await Task.Delay(TimeSpan.FromMilliseconds(300));
    }

    Console.WriteLine("Subscription completed");
});

// Make sure the subscription is ready
while (sync == false)
{
    await Task.Delay(TimeSpan.FromMilliseconds(100));
}

```

--------------------------------

### Authenticate MQTT Users with JWT in Operator Mode

Source: https://docs.nats.io/running-a-nats-service/configuration/mqtt/mqtt_config

This example shows how to authenticate MQTT clients using JWTs in NATS operator mode. The JWT is passed as the MQTT password, and any valid username can be used. The JWT must have the `Bearer` boolean set to true, which can be configured using the `nsc` command-line tool.

```shell
nsc edit user --name U --account A --bearer
```

--------------------------------

### Subscribe to Specific Wildcard Subjects with NATS in JavaScript

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This JavaScript example demonstrates subscribing to more specific wildcard subjects within the 'time.us.*' pattern. It uses a switch statement to handle different specific subjects like 'time.us.east', 'time.us.central', etc., and formats the time based on the timezone. Note that proper timezone handling in Node.js might require additional libraries.

```JavaScript
nc.subscribe("time.us.*", (_err, msg) => {
    // converting timezones correctly in node requires a library
    // this doesn't take into account *many* things.
    let time;
    switch (msg.subject) {
      case "time.us.east":
        time = new Date().toLocaleTimeString("en-us", {
          timeZone: "America/New_York",
        });
        break;
      case "time.us.central":
        time = new Date().toLocaleTimeString("en-us", {
          timeZone: "America/Chicago",
        });
        break;
      case "time.us.mountain":
        time = new Date().toLocaleTimeString("en-us", {
          timeZone: "America/Denver",
        });
        break;
      case "time.us.west":
        time = new Date().toLocaleTimeString("en-us", {
          timeZone: "America/Los_Angeles",
        });
        break;
      default:
        time = "I don't know what you are talking about Willis";
    }
    t.log(subject, time);
});
```

--------------------------------

### NATS Object Store Manager Interface (Go)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/object

Defines the ObjectStoreManager interface for managing object stores in NATS. It includes methods for binding to existing stores, creating, updating, and deleting stores, and listing store names and statuses. Dependencies include 'context' and 'nats.go'.

```Go
package nats

import (
	"context"
	"io"
)

// ObjectStoreManager is used to manage object stores. It provides methods
// for CRUD operations on object stores.
type ObjectStoreManager interface {
	// ObjectStore will look up and bind to an existing object store
	// instance.
	//
	// If the object store with given name does not exist, ErrBucketNotFound
	// will be returned.
	ObjectStore(ctx context.Context, bucket string) (ObjectStore, error)

	// CreateObjectStore will create a new object store with the given
	// configuration.
	//
	// If the object store with given name already exists, ErrBucketExists
	// will be returned.
	CreateObjectStore(ctx context.Context, cfg ObjectStoreConfig) (ObjectStore, error)

	// UpdateObjectStore will update an existing object store with the given
	// configuration.
	//
	// If the object store with given name does not exist, ErrBucketNotFound
	// will be returned.
	UpdateObjectStore(ctx context.Context, cfg ObjectStoreConfig) (ObjectStore, error)

	// CreateOrUpdateObjectStore will create a new object store with the given
	// configuration if it does not exist, or update an existing object store
	// with the given configuration.
	CreateOrUpdateObjectStore(ctx context.Context, cfg ObjectStoreConfig) (ObjectStore, error)

	// DeleteObjectStore will delete the provided object store.
	//
	// If the object store with given name does not exist, ErrBucketNotFound
	// will be returned.
	DeleteObjectStore(ctx context.Context, bucket string) error

	// ObjectStoreNames is used to retrieve a list of bucket names.
	// It returns an ObjectStoreNamesLister exposing a channel to receive
	// the names of the object stores.
	//
	// The lister will always close the channel when done (either all names
	// have been read or an error occurred) and therefore can be used in a
	// for-range loop.
	ObjectStoreNames(ctx context.Context) ObjectStoreNamesLister

	// ObjectStores is used to retrieve a list of bucket statuses.
	// It returns an ObjectStoresLister exposing a channel to receive
	// the statuses of the object stores.
	//
	// The lister will always close the channel when done (either all statuses
	// have been read or an error occurred) and therefore can be used in a
	// for-range loop.
	ObjectStores(ctx context.Context) ObjectStoresLister
}

// ObjectStore contains methods to operate on an object store.
// Using the ObjectStore interface, it is possible to:
//
// - Perform CRUD operations on objects (Get, Put, Delete).
//   Get and put expose convenience methods to work with
//   byte slices, strings and files, in addition to streaming [io.Reader]
// - Get information about an object without retrieving it.
// - Update the metadata of an object.
// - Add links to other objects or object stores.
// - Watch for updates to a store
// - List information about objects in a store
// - Retrieve status and configuration of an object store.
type ObjectStore interface {
	// Put will place the contents from the reader into a new object. If the
	// object already exists, it will be overwritten. The object name is
	// required and is taken from the ObjectMeta.Name field.
	//
	// The reader will be read until EOF. ObjectInfo will be returned, containing
	// the object's metadata, digest and instance information.
	Put(ctx context.Context, obj ObjectMeta, reader io.Reader) (*ObjectInfo, error)

	// PutBytes is convenience function to put a byte slice into this object
	// store under the given name.
	//
	// ObjectInfo will be returned, containing the object's metadata, digest
	// and instance information.
	PutBytes(ctx context.Context, name string, data []byte) (*ObjectInfo, error)

	// PutString is convenience function to put a string into this object
	// store under the given name.
	//
	// ObjectInfo will be returned, containing the object's metadata, digest
	// and instance information.
	PutString(ctx context.Context, name string, data string) (*ObjectInfo, error)

	// PutFile is convenience function to put a file contents into this
	// object store. The name of the object will be the path of the file.
	//
	// ObjectInfo will be returned, containing the object's metadata, digest
	// and instance information.
	PutFile(ctx context.Context, file string) (*ObjectInfo, error)

	// Get will pull the named object from the object store. If the object
	// does not exist, ErrObjectNotFound will be returned.
	//
	// The returned ObjectResult will contain the object's metadata and a
	// reader to read the object's contents. The reader will be closed when

```

--------------------------------

### List and Select NATS Contexts

Source: https://docs.nats.io/using-nats/nats-tools/nats_cli

Shows how to list all known NATS contexts and select a specific context to be the default. The output indicates the current default context with an asterisk.

```shell
nats context ls
nats context select
```

--------------------------------

### Subscribe to a NATS Subject using SUB Command

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol-demo

Subscribes to the wildcard subject 'foo.*' using a specific subscription ID '90'. An '+OK' response confirms the successful subscription to the specified subject.

```text
SUB foo.* 90
```

--------------------------------

### NATS Leaf Node Remote Connection Configuration

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes

This configuration block defines the remote connections for a NATS leaf node. It specifies the URL of the remote NATS server (NGS in this case) and the path to the user credentials file for authentication. This setup allows the leaf node to connect to a global NATS service.

```hcl
leafnodes {
    remotes = [
        {
          url: "tls://connect.ngs.global"
          credentials: "/Users/alberto/.nkeys/creds/synadia/leaftest/leaftestuser.creds"
        },
    ]
}

```

--------------------------------

### Connect to NATS Cluster in Java

Source: https://docs.nats.io/using-nats/developer/connecting/cluster

Illustrates connecting to a NATS cluster using the Java NATS client. It utilizes the `Options.Builder` to specify a comma-separated string of server URLs. The connection is then established, and a placeholder comment indicates where to perform operations before closing the connection.

```Java
Options options = new Options.Builder()
    .server("nats://127.0.0.1:1222,nats://127.0.0.1:1223,nats://127.0.0.1:1224")
    .build();
Connection nc = Nats.connect(options);

// Do something with the connection

nc.close();
```

--------------------------------

### Add NATS Operator using NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This command adds a new operator to your NATS environment using the `nsc` command-line tool. It initializes the operator's keys and configuration. This is a prerequisite for setting up accounts and users within that operator's domain.

```shell
nsc add operator -n memory
```

--------------------------------

### Connect to Default NATS Server in C#

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Shows how to connect to the default NATS server using the NATS.Net library in C#. It utilizes `await using` for resource management and highlights that `ConnectAsync` is optional as it's called automatically when needed.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient();

// It's optional to call ConnectAsync()
// as it will be called when needed automatically
await client.ConnectAsync();
```

--------------------------------

### Configure NATS Server OCSP Stapling Modes

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/ocsp

These configuration examples show advanced OCSP stapling modes for the NATS Server. The 'must' mode enforces stapling for Must-Staple certificates and shuts down the server if a staple is revoked. The 'always' mode enforces stapling for all certificates and also shuts down the server if a staple is revoked. Both modes allow overriding the OCSP responder URL.

```yaml
ocsp {
  mode: must
  url: "http://ocsp.example.net"
}
```

```yaml
ocsp {
  mode: always
  url: "http://ocsp.example.net"
}
```

--------------------------------

### Get Account Statistics (accstatz) - JSON Response

Source: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring

This JSON structure represents the response from the /accstatz endpoint, providing statistics for NATS accounts. It includes details on active connections, leaf nodes, subscriptions, and message transfer rates for sent and received data. The 'unused' argument can be set to include accounts without current connections.

```json
{
  "server_id": "NDJ5M4F5WAIBUA26NJ3QMH532AQPN7QNTJP3Y4SBHSHL4Y7QUAKNJEAF",
  "now": "2022-10-19T17:16:20.881296749Z",
  "account_statz": [
    {
      "acc": "default",
      "conns": 31,
      "leafnodes": 2,
      "total_conns": 33,
      "num_subscriptions": 45,
      "sent": {
        "msgs": 1876970,
        "bytes": 246705616
      },
      "received": {
        "msgs": 1347454,
        "bytes": 219438308
      },
      "slow_consumers": 29
    },
    {
      "acc": "$G",
      "conns": 1,
      "leafnodes": 0,
      "total_conns": 1,
      "num_subscriptions": 3,
      "sent": {
        "msgs": 0,
        "bytes": 0
      },
      "received": {
        "msgs": 107,
        "bytes": 1094
      },
      "slow_consumers": 0
    }
  ]
}
```

--------------------------------

### Configure NATS Authentication Timeout

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/auth_timeout

Example configuration block for setting the authentication timeout in NATS. The timeout is specified in seconds (can be fractional). A value of '0' or no value defaults to 1 second more than the tls_timeout. Invalid values also default to 1 second. This setting limits the time a client has to authenticate before disconnection.

```hocon
authorization:
    timeout: 3
    users:
        - {user: a, password b}
        - {user: b, password a}
```

--------------------------------

### Run Second NATS Leaf Node (Docker)

Source: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-leafnodes-docker

Launches the second NATS leaf node in a Docker container, similar to the first. It uses a different local port (4333) because 4222 is already in use and mounts the 'blue' user's credentials. This is for the 'blue' user's leaf node.

```shell
docker run  -p 4333:4222 -v leafnode.conf:/leafnode.conf -v /etc/ssl/cert.pem:/etc/ssl/cert.pem -v default-blue.creds:/ngs.creds  nats:latest -c /leafnode.conf

```

--------------------------------

### Key-Value Store Operations

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Provides details on various operations for the NATS Key-Value store, including creating, updating, deleting, and retrieving keys.

```APIDOC
## Key-Value Store Operations

This section covers the fundamental operations for interacting with the NATS Key-Value store.

### Creating a Key

These functions create a key-value pair in the store if the key does not already exist.

#### `kvStore_CreateString`

*   **Description**: Places the value (as a string) for the key into the store if and only if the key does not exist.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `rev` (uint64_t *) - Output parameter for the revision number.
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `key` (const char *) - The key to create.
    *   `data` (const char *) - The string data to store.
*   **Returns**: `natsStatus` indicating success or failure.

### Updating a Key

These functions update an existing key-value pair in the store, provided the latest revision matches.

#### `kvStore_Update`

*   **Description**: Updates the value for the key into the store if and only if the latest revision matches.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `rev` (uint64_t *) - Output parameter for the revision number.
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `key` (const char *) - The key to update.
    *   `data` (const void *) - Pointer to the data to store.
    *   `len` (int) - The length of the data.
    *   `last` (uint64_t) - The last known revision of the key.
*   **Returns**: `natsStatus` indicating success or failure.

#### `kvStore_UpdateString`

*   **Description**: Updates the value (as a string) for the key into the store if and only if the latest revision matches.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `rev` (uint64_t *) - Output parameter for the revision number.
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `key` (const char *) - The key to update.
    *   `data` (const char *) - The string data to store.
    *   `last` (uint64_t) - The last known revision of the key.
*   **Returns**: `natsStatus` indicating success or failure.

### Deleting Keys

Operations to delete specific keys or purge entire key/value buckets.

#### `kvStore_Delete`

*   **Description**: Deletes a key by placing a delete marker and leaving all revisions.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `key` (const char *) - The key to delete.
*   **Returns**: `natsStatus` indicating success or failure.

#### `kvStore_Purge`

*   **Description**: Deletes a key by placing a purge marker and removing all revisions.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `key` (const char *) - The key to purge.
    *   `opts` (kvPurgeOptions *) - Options for purging.
*   **Returns**: `natsStatus` indicating success or failure.

#### `kvStore_PurgeDeletes`

*   **Description**: Purges and removes delete markers.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `opts` (kvPurgeOptions *) - Options for purging.
*   **Returns**: `natsStatus` indicating success or failure.

### Retrieving Keys

Operations to get a list of all keys currently having a value associated.

#### `kvStore_Keys`

*   **Description**: Returns all keys in the bucket.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `list` (kvKeysList *) - Pointer to a structure to hold the list of keys.
    *   `kv` (kvStore *) - Pointer to the key-value store.
    *   `opts` (kvWatchOptions *) - Options for watching keys.
*   **Returns**: `natsStatus` indicating success or failure.

#### `kvKeysList_Destroy`

*   **Description**: Destroys the list of KeyValue store key strings.
*   **Method**: NATS_EXTERN
*   **Parameters**:
    *   `list` (kvKeysList *) - Pointer to the keys list to destroy.


```

--------------------------------

### NATS CLI: Subscribe and Publish to Server A

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This command sequence demonstrates subscribing to a 'hello' subject on server A (port 4222) and then publishing a message 'world_4222' to the same subject and server. It verifies that messages can be sent and received within the NATS cluster.

```bash
nats sub -s "nats://127.0.0.1:4222" hello &
nats pub -s "nats://127.0.0.1:4222" hello world_4222
```

--------------------------------

### Discovering Servers

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

Publishes a request to discover servers within the NATS cluster, returning their ID and name.

```APIDOC
## GET $SYS.REQ.SERVER.PING.IDZ

### Description
To discover servers in the cluster to get their ID and name, publish a request to `$SYS.REQ.SERVER.PING.IDZ`.

### Method
REQUEST

### Endpoint
$SYS.REQ.SERVER.PING.IDZ

### Parameters
#### Query Parameters
None

#### Request Body
None

### Request Example
```shell
nats request --creds ~/.nkeys/SAOP/accounts/SYS/users/SYSU.creds $SYS.REQ.SERVER.PING.IDZ ""
```

### Response
#### Success Response (200)
- **host** (string) - The host address of the server.
- **id** (string) - The unique identifier of the server.
- **name** (string) - The name of the server.

#### Response Example
```json
{
  "host": "0.0.0.0",
  "id": "NC7AKPQRC6CIZGWRJOTVFIGVSL7VW7WXTQCTUJFNG7HTCMCKQTGE5PUL",
  "name": "n1"
}
```
```

--------------------------------

### Add Durable Pull Consumer in NATS JetStream

Source: https://docs.nats.io/running-a-nats-service/configuration/leafnodes/jetstream_leafnodes

This command adds a new durable pull consumer to a NATS JetStream stream. It prompts for consumer name, delivery target, start and replay policies, filter subjects, maximum deliveries, acknowledgements, and the target stream. The output confirms the consumer creation and displays its configuration and state.

```bash
nats  --server nats://acc:acc@localhost:4222 consumer add  --js-domain hub
```

--------------------------------

### Connect to Default NATS Server in Java

Source: https://docs.nats.io/using-nats/developer/connecting/default_server

Establishes a connection to the default NATS server using the Java NATS client. The code snippet shows how to create a connection object and includes a placeholder for operations and connection closure.

```java
Connection nc = Nats.connect();

// Do something with the connection

nc.close();
```

--------------------------------

### Connect to NATS using C#

Source: https://docs.nats.io/using-nats/developer/connecting/userpass

Demonstrates how to add the NATS.Net package and establish a connection to a NATS server with authentication options. Note that username/password are not supported in the URL for the .NET client.

```csharp
// dotnet add package NATS.Net
using NATS.Net;
using NATS.Client.Core;

await using var nc = new NatsClient(new NatsOpts
{
    // .NET client doesn't support username/password in URLs
    // use `Username` and `Password` options.
    Url = "nats://demo.nats.io:4222",
    AuthOpts = new NatsAuthOpts
    {
        Username = "myname",
        Password = "password",
    }
});
```

--------------------------------

### Publish with NKEY Authentication

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Shows how to publish a message to a NATS subject using the `nats` CLI with NKEY authentication. The NKEY file is specified using the `--nkey` flag.

```shell
nats -s nats://localhost:4222 pub --nkey=b.nk foo nkey
```

--------------------------------

### NATS Server Clustering Command Line Options (Bash)

Source: https://docs.nats.io/running-a-nats-service/configuration/clustering

This snippet lists essential command-line options for configuring NATS server clustering. It specifies how to define routes for server discovery and connection, and the cluster URL for soliciting routes. It also notes that the '-cluster' option must be specified when using '-routes'.

```bash
--routes [rurl-1, rurl-2]     Routes to solicit and connect
--cluster nats://host:port    Cluster URL for solicited routes
```

--------------------------------

### NATS Configuration File Syntax

Source: https://docs.nats.io/nats-server/configuration

Illustrates the basic syntax for NATS configuration files, including comments, key-value assignments, arrays, and maps. Supports various delimiters and optional semicolons.

```plaintext
# Lines can be commented with # and //
# Values can be assigned to properties with delimiters:
# Equals sign: foo = 2
# Colon: foo: 2
# Whitespace: foo 2
# Arrays are enclosed in brackets: ["a", "b", "c"]
# Maps are enclosed in braces: {foo: 2}
# Maps can be assigned with no delimiter accounts { SYS {...}, cloud-user {...} }
# Semicolons can be optionally used as terminators host: 127.0.0.1; port: 4222;
```

--------------------------------

### NATS Configuration: Variable Definition and Usage

Source: https://docs.nats.io/running-a-nats-service/configuration

Demonstrates how to define and reference variables within NATS configuration files. It covers block scoping, the '$' prefix for references, and how undefined variables can be resolved from environment variables.

```shell
# Define a variable in the config
TOKEN: "secret"

# Reference the variable
authorization {
    token: $TOKEN
}
```

```shell
# Define a variable in the config
# But TOKEN is never used resulting in a config parsing error
TOKEN: "secret"

# Reference the variable
authorization {
    token: "another secret"
}
```

```shell
# TOKEN is defined in the environment
authorization {
    token: $TOKEN
}
```

--------------------------------

### NATS Include Directive for Modular Configuration

Source: https://docs.nats.io/nats-server/configuration

Explains and demonstrates the `include` directive in NATS configuration files. This allows splitting configurations into multiple files, which must use relative paths from the main configuration file.

```plaintext
listen: 127.0.0.1:4222
include ./auth.conf
```

```plaintext
authorization: {
    token: "f0oBar"
}
```

```shell
> nats-server -c server.conf
```

--------------------------------

### NATS CLI: Pushing Account Changes to Servers

Source: https://docs.nats.io/using-nats/nats-tools/nsc/revocation

Demonstrates commands for pushing local account changes, such as revocations, to NATS servers. This is necessary for the servers to recognize and apply the updated account configurations. It shows interactive and direct connection methods.

```bash
nsc push -i
```

```bash
nsc push -a B -u nats://localhost
```

--------------------------------

### Configure NATS Discovered Servers Callback (C)

Source: https://docs.nats.io/using-nats/developer/connecting/events/events

Sets a callback function for discovered servers events in the NATS C client. It demonstrates creating options, setting the discovered servers callback, connecting, and then retrieving and freeing the discovered server list.

```c
static void
discoveredServersCB(natsConnection *conn, void *closure)
{
    natsStatus  s         = NATS_OK;
    char        **servers = NULL;
    int         count     = 0;

    s = natsConnection_GetDiscoveredServers(conn, &servers, &count);
    if (s == NATS_OK)
    {
        int i;

        // Do something...
        for (i=0; i<count; i++)
            printf("Discovered server: %s\n", servers[i]);

        // Free allocated memory
        for (i=0; i<count; i++)
            free(servers[i]);
        free(servers);
    }
}

(...)

natsConnection      *conn      = NULL;
natsOptions         *opts      = NULL;
natsStatus          s          = NATS_OK;

s = natsOptions_Create(&opts);
if (s == NATS_OK)
    s = natsOptions_SetDiscoveredServersCB(opts, discoveredServersCB, NULL);
if (s == NATS_OK)
    s = natsConnection_Connect(&conn, opts);

(...)


// Destroy objects that were created
natsConnection_Destroy(conn);
natsOptions_Destroy(opts);
```

--------------------------------

### Add NATS JetStream Stream Configuration and Publish

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/streams

Initializes a stream configuration, adds a stream to the JetStream context, and publishes messages to the stream. It handles both synchronous and asynchronous publishing, including completing asynchronous operations and handling timeouts. It also retrieves and prints stream information before and after publishing.

```c
            // Initialize the configuration structure.
            jsStreamConfig_Init(&cfg);
            cfg.Name = stream;
            // Set the subject
            cfg.Subjects = (const char*[1]){subj};
            cfg.SubjectsLen = 1;
            // Make it a memory stream.
            cfg.Storage = js_MemoryStorage;
            // Add the stream,
            s = js_AddStream(&si, js, &cfg, NULL, &jerr);
        }
        if (s == NATS_OK)
        {
            printf("Stream %s has %" PRIu64 " messages (%%" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            // Need to destroy the returned stream object.
            jsStreamInfo_Destroy(si);
        }
    }

    if (s == NATS_OK)
        s = natsStatistics_Create(&stats);

    if (s == NATS_OK)
    {
        printf("\nSending %" PRId64 " messages to subject '%s'\n", total, stream);
        start = nats_Now();
    }

    for (count = 0; (s == NATS_OK) && (count < total); count++)
    {
        if (async)
            s = js_PublishAsync(js, subj, (const void*) payload, dataLen, NULL);
        else
        {
            jsPubAck *pa = NULL;

            s = js_Publish(&pa, js, subj, (const void*) payload, dataLen, NULL, &jerr);
            if (s == NATS_OK)
            {
                if (pa->Duplicate)
                    printf("Got a duplicate message! Sequence=%%" PRIu64 "\n", pa->Sequence);

                jsPubAck_Destroy(pa);
            }
        }
    }

    if ((s == NATS_OK) && async)
    {
        jsPubOptions    jsPubOpts;

        jsPubOptions_Init(&jsPubOpts);
        // Let's set it to 30 seconds, if getting "Timeout" errors,
        // this may need to be increased based on the number of messages
        // being sent.
        jsPubOpts.MaxWait = 30000;
        s = js_PublishAsyncComplete(js, &jsPubOpts);
        if (s == NATS_TIMEOUT)
        {
            // Let's get the list of pending messages. We could resend,
            // etc, but for now, just destroy them.
            natsMsgList list;

            js_PublishAsyncGetPendingList(&list, js);
            natsMsgList_Destroy(&list);
        }
    }

    if (s == NATS_OK)
    {
        jsStreamInfo *si = NULL;

        elapsed = nats_Now() - start;
        printStats(STATS_OUT, conn, NULL, stats);
        printPerf("Sent");

        if (errors != 0)
            printf("There were %d asynchronous errors\n", errors);

        // Let's report some stats after the run
        s = js_GetStreamInfo(&si, js, stream, NULL, &jerr);
        if (s == NATS_OK)
        {
            printf("\nStream %s has %" PRIu64 " messages (%%" PRIu64 " bytes)\n",
                si->Config->Name, si->State.Msgs, si->State.Bytes);

            jsStreamInfo_Destroy(si);
        }
    }
    if (delStream && (js != NULL))
    {
        printf("\nDeleting stream %s: ", stream);
        s = js_DeleteStream(js, stream, NULL, &jerr);
        if (s == NATS_OK)
            printf("OK!");
        printf("\n");
    }
    if (s != NATS_OK)
    {
        printf("Error: %u - %s - jerr=%u\n", s, natsStatus_GetText(s), jerr);
        nats_PrintLastErrorStack(stderr);
    }

    // Destroy all our objects to avoid report of memory leak
    jsCtx_Destroy(js);
    natsStatistics_Destroy(stats);
    natsConnection_Destroy(conn);
    natsOptions_Destroy(opts);

    // To silence reports of memory still in used with valgrind
    nats_Close();

    return 0;
}

```

--------------------------------

### Interactive Subject Mapping

Source: https://docs.nats.io/nats-concepts/subject_mapping

Shows how to use the NATS CLI in interactive mode to test subject mappings. This allows for dynamic testing of different subject transformations.

```shell
nats server mapping foo bar
> Enter subjects to test, empty subject terminates.
> 
> ? Subject foo
> bar

> ? Subject test
> Error: no matching transforms available
```

--------------------------------

### Add NATS Account using NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This command adds a new account within an existing operator using the `nsc` tool. It generates the necessary keys and configuration files for the account. This step is crucial for organizing NATS resources and permissions.

```shell
nsc add account --name A
```

--------------------------------

### NATS CLI: Revocations Command Structure

Source: https://docs.nats.io/using-nats/nats-tools/nsc/revocation

Provides the usage and available commands for managing NATS revocations via the 'nsc' CLI. This includes adding, deleting, and listing user revocations and activation revocations for exports. Flags like '--at' for timestamp and global flags for interactivity and private key are also shown.

```bash
Manage revocation for users and activations from an account

Usage:
  nsc revocations [command]

Available Commands:
  add-user          Revoke a user
  add_activation    Revoke an accounts access to an export
  delete-user       Remove a user revocation
  delete_activation Remove an account revocation from an export
  list-users        List users revoked in an account
  list_activations  List account revocations for an export

Flags:
  -h, --help   help for revocations

Global Flags:
  -i, --interactive          ask questions for various settings
  -K, --private-key string   private key

Use "nsc revocations [command] --help" for more information about a command.
```

--------------------------------

### Enable and Verify JetStream

Source: https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

This snippet demonstrates how to enable JetStream on a NATS server and verify its status. JetStream must be enabled to use Object Stores. The `nats account info` command checks for JetStream information.

```shell
nats-server -js
nats account info
```

--------------------------------

### Create Account with nsc

Source: https://docs.nats.io/using-nats/nats-tools/nsc/services

Creates a new NATS account using the `nsc` command-line tool. This is a prerequisite for importing services from other accounts.

```bash
nsc add account B
```

--------------------------------

### Add Consumer with JSON Config (CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/consumers

Shows how to add a NATS consumer by providing its configuration in a JSON file. This is an alternative to specifying all parameters via CLI flags, allowing for more complex or repeatable configurations.

```shell
nats con add ORDERS --config monitor.json
```

--------------------------------

### Add a NATS User to an Account

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This command adds a new user named 'foo' to the 'TEST' account. It generates user credentials, stores them in a file, and associates the user with the specified account. The output confirms the user creation and the path to the credentials file.

```shell
nsc add user -a TEST -n foo
```

--------------------------------

### Publish to NATS JetStream Stream (Python)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/publish

This Python code snippet shows how to connect to NATS using asyncio, create a JetStream context, add a stream, and publish messages. It demonstrates asynchronous publishing and includes basic error handling for connection and publishing.

```python
import asyncio

import nats
from nats.errors import TimeoutError


async def main():
    nc = await nats.connect("localhost")

    # Create JetStream context.
    js = nc.jetstream()

    # Persist messages on 'example-subject'.
    await js.add_stream(name="example-stream", subjects=["example-subject"])

    for i in range(0, 10):
        ack = await js.publish("example-subject", f"hello world: {i}".encode())
        print(ack)

    await nc.close()

if __name__ == '__main__':
    asyncio.run(main())
```

--------------------------------

### Publish Message to NATS Server

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

This command publishes a simple message ('bar') to the subject 'foo' using specified user credentials. This action can trigger server events that are then observable by subscribers, as demonstrated in the tutorial.

```shell
nats pub --creds ~/.nkeys/SAOP/accounts/SYS/users/SYSU.creds foo bar
```

--------------------------------

### Add NATS User using NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This command creates a new user associated with a specific NATS account using the `nsc` tool. It generates user credentials, including private keys and a .creds file, which are necessary for users to connect to the NATS server. This is fundamental for user authentication and authorization.

```shell
nsc add user --name TA
```

--------------------------------

### List All NATS Keys and Entities

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

This command displays all generated keys and associated entities (operators, accounts, users) managed by nsc. It provides a summary of key information like entity name, key identifier, signing key status, and storage.

```shell
nsc list keys --all

```

--------------------------------

### Send Request to Echo Service using NATS CLI

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-services

This command sends a request to the 'svc.echo' subject of the EchoService. It demonstrates how to interact with a deployed service and measure the round-trip time (RTT).

```bash
nats req svc.echo 'hey'
19:40:22 Sending request on "svc.echo"
19:40:22 Received with rtt 446.27µs
hey
```

--------------------------------

### Authenticate NATS Client with Credentials File (C#)

Source: https://docs.nats.io/using-nats/developer/connecting/creds

Initializes a NATS client connection using a credentials file. The `credsFile` parameter in the `NatsClient` constructor specifies the path to the credentials file for authentication.

```csharp
// dotnet add package NATS.Net
using NATS.Net;

await using var client = new NatsClient("127.0.0.1", credsFile: "/path/to/file.creds");
```

--------------------------------

### Publish and Subscribe with NATS in Go

Source: https://docs.nats.io/using-nats/developer/receiving/wildcards

This Go program demonstrates publishing messages to NATS subjects with time zone information and subscribing to receive them. It uses the `nats-client` Go library and handles time zone conversions.

```go
package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
defer nc.Close()

	// Publish messages with time zone information for Eastern US
	zoneID, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	zoneDateTime := now.In(zoneID)
	formatted := zoneDateTime.String()

	nc.Publish("time.us.east", []byte(formatted))
	nc.Publish("time.us.east.atlanta", []byte(formatted))

	// Publish messages with time zone information for Europe
	zoneID, err = time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.Fatal(err)
	}
	zoneDateTime = now.In(zoneID)
	formatted = zoneDateTime.String()

	nc.Publish("time.eu.east", []byte(formatted))
	nc.Publish("time.eu.east.warsaw", []byte(formatted))

	log.Println("Published messages.")

	// Subscribe to receive messages
	sub, err := nc.SubscribeSync("time.>")
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for i := 0; i < 4; i++ {
		msg, err := sub.NextMsg(10 * time.Second) // Wait up to 10 seconds for a message
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			break
		}
		log.Printf("Received message: Subject='%s', Data='%s'", msg.Subject, string(msg.Data))
	}
}

```

--------------------------------

### Trigger WebAssembly Function via NATS

Source: https://docs.nats.io/using-nats/nex/getting-started/deploying-functions

This command sends a request to the 'wasm.echo' subject on the NATS server with the payload 'hello'. The WebAssembly function, deployed to this subject, should respond by concatenating the payload with the subject. The output displays the time of sending and receiving the request, and the received response.

```bash
$ nats req wasm.echo 'hello'
09:45:24 Sending request on "wasm.echo"
09:45:24 Received with rtt 42.867014ms
hellowasm.echo
```

--------------------------------

### Describe NATS Account using NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt/mem_resolver

This command retrieves and displays detailed information about a specific NATS account using the `nsc` tool. It shows account ID, issuer, expiry, and various limits. This is useful for auditing and understanding an account's configuration.

```shell
nsc describe account -W
```

--------------------------------

### Create NATS Operator, Account, and User with NSC

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts/sys_accounts

These shell commands use the 'nsc' tool to create a new NATS operator, followed by an account within that operator, and then a user associated with that account. This process generates private keys and credential files necessary for authentication.

```shell
nsc add operator -n SAOP
nsc add account -n SYS
nsc add user -n SYSU
```

--------------------------------

### NATS.IO CLI for System Event Subscription

Source: https://docs.nats.io/running-a-nats-service/configuration/sys_accounts

This command subscribes to all system events using the SYS account with administrative credentials. It connects to the NATS server running on localhost at the default port.

```bash
nats sub -s nats://admin:changeit@localhost:4222 ">"
```

--------------------------------

### Manage KV Buckets (Python)

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

This Python code outlines the methods for managing NATS Key/Value (KV) buckets. It defines functions for creating a new KV store, binding to an existing one, and destroying a specified bucket. These operations are essential for programmatic control over KV data storage within a Python application using the NATS client library.

```python
from nats.aio.client import Client as NATS
from nats.js.kv import KeyValue

async def create_kv_bucket(nc: NATS, bucket_name: str) -> KeyValue:
    """Creates a new Key/Value bucket."""
    # ... implementation details ...
    return None

async def bind_to_kv_bucket(nc: NATS, bucket_name: str) -> KeyValue:
    """Binds to an existing Key/Value bucket."""
    # ... implementation details ...
    return None

async def destroy_kv_bucket(kv: KeyValue) -> bool:
    """Destroys a Key/Value bucket."""
    # ... implementation details ...
    return False

```

--------------------------------

### Import System Account JWT (nsc)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Imports the system account JWT from a specified file into the current environment. This command establishes the 'SYS' account's presence and configuration.

```shell
nsc import account --file SYS.jwt
```

--------------------------------

### Echo Function in JavaScript for Nex

Source: https://docs.nats.io/using-nats/nex/getting-started/building-function

A simple JavaScript function for Nex that echoes the incoming payload. It accepts a subject and payload, logs the subject, and returns the payload. This function demonstrates basic input/output handling in Nex JavaScript functions.

```javascript
(subject, payload) => {
  console.log(subject);
  return payload;
};
```

--------------------------------

### Add NATS Servers to a Cluster with Docker

Source: https://docs.nats.io/running-a-nats-service/nats_docker

These commands add additional NATS servers to an existing cluster. Each server is given a unique name, connected to the 'nats' network, and configured to route to the seed server using provided credentials. The `--routes` parameter is crucial for cluster formation.

```bash
docker run --name nats-1 --network nats --rm nats --cluster_name NATS --cluster nats://0.0.0.0:6222 --routes=nats://ruser:T0pS3cr3t@nats:6222
docker run --name nats-2 --network nats --rm nats --cluster_name NATS --cluster nats://0.0.0.0:6222 --routes=nats://ruser:T0pS3cr3t@nats:6222
```

--------------------------------

### Manage Consumers (CLI)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/consumers

Provides essential commands for managing NATS consumers. Includes listing all consumers for a specific stream and querying detailed information about a particular consumer, showing its configuration and state.

```shell
nats con ls ORDERS
nats con info ORDERS DISPATCH
```

--------------------------------

### Create NATS User with Signing Key using nsc

Source: https://docs.nats.io/using-nats/nats-tools/nsc/signing_keys

Creates a new NATS user for an account and signs it using the account's signing key. This command also generates the user's credentials file and adds the user to the account. The command `nsc add user <user_name> -K <account_key_path>` is utilized.

```bash
nsc add user U -K ~/.nkeys/keys/A/DU/ADUQTJD4TF4O6LTTHCKDKSHKGBN2NECCHHMWFREPKNO6MPA7ZETFEEF7.nk
```

--------------------------------

### NATS Client Connection Options

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol

Parameters that can be included in the initial CONNECT message from a client to a NATS server. These options configure client behavior and authentication.

```APIDOC
## Client Connection Options

These fields can be included in the CONNECT message to configure client behavior and authentication.

### Fields

- **jwt** (string) - The JWT that identifies a user permissions and account.
- **no_responders** (bool) - Enable quick replies for cases where a request is sent to a topic with no responders.
- **headers** (bool) - Whether the client supports headers.
- **nkey** (string) - The public NKey to authenticate the client. This will be used to verify the signature (`sig`) against the `nonce` provided in the `INFO` message.

### Example (Go Client)

```
CONNECT {"verbose":false,"pedantic":false,"tls_required":false,"name":"","lang":"go","version":"1.2.2","protocol":1}
```

*Note: Most clients set `verbose` to `false` by default.*
```

--------------------------------

### INFO Protocol Options

Source: https://docs.nats.io/reference/reference-protocols/nats-protocol

The INFO protocol allows clients to retrieve details about the NATS server upon connection. The options are provided as a JSON object.

```APIDOC
## INFO Protocol Options

### Description

The INFO protocol is used by NATS servers to provide clients with essential information about the server's configuration and status. This information is transmitted as a JSON object.

### Method

N/A (This is a protocol message, not a typical HTTP API endpoint)

### Endpoint

N/A

### Parameters

#### Request Body

This section describes the options available within the JSON payload of the INFO message.

- **server_id** (string) - always - The unique identifier of the NATS server.
- **server_name** (string) - always - The name of the NATS server.
- **version** (string) - always - The version of NATS.
- **go** (string) - always - The version of golang the NATS server was built with.
- **host** (string) - always - The IP address used to start the NATS server.
- **port** (int) - always - The port number the NATS server is configured to listen on.
- **headers** (bool) - always - Whether the server supports headers.
- **max_payload** (int) - always - Maximum payload size, in bytes, that the server will accept from the client.
- **proto** (int) - always - An integer indicating the protocol version of the server.
- **client_id** (uint64) - optional - The internal client identifier in the server.
- **auth_required** (bool) - optional - If this is true, then the client should try to authenticate upon connect.
- **tls_required** (bool) - optional - If this is true, then the client must perform the TLS/1.2 handshake.
- **tls_verify** (bool) - optional - If this is true, the client must provide a valid certificate during the TLS handshake.
- **tls_available** (bool) - optional - If this is true, the client can provide a valid certificate during the TLS handshake.
- **connect_urls** ([string]) - optional - List of server urls that a client can connect to.
- **ws_connect_urls** ([string]) - optional - List of server urls that a websocket client can connect to.
- **ldm** (bool) - optional - If the server supports *Lame Duck Mode* notifications, and the current server has transitioned to lame duck, `ldm` will be set to `true`.
- **git_commit** (string) - optional - The git hash at which the NATS server was built.
- **jetstream** (bool) - optional - Whether the server supports JetStream.
- **ip** (string) - optional - The IP of the server.

### Request Example

```json
{
  "server_id": "nats-server-1",
  "server_name": "NATS-Server",
  "version": "2.9.19",
  "go": "go1.21.6",
  "host": "0.0.0.0",
  "port": 4222,
  "headers": true,
  "max_payload": 33554432,
  "proto": 1,
  "client_id": 12345,
  "auth_required": false,
  "tls_required": false,
  "tls_verify": false,
  "tls_available": true,
  "connect_urls": [
    "nats://localhost:4222"
  ],
  "ws_connect_urls": [
    "ws://localhost:8080"
  ],
  "ldm": false,
  "git_commit": "a1b2c3d4e5f6",
  "jetstream": true,
  "ip": "192.168.1.100"
}
```

### Response

#### Success Response (INFO)

The INFO message itself is the response, containing the JSON object detailed in the Parameters section.

- **server_id** (string) - The unique identifier of the NATS server.
- **server_name** (string) - The name of the NATS server.
- **version** (string) - The version of NATS.
- **go** (string) - The version of golang the NATS server was built with.
- **host** (string) - The IP address used to start the NATS server.
- **port** (int) - The port number the NATS server is configured to listen on.
- **headers** (bool) - Whether the server supports headers.
- **max_payload** (int) - Maximum payload size, in bytes, that the server will accept from the client.
- **proto** (int) - An integer indicating the protocol version of the server.
- **client_id** (uint64) - optional - The internal client identifier in the server.
- **auth_required** (bool) - optional - If this is true, then the client should try to authenticate upon connect.
- **tls_required** (bool) - optional - If this is true, then the client must perform the TLS/1.2 handshake.
- **tls_verify** (bool) - optional - If this is true, the client must provide a valid certificate during the TLS handshake.
- **tls_available** (bool) - optional - If this is true, the client can provide a valid certificate during the TLS handshake.
- **connect_urls** ([string]) - optional - List of server urls that a client can connect to.
- **ws_connect_urls** ([string]) - optional - List of server urls that a websocket client can connect to.
- **ldm** (bool) - optional - If the server supports *Lame Duck Mode* notifications, and the current server has transitioned to lame duck, `ldm` will be set to `true`.
- **git_commit** (string) - optional - The git hash at which the NATS server was built.
- **jetstream** (bool) - optional - Whether the server supports JetStream.
- **ip** (string) - optional - The IP of the server.

#### Response Example

```json
{
  "server_id": "nats-server-1",
  "server_name": "NATS-Server",
  "version": "2.9.19",
  "go": "go1.21.6",
  "host": "0.0.0.0",
  "port": 4222,
  "headers": true,
  "max_payload": 33554432,
  "proto": 1,
  "client_id": 12345,
  "auth_required": false,
  "tls_required": false,
  "tls_verify": false,
  "tls_available": true,
  "connect_urls": [
    "nats://localhost:4222"
  ],
  "ws_connect_urls": [
    "ws://localhost:8080"
  ],
  "ldm": false,
  "git_commit": "a1b2c3d4e5f6",
  "jetstream": true,
  "ip": "192.168.1.100"
}
```
```

--------------------------------

### View NSC Environment Configuration

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

This command displays the current NSC environment settings, including the default operator, account, and effective values of relevant environment variables. It helps users understand their current context within the NATS Service Connect system.

```shell
nsc env
```

--------------------------------

### List NATS Entity Keys (Bash)

Source: https://docs.nats.io/using-nats/nats-tools/nsc/basics

Lists the current NATS entities (operators, accounts, users) and their public keys. It indicates whether the keys are stored locally. The output is a formatted table.

```bash
nsc list keys
```

--------------------------------

### Watch Key/Value Store Updates in Go, Java, JavaScript, C#, C

Source: https://docs.nats.io/using-nats/developer/develop_jetstream/kv

Subscribes to updates for keys in the Key/Value store. Supports watching specific keys, multiple keys, or all keys within a bucket. Callbacks are invoked upon changes.

```Go
// Watch for any updates to keys that match the keys argument which could include wildcards.
// Watch will send a nil entry when it has received all initial values.
Watch(keys string, opts ...WatchOpt) (KeyWatcher, error)
// WatchAll will invoke the callback for all updates.
WatchAll(opts ...WatchOpt) (KeyWatcher, error)
```

```Java
/**
 * Watch updates for a specific key
 */
NatsKeyValueWatchSubscription watch(String key, KeyValueWatcher watcher, KeyValueWatchOption... watchOptions) throws IOException, JetStreamApiException, InterruptedException;

/**
 * Watch updates for all keys
 */
NatsKeyValueWatchSubscription watchAll(KeyValueWatcher watcher, KeyValueWatchOption... watchOptions) throws IOException, JetStreamApiException, InterruptedException;
```

```JavaScript
  async watch(
    opts: {
      key?: string;
      headers_only?: boolean;
      initializedFn?: callbackFn;
    } = {},
  ): Promise<QueuedIterator<KvEntry>>
```

```C#
// dotnet add package NATS.Net

// Start a watcher for specific keys
// Key to watch is subject-based and wildcards may be used
IAsyncEnumerable<NatsKVEntry<T>> WatchAsync<T>(string key, INatsDeserialize<T>? serializer = default, NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);

// Start a watcher for specific keys
// Key to watch are subject-based and wildcards may be used
IAsyncEnumerable<NatsKVEntry<T>> WatchAsync<T>(IEnumerable<string> keys, INatsDeserialize<T>? serializer = default, NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);

// Start a watcher for all the keys in the bucket
IAsyncEnumerable<NatsKVEntry<T>> WatchAsync<T>(INatsDeserialize<T>? serializer = default, NatsKVWatchOpts? opts = default, CancellationToken cancellationToken = default);

//
```

```C
NATS_EXTERN natsStatus 	kvStore_Watch (kvWatcher **new_watcher, kvStore *kv, const char *keys, kvWatchOptions *opts)
 	Returns a watcher for any updates to keys that match the keys argument.
 
NATS_EXTERN natsStatus 	kvStore_WatchAll (kvWatcher **new_watcher, kvStore *kv, kvWatchOptions *opts)
 	Returns a watcher for any updates to any keys of the KeyValue store bucket.
```

--------------------------------

### Download NATS CPU Profile (Shell)

Source: https://docs.nats.io/running-a-nats-service/nats_admin/profiling

This snippet shows how to download a CPU profile from a NATS instance over a specified duration using curl. The profile is saved to 'cpu.prof'. The 'seconds' query parameter controls the sampling duration, allowing for analysis of CPU-intensive operations.

```shell
curl -o cpu.prof http://localhost:65432/debug/pprof/profile?seconds=30

```

--------------------------------

### Request-Reply with NATS.io in Go

Source: https://docs.nats.io/using-nats/developer/sending/replyto

Sends a request to the 'time' subject and listens for a single synchronous response on a unique inbox subject. Handles connection, subscription, publishing, and message retrieval. Requires the 'nats-io/nats.go' library.

```go
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Create a unique subject name for replies.
uniqueReplyTo := nats.NewInbox()

// Listen for a single response
sub, err := nc.SubscribeSync(uniqueReplyTo)
if err != nil {
    log.Fatal(err)
}

// Send the request.
// If processing is synchronous, use Request() which returns the response message.
if err := nc.PublishRequest("time", uniqueReplyTo, nil); err != nil {
    log.Fatal(err)
}

// Read the reply
msg, err := sub.NextMsg(time.Second)
if err != nil {
    log.Fatal(err)
}

// Use the response
log.Printf("Reply: %s", msg.Data)
```

--------------------------------

### Connect to NATS with TLS in Python

Source: https://docs.nats.io/using-nats/developer/connecting/tls

This Python code snippet demonstrates establishing a TLS connection to a NATS server. It utilizes the `ssl` module to create a default context, load CA certificates, and load the client's certificate chain before connecting.

```python
nc = NATS()

ssl_ctx = ssl.create_default_context(purpose=ssl.Purpose.SERVER_AUTH)
ssl_ctx.load_verify_locations('rootCA.pem')
ssl_ctx.load_cert_chain(certfile='client-cert.pem',
                        keyfile='client-key.pem')
await nc.connect(io_loop=loop, tls=ssl_ctx)

await nc.connect(servers=["nats://demo.nats.io:4222"], tls=ssl_ctx)
```

--------------------------------

### NATS CLI: Subscribe and Publish within Accounts

Source: https://docs.nats.io/running-a-nats-service/nats_admin/security/jwt

Demonstrates subscribing to all messages and publishing messages within specific NATS accounts using the NATS CLI. Requires NATS server to be running with account configurations.

```shell
nats -s nats://a:a@localhost:4222 sub ">"
nats -s nats://b:b@localhost:4222 pub "foo" "user b"
nats -s nats://a:a@localhost:4222 pub "foo" "user a"
```
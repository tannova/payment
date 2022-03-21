# Alopay - monorepo - backend

## Install Go

https://golang.org/dl/
Version 1.16

At the time 2021-03-21: 1.16.2 (stable latest)

## Install Protobuf
### Step 1
To build and install the C++ Protocol Buffer runtime and the Protocol Buffer compiler (protoc) execute the following
Go to this (link)[https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.4]
And download the pre-build binary for your OS

Unzip the file (read protoc-xxx/README.md)
```
1. copy protoc-xxx/bin/protoc to your bin directory
2. copy protoc-xxx/includes and put into the same level with bin directory
```

### Step 2
Install Go support for Google's protocol buffers
Go to this (link)[https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.25.0]
And download the pre-build binary for your OS & copy to your bin directory

### Step 3
Install The Go language implementation of gRPC. HTTP/2 based RPC
Go to this (link)[https://github.com/grpc/grpc-go/releases/tag/v1.31.0]
And download the pre-build binary for your OS & copy to your bin directory

### Step 4
Install gRPC for Web Clients
Go to this (link)[https://github.com/grpc/grpc-web/releases/tag/1.2.1]
And download the pre-build binary for your OS & copy to your bin directory

### Step 5
Install 
We're using version (v0.4.1)[https://github.com/envoyproxy/protoc-gen-validate/releases/tag/v0.4.1]
You can checkout exactly this version and run `go install` in the root on this project
```
go install github.com/envoyproxy/protoc-gen-validate
```

### Step 6
Install ent for generate code for storage
Go to this (link)[https://github.com/ent/ent/releases/tag/v0.7.0]
You can checkout the repository to your GOPATH
```
cd cmd/ent
go install
```

### Step 7
Install ent for generate code for storage
Go to this (link)[https://github.com/grpc-ecosystem/grpc-gateway/releases/tag/v2.2.0]
You can checkout the repository to your GOPATH
```
cd protoc-gen-openapiv2
go install
```

### Summary
```bash
protoc --version
# libprotoc 3.12.4

protoc-gen-go --version
# protoc-gen-go v1.25.0
```

## Setup Dev Env on EC2

### SSH to Dev service on STG
```bash
ssh -i ~/.ssh/your-private-key-cmgp ec2-user@ec2-13-229-147-201.ap-southeast-1.compute.amazonaws.com
```

### Checkout source code
```bash
git clone https://<key-name>:<access-key>@gitlab.com/mcuc/monorepo.git
```

### Docker CE Install

```sh
sudo amazon-linux-extras install docker
sudo service docker start
sudo usermod -a -G docker ec2-user
```

Make docker auto-start

```
sudo chkconfig docker on
```

Because you always need it....

```
sudo yum install -y git
```

Reboot to verify it all loads fine on its own.

```
sudo reboot
```

### docker-compose install

Copy the appropriate `docker-compose` binary from GitHub:

```
sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
```

Fix permissions after download:

```
sudo chmod +x /usr/local/bin/docker-compose
```

Verify success:

```
docker-compose version
```

### run services
```
make up-stg
```
or
```
docker-compose -f docker-compose.yml -f docker-compose.stg.yml up --build -d
```

### update from master
```
cd /home/ec2-user/monorepo
git pull
make up-stg
```

### check logs of services
```
docker-compose logs -f
```

## Using Makefile
```
make dep

make test

make coverhtml
```



## Install GitLab runner
https://docs.gitlab.com/runner/install/linux-manually.html

```
curl -LJO https://gitlab-runner-downloads.s3.amazonaws.com/latest/rpm/gitlab-runner_amd64.rpm

```

And then register runner https://docs.gitlab.com/runner/register/index.html#linux

# Proto Guidelines

- Always use `proto3`. [Google Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
- Always use a distinct message for RPC requests and responses.

https://medium.com/@akhaku/protobuf-definition-best-practices-87f281576f31

##  Enums

All `enum` values should have a zero value of `UNSPECIFIED`. This signifies that the value was never filled in.
If a real value was used for the zero value, then that value may be erroneously presented in an object despite not being
filled in.

`UNKNOWN` is also a useful value to be used when the value needed is not present in the enum for some reason, for example when
translating between two systems where a string will become an enumerated value.

```proto
enum Foo {
  UNSPECIFIED = 0;  // Value was not filled in.
  UNKNOWN = 1;      // Value could not be represented by existing enum values.
  MY_VALUE = 2;
  MY_OTHER_VALUE = 3;
}
```

# How to run in local machine
You can create a key for yourself over there https://gitlab.com/-/profile/personal_access_tokens
```
BOT_USER=<key-name> BOT_PRIVATE_TOKEN=<access-key> CI_PROJECT_NAMESPACE=crypto-game-portal docker-compose -f docker-compose.yml -f docker-compose.db.yml --env-file ./.env.local up --build
```
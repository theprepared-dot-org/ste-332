# ste332
[![Build Status](https://cloud.drone.io/api/badges/theprepared-dot-org/ste332/status.svg)](https://cloud.drone.io/theprepared-dot-org/ste332)

This repo houses tools for the prepared.org's workshop (Suite 332). The two primary use cases are the Shop Manager
application and the wiki page hosted on Github Pages.

## Tests
To run the unit tests for this repo, execute the following:
```bash
make test
```

To run unit and integration tests, execute the following:
```bash
make test-full
```

## Shop Manager
To run the shop manager locally, execute the following:
```bash
make run
```

### API Interaction
Below is an example that lists users using [grpcurl](https://github.com/fullstorydev/grpcurl):
```bash
grpcurl -plaintext -import-path ./api/shopmanager/ -proto shop_manager.proto localhost:50051 shopmanager.ShopManager.ListUsers
```

### Generate gRPC Definitions
In order to regenerate the gRPC definitions, you will need to have 
[protoc](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1) installed. From there, the following command
will take care of the rest:
```bash
make rpc
```

## Wiki
To run the wiki locally, run the following:
```bash
make serve-wiki
```

### Updates
The wiki is built and published automatically on master builds. To build the wiki for production locally, run the
following:
```bash
make wiki
```

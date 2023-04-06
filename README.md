# PGN
PGN stand for `Playground_N` because I keep debuging xds server in so many environment that I forget which attempt it is. 

From some reason, it happens to work at specifically playground n not n+1.
# Requirement
- [just](https://github.com/casey/just) (Optional: scripting tool)
- kubectl
- [k3d](https://github.com/k3d-io/k3d)


# Start
## Create cluster with k3d
```
just k3d create
```

## Clone [xds server](https://github.com/wong/xds)
```
just xds clone
```

## Build and deploy
```
just build
just deploy
```

# Reference
## Example
- [asishrs/proxyless-grpc-lb](https://github.com/asishrs/proxyless-grpc-lb)
- [salrashid123/grpc_xds](https://github.com/salrashid123/grpc_xds)
## API Reference
- [envoyproxy/go-control-plane](https://github.com/envoyproxy/go-control-plane)
- [xDS REST and gRPC protocol](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol)
## xDS Server
- [wongnai/xds](https://github.com/wongnai/xds)

# Issue
- Hard to debug
- xDS Server fetch interval
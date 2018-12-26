# abuse-mesh-go-stubs
This repository contains the go code grpc stubs for abuse mesh

## Generate the stubs

1. (first time only) clone the abuse mesh protocol repo `git clone https://github.com/abuse-mesh/abuse-mesh-protocol.git vendor/abuse-mesh-protocol`
2. checkout and pull the desired version `git -C vendor/abuse-mesh-protocol checkout {{branch}} && git -C vendor/abuse-mesh-protocol pull`
3. generate the go stubs from the protocol files `docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I vendor/abuse-mesh-protocol abuse-mesh-common.proto abuse-mesh.proto --go_out=plugins=grpc:abusemesh`
4. if the protocol version has changed the version constant in 'vendor/abusemesh/abuse-mesh-version.go' should be updated by hand
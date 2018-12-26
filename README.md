# abuse-mesh-go-stubs
This repository contains the go code grpc stubs for abuse mesh

NOTE at the moment the repository also includes the .proto file, which should be moved into a dedicated repo when the time is right

## Generate the stubs
`docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I . abuse-mesh.proto --go_out=plugins=grpc:abusemesh`
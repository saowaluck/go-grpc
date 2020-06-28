#!bin/bash

# Generate Proto3
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
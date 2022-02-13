#!/bin/bash

cd pkg/protocol/proto
protoc --go_out=plugins=grpc,paths=source_relative:../../pb *.proto

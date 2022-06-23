# Golang Stuff
GOCMD=go

GORUN=$(GOCMD) run

ENV=local

GOPRIVATE=github.com/leonardochristofer/*

SERVICE=article-service

# Make for Unix
init:
	$(GOCMD) mod init $(SERVICE)
	
reinit:
	rm go.mod go.sum && $(GOCMD) mod init $(SERVICE)

tidy:
	ENV=local GOPRIVATE=$(GOPRIVATE) $(GOCMD) mod tidy

run:
	ENV=$(ENV) $(GORUN) main.go

run-grpc:
	ENV=$(ENV) $(GORUN) main.go --mode grpc

# Make for Windows
inits:
	$(GOCMD) mod init $(SERVICE)
	
reinits:
	del go.mod go.sum && $(GOCMD) mod init $(SERVICE)

tidys:
	$(GOCMD) mod tidy

runs-grpc:
	$(GORUN) main.go --mode grpc

# Proto Generator
proto-gen:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional

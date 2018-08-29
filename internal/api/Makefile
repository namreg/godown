PWD?=$(shell pwd)

gen:
	@docker run --rm -v ${PWD}:${PWD} -w ${PWD} znly/protoc --gogofast_out=plugins=grpc:. -I. api.proto
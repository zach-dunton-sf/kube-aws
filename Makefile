all: kube-aws

kube-aws: codegen
	go build ./cmd/kube-aws

test: codegen
	go test $$(go list ./... | grep -v /vendor)

gen: $(wildcard _codegen/*.go _codegen/*/*.go)
	go build -o $@ ./_codegen

codegen: gen
	cd types/ec2 && ../../gen
	cd types && ../gen


.PHONY: all kube-aws codegen
.DELETE_ON_ERROR:


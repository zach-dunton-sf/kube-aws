
all: kube-aws

kube-aws:
	go build ./cmd/kube-aws

test:
	go test $$(go list ./... | grep -v /vendor)

.PHONY: all kube-aws


all: kube-aws

kube-aws:
	go build ./cmd/kube-aws


.PHONY: all kube-aws

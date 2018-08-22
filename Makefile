.PHONY: test

build: graph.go node.go edge.go
	go build github.com/nosarthur/graph

test:
	go test

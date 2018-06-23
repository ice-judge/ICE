deps-go: 
	cd ice && dep ensure --vendor-only

deps-js:
	cd ./ice/public && npm install 

build-judger:
	cd ./ice/judger && mkdir out && make release

build-scheduler:
	go build ./ice/scheduler

build-api:
	go build ./ice/api

build-public:
	cd ./ice/public && npm build

test-go:
	go vet ./ice/...
	go test -cover -v ./ice/...

deps: deps-go deps-js

build: build-scheduler build-api build-public build-judger

test: test-go

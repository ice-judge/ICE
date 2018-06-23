deps-go: 
	cd ice && dep ensure --vendor-only

deps-js:
	cd ./ice/public && npm install 

build-judger:
	cd ./ice/judger && mkdir out && make release

build-scheduler:
	cd ice/scheduler && go build

build-api:
	cd ice/api && go build

build-public:
	cd ./ice/public && npm build

test-go:
	go vet ./ice/...
	go test -cover -v ./ice/...

deps: deps-go deps-js

build: build-scheduler build-api build-public build-judger

test: test-go

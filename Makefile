deps: deps-go deps-js

deps-go: 
	cd ./ice && dep ensure --vendor-only

deps-js:
	cd ./ice/public && npm install 


build: build-scheduler build-api build-public build-judger

build-judger:
	cd ./ice/judger && mkdir out && make release

build-scheduler:
	cd ./ice/scheduler && go build

build-api:
	cd ./ice/api && go build

build-public:
	cd ./ice/public && npm build


test: test-go

test-go:
	cd ./ice && go vet ./...
	cd ./ice && go test -cover -v ./...

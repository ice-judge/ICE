#!/bin/bash
set -e
set -x

ROOTDIR=$1
BINDIR=../bin

PKG_C=$(./check-changes.sh pkg)
API_C=$(./check-changes.sh api)
SCH_C=$(./check-changes.sh scheduler)

if [ $PKG_C = "no" ] && [ $API_C = "no" ] && [ $SCH_C = "no" ]; then
	echo "no changes skipping"
	# pretend successful op
	# so that drone ci won't stop building
	exit 0
fi

# we can use makefile here
gotest(){
	go vet ../$1/...
	go test -cover ../$1/...
}

gobuild(){
	GOOS=linux GOARCH=amd64 go build -o $BINDIR/$1-amd64 ../$1
	GOOS=linux GOARCH=arm go build -o $BINDIR/$1-arm ../$1
	GOOS=linux GOARCH=386 go build -o $BINDIR/$1-i386 ../$1
}

(cd .. && dep ensure --vendor-only)

# since bot api and scheduler use pkg, they must be rebuilt
if [ $PKG_C = "yes" ]; then 
	gotest pkg
	API_C="yes"
	SCH_C="yes"
fi

if [ $API_C = "yes" ]; then
	gobuild api
	gotest api
	if [$CI = "drone"]; then
		mv ../api.Dockerfile ../ci.api.Dockerfile
	fi
fi

if [ $SCH_C = "yes" ]; then
	gobuild scheduler
	gotest scheduler
	if [ $CI = "drone" ]; then 
		mv $ROOTDIR/scheduler.Dockerfile $ROOTDIR/ci.scheduler.Dockerfile
	fi
fi

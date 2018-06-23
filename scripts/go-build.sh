#!/bin/bash
set -e
set -x

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
	go vet ../ice/$1/...
	go test -cover ../ice/$1/...
}

gobuild(){
	go build ../ice/$1
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
		mv ../scheduler.Dockerfile ../ci.scheduler.Dockerfile
	fi
fi

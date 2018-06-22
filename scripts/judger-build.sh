#!/bin/bash
set -e
set -x

if [ $(./check-changes.sh judger) = "no" ]; then
	echo "no changes skipping"
	# pretend successful op
	# so that drone ci won't stop building
	exit 0
fi

(cd ../judger && make release)
if [$CI = "dorne"]; then 
	mv ../judger.Dockerfile ../ci.judger.Dockerfile
fi

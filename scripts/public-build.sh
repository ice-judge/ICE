#!/bin/bash
set -e
set -x

if [ $(./check-changes.sh scheduler) = "no" ];then 
	echo "no changes skipping"
	# to pretend successful op
	# so that drone ci won't stop building
	exit 0
fi

(cd .. && npm install && npm build)

if [ $CI = "drone" ]; then
	mv ../public.Dockerfile ../ci.public.Dockerfile
fi

#!/bin/bash
set -e
set -x

if [ $(./check-changes.sh public) = "no" ];then 
	echo "no changes skipping"
	# to pretend successful op
	# so that drone ci won't stop building
	exit 0
fi

(cd ../ice/public && npm install && npm build)

mv ../public.Dockerfile ../ci.public.Dockerfile

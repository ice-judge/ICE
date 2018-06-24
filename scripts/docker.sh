#!/bin/bash
HASH=$(git log -n 1 --pretty=format:'%H' | cut -c1-8)
BRANCH=$(git branch | grep \* | cut -d ' ' -f2)
repos=( judger scheduler web )

push() {
	for repo in "${repos[@]}"
	do
		echo -e "\n\n"
		echo "%%%%%%%%%pushing $repo:$1"
		echo -e "\n\n"
		docker tag icejudge/$repo:$1 icejudge/$repo
		docker push icejudge/$repo:$1
	done
}

if [[ $1 == "build" ]]; then
	for repo in "${repos[@]}"
	do
		echo -e "\n\n"
		echo "%%%%%%%%%building $repo"
		echo -e "\n\n"
		docker build -t icejudge/$repo -f $repo.Dockerfile .
	done
fi

if [[ $1 == "push" ]]; then 
	if [[ $BRANCH == "master" ]]; then
		push latest
		push "$HASH"
	fi
fi

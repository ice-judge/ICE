#!/bin/bash
HASH=$(git log -n 1 --pretty=format:'%H' | cut -c1-8)
BRANCH=$(git branch | grep \* | cut -d ' ' -f2)
repos=( judger scheduler web )

push() {
	for repo in "${repos[@]}"
	do
		echo "%%%%%%%%%pushing $repo:$1"
		echo -e "\n\n\n\n"
		docker tag icejudge/$repo:$1 icejudge/$repo
		docker push icejudge/$repo:$1
	done
}

if [[ $1 == "build" ]]; then
	for repo in "${repos[@]}"
	do
		echo "%%%%%%%%%building $repo"
		echo -e "\n\n\n\n"
		docker build -t icejudge/$repo -f $repo.Dockerfile .
	done
fi

if [[ $1 == "push" ]]; then 
	if [[ $BRANCH == "master" ]]; then
		push latest
	fi
	push "$BRANCH-$HASH"
fi

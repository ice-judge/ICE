#!/usr/bin/env groovy

String orgPath = "src/github.com/ice-judge"
String repoPath = "${orgPath}/ICE"

pipeline {
	agent {
		docker {
			image "icejudge/build-agent"
			customWorkspace "ice-judge"
		}
	}
	
	environment {
		GOPATH="${WORKSPACE}/go"
		npm_config_cache="npm-cache"
	}

	stages {
		stage("Dependencys") {
			steps {
				sh "mkdir -p $GOPATH/${orgPath}"
				sh "ln -f -s ${WORKSPACE} $GOPATH/${repoPath}"

				sh "cd $GOPATH/${repoPath} && make deps"
			}
		}

		stage("Build") {
			steps {
				sh "cd $GOPATH/{$repoPath} && make build"				
			}
		}

		stage("Unit Test") {
			steps {
				sh "cd $GOPATH/${repoPath} && make test"
			}		
		}
	}
}

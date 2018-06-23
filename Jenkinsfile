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
	}

	stages {
		stage("Dependencys") {
			steps {
				sh "chmod 777 $GOPATH"
				sh "mkdir -p $GOPATH/${orgPath}"
				sh "ln -f -s ${WORKSPACE} $GOPATH/${repoPath}"
				sh "cd $GOPATH/${repoPath} ICE && make deps"
			}
		}

		stage("Build") {
			steps {
				sh "make build"				
			}
		}

		stage("Unit Test") {
			steps {
				sh "make test"
			}		
		}
	}
}

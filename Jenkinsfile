#!/usr/bin/env groovy

String goPath = "/go/src/github.com/ice-judge/ICE"

pipeline {
	agent {
		docker {
			image "icejudge/build-agent"
			customWorkspace "ice-judge"
		}
	}
	
	environment {
		GOPATH="${WORKSPACE}"
	}

	stages {
		stage("Dependencys") {
			steps {
				sh "mkdir -p $GOPATH/src/github.com/ice-judge"
				sh "ln -f -s ${WORKSPACE} $GOPATH/src/github.com/ice-judge/ICE"
				sh "cd $GOPATH/src/github.com/ice-judge/ICE && make deps"
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

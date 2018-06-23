#!/usr/bin/env groovy

String goPath = "/go/src/github.com/ice-judge/ICE"

pipeline {
	agent {
		docker {
			image "icejudge/build-agent"
		}
	}

	environment {
		GOPATH = "/go"
	}

	stages {
		stage("Dependencys") {
			steps {
				sh "mkdir -p $GOPATH/src/github.com/ice-judge/ICE"
				sh "ln -sf $WORKSPACE $GOPATH/src/github.com/ice-judge/ICE"
				sh "make deps"
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

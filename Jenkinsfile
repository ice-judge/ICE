#!/usr/bin/env groovy

String goPath = "/go/src/github.com/ice-judge/ICE"

pipeline {
	agent {
		docker {
			image "icejudge/build-agent"
			args 	"-v ${pwd()}:/go/src/github.com/ice-judge/ICE"

		}
	}

	environment {
		GOPATH = "/go"
	}

	stages {
		stage("Dependencys") {
			steps {
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

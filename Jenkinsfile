#!/usr/bin/env groovy

String goPath = "/go/src/github.com/ice-judge/ICE"

pipeline {
	agent {
		docker {
			image "icejudge/build-agent"
			customWorkspace "${env.HOME}/go/src/github.com/ice-judge/ICE"
		}
	}
	
	environment {
		GOPATH = "${env.HOME}/go"
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

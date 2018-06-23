#!/usr/bin/env groovy

String goPath = "/go/src/github.com/ice-judge/ICE"

pipeline {
	agent any

	stages {
		stage("Dependencys") {
			agent {
				docker {
					image "icejudge/build-agent"
					args 	"-v ${PWD}:/go/src/github.com/ice-judge/ICE"
					reuseNode true
				}
			}

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

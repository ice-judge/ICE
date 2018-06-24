#!/usr/bin/env groovy
import hudson.model.*

pipeline {
	agent any

  stages {
    stage("Build") {
			steps {
				sh "scripts/docker.sh build"
			}
		}

		stage("Unit Test") {
			steps {
				sh "mkdir reports"
				sh "docker-compose -f docker/docker-compose-test-go-js.yml run web make test-go 2>&1 | go-junit-report > reports/go.xml"
			}
		}

		stage("Publish to Docker") {
			steps {
				script {
					docker.withRegistry("", "icejudge-docker-credentials") {
						sh "scripts/docker.sh push"
					}
				}
			}
		}
	}

	post {
		always {
			junit "reports/*.xml"
			archiveArtifacts "reports/*.xml"
		}
	}
}

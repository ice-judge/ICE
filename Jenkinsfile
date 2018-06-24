#!/usr/bin/env groovy
import hudson.model.*

def hash = sh (script: "git log -n 1 --pretty=format:'%H' | cut -c1-8", returnStdout: true)
def tag = "${BRANCH_NAME}-${hash}"

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
					withDockerRegistry([credentialsId: 'docker-registry-credentials']) {
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

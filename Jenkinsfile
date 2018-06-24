#!/usr/bin/env groovy
import hudson.model.*

def orgPath = "src/github.com/ice-judge"
def repoPath = "${orgPath}/ICE"

pipeline {
  agent any

  stages {
    stage("Build") {
			steps {
				sh "docker build -t icejudge/web -f scheduler.Dockerfile ."
			}
		}

		stage("Unit Test") {
			steps {
				sh "mkdir reports"
				sh "docker-compose -f docker/docker-compose-test-go-js.yml run web make test-go 2>&1 | go-junit-report > reports/go.xml"
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

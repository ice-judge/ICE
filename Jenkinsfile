#!/usr/bin/env groovy

node('docker') {
	String apiPath = "/go/src/github.com/ice-judge/ICE/api"

	stage('Checkout from GitHub') {
		checkout scm
	}

	stage("Create binaries") {
		docker.image("golang:1.8.0-alpine").inside("-v ${pwd()}:${apiPath}") {
			sh "cd ${apiPath} && GOOS=darwin GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/asdf
		}
  }

  stage("Archive artifacts") {
    archiveArtifacts artifacts: 'binaries/**', fingerprint: true
  }
}

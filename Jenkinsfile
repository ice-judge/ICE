#!/usr/bin/env groovy



def setBuildStatus = { String message, String state ->
  step([
      $class: "GitHubCommitStatusSetter",
      reposSource: [$class: "ManuallyEnteredRepositorySource", url: "https://github.com/ice-judge/ICE"],
      contextSource: [$class: "ManuallyEnteredCommitContextSource", context: "ci/jenkins/build-status"],
      errorHandlers: [[$class: "ChangingBuildStatusErrorHandler", result: "UNSTABLE"]],
      statusResultSource: [ $class: "ConditionalStatusResultSource", results: [[$class: "AnyBuildResult", message: message, state: state]] ]
  ]);
}

String goPath = "/go/src/github.com/ice-judge/ICE"
 
node('docker') {
	setBuildStatus("building", "PENDING")

	stage('checkout from gitHub') {
		checkout scm
	}

	stage("build go") {
		docker.image("golang:strecth").inside("-v ${pwd()}:${goPath}") {
			sh "cd ${goPath}/scripts && ./go-build.sh"
		}
  }

	setBuildStatus("success", "SUCCESS")
}

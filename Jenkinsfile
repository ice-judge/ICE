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
 
pipeline {
	agent any 
	stages {
		stage("checkout") {
			checkout scm
		}

		setBuildStatus("building", "PENDING")
		
		stage("check changes") {

		}	
		
		setBuildStatus("success", "SUCCESS")
	}
}
	

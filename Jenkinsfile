#!/usr/bin/env groovy

String orgPath = "src/github.com/ice-judge"
String repoPath = "${orgPath}/ICE"

node {
	stage('Checkout'){
		checkout scm
	}

	def scheduler	
	def judger
	def web
	
	stage('Build') {
		scheduler = docker.build("image", "./scheduler.Dockerfile")
		web = docker.build("image3", "./web.Dockerfile")
		judger = docker.build("image2", "./judger.Dockerfile")
	}
}

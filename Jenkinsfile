#!/usr/bin/env groovy

String orgPath = "src/github.com/ice-judge"
String repoPath = "${orgPath}/ICE"

node {
	try {
		stage("Checkout"){
			checkout scm
		}

		def scheduler	
		def judger
		def web
		
		stage("Build") {
			scheduler = docker.build("icejudge/scheduler", "-f ./scheduler.Dockerfile .")
			//web = docker.build("icejudge/web", "-f ./web.Dockerfile .")
			//judger = docker.build("icejudge/judger", "-f ./judger.Dockerfile .")
			sh "docker-compose up -d"
		}
		
		stage("Unit Tests") {
			stage "Go Tests"
			sh "docker-compose exec -T scheduler make test-go"
		}
	} catch (e) {
		throw e
	} finally {
		sh "docker-compose down || true"
	}
}	

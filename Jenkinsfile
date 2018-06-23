#!/usr/bin/env groovy

def orgPath = "src/github.com/ice-judge"
def repoPath = "${orgPath}/ICE"

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
			web = docker.build("icejudge/web", "-f ./web.Dockerfile .")
			judger = docker.build("icejudge/judger", "-f ./judger.Dockerfile .")

			sh "docker-compose up -d"
		}
		
		stage("Unit Tests") {
			sh "docker-compose exec -T scheduler make test-go"
		}

		stage("Publish to Docker") {
			withDockerRegistry([ credentialsId: "icejudge-docke-credentials", url: "" ]) {
				def commit = "${GIT_COMMIT:0:8}"
				def tag = "${BRANCH_NAME}-${commit}"

				scheduler.push("${tag}")
				web.push("${tag}")
				judger.push("${tag}")
			}
		}
	} catch (e) {
		throw e
	} finally {
		sh "docker-compose down || true"
	}
}	

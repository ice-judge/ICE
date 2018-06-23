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
				def COMMIT = "${GIT_COMMIT.substring(0,8)}"
				if ("${BRANCH_NAME}" == "master") {
					scheduler.push("latest")	
					web.push("latest")	
					judger.push("latest")	
				}
				scheduler.push("${BRANCH_NAME}-${COMMIT}")
				web.push("${BRANCH_NAME}-${COMMIT}")
				judger.push("${BRANCH_NAME}-${COMMIT}")
			}
		}
	} catch (e) {
		throw e
	} finally {
		sh "docker-compose down || true"
	}
}	

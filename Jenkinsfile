pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS_ID = 'dockerhub-credentials'
        KUBECONFIG_CREDENTIALS_ID = 'kubeconfig-credentials'
        DOCKER_REPO = 'josuejero/selestino'
    }

    triggers {
        githubPush()
    }

    stages {
        stage('Build') {
            steps {
                script {
                    sh 'docker context use default'
                    sh 'docker --version'
                    docker.build("${DOCKER_REPO}:latest")
                }
            }
        }

        stage('Push') {
            steps {
                script {
                    sh 'docker context ls'
                    sh 'docker context use default'
                    withDockerRegistry([url: 'https://index.docker.io/v1/', credentialsId: DOCKER_CREDENTIALS_ID]) {
                        docker.image("${DOCKER_REPO}:latest").push()
                    }
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    docker.image('golang:1.20-alpine').inside {
                        sh 'apk --no-cache add git'
                        sh 'go test -v ./... -coverprofile=coverage.out'
                        sh 'go tool cover -html=coverage.out -o coverage.html'
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    withKubeConfig([credentialsId: KUBECONFIG_CREDENTIALS_ID]) {
                        sh 'kubectl apply -f k8s/elasticsearch-deployment.yaml'
                        sh 'kubectl apply -f k8s/postgres-deployment.yaml'
                        sh 'kubectl apply -f k8s/selestino-deployment.yaml'
                        sh 'kubectl apply -f k8s/redis-deployment.yaml'
                    }
                }
            }
        }

        stage('Publish Coverage Report') {
            steps {
                publishHTML(target: [
                    reportDir: '.',
                    reportFiles: 'coverage.html',
                    reportName: 'Coverage Report'
                ])
            }
        }
    }
}

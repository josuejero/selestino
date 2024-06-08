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
                    docker.image('golang:1.20-alpine').inside('-u root:root') {
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
                        // Install kubectl
                        sh '''
                        apk --no-cache add curl
                        curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
                        chmod +x ./kubectl
                        mv ./kubectl /usr/local/bin/kubectl
                        '''

                        sh 'kubectl version --client'

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

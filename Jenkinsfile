pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS_ID = 'dockerhub-credentials'
        KUBECONFIG_CREDENTIALS_ID = 'kubeconfig-credentials'
        DOCKER_REPO = 'josuejero/selestino'
        KUBECONFIG = "${WORKSPACE}/kubeconfig"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker context use default'
                    sh 'docker --version'
                    echo "Building Docker image..."
                    docker.build("${DOCKER_REPO}:latest")
                    echo "Docker image built successfully"
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh 'docker context ls'
                    sh 'docker context use default'
                    echo "Pushing Docker image to repository..."
                    withDockerRegistry([url: 'https://index.docker.io/v1/', credentialsId: DOCKER_CREDENTIALS_ID]) {
                        docker.image("${DOCKER_REPO}:latest").push()
                    }
                    echo "Docker image pushed successfully"
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    docker.image('golang:1.20-alpine').inside('-u root:root') {
                        sh 'apk --no-cache add git'
                        echo "Running tests..."
                        sh 'go test -v ./... -coverprofile=coverage.out'
                        sh 'go tool cover -html=coverage.out -o coverage.html'
                        echo "Tests completed"
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    withKubeConfig([credentialsId: KUBECONFIG_CREDENTIALS_ID]) {
                        docker.image('alpine:latest').inside('-u root') {
                            echo "Installing kubectl..."
                            sh '''
                            apk --no-cache add curl
                            curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
                            chmod +x ./kubectl
                            mv ./kubectl /usr/local/bin/kubectl
                            '''
                            sh 'kubectl version --client'
                            echo "kubectl installed successfully"

                            echo "Validating kubeconfig path and contents..."
                            sh 'ls -l ${KUBECONFIG}'
                            sh 'cat ${KUBECONFIG}'

                            echo "Applying Kubernetes configurations..."
                            sh 'KUBECONFIG=${KUBECONFIG} kubectl apply -f k8s/elasticsearch-deployment.yaml'
                            sh 'KUBECONFIG=${KUBECONFIG} kubectl apply -f k8s/postgres-deployment.yaml'
                            sh 'KUBECONFIG=${KUBECONFIG} kubectl apply -f k8s/selestino-deployment.yaml'
                            sh 'KUBECONFIG=${KUBECONFIG} kubectl apply -f k8s/redis-deployment.yaml'
                            echo "Kubernetes configurations applied"

                            echo "Listing Kubernetes pods..."
                            sh 'KUBECONFIG=${KUBECONFIG} kubectl get pods -o wide'
                        }
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

    post {
        success {
            echo 'Pipeline completed successfully'
        }
        failure {
            echo 'Pipeline failed'
        }
        always {
            cleanWs()
        }
    }
}

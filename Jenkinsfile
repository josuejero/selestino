pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS_ID = 'dockerhub-credentials'
        KUBECONFIG_CREDENTIALS_ID = 'kubeconfig-credentials'
        DOCKER_REPO = 'josuejero/selestino'
        KUBECONFIG = "${WORKSPACE}/kubeconfig"  
    }

    triggers {
        githubPush()
    }

    stages {
        stage('Build') {
            steps {
                script {
                    echo "Setting Docker context to default..."
                    sh 'docker context use default'
                    sh 'docker --version'
                    echo "Building Docker image..."
                    docker.build("${DOCKER_REPO}:latest")
                    echo "Docker image built successfully"
                }
            }
        }

        stage('Push') {
            steps {
                script {
                    echo "Listing Docker contexts..."
                    sh 'docker context ls'
                    echo "Setting Docker context to default..."
                    sh 'docker context use default'
                    echo "Pushing Docker image to repository..."
                    withDockerRegistry([url: 'https://index.docker.io/v1/', credentialsId: DOCKER_CREDENTIALS_ID]) {
                        docker.image("${DOCKER_REPO}:latest").push()
                    }
                    echo "Docker image pushed successfully"
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    docker.image('golang:1.20-alpine').inside('-u root:root') {
                        echo "Installing Git inside the Docker container..."
                        sh 'apk --no-cache add git'
                        echo "Running tests..."
                        sh 'go test -v ./... -coverprofile=coverage.out'
                        sh 'go tool cover -html=coverage.out -o coverage.html'
                        echo "Tests completed"
                    }
                }
            }
        }

        stage('Prepare Kubeconfig') {
            steps {
                script {
                    echo "Creating .minikube directory in workspace..."
                    sh 'mkdir -p ${WORKSPACE}/.minikube'
                    echo "Copying ca.crt to workspace..."
                    sh 'cp /Users/wholesway/.minikube/ca.crt ${WORKSPACE}/.minikube/ca.crt'
                    echo "Copying client.crt to workspace..."
                    sh 'cp /Users/wholesway/.minikube/profiles/minikube/client.crt ${WORKSPACE}/.minikube/client.crt'
                    echo "Copying client.key to workspace..."
                    sh 'cp /Users/wholesway/.minikube/profiles/minikube/client.key ${WORKSPACE}/.minikube/client.key'

                    echo "Updating paths in kubeconfig file..."
                    sh """
                        sed -i 's|/Users/wholesway/.minikube/ca.crt|${WORKSPACE}/.minikube/ca.crt|g' ${KUBECONFIG}
                        sed -i 's|/Users/wholesway/.minikube/profiles/minikube/client.crt|${WORKSPACE}/.minikube/client.crt|g' ${KUBECONFIG}
                        sed -i 's|/Users/wholesway/.minikube/profiles/minikube/client.key|${WORKSPACE}/.minikube/client.key|g' ${KUBECONFIG}
                    """
                    echo "Updated kubeconfig file:"
                    sh 'cat ${KUBECONFIG}'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    echo "Entering withKubeConfig block..."
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
                echo "Coverage report published"
            }
        }
    }

    post {
        always {
            echo "Pipeline execution completed"
        }
        success {
            echo "Pipeline succeeded"
        }
        failure {
            echo "Pipeline failed"
        }
    }
}

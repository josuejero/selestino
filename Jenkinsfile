pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS_ID = 'dockerhub-credentials'
        KUBECONFIG_CREDENTIALS_ID = 'kubeconfig-credentials'
        DOCKER_REPO = 'josuejero/selestino'
        KUBECONFIG = "/root/.kube/config"
    }

    stages {
        stage('Checkout') {
            steps {
                echo "Checking out SCM..."
                checkout scm
                echo "SCM checkout completed."
            }
        }

        stage('Prepare Environment') {
            steps {
                script {
                    echo "Preparing environment..."
                    // Copy cert files to the workspace
                    sh 'cp /ca.crt $WORKSPACE/ca.crt'
                    sh 'cp /client.crt $WORKSPACE/client.crt'
                    sh 'cp /client.key $WORKSPACE/client.key'
                    echo "Cert files copied to workspace:"
                    sh 'ls -l $WORKSPACE'
                }
            }
        }

        stage('Build Docker Image') {
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

        stage('Push Docker Image') {
            steps {
                script {
                    echo "Listing Docker contexts..."
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
                    echo "Running tests in Docker container..."
                    docker.image('golang:1.20-alpine').inside('-u root:root') {
                        sh 'apk --no-cache add git'
                        echo "Running tests..."
                        sh 'go test -v ./... -coverprofile=coverage.out'
                        sh 'go tool cover -html=coverage.out -o coverage.html'
                        echo "Tests completed"
                        echo "Test output:"
                        sh 'cat coverage.out'
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {
                    withKubeConfig([credentialsId: KUBECONFIG_CREDENTIALS_ID]) {
                        docker.image('alpine:latest').inside('-u root') {
                            echo "Installing Docker and kubectl..."
                            sh '''
                            apk --no-cache add curl docker openrc
                            if ! (service docker status | grep -q 'is running'); then
                                if ! (service docker status | grep -q 'is starting'); then
                                    service docker start
                                    echo "Checking Docker status..."
                                    while ! docker info > /dev/null 2>&1; do
                                        echo "Waiting for Docker to start..."
                                        sleep 1
                                    done
                                else
                                    echo "Docker is already starting"
                                fi
                            else
                                echo "Docker is already running"
                            fi
                            docker --version
                            curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
                            chmod +x ./kubectl
                            mv ./kubectl /usr/local/bin/kubectl
                            kubectl version --client
                            '''
                            echo "Docker and kubectl installed successfully"

                            echo "Creating .kube directory and copying kubeconfig"
                            sh 'mkdir -p /root/.kube'
                            sh 'cp $WORKSPACE/ca.crt /root/.kube/ca.crt'
                            sh 'cp $WORKSPACE/client.crt /root/.kube/client.crt'
                            sh 'cp $WORKSPACE/client.key /root/.kube/client.key'
                            sh 'cp $WORKSPACE/kubeconfig /root/.kube/config'

                            echo "Validating kubeconfig path and contents..."
                            sh 'ls -l /root/.kube/config'
                            sh 'cat /root/.kube/config'

                            echo "Checking for certificate files..."
                            sh 'ls -l /root/.kube/'
                            sh 'cat /root/.kube/client.crt'
                            sh 'cat /root/.kube/client.key'
                            sh 'cat /root/.kube/ca.crt'

                            echo "Validating Minikube status..."
                            sh '''
                            curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-arm64
                            chmod +x minikube
                            mv minikube /usr/local/bin/
                            minikube status || minikube start --driver=docker
                            '''

                            echo "Applying Kubernetes configurations..."
                            sh '''
                            kubectl config set-cluster minikube --certificate-authority=/root/.kube/ca.crt --embed-certs=true
                            kubectl config set-credentials minikube --client-certificate=/root/.kube/client.crt --client-key=/root/.kube/client.key --embed-certs=true
                            kubectl config set-context minikube --cluster=minikube --user=minikube
                            kubectl config use-context minikube
                            kubectl apply -f k8s/elasticsearch-deployment.yaml --validate=false
                            kubectl apply -f k8s/postgres-deployment.yaml --validate=false
                            kubectl apply -f k8s/selestino-deployment.yaml --validate=false
                            kubectl apply -f k8s/redis-deployment.yaml --validate=false
                            '''
                            echo "Kubernetes configurations applied"

                            echo "Listing Kubernetes pods..."
                            sh 'kubectl get pods -o wide'
                        }
                    }
                }
            }
        }

        stage('Publish Coverage Report') {
            steps {
                echo "Publishing coverage report..."
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
            echo "Cleaning up workspace..."
            cleanWs()
        }
    }
}

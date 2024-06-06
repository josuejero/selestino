// Jenkinsfile

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
                    // Ensure Docker context is set to default
                    sh 'docker context use default'
                    
                    // Verify Docker version
                    sh 'docker --version'
                    
                    // Build Docker image
                    docker.build("${DOCKER_REPO}:latest")
                }
            }
        }

        stage('Push') {
            steps {
                script {
                    echo 'Docker context:'
                    sh 'docker context ls'

                    // Explicitly set Docker context to default
                    sh 'docker context use default'

                    // Login to Docker Hub and push the image
                    withDockerRegistry([url: 'https://index.docker.io/v1/', credentialsId: DOCKER_CREDENTIALS_ID]) {
                        docker.image("${DOCKER_REPO}:latest").push()
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Apply Kubernetes configurations using kubectl
                    withKubeConfig([credentialsId: KUBECONFIG_CREDENTIALS_ID]) {
                        sh 'kubectl apply -f k8s/redis-deployment.yaml'
                        sh 'kubectl apply -f k8s/postgres-deployment.yaml'
                        sh 'kubectl apply -f k8s/selestino-deployment.yaml'
                    }
                }
            }
        }
    }
}

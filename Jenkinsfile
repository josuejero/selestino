pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS_ID = 'your-docker-credentials-id'
        KUBECONFIG_CREDENTIALS_ID = 'your-kubeconfig-credentials-id'
    }

    stages {
        stage('Build') {
            steps {
                script {
                    docker.build("selestino:latest")
                }
            }
        }

        stage('Push') {
            steps {
                script {
                    echo 'Docker context:'
                    sh 'docker context ls'

                    docker.withRegistry('https://index.docker.io/v1/', DOCKER_CREDENTIALS_ID) {
                        docker.image("selestino:latest").push()
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                withKubeConfig(credentialsId: KUBECONFIG_CREDENTIALS_ID) {
                    sh 'kubectl apply -f k8s/postgres-deployment.yaml'
                    sh 'kubectl apply -f k8s/selestino-deployment.yaml'
                }
            }
        }
    }
}

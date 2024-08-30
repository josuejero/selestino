pipeline {
    agent any

    environment {
        // Environment variables
        DOCKER_IMAGE = "selestino-web"
        REGISTRY_CREDENTIALS = credentials('docker')
        GITHUB_CREDENTIALS = credentials('github')
        GOOGLE_CLOUD_CREDENTIALS = credentials('GOOGLE_CLOUD_CREDENTIALS')
        DB_NAME = "selestino"
        DB_USER = "josuejero"
        DB_PASSWORD = credentials('DB_PASSWORD')
        GOOGLE_PROJECT_ID = "selestino-434015"
        GOOGLE_COMPUTE_ZONE = "us-central1-a"
    }

    stages {
        stage('Check Environment') {
            steps {
                script {
                    echo "User: $(whoami) [DEBUG-001]"
                    echo "Current directory: $(pwd) [DEBUG-002]"
                    echo "PATH: $PATH [DEBUG-003]"
                    echo "Checking if Docker is in PATH: $(which docker || echo 'Docker not found in PATH') [DEBUG-004]"
                    sh 'ls -la /var/run/docker.sock [DEBUG-005]'
                    sh 'ls -la /usr/bin/docker || echo "/usr/bin/docker not found" [DEBUG-006]'
                    sh 'echo $SHELL [DEBUG-007]'
                }
            }
        }

        stage('Check Docker Installation') {
            steps {
                script {
                    try {
                        sh 'docker --version'
                        echo "Docker is installed and accessible. [DEBUG-008]"
                    } catch (Exception e) {
                        echo "Docker is not installed or not accessible: ${e.message} [ERROR-108]"
                        error("Failed at stage: Check Docker Installation [ERROR-108]")
                    }
                }
            }
        }

        stage('Verify Docker Installation') {
            steps {
                script {
                    try {
                        sh 'sudo docker --version'
                        echo "Docker is installed and accessible with sudo. [DEBUG-009]"
                    } catch (Exception e) {
                        echo "Docker is not accessible with sudo: ${e.message} [ERROR-109]"
                        error("Failed at stage: Verify Docker Installation [ERROR-109]")
                    }
                }
            }
        }

        stage('Checkout Code') {
            steps {
                echo "Starting code checkout from GitHub... [DEBUG-010]"
                script {
                    try {
                        git branch: 'master', credentialsId: "${GITHUB_CREDENTIALS}", url: 'https://github.com/josuejero/selestino.git'
                        echo "Code checkout completed successfully. [DEBUG-011]"
                    } catch (Exception e) {
                        echo "Error during code checkout: ${e.message} [ERROR-101]"
                        error("Failed at stage: Checkout Code [ERROR-101]")
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Starting Docker image build... [DEBUG-012]"
                script {
                    try {
                        dockerImage = docker.build("${DOCKER_IMAGE}:${env.BUILD_ID}")
                        echo "Docker image built successfully: ${DOCKER_IMAGE}:${env.BUILD_ID} [DEBUG-013]"
                    } catch (Exception e) {
                        echo "Error during Docker image build: ${e.message} [ERROR-102]"
                        error("Failed at stage: Build Docker Image [ERROR-102]")
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo "Starting tests... [DEBUG-014]"
                script {
                    try {
                        dockerImage.inside('-v /var/run/docker.sock:/var/run/docker.sock') {
                            sh 'pytest selestino/tests/'
                        }
                        echo "Tests completed successfully. [DEBUG-015]"
                    } catch (Exception e) {
                        echo "Error during testing: ${e.message} [ERROR-103]"
                        error("Failed at stage: Run Tests [ERROR-103]")
                    }
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                echo "Starting Docker image push to registry... [DEBUG-016]"
                script {
                    try {
                        docker.withRegistry('', "${REGISTRY_CREDENTIALS}") {
                            dockerImage.push("${env.BUILD_ID}")
                            dockerImage.push('latest')
                        }
                        echo "Docker image pushed successfully: ${DOCKER_IMAGE}:${env.BUILD_ID} [DEBUG-017]"
                    } catch (Exception e) {
                        echo "Error during Docker image push: ${e.message} [ERROR-104]"
                        error("Failed at stage: Push Docker Image [ERROR-104]")
                    }
                }
            }
        }

        stage('Deploy to Google Cloud') {
            steps {
                echo "Starting deployment to Google Cloud... [DEBUG-018]"
                script {
                    try {
                        sh """
                            gcloud auth activate-service-account --key-file=${GOOGLE_CLOUD_CREDENTIALS}
                            gcloud config set project ${GOOGLE_PROJECT_ID}
                            gcloud config set compute/zone ${GOOGLE_COMPUTE_ZONE}
                            gcloud run deploy selestino-app \
                                --image gcr.io/${GOOGLE_PROJECT_ID}/${DOCKER_IMAGE}:${env.BUILD_ID} \
                                --platform managed \
                                --region us-central1 \
                                --allow-unauthenticated
                        """
                        echo "Deployment to Google Cloud completed successfully. [DEBUG-019]"
                    } catch (Exception e) {
                        echo "Error during deployment to Google Cloud: ${e.message} [ERROR-105]"
                        error("Failed at stage: Deploy to Google Cloud [ERROR-105]")
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Build and Deployment were successful! [DEBUG-020]'
        }
        failure {
            echo 'Build or Deployment failed. Please check the logs. [ERROR-106]'
        }
    }
}

pipeline {
    agent any

    environment {
        PATH = "/usr/bin:$PATH"
        DB_NAME = "selestino"
        DB_USER = "josuejero"
        GOOGLE_PROJECT_ID = "selestino-434015"
        GOOGLE_COMPUTE_ZONE = "us-central1-a"
        GITHUB_CREDENTIALS = credentials('GITHUB_CREDENTIALS_ID')  // Add GitHub credentials to the environment
    }

    stages {
        stage('Check Environment Variables') {
            steps {
                script {
                    echo "Checking environment variables... [DEBUG-000]"
                    sh 'echo DB_NAME: $DB_NAME'
                    sh 'echo DB_USER: $DB_USER'
                    sh 'echo GOOGLE_PROJECT_ID: $GOOGLE_PROJECT_ID'
                    sh 'echo GOOGLE_COMPUTE_ZONE: $GOOGLE_COMPUTE_ZONE'
                    sh 'echo GITHUB_CREDENTIALS: ****'  // Masked echo of the GitHub credentials
                }
            }
        }

        stage('Check Environment') {
            steps {
                script {
                    sh 'echo "User: $(whoami)" [DEBUG-001]'
                    sh 'echo "Current directory: $(pwd)" [DEBUG-002]'
                    sh 'echo "PATH: $PATH" [DEBUG-003]'
                }
            }
        }

        stage('Checkout Code') {
            steps {
                script {
                    echo "Starting code checkout... [DEBUG-010]"
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

        // Skipped Docker build and related stages

        stage('Run Tests') {
            steps {
                echo "Skipping test stage... [DEBUG-014]"
            }
        }

        stage('Deploy to Google Cloud') {
            steps {
                echo "Skipping deployment to Google Cloud... [DEBUG-018]"
            }
        }
    }

    post {
        success {
            echo 'Build was successful! [DEBUG-020]'
        }
        failure {
            echo 'Build failed. Please check the logs. [ERROR-106]'
        }
    }
}

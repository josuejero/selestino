pipeline {
    agent any

    environment {
        PATH = "/usr/bin:$PATH"
        DB_NAME = "selestino"
        DB_USER = "josuejero"
        GOOGLE_PROJECT_ID = "selestino-434015"
        GOOGLE_COMPUTE_ZONE = "us-central1-a"
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
                echo "Skipping code checkout... [DEBUG-010]"
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

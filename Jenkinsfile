pipeline {
    agent {
        docker {
            image 'python:3.8'
            args '-u root'
        }
    }

    environment {
        PATH = "/usr/local/bin:$PATH"
        GITHUB_CREDENTIALS = credentials('GITHUB_CREDENTIALS_ID')
        DB_CREDENTIALS = credentials('DB_CREDENTIALS_ID')
    }

    stages {
        stage('Load Environment Variables') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'ENV_FILE_CREDENTIALS', variable: 'ENV_FILE')]) {
                        sh 'export $(cat $ENV_FILE | xargs)'
                    }
                }
            }
        }

        stage('Check Environment') {
            steps {
                script {
                    sh 'echo "User: $(whoami)"'
                    sh 'echo "Current directory: $(pwd)"'
                }
            }
        }

        stage('Install Python 3 and Setup Virtual Environment') {
            steps {
                script {
                    sh '''
                        if ! command -v python3.8 &> /dev/null
                        then
                            sudo apt update -y
                            sudo apt install -y python3.8 python3-pip python3-venv
                        fi
                        python3.8 -m venv venv
                        . venv/bin/activate
                    '''
                }
            }
        }

        stage('Install Python Dependencies') {
            steps {
                script {
                    sh '''
                        . venv/bin/activate
                        pip install --cache-dir $WORKSPACE/.pip_cache --upgrade pip
                        pip install -r requirements.txt
                    '''
                }
            }
        }

        stage('Checkout Code') {
            steps {
                script {
                    retry(3) {
                        git branch: 'master', credentialsId: "${GITHUB_CREDENTIALS}", url: 'https://github.com/josuejero/selestino.git'
                    }
                }
            }
        }

        stage('Start PostgreSQL Service and Check Database') {
            steps {
                script {
                    try {
                        sh '''
                            service postgresql start || echo "PostgreSQL is already running"
                            DB_NAME=$(echo $DB_NAME | sed 's/[^a-zA-Z0-9_]//g')
                            DB_USER=$(echo $DB_USER | sed 's/[^a-zA-Z0-9_]//g')
                            psql -h $DB_HOST -U postgres -c "CREATE DATABASE $DB_NAME;" 2>/dev/null || echo "Database already exists"
                            psql -h $DB_HOST -U postgres -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';" 2>/dev/null || echo "User already exists"
                            psql -h $DB_HOST -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
                        '''
                    } catch (Exception e) {
                        error("Failed at stage: Start PostgreSQL Service")
                    }
                }
            }
        }

        stage('Verify Database Connectivity') {
            steps {
                script {
                    try {
                        sh '''
                            PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT 1;"
                        '''
                    } catch (Exception e) {
                        error("Failed at stage: Verify Database Connectivity")
                    }
                }
            }
        }

        stage('Apply Migrations') {
            steps {
                script {
                    try {
                        sh '''
                            . venv/bin/activate
                            python selestino/manage.py migrate
                        '''
                    } catch (Exception e) {
                        error("Failed at stage: Apply Migrations")
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    try {
                        if (fileExists('tests')) {
                            sh '''
                                . venv/bin/activate
                                pytest --cov=selestino tests/
                            '''
                        } else if (fileExists('selestino/tests')) {
                            dir('selestino') {
                                sh '''
                                    . ../venv/bin/activate
                                    pytest --cov=selestino tests/
                                '''
                            }
                        } else {
                            error("Tests directory not found")
                        }
                    } catch (Exception e) {
                        error("Failed at stage: Run Tests")
                    }
                }
            }
        }

        stage('Build and Deploy to Kubernetes') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'GCLOUD_CREDENTIALS', variable: 'GCLOUD_KEYFILE'),
                                    string(credentialsId: 'GCLOUD_PROJECT_ID', variable: 'PROJECT_ID'),
                                    string(credentialsId: 'CLUSTER_NAME', variable: 'CLUSTER_NAME'),
                                    string(credentialsId: 'CLUSTER_ZONE', variable: 'CLUSTER_ZONE')]) {
                        sh '''
                            gcloud auth activate-service-account --key-file=$GCLOUD_KEYFILE
                            gcloud config set project $PROJECT_ID
                            docker build -t gcr.io/$PROJECT_ID/selestino:${BUILD_NUMBER} .
                            gcloud auth configure-docker
                            docker push gcr.io/$PROJECT_ID/selestino:${BUILD_NUMBER}
                            gcloud container clusters get-credentials $CLUSTER_NAME --zone $CLUSTER_ZONE
                            kubectl apply -f k8s/configmap.yaml
                            kubectl apply -f k8s/secret.yaml
                            kubectl apply -f k8s/deployment.yaml
                            kubectl apply -f k8s/service.yaml
                            kubectl apply -f k8s/ingress.yaml
                            kubectl set image deployment/selestino-deployment selestino=gcr.io/$PROJECT_ID/selestino:${BUILD_NUMBER}
                            kubectl rollout status deployment/selestino-deployment
                        '''
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Build and deployment were successful!'
        }
        failure {
            echo 'Build or deployment failed. Please check the logs.'
        }
        always {
            script {
                currentBuild.result = currentBuild.result ?: 'SUCCESS'
                emailext(
                    to: 'josuejero@hotmail.com',
                    subject: "Jenkins Build: ${currentBuild.fullDisplayName} - ${currentBuild.result}",
                    body: "Check the console output at ${env.BUILD_URL} to view the results."
                )
            }
        }
    }
}

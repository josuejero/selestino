pipeline {
    agent any

    environment {
        PATH = "/usr/local/bin:$PATH"
        GITHUB_CREDENTIALS = credentials('GITHUB_CREDENTIALS_ID') // Use Jenkins credentials for GitHub
        DB_CREDENTIALS = credentials('DB_CREDENTIALS_ID') // Securely manage database credentials
    }

    stages {
        stage('Load Environment Variables') {
            steps {
                script {
                    // Securely load .env variables without exposing them in logs
                    withCredentials([file(credentialsId: 'ENV_FILE_CREDENTIALS', variable: 'ENV_FILE')]) {
                        sh 'export $(cat $ENV_FILE | xargs)'
                    }
                }
            }
        }

        stage('Check Environment') {
            steps {
                script {
                    sh 'echo "User: $(whoami)" [DEBUG-001]'
                    sh 'echo "Current directory: $(pwd)" [DEBUG-002]'
                    // Removed sensitive environment information logging
                }
            }
        }

        stage('Install Python 3 and Setup Virtual Environment') {
            steps {
                script {
                    echo "Checking for Python 3 and setting up virtual environment if necessary... [DEBUG-004]"
                    sh '''
                        if ! command -v python3 &> /dev/null
                        then
                            echo "Python 3 not found, installing Python 3... [DEBUG-005]"
                            sudo apt-get update -y
                            sudo apt-get install -y python3 python3-pip python3-venv
                        fi
                        python3 -m venv venv
                        . venv/bin/activate
                    '''
                }
            }
        }

        stage('Install Python Dependencies') {
            steps {
                script {
                    echo "Installing Python dependencies from requirements.txt... [DEBUG-007]"
                    sh '''
                        . venv/bin/activate
                        pip install --upgrade pip
                        pip install -r requirements.txt
                    '''
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

        stage('Start PostgreSQL Service and Check Database') {
            steps {
                script {
                    echo "Starting PostgreSQL service and verifying database setup... [DEBUG-012]"
                    try {
                        sh '''
                            # Start PostgreSQL without sudo
                            service postgresql start || echo "PostgreSQL is already running"
                            
                            # Securely create the database and user
                            DB_NAME=$(echo $DB_NAME | sed 's/[^a-zA-Z0-9_]//g') # Sanitize DB_NAME
                            DB_USER=$(echo $DB_USER | sed 's/[^a-zA-Z0-9_]//g') # Sanitize DB_USER
                            
                            psql -h $DB_HOST -U postgres -c "CREATE DATABASE $DB_NAME;" 2>/dev/null || echo "Database $DB_NAME already exists"
                            psql -h $DB_HOST -U postgres -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';" 2>/dev/null || echo "User $DB_USER already exists"
                            psql -h $DB_HOST -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
                            psql -h $DB_HOST -U postgres -c "GRANT ALL PRIVILEGES ON SCHEMA public TO $DB_USER;"
                            psql -h $DB_HOST -U postgres -c "ALTER USER $DB_USER WITH SUPERUSER;"
                            psql -h $DB_HOST -U postgres -c "ALTER SCHEMA public OWNER TO $DB_USER;"
                        '''
                        echo "PostgreSQL service started, and database setup verified. [DEBUG-013]"
                    } catch (Exception e) {
                        echo "Error during PostgreSQL setup: ${e.message} [ERROR-102]"
                        error("Failed at stage: Start PostgreSQL Service and Check Database [ERROR-102]")
                    }
                }
            }
        }

        stage('Verify Database Connectivity') {
            steps {
                script {
                    echo "Verifying database connectivity... [DEBUG-014]"
                    try {
                        sh '''
                            PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT 1;"
                        '''
                        echo "Database connectivity verified. [DEBUG-015]"
                    } catch (Exception e) {
                        echo "Error verifying database connectivity: ${e.message} [ERROR-103]"
                        error("Failed at stage: Verify Database Connectivity [ERROR-103]")
                    }
                }
            }
        }

        stage('Apply Migrations') {
            steps {
                script {
                    echo "Applying database migrations... [DEBUG-019]"
                    try {
                        sh '''
                            . venv/bin/activate
                            python selestino/manage.py migrate
                        '''
                        echo "Migrations applied successfully. [DEBUG-020]"
                    } catch (Exception e) {
                        echo "Error applying migrations: ${e.message} [ERROR-107]"
                        error("Failed at stage: Apply Migrations [ERROR-107]"
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    echo "Running tests... [DEBUG-016]"
                    try {
                        if (fileExists('tests')) {
                            echo "tests directory found in the current directory. [DEBUG-017]"
                            sh '''
                                . venv/bin/activate
                                pytest tests/
                            '''
                        } else if (fileExists('selestino/tests')) {
                            echo "tests directory found in the selestino directory. Changing directory... [DEBUG-018]"
                            dir('selestino') {
                                sh '''
                                    . ../venv/bin/activate
                                    pytest tests/
                                '''
                            }
                        } else {
                            error("tests directory not found in either the current or selestino directory [ERROR-105]")
                        }
                    } catch (Exception e) {
                        echo "Error during testing: ${e.message} [ERROR-104]"
                        error("Failed at stage: Run Tests [ERROR-104]")
                    }
                }
            }
        }

        stage('Deploy to Google Cloud') {
            steps {
                script {
                    echo "Skipping deployment to Google Cloud... [DEBUG-020]"
                    // Include your deployment logic here when ready
                }
            }
        }
    }

    post {
        success {
            echo 'Build was successful! [DEBUG-022]'
        }
        failure {
            echo 'Build failed. Please check the logs. [ERROR-106]'
        }
    }
}

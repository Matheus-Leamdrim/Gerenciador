pipeline {
    agent any

    environment {

        FRONTEND_IMAGE = 'frontend'
        BACKEND_IMAGE = 'backend'
        MOBILE_IMAGE = 'mobile'
        KUBE_NAMESPACE = 'dev'
    }

    stages {

        stage('Instalar Dependências - Frontend') {
            steps {
                dir('frontend') {
                    sh 'npm install'
                }
            }
        }

        stage('Instalar Dependências - Backend') {
            steps {
                dir('backend') {
                    sh 'go mod tidy'
                }
            }
        }
        stage('Executar Testes') {
            steps {
                dir('backend') {
                    sh 'go test ./... -v'
                }
            }
        }

        stage('Buildar Imagens Docker') {
            steps {
                script {
                    sh "docker build -t ${FRONTEND_IMAGE}:latest ./frontend"
                    sh "docker build -t ${BACKEND_IMAGE}:latest ./backend"
                    sh "docker build -t ${MOBILE_IMAGE}:latest ./mobile"
                }
            }
        }
        
        stage('Implantar no Kubernetes') {
            steps {
                script {
                    sh "kubectl apply -f kubernetes/${KUBE_NAMESPACE}/frontend-deployment.yaml"
                    sh "kubectl apply -f kubernetes/${KUBE_NAMESPACE}/backend-deployment.yaml"
                    sh "kubectl apply -f kubernetes/${KUBE_NAMESPACE}/mobile-deployment.yaml"
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline executado com sucesso!'
        }
        failure {
            echo 'Pipeline falhou.'
        }
    }
}

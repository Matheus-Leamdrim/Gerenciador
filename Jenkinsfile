pipeline {
    agent any

    environment {
        DOCKER_REGISTRY = 'b7fc73f3-32cb-40f8-833e-bc684689b0f2'
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
        stage('Executar Testes - Frontend') {
            steps {
                dir('frontend') {
                    sh 'npm test'
                }
            }
        }

        stage('Executar Testes - Backend') {
            steps {
                dir('backend') {
                    sh 'go test ./... -v'
                }
            }
        }

        stage('Buildar Imagens Docker') {
            steps {
                script {
                    sh "docker build -t ${DOCKER_REGISTRY}/${FRONTEND_IMAGE}:${BUILD_NUMBER} ./frontend"
                    sh "docker build -t ${DOCKER_REGISTRY}/${BACKEND_IMAGE}:${BUILD_NUMBER} ./backend"
                    sh "docker build -t ${DOCKER_REGISTRY}/${MOBILE_IMAGE}:${BUILD_NUMBER} ./mobile"
                }
            }
        }

        stage('Push das Imagens Docker') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh 'docker login -u ${DOCKER_USER} -p ${DOCKER_PASS} ${DOCKER_REGISTRY}'
                        sh "docker push ${DOCKER_REGISTRY}/${FRONTEND_IMAGE}:${BUILD_NUMBER}"
                        sh "docker push ${DOCKER_REGISTRY}/${BACKEND_IMAGE}:${BUILD_NUMBER}"
                        sh "docker push ${DOCKER_REGISTRY}/${MOBILE_IMAGE}:${BUILD_NUMBER}"
                    }
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
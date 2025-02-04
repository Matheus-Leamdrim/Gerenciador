pipeline {
    agent any

    environment {
        DOCKER_REGISTRY = 'seu-registro-docker'
        KUBE_CONFIG = credentials('kubeconfig')
    }

    stages {
        stage('Instalar Dependências') {
            steps {
                sh 'cd backend && go mod download'
                sh 'cd frontend && npm install'
            }
        }

        stage('Executar Testes') {
            steps {
                sh 'cd backend && go test ./...'
                sh 'cd frontend && npm run test'
            }
        }

        stage('Buildar Aplicação') {
            steps {
                sh 'cd backend && docker build -t ${DOCKER_REGISTRY}/backend:${BUILD_NUMBER} .'
                sh 'cd frontend && docker build -t ${DOCKER_REGISTRY}/frontend:${BUILD_NUMBER} .'
            }
        }

        stage('Push das Imagens Docker') {
            steps {
                sh 'docker push ${DOCKER_REGISTRY}/backend:${BUILD_NUMBER}'
                sh 'docker push ${DOCKER_REGISTRY}/frontend:${BUILD_NUMBER}'
            }
        }

        stage('Implantar no Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/dev/ --kubeconfig=${KUBE_CONFIG}'
            }
        }
    }
}
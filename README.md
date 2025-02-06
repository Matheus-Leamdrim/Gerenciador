# Gerenciador
Projeto de Gerenciamento de Tarefas - Tecnologias: Jenkins, Kubernetes, Flutter, GoLang, React, Next.js e TypeScript.

📝 Gerenciador - Sistema de Gerenciamento de Tarefas
Este repositório contém três projetos independentes para um sistema completo de gerenciamento de tarefas:

Frontend - Aplicação Web com React, Next.js e TypeScript
Mobile - Aplicativo Mobile desenvolvido com Flutter
Backend - API RESTful criada em GoLang
Cada projeto segue boas práticas de arquitetura e inclui testes automatizados.

Os projetos utilizam as seguintes tecnologias:

Frontend: React, Next.js, TypeScript, Jest
Mobile: Flutter, Bloc (gerenciamento de estado)
Backend: GoLang, PostgreSQL, JWT para autenticação
📂 Estrutura do Repositório

Gerenciador/
│── frontend/     # Aplicação web (React + Next.js + TypeScript)
│── mobile/       # Aplicação mobile (Flutter)
│── backend/      # API backend (GoLang + PostgreSQL)


🖥️ Frontend
A aplicação web permite visualizar uma lista de tarefas obtidas do backend simulado e inclui um filtro de busca por nome.

📦 Instalação e Execução
Acesse a pasta do projeto:

cd frontend
Instale as dependências:

npm install
Execute a aplicação em ambiente de desenvolvimento:
npm run dev ou npx next dev

O servidor estará disponível em http://localhost:3000

🧪 Testes
Para rodar os testes unitários com Jest:
npm test


📱 Mobile
O aplicativo mobile permite criar, editar e excluir tarefas, sincronizando os dados com o backend simulado em tempo real.

📦 Instalação e Execução
Acesse a pasta do projeto:

cd mobile
Instale as dependências do Flutter:


flutter pub get
Execute: dart websocket_server.dart no bash, dentro da pasta Mobile.
execute também: flutter run

🧪 Testes
Para rodar os testes:
flutter test


⚙️ Backend
A API backend fornece endpoints RESTful para operações CRUD de tarefas e inclui autenticação com JWT.

📦 Instalação e Execução
Acesse a pasta do backend:

cd backend
Instale as dependências:

go mod tidy
Configure as variáveis de ambiente:

Execute o servidor:
go run cmd/backend/main.go

A API estará disponível em http://localhost:8080
🔐 Autenticação
Os endpoints protegidos exigem um token JWT no cabeçalho:

🧪 Testes
Para rodar os testes de integração:
go test ./...
📜 Licença
Este projeto é de código aberto e está licenciado sob a MIT License.

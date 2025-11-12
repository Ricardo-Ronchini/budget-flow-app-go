# Budget Flow App - Go

Uma aplicação de gerenciamento de orçamentos desenvolvida em Go para fins educacionais, com foco em estudar mecanismos avançados da linguagem e manipulação de estruturas de API.

## 📋 Sobre o Projeto

Este projeto é uma **API REST** para cadastro e gerenciamento de contas.

**Autenticação e Segurança:**
Sistema de login implementado com **JWT (JSON Web Tokens)** para autenticação segura, simulando a base back-end de um projeto real com validação de credenciais.

**Testes e Qualidade:**
Suite de testes integrada para melhorar a segurança da entrega e garantir a confiabilidade das funcionalidades.

**Logging e Monitoramento:**
Sistema de logs internalizado para melhor controle, rastreamento de operações e diagnóstico de problemas.

**Arquitetura de Rotas:**
Sistema de handler de APIs utilizando **Echo do Golang** para padronizar a declaração de novas rotas e manter a consistência do código.

**Gerenciamento de Banco de Dados:**
Sistema de migration dedicado para novas implementações, facilitando versionamento e evolução do banco de dados.

## 🛠️ Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework Web**: Echo
- **Banco de Dados**: PostgreSQL
- **Containerização**: Docker & Docker Compose

## 🚀 Como Começar

### Pré-requisitos

- Docker e Docker Compose instalados
- Go 1.x ou superior (para desenvolvimento local)

### Instalação

#### 1. Criar e Iniciar o Banco de Dados

```bash
docker compose -f ./scripts/docker-compose.yml -p budget-app up -d
```

#### 2. Executar Migrations

```bash
go run ./cmd/migrations
```

#### 3. Executar Script SQL

Execute o script SQL para inicializar o banco de dados:

```bash
psql -U {{user}} -d {{database}} -f ./db/scripts/init.sql
```

#### 4. Iniciar a API

```bash
go run ./cmd/api
```

A API estará disponível em `http://localhost:8080` (ou a porta configurada).

## 📁 Estrutura do Projeto

```
budget-flow-app-go/
├── cmd/
│   ├── api/          # Aplicação principal
│   └── migrations/   # Scripts de migração
├── scripts/
│   └── docker-compose.yml
└── ...
```

## 📝 Notas

- Certifique-se de que o Docker está em execução antes de criar o banco de dados
- As migrations devem ser executadas antes de executar os scripts SQL
- Atualize as credenciais do banco de dados conforme sua configuração

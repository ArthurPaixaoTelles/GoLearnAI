# ☁️ IsCoolGPT - Cloud Native AI Assistant

> Assistente de Inteligência Artificial Fullstack desenvolvido em Go, containerizado com Docker e orquestrado na AWS (ECS Fargate) através de uma robusta pipeline CI/CD.

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Multi_Stage-2496ED?logo=docker&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-ECS_Fargate-FF9900?logo=amazon-aws&logoColor=white)
![CI/CD](https://img.shields.io/badge/GitHub_Actions-Automated-2088FF?logo=github-actions&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

---

## 📖 Sobre o Projeto

O **IsCoolGPT** é um projeto da disciplina de Cloud Computing, cujo foco é a implementação de uma arquitetura de microsserviços modular moderna. Ela consiste na integração do modelo de texto **gpt-oss-20b**, que está hospedado na plataforma **Hugging Face**, com os serviços da AWS, permitindo seu uso de maneira fácil e simples.

O diferencial do projeto é o foco em **DevOps**, **FinOps** (otimização de custos na nuvem) e **Automação**, utilizando infraestrutura como código e esteiras de deploy contínuo. Após integrar tudo, foi construída uma página simples de frontend para facilitar a interação.

---

## 🏗️ Arquitetura e Tecnologias

### 1. Backend & Frontend (Go Monolith)
* **Linguagem:** Go (Golang) 1.22.
* **Roteamento:** Biblioteca `Gorilla Mux`.
* **Frontend:** HTML5/JS Vanilla servido estaticamente pelo próprio binário Go (reduzindo complexidade de infraestrutura).
* **Integração:** Consumo da API de Inferência Serverless da Hugging Face.

### 2. Containerização (Docker)
Utilização de **Multi-stage Builds** para alta otimização:
* **Stage 1 (Builder):** Imagem `golang:alpine` para compilação.
* **Stage 2 (Runtime):** Imagem `alpine:latest` para execução.
* **Resultado:** Imagem final de apenas **32.4MB** no Docker e **~9.17MB** no AWS ECR, garantindo deploys rápidos e menor custo de armazenamento.

### 3. Infraestrutura Cloud (AWS)
* **Compute:** AWS ECS (Elastic Container Service) com **Fargate Spot** (Serverless).
    * *Decisão:* O uso de Spot Instances reduz o custo computacional em até 70%.
* **Networking:** Application Load Balancer (ALB) para gestão de tráfego e endpoint fixo.
* **Registry:** AWS ECR (Elastic Container Registry) para armazenamento seguro das imagens.
* **Segurança:** Security Groups configurados em camadas (Firewall Virtual):
    * O **ALB** aceita tráfego público na porta 80.
    * Os **Containers** aceitam tráfego apenas vindo do ALB na porta 8080 (isolamento total).
    * **Aplicação:** Injeção segura de segredos (`HF_API_KEY`) em tempo de execução via Variáveis de Ambiente, garantindo que credenciais sensíveis nunca fiquem expostas no código.

---

## ⚙️ Pipeline de CI/CD (DevOps)

O projeto implementa um fluxo de **GitFlow Simplificado** automatizado via GitHub Actions:

| Branch | Ambiente | Ação Automática |
| :--- | :--- | :--- |
| `dev` | Desenvolvimento | **CI:** Execução de Testes Unitários e Integração (Mock). |
| `staging` | Pré-Produção | **CI + CD:** Testes + Build Docker + Deploy automático no ECS (Ambiente de Validação). |
| `main` | Produção | **CI + CD:** Testes + Build Docker (Tag `latest`) + Deploy final no ECS. |

### Testes Automatizados
Os testes (`*_test.go`) utilizam **Mocking** para simular a API da Hugging Face, permitindo validar a lógica de negócios sem depender de conectividade externa ou gastar cotas de API durante a pipeline.

---

## 🚀 Como Acessar e Testar

A aplicação está implantada e rodando em alta disponibilidade na AWS.

### 🌐 Acesso via Navegador (Frontend)
Acesse o chat interativo através do Load Balancer da AWS:
👉 **[http://alb-golearnai-2009330452.us-east-2.elb.amazonaws.com](http://alb-golearnai-2009330452.us-east-2.elb.amazonaws.com)**

### 🧪 Testando via Terminal (API)


### Linux / macOS / Bash
curl -X POST <http://alb-golearnai-2009330452.us-east-2.elb.amazonaws.com/api/prompt> \
     -H "Content-Type: application/json" \
     -d '{"question": "Qual a vantagem de usar Docker?"}'


### Windows (PowerShell)

$body = @{ question = "Qual a vantagem de usar Docker?" } | ConvertTo-Json

Invoke-RestMethod -Uri <http://alb-golearnai-2009330452.us-east-2.elb.amazonaws.com/api/prompt> `
    -Method POST `
    -Body $body `
    -ContentType "application/json"



---

## 🛡️ Monitoramento e Custos

- **CloudWatch Logs:** logs completos + alarmes de CPU (>80%)  
- **Auto Scaling:** 1–3 tarefas Fargate, baseado em CPU  
- **AWS Budgets:** alertas configurados  
- **FinOps:** uso de Fargate Spot + imagem leve reduz drasticamente o custo

---

## 📄 Licença

Distribuído sob a licença **MIT**.

---

## 🤝 Contribuição

Contribuições são bem-vindas!  
Sinta-se à vontade para abrir uma *issue* ou enviar um *pull request*.

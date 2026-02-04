# Treinamento Docker 

## Visão geral

Este repositório contém exemplos e exercícios práticos para o treinamento de Docker. Ele abrange desde conceitos básicos até tópicos avançados, incluindo a criação e gerenciamento de contêineres, redes, volumes e orquestração com Docker Compose.

## Programa

[![Notion](https://img.shields.io/badge/Notion-000000?style=flat-square&logo=notion&logoColor=white)](https://www.notion.so/Treinamento-Docker-2fdd3a700afe80b599e2f4c3ccfb9eed?source=copy_link) **Conteúdo do treinamento**
- Página no Notion com todo o conteúdo do treinamento
- Nele estão explicações das implementações presentes neste repositório

[![DOCS](https://img.shields.io/badge/DOCS-4B0082?style=flat-square&logo=read-the-docs&logoColor=white)](./COMPILING.md) **Introdução Sistemas Operacionais**
- Material introdutório sobre sistemas operacionais
- Aborda conceitos como kernel, chamadas de sistema e processos 

[![Python](https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white)](./hello-py/) **Nossa própria imagem, "Hello World!**
- Exemplo simples em Python com Docker
- Contém `Dockerfile` e `hello.py` em `hello-py/`
- Como rodar: `docker build -t hello-py ./hello-py && docker run --rm hello-py`

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](./producer-consumer/) **Um produz, o outro consume**
- Exemplo de producer/consumer usando Redis
- Contém `Producer.dockerfile`, `Consumer.dockerfile`, `docker-compose.yaml` e código Go em `producer/` e `consumer/`
- Como rodar: `docker compose up --build` (a partir de `producer-consumer/`)


[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white)](./postgres-server/) **Persistindo os registros de um banco de dados**
- Serviço PostgreSQL com script de migrações
- Contém `docker-compose.yaml` e scripts em `scripts/`
- Como rodar: `docker compose up -d` (a partir de `postgres-server/`)



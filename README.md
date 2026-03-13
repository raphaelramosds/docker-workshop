# Treinamento Docker 

## Visão geral

Este repositório contém exemplos e exercícios práticos para o treinamento de Docker. Ele abrange desde conceitos básicos até tópicos avançados, incluindo a criação e gerenciamento de contêineres, redes, volumes e orquestração com Docker Compose.

## Programa

[![Notion](https://img.shields.io/badge/Notion-000000?style=flat-square&logo=notion&logoColor=white)](https://www.notion.so/Treinamento-Docker-2fdd3a700afe80b599e2f4c3ccfb9eed?source=copy_link) **Conteúdo do treinamento**
- Página no Notion com todo o conteúdo do treinamento
- Nele estão explicações das implementações presentes neste repositório

[![DOCS](https://img.shields.io/badge/DOCS-4B0082?style=flat-square&logo=read-the-docs&logoColor=white)](./intro-so/README.md) **Introdução Sistemas Operacionais**
- Material introdutório sobre sistemas operacionais
- Aborda conceitos como kernel, chamadas de sistema e processos 

[![Python](https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white)](./hello-py/) **Nossa própria imagem, "Hello World!"**
- Exemplo simples em Python com Docker
- Demonstra como criar uma imagem

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](./hello-go/) **Hello Go: Docker com e sem multi-stage**
- Exemplo básico em Go com Dockerfile multi-stage e versão simples
- Inclui Makefile para build e limpeza de imagens

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](./producer-consumer/) **Um produz, o outro consume**
- Exemplo de producer/consumer usando Redis
- Demonstra a facilidade de controlar containers com Docker Compose

[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white)](./postgres-server/) **Persistindo os registros de um DB**
- Serviço PostgreSQL com script de migrações
- Demonstra o uso de volumes para persistência de dados

[![Linux](https://img.shields.io/badge/Linux-FCC624?style=flat-square&logo=linux&logoColor=black)](./intro-so/) **Prática de Sistema Operacional**
- Exercício com comandos UNIX e análise de chamadas de sistema com strace
- Inclui arquivos de exemplo para observação de interação com o kernel
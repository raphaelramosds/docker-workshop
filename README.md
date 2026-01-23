# Workshop Docker

## Programa

Introdução
- O que é um sistema operacional? (discussão de kernel e aplicações)
- O que é um processo?
- O que são distribuições Linux?
- Formas de virtualização: hipervisores e containers
- E o WSL? (Execute o comando uname -r e veja o kernel do Linux rodando no WSL e numa VM Hyper-V)

Docker 
- Arquitetura
- Imagem e container (imagem base e camadas)
- Docker Hub 
- Ciclo de vida de uma imagem e de um container: criação, execução, parada, remoção
- Orquestração com Docker Compose
- Volumes
- Redes

Hands-on
- Instalação do Docker (Linux e Windows)
- "Hello, World!" no Docker
- Um container que apenas dorme: `docker run -d alpine sleep 3600` (perfeito para exemplificar que o container precisa de pelo menos um processo rodando, isto é, seu entrypoint) 
- Produtor e consumidor: criando uma aplicação simples com dois containers se comunicando via socket TCP
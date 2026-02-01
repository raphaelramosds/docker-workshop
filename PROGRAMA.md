# Workshop Docker

## Programa

Introdução
- O que é um sistema operacional? (discussão de kernel e aplicações)
- O que é um processo? (exemplificar executando um script python de um loop que imprime Hello a cada 2s, e observá-lo no htop; o processo é um programa em execução!)
- O que são distribuições Linux?
- Formas de virtualização: hipervisores e containers
- WSl (um caso especial): execute o comando uname -r no WSL e numa VM Hyper-V (com Linux Mint) para ver a implementação do kernel de cada uma

Docker 
- Arquitetura: docker daemon
- Imagem e container: introduzir o que é imagem base, o conceito de camadas (layers) e caching
- Ciclo de vida de uma imagem e de um container: criação, execução, parada, remoção
- Docker Hub 
- Orquestração com Docker Compose: comparar um comando longo com docker run (subindo uma API e um client consumindo ela) e a sua versão num arquivo de configuração docker compose
- Volumes: docker compose down versus docker compose down --volumes
- Redes: link entre containers; internal versus external network; DNS (resolver IP pelo nome do container) 

Hands-on
- Instalação do Docker (Linux e Windows)
- "Hello, World!" 
- Um container que apenas dorme: docker run -d alpine sleep 3600 (exemplificar que o container precisa de pelo menos um processo rodando, isto é, seu entrypoint) 
- Produtor e consumidor: criando uma aplicação em Go simples com dois containers se comunicando via socket TCP

Publicando sua imagem: vamos aproveitar a imagem do "container que apenas dorme" para essa parte
- Criando uma conta no Docker Hub
- Tagueando sua imagem
- Publicando sua imagem no Docker Hub
- Baixando uma imagem do Docker Hub

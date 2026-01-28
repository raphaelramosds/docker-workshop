# Plano

**Título:** Treinamento - Docker

## Sistemas operacionais

### O que é?

É um software que coordena e controla o uso do hardware, requisitados pelas aplicações utilizadas pelos usuários.

### Kernel e aplicações 

Um sistema operacional é composto por um **kernel** (núcleo) e aplicações. O usuário utiliza aplicações que se comunicam com o kernel através de chamadas de sistema.

Alguns exemplos de chamadas de Sistema do UNIX

- execv
- read
- open
- close

### Exemplo: o binário `echo`

Abaixo usamos a aplicação `echo` para solicitar ao sistema operacional a escrita de uma nova linha em um arquivo de texto e salvá-lo no disco

```bash
/usr/sbin/echo "Mais uma linha" >> arquivo.txt
```

O utilitário `which <aplicacao>` do UNIX informa o caminho de um binário instalado no sistema operacional. No exemplo acima, o programa `echo` está localizado na pasta `/usr/sbin`

Caso queira visualizar as chamadas de sistema envolvidas no comando acima

```bash
strace -f -o log_strace.txt sh -c 'echo "Mais uma linha" >> arquivo.txt'
```

### O Kernel Linux

Lançado em 1991 por Linus Torvalds, como um software livre de código aberto.

Atualmente o kernel Linux é mantido pela comunidade, sob supervisão do próprio Linus Torvalds, no repositório [torvalds/linux](https://github.com/torvalds/linux)

Você pode visualizar qual release do Kernel linux está instalada na sua máquina com o programa `uname` 

```bash
uname -sr
```

### Distribuições Linux

Distribuições Linux são sistemas operacionais que possuem o kernel Linux.

As distribuições se diferem entre si pelas aplicações presentes nelas. Observe abaixo como o requisito de **espaço livre no HD** necessário para instalá-las varia entre elas 

- [Ubuntu 24.04.4 LTS](https://ubuntu.com/download/desktop): pelo menos 25 GB
- [Arch Linux](https://wiki.archlinux.org/title/Installation_guide): até 2 GB 
- [Alpine Linux](https://wiki.alpinelinux.org/wiki/Installation): até 700 MB

Quanto mais aplicações presentes em uma distribuição, mais espaço ela ocupa 

## Virtualização

### Contextualizando

Um sistema operacional pode ser virtualizado, isto é, emulado dentro de um outro sistema operacional

### Formas de virtualização

Existem duas formas de se fazer isso: hipervisores e containers.

1. Hipervisor (Oracle Virtual Box, VMWare), no qual toda a estrutura do sistema operacional (kernel e aplicações) são executados em cima do sistema operacional hospedeiro (host).

2. Execução de um container (Docker, Podman, Containerd), no qual as aplicações se comunicam com o kernel do sistema operacional hospedeiro

## Docker

### O que é? 

Solução que possibilita a execução de containers a partir de uma arquitetura cliente e servidor

### Arquitetura (Visão Geral)

Os componentes principais da sua arquitetura são

- Docker client (`docker`)
- Docker Host: Docker daemon (`dockerd`), Imagens, Containers
- Registry

O Docker client se comunica com o Docker daemon via API REST a partir de comandos como `docker run` e `docker build`. 

https://docs.docker.com/get-started/docker-overview/

### Imagem 

Todas as imagens Docker são baseadas em alguma distribuição Linux.

O ciclo de vida de uma imagem Docker baseia-se em

1. Construção (`docker build <argumentos>`): necessário uma arquivo de definição Dockerfile
2. Execução (`docker run <argumentos> imagem`) 
3. Encerramento (`docker stop imagem`)

Ao executar `docker run imagem`, o *daemon* do Docker vai procurá-la no registry local (isto é, dentro do seu disco).

Caso não encontre, o comando `docker pull imagem` é executado automaticamente para fazer o download dessa imagem no Docker Hub (registry remoto padrão do Docker) 

```bash
docker run imagem
```
# Treinamento Docker  

## Distribuições Linux

Distribuições Linux são sistemas operacionais que possuem o kernel Linux, necessariamente.

As distribuições se diferem entre si pelas aplicações presentes nelas. Observe abaixo como o requisito de **espaço livre no HD** necessário para instalá-las varia entre elas 

- [Ubuntu 24.04.4 LTS](https://ubuntu.com/download/desktop) (pelo menos 25 GB)
- [Arch Linux](https://wiki.archlinux.org/title/Installation_guide) (até 2 GB)
- [Alpine Linux](https://wiki.alpinelinux.org/wiki/Installation) (até 700 MB)

Quanto mais aplicações presentes em uma distribuição, mais espaço ela ocupa 

## Virtualização

### Contextualizando

Um sistema operacional pode ser virtualizado, isto é, emulado dentro de um outro sistema operacional

### Formas de virtualização

Existem duas formas de se fazer isso: hipervisores e containers.

1. Hipervisor (Oracle Virtual Box, VMWare), no qual toda a estrutura do sistema operacional (kernel e aplicações) são executados em cima do sistema operacional hospedeiro (host).

2. Execução de um container (Docker, Podman, Containerd), no qual as aplicações se comunicam com o kernel do sistema operacional hospedeiro

### Por que virtualizar?

- Isolamento de aplicações: garante que apenas as aplicações necessárias estejam presentes no ambiente
- Facilidade de reprodução aplicações: possibilita a criação de ambientes idênticos em diferentes máquinas
- Padronização de ambientes de desenvolvimento, teste e produção

## Docker

### O que é? 

Solução que possibilita a execução de containers a partir de uma arquitetura client e servidor

### Arquitetura (Visão Geral)

Os componentes principais da sua arquitetura são

- Docker Client (`docker`) 
- Docker Daemon 
- Docker Registry

O Docker client se comunica com o Docker Host via API REST a partir de comandos do tipo `docker <nome do comando>`

Mais em: https://docs.docker.com/get-started/docker-overview/

### Instalação

#### Windows

O principal requisito para instalar o Docker no Windows é que o WSL (Windows Subsystem for Linux) esteja instalado e configurado com uma distribuição Linux. Para isso, siga as instruções em https://learn.microsoft.com/pt-br/windows/wsl/install

Com o WSL instalado baixe o instalador do Docker Desktop em https://docs.docker.com/desktop/install/windows-install/

**Observação:** É possível utilizar o Docker no Windows sem o Docker Desktop, mas optando por esse caminho, você terá que instalar e configurar o Docker Engine manualmente dentro do WSL.

#### Linux 

Seguem links para instalação do Docker em algumas distribuições Linux populares

- [Arch Linux](https://docs.docker.com/desktop/setup/install/linux/archlinux/)
- [Ubuntu](https://www.notion.so/Como-instalar-o-Docker-no-Linux-2e7d3a700afe80de8194dfb70cd32b65?source=copy_link)

### Imagem 

"Uma imagem Docker é a materialização de um modelo de um sistema de arquivos, modelo este produzido através de um processo chamado *build*."

De forma simples, a imagem é o componente que guarda as definições de aplicações e configurações que estarão presentes em um ambiente de execução (*runtime*).

Todas as imagens Docker são baseadas em alguma distribuição Linux ou algum SO que emule o kernel Linux.

Liste as imagens do seu *registry local* com

```bash
docker images
```

Podemos tanto realizar o download de uma imagem como construí-la.

#### Download de uma imagem

O que acontece ao executar o comando abaixo?

```bash
docker pull hello-world
```

1. Docker Client envia a requisição para obter a imagem `hello-world`
2. Docker Daemon procura essa imagem no *registry local*
    - Caso não encontre, realiza o download dela do Docker Hub no repositório público do *registry* docker.io 
    - Caso a encontre, informa que a imagem já existe 

O docker assume `latest` como a tag padrão do comando acima, isto é, a sua *release* mais recente. Mas, é possível definir uma outra tag, como demonstrado abaixo

```bash
docker pull hello-world:linux
```

As tags disponíveis para a imagem `hello-world` estão listadas em https://hub.docker.com/_/hello-world/tags

#### Construindo sua imagem 

A construção de uma imagem Docker envolve a escrita de um arquivo de *definição de imagem* chamado Dockerfile 

Podemos dividir a estrutura de um Dockerfile em

1. Imagem base: contém a imagem da qual vamos aproveitar as aplicações e configurações (pode ser uma imagem construída de outro Dockerfile, inclusive)
2. Camadas intermediárias: execução de instruções como WORKDIR, COPY e RUN 
3. Ponto de partida: processo que será executado ou argumentos enviados para o processo já existente na imagem base

**Exemplo:** vamos escrever uma imagem cujo ponto de partida seja um programa em Python que imprime "Hello World" indefinidamente a cada 5 segundos

```Dockerfile
# Imagem base
FROM python:3.11-slim

# Evitar buffering de saida (imprime imediatamente nos logs)
ENV PYTHONUNBUFFERED=1

# Camadas intermediárias
WORKDIR /usr/local
COPY hello.py .

# Ponto de partida 
ENTRYPOINT ["python", "hello.py"]
```

Para construir a imagem a partir deste Dockerfile execute o comando abaixo 

```bash
docker build -t hello-py:local -f hello/Dockerfile hello-py
```

- A imagem será nomeada como `hello-py` e sua tag será `local`
- O arquivo Dockerfile está localizado no diretório `hello-py`
- O contexto de build é o diretório `hello-py` (arquivos copiados via instrução `COPY` são relativos a esse diretório)

### Container

"Container é um sistema de arquivos criado a partir de uma imagem onde será executado um processo isolado de todo o resto do ambiente"

De forma simples, um container possibilita a execução de uma imagem.

Para executar a imagem `hello-world` que obtemos anteriormente e logo após executar um container a partir dela

```bash
docker run --name hello-docker hello-world 
```

Como essa imagem não executa um processo persistente, isto é, ele apenas imprime uma mensagem e para, o container para sua execução também. 

Para visualizar os containers (em execução e parados), execute 

```bash
docker ps --all
```

A saída desse comando se parece com

```bash
CONTAINER ID   IMAGE         COMMAND    CREATED          STATUS                      PORTS     NAMES
f65ce2afbf99   hello-world   "/hello"   20 seconds ago   Exited (0) 20 seconds ago             jolly_panini
```

- CONTAINER ID: identificador único do container
- IMAGE: nome da imagem utilizada pelo container
- COMMAND: processo que esse container executou (script `/hello`)
- CREATED: tempo desde o momento em que o container foi iniciado
- STATUS: estado do container

Caso queira executar o container novamente 

```bash
docker container start -a hello-docker
```
O uso da flag `-a (--attach)` no comando acima faz com que o Docker Client anexe a saída do container ao terminal atual. Sem essa flag não conseguiríamos ver a mensagem impressa pelo script `/hello`, e o container pararia imediatamente após a execução desse script.

#### Nossa própria imagem "Hello, World!"

Escrevemos nossa própria versão da imagem `hello-world`. Mas, ao invés de seu entrypoint ser um script apenas imprima "Hello World" e termine sua execução, fizemos com que ela execute um script Python que imprima "Hello World" a cada 5 segundos indefinidamente.

```bash
# Build da imagem
docker build -t hello-py:local -f hello-py/Dockerfile hello-py

# Execução do container
docker run --name sleep-py hello-py:local
```

OBS: Você pode adicionar a flag `-d` (detached) no comando `docker run` para executar o container em segundo plano (*detached mode*). Por padrão, o container é atrelado ao terminal atual, isto é, a flag `-a` (attach) é assumida implicitamente.

Podemos ver a saída do nosso container com o comando

```bash
docker logs sleep-py 
```

Ainda, podemos "entrar" no container para interagir com ele, utilizando o comando

```bash
docker exec -it sleep-py /bin/bash
```

Remova o container com a sequência de comandos abaixo 

```bash
# Pare a execução (caso esteja em execução)
docker container stop sleep-py

# Remova o container
docker container rm sleep-py
```

Agora faça um experimento: comente a linha do `ENTRYPOINT` no Dockerfile da imagem `hello-py`, faça o build novamente e execute o container. Porque não conseguimos mais "entrar" nesse container para interagir com ele? 

Por que não há mais um processo persitente no seu ponto de partida! Exemplos de processos persistentes são servidores web, bancos de dados, shells interativos, etc.
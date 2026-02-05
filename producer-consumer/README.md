# Produtor e Consumidor 

Exemplo de aplicação utilizando Redis como broker de mensagens entre um produtor e um consumidor.

## Executando os processos

### Manualmente

A inicialização dos processos consumidor e produtor precisam que o Go esteja instalado na máquina local.

```bash
# Execute o produtor
sh execute producer

# Execute o consumidor (em outro terminal)
sh execute consumer
```

### Usando Docker Compose

Inicie todos os serviços (redis, produtor e consumidor)

```bash
docker compose up -d
```

## Interagindo com o Redis CLI

No container, entre no Redis CLI

```bash
redis-cli -a local
```

Veja se a mensagem publicada existe (será retornado 1 se existir, 0 caso contrário)

```bash
EXISTS nome_da_mensagem
```

Recupere o valor da mensagem

```bash
GET nome_da_mensagem
```

No caso de ser um array

```bash
LRANGE nome_da_mensagem 0 -1
```
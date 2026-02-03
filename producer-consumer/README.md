# Produtor e Consumidor 

Exemplo de aplicação utilizando Redis como broker de mensagens entre um produtor e um consumidor.

## Interagindo com o Redis
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
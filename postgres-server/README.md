# Persistindo os dados de um DB

## Executando

```bash
# Inicialize o servidor Postgres 
docker compose up -d

# Execute as migracoes dentro do container
docker exec postgres-server-db-1 bash /etc/scripts/migrar 
```

## Acessando o banco

```bash
# Acesse o container do Postgres
docker exec -it postgres-server-db-1 bash

# Entre na linha de comando do Postgres
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

# Liste as tabelas
\dt

# Veja os dados inseridos
SELECT * FROM users;
```

## Experimentos 

Faça o que se pede no `docker-compose.yaml`

```yaml
volumes:
    # Comente a próxima linha
    - pgdata:/var/lib/postgresql/data
```

Remova o container e execute-o novamente

```bash
# Remover o container
docker compose down

# Iniciar o container novamente
docker compose up -d
```

O que acontece com as tabelas e registros inseridos pelo script `/etc/scripts/migrar`?

Agora, descomente a linha, remova o container e execute-o novamente. O que acontece com os dados criados anteriormente?

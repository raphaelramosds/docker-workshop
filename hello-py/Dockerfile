# Imagem base
FROM python:3.11-slim

# Camadas intermediárias
WORKDIR /usr/local
COPY hello.py .

# Ponto de partida 
ENTRYPOINT ["python", "hello.py"]
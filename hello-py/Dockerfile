# Imagem base
FROM python:3.11-slim

# Evitar buffering de saida (imprime imediatamente nos STDOUT/STDERR)
ENV PYTHONUNBUFFERED=1

# Camadas intermediárias
WORKDIR /usr/local
COPY hello.py .

# Ponto de partida 
ENTRYPOINT ["python", "hello.py"]
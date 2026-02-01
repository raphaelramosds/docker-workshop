# Introdução Sistemas operacionais

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
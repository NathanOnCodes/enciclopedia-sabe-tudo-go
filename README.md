# Projeto Web Crawler, Enciclopédia Sabe tudo.
### Feito com a linguagem de programção Go:

Este projeto consiste em um Web Crawler que busca informações em uma API e as armazena em um arquivo de log. 
O usuário pode pesquisar informações através de uma flag -search, ler as informações armazenadas no arquivo de log através 
da flag -read e apagar todo o histórico do arquivo através da flag -del.

## Pacotes nativos utilizados

Os pacotes nativos utilizados no projeto foram:

- `fmt`: para formatação de strings.
- `os`: para operações de sistema, como abrir arquivos.
- `bufio`: para leitura de arquivos linha por linha.
- `flag`: para definir as flags do programa.
- `net/http`: para realizar requisições HTTP na API.
- `encoding/json`: para decodificar a resposta da API em formato JSON.
- `time`: para inserir o horario da requisição na struct Person.

## Como usar

Para usar siga os passos abaixo:

1. Faça o download do código fonte do projeto.
2. Abra o terminal e navegue até o diretório onde o código fonte foi baixado.
3. Compile o código com o comando: `go build -o mini-wiki main.go`.
4. Rode o comando `./mini-wiki` com as seguintes flags:
   - `-search`: pesquisa informações na API e armazena no arquivo de log.
   - `-read`: lê as informações armazenadas no arquivo de log e exibe no terminal.
   - `-del`: apaga todo o histórico do arquivo de log.
   
   Exemplo de comando para pesquisa: `./mini-wiki -search=sao paulo`.
   Exemplo de comando para leitura: `./mini-wiki -read`.
   Exemplo de comando para deletar: `./mini-wiki -del`.

## Cross-compilação

Este projeto foi feito para ser executado em diferentes plataformas, 
por isso é necessário compilar o código para cada arquitetura desejada. sendo assim, não deixei um executável pré-compilado.



# Desafio 2 Golang - Imersão Full Stack && Full Cycle

<hr>

<img src="https://camo.githubusercontent.com/b54d2e6bf5f15ddf3dd884b7d1bf21c7d5cc8798d119d74a6538c1a1b583a49b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f446f636b65722d3234393645443f7374796c653d666f722d7468652d6261646765266c6f676f3d646f636b6572266c6f676f436f6c6f723d7768697465" /> <img src="https://camo.githubusercontent.com/b24523991dccb1991a2b8faa8dd97f48a0944157eacb7316b32834c2e070cf09/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f2d2532333030414444382e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d676f266c6f676f436f6c6f723d7768697465" /> <img src="https://camo.githubusercontent.com/b310667470594171440f9b80f624787ea58555296d88af177788509b0d73a40b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f73716c6974652d2532333037343035652e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d73716c697465266c6f676f436f6c6f723d7768697465" />


Informações do desafio

Neste desafio, você deverá criar uma aplicação Golang que realiza transações bancárias.

A aplicação será uma API Rest contendo 2 endpoints:

```bash
- POST /bank-accounts - Criar contas bancárias
```

No corpo da requisição deverá ser enviado:

```bash
{
"account_number": "1111-11"
}
```

A resposta HTTP deverá ser 201, contendo o ID da conta criada e o "account_number"

```bash
- POST /bank-accounts/transfer - Transferência entre contas bancárias
```

No corpo da requisição deverá ser enviado:

```bash
{
"from": "1111-11"
"to": "1111-11"
"amount": 100
}
```

A aplicação deverá persistir os dados no banco de dados SQLite.

A resposta HTTP deverá conter o saldo da conta from e o saldo da conta to.

Disponibilize esta aplicação Golang com Docker Compose na porta 8000.

Ao rodar docker compose up todo ambiente deverá já estar disponível.


# Como executar o projeto

```bash
#Baixar o repositório do projeto
git clone git clone https://gabrielalves962008@bitbucket.org/gabrielalves962008/golang-api.git

#acessando a pasta com a aplicação 
cd golang-api

#inicializando o projeto no docker
docker-compose up

```

# Cute dogs email

Olá, se você gosta de cachorrinhos fofos assim como eu, então irá amar essa aplicação ❤️

O cute dogs email é um encaminhador de imagem de cachorrinhos para o seu email, sendo entregue às 6:00 horário de Brasília.

Todos os créditos das imagens vão para **[Dog API](https://dog.ceo/dog-api/)**

# Sumário
- ⚒️ [Ferramentas utilizadas](#ferramentas-utilizadas)
- ⚙️ [Configurando variáveis de ambiente](#configurando-variáveis-de-ambiente)
- 🌐 [Configurando o arquivo nginx.conf (apenas em caso de - deploy)](#configurando-o-arquivo-nginxconf-apenas-em-caso-de-deploy)
- 🖥️ [Como rodar a aplicação](#como-rodar-a-aplicação)
- 💻 [Rodando localmente](#rodando-localmente)
- 🐋 [Rodando no docker-compose](#rodando-no-docker-compose)
- 🥭 [Criar uma database e collection no MongoDB](#3-criar-uma-database-e-collection-no-mongodb)
- ⛵ [Deploy por meio do Kubernetes](#deploy-por-meio-do-kubernetes)

# Ferramentas utilizadas
Frontend:
- React com Typescript
- Uma pitada de css

Backend:
- Go
- Gofiber
- Gomail
- MongoDB

Infra:
- Docker
- Kubernetes
- Google Cloud Provider (Não abordei seu uso aqui, pois quis deixar livre para o uso em qualquer outro provider)
- Nginx



# Configurando variáveis de ambiente

Existem dois arquivos dotenv exemplo na aplicação: um servirá para a aplicação backend e o segundo para a frontend.

> **_OBSERVAÇÃO:_** Não esqueça de substituir o nome **_.env.example_** para **_.env_** 

O primeiro localizado no diretório raiz **.env.example** recebe os seguintes valores

```
FROM_EMAIL= Seu email que será utilizado como remetente

FROM_PASSWORD= senha ou token referente ao email acima
```

> **_OBSERVAÇÃO:_** As variáveis abaixo podem ser preenchidas de acordo com o mongo instalado em sua máquina ou caso você opte, pela cloud do MongoDB. É possível criá-lo de forma gratuita. Para saber mais **[acesse aqui](https://cloud.mongodb.com/)**.

```
MONGO_CREDENTIALS=mongodb://127.0.0.1:27017

MONGO_DB=nome da database do mongo

MONGO_COLLECTION=nome da collection
```

O segundo localizado no diretório **/frontend/cute-dogs/.env.example** recebe o seguinte valor:

```
REACT_APP_YOUR_DOMAIN=http://localhost:3001
```

> **_OBSERVAÇÃO:_** o valor inserido aponta para o servidor backend localmente, mas caso seja realizado deploy o valor deverá ser o DNS, configurado no arquivo arquivo **[nginx.conf](#configurando-o-arquivo-nginxconf-apenas-em-caso-de-deploy)**.

## Configurando o arquivo nginx.conf (apenas em caso de deploy) 

Altere o nome do arquivo **_nginx.conf.example_** para **_nginx.conf_** localizado em **./frontend/cute-dogs/nginx.conf.example**

Altere os valores:

```
$your-dns=seu-dns.com
$your-port=porta (Ex.: 3001)
```

> **_OBSERVAÇÃO:_** para portas com valores default HTTP (80) ou HTTPS (403) não é necessário informar.

# Como rodar a aplicação

O projeto pode ser executado:
1. [localmente (instalando as ferramentas necessárias)](#rodando-localmente)
2. [docker-compose](#rodando-no-docker-compose)
3. [Deploy por meio do Kubernetes](#deploy-por-meio-do-kubernetes)

## Rodando localmente

Será necessário as seguintes ferramentas instaladas:
- [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm/)
- [go](https://go.dev/doc/install)
- [mongodb](https://www.mongodb.com/docs/manual/installation/)

1. [Configure as variaveis de ambiente](#configurando-variáveis-de-ambiente) 
2. No mongodb [crie uma database e collection](#3-criar-uma-database-e-collection-no-mongodb) com o mesmo valor informado nas variáveis de ambiente.
3. Na raiz do projeto rode o comando ```go run main.go``` para rodar o servidor.
4. Vá para dentro da pasta **./frontend/cute-dogs** e execute o comando ```npm start``` para iniciar o frontend.
5. Acesse a aplicação em http://localhost:3000

## Rodando no docker-compose
Acesse o arquivo **_./docker-compose-example.yml_** e altere seu nome para **docker-compose.yml**
1. [Configure as variaveis de ambiente](#configurando-variáveis-de-ambiente)
2. Rode o comando ```docker-compose up -d``` para executar os containers em background.
### 3. Criar uma database e collection no MongoDB
3.1. Para criar o banco no mongo e a collection acesse o container pelo seguinte comando ```docker exec -it mongodb-cute-dogs bash ```

3.2. Digite ```mongosh``` para acessar o terminal do mongo.

3.3. Escolha um nome para o seu database, para fins de exemplo escolherei o nome cute-dogs-db. Rodando o seguinte comando: ```use cute-dogs-db``` . O comando ```use``` permite utilizar um database ou criar um novo quando não existe.

3.4. Ao ter acessado o database basta criar a collecition com o comando: ```db.createCollection("nome-da-sua-collection")```
É esperado o resultado:
> { ok: 1 }

> **_OBSERVAÇÃO:_** O nome escolhida para a database e a collection deverão ser informadas no docker-compose.

7. Digite ```quit```e logo em seguida ```exit``` para sair do mongosh e do terminal do container respectivamente.
8. Acesse a aplicação em http://localhost:3000 

## Deploy por meio do Kubernetes
Os arquivos encontram-se no diretorio raiz **_./k8s-example_** e voce pode renomeá-lo removendo o sufixo **-example** caso queira. Escolha o servidor cloud de sua preferência e os adicione por meio do comando ```kubectl apply -f ./k8s-example```.

> **_OBSERVAÇÃO:_** No arquivo **secret.yaml** converta os valores para base64 ao inserir nas variáveis de ambiente sensíveis. Também não esqueça de informar as variáveis de ambiente do arquivo **deployment-backend.yaml**

Espero que tenha gostado ❤️

![doguinho-fofo](./assets/doguinho.jpg)
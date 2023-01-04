# Cute dogs email

Olá, se você gosta de cachorrinhos fofos assim como eu, então irá amar essa aplicação ❤️

O cute dogs email é um encaminhador de imagem de cachorrinhos para o seu email, sendo entregue às 6:00 horário de Brasília.

Todos os créditos das imagens vão para **[Dog API](https://dog.ceo/dog-api/)**

# Configurando variáveis de ambiente

Existem dois arquivos dotenv exemplo na aplicação: um servirá para a aplicação backend e o segundo para a frontend.

O primeiro localizado no diretório raiz **/.env.example** recebe os seguintes valores

```
FROM_EMAIL= Seu email que será utilizado como remetente

FROM_PASSWORD= senha ou token referente ao email acima
```

> **_OBSERVAÇÃO:_ As variáveis abaixo podem ser preenchidas de acordo com o mongo instalado em sua máquina ou caso você opte, pela cloud do MongoDB. É possível criá-lo de forma gratuita. Para saber mais **[acesse aqui](https://cloud.mongodb.com/)**.

```
MONGO_CREDENTIALS=mongodb://127.0.0.1:27017

MONGO_DB=nome da database do mongo

MONGO_COLLECTION=nome da collection
```


# Como rodar a aplicação

O projeto pode ser executado tanto localmente quanto em fase de produção utilizando Kubernetes.

## Rodando localmente

Para executar localmente há duas formas:
Utilizando **docker** ou pelo **npm** e **go** instalados.

1. Configure as variaveis de ambiente
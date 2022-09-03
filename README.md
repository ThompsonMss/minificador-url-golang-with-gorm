# Minificador de URL feito em GoLang com GORM (PoC)

## Intro

### Um simples sistema que gera urls curtas e contabiliza os números de <b>redirects</b> daquela url, desenvolvido em GoLang utilizando o ORM [GORM](https://gorm.io/index.html).<br>
A intenção desse projeto é aprender sobre GoLang aplicando o Design Patterns Repository e realizando conexões com o banco de dados.

<br>

[Vídeo de apresentação](https://user-images.githubusercontent.com/30129295/188287101-ed4289c2-90e5-49d3-9eee-09b459612274.mp4)

-------------------------------------
<br>

## Prototipação (Figma)
[Link da prototipação (Figma)](https://www.figma.com/file/n7a8f9R2jOgcVifNFr5pEM/Untitled)

<br>

## Instalação

Siga os comandos abaixo para instalar essa aplicação em sua máquina.

> ❗ É necessário que você tenha docker e docker-compose em seu ambiente de desenvolvimento.

```bash
# Clonando o repositório da aplicação.

git clone https://github.com/ThompsonMss/minificador-url-golang-with-gorm gourl
```

```bash
# Entrando na pasta da aplicação.

cd gourl/
```

> Crie um arquivo <b><i>.env</i></b> na raiz do projeto, você pode pegar o arquivo <b><i>.env.example</i></b> como exemplo que está disponível na raiz do projeto.

```bash
# Comando para subir os conteiners do Docker

docker-compose up -d
```

> Pronto! Sua aplicação está disponível em [http://localhost:8090](http://localhost:8090)

<br>
 
 ## TODO
  - Cadastro
  - Login
  - CRUD de URLs por usuário.
  - Exibir o número de clique na url.
 <br>
 <br>

## Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://linkedin.com/in/thompson-silva)

<br>

## Licença

[MIT](https://choosealicense.com/licenses/mit/)

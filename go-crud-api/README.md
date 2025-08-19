# 🛠️ API RESTful com Golang, Clean Architecture, MongoDB e Docker

Este repositório contém o projeto desenvolvido durante o curso **"Aprenda a desenvolver uma API RESTful com Golang, Clean Architecture, MongoDB e Docker"** da Udemy. O objetivo do curso é ensinar desde os fundamentos até técnicas mais avançadas para criar APIs eficientes e escaláveis utilizando Go, MongoDB, Docker e conceitos de Clean Architecture.

## 🚀 Tecnologias Utilizadas

* **Golang**: 🖥️ Linguagem de programação para construção da API.
* **Clean Architecture**: 🏗️ Estrutura de projeto que organiza o código de forma escalável e manutenível.
* **MongoDB**: 🗃️ Banco de dados NoSQL utilizado para persistência de dados.
* **Docker**: 📦 Containerização da aplicação para facilitar o desenvolvimento e implantação.
* **Swagger (OpenAPI)**: 📄 Para documentação e testes da API.

## 📂 Estrutura do Projeto

O projeto está organizado seguindo os princípios da **Clean Architecture**, dividindo a aplicação em camadas distintas para garantir a escalabilidade e a manutenibilidade do código.

1. **Core**: 🔑 Contém os casos de uso e entidades do domínio.
2. **API**: 🌐 Implementação dos endpoints RESTful.
3. **Infrastructure**: 🔌 Conexão com o banco de dados MongoDB e integração com outras dependências externas.
4. **Docs**: 📚 Documentação da API utilizando Swagger.

## ⚙️ Como Executar o Projeto

### Pré-requisitos

* [Go](https://golang.org/dl/) versão 1.18 ou superior.
* [Docker](https://www.docker.com/) instalado para rodar os containers.
* [MongoDB](https://www.mongodb.com/) rodando localmente ou via Docker.

### Passos

1. Clone o repositório:

   ```bash
   git clone https://github.com/rogeriobgregorio/estudos-go
   cd estudos-go
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Para rodar a API localmente, use o comando:

   ```bash
   go run main.go
   ```

4. Para rodar a aplicação com Docker, utilize o comando:

   ```bash
   docker-compose up --build
   ```

5. A API estará disponível em [http://localhost:8080](http://localhost:8080).

6. A documentação da API estará disponível em [http://localhost:8080/swagger](http://localhost:8080/swagger).

## 🧪 Testes

* Para testar a API, você pode usar ferramentas como [Postman](https://www.postman.com/) ou [Insomnia](https://insomnia.rest/).
* A documentação gerada pelo Swagger também pode ser usada para testar os endpoints diretamente.


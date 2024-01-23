# Projeto simples que realiza operações básicas do CRUD para vagas de emprego

## Descrição das tecnologias utilizadas
O projeto foi desenvolvido utilizando a linguagem **Go** e o banco de dados **SQLite**. Para realizar a conexão com o banco de dados, foi utilizado o **ORM GORM**. Para realizar a documentação das rotas, foi utilizado o **Swagger**. E por fim, foi utilizado o **GIN** para realizar o roteamento das requisições.

## Executando o projeto

Para executar o projeto é necessário ter o `make` instalado na sua máquina. Após isso, acesse a raiz do projeto em **gopportunities** e execute o comando abaixo.
```bash
make
```

Para saber quais rotas e informações sobre as operações **CRUD**, acesse o **URL** abaixo.
```bash
localhost:8080/swagger/index.html
```

## TODO:
- [ ] Implementar testes unitários
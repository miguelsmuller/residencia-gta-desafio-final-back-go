![Imgur](https://i.imgur.com/j9JmM4L.png)

# **Trilha Específica - Desafio Final - Back End - GO**

## **Equipe 2 - 11/04/2022 a 14/04/2022**

- [Alexia Assumpção](https://github.com/alexiaassumpcao)
- [Carlos Lopes](https://github.com/devcarlosl)
- [Heloise Meirelles](https://github.com/Heloisemeirelles)
- [Miguel Müller](https://github.com/miguelsmuller)

## Projeto de Front-End e o Back-End principal

- [Link do Repositório de Front End](https://github.com/Heloisemeirelles/theJasonsProject)
- [Link do Repositório de Back End](https://github.com/miguelsmuller/residencia-gta-desafio-final-back)

## Planning

- [Trello Board](https://trello.com/b/KfE5ZTRF/the-jasons-projectd)

## Overview
API simples em Go com rotas para listagem de restaurantes, busca exata de restaurnate(por nome) e remoção de um restaurante.

### Dependencias
- [Go](https://go.dev/doc/)
- [Fiber](https://github.com/gofiber/fiber)
- [PG](https://github.com/lib/pq)

### _Executar somente Postgres Data Base_

```
docker-compose up --detach --force-recreate --build database
```
- Ou
```
docker-compose up --detach --force-recreate --build database pgadmin
```

### _Executar back end em  Go_
```
go run main.go
```
- Irá executar a API em: `http://localhost:3009`

# **Comando úteis**

| Comando                                  | Descrição                                        |
| ---------------------------------------- | ------------------------------------------------ |
| `docker-compose stop`                    | Interrompe os container do projeto               |
| `docker-compose down -v`                 | Apaga o conjunto removendo os volumes do projeto |
| `netstat -a -b`                          | Lista as portas abertas na máquina               |
| `docker build -t {NOME-IMAGEM} .`        | Cria uma imagem - comando para o Front End       |
| `docker run -p 8001:80 -d {NOME-IMAGEM}` | Executa uma image - comando para o Front End     |

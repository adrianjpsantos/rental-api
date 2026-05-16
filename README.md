# Rental API

API para uma plataforma de aluguel de itens desenvolvida em Go.

## Tecnologias

* Go
* Fiber v3
* PostgreSQL
* Swagger (WIP)

---

## Estrutura

```txt
internal/
 ├── application/
 ├── domain/
 ├── infrastructure/
 │    ├── http/
 │    └── persistence/
```

O projeto foi separado em camadas para manter a organização entre:

* handlers
* services
* repositories
* domínio

---

## Funcionalidades

### Usuários

* Cadastro
* Atualização
* Busca por ID
* Verificação de email e CPF

### Itens

* Cadastro de itens
* Atualização
* Busca por ID
* Busca com filtros

### Rentals

* Criação de aluguel
* Cancelamento
* Atualização de status
* Listagem por usuário

### Reviews

* Avaliações entre usuários
* Controle de review duplicado

---

## TODO

* JWT
* Upload de imagens
* Testes
* Redis
* Paginação
* Docs (Swagger)

---

Feito para estudos e evolução da arquitetura backend em Go.

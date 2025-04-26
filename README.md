# Library API server

Server implement simple RESTful API for get and set information about books and authors

# API

All ID's generating on server side, when creating new entities field `id` in json structures is ignored

## Available requests

| Relative address | Method | Value | Description |
| ---------------- | ------ | ------------ | ----------- |
| `/books` | `GET` | Array of `Book` structures | Get list of all books |
| `/books` | `POST` | `Book` structure, `id` is ignored | Add new book |
| `/books/:id` | `GET` | `Book` structure | Get info about book with id |
| `/authors` | `GET` | Array of `Author` structures | Get list of all authors |
| `/authors` | `POST` | `Author` structure, `id` is ignored | Add new author |
| `/authors/:id` | `GET` | `Author` structure | Get info about author with id |

## JSON structs

### Book
| Field name  | Descripton |
| ----------- | ---------- |
| `id` | Unique identifier |
| `authorId` | Identifier of author |
| `title` | Name of book |
| `description` | Some info about it |


### Author
| Field name | Description |
| ---------- | ----------- |
| `id` | Unique identifier |
| `firstName` | Author's name |
| `lastName` | Author's last name |

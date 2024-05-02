# Book Management System

It is a simple RESTful API for managing a collection of books written in Go.

## Gin Framework

Gin is a web framework written in Go (or Golang) which is known for its minimalistic design and high performance. It's commonly used for building web applications and RESTful APIs due to its fast routing and processing capabilities.

Usage
-----

Once the application is running, you can interact with it using HTTP requests. Here are the available endpoints:

- `GET /books`: Retrieve all books.
- `POST /books`: Add a new book to the collection.
- `GET /books/:id`: Retrieve details of a specific book by its ID.
- `PATCH /checkout?id=:id`: Check out a book by its ID.
- `PATCH /return?id=:id`: Return a book by its ID.

- **Retrieve all books**: 
  `GET /books`  

```bash
  curl localhost:8080/books
  ```

- **Add a new book** :
  `POST /books`
  ```bash
  curl localhost:8080/books --include --header "Content-Type: applicationo/json" -d @<file_name_json> --request "POST"
  ```
- **Retrieve details of a specific book**:
  `GET /books/:id`
  ```bash
  curl localhost:8080/books/1 
  ```
- **Check out a book**:
  `PATCH /checkout?id=:id`
  ```bash
  curl 'localhost:8080/checkout?id=1' --request "PATCH"
  ```
- **Return a book**:
  `PATCH /return?id=:id`
  ```bash
  curl 'localhost:8080/return?id=1' --request "PATCH"
  ```
  ### Preview of different endpoints
  
  ![Terminal Output image](https://github.com/Shreyank031/Book-Management-System/blob/master/terminal_output.png)



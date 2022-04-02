## Homework | Week 5

In this project, you can add books in SQLite and you can see them using APIs. 

Technologies which are used:
    -   GORM (SQLite Database)
    -   Fiber (For APIs)

We have 3 operations:
    - GET
    - POST
    - DELETE

## Usage

### GET Books
```
curl http://localhost:3000/book
curl http://localhost:3000/books
```
### GET Book with ID
```
curl http://localhost:3000/book/:id
```

### POST Books (Spesifics)
```
curl -X POST http://localhost:3000/book
```
### POST Books (Spesifics)
```
curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\" \"stock\": 4,\"price\": 4,\"page_number\": 4}" http://localhost:3000/books
```
### DELETE Book with ID
```
curl -X DELETE http://localhost:3000/book/:id
```
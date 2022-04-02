## Homework | Week 5

In this project, you can add books in SQLite and you can see them using APIs. 

Technologies which are used:
-   SQLite
-   GORM (SQLite Database Driver)
-   Fiber (For APIs)

We have 3 operations:
- GET
- POST
- DELETE

## Usage

You can use the program using with terminal

### GET Books
```
curl http://localhost:3000/book
```
### GET Book with ID
```
curl http://localhost:3000/book/:id
```

### POST Books (Spesifics)
```
curl -X POST http://localhost:3000/book 
```
### POST Books (Custom)

You can insert any book giving details like example shown as below

```
curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\",   \"stock\": 4,\"price\": 4,\"page_number\": 4}" http://localhost:3000/insert
```
### DELETE Book with ID
```
curl -X DELETE http://localhost:3000/book/:id
```

### Buy Book with ID
```
curl -X POST http://localhost:3000/buy/:id
```
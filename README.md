# simple_golang_restApi
Simple RESTful API to create, read, update and delete books. No database implementation 

# Quick Start
Install mux router 

go get -u github.com/gorilla/mux

go build 

./go_restapi

# Endpoints
Get All Books

GET api/books

Get Single Book

GET api/books/{id}

Delete Book

DELETE api/books/{id}

Create Book

POST api/books

# Request sample
 {
   "isbn":"4545454",
   "title":"Book Three",
   "author":{"firstname":"Harry",  "lastname":"White"}
 }
Update Book

PUT api/books/{id}

# Request sample
 {
   "isbn":"4545454",
   "title":"Updated Title",
   "author":{"firstname":"Harry",  "lastname":"White"}
}

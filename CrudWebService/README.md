# Run the web service
To run the webservice project only you need to use the command `go run main.go`. With this command the server starts listening on port 3000.

Note: This project doesn't have database, it use a slice of Users so each time when you run the program there aren't Users.

# Endpoints to test the project
 - GET `http://localhost:3000/users` with this method you will get all the users.
 - GET `http://localhost:3000/users/{id}` get a specific user. 
 - POST `http://localhost:3000/users` create a User in the database.
 ```json
{
    "FirstName" : "Santiago",
    "LastName" : "Medina"
} 
```
- PUT `http://localhost:3000/users/{id}` edit an User.
 ```json
{
    "ID" : 1,
    "FirstName" : "Santiago",
    "LastName" : "Medina"
} 
```
- DELETE `http://localhost:3000/users/{id}` Delete a User by Id.

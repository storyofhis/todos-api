# todos-api
simple web service for create read update delete (CRUD) for todos activities

## directory structure 
```
   .
    ├── common                   # Response JSON data
    ├── config                   # Configure Database Connection 
    ├── controller               # Handle Function 
    ├── docs                     # Documentation Swagger.io (alternatively `docs`)
    ├── entity                   # Specific Entiies for databases
    ├── repository               # Queries Database Postgres
    ├── router                   # Routing API
    ├── service                  # Business Logic from Repository to Controller
    ├── .env.example             # copying secret variable to connect db from .env
    ├── .gitignore               # hiding .env
    ├── go.mod                 
    ├── go.sum                   
    ├── LICENSE
    └── main.go                  # main function
```
### run todos-api 
```
go run *.go
```

### run swagger 
```
swag init -g main.go --parseDependency
```

if you have trouble with your swagger i recommend you to execute (for macOS)
```
export PATH=$(go env GOPATH)/bin:$PATH
```
see our result on swagger [link](https://editor.swagger.io)

## list endpoint

### CreateTodo : Create a list activity 
```
POST https://todos-api-production.up.railway.app/v1/todos
```

### GetTodos : Display all of your list activity 
```
GET https://todos-api-production.up.railway.app/v1/todos
```

### GetTodoByID : Display all of your list activity with specific id  
```
GET https://todos-api-production.up.railway.app/v1/todos/:id
```

### UpdateTodo : Edit and update a list activity
```
PUT https://todos-api-production.up.railway.app/v1/todos/:id
```

### DeleteTodo : Delete list from your activity 
```
DELETE https://todos-api-production.up.railway.app/v1/todos/:id
```

### Jobdesk member

- MAULA IZZA AZIZI (GLNG-KS04-020) : 
   -  Initialize Project 
   -  CreateTodo
   -  UpdateTodo
   -  GetAllTodos
   -  GetTodoByID
   -  Swaggo
   -  Deployment

- HEZKYA NATANAEL RAMLI (GLNG-KS04-008) : 
   -  CreateTodo
   -  UpdateTodo
   -  GetAllTodos
   -  GetTodoByID
   -  Mocking
   -  Unit Test

- MUHAMAD RESTU FADILLAH (GLNG-KS04-002) : 
   -  UpdateTodo
   -  DeleteTodo
   -  Swaggo
   -  Deployment
   -  Postman 
   -  Collection


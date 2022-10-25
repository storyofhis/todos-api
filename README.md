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

### list endpoint

#### CreateTodo ->Membuat todo list 
- POST 
```
https://todos-api-production.up.railway.app/v1/todos
```

GetTodos ->	  GET "/todos" -> Memunculkan seluruh todo list

GetTodoByID	-> GET "/todos/:id" -> memunculkan todo list berdasarkan ID

UpdateTodo -> PUT "/todos/:id" -> mengedit dan mengupdate todo list

DeleteTodo -> DELETE "/todos/:id" -> menghapus todo list 


### tugas anggota (team 9)

GLNG-KS04-020 - MAULA IZZA AZIZI      : Initialize Project, CreateTodo, UpdateTodo, GetAllTodos, GetTodoByID, Swaggo, Deployment

GLNG-KS04-008 - HEZKYA NATANAEL RAMLI : CreateTodo, UpdateTodo, GetAllTodos, GetTodoByID, Mocking, Unit Test

GLNG-KS04-002 - MUHAMAD RESTU FADILLAH: UpdateTodo, DeleteTodo, Swaggo, Deployment, Postman Collection


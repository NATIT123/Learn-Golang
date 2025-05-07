# 🧩 User Module - Clean Architecture Example

This module demonstrates a clean and maintainable structure following principles of Clean Architecture and Domain-Driven Design (DDD). It is designed to be easy to test, extend, and refactor.

---

## 📁 Folder Structure

user/
├── models/ # Defines data structures (entities, DTOs, types)
├── biz/ # Business logic (use cases, services)
├── storage/ # Data access layer (database interaction)
└── transport/ # Interface layer (e.g., HTTP or gRPC handlers)

## ⚙️ Responsibility Breakdown

### ✅ Models (`models/`)

Defines the core data structures used across the module.

```go
// models/user.model.go
package models

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserDto struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## ✅ Biz Layer (biz/)

Contains the business rules and core logic.

```go
// biz/user.service.go
package biz

import (
"errors"
"example.com/project/storage"
"example.com/project/models"
)

type UserService struct {
userRepo \*storage.UserRepository
}

func NewUserService(userRepo *storage.UserRepository) *UserService {
return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(data models.CreateUserDto) (*models.User, error) {
existing, err := s.userRepo.FindByEmail(data.Email)
if err != nil {
return nil, err
}
if existing != nil {
return nil, errors.New("email already in use")
}
return s.userRepo.Create(data)
}
```

## ✅ Storage Layer (storage/)

Responsible for all data persistence and querying logic.

```go
// storage/user.repository.go
package storage

import (
"example.com/project/models"
)

type UserRepository struct {
// Ideally, would connect to a database here
}

func NewUserRepository() \*UserRepository {
return &UserRepository{}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
// Simulate checking the database
// return nil if not found
return nil, nil
}

func (r *UserRepository) Create(data models.CreateUserDto) (*models.User, error) {
// Simulate creating a new user in the database
return &models.User{ID: "123", Name: data.Name, Email: data.Email}, nil
}
```

## ✅ Transport Layer (transport/)

```go
Handles communication with the outside world (HTTP, gRPC, etc.).
// transport/user.controller.go
package transport

import (
"encoding/json"
"net/http"
"example.com/project/biz"
"example.com/project/models"
)

type UserController struct {
userService \*biz.UserService
}

func NewUserController(userService *biz.UserService) *UserController {
return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
var dto models.CreateUserDto
if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
http.Error(w, "Invalid request", http.StatusBadRequest)
return
}

    user, err := c.userService.CreateUser(dto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

}
```

## 🔄 Flow: Create User

1. **[Client]** sends a request.
2. **[Transport Layer]** receives the request.
3. **[Biz Layer]** applies business rules (e.g., validation, business logic).
4. **[Storage Layer]** fetches or saves data.
5. **[Database]** persists data.
6. **[Transport Layer]** returns the response to the **[Client]**.

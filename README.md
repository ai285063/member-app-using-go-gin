# CRUD Member Api using Go, Gin framework and Docker

### Environment Requirements:
- MySQL v1.1.3
- Redis v8.11.4
- optional: Docker

### Installing steps:
1. Clone project
```shell
git clone https://github.com/ai285063/member_app_gin.git
```

2. Start testing by setting up containers using docker-compose
```shell
cd member_app_gin
docker-compose up
```

3. Test routes on Postman

### Routes:
| Method  |   | Route  |
|---|---|---|
| GET  | get all users  | [localhost:8080/users](http://localhost:8080/users)  |
| POST | create new register <br> - required fields in body: account, email, password | [localhost:8080/users/register](http://localhost:8080/users/register)  |
| PUT  | update user |  [http://localhost:8080/users/:id](http://localhost:8080/users/:id) |
| DELETE  | delete user  |  [http://localhost:8080/users/:id](http://localhost:8080/users/:id) |

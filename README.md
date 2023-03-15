
# Monorepo Auth & Fecth

## Tech Stack

**Auth:** Pyton

**Fetch:** Go

## 🔗 Links
Auth Swagger
http://95.111.195.37:8881/docs#/default/auth_controller_auth_post

Fetch Swagger
http://95.111.195.37:8888/swagger/index.html

## Running
### Auth
To run, run the following command

```bash
  cd auth
  make run
```

```http
  POST /register
    {
      "nik": "string",
      "role": "string"
    }
```
```http
  POST /login
    {
      "nik": "string",
      "password": "string"
    }
```
```http
  POST /auth
      Header
      Authorization : Bearer {{token}}
```

### Fetch
To run,  run the following command
```bash
  cd fetch
  make run
```
To Run Tes,  run the following command
```bash
  cd fetch
  make test
```

```http
  GET /api/v1/products
    Header
    Authorization : Bearer {{token}}
```

```http
  GET /api/v1/products/recommended
    Header
    Authorization : Bearer {{token}}
```

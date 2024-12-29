# go-docker-restapi-sample

golang docker REST API sample

## Technical configuration

- go
- gorm
- gorilla
- crypto
- godotenv
- mysql
- jwt

## API

### Top

|                                          | Method | URI     | Authority |
| :--------------------------------------- | :------- | :------ | :--- |
| Confirming endpoint that only returns a strings  | GET      | /api/v1 | - |

### Authentication

|                                          | Method | URI     | Authority |
| :--------------------------------------- | :------- | :------ | :--- |
| Signin | POST      | /api/v1/signin | - |
| Signup | POST      | /api/v1/signup | - |

### Todo

|                                          | Method | URI     | Authority |
| :--------------------------------------- | :------- | :------ | :--- |
| Getting all todo list associated with user| GET      | /api/v1/todo | Verified |
| Getting a single Todo associated with Todo id  | GET      | /api/v1/todo/:id | Verified |
| Creating a Todo | POST      | /api/v1/todo | Verified |
| Updating a Todo | PUT      | /api/v1/todo/:id | Verified |
| Deleting a Todo | DELETE      | /api/v1/todo/:id | Verified |

## Setting environment

### 1. Create an env file

- Create ".env" file directly under the root directory
- Copy the description of ".env.sample"

```
touch .env
```

- Go to the app directory and create a ".env" file
- Coty the description of "app/.env.sample"

```
cd app
touch .env
```

### 2. Startup docker

- Build

```
docker compose build
```

- Startup container

```
docker compose up
```

- Access go container

```
make backend-ssh
```

### 3. Preparing data

- Perform migration
- In the go container, execute the following commands.

```
make db-migrate
```

- A table will be created, so connect to the DB(MySQL) and check it.

  - "Sequel Ace" is the recommended connection application
  - https://qiita.com/ucan-lab/items/b1304eee2157dbef7774

- Perform seeding
- In the go container, execute the following commands.

```
make db-seed
```

- Data is created, so connect to the DB(MySQL) and check it.

### 4. Startup API

- Now that the data has been created in the DB, restart docker again to launch the API(first time only)

```
docker compose restart
```

- Connct to the following url and confirm that a response is returned
  - http://localhost:4000/api/v1

## Commands during development
※Nothing should be run with the API started in docker.

### Test
```
make test
```

### Static analysis
```
make lint
```

### Add go library
```
go-add-library name="[library name]"

// When specifying multiple libraries, enclose them in "", such as name="xxx yyy"
```

### To reset DB data
Since gorm does not have a rollback function, erase the entire DB with the following command
```
docker compose down -v
```

Then start docker and do the migration again to initialize the table

## docker commands

```
// build
docker-compose build

// start container
docker-compose up

// start container (background execution)
docker-compose up -d

// stop container
docker-compose down

// Stop container and delete volume (delete DB data)
docker-compose down -v

// Login to app container
docker exec -it 20211105_go_rest_server sh

// Login to DB container
docker exec -it 20211105_go_rest_db /bin/bash

```

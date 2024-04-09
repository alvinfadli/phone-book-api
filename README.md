
# Phone Book API

Phone Book API

Note : Go version 1.22.2

## Setup Database

- Run and open PostgreSQL

Create new database

```bash
  CREATE DATABASE phone_book;
```


Create new user and grant access to database

```bash
  CREATE USER alvin WITH ENCRYPTED PASSWORD '123';
  GRANT ALL PRIVILEGES ON DATABASE phone_book TO myuser;
```

P.S : You can edit the user and database on the .env file


    
## Running

Run server (default port 8080)

```bash
  go run ./
```


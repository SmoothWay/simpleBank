# Simplebank
Taken from a Udemy Course named [Backend Master Class (Golang + PostgreSQL + Kubernetes)](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/?referralCode=DD082CB0A39D22EC43EE).  
This project implements all subjects from the course, except AWS, Kubernetes, and gRPC.  
I also modified some code on API unit test and application configuration.

## Simple bank service

This project will provide APIs for the frontend to do following things:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

### Tools

- [Docker](https://docs.docker.com/engine/)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)


### Setup infrastructure

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```


- Run test:

    ```bash
    make test
    ```
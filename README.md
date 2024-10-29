# Bank\_Panda

## About bank\_panda

Its implements a functioning backend system, Restful api's and postgres database
it has api's to handle, user creation, login, account creation, getting accounts, making transfer transaction.
Uses github actions to handle CI as the complexity of the project increases.

This a growing project as it introduces more complexities of backend development, which I implement to become better at backend development.

## Runing the code locally on your machine

To have to the entire project up and running in your local machine, you need docker engine.

then use this command in the root directory:

```bash
docker compose up --build
```

There are currently three resources exposed using rest api endpoints:

- users, two post endpoints to create users and users/login to login
- accounts, one post endpoint to create an account and a get accounts/{id} endpoint to get and account
- transfers, one post endpoint to create a transfer

After login, access to resouces is authorized and authenticated using web tokens

## Future Updates

Using grpc, and an api gate way to serve but rest and grpc endpoints

# Beer API

Challenge beer API

## Run project 🚀

```bash
export CURRENCY_API_SECRET=YourKeySecret
export POSTGRES_PASSWORD=password
Docker-compose up
```

This command creates and raises the api for the port `8080` and a database postgres for the port `5432`, you can consume the service from the following url [http://localhost:8080](http://localhost:8080), for more information about accesses check [docker-compose.yml](./docker-compose.yml).

## Docs

If you want to generate the swagger documentation, execute the following command `make doc` but if you want to view it, you just have to have the service above and enter the following url
[http://localhost:8080/swagger/index.html#/](http://localhost:8080/swagger/index.html#/)

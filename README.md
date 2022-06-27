# ShortLinksService
Links shortener with Golang

## Endpoints

- ### `/create` - generate new short link
  **Request**: `POST`
    ```json
    { "long_url": "https://" }
    ```
  **Response**:
    ```json
    {
        "short_url": "http://localhost:8080/"
    }
    ```

- ### `/url/{short_url}` - get short link
  **Request**: `GET`  
  **Response**:
    ```json
    {
        "long_url": "https://stepik.org/catalog?auth=registration"
    }
    ```
- ### `/{short_url}` - auto redirect
  **Request**: `GET`  
---


### Environmental variables
1. `HTTP_PORT`: default `8080`
2. `DB_HOST`: default = `ec2-52-212-228-71.eu-west-1.compute.amazonaws.com`
3. `DB_PORT`: default = `5432`
4. `DB_USERNAME`: default = `fzcamrgntritxl`
5. `DB_NAME`: default = `d7r6s29qsbs7ah`
6. `DB_PASSWORD`: default = `test`
7. `DB_SSL_MODE`: default = `require`
8. `MODE`: default = `postgres`

### Build

To build docker image:
```
    make build-image
```

To start service in docker container:
```
    make docker-start
```

To change working mode from `postgres` to `internal`:

```
docker run --env MODE=internal  --name=links-service --publish 8080:8080 --rm links-service;
```

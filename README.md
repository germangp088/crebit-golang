# crebit-golang
Transactional work

# Financial transactions logic

Web service money account system to perform transactions logic (debit and credit).

  - Credit (+)
  - Debit (-)

## Run on Docker
You can run the project using docker to avoid install dependencies, packages or node, the following commands will allow you to get the project running on a virtualize container.

# Run docker
```sh
$ docker-compose build
$ docker-compose up
```

### How to use it

You could use Postman to create and consume transactions throgh Rest API.

Client runs on localhost:3000
Api runs on localhost:8080


###### ENDPOINTS

-- GET /{URL}/transactions

-- GET /{URL}/transactions/:id

-- POST /{URL}/transactions

-- GET /{URL}/balance

**Transaction Body Example**
```json
{
  "type": "credit",
  "amount": 8000
}
```
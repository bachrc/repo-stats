# Canvas for Backend Technical Test at Scalingo

## Instructions

* From this canvas, respond to the project which has been communicated to you by our team
* Feel free to change everything

## Execution

```
docker-compose up
```

Application will be then running on port `5000`

## Test

```
$ curl localhost:5000/ping
{ "status": "pong" }
```

# What did this person do ?

This API will show you insights of the repositories of a user.

What we will test :

- Récupérer la liste des dépots d'un utilisateur
- Permettre de limiter sur :
  - Le langage
  - La license
- 

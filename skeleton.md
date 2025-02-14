# e-commerce

## requisites

create a web-app to manage an e-commerce

List of available products
- simple, straightforward

Shopping cart management
- add / remove products from cart
- show/change amount in cart
- show total price

Database
- save available products (and their amount)
- save user data?
- save purchase info

Containerization
- make Docker containers
- make docker-compose to manage the container set

Additionally:
- JWT authenitcation
- use React's Context API to manage app state (frontend)

## Architecture



### Frontend
### Backend

Im using GO. Don't care about python.

Go chi for api routing
sqlc for building go-mysql queries

REST API for each data type probably...

### Database

goose to manage migrations
sqlc to generate go-mysql code
mysql for DB access

### Containerization

docker and docker-compose to deal with this...

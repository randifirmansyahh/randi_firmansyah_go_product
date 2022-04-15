# Randi Firmansyah

## Getting started

Welcome!

## To run this project you can type text below on the terminal / bash :

### go run .

# LOGIN

- POST http://localhost:5000/login

```json
{
  "id": 1,
  "username": "randi"
}
```

# Generate token without login

- GET http://localhost:5000/globaltoken

# USER (Login Required)

- GET http://localhost:5000/user
- GET http://localhost:5000/user/{id}
- DELETE http://localhost:5000/user/{id}
- PUT http://localhost:5000/user/{id}

```json
{
  "nama": "Randi Firmansyah",
  "username": "randifirmansyahh",
  "password": "randi123!",
  "email": "randykelvin26@gmail.com",
  "no_hp": "0854545454",
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

- POST http://localhost:5000/user

```json
{
  "nama": "Randi Firmansyah",
  "username": "randifirmansyahh",
  "password": "randi123!",
  "email": "randykelvin26@gmail.com",
  "no_hp": "0854545454",
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

# PRODUCT (Login Required)

- GET http://localhost:5000/product
- GET http://localhost:5000/product/{id}
- DELETE http://localhost:5000/product/{id}
- PUT http://localhost:5000/product/{id}

```json
{
  "nama": "Ice Cream",
  "harga": 20000,
  "qty": 12,
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

- POST http://localhost:5000/user

```json
{
  "nama": "Ice Cream",
  "harga": 20000,
  "qty": 12,
  "image": "https://1757140519.rsc.cdn77.org/blog/wp-content/uploads/2018/05/1-google-logo.png"
}
```

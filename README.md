# Modern Version Of Pokedex

## Demo app of implementing DevOps solutions like docker image and Github Actions

### Technologies

- Golang
- Docker
- Docker Compose
- PostgreSQL
- Redis
- Git
- Github Actions

### Results times

| Command | Time | Cached |
| --- | --- | --- |
| http :8080/2 | 598.417µs | false |
| http :8080/2 | 119.962µs | true |
| http :8080/Mew | 489.174µs | false |
| http :8080/Mew | 140.2µs| true |
| http :8080 | 586.181µs| false |
| http :8080 | 443.069µs| true |

### How to run

```bash
make run
```

### How to run tests

```bash
make test
```

### How to clean

```bash
make clean
```

### After Running There Are Api Endpooints

#### GET all pokemons

```bash
http :8080 -v
```

#### GET pokemon by id

```bash
http :8080/:id -v
```

#### GET pokemon by name

```bash
http :8080/:name -v
```

##### -v flag for redis cache headers

#### Example

##### Without cache

```bash
http :8080/Bulbasaur -v
```

```http
HTTP/1.1 200 OK
Content-Length: 60
Content-Type: application/json; charset=utf-8
Date: Wed, 26 Nov 2025 21:50:25 GMT
X-Cache: MISS

{
    "ID": 1,
    "Name": "Bulbasaur",
    "Type1": "Grass",
    "Type2": "Poison"
}
```

##### With cache

```bash
http :8080/4 -v
```

```http
HTTP/1.1 200 OK
Content-Length: 54
Content-Type: application/json; charset=utf-8
Date: Wed, 26 Nov 2025 21:51:52 GMT
X-Cache: HIT

{
    "ID": 4,
    "Name": "Charmander",
    "Type1": "Fire",
    "Type2": ""
}
```

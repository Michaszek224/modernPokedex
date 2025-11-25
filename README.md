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
curl http://localhost:8080/
```

#### GET pokemon by id

```bash
curl http://localhost:8080/:id
```

#### GET pokemon by name

```bash
curl http://localhost:8080/:name
```

#### Example

```bash
curl http://localhost:8080/Bulbasaur
```

```json
{
    "ID": 1,
    "Name": "Bulbasaur",
    "Type1": "Grass",
    "Type2": "Poison"
} 
```

```bash
curl http://localhost:8080/4
```

```json
{
    "ID": 4,
    "Name": "Charmander",
    "Type1": "Fire",
    "Type2": ""
}
```

# weaver-pokemon-type

## Init

``` bash
make init
direnv allow
```

## Run

``` bash
docker compose up -d
make seed
make
./dist/server
```

## Dev

Run `weaver generate` and other code generations.

``` bash
make generate
```

## API

### /type

``` bash
❯ curl "localhost:21099/type" -d '{"name":"エスパー"}'
{"item":{"id":11,"name":"エスパー"}}
```

### /attack

``` bash
❯ curl "localhost:21099/attack" -d '{"id":11,"index":1}'|jq '.items[0]'
{
  "pile": [
    {
      "id": 181,
      "attack": {
        "id": 11,
        "name": "エスパー"
      },
      "defense": {
        "id": 1,
        "name": "ノーマル"
      },
      "multiplier": 1
    }
  ],
  "multiplier": 1
}
```

``` bash
❯ curl "localhost:21099/attack" -d '{"id":11,"index":2}'|jq '.items[0]'
{
  "pile": [
    {
      "id": 181,
      "attack": {
        "id": 11,
        "name": "エスパー"
      },
      "defense": {
        "id": 1,
        "name": "ノーマル"
      },
      "multiplier": 1
    },
    {
      "id": 182,
      "attack": {
        "id": 11,
        "name": "エスパー"
      },
      "defense": {
        "id": 2,
        "name": "ほのお"
      },
      "multiplier": 1
    }
  ],
  "multiplier": 1
}
```

### /defense

``` bash
❯ curl "localhost:21099/defense" -d '{"ids":[11]}'|jq '.items[0]'
{
  "pile": [
    {
      "id": 119,
      "attack": {
        "id": 7,
        "name": "かくとう"
      },
      "defense": {
        "id": 11,
        "name": "エスパー"
      },
      "multiplier": 0.5
    }
  ],
  "multiplier": 0.5
}
```

``` bash
❯ curl "localhost:21099/defense" -d '{"ids":[11,7]}'|jq '.items[0]'
{
  "pile": [
    {
      "id": 65,
      "attack": {
        "id": 4,
        "name": "でんき"
      },
      "defense": {
        "id": 11,
        "name": "エスパー"
      },
      "multiplier": 1
    },
    {
      "id": 61,
      "attack": {
        "id": 4,
        "name": "でんき"
      },
      "defense": {
        "id": 7,
        "name": "かくとう"
      },
      "multiplier": 1
    }
  ],
  "multiplier": 1
}
```

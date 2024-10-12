# Greenlight

Me working through the book ["Let's Go Further"][letsgof] by [Alex Edwards](https://www.alexedwards.net/).

## Prerequisites

- [Go 1.23](https://go.dev/) or later
- [Air](https://github.com/air-verse/air) for live reloading:
  ```shell
  go install github.com/air-verse/air
  ```
- [Docker](https://www.docker.com/) for running Postgres in a container

## Development

```shell
# Start the development server (Linux / MacOS)
air

# Windows
air -c .air-win.toml
```


[letsgof]: https://lets-go-further.alexedwards.net/ "Let's Go Further"

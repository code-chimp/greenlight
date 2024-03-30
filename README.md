# Greenlight

It is based on the book ["Let's Go Further"][letsgof] by [Alex Edwards](https://www.alexedwards.net/).

## Prep the Database

Create a Docker instance of MariaDB for development purposes:

```shell
docker run --name greenlight_dev \
    -p 5432:5432 \
    -e 'POSTGRES_PASSWORD=P@ssw0rd' \
    -d postgres:14
```

**Seed DB**?

```postgresql

```

[letsgof]: https://lets-go-further.alexedwards.net/ "Let's Go Further"

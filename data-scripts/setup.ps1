#!/usr/bin/env pwsh

docker run --name greenlight_dev `
    -p 5432:5432 `
    -e 'POSTGRES_PASSWORD=P@ssw0rd' `
    -d postgres:14

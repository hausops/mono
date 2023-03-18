# dashboard-api

github.com/hausops/mono/apps/dashboard-api

Providing data access (BFF) to the Dashboard.

Resolvers should simply delegates to services. No business logic in the API; they belong in the services.

## Development

Copy `.env.example` to `.env`

```sh
# from apps/dashboard-api
make run
```

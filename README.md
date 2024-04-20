# workflou/template

A simple project template using [Echo](http://echo.labstack.com), [HTMX](http://htmx.org), [AlpineJS](https://alpinejs.dev), [Tailwind](http://tailwindcss.com) and [Air](https://github.com/cosmtrek/air).

## Local development

Copy `.env.example` to `.env` and customize environment variables. 

```
make dev
```

This will run a local development server on port `4000` with hot reloading.

## Database migrations

### Create a new migration file

```
make make-migration
```

### Run migrations

```
make migrate
```

This will run migrations from `/migrations` directory.

### Reset database

```
make reset
```

## Adding handlers

Create a package for your domain, e.g. `auth`, `home`, `blog`, `billing`, `store`, etc. Create a new `handler.go` with a new struct. Pass stores, mailers and validators as a dependency when creating a new instance of a handler. See `home/handler.go` for reference.

## Adding stores

Similarly, create `store.go` with a struct which accepts `*sqlx.DB` as a dependency. Create a method for each query.

## Running demo app

```
make migrate
make testuser
make
```
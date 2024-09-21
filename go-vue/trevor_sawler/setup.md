# Adding `Vetur` extension to VSCode

- Vetur (Vue tooling for VS Code)

## Installing `make`

- [MacOS] `brew install make`

## Installing Vue (on `.html`)

- `https://getbootstrap.com/docs/5.1/getting-started/download/`
  - Copy CDN
  ```html
  <link
    href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css"
    rel="stylesheet"
    integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3"
    crossorigin="anonymous"
  />
  <script
    src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"
    integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p"
    crossorigin="anonymous"
  ></script>
  ```
- `https://www.jsdelivr.com/`

## Installing Vue-CLI

- `npm install -g @vue/cli`
  - To check if vue-cli is installed, type `vue --version`
- `vue create <project_name>`

## `notie`

- For custom error messages alert box
- `npm install notie`
- Add to `index.html`
  - `<link rel="stylesheet" type="text/css" href="https://unpkg.com/notie/dist/notie.min.css" />`

# Golang

- chi: `go get -u github.com/go-chi/chi/v5`
- CORS: `go get github.com/go-chi/cors`
- jackc/pgx: `go get github.com/jackc/pgconn` and `go get github.com/jackc/pgx/v4`
- DATADOG: `go get github.com/DATA-DOG/go-sqlmock`

- Running with env, but troublesome: `env DSN="host=localhost port=5434 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5" go run ./cmd/api`

## Database

- Used Beekeeper Studio
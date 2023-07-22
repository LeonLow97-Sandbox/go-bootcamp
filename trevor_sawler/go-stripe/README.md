# Setup

- go mod init myapp
- `go get github.com/go-chi/chi/v5`
- `go run ./cmd/web`

# Hot Reloading

- `https://github.com/cosmtrek/air`
- `curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin`
- `sudo air`

## Install stripe

- `https://github.com/stripe/stripe-go`
- `go get -u github.com/stripe/stripe-go/v72`

## CORS

- `https://github.com/go-chi/cors`
- `go get github.com/go-chi/cors`

## Adding `Makefile`

- Hot reloading for both frontend and backend
- Ensure `Makefile` is in root directory
- `make start_back`

## Stripe Test Error Code

- `https://stripe.com/docs/testing`
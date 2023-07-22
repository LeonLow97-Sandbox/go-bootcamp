# setup

- go mod init myapp
- `go get github.com/go-chi/chi/v5`
- `go run ./cmd/web`

# Hot Reloading

- `https://github.com/cosmtrek/air`
- `curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin`
- `sudo air`
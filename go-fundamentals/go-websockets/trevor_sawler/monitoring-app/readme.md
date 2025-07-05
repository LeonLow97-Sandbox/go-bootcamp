# Vigilate

This is the source code for the second project in the Udemy course Working with Websockets in Go (Golang).

A dead simple monitoring service, intended to replace things like Nagios.

## Build

Build in the normal way on Mac/Linux:

```
go build -o vigilate cmd/web/*.go
```

Or on Windows:

```
go build -o vigilate.exe cmd/web/.
```

Or for a particular platform:

```
env GOOS=linux GOARCH=amd64 go build -o vigilate cmd/web/*.go
```

## Requirements

Vigilate requires:

- Postgres 11 or later (db is set up as a repository, so other databases are possible)
- An account with [Pusher](https://pusher.com/), or a Pusher alternative
  (like [ipê](https://github.com/dimiro1/ipe))

## Run

First, make sure ipê is running (if you're using ipê):

On Mac/Linux

```
cd ipe
./ipe
```

On Windows

```
cd ipe
ipe.exe
```

Run with flags:

```
./vigilate \
-dbuser='tcs' \
-pusherHost='localhost' \
-pusherPort='4001' \
-pusherKey='123abc' \
-pusherSecret='abc123' \
-pusherApp="1" \
-pusherSecure=false
```

## All Flags

```
tcs@grendel vigilate-udemy % ./vigilate -help
Usage of ./vigilate:
  -db string
        database name (default "vigilate")
  -dbhost string
        database host (default "localhost")
  -dbport string
        database port (default "5432")
  -dbssl string
        database ssl setting (default "disable")
  -dbuser string
        database user
  -domain string
        domain name (e.g. example.com) (default "localhost")
  -identifier string
        unique identifier (default "vigilate")
  -port string
        port to listen on (default ":4000")
  -production
        application is in production
  -pusherApp string
        pusher app id (default "9")
  -pusherHost string
        pusher host
  -pusherKey string
        pusher key
  -pusherPort string
        pusher port (default "443")
  -pusherSecret string
        pusher secret
   -pusherSecure
        pusher server uses SSL (true or false)
```

## PG

```sql
-- password
psql -U postgres

GRANT ALL ON DATABASE vigilate TO leon;
ALTER DATABASE vigilate OWNER TO leon;

\c vigilate;

-- to view all tables
\dt;
```

```
# Run to setup database
soda migrate
./run.sh

# Application Account (localhost:4000)
U: admin@example.com
P: password
```

## Pusher Service

- [GitHub Pusher Service](https://github.com/dimiro1/ipe)
- Why use Pusher Service?
  - Pusher is very popular, and there are many companies that hire people with this skill set
  - Pusher makes it easy to work with WebSockets
  - Moving the WebSockets load to a different application server (or even host) spreads the load around
- Pusher Client is in `js.jet` file
- Pusher Server has `app.WsClient.Trigger` to send to clients who are subscribed to this event

```js
// Pusher Client
<script src="/static/admin/js/pusher.min.js"></script>

<script>
    let pusher = new Pusher("{{.PreferenceMap["pusher-key"]}}", {
        authEndPoint: "/pusher/auth",
        wsHost: "localhost",
        wsPort: 4001,
        forceTLS: false,
        enabledTransports: ["ws", "wss"],
        disabledTransports: []
    });

    let publicChannel = pusher.subscribe("public-channel");

    publicChannel.bind("app-starting", function (data) {
        successAlert(data.message);
    })

    publicChannel.bind("app-stopping", function (data) {
        warningAlert(data.message);
    })
```

```go
// Pusher Server
data := make(map[string]string)
data["message"] = "Monitoring is starting..."               // message pushed to all clients
err := app.WsClient.Trigger("public-channel", "app-starting", data) // push to all clients in public channel
if err != nil {
      log.Println(err)
}
```

## Private Channel

- Must begin with "private" in the frontend

```js
// private channel must start with "private"
let privateChannel = pusher.subscribe('private-channel-{{.User.ID}}');

privateChannel.bind('private-message', function (data) {
  attention.alert({
    html: data.message,
    icon: 'success',
  });
});
```

## Start Mailhog

- `brew services start mailhog`
- `localhost:8025`
- `brew stop mailhog`

## Insert other services

```sql
INSERT INTO services (service_name, active, icon, created_at, updated_at)
VALUES ('HTTPS', 1, 'fas fa-server', NOW(), NOW());

INSERT INTO services (service_name, active, icon, created_at, updated_at)
VALUES ('SSL', 1, 'fas fa-server', NOW(), NOW());
```

## SSL Certificate

- `https://github.com/tsawler/checkhttp2`
- `git clone https://github.com/tsawler/checkhttp2`
  - `go build -o checkhttp2 main.go`
  - `./checkhttp2 -host google.com -cert`

## How to Run?

- `cd monitoring-app`
  - `./run.sh`
- `cd monitoring-app/ipe`
  - `./ipe`
- `cd hello-world-web` (For HTTP service)
  - `go run main.go`

## Final Source Code

- [Source Code](https://github.com/tsawler/vigilate/releases/tag/v59)

package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// middleware for rate limiting requests
func (app *application) rateLimit(next http.Handler) http.Handler {
	// create a client type
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	// create a rate limiter for the client
	var (
		cleanMu   sync.Mutex
		ipCheckMu sync.Mutex
		// key is IP Address of the client
		clients = make(map[string]*client)
	)

	// launch a background goroutine that removes old
	// entries from the `clients` map once every minute
	go func() {
		for {
			time.Sleep(time.Minute)
			// Lock before starting to clean
			cleanMu.Lock()
			for ip, client := range clients {
				// if last seen more than 3 minutes, remove this ip from the clients map
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			cleanMu.Unlock()
		}
	}()

	// 1st argument: how many requests per second?
	// limit := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the IP address of the request
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err)
			return
		}

		// Lock()
		ipCheckMu.Lock()
		// check if the IP address is in the map
		if _, found := clients[ip]; !found {
			clients[ip] = &client{
				limiter: rate.NewLimiter(2, 4),
			}
		}

		// update the last seen time of the client
		clients[ip].lastSeen = time.Now()

		// check if request allowed
		if !clients[ip].limiter.Allow() {
			ipCheckMu.Unlock()
			app.rateLimitExceededResponse(w)
			return
		}
		ipCheckMu.Unlock()

		next.ServeHTTP(w, r)
	})
}

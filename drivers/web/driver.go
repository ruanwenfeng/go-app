// Package web is the driver to be used for web applications.
package web

import (
	"net/http"

	"github.com/murlokswarm/app"
)

// Driver is an app.Driver implementation for web.
type Driver struct {
	// The URL of the component to load when a navigating on the website root.
	DefaultURL string

	// The URL of the component to load when a 404 errors occurs.
	// Default is /web.NotFound
	NotFoundURL string

	// The server used to save request.
	// Default is a server that listens on port 7042.
	Server *http.Server

	factory     app.Factory
	fileHandler http.Handler
}

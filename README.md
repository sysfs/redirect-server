## This program redirects all web traffic to a website you specify.

**How to build**

`go build redirect-server.go`

**How to use**

`./redirect-server --redirect-url=https://go.dev` (replace with the actual website URL of your choice)

This will start a webserver on http://localhost:8080 which redirects all visits to your chosen website.

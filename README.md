<!-- vim: set et ts=2 sw=2 sts=2 tw=79 cc=80 list listchars=trail\:·: -->
## This program redirects all web traffic to a URL you specify.

**How to build**

`go build redirect-server.go`

**How to use**

`./redirect-server --redirect-url=https://go.dev` (replace with the actual URL of your choice)

This will start a webserver on port 8080 which redirects all visits to your chosen URL.

Alternatively you can set the environment variable REDIRECT_URL before running instead of
using a commandline parameter.

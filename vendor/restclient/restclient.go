package restclient

import (
	"net/http"
)

var (
	client    *http.Client
	Adres     string
	transport *http.Transport
)

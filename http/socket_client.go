package http

import (
	"crypto/tls"
	"net/http"
	"context"
	"appengine/socket"
	"net"
)

func NewSocketClient(ctx context.Context) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Dial: func(network, addr string) (net.Conn, error) {
			return socket.Dial(ctx, network, addr)
		},
	}
	client := &http.Client{
		Transport: transport,
	}
	return client
}
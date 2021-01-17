package headerstransformation

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/rwxrob/uniq"
)

const defaultHeader = "X-Traefik-Uuid"

type Config struct {
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

type HeadersTransformation struct {
	next       http.Handler
	headerName string
	name       string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.HeaderName) == 0 {
		return nil, fmt.Errorf("HeaderName cannot be empty")
	}
	return &HeadersTransformation{
		next:       next,
		headerName: config.HeaderName,
		name:       name,
	}, nil
}

func (ht *HeadersTransformation) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	uid := uniq.UUID()
	req.Header.Set(ht.headerName, uid)
	rw.Header().Add(ht.headerName, uid)
	ht.next.ServeHTTP(rw, req)
}

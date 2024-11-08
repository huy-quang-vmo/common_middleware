package common_middleware

import "net/http"

type MiddlewareHandler interface {
	Handle(next http.Handler) http.Handler
}

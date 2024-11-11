package maintenance

import "net/http"

type MiddlewareHandler interface {
	Handle(next http.Handler) http.Handler
}

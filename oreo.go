package oreo

import (
	"fmt"
	"net/http"
)

type OreoRouter struct {
	routes map[string]http.HandlerFunc
}

// New Router function .
// For creating new instance of OreoRouter
func Oreo() *OreoRouter {
	return &OreoRouter{
		routes: make(map[string]http.HandlerFunc),
	}
}

// Request a path
// e.g. /login
// Handler is ServerHTTP
func (r *OreoRouter) Req(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *OreoRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for path, handler := range r.routes {
		if req.URL.Path == path {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// Listening and serving on port
// Result ensures that serving is done
// Handler might be nil
func (r *OreoRouter) Listen(port string, handler http.Handler, result string) {
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		panic("Problem in serving : " + err.Error());
	}
	fmt.Println(result);
}

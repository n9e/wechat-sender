package middleware

import (
	"net/http"
	"runtime"

	"github.com/n9e/wechat-sender/http/render"

	"github.com/toolkits/pkg/errors"
	"github.com/toolkits/pkg/logger"
)

// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
type Recovery struct {
	StackAll  bool
	StackSize int
}

// NewRecovery returns a new instance of Recovery
func NewRecovery() *Recovery {
	return &Recovery{
		StackAll:  false,
		StackSize: 1024 * 8,
	}
}

func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(errors.PageError); ok {
				// custom error
				// if e.Error() == "unauthorized" {
				// 	http.Redirect(w, r, "/?callback="+r.RequestURI, 302)
				// 	return
				// }
				render.Message(w, e)
				return
			}

			// Negroni part
			w.WriteHeader(http.StatusInternalServerError)
			stack := make([]byte, rec.StackSize)
			stack = stack[:runtime.Stack(stack, rec.StackAll)]

			logger.Errorf("PANIC: %s\n%s", err, stack)
		}
	}()

	next(w, r)
}

func isAjax(r *http.Request) bool {
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

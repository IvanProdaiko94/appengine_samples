package http

import (
	"context"
	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

type Auth struct {
	AdminOnly bool
	RedirectTo string
}

type Handler struct {
	Method string
	Params []string
	Auth *Auth
	OnRequest func(ctx context.Context, w http.ResponseWriter, r *http.Request, handler *Handler)
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ctx = appengine.NewContext(r)
	if !MethodCheck(w, r, handler.Method, ctx) {
		return
	}
	if handler.Params != nil && len(handler.Params) > 0 {
		var params url.Values
		if r.Method == http.MethodPost || r.Method == http.MethodPatch || r.Method == http.MethodPut {
			if err := r.ParseForm(); err != nil {
				HandleError(ctx, w, err, http.StatusBadRequest)
				return
			}
			params = r.PostForm
		} else if r.Method == http.MethodGet {
			params = r.URL.Query()
		}
		if !ParamsCheck(w, handler.Params, params, ctx) {
			return
		}
	}
	if handler.Auth != nil {
		u := user.Current(ctx)
		if !AuthCheck(w, r, ctx, u, *handler.Auth) {
			return
		}
	}
	handler.OnRequest(ctx, w, r, handler)
}

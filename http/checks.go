package http

import (
  "context"
  "errors"
  "fmt"
  "net/http"
  "net/url"
  "strings"

  "google.golang.org/appengine/log"
  "google.golang.org/appengine/user"
)

func MethodCheck(w http.ResponseWriter, r *http.Request, method string, ctx context.Context) (next bool) {
  if r.Method != method {
    w.Header().Set("Access-Control-Allow-Methods", method)
    err := fmt.Errorf("method not allowed. Expected %s, but got %s", method, r.Method)
    HandleError(ctx, w, err, http.StatusBadRequest)
    return false
  }
  return true
}

func ParamsCheck(w http.ResponseWriter, params []string, qs url.Values, ctx context.Context) (next bool) {
  for _, param := range params {
    if _, ok := qs[param]; !ok {
      err := fmt.Errorf("parameters \"%s\" are required", strings.Join(params, ", "))
      HandleError(ctx, w, err, http.StatusBadRequest)
      return false
    }
  }
  return true
}

func AuthCheck(w http.ResponseWriter, r *http.Request, ctx context.Context, u *user.User, config Auth) bool {
  if u == nil {
    err := errors.New("user not exists")
    log.Errorf(ctx, err.Error())
    if config.RedirectTo != "" {
      redirectUrl, err := user.LoginURL(ctx, config.RedirectTo)
      if err != nil {
        HandleError(ctx, w, err, http.StatusForbidden)
        return false
      }
      http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
    } else {
      HandleError(ctx, w, err, http.StatusForbidden)
    }
    return false
  }
  if config.AdminOnly && !u.Admin {
    err := errors.New("not authorized")
    HandleError(ctx, w, err, http.StatusForbidden)
    return false
  }
  log.Infof(ctx, "Current user is administrator: %s", u)
  return true
}
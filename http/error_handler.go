package http

import (
  "context"
  "net/http"

  "google.golang.org/appengine/log"
)

func HandleError(ctx context.Context, w http.ResponseWriter, err error, status int) {
  log.Errorf(ctx, err.Error())
  http.Error(w, err.Error(), status)
}

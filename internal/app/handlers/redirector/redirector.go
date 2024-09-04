package redirector

import (
	"github.com/beaverkiria/urlshortener/internal/app"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	redirectPath := r.PathValue("id")
	redirectTo, err := app.GetRedirect(redirectPath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", redirectTo)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

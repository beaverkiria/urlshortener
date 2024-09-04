package shortener

import (
	"fmt"
	"github.com/beaverkiria/urlshortener/internal/app"
	"io"
	"net/http"
	"net/url"
)

const shortenUrlLength = 10
const shortenUrlCharset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
const siteName = "localhost:8080"

func ShortenLink(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	shortenedUrl := app.RandomStringWithCharset(shortenUrlLength, shortenUrlCharset)
	_ = app.SaveRedirect(string(body), shortenedUrl)

	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte(fmt.Sprintf("%s/%s", siteName, shortenedUrl)))
}

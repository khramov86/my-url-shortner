package app

import (
	"bytes"
	"fmt"
	"net/http"

	"log/slog"

	"math/rand"

	"github.com/khramov86/my-url-shortner/internal/config"
)

type URLMap map[string]string

var urlMap = make(URLMap)

func generateID() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, 8)
	for idx := range id {
		id[idx] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

func Init(cfg *config.Config) {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port), nil)
}

func getScheme(r *http.Request) string {
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		return proto
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Path[1:]
		slog.Info(fmt.Sprintf("ID: %v", id))
		if url, ok := urlMap[id]; ok {
			w.Header().Set("Location", url)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}
		http.NotFound(w, r)

	case http.MethodPost:
		url := r.Body
		defer url.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(url)
		urlStr := buf.String()
		if urlStr == "" {
			http.Error(w, "Empty URL", http.StatusBadRequest)
			return
		}
		id := generateID()
		urlMap[id] = urlStr
		slog.Info(fmt.Sprintf("Short URL: %v", urlMap))
		// TODO: подумать, как лучше формировать хост, из запроса или из конфига
		host := r.Host
		scheme := getScheme(r)
		w.Write([]byte(fmt.Sprintf("%s://%s/%s", scheme, host, id)))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

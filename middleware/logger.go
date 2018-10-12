package middleware

import (
	"github.com/listmera/frank/env"
	"github.com/listmera/frank/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func logger(h http.Handler) http.Handler {
	PADDING := "================================"
	consumeBody := env.GetOr("GO_ENV", "development") == "development"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc, err := httputil.DumpRequest(r, consumeBody)
		utils.CheckErr(err)
		log.Printf("\n%s\n%s requested %s:\n    %s\n%s", PADDING, r.RemoteAddr, r.URL, string(rc), PADDING)

		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, r)

		wc, err := httputil.DumpResponse(rw.Result(), consumeBody)
		utils.CheckErr(err)

		log.Printf("\n%s\nfrank replied to %s:\n    %s\n%s", PADDING, r.RemoteAddr, string(wc), PADDING)

		for k, v := range rw.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		body := rw.Body.Bytes()

		w.Write(body)
	})
}
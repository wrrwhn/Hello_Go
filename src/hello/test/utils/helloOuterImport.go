package utils

import (
	// urlshortener "google.golang.org/api/urlshortener/v1"
	// urlshortener "github.com/google/google-api-go-client/urlshortener/v1"
	"log"
	// "net/http"
)

func HelloUrlShorter(url string) {
	// svc, _ := urlshortener.New(http.DefaultClient)
	// sUrl, _ := svc.Url.Insert(&urlshortener.Url{LongUrl: url}).Do()
	// lUrl, _ := svc.Url.Get(sUrl).Do()

	sUrl := url
	lUrl := url
	log.Println(sUrl, lUrl)
}

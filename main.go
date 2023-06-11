package main

import (
	"net/http"

	"github.com/r-leafar/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	//http.ListenAndServeTLS("0.0.0.0:8000", "cert.pem", "key.pem", nil)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

package main

import (
	"net/http"
	"log"
	"crypto/tls"
	"io/ioutil"
	"crypto/x509"
	"path/filepath"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from test server.\n"))
}

func handleError(err error) {
	if err != nil {
		log.Fatal("Fatal", err)
	}
}

func main() {
	absPathServerCrt, err := filepath.Abs("certs/server.crt")
	handleError(err)
	absPathServerKey, err := filepath.Abs("certs/server.key")
	handleError(err)


	clientCACert, err := ioutil.ReadFile(absPathServerCrt)
	handleError(err)

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: clientCertPool,
		PreferServerCipherSuites: true,
		MinVersion: tls.VersionTLS12,
	}

	tlsConfig.BuildNameToCertificate()


	http.HandleFunc("/", HelloServer)
	httpServer := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	err = httpServer.ListenAndServeTLS(absPathServerCrt,absPathServerKey)
	handleError(err)
}

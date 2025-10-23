package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(c *gin.Context) {
	if c.Request.TLS != nil && len(c.Request.TLS.PeerCertificates) > 0 {
		clientCert := c.Request.TLS.PeerCertificates[0]

		c.JSON(200, gin.H{
			"subject":    clientCert.Subject,
			"issuer":     clientCert.Issuer,
			"serial":     clientCert.SerialNumber.String(),
			"not_before": clientCert.NotBefore,
			"not_after":  clientCert.NotAfter,
		})
		return
	}

	c.JSON(401, gin.H{"error": "no client certificate provided"})
}

func main() {
	ginServer := gin.New()

	// Load CA cert
	caCert, err := os.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Fatalln("fatal: can't find root CA certificate")
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	server := &http.Server{
		Addr:    "0.0.0.0:443",
		Handler: ginServer,
		TLSConfig: &tls.Config{
			ClientCAs:  caPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	// Add routing
	ginServer.GET("/profile", GetProfileHandler)

	err = server.ListenAndServeTLS("./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Fatal(err)
	}
}

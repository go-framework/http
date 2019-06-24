package http

import (
	"crypto/tls"
	"encoding/pem"

	"golang.org/x/crypto/pkcs12"
)

// PKCS12 convert to Pem
// https://blog.fish2bird.com/?p=555
func PKCS12ToPem(p12Data []byte, password string) (cert tls.Certificate, err error) {

	var blocks []*pem.Block
	blocks, err = pkcs12.ToPEM(p12Data, password)
	if err != nil {
		return
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err = tls.X509KeyPair(pemData, pemData)

	return
}

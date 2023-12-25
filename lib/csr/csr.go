package csr

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

type Generate struct {
	Country          string
	State            string
	Locality         string
	Organization     string
	OrganizationUnit string
	CommonName       string
}

func (c *Generate) Create() ([]byte, error) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	err = os.WriteFile(c.CommonName+"/"+"private.key", pemEncoded, 0777)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println("Error al generar la clave privada:", err)
		return nil, err
	}

	subject := pkix.Name{
		CommonName:         c.CommonName,
		OrganizationalUnit: []string{c.OrganizationUnit},
		Organization:       []string{c.Organization},
		Locality:           []string{c.Locality},
		Province:           []string{c.State},
		Country:            []string{c.Country},
	}

	csrTemplate := x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	csrDER, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		fmt.Println("Error al crear la CSR:", err)
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrDER}), nil
}

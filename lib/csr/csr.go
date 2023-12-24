package csr

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
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
	// Generar una clave privada ECDSA
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
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

	// Crear la solicitud de firma del certificado (CSR)
	csrTemplate := x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	// Codificar la CSR
	csrDER, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		fmt.Println("Error al crear la CSR:", err)
		return nil, err
	}

	// Codificar la CSR en formato PEM
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrDER}), nil
}

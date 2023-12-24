// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    certificates, err := UnmarshalCertificates(bytes)
//    bytes, err = certificates.Marshal()

package responses

import "encoding/json"

func UnmarshalCertificates(data []byte) (Certificates, error) {
	var r Certificates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Certificates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Certificates struct {
	TotalCount     int64    `json:"total_count"`
	ResultCount    int64    `json:"result_count"`
	Page           string   `json:"page"`
	Limit          int64    `json:"limit"`
	ACMEUsageLevel string   `json:"acmeUsageLevel"`
	ACMELocked     bool     `json:"acmeLocked"`
	Results        []Result `json:"results"`
}

type Result struct {
	ID                string      `json:"id"`
	Type              string      `json:"type"`
	CommonName        string      `json:"common_name"`
	AdditionalDomains string      `json:"additional_domains"`
	Created           string      `json:"created"`
	Expires           string      `json:"expires"`
	Status            string      `json:"status"`
	ValidationType    interface{} `json:"validation_type"`
	ValidationEmails  interface{} `json:"validation_emails"`
	ReplacementFor    string      `json:"replacement_for"`
	FingerprintSha1   interface{} `json:"fingerprint_sha1"`
	BrandValidation   interface{} `json:"brand_validation"`
	Validation        Validation  `json:"validation"`
}

type Validation struct {
	EmailValidation EmailValidation        `json:"email_validation"`
	OtherMethods    map[string]OtherMethod `json:"other_methods"`
}

type EmailValidation struct {
	FedeCOM []string `json:"fede.com"`
}

type OtherMethod struct {
	FileValidationURLHTTP  string   `json:"file_validation_url_http"`
	FileValidationURLHTTPS string   `json:"file_validation_url_https"`
	FileValidationContent  []string `json:"file_validation_content"`
	CnameValidationP1      string   `json:"cname_validation_p1"`
	CnameValidationP2      string   `json:"cname_validation_p2"`
}

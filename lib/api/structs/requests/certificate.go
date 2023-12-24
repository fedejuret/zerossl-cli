package requests

type CreateCertificationStructure struct {
	Domains      string `json:"certificate_domains"`
	Csr          string `json:"certificate_csr"`
	ValidityDays uint16 `json:"certificate_validity_days"`
}

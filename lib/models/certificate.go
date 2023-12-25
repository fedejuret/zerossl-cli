package models

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func UnmarshalCertificate(data []byte) (Certificate, error) {
	var r Certificate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Certificate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Certificate struct {
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
	Validation        Validation  `json:"validation"`
}

type Validation struct {
	EmailValidation map[string]interface{} `json:"email_validation"`
	OtherMethods    map[string]interface{} `json:"other_methods"`
}

func (c *Certificate) GetFileValidationURLHTTPS() (string, error) {
	fileValidationURLHTTPS, ok := c.Validation.OtherMethods[c.CommonName].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("%s is not a map[string]interface{}", c.CommonName)
	}

	parsedURL, ok := fileValidationURLHTTPS["file_validation_url_https"].(string)
	if !ok {
		return "", fmt.Errorf("file_validation_url_https is not a string")
	}

	return parsedURL, nil
}

func (c *Certificate) GetDNSValidation() (cname string, content string, err error) {
	certData, ok := c.Validation.OtherMethods[c.CommonName].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("%s is not a map[string]interface{}", c.CommonName)
	}

	cname, ok = certData["cname_validation_p1"].(string)
	if !ok {
		return "", "", fmt.Errorf("cname_validation_p1 is not a string")
	}

	content, ok = certData["cname_validation_p2"].(string)
	if !ok {
		return "", "", fmt.Errorf("cname_validation_p2 is not a string")
	}

	return cname, content, nil
}

func (c *Certificate) GetFileValidationContent() ([]string, error) {
	otherMethods, ok := c.Validation.OtherMethods[c.CommonName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%s is not a map[string]interface{}", c.CommonName)
	}

	contentInterface, ok := otherMethods["file_validation_content"]
	if !ok {
		return nil, fmt.Errorf("file_validation_content not found")
	}

	if reflect.TypeOf(contentInterface).Kind() != reflect.Slice {
		return nil, fmt.Errorf("file_validation_content is not an array")
	}

	contentSlice, ok := contentInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("file_validation_content is not a []interface{}")
	}

	var content []string
	for _, val := range contentSlice {
		if strVal, ok := val.(string); ok {
			content = append(content, strVal)
		} else {
			return nil, fmt.Errorf("file_validation_content contains non-string elements")
		}
	}

	return content, nil
}

package showMe

import (
	"encoding/json"
	"encoding/xml"
)

func jsonToString(data interface{}) (string, error) {
	j, err := json.MarshalIndent(data, "", ident)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func xmlToString(data interface{}) (string, error) {
	j, err := xml.MarshalIndent(data, "", ident)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

package common

import (
	"encoding/json"
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/model"
	"net/http"
	"net/url"
)

func ParseUrl(urlPath *string, allowEmptyScheme bool, allowEmptyHost bool, allowEmptyPath bool) (*url.URL, error) {
	address, err := url.Parse(*urlPath)
	if err != nil {
		return nil, err
	}
	if address.Scheme == "" && !allowEmptyScheme {
		return nil, fmt.Errorf("missing scheme from url")
	}
	if address.Host == "" && !allowEmptyHost {
		return nil, fmt.Errorf("missing host from url")
	}

	if address.Path == "" && !allowEmptyPath {
		return nil, fmt.Errorf("missing path from url")
	}
	return address, nil
}

func RespondAsJson(responseWriter http.ResponseWriter, response model.Documents) error {
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	_, err = responseWriter.Write(marshaledResponse)
	if err != nil {
		return err
	}
	return nil
}

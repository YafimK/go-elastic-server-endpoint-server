package common

import (
	"fmt"
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

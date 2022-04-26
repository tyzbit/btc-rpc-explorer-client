package btcapi

import (
	"fmt"
)

const VersionRoute string = "/version"

// ExtendedPublicKeyDetails returns details for an extended public key.
func (c Config) Version() (version string, err error) {
	body, err := getAPI(c.ExplorerURL + api + VersionRoute)
	if err != nil {
		return version, err
	}

	version = string(body)
	if err != nil {
		return version, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return version, nil
}

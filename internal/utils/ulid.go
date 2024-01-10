/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/subhradip-bose/openfga-go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package internalutils

import (
	"regexp"
)

// cUlidRegex contains the regex for valid ULID
const cUlidRegex = "^[0-7][0-9A-HJKMNP-TV-Z]{25}$"

// IsWellFormedUlidString returns whethr the ulidString is a properly formatted ulid string
func IsWellFormedUlidString(ulidString string) bool {
	re := regexp.MustCompile(cUlidRegex)
	return re.MatchString(ulidString)
}

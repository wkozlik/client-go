/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// HashHashType : Specifies the hash algorithm, if any.   - NONE: No hash requested.  - SHA256: A sha256 hash.
type HashHashType string

// List of HashHashType
const (
	NONE HashHashType = "NONE"
	SHA256 HashHashType = "SHA256"
)
/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Details of a discovery occurrence.
type V1beta1discoveryDetails struct {
	// Required. Analysis status for the discovered resource.
	Discovered *DiscoveryDiscovered `json:"discovered,omitempty"`
}

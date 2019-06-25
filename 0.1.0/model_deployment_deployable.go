/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// An artifact that can be deployed in some runtime.
type DeploymentDeployable struct {
	// Required. Resource URI for the artifact being deployed.
	ResourceUri []string `json:"resource_uri,omitempty"`
}

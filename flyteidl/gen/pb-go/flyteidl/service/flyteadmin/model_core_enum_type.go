/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Enables declaring enum types, with predefined string values For len(values) > 0, the first value in the ordered list is regarded as the default value. If you wish To provide no defaults, make the first value as undefined.
type CoreEnumType struct {
	// Predefined set of enum values.
	Values []string `json:"values,omitempty"`
}

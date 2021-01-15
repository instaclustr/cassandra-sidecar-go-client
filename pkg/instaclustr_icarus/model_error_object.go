/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.4
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type ErrorObject struct {
	// hostame of a node where this error has occurred
	Source string `json:"source,omitempty"`
	// message explaining the error
	Message string `json:"message,omitempty"`
}

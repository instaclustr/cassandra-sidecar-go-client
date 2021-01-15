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

// bandwidth used during uploads 
type DataRate struct {
	// quantified value of bandwidth, an integer 
	Value int32 `json:"value,omitempty"`
	// unit of 'data bandwidth' 
	Unit string `json:"unit"`
}

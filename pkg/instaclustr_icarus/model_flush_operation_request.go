/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.7
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

// flushes tables of a keyspace or all tables of a keyspace when tables are not specified 
type FlushOperationRequest struct {
	Type_ string `json:"type"`
	// keyspace to flush 
	Keyspace string `json:"keyspace"`
	// tables to flush, if not provided or empty, flush all tables in a keyspace 
	Tables []string `json:"tables,omitempty"`
}

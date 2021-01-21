/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.5
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type UpgradeSsTablesOperationRequest struct {
	Type_ string `json:"type"`
	// keyspace to upgrade SSTables of 
	Keyspace string `json:"keyspace"`
	// an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace 
	Tables []string `json:"tables,omitempty"`
	// the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads 
	Jobs int32 `json:"jobs,omitempty"`
	// include all sstables, even those already on the current version, defaults to false
	IncludeAllSStables bool `json:"includeAllSStables,omitempty"`
}

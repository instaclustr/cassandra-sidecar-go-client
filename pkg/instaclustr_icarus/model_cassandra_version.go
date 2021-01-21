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

type CassandraVersion struct {
	Major int32 `json:"major,omitempty"`
	Minor int32 `json:"minor,omitempty"`
	Patch int32 `json:"patch,omitempty"`
	DsePatch int32 `json:"dsePatch,omitempty"`
	// Labels for Cassandra version, e.g. SNAPSHOT
	PreReleaseLabels []string `json:"preReleaseLabels,omitempty"`
}

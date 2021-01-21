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

type Retry struct {
	// Time gap between retries, linear strategy will have always this gap constant, exponential strategy will make the gap bigger exponentially (power of 2) on each attempt 
	Interval int32 `json:"interval,omitempty"`
	// Strategy how retry should be driven, might be either 'LINEAR' or 'EXPONENTIAL' 
	Strategy string `json:"strategy,omitempty"`
	// Number of repetitions of an upload / download operation in case it fails before giving up completely. 
	MaxAttempts int32 `json:"maxAttempts,omitempty"`
	// Defaults to false if not specified. If false, retry mechanism on upload / download operations in case they fail will not be used. 
	Enabled bool `json:"enabled,omitempty"`
}

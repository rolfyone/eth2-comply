/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// ErrorMessage struct for ErrorMessage
type ErrorMessage struct {
	// Either specific error code in case of invalid request or http status code
	Code float32 `json:"code,omitempty"`
	// Message describing error
	Message string `json:"message,omitempty"`
	// Optional stacktraces, sent when node is in debug mode
	Stacktraces []string `json:"stacktraces,omitempty"`
}
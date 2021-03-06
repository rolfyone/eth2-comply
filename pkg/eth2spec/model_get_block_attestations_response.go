/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// GetBlockAttestationsResponse struct for GetBlockAttestationsResponse
type GetBlockAttestationsResponse struct {
	Data []GetBlockAttestationsResponseData1 `json:"data,omitempty"`
}

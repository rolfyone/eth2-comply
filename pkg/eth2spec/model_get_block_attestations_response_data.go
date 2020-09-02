/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// GetBlockAttestationsResponseData The [`AttestationData`](https://github.com/ethereum/eth2.0-specs/blob/v0.12.2/specs/phase0/beacon-chain.md#attestationdata) object from the Eth2.0 spec.
type GetBlockAttestationsResponseData struct {
	Slot string `json:"slot,omitempty"`
	Index string `json:"index,omitempty"`
	BeaconBlockRoot string `json:"beacon_block_root,omitempty"`
	Source GetStateFinalityCheckpointsResponseDataPreviousJustified `json:"source,omitempty"`
	Target GetStateFinalityCheckpointsResponseDataPreviousJustified `json:"target,omitempty"`
}

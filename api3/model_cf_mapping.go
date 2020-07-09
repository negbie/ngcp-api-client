/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CfMapping struct for CfMapping
type CfMapping struct {
	// Call Forward Unavailable, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber.
	Cfna []CfMappingCfna `json:"cfna"`
	// Call Forward Timeout, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber.
	Cft []CfMappingCft `json:"cft"`
	// Call Forward Busy, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber.
	Cfb []CfMappingCfb `json:"cfb"`
	// Call Forward Rerouting, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber.
	Cfr []CfMappingCfr `json:"cfr"`
	// Call Forward Unconditional, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber.
	Cfu []CfMappingCfu `json:"cfu"`
	// Call Forward SMS, Number of Objects, each containing the keys \"destinationset\", \"timeset\" and \"sourceset\". The values must be the name of a corresponding set which belongs to the same subscriber. Alternatively, you can pass destinationset_id, timeset_id and sourceset_id instead of names.
	Cfs []CfMappingCfr `json:"cfs"`
	// Cft ringtimeout
	CftRingtimeout float32 `json:"cft_ringtimeout"`
}
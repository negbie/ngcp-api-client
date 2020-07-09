/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// LnpNumber struct for LnpNumber
type LnpNumber struct {
	// An optional routing number replacing the ported number.
	RoutingNumber string `json:"routing_number"`
	// The optional date when the porting gets active in format YYYY-MM-DD.
	Start string `json:"start"`
	// The ported number.
	Number string `json:"number,omitempty"`
	// The LNP carrier this number is ported to.
	CarrierId float32 `json:"carrier_id,omitempty"`
	// The optional LNP number type tag, for CDR exports.
	Type string `json:"type"`
	// The optional date when the porting gets inactive again in format YYYY-MM-DD.
	End string `json:"end"`
}
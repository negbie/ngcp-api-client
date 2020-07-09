/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Number struct for Number
type Number struct {
	// The id of the subscriber the number is assigned to.
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// The reseller this number belongs to.
	ResellerId float32 `json:"reseller_id,omitempty"`
	// Area Code, e.g. 212 for NYC or 1 for Vienna (read-only)
	Ac string `json:"ac"`
	// Subscriber Number, e.g. 12345678 (read-only)
	Sn string `json:"sn"`
	// Whether the number is a primary number or not (read-only).
	IsPrimary bool `json:"is_primary"`
	// Country Code, e.g. 1 for US or 43 for Austria (read-only)
	Cc string `json:"cc"`
}

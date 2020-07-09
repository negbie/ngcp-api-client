/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SubscriberAliasNumbers struct for SubscriberAliasNumbers
type SubscriberAliasNumbers struct {
	// Subscriber Number, e.g. 12345678
	Sn string `json:"sn"`
	// Country Code, e.g. 1 for US or 43 for Austria
	Cc string `json:"cc"`
	// Area Code, e.g. 212 for NYC or 1 for Vienna
	Ac string `json:"ac"`
}

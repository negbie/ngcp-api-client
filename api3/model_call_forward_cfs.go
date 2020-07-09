/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallForwardCfs Call Forward SMS, Contains the keys \"destinations\", \"times\" and \"sources\". \"destinations\" is an Array of Objects having a \"destination\", \"priority\" and \"timeout\" field. \"times\" is an Array of Objects having the fields \"minute\", \"hour\", \"wday\", \"mday\", \"month\", \"year\". \"times\" can be empty, then the CF is applied always. \"sources\" is an Array of Objects having one field \"source\". \"sources\" can be empty.
type CallForwardCfs struct {
	// Destinations
	Destinations []CallForwardCfsDestinations `json:"destinations"`
	// Bnumbers mode
	BnumbersMode string `json:"bnumbers_mode"`
	// Sources is regex
	SourcesIsRegex bool `json:"sources_is_regex"`
	// Bnumbers is regex
	BnumbersIsRegex bool `json:"bnumbers_is_regex"`
	// Bnumbers
	Bnumbers []CallForwardCftBnumbers `json:"bnumbers"`
	// Sources mode
	SourcesMode string `json:"sources_mode"`
	// Times
	Times []CallForwardCfsTimes `json:"times"`
	// Sources
	Sources []CallForwardCftSources `json:"sources"`
}

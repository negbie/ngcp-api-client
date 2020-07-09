/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CfSourceSet struct for CfSourceSet
type CfSourceSet struct {
	// The name of the source set
	Name string `json:"name,omitempty"`
	// The subscriber id this source set belongs to
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// A flag indicating, whether the numbers in this set are regular expressions. If true, all sources will be interepreted as perl compatible regular expressions and matched against the calling party number (in E164 format) of the calls. If false, the whole numbers are plainly matched while shell patterns like 431* or 49123[1-5]67 are possible.
	IsRegex bool `json:"is_regex"`
	// Source set mode. If set to \"blacklist\" it enables forwarding for everything except numbers in the list, and \"whitelist\" only enables forwards for numbers defined in this list. This field is mandatory.
	Mode string `json:"mode,omitempty"`
	// An array of sources, each containing the key \"source\" which will be matched against the calling party number to determine whether to apply the callforward or not. \"source\" is the calling party number in E164 format to match. Regular expressions or shell patterns can be used depending on the is_regex flag. Use \"anonymous\" to match suppressed numbers.
	Sources []CfSourceSetSources `json:"sources"`
}

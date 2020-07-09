/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PbxDeviceModelLinerange struct for PbxDeviceModelLinerange
type PbxDeviceModelLinerange struct {
	// Lines/Keys in this range can be used as Busy Lamp Field. Value is accessible in the config template via phone.lineranges.lines.can_blf
	CanBlf bool `json:"can_blf"`
	// Lines/Keys in this range can be used as regular phone lines. Value is accessible in the config template via phone.lineranges.lines.can_private
	CanPrivate bool `json:"can_private"`
	// The Name of this range, e.g. Phone Keys or Attendant Console 1 Keys, accessible in the config template array via phone.lineranges.name
	Name string `json:"name"`
	// Lines/Keys in this range can be used as shared lines. Value is accessible in the config template via phone.lineranges.lines.can_shared
	CanShared bool `json:"can_shared"`
	// The position of the keys on the front image. Attributes are x, y, labelpos (how the label for the key is displayed in the web interface, relative to the given coordinates; one of top, bottom, left, right).
	Keys []PbxDeviceModelKeys `json:"keys"`
}

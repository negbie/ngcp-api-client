/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ManagerSecretary struct for ManagerSecretary
type ManagerSecretary struct {
	// Secretary numbers
	SecretaryNumbers []ManagerSecretarySecretaryNumbers `json:"secretary_numbers"`
	// The subscriber UUID.
	Uuid string `json:"uuid,omitempty"`
}

/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BillingProfilePeaktimeWeekdays struct for BillingProfilePeaktimeWeekdays
type BillingProfilePeaktimeWeekdays struct {
	// Wochentag
	Weekday float32 `json:"weekday,omitempty"`
	// Start
	Start string `json:"start"`
	// Stop
	Stop string `json:"stop"`
}
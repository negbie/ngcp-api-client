/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Reminder struct for Reminder
type Reminder struct {
	// The reminder recurrence (one of never, weekdays, always).
	Recur string `json:"recur,omitempty"`
	// Enables or disables the reminder. When Repeat is set to once, reminder is also disabled after triggering once.
	Active bool `json:"active"`
	// The subscriber this reminder belongs to.
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// The time the reminder call is triggered.
	Time string `json:"time,omitempty"`
}
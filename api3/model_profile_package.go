/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProfilePackage struct for ProfilePackage
type ProfilePackage struct {
	// The lock level to reset all customer's subscribers to after a successful top-up (usually 0).
	TopupLockLevel string `json:"topup_lock_level"`
	// Arbitrary text.
	Description string `json:"description,omitempty"`
	// The balance threshold (in cents) for the underrun lock level to come into effect.
	UnderrunLockThreshold float32 `json:"underrun_lock_threshold"`
	// An array of objects with keys \"profile_id\" and \"network_id\" to create profile mappings from when applying this profile package to a customer.
	InitialProfiles []ProfilePackageInitialProfiles `json:"initial_profiles"`
	// The temporal unit for the balance interval.
	BalanceIntervalUnit string `json:"balance_interval_unit"`
	// Options to carry over the customer's balance to the next balance interval.
	CarryOverMode string `json:"carry_over_mode"`
	// The balance interval in temporal units.
	BalanceIntervalValue float32 `json:"balance_interval_value"`
	// The initial balance (in cents) that will be set for the very first balance interval.
	InitialBalance float32 `json:"initial_balance"`
	// The reseller id this profile package belongs to.
	ResellerId float32 `json:"reseller_id"`
	// An array of objects with keys \"profile_id\" and \"network_id\" to create profile mappings from when the balance underruns the \"underrun_profile_threshold\" value.
	UnderrunProfiles []ProfilePackageInitialProfiles `json:"underrun_profiles"`
	// The lock level to set all customer's subscribers to in case the balance underruns \"underrun_lock_threshold\".
	UnderrunLockLevel string `json:"underrun_lock_level"`
	// The balance will be discarded if no top-up happened for the given number of balance interval units.
	NotopupDiscardIntervals float32 `json:"notopup_discard_intervals"`
	// The balance threshold (in cents) for underrun profiles to come into effect.
	UnderrunProfileThreshold float32 `json:"underrun_profile_threshold"`
	// The service charge amount (in cents) will be subtracted from the voucher amount.
	ServiceCharge float32 `json:"service_charge"`
	// The temporal unit for the \"timely\" interval.
	TimelyDurationUnit string `json:"timely_duration_unit"`
	// This mode determines when balance intervals start.
	BalanceIntervalStartMode string `json:"balance_interval_start_mode"`
	// The unique name of the profile package.
	Name string `json:"name,omitempty"`
	// An array of objects with keys \"profile_id\" and \"network_id\" to create profile mappings from when a customer top-ups with a voucher associated with this profile package.
	TopupProfiles []ProfilePackageTopupProfiles `json:"topup_profiles"`
	// The \"timely\" interval in temporal units.
	TimelyDurationValue float32 `json:"timely_duration_value"`
}
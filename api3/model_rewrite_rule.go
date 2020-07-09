/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RewriteRule struct for RewriteRule
type RewriteRule struct {
	// Replacement pattern.
	ReplacePattern string `json:"replace_pattern,omitempty"`
	// Arbitrary text.
	Description string `json:"description,omitempty"`
	// Match pattern, a PCRE regular expression.
	MatchPattern string `json:"match_pattern,omitempty"`
	// The rewrite rule priority.
	Priority float32 `json:"priority"`
	// Inbound (in), Outbound (out) or LNP (lnp).
	Direction string `json:"direction"`
	// caller or callee.
	Field string `json:"field"`
	// Rule enabled state.
	Enabled bool `json:"enabled"`
	// The rewrite rule set this rule belongs to.
	SetId float32 `json:"set_id,omitempty"`
}

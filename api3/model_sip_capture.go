/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SipCapture struct for SipCapture
type SipCapture struct {
	// CSEQ Method of the sip packet
	CseqMethod string `json:"cseq_method,omitempty"`
	// Source port of the sip packet
	SrcPort string `json:"src_port,omitempty"`
	// Destination port of the sip packet
	DstPort float32 `json:"dst_port,omitempty"`
	// Protocol of the sip packet
	Protocol string `json:"protocol,omitempty"`
	// From URI of the sip packet
	FromUri string `json:"from_uri,omitempty"`
	// Request URI of the sip packet
	RequestUri string `json:"request_uri,omitempty"`
	// Method of the sip packet
	Method string `json:"method,omitempty"`
	// IP transport of the sip packet (TCP/UDP)
	Transport string `json:"transport,omitempty"`
	// Timestamp of the sip packet
	Timestamp string `json:"timestamp,omitempty"`
	// Source IP of the sip packet
	SrcIp string `json:"src_ip,omitempty"`
	// Call id of the sip packet
	CallId string `json:"call_id,omitempty"`
	// Destination IP of the sip packet
	DstIp string `json:"dst_ip,omitempty"`
}

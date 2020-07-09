/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CustomerContact struct for CustomerContact
type CustomerContact struct {
	// A general purpose field for free use.
	Gpp3 string `json:"gpp3"`
	// The given name of the contact.
	Firstname string `json:"firstname"`
	// A general purpose field for free use.
	Gpp8 string `json:"gpp8"`
	// The timezone of the contact.
	Timezone float32 `json:"timezone"`
	// The VAT number of the contact.
	Vatnum string `json:"vatnum"`
	// A general purpose field for free use.
	Gpp0 string `json:"gpp0"`
	// The mobile number of the contact.
	Mobilenumber string `json:"mobilenumber"`
	// The email address of the contact.
	Email string `json:"email,omitempty"`
	// The reseller id this contact belongs to.
	ResellerId float32 `json:"reseller_id,omitempty"`
	// A general purpose field for free use.
	Gpp6 string `json:"gpp6"`
	// The IBAN (International Bank Account Number) of the contact bank details.
	Iban string `json:"iban"`
	// The BIC (Business Identifier Code) of the contact bank details.
	Bic string `json:"bic"`
	// The company name of the contact.
	Company string `json:"company"`
	// The fax number of the contact.
	Faxnumber string `json:"faxnumber"`
	// The city name of the contact.
	City string `json:"city"`
	// A general purpose field for free use.
	Gpp2 string `json:"gpp2"`
	// The surname of the contact.
	Lastname string `json:"lastname"`
	// The two-letter ISO 3166-1 country code of the contact (e.g. US or DE).
	Country string `json:"country"`
	// A general purpose field for free use.
	Gpp5 string `json:"gpp5"`
	// A general purpose field for free use.
	Gpp1 string `json:"gpp1"`
	// The bank name of the contact bank details.
	Bankname string `json:"bankname"`
	// The company registration number of the contact.
	Comregnum string `json:"comregnum"`
	// A general purpose field for free use.
	Gpp7 string `json:"gpp7"`
	// The street name of the contact.
	Street string `json:"street"`
	// A general purpose field for free use.
	Gpp9 string `json:"gpp9"`
	// The phone number of the contact.
	Phonenumber string `json:"phonenumber"`
	// A general purpose field for free use.
	Gpp4 string `json:"gpp4"`
	// The postal code of the contact.
	Postcode string `json:"postcode"`
}
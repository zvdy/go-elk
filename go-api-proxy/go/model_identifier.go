// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Identifier - Identifier within the parent organisation for the product. Must be unique in the organisation. GenericIdentification|GenericIdentification (https://www.iso20022.org/standardsrepository/type/GenericIdentification)|Standard|ISO20022 BM ()
type Identifier struct {

	IdentifierValue Value `json:"IdentifierValue,omitempty"`

	IdentifierIssuingAuthority Involvedparty `json:"IdentifierIssuingAuthority,omitempty"`

	IdentifierStartDate Datetime `json:"IdentifierStartDate,omitempty"`

	IdentifierEndDate Datetime `json:"IdentifierEndDate,omitempty"`
}

// AssertIdentifierRequired checks if the required fields are not zero-ed
func AssertIdentifierRequired(obj Identifier) error {
	if err := AssertValueRequired(obj.IdentifierValue); err != nil {
		return err
	}
	if err := AssertInvolvedpartyRequired(obj.IdentifierIssuingAuthority); err != nil {
		return err
	}
	if err := AssertDatetimeRequired(obj.IdentifierStartDate); err != nil {
		return err
	}
	if err := AssertDatetimeRequired(obj.IdentifierEndDate); err != nil {
		return err
	}
	return nil
}

// AssertIdentifierConstraints checks if the values respects the defined constraints
func AssertIdentifierConstraints(obj Identifier) error {
	if err := AssertValueConstraints(obj.IdentifierValue); err != nil {
		return err
	}
	if err := AssertInvolvedpartyConstraints(obj.IdentifierIssuingAuthority); err != nil {
		return err
	}
	if err := AssertDatetimeConstraints(obj.IdentifierStartDate); err != nil {
		return err
	}
	if err := AssertDatetimeConstraints(obj.IdentifierEndDate); err != nil {
		return err
	}
	return nil
}

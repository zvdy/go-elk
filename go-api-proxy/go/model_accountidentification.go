// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Accountidentification - The identifier of account like Account Number.
type Accountidentification struct {

	AccountIdentificationType Accountidentificationtypevalues `json:"AccountIdentificationType,omitempty"`

	AccountIdentification Identifier `json:"AccountIdentification,omitempty"`
}

// AssertAccountidentificationRequired checks if the required fields are not zero-ed
func AssertAccountidentificationRequired(obj Accountidentification) error {
	if err := AssertIdentifierRequired(obj.AccountIdentification); err != nil {
		return err
	}
	return nil
}

// AssertAccountidentificationConstraints checks if the values respects the defined constraints
func AssertAccountidentificationConstraints(obj Accountidentification) error {
	if err := AssertIdentifierConstraints(obj.AccountIdentification); err != nil {
		return err
	}
	return nil
}

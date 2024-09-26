// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Branch - The name of the branch where the customer will collect the device
type Branch struct {

	BranchIdentification Identifier `json:"BranchIdentification,omitempty"`

	BranchLegalEntityIdentification Identifier `json:"BranchLegalEntityIdentification,omitempty"`

	BranchAddress Address `json:"BranchAddress,omitempty"`

	BranchName Name `json:"BranchName,omitempty"`
}

// AssertBranchRequired checks if the required fields are not zero-ed
func AssertBranchRequired(obj Branch) error {
	if err := AssertIdentifierRequired(obj.BranchIdentification); err != nil {
		return err
	}
	if err := AssertIdentifierRequired(obj.BranchLegalEntityIdentification); err != nil {
		return err
	}
	if err := AssertAddressRequired(obj.BranchAddress); err != nil {
		return err
	}
	if err := AssertNameRequired(obj.BranchName); err != nil {
		return err
	}
	return nil
}

// AssertBranchConstraints checks if the values respects the defined constraints
func AssertBranchConstraints(obj Branch) error {
	if err := AssertIdentifierConstraints(obj.BranchIdentification); err != nil {
		return err
	}
	if err := AssertIdentifierConstraints(obj.BranchLegalEntityIdentification); err != nil {
		return err
	}
	if err := AssertAddressConstraints(obj.BranchAddress); err != nil {
		return err
	}
	if err := AssertNameConstraints(obj.BranchName); err != nil {
		return err
	}
	return nil
}

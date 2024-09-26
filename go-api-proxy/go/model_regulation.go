// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Regulation - Reference to the regulation which is defined in Leasing Arrangement ||  |
type Regulation struct {

	// 
	RegulationDefinition string `json:"RegulationDefinition,omitempty"`
}

// AssertRegulationRequired checks if the required fields are not zero-ed
func AssertRegulationRequired(obj Regulation) error {
	return nil
}

// AssertRegulationConstraints checks if the values respects the defined constraints
func AssertRegulationConstraints(obj Regulation) error {
	return nil
}

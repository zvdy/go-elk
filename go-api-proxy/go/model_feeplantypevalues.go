// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi


import (
	"fmt"
)


// Feeplantypevalues : >  * `StandardFeePlan` -   * `FinalFeePlan` -   
type Feeplantypevalues string

// List of Feeplantypevalues
const (
	STANDARD_FEE_PLAN Feeplantypevalues = "StandardFeePlan"
	FINAL_FEE_PLAN Feeplantypevalues = "FinalFeePlan"
)

// AllowedFeeplantypevaluesEnumValues is all the allowed values of Feeplantypevalues enum
var AllowedFeeplantypevaluesEnumValues = []Feeplantypevalues{
	"StandardFeePlan",
	"FinalFeePlan",
}

// validFeeplantypevaluesEnumValue provides a map of Feeplantypevaluess for fast verification of use input
var validFeeplantypevaluesEnumValues = map[Feeplantypevalues]struct{}{
	"StandardFeePlan": {},
	"FinalFeePlan": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Feeplantypevalues) IsValid() bool {
	_, ok := validFeeplantypevaluesEnumValues[v]
	return ok
}

// NewFeeplantypevaluesFromValue returns a pointer to a valid Feeplantypevalues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeplantypevaluesFromValue(v string) (Feeplantypevalues, error) {
	ev := Feeplantypevalues(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for Feeplantypevalues: valid values are %v", v, AllowedFeeplantypevaluesEnumValues)
}



// AssertFeeplantypevaluesRequired checks if the required fields are not zero-ed
func AssertFeeplantypevaluesRequired(obj Feeplantypevalues) error {
	return nil
}

// AssertFeeplantypevaluesConstraints checks if the values respects the defined constraints
func AssertFeeplantypevaluesConstraints(obj Feeplantypevalues) error {
	return nil
}

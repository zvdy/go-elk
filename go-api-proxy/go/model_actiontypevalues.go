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


// Actiontypevalues : >  * `Initiate` -   * `Execute` -   * `Create` -   * `Transfer` -   * `Pay` -   * `Deliver` -   * `Apply` -   * `Calculate` -   
type Actiontypevalues string

// List of Actiontypevalues
const (
	INITIATE Actiontypevalues = "Initiate"
	EXECUTE Actiontypevalues = "Execute"
	CREATE Actiontypevalues = "Create"
	TRANSFER Actiontypevalues = "Transfer"
	PAY Actiontypevalues = "Pay"
	DELIVER Actiontypevalues = "Deliver"
	APPLY Actiontypevalues = "Apply"
	CALCULATE Actiontypevalues = "Calculate"
)

// AllowedActiontypevaluesEnumValues is all the allowed values of Actiontypevalues enum
var AllowedActiontypevaluesEnumValues = []Actiontypevalues{
	"Initiate",
	"Execute",
	"Create",
	"Transfer",
	"Pay",
	"Deliver",
	"Apply",
	"Calculate",
}

// validActiontypevaluesEnumValue provides a map of Actiontypevaluess for fast verification of use input
var validActiontypevaluesEnumValues = map[Actiontypevalues]struct{}{
	"Initiate": {},
	"Execute": {},
	"Create": {},
	"Transfer": {},
	"Pay": {},
	"Deliver": {},
	"Apply": {},
	"Calculate": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Actiontypevalues) IsValid() bool {
	_, ok := validActiontypevaluesEnumValues[v]
	return ok
}

// NewActiontypevaluesFromValue returns a pointer to a valid Actiontypevalues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActiontypevaluesFromValue(v string) (Actiontypevalues, error) {
	ev := Actiontypevalues(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for Actiontypevalues: valid values are %v", v, AllowedActiontypevaluesEnumValues)
}



// AssertActiontypevaluesRequired checks if the required fields are not zero-ed
func AssertActiontypevaluesRequired(obj Actiontypevalues) error {
	return nil
}

// AssertActiontypevaluesConstraints checks if the values respects the defined constraints
func AssertActiontypevaluesConstraints(obj Actiontypevalues) error {
	return nil
}

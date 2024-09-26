// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Businessservice - The Leasing Arrangement specific Business Service
type Businessservice struct {

	BusinessServiceType Businessservicetypevalues `json:"BusinessServiceType,omitempty"`
}

// AssertBusinessserviceRequired checks if the required fields are not zero-ed
func AssertBusinessserviceRequired(obj Businessservice) error {
	if err := AssertBusinessservicetypevaluesRequired(obj.BusinessServiceType); err != nil {
		return err
	}
	return nil
}

// AssertBusinessserviceConstraints checks if the values respects the defined constraints
func AssertBusinessserviceConstraints(obj Businessservice) error {
	if err := AssertBusinessservicetypevaluesConstraints(obj.BusinessServiceType); err != nil {
		return err
	}
	return nil
}

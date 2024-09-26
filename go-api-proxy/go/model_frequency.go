// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Frequency - Specifies the periodicity of the calculation of the interest.
type Frequency struct {

	FrequencyCode Frequencytypevalues `json:"FrequencyCode,omitempty"`

	FrequencyName Name `json:"FrequencyName,omitempty"`

	FrequencyDefinition Text `json:"FrequencyDefinition,omitempty"`
}

// AssertFrequencyRequired checks if the required fields are not zero-ed
func AssertFrequencyRequired(obj Frequency) error {
	if err := AssertNameRequired(obj.FrequencyName); err != nil {
		return err
	}
	if err := AssertTextRequired(obj.FrequencyDefinition); err != nil {
		return err
	}
	return nil
}

// AssertFrequencyConstraints checks if the values respects the defined constraints
func AssertFrequencyConstraints(obj Frequency) error {
	if err := AssertNameConstraints(obj.FrequencyName); err != nil {
		return err
	}
	if err := AssertTextConstraints(obj.FrequencyDefinition); err != nil {
		return err
	}
	return nil
}

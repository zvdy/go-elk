// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Arrangement - 
type Arrangement struct {

	ArrangementAction Action `json:"ArrangementAction,omitempty"`

	ArrangementStartDate Datetime `json:"ArrangementStartDate,omitempty"`

	ArrangementEndDate Datetime `json:"ArrangementEndDate,omitempty"`

	ArrangementStatus Arrangementlifecyclestatus `json:"ArrangementStatus,omitempty"`

	ArrangementSubjectMatter Subject `json:"ArrangementSubjectMatter,omitempty"`

	Arrangementtype Arrangementtypevalues `json:"Arrangementtype,omitempty"`
}

// AssertArrangementRequired checks if the required fields are not zero-ed
func AssertArrangementRequired(obj Arrangement) error {
	if err := AssertActionRequired(obj.ArrangementAction); err != nil {
		return err
	}
	if err := AssertDatetimeRequired(obj.ArrangementStartDate); err != nil {
		return err
	}
	if err := AssertDatetimeRequired(obj.ArrangementEndDate); err != nil {
		return err
	}
	if err := AssertArrangementlifecyclestatusRequired(obj.ArrangementStatus); err != nil {
		return err
	}
	if err := AssertSubjectRequired(obj.ArrangementSubjectMatter); err != nil {
		return err
	}
	return nil
}

// AssertArrangementConstraints checks if the values respects the defined constraints
func AssertArrangementConstraints(obj Arrangement) error {
	if err := AssertActionConstraints(obj.ArrangementAction); err != nil {
		return err
	}
	if err := AssertDatetimeConstraints(obj.ArrangementStartDate); err != nil {
		return err
	}
	if err := AssertDatetimeConstraints(obj.ArrangementEndDate); err != nil {
		return err
	}
	if err := AssertArrangementlifecyclestatusConstraints(obj.ArrangementStatus); err != nil {
		return err
	}
	if err := AssertSubjectConstraints(obj.ArrangementSubjectMatter); err != nil {
		return err
	}
	return nil
}

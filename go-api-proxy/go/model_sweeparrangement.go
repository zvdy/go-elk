// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Sweeparrangement - The set-up of the sweep facility
type Sweeparrangement struct {

	SweepType Text `json:"SweepType,omitempty"`

	SweepModalitites Servicemodality `json:"SweepModalitites,omitempty"`

	SweepCounterAccountReference Account `json:"SweepCounterAccountReference,omitempty"`

	SweepSchedule Schedule `json:"SweepSchedule,omitempty"`

	SweepMandate Mandate `json:"SweepMandate,omitempty"`

	SweepDirection Text `json:"SweepDirection,omitempty"`
}

// AssertSweeparrangementRequired checks if the required fields are not zero-ed
func AssertSweeparrangementRequired(obj Sweeparrangement) error {
	if err := AssertTextRequired(obj.SweepType); err != nil {
		return err
	}
	if err := AssertServicemodalityRequired(obj.SweepModalitites); err != nil {
		return err
	}
	if err := AssertAccountRequired(obj.SweepCounterAccountReference); err != nil {
		return err
	}
	if err := AssertScheduleRequired(obj.SweepSchedule); err != nil {
		return err
	}
	if err := AssertMandateRequired(obj.SweepMandate); err != nil {
		return err
	}
	if err := AssertTextRequired(obj.SweepDirection); err != nil {
		return err
	}
	return nil
}

// AssertSweeparrangementConstraints checks if the values respects the defined constraints
func AssertSweeparrangementConstraints(obj Sweeparrangement) error {
	if err := AssertTextConstraints(obj.SweepType); err != nil {
		return err
	}
	if err := AssertServicemodalityConstraints(obj.SweepModalitites); err != nil {
		return err
	}
	if err := AssertAccountConstraints(obj.SweepCounterAccountReference); err != nil {
		return err
	}
	if err := AssertScheduleConstraints(obj.SweepSchedule); err != nil {
		return err
	}
	if err := AssertMandateConstraints(obj.SweepMandate); err != nil {
		return err
	}
	if err := AssertTextConstraints(obj.SweepDirection); err != nil {
		return err
	}
	return nil
}

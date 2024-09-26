// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// FinancialAccounting - The product features/services available with a financical facility
type FinancialAccounting struct {

	FinancialAccountingPreconditions Condition `json:"FinancialAccountingPreconditions,omitempty"`

	FinancialAccountingFeatureSchedule Schedule `json:"FinancialAccountingFeatureSchedule,omitempty"`

	FinancialAccounting Businessservice `json:"FinancialAccounting,omitempty"`

	FinancialAccountingPostconditions Condition `json:"FinancialAccountingPostconditions,omitempty"`

	FinancialAccountingServiceType Text `json:"FinancialAccountingServiceType,omitempty"`

	FinancialAccountingServiceDescription Text `json:"FinancialAccountingServiceDescription,omitempty"`

	FinancialAccountingServiceInputsandOuputs Text `json:"FinancialAccountingServiceInputsandOuputs,omitempty"`

	FinancialAccountingServiceWorkProduct Workproduct `json:"FinancialAccountingServiceWorkProduct,omitempty"`

	FinancialAccountingServiceName Name `json:"FinancialAccountingServiceName,omitempty"`
}

// AssertFinancialAccountingRequired checks if the required fields are not zero-ed
func AssertFinancialAccountingRequired(obj FinancialAccounting) error {
	if err := AssertConditionRequired(obj.FinancialAccountingPreconditions); err != nil {
		return err
	}
	if err := AssertScheduleRequired(obj.FinancialAccountingFeatureSchedule); err != nil {
		return err
	}
	if err := AssertBusinessserviceRequired(obj.FinancialAccounting); err != nil {
		return err
	}
	if err := AssertConditionRequired(obj.FinancialAccountingPostconditions); err != nil {
		return err
	}
	if err := AssertTextRequired(obj.FinancialAccountingServiceType); err != nil {
		return err
	}
	if err := AssertTextRequired(obj.FinancialAccountingServiceDescription); err != nil {
		return err
	}
	if err := AssertTextRequired(obj.FinancialAccountingServiceInputsandOuputs); err != nil {
		return err
	}
	if err := AssertWorkproductRequired(obj.FinancialAccountingServiceWorkProduct); err != nil {
		return err
	}
	if err := AssertNameRequired(obj.FinancialAccountingServiceName); err != nil {
		return err
	}
	return nil
}

// AssertFinancialAccountingConstraints checks if the values respects the defined constraints
func AssertFinancialAccountingConstraints(obj FinancialAccounting) error {
	if err := AssertConditionConstraints(obj.FinancialAccountingPreconditions); err != nil {
		return err
	}
	if err := AssertScheduleConstraints(obj.FinancialAccountingFeatureSchedule); err != nil {
		return err
	}
	if err := AssertBusinessserviceConstraints(obj.FinancialAccounting); err != nil {
		return err
	}
	if err := AssertConditionConstraints(obj.FinancialAccountingPostconditions); err != nil {
		return err
	}
	if err := AssertTextConstraints(obj.FinancialAccountingServiceType); err != nil {
		return err
	}
	if err := AssertTextConstraints(obj.FinancialAccountingServiceDescription); err != nil {
		return err
	}
	if err := AssertTextConstraints(obj.FinancialAccountingServiceInputsandOuputs); err != nil {
		return err
	}
	if err := AssertWorkproductConstraints(obj.FinancialAccountingServiceWorkProduct); err != nil {
		return err
	}
	if err := AssertNameConstraints(obj.FinancialAccountingServiceName); err != nil {
		return err
	}
	return nil
}

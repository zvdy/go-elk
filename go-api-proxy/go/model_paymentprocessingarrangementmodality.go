// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Paymentprocessingarrangementmodality - Definition of the processing option and impact (e.g. frequency and cumulative amount constraints)
type Paymentprocessingarrangementmodality struct {

	PaymentProcessingPriority Priorityvaluetypes `json:"PaymentProcessingPriority,omitempty"`

	PaymentProcessingArrangement Arrangement `json:"PaymentProcessingArrangement,omitempty"`

	PaymentProcessingClearingChannel Clearingchanneltypevalues `json:"PaymentProcessingClearingChannel,omitempty"`

	PaymentLocalInstument Paymentlocalinstumenttypevalues `json:"PaymentLocalInstument,omitempty"`

	PaymentCategoryPurpose Paymentcategorypurposetypevalues `json:"PaymentCategoryPurpose,omitempty"`

	PaymentSequence Paymentsequencetypevalues `json:"PaymentSequence,omitempty"`
}

// AssertPaymentprocessingarrangementmodalityRequired checks if the required fields are not zero-ed
func AssertPaymentprocessingarrangementmodalityRequired(obj Paymentprocessingarrangementmodality) error {
	if err := AssertArrangementRequired(obj.PaymentProcessingArrangement); err != nil {
		return err
	}
	return nil
}

// AssertPaymentprocessingarrangementmodalityConstraints checks if the values respects the defined constraints
func AssertPaymentprocessingarrangementmodalityConstraints(obj Paymentprocessingarrangementmodality) error {
	if err := AssertArrangementConstraints(obj.PaymentProcessingArrangement); err != nil {
		return err
	}
	return nil
}

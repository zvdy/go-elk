// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Deposittransaction - Details of the deposit transaction
type Deposittransaction struct {

	DepositAmount Amount `json:"DepositAmount,omitempty"`
}

// AssertDeposittransactionRequired checks if the required fields are not zero-ed
func AssertDeposittransactionRequired(obj Deposittransaction) error {
	if err := AssertAmountRequired(obj.DepositAmount); err != nil {
		return err
	}
	return nil
}

// AssertDeposittransactionConstraints checks if the values respects the defined constraints
func AssertDeposittransactionConstraints(obj Deposittransaction) error {
	if err := AssertAmountConstraints(obj.DepositAmount); err != nil {
		return err
	}
	return nil
}

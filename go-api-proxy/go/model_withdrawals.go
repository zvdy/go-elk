// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Withdrawals - The product features/services available with a financical facility
type Withdrawals struct {

	WithdrawalTransactionDescription Text `json:"WithdrawalTransactionDescription,omitempty"`

	// Identifies the target for the withdrawal
	WithdrawalTransactionSourceReference map[string]interface{} `json:"WithdrawalTransactionSourceReference,omitempty"`

	WithdrawalType Withdrawaltypevalues `json:"WithdrawalType,omitempty"`

	WithdrawalTransaction Paymenttransaction `json:"WithdrawalTransaction,omitempty"`

	WithdrawalTransactionAmount Amount `json:"WithdrawalTransactionAmount,omitempty"`

	WithdrawalTransactionDate Date `json:"WithdrawalTransactionDate,omitempty"`

	WithdrawalTransactionWithdrawalType Withdrawaltypevalues `json:"WithdrawalTransactionWithdrawalType,omitempty"`
}

// AssertWithdrawalsRequired checks if the required fields are not zero-ed
func AssertWithdrawalsRequired(obj Withdrawals) error {
	if err := AssertTextRequired(obj.WithdrawalTransactionDescription); err != nil {
		return err
	}
	if err := AssertPaymenttransactionRequired(obj.WithdrawalTransaction); err != nil {
		return err
	}
	if err := AssertAmountRequired(obj.WithdrawalTransactionAmount); err != nil {
		return err
	}
	if err := AssertDateRequired(obj.WithdrawalTransactionDate); err != nil {
		return err
	}
	return nil
}

// AssertWithdrawalsConstraints checks if the values respects the defined constraints
func AssertWithdrawalsConstraints(obj Withdrawals) error {
	if err := AssertTextConstraints(obj.WithdrawalTransactionDescription); err != nil {
		return err
	}
	if err := AssertPaymenttransactionConstraints(obj.WithdrawalTransaction); err != nil {
		return err
	}
	if err := AssertAmountConstraints(obj.WithdrawalTransactionAmount); err != nil {
		return err
	}
	if err := AssertDateConstraints(obj.WithdrawalTransactionDate); err != nil {
		return err
	}
	return nil
}

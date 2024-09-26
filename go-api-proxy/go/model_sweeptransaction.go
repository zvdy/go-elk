// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Sweeptransaction - The transaction record for an applied sweep
type Sweeptransaction struct {

	// An transaction which is the result of fulfilling a sweep arrangement
	SweepTransaction string `json:"SweepTransaction,omitempty"`
}

// AssertSweeptransactionRequired checks if the required fields are not zero-ed
func AssertSweeptransactionRequired(obj Sweeptransaction) error {
	return nil
}

// AssertSweeptransactionConstraints checks if the values respects the defined constraints
func AssertSweeptransactionConstraints(obj Sweeptransaction) error {
	return nil
}

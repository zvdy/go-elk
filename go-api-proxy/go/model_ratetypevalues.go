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


// Ratetypevalues : >  * `Fixed` - Rate is fixed.  * `Forfeit` - No specific repurchase rate applies to the transaction Repo, only a forfeit.  * `Variable` - Rate is variable.  * `Open` - Rate has not been established.  * `Unknown` - Rate is unknown by the sender or has not been established.  * `Nilpayment` - Rate will not be paid.  * `Additionaltax` - Rate used for additional tax that cannot be categorised.  * `Charges` - Rate used to calculate the amount of the charges/fees that cannot be categorised.  * `Cashinlieuofsecurities` - Rate used to calculate the cash disbursement in lieu of a fractional quantity of, for example, equity.  * `Gross` - Cash dividend per equity before deductions or allowances have been made.  * `Cashincentive` - Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.  * `Net` - Relates to the netting of settlement instructions.  * `Sollication` - Cash rate made available in an offer in order to encourage participation in the offer.  * `Stampduty` - Financial instrument has not been stamped and/or duly signed.  * `Stockexchangetax` - Rate of stock exchange tax.  * `Withholdingtax` - Relates to a tax refund from the authorities on the associated corporate action event.  * `Transfertax` - Transaction has been generated due to transformation following a corporate action.  * `Transactiontax` - Rate used to calculate the amount of transaction tax.  * `Taxdeferred` - Rate relating to the underlying security for which tax is deferred.  * `Taxfeeamount` - Rate relating to the underlying security which is not taxable.  * `Withholdingofforeigntax` - Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.  * `Withholdingoflocaltax` - Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.  * `Imputed` - Imputed tax.  * `Precompte` - Rate is a precompte.  * `Onetiertax` - Rate is a one tier tax.  * `Localtax` - Interest liable for interest down payment tax (proportion of gross interest per unit/interim profits that is not covered by the credit in the interest pool).  * `Scheduled` - Rate of the scheduled payment.  * `Unscheduled` - Rate of the unscheduled payment.  * `Anyandall` - Any and all rate is sought.  
type Ratetypevalues string

// List of Ratetypevalues
const (
	FIXED Ratetypevalues = "Fixed"
	FORFEIT Ratetypevalues = "Forfeit"
	VARIABLE Ratetypevalues = "Variable"
	OPEN Ratetypevalues = "Open"
	UNKNOWN Ratetypevalues = "Unknown"
	NILPAYMENT Ratetypevalues = "Nilpayment"
	ADDITIONALTAX Ratetypevalues = "Additionaltax"
	CHARGES Ratetypevalues = "Charges"
	CASHINLIEUOFSECURITIES Ratetypevalues = "Cashinlieuofsecurities"
	GROSS Ratetypevalues = "Gross"
	CASHINCENTIVE Ratetypevalues = "Cashincentive"
	NET Ratetypevalues = "Net"
	SOLLICATION Ratetypevalues = "Sollication"
	STAMPDUTY Ratetypevalues = "Stampduty"
	STOCKEXCHANGETAX Ratetypevalues = "Stockexchangetax"
	WITHHOLDINGTAX Ratetypevalues = "Withholdingtax"
	TRANSFERTAX Ratetypevalues = "Transfertax"
	TRANSACTIONTAX Ratetypevalues = "Transactiontax"
	TAXDEFERRED Ratetypevalues = "Taxdeferred"
	TAXFEEAMOUNT Ratetypevalues = "Taxfeeamount"
	WITHHOLDINGOFFOREIGNTAX Ratetypevalues = "Withholdingofforeigntax"
	WITHHOLDINGOFLOCALTAX Ratetypevalues = "Withholdingoflocaltax"
	IMPUTED Ratetypevalues = "Imputed"
	PRECOMPTE Ratetypevalues = "Precompte"
	ONETIERTAX Ratetypevalues = "Onetiertax"
	LOCALTAX Ratetypevalues = "Localtax"
	SCHEDULED Ratetypevalues = "Scheduled"
	UNSCHEDULED Ratetypevalues = "Unscheduled"
	ANYANDALL Ratetypevalues = "Anyandall"
)

// AllowedRatetypevaluesEnumValues is all the allowed values of Ratetypevalues enum
var AllowedRatetypevaluesEnumValues = []Ratetypevalues{
	"Fixed",
	"Forfeit",
	"Variable",
	"Open",
	"Unknown",
	"Nilpayment",
	"Additionaltax",
	"Charges",
	"Cashinlieuofsecurities",
	"Gross",
	"Cashincentive",
	"Net",
	"Sollication",
	"Stampduty",
	"Stockexchangetax",
	"Withholdingtax",
	"Transfertax",
	"Transactiontax",
	"Taxdeferred",
	"Taxfeeamount",
	"Withholdingofforeigntax",
	"Withholdingoflocaltax",
	"Imputed",
	"Precompte",
	"Onetiertax",
	"Localtax",
	"Scheduled",
	"Unscheduled",
	"Anyandall",
}

// validRatetypevaluesEnumValue provides a map of Ratetypevaluess for fast verification of use input
var validRatetypevaluesEnumValues = map[Ratetypevalues]struct{}{
	"Fixed": {},
	"Forfeit": {},
	"Variable": {},
	"Open": {},
	"Unknown": {},
	"Nilpayment": {},
	"Additionaltax": {},
	"Charges": {},
	"Cashinlieuofsecurities": {},
	"Gross": {},
	"Cashincentive": {},
	"Net": {},
	"Sollication": {},
	"Stampduty": {},
	"Stockexchangetax": {},
	"Withholdingtax": {},
	"Transfertax": {},
	"Transactiontax": {},
	"Taxdeferred": {},
	"Taxfeeamount": {},
	"Withholdingofforeigntax": {},
	"Withholdingoflocaltax": {},
	"Imputed": {},
	"Precompte": {},
	"Onetiertax": {},
	"Localtax": {},
	"Scheduled": {},
	"Unscheduled": {},
	"Anyandall": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Ratetypevalues) IsValid() bool {
	_, ok := validRatetypevaluesEnumValues[v]
	return ok
}

// NewRatetypevaluesFromValue returns a pointer to a valid Ratetypevalues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatetypevaluesFromValue(v string) (Ratetypevalues, error) {
	ev := Ratetypevalues(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for Ratetypevalues: valid values are %v", v, AllowedRatetypevaluesEnumValues)
}



// AssertRatetypevaluesRequired checks if the required fields are not zero-ed
func AssertRatetypevaluesRequired(obj Ratetypevalues) error {
	return nil
}

// AssertRatetypevaluesConstraints checks if the values respects the defined constraints
func AssertRatetypevaluesConstraints(obj Ratetypevalues) error {
	return nil
}

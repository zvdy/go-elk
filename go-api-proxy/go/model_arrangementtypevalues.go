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


// Arrangementtypevalues : >  * `SweepArrangement` - An arrangement into which is promised for the set-up and execution of structured sweeps to and from the facility to a linked account.  * `WithdrawalArrangement` - An arrangement into which is promised to deliver Withdraw Service  to take fund out of an account.   * `DepositArrangement` - An arrangement into which is promised to deliver Deposit Service  to place funds into an account for safekeeping.   * `InterestArrangement` - An arrangement into which is promised to apply a pricing element expressed in a percentage  to a principal amount.  * `FeeArrangement` - An arrangement into which is promised to apply a charge for services rendered.   * `LienArrangement` - An arrangement into which is promised for placing a lien or block on the account for a specified purpose, amount, and period.   * `InformationArrangement` - An arrangement concerning the information the bank will provide about the agreement/account  * `PaymentArrangement` - An arrangement into which is promised to deliver a specific payment service including regular payments, standing orders, direct debits and bill pay instructions.   * `EntitlementArrangement` - Right or obligation linked to the involvement of a party in an agreement  * `PeriodArrangement` -   * `CollateralArrangement` - An arrangement into which is promised for handling the oversight of the allocation, valuation and administration of collateral associated with the facility.   * `LimitArrangement` - An arrangement to limit something e.g., the number of transaction per week, the amount of money allowed to withdraw per week/day and etc.   * `AccessArrangement` - An arrangement that allows or restricts the access to a service or product, channel or device through a specific access point  * `CardPaymentArrangement` - An arrangement into which is promised to deliver Card payment Service to execute payment transactions through a payment card, including online payments.   * `StandingOrderArrangement` - An arrangement into which is promised to deliver Standing Order Service to make regular transfers on given dates to a named beneﬁciary.   * `OverdraftArrangement` - An arrangement into which is promised to deliver Overdraft Service  that allows fund transfer even if there is no provision on the account.   * `RepaymentArrangement` - An arrangement into which is promised for paying back money previously borrowed from a lender.   * `StatementArrangement` - An arrangement into which is promised to deliver a reporting service on payment transactions registered on an account.   * `CreditTransferArrangement` - An arrangement into which is promised to allow a payer to transfer fund to a beneficiary.   * `ProductAndServiceArrangement` -   * `FactoringArrangement` -   * `RolloverArrangement` - Arrangement stipulating the modalities for an automatic renewal of the agreement  i.e. the replacement of the agreement with a new one without renewed negociation and witout explicit agreement of the customer   * `RestructuringArrangement` -   * `InsuranceArrangement` -   * `CollectionArrangement` - An arrangement into which is promised to collect past-due debts from borrowers  * `UnderwritingArrangement` - A contractual agreement between parties that commits the underwriter to assuming risk.  * `TerminationArrangement` -   * `MaturityArrangement` -   
type Arrangementtypevalues string

// List of Arrangementtypevalues
const (
	SWEEP_ARRANGEMENT Arrangementtypevalues = "SweepArrangement"
	WITHDRAWAL_ARRANGEMENT Arrangementtypevalues = "WithdrawalArrangement"
	DEPOSIT_ARRANGEMENT Arrangementtypevalues = "DepositArrangement"
	INTEREST_ARRANGEMENT Arrangementtypevalues = "InterestArrangement"
	FEE_ARRANGEMENT Arrangementtypevalues = "FeeArrangement"
	LIEN_ARRANGEMENT Arrangementtypevalues = "LienArrangement"
	INFORMATION_ARRANGEMENT Arrangementtypevalues = "InformationArrangement"
	PAYMENT_ARRANGEMENT Arrangementtypevalues = "PaymentArrangement"
	ENTITLEMENT_ARRANGEMENT Arrangementtypevalues = "EntitlementArrangement"
	PERIOD_ARRANGEMENT Arrangementtypevalues = "PeriodArrangement"
	COLLATERAL_ARRANGEMENT Arrangementtypevalues = "CollateralArrangement"
	LIMIT_ARRANGEMENT Arrangementtypevalues = "LimitArrangement"
	ACCESS_ARRANGEMENT Arrangementtypevalues = "AccessArrangement"
	CARD_PAYMENT_ARRANGEMENT Arrangementtypevalues = "CardPaymentArrangement"
	STANDING_ORDER_ARRANGEMENT Arrangementtypevalues = "StandingOrderArrangement"
	OVERDRAFT_ARRANGEMENT Arrangementtypevalues = "OverdraftArrangement"
	REPAYMENT_ARRANGEMENT Arrangementtypevalues = "RepaymentArrangement"
	STATEMENT_ARRANGEMENT Arrangementtypevalues = "StatementArrangement"
	CREDIT_TRANSFER_ARRANGEMENT Arrangementtypevalues = "CreditTransferArrangement"
	PRODUCT_AND_SERVICE_ARRANGEMENT Arrangementtypevalues = "ProductAndServiceArrangement"
	FACTORING_ARRANGEMENT Arrangementtypevalues = "FactoringArrangement"
	ROLLOVER_ARRANGEMENT Arrangementtypevalues = "RolloverArrangement"
	RESTRUCTURING_ARRANGEMENT Arrangementtypevalues = "RestructuringArrangement"
	INSURANCE_ARRANGEMENT Arrangementtypevalues = "InsuranceArrangement"
	COLLECTION_ARRANGEMENT Arrangementtypevalues = "CollectionArrangement"
	UNDERWRITING_ARRANGEMENT Arrangementtypevalues = "UnderwritingArrangement"
	TERMINATION_ARRANGEMENT Arrangementtypevalues = "TerminationArrangement"
	MATURITY_ARRANGEMENT Arrangementtypevalues = "MaturityArrangement"
)

// AllowedArrangementtypevaluesEnumValues is all the allowed values of Arrangementtypevalues enum
var AllowedArrangementtypevaluesEnumValues = []Arrangementtypevalues{
	"SweepArrangement",
	"WithdrawalArrangement",
	"DepositArrangement",
	"InterestArrangement",
	"FeeArrangement",
	"LienArrangement",
	"InformationArrangement",
	"PaymentArrangement",
	"EntitlementArrangement",
	"PeriodArrangement",
	"CollateralArrangement",
	"LimitArrangement",
	"AccessArrangement",
	"CardPaymentArrangement",
	"StandingOrderArrangement",
	"OverdraftArrangement",
	"RepaymentArrangement",
	"StatementArrangement",
	"CreditTransferArrangement",
	"ProductAndServiceArrangement",
	"FactoringArrangement",
	"RolloverArrangement",
	"RestructuringArrangement",
	"InsuranceArrangement",
	"CollectionArrangement",
	"UnderwritingArrangement",
	"TerminationArrangement",
	"MaturityArrangement",
}

// validArrangementtypevaluesEnumValue provides a map of Arrangementtypevaluess for fast verification of use input
var validArrangementtypevaluesEnumValues = map[Arrangementtypevalues]struct{}{
	"SweepArrangement": {},
	"WithdrawalArrangement": {},
	"DepositArrangement": {},
	"InterestArrangement": {},
	"FeeArrangement": {},
	"LienArrangement": {},
	"InformationArrangement": {},
	"PaymentArrangement": {},
	"EntitlementArrangement": {},
	"PeriodArrangement": {},
	"CollateralArrangement": {},
	"LimitArrangement": {},
	"AccessArrangement": {},
	"CardPaymentArrangement": {},
	"StandingOrderArrangement": {},
	"OverdraftArrangement": {},
	"RepaymentArrangement": {},
	"StatementArrangement": {},
	"CreditTransferArrangement": {},
	"ProductAndServiceArrangement": {},
	"FactoringArrangement": {},
	"RolloverArrangement": {},
	"RestructuringArrangement": {},
	"InsuranceArrangement": {},
	"CollectionArrangement": {},
	"UnderwritingArrangement": {},
	"TerminationArrangement": {},
	"MaturityArrangement": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Arrangementtypevalues) IsValid() bool {
	_, ok := validArrangementtypevaluesEnumValues[v]
	return ok
}

// NewArrangementtypevaluesFromValue returns a pointer to a valid Arrangementtypevalues
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewArrangementtypevaluesFromValue(v string) (Arrangementtypevalues, error) {
	ev := Arrangementtypevalues(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for Arrangementtypevalues: valid values are %v", v, AllowedArrangementtypevaluesEnumValues)
}



// AssertArrangementtypevaluesRequired checks if the required fields are not zero-ed
func AssertArrangementtypevaluesRequired(obj Arrangementtypevalues) error {
	return nil
}

// AssertArrangementtypevaluesConstraints checks if the values respects the defined constraints
func AssertArrangementtypevaluesConstraints(obj Arrangementtypevalues) error {
	return nil
}

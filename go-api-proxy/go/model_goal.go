// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi




// Goal - 
type Goal struct {

	// The type of goal
	GoalType string `json:"GoalType,omitempty"`

	// A definition of the goal, including the intended actions required to achieve the goal
	GoalDefinition string `json:"GoalDefinition,omitempty"`

	// Identifies the interested parties and their roles and responsibilities for the specific actions supporting the achievement of the goal
	GoalOrganisation string `json:"GoalOrganisation,omitempty"`

	// Assessment of how the work is progressing towards achieving the goal, including projections of likely outcomes
	GoalResult string `json:"GoalResult,omitempty"`
}

// AssertGoalRequired checks if the required fields are not zero-ed
func AssertGoalRequired(obj Goal) error {
	return nil
}

// AssertGoalConstraints checks if the values respects the defined constraints
func AssertGoalConstraints(obj Goal) error {
	return nil
}

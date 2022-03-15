/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SelectorItem struct {

	// Example \"SPAN\"
	LocationName string `json:"locationName,omitempty"`

	// Example \"operation\"
	PropertyName string `json:"propertyName,omitempty"`

	// Example \"POST /users/verify\"
	Value string `json:"value,omitempty"`

	// TODO(pov) think about value types?
	ValueType string `json:"valueType,omitempty"`
}

// AssertSelectorItemRequired checks if the required fields are not zero-ed
func AssertSelectorItemRequired(obj SelectorItem) error {
	return nil
}

// AssertRecurseSelectorItemRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SelectorItem (e.g. [][]SelectorItem), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSelectorItemRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSelectorItem, ok := obj.(SelectorItem)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSelectorItemRequired(aSelectorItem)
	})
}

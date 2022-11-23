/*
 * Sample API
 *
 * A sample API that uses a sample-site as an example to demonstrate features in the OpenAPI 3.0 specification
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type User struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	// user status
	Status string `json:"status,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}

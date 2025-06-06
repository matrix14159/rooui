package dom

import (
	"syscall/js"
)

// GetByPath recursively gets the Object's properties, returning a TypeMismatchError if it encounters a non-object while
// descending through the object.
func (p *Object) GetByPath(path ...string) (js.Value, error) {
	current := p.Value
	for _, v := range path {
		if current.Type() != js.TypeObject {
			return js.Value{}, TypeMismatchError{
				Expected: js.TypeObject,
				Actual:   current.Type(),
			}
		}

		current = current.Get(v)
	}
	return current, nil
}

// ExpectByPath is a helper function that calls Get and checks the type of the final result.
// It returns a TypeMismatchError if a non-object is encountered while descending the path or the final type does not
// match with the provided expected type.
func (p *Object) ExpectByPath(expectedType js.Type, path ...string) (js.Value, error) {
	value, err := p.GetByPath(path...)
	if err != nil {
		return js.Value{}, err
	}

	if value.Type() != expectedType {
		return js.Value{}, TypeMismatchError{
			Expected: expectedType,
			Actual:   value.Type(),
		}
	}

	return value, nil
}

package misc

import "slices"

// Embeds calls the Embeds method on the given element and returns the result.
//
// Parameters:
//   - elem: The element to call the Embeds method on.
//
// Returns:
//   - T: The result of the Embeds method.
//   - bool: True if the element has an Embeds method and false otherwise.
func Embeds[T any](elem any) (T, bool) {
	if elem == nil {
		return *new(T), false
	}

	e, ok := elem.(interface{ Embeds() T })
	if !ok {
		return *new(T), false
	}

	return e.Embeds(), true
}

// Innermost retrieves the innermost embedded element of the given type from the provided element.
//
// The function continuously calls the Embeds method to access nested elements until it encounters
// an element that either does not have an Embeds method or is equivalent to the zero value of the
// specified type.
//
// Parameters:
//   - elem: The element from which to extract the innermost embedded element.
//
// Returns:
//   - T: The innermost embedded element of type T.
//   - bool: True if such an element is found and convertible to type T, false otherwise.
func Innermost[T comparable](elem any) (T, bool) {
	if elem == nil {
		return *new(T), false
	}

	zero := *new(T)

	for {
		inner, ok := Embeds[T](elem)
		if !ok || inner == zero {
			break
		}

		elem = inner
	}

	e, ok := elem.(T)
	return e, ok
}

// TowerOfEmbeds returns the tower of embeds of the given element where
// the first element is the innermost base of the element and the last one
// is the given element itself.
//
// The tower stops at the first nil element.
//
// Parameters:
//   - elem: The element.
//
// Returns:
//   - []T: The tower of embeds.
//   - bool: True if the tower is not empty and the last element is of type T, false otherwise.
func TowerOfEmbeds[T comparable](elem any) ([]T, bool) {
	if elem == nil {
		return nil, false
	}

	var tower []T

	for {
		inner, ok := Embeds[T](elem)
		if !ok {
			e, ok := elem.(T)
			if ok {
				tower = append(tower, e)
			}

			return tower, ok
		} else if inner == *new(T) {
			break
		}

		tower = append(tower, inner)
	}

	tower = tower[:len(tower):len(tower)]

	slices.Reverse(tower)

	return tower, true
}

// WithContext calls the WithContext method on the given element and returns the result.
//
// Parameters:
//   - elem: The element to call the WithContext method on.
//   - key: The key to add to the context.
//   - value: The value to add to the context.
//
// Returns:
//   - T: The result of the WithContext method. The type of this value must be T.
//   - bool: True if the element has a WithContext method and false otherwise.
func WithContext[T comparable](elem any, key string, value any) (T, bool) {
	if elem == nil {
		return *new(T), false
	}

	e := elem

	for {
		if f, ok := e.(interface{ WithContext(string, any) T }); ok {
			e := f.WithContext(key, value)
			return e, true
		}

		if f, ok := e.(interface{ Embeds() T }); ok {
			e = f.Embeds()
		} else {
			e, ok := elem.(T)
			if !ok {
				panic("elem is not convertible to T")
			}

			return e, false
		}
	}
}

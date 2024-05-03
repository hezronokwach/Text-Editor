package reloaded

// RemoveElement function removes an element from a slice at the specified index.
// It takes a pointer to a slice and the index of the element to remove as input.
// It modifies the original slice by removing the element at the specified index.
func RemoveElement(slice *[]string, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

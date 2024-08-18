package stack

var (
	isEmpty bool = true
	element []int
	size    int = 0
)

// type element int

func CreateStack() bool {
	return isEmpty

}

func Push(stack_element int) int {
	isEmpty = false
	element = append(element, stack_element)
	size++

	return stack_element
}

func Pop() int {

	isEmpty = true
	var poppedElement int
	size--
	if size > 0 {
		isEmpty = false
	}
	

	poppedElement, element = element[len(element)-1], element[:len(element)-1] // append(element[:])

	return poppedElement

	// return errors.New("popError")
}

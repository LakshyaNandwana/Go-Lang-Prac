package openandcount_test

import (
	"fmt"
	openandcount "writer/openAndCount"
)

func ExampleLineCount() {
	fmt.Println(openandcount.LineCount("./testdata/road.txt"))

	// Output:
	// 8 <nil>
}

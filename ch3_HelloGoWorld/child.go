package ch3_HelloGoWorld

import (
	"fmt"
)

func AddLiteral(a, b int) (addedResult int) {
	addedResult = a + b
	fmt.Println("Added Result: ", addedResult)
	return addedResult
}

package dog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says " + strings.ToUpper(s)
}

func WhenSmall(s string) string {
	return "When the puppy is small it says " + strings.ToLower(s)
}

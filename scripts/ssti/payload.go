package ssti

import (
	"fmt"
	"strings"
)

// Replaces the UNIQUE_NUM placeholder with the given number.
func ReplaceUniqueNum(template string, uniqueNum int) string {
	return strings.ReplaceAll(template, "UNIQUE_NUM", fmt.Sprintf("%d", uniqueNum))
}

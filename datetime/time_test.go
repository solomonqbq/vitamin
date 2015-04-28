package datetime

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	time, e := Parse("2015-08-17 22:00:00")
	fmt.Println(e)
	fmt.Println(time)
}

package numeric

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	fmt.Println(Power(2, -3))
	fmt.Println(Power(2, 3))
	fmt.Println(Power(0, 3))
	fmt.Println(Power(2, 0))
}

package cast

import (
	"fmt"
	"testing"
)

func TestToFloat64SliceE(t *testing.T) {

	list := []string{"10.00001", "20.432414", "30.45245"}
	ret := ToFloat64Slice(list)
	fmt.Println(ret)
}

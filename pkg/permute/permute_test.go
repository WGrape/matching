package permute

import (
	"fmt"
	"strings"
	"testing"
)

func TestCombinationAndImplode(t *testing.T) {
	s := strings.Join(CombinationAndImplode([]string{"123", "456", "789"}), "/")
	if s != "123;456;789/123;456/123;789/456;789/123/456/789" {
		fmt.Println()
		t.Fail()
	}
}

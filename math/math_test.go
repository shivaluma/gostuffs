package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max(t *testing.T) {
	var datas = []struct {
		a, b int
		want int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{1, 1, 1},
		{0, 0, 0},
		{-2, -1, -1},
	}

	for _, data := range datas {
		t.Run(fmt.Sprintf("max(%v, %v)", data.a, data.b), func(t *testing.T) {
			got := Max(data.a, data.b)
			assert.Equal(t, data.want, got)
		})
	}
}

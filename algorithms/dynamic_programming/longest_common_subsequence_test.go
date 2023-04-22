package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestCommonSubsequence(t *testing.T) {

	var datas = []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}

	for _, data := range datas {
		t.Run(data.text1+" "+data.text2, func(t *testing.T) {
			got := LongestCommonSubsequence(data.text1, data.text2)
			assert.Equal(t, data.want, got)
		})
	}
}

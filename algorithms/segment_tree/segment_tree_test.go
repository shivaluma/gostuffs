package segment_tree

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestSegmentTree_Query(t *testing.T) {
	var datas = []struct {
		nums       []int
		start, end int
		want       int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 0, 5, 21},
	}

	for i, data := range datas {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			segmentTree := NewSegmentTree(data.nums)
			got := segmentTree.Query(data.start, data.end)
			assert.Equal(t, data.want, got)
		})
	}
}

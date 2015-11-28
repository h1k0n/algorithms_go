package insertion

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := []int{32, 23, 46, 667, 234, 663, 675, 213, 345, 45, 234, 43}
	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

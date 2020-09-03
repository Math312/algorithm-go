package select_sort

import (
	"algorithm/mysort"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		source := []int{1,2}
		var sort mysort.Sort = SelectSort{}
		sort.Sort(&source)
		if !reflect.DeepEqual(source,[]int{1,2}) {
			t.FailNow()
		}
	})
	t.Run("Common", func(t *testing.T) {
		source := []int{52,88,89,0,1,42,46,2,68,77}
		var sort mysort.Sort = SelectSort{}
		sort.Sort(&source)
		if !reflect.DeepEqual(source,[]int{0,1,2,42,46,52,68,77,88,89}) {
			t.FailNow()
		}
	})
	t.Run("One item", func(t *testing.T) {
		source := []int{1}
		var sort mysort.Sort = SelectSort{}
		sort.Sort(&source)
		if !reflect.DeepEqual(source,[]int{1}) {
			t.FailNow()
		}
	})
	t.Run("None item", func(t *testing.T) {
		source := []int{}
		var sort mysort.Sort = SelectSort{}
		sort.Sort(&source)
		if !reflect.DeepEqual(source,[]int{}) {
			t.FailNow()
		}
	})

}



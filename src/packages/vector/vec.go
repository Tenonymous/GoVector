package vector

import (
	"errors"
	"sort"
)

type Vec []interface{}


func (v *Vec) Size() int {
	return len(*v)
}
func (v *Vec) Empty() bool {
	return v.Size() == 0
}
func (v *Vec) At(index int) (interface{}, error) {
	if index > v.Size() {
		return nil, errors.New("Something")
	}
	return (*v)[index], nil
}

func (v *Vec) Find(elem interface{}) (int, bool) {
	for i := 0; i < v.Size(); i++ {
		if (*v)[i] == elem {
			return i, true
		}
	}
	return -1, false
}

func (v *Vec) Sort(less func(i, j int) bool) {
	sort.Slice(*v, less)
}

func (v *Vec) PushBack(elem interface{}) {
	*v = append(*v, elem)
}

func (v *Vec) PopBack() {
	*v = (*v)[0 : v.Size() - 1]
}
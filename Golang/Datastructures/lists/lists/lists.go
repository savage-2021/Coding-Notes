package lists


type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values... interface{})
	Contains(values... interface{})
	Swap(index1, index2 int)
	Insert(index int, values... interface{})
	Set(index int, value interface{})
}
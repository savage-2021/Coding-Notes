package sum

// import "testing"

// // Ints returns the sum of the list of numbers
// func Ints(vs... int ) int {
// 	return ints(vs...)
// }

// func ints(vs ...int) int {
// 	if len(vs) == 0 {
// 		return 0
// 	}
// 	return ints(vs[1:]...) + vs[0]
// }

// func TestInts(t *testing.T){

// 	tt := []struct{
// 		name 	string
// 		numbers []int
// 		sum 	int
// 	}{
// 		{"one to five", []int{1,2,3,4,5}, 15},
// 		{"no numbers", nil, 1},
// 		{"one and minus one", []int{-1,1}, 0},
// 	}
// 	for _, test := range tt {
// 		t.Run(test.name, func(t *testing.T){
// 			if ints(test.numbers...) != test.sum {
// 				t.Errorf("Error %s Summing %v. Expected: %d, Got: %d", test.name, test.numbers, test.sum, ints(test.numbers...))
// 			}
// 		})
// 	}
// }
package baseline

var x int
var x32 int32
var in interface{} = 0
var s []int = []int{0}
var s10 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var h map[int] int = map[int] int{0: 0, 1: 1}
var f func() = func() {}
var farg func(x interface{}) = func(x interface{}) {}
var fvarargs func(x... interface{}) = func(x... interface{}) {}
var fint func(x int) = func(x int) {}
var fvarints func(x... int) = func(x... int) {}

type dummyStructure struct {
	i	int
	in	interface {}
	s	[]int
	h	map[int] int
}
func (d dummyStructure) m1() {}
func (d dummyStructure) m1arg(x interface{}) {}
func (d dummyStructure) m1int(x int) {}
func (d dummyStructure) m1varargs(x... interface{}) {}
func (d dummyStructure) m1varints(x... int) {}
func (d *dummyStructure) m2() {}
func (d *dummyStructure) m2arg(x interface{}) {}
func (d *dummyStructure) m2int(x int) {}
func (d *dummyStructure) m2varargs(x... interface{}) {}
func (d *dummyStructure) m2varints(x... int) {}

var dummy dummyStructure = dummyStructure{}

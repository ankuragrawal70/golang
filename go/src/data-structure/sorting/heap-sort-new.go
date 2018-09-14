package sorting


//func main() {
//	var heap = new(Heap)
//
//	// heap.HeapSort(array)
//	var initial_array = make([]int, cap)
//
//	for i:=0; i<cap; i++{
//		initial_array[i] = rand.Intn(cap)
//	}
//	// var array = []int{1, 7, 7, 9, 1, 8, 5, 0, 6, 0}
//	c_time := time.Now()
//	fmt.Println(c_time)
//	// fmt.Println("Initial array is", initial_array[0:len(initial_array)])
//	heap.HeapSort(initial_array)
//	e_time := time.Now()
//	fmt.Println(e_time.Sub(c_time))
//
//
//}
//All the code below here should go in another package
type Heap struct {
}

func (heap *Heap) HeapSort(array []int) {
	heap.BuildHeap(array)
	// fmt.Println("after heap cretion is", array)
	for length:= len(array); length > 1; length-- {
		heap.RemoveTop(array, length)
		// fmt.Println("after top removing", array)
	}
}

func (heap *Heap) BuildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
}

func (heap *Heap) RemoveTop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = heap.Left(array, root), heap.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

func (*Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}
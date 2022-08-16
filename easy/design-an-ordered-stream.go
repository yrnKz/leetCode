package main

type OrderedStream struct {
	steam map[int]string
	ptr   int
}

func Constructor(n int) OrderedStream {
	o := OrderedStream{
		steam: make(map[int]string),
		ptr:   1,
	}
	return o
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.steam[idKey] = value
	var r []string
	values, ok := this.steam[this.ptr]
	for ok {
		r = append(r, values)
		this.ptr++
		values, ok = this.steam[this.ptr]
	}
	return r
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

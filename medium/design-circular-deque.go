package main

func main() {
	circularDeque := Constructor(3) // 设置容量大小为3
	circularDeque.InsertLast(1)     // 返回 true
	circularDeque.InsertLast(2)     // 返回 true
	circularDeque.InsertFront(3)    // 返回 true
	circularDeque.InsertFront(4)    // 已经满了，返回 false
	circularDeque.GetRear()         // 返回 2
	circularDeque.IsFull()          // 返回 true
	circularDeque.DeleteLast()      // 返回 true
	circularDeque.InsertFront(4)    // 返回 true
	circularDeque.GetFront()        // 返回 4
}

type MyCircularDeque struct {
	CircularDeque []int
	max           int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{max: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.CircularDeque) >= this.max {
		return false
	}
	this.CircularDeque = append([]int{value}, this.CircularDeque...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.CircularDeque) >= this.max {
		return false
	}
	this.CircularDeque = append(this.CircularDeque, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.CircularDeque) > 0 {
		this.CircularDeque = this.CircularDeque[1:]
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.CircularDeque) > 0 {
		this.CircularDeque = this.CircularDeque[:len(this.CircularDeque)-1]
		return true
	}
	return false
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.CircularDeque) == 0 {
		return -1
	}
	return this.CircularDeque[0]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.CircularDeque) == 0 {
		return -1
	}
	return this.CircularDeque[len(this.CircularDeque)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.CircularDeque) == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if len(this.CircularDeque) == this.max {
		return true
	}
	return false
}

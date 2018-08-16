package queue

/**
 * 通过定义别名的方式扩展系统类型或者别人的类型
 */
//只支持int
//type Queue []int
// interface{} 表示 支持任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	//*q = append(*q, v.(int)) // 限定传入值是 int 类型
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

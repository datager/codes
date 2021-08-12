package pool

type Node struct {
	val  interface{}
	next *Node
}

type Link struct {
	head *Node
	size uint64
	ptr  uint64
}

func (link *Link) Append(value interface{}) *Node {
	node := &Node{
		val:  value,
		next: nil,
	}

	if link.head == nil {
		link.head = node
	} else {
		point := link.head
		for {
			if point.next == nil {
				break
			}
			point = point.next
		}

		point.next = node
	}

	link.size++

	return node
}

func (link *Link) Update(node *Node, value interface{}) {

	if node == nil {
		return
	}

	if link.head == nil {
		return
	}

	var found bool
	point := link.head
	for {
		if point == nil {
			break
		}

		if point == node {
			found = true
			break
		}

		point = point.next
	}

	if found {
		point.val = value
	}
}

func (link *Link) Remove(node *Node) {

	if node == nil {
		return
	}

	if link.head == nil {
		return
	}

	point := link.head
	var pre *Node
	var found bool

	for {
		if point == nil {
			break
		}

		if point == node {
			found = true
			break
		}

		pre = point
		point = point.next
	}

	if found {
		if pre == nil {
			link.head = link.head.next
		} else {
			pre.next = point.next
		}

		link.size--
	}

}

func (link *Link) Next() *Node {

	if link.head == nil {
		return nil
	}

	if link.ptr >= link.size {
		link.ptr = 0
		return nil
	}

	var count uint64
	point := link.head
	for {
		if count == link.ptr {
			break
		}

		point = point.next
		count++
	}

	link.ptr++
	return point
}

func (link *Link) Len() uint64 {
	return link.size
}

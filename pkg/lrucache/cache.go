package lrucache

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value string
}
type Hash map[string]*Node

type Queue struct {
	Head     *Node
	Tail     *Node
	Size     int
	Capicity int
}

type LRUCache struct {
	Queue *Queue
	Hash  Hash
}

func (c *LRUCache) Add(n *Node) {
	fmt.Println("Adding", n.Value)
	tmp := c.Queue.Head.Right // First Element
	tmp.Left = n
	n.Right = tmp
	n.Left = c.Queue.Head
	c.Queue.Head.Right = n // Point head to new node
	c.Queue.Size++         // Increase size of queue

	// Remove from tail if capacity is overflowed
	if c.Queue.Size > c.Queue.Capicity {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *LRUCache) Remove(n *Node) *Node {
	fmt.Println("Removing", n.Value)
	if c.Queue.Size <= 0 {
		panic("Cannot remove from empty queue")
	}

	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left

	delete(c.Hash, n.Value)
	c.Queue.Size--

	// Reset Left and Right
	n.Left = nil
	n.Right = nil

	return n

}
func (c *LRUCache) Check(key string) {
	// Check if the value is present in the cache. If it's present, move it to first of queue
	// If it's not, add it to first of queue
	var node *Node
	if val, ok := c.Hash[key]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: key}
	}
	c.Add(node)
	c.Hash[key] = node
}
func (c *LRUCache) Display() {
	fmt.Print("Head <--> ")
	node := c.Queue.Head.Right
	for i := 0; i < c.Queue.Size; i++ {

		fmt.Printf("%s <--> ", node.Value)
		node = node.Right
	}

	fmt.Println("Tail")

	fmt.Println("Hashed Values:")
	for _, v := range c.Hash {
		fmt.Printf("%s,", v.Value)
	}
	fmt.Println()

}
func NewCache(capicity int) *LRUCache {
	return &LRUCache{Queue: NewQueue(capicity), Hash: Hash{}}
}

func NewQueue(capicity int) *Queue {
	head := &Node{}
	tail := &Node{Left: head}

	head.Right = tail

	return &Queue{Head: head, Tail: tail, Size: 0, Capicity: capicity}
}

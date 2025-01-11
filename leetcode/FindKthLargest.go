package leetcode

type KthLargest struct {
	root *Node
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{k: k}
	for _, num := range nums {
		kthLargest.root = kthLargest.insert(kthLargest.root, num)
	}
	return kthLargest
}

func (this *KthLargest) Add(value int) int {
	this.root = this.insert(this.root, value)
	return this.search(this.root, this.k)
}

func (kl *KthLargest) search(root *Node, k int) int {
	if root == nil {
		return -1
	}

	rightCount := 0
	if root.right != nil {
		rightCount = root.right.count
	}

	if k == rightCount+1 {
		return root.value
	} else if k > rightCount {
		return kl.search(root.left, k-rightCount-1)
	} else {
		return kl.search(root.right, k)
	}
}

func (this *KthLargest) insert(root *Node, value int) *Node {
	if root == nil {
		return newNode(value, 1)
	}
	if root.value > value {
		root.left = this.insert(root.left, value)
	} else {
		root.right = this.insert(root.right, value)
	}
	root.count++
	return root
}

func newNode(value, count int) *Node {
	return &Node{value: value, count: count, left: nil, right: nil}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

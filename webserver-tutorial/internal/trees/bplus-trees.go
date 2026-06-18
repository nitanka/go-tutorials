package trees

/*
- A node can be either a internal or Leaf.
- Node can be internal which has keys and pointer to the children.
- Node can be leaf which has keys and values and pointers to leaf nodes.
- A node cannot have keys > degree, if it is there, then the keys are pushed to the parent.
- Root is the entry point of the B+ Tree

Root:
- Root is also an leaf node if there is less data or number of keys <= degree.
- If number of Keys > degree, then keys in the root will have all the KEYS from index number of keys/2 to last and pointer to the child.
- Then a new node will be created which can be leaf/internal.
- When number or keys > degree, then we split the root to two new nodes and create a new node from the mid element.

Internal:
- Keys Array which are like the index.
- Pointer to the child.
- Child can be leaf/internal.
- Contain only routing information
- When number of keys > degree, the node is split into two nodes.

Leaf:
- Contains the keys and values.
- prev/next pointer to leaf nodes.


- A node is either Internal or Leaf.

- Internal nodes contain separator keys and child pointers.

- Leaf nodes contain keys, values, and next/prev leaf pointers.

- When a node exceeds maxKeys, it is split into two nodes.

- A separator key is promoted (or copied up) to the parent.

- The root is the entry point of the tree.

- Initially the root is also a leaf.

- After the first split, the root becomes an internal node and points to child nodes.

- All actual data exists only in leaf nodes.

- Number of Children are equal to degree + 1.

*/

/*

- How to search:
	- Initial call passes root.

	- If root is a leaf:
		search key in keys[]
		if found:
				return value, true
		else:
				return nil, false

	- If root is an internal node:
       	determine which key range contains the search key

       	choose the corresponding child

       	recursively search that child

	- Continue until a leaf node is reached.

	- Search the leaf node and return the result.
*/

/*

- INSERT(root, key, value)

1. Find target leaf.

2. Insert key in sorted order.

3. If leaf does not overflow:
       done.

4. If leaf overflows:
       split leaf into left/right.

       promotedKey = first key of right leaf

       insert promotedKey into parent.

5. If parent overflows:
       split internal node.

       promote middle key.

6. Continue recursively until:

       a) no overflow

       OR

       b) root overflows

*/

/*
	A much better approach will be:

	the implementation is very messy now. better approach.
	- get to the correct leaf node. - add it there.
	- if after adding the keycount > maxkeys
	- then get the midindex and try to push it to the parent.
	- bubble the split towards the top


	INSERT(root, key)

    1. Find target leaf

    2. Insert key

    3. If leaf overflow:
           split leaf
           return promotedKey upward

    4. Parent receives promotedKey

    5. Parent inserts promotedKey

    6. If parent overflow:
           split parent
           return promotedKey upward

    7. Continue until root

    8. If root overflows:
           create new root

*/

const MAX_KEYS = 2

type node struct {
	keys     [MAX_KEYS + 1]int //to handle overflow while creating the split condition
	keyCount int
	values   [MAX_KEYS + 1]int //to handle overflow while creating the split condition
	isLeaf   bool
	children [MAX_KEYS + 2]*node //to handle internal node overflow while splitting
	nextLeaf *node
	prevLeaf *node
}

type BPtree struct {
	root *node
}

func findLeaf(root *node, key int) *node {

	n := root

	for !n.isLeaf {

		i := 0

		for i < n.keyCount && key >= n.keys[i] {
			i++
		}

		n = n.children[i]
	}

	return n
}

func insert(n *node, key int, value int) (newnode *node, promoted int) {

	if n.isLeaf {
		i := 0
		for i < n.keyCount && key >= n.keys[i] {
			i++
		}
		copy(n.keys[i+1:], n.keys[i:])
		copy(n.values[i+1:], n.values[i:])
		n.keys[i] = key
		n.values[i] = value

		n.keyCount++
		if n.keyCount <= MAX_KEYS {
			return nil, 0
		}
		mid := n.keyCount / 2

		rightNode := &node{isLeaf: true}
		rightNode.nextLeaf = n.nextLeaf

		if n.nextLeaf != nil {
			n.nextLeaf.prevLeaf = rightNode
		}

		n.nextLeaf = rightNode
		rightNode.prevLeaf = n
		r := 0
		for i := mid; i < n.keyCount; i++ {
			rightNode.keys[r] = n.keys[i]
			rightNode.values[r] = n.values[i]
			r++
		}
		n.keyCount = mid
		rightNode.keyCount = r
		promoted = rightNode.keys[0]
		return rightNode, promoted
	}

	i := 0
	for i < n.keyCount && key >= n.keys[i] {
		i++
	}

	rightChild, promoted := insert(n.children[i], key, value)

	if rightChild == nil {
		return nil, 0
	}

	copy(n.keys[i+1:], n.keys[i:])
	n.keys[i] = promoted

	copy(
		n.children[i+2:],
		n.children[i+1:],
	)

	n.children[i+1] = rightChild

	n.keyCount++

	if n.keyCount <= MAX_KEYS {
		return nil, 0
	}

	// split internal node

	mid := n.keyCount / 2

	promoteUp := n.keys[mid]

	right := &node{
		isLeaf: false,
	}

	r := 0
	for j := mid + 1; j < n.keyCount; j++ {
		right.keys[r] = n.keys[j]
		r++
	}
	right.keyCount = r

	c := 0
	for j := mid + 1; j <= n.keyCount; j++ {
		right.children[c] = n.children[j]
		c++
	}

	n.keyCount = mid

	return right, promoteUp

}

func (t *BPtree) Insert(key int, value int) {
	if t.root == nil {
		t.root = &node{isLeaf: true}
	}

	rightChild, promoted := insert(t.root, key, value)

	if rightChild == nil {
		return
	}

	// root split — create a new root above the old root and the new right child
	newRoot := &node{isLeaf: false}
	newRoot.keys[0] = promoted
	newRoot.keyCount = 1
	newRoot.children[0] = t.root
	newRoot.children[1] = rightChild
	t.root = newRoot
}
func search(node *node, key int) (int, bool) {

	if node == nil {
		return 0, false
	}

	if node.isLeaf {
		for i := 0; i < node.keyCount; i++ {
			if node.keys[i] == key {
				return node.values[i], true
			}
		}
		return 0, false
	}

	i := 0
	for i < node.keyCount && key >= node.keys[i] {
		i++
	}

	return search(node.children[i], key)
}

func (t *BPtree) Search(key int) (int, bool) {
	return search(t.root, key)
}

// func findLeaf(node *node) *node {
// 	if node.isLeaf {
// 		return node
// 	}
// 	return findLeaf(node.children)
// }

// func insert(n *node, key int, value int) (split bool, pushedUp int, leftchild *node, rightchild *node) {
// 	//find the leaf
// 	i := 0
// 	if n.isLeaf {
// 		for i < n.keyCount && key >= n.keys[i] {
// 			i++
// 		}
// 		n.keyCount++ //as we found the correct place
// 		if n.keyCount <= MAX_KEYS {
// 			copy(n.keys[i:], n.keys[i+1:])
// 			n.keys[i] = key
// 			n.values[i] = value
// 			return false, 0, nil, nil
// 		}
// 		//we will have to split
// 		midIndex := n.keyCount / 2

// 		//create 2 new nodes left and right
// 		leftleaf := &node{isLeaf: true}
// 		rightleaf := &node{isLeaf: true}

// 		//setting the left keys and values
// 		for i := 0; i < midIndex; i++ {
// 			leftleaf.keys[i] = n.keys[i]
// 			leftleaf.values[i] = n.values[i]
// 		}
// 		leftleaf.keyCount = midIndex
// 		//setting the left keys and values
// 		r := 0
// 		for i := midIndex; i < n.keyCount; i++ {
// 			rightleaf.keys[r] = n.keys[i]
// 			rightleaf.values[r] = n.values[i]
// 			r++
// 		}
// 		rightleaf.keyCount = r

// 		//Now the new nodes don;t have parents, we need to add the left and right parent.
// 		// if parent is root, then promoted index will be added to it. else the internal node keys needs to be updated.
// 		// provided the keyCount <= maxcount
// 		leftleaf.nextLeaf = rightleaf
// 		rightleaf.prevLeaf = leftleaf

// 	//not a leaf, find the proper child, and then go to that child

// 		return true,right.keys[0],leftleaf,rightleaf,
// }

// func (t *BPtree) Insert(key int , value int) {
// 	split, pushedUp, leftchild, rightchild := insert(t.root, key, value)

// 	if !split {
// 		return
// 	}
// 	newRoot := &node{
// 		isLeaf: false,
// 	}

// 	newRoot.keys[0] = pushedUp
// 	newRoot.keyCount = 1

// 	newRoot.children[0] = leftchild
// 	newRoot.children[1] = rightchild

// 	t.root = newRoot
// }

func main() {
	tree := &BPtree{}

	for _, kv := range [][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}} {
		tree.Insert(kv[0], kv[1])
	}

	for _, k := range []int{1, 3, 5, 99} {
		if v, ok := tree.Search(k); ok {
			println("found", k, "->", v)
		} else {
			println("not found", k)
		}
	}
}

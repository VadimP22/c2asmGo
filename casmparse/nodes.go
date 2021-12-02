package casmparse


type Node struct {
	valueType  string
	valueValue string
	childs []*Node
	parent *Node
}


func (n *Node) GetType() string {
	return n.valueType
}


func (n *Node) GetValue() string {
	return n.valueValue
}


func (n *Node) GetChilds() []*Node {
	return n.childs
}


func (n *Node) GetParent() *Node {
	return n.parent
}


func (n *Node) AddChild(childType, childValue string) *Node{
	child := NewNode(childType, childValue)
	child.parent = n
	n.childs = append(n.childs, child)
	return child
}


func NewNode(type_, value string) *Node {
	node := new(Node)
	node.valueType = type_
	node.valueValue = value
	node.parent = nil
	return node
}
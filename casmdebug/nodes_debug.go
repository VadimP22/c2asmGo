package casmdebug


import (
	"../casmparse"
	"fmt"
)


func DisplayTree(root *casmparse.Node)  {
	fmt.Println("BEGIN: casmdebug.DisplayTree()")
	visualizeTree(root, 0)
	fmt.Println("END: casmdebug.DisplayTree()")
}


func visualizeTree(root *casmparse.Node, level int) {
	
	fmt.Println(mulStr("-", level), root.GetType(), ":", root.GetValue())
	
	for _, node := range root.GetChilds() {
		visualizeTree(node, level + 1)
	}

}


func mulStr(str string, n int) string {
	for i := 0; i < n; i++ {
		str = str + str
	}
	return str
}
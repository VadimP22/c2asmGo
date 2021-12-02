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


func visualizeTree(root *casmparse.Node, level uint) {
	printLevelSymbol(level)
	fmt.Println(root.GetType(), ":", root.GetValue())
	
	for _, node := range root.GetChilds() {
		visualizeTree(node, level + 1)
	}

}


func printLevelSymbol(level uint) {
	for i := 0; i < int(level); i++ {
		fmt.Print("-")
	}
}
/* author:owl date:2019/1/28 description:nil */
package main

import "fmt"
import ".."

func main() {
	var root hello.TreeNode
	root = hello.TreeNode{Value: 3}
	root.Left = &hello.TreeNode{}
	root.Right = &hello.TreeNode{5, nil, nil}
	root.Right.Left = new(hello.TreeNode)
	fmt.Println(root.Printv(5))
}

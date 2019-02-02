/* author:owl date:2019/1/27 description:面向对象 */
package hello

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node TreeNode) Printv(a int) int {
	fmt.Println(node.Value)
	return a
}

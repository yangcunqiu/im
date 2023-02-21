package model

type TreeNode struct {
	ID            int
	Pid           int
	Name          string
	Level         int
	TreeChildList []TreeNode
}

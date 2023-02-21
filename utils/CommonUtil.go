package utils

import (
	"im/model"
)

func ListToTree(list []model.TreeNode, pid int) []model.TreeNode {
	var result []model.TreeNode
	for _, node := range list {
		if node.Pid == pid {
			node.TreeChildList = ListToTree(list, node.ID)
			result = append(result, node)
		}
	}
	return result
}

package utils

import "im/model"

func ListToTree(list []model.District, pid uint) []model.DistrictNode {
	// 初始化slice
	nodes := make([]model.DistrictNode, 0)
	for _, district := range list {
		if district.Pid == pid {
			// 构建最顶层节点
			rootNode := model.DistrictNode{
				ID:                    district.ID,
				Pid:                   district.Pid,
				Name:                  district.Name,
				Level:                 district.Level,
				DistrictWrapChildList: getChildByPid(list, district.ID),
			}
			nodes = append(nodes, rootNode)
		}
	}
	return nodes
}

func getChildByPid(list []model.District, pid uint) []model.DistrictNode {
	nodes := make([]model.DistrictNode, 0)
	for _, district := range list {
		if district.Pid == pid {
			node := model.DistrictNode{
				ID:                    district.ID,
				Pid:                   district.Pid,
				Name:                  district.Name,
				Level:                 district.Level,
				DistrictWrapChildList: getChildByPid(list, district.ID),
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

package cladminentity

import "sync"

type DeptInfo struct {
	ID       uint64 `json:"deptId"`
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name"`
	Sore     int32  `json:"sort"`
}

type DeptList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*DeptInfo
}

type DeptTree struct {
	ID       uint64      `json:"deptId"`
	ParentID uint64      `json:"parentId"`
	Name     string      `json:"name"`
	Sort     int32       `json:"sort"`
	Children []*DeptTree `json:"children,omitempty"`
}

func (d *DeptTree) IsEqual(father *DeptTree, child *DeptTree) bool {
	return father.ID == child.ParentID
}
func (d *DeptTree) SetChildren(father *DeptTree, children []*DeptTree) {
	father.Children = children
}
func (d *DeptTree) RetFather(father *DeptTree) bool {
	return father.ParentID == 0
}

package user

import (
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/deptservice"
	"cladmin/service/userservice"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.ID = r.ID
	userModel, errNo := userService.Get([]field.Expr{
		cladminquery.Q.SysUser.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUser.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	roleIDList := make([]uint64, 0)
	slice.ForEach(userModel.Roles, func(_ int, roleModel *cladminmodel.SysRole) {
		roleIDList = append(roleIDList, roleModel.ID)
	})
	//
	var (
		deptModel *cladminmodel.SysDept
		deptName  string
	)
	deptService := deptservice.NewDeptService(c)
	if deptModel, errNo = deptService.Get([]field.Expr{
		cladminquery.Q.SysDept.ID,
		cladminquery.Q.SysDept.Name,
	}, []gen.Condition{
		cladminquery.Q.SysDept.ID.Eq(userModel.DeptID),
	}); errNo == nil && deptModel != nil {
		deptName = deptModel.Name
	}
	//
	handler.SendResponse(c, nil, GetResponse{
		UserID:       userModel.ID,
		Username:     userModel.Username,
		CreateTime:   userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserID: userModel.CreateUserID,
		DeptID:       userModel.DeptID,
		DeptName:     deptName,
		Email:        userModel.Email,
		Mobile:       userModel.Mobile,
		Gender:       userModel.Gender,
		SuperAdmin:   userModel.SuperAdmin,
		Status:       userModel.Status,
		RoleIDList:   roleIDList,
	})
}

func GetPersonalInfo(c *gin.Context) {
	id := c.GetUint64("userID")
	userService := userservice.NewUserService(c)
	userModel, errNo := userService.Get([]field.Expr{
		cladminquery.Q.SysUser.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUser.ID.Eq(id),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	roleIDList := make([]uint64, 0)
	slice.ForEach(userModel.Roles, func(_ int, roleModel *cladminmodel.SysRole) {
		roleIDList = append(roleIDList, roleModel.ID)
	})
	//
	var (
		deptModel *cladminmodel.SysDept
		deptName  string
	)
	deptService := deptservice.NewDeptService(c)
	if deptModel, errNo = deptService.Get([]field.Expr{
		cladminquery.Q.SysDept.ID,
		cladminquery.Q.SysDept.Name,
	}, []gen.Condition{
		cladminquery.Q.SysDept.ID.Eq(userModel.DeptID),
	}); errNo == nil && deptModel != nil {
		deptName = deptModel.Name
	}
	//
	handler.SendResponse(c, nil, GetResponse{
		UserID:       userModel.ID,
		Username:     userModel.Username,
		CreateTime:   userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserID: userModel.CreateUserID,
		DeptID:       userModel.DeptID,
		DeptName:     deptName,
		Email:        userModel.Email,
		Mobile:       userModel.Mobile,
		Gender:       userModel.Gender,
		SuperAdmin:   userModel.SuperAdmin,
		Status:       userModel.Status,
		RoleIDList:   roleIDList,
	})
}

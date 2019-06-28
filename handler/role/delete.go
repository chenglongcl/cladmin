package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/roleservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"sync"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	finished := make(chan bool, 1)
	errorChanel := make(chan *errno.Errno, 1)
	wg := sync.WaitGroup{}
	for _, id := range r.Ids {
		wg.Add(1)
		go func(id uint64) {
			defer wg.Done()
			roleService := roleservice.Role{
				Id: id,
			}
			role, _ := roleService.Get()
			if errNo := roleService.Delete(); errNo != nil {
				errorChanel <- errNo
				return
			}
			inject.Obj.Enforcer.DeleteRole(role.RoleName)
		}(id)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case errNo := <-errorChanel:
		SendResponse(c, errNo, nil)
	case <-finished:
		SendResponse(c, nil, nil)
	}
}

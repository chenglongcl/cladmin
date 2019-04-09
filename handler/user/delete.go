package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/user_service"
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
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	errorChanel := make(chan *errno.Errno, 1)
	for _, id := range r.Ids {
		wg.Add(1)
		go func(id uint64) {
			defer wg.Done()
			userService := user_service.User{
				Id: id,
			}
			user, _ := userService.Get()
			if errNo := userService.Delete(); errNo != nil {
				errorChanel <- errNo
				return
			}
			inject.Obj.Enforcer.DeleteUser(user.Username)
		}(id)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
		SendResponse(c, nil, nil)
	case errNo := <-errorChanel:
		SendResponse(c, errNo, nil)
	}
}

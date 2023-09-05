package user

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/auth"
	"cladmin/pkg/errno"
	"cladmin/pkg/token"
	"cladmin/service/userservice"
	"cladmin/service/usertokenservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"time"
)

// Login
// @Description: 用户登录
// @param c
func Login(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userModel, errNo := userService.Get([]field.Expr{
		cladminquery.Q.SysUser.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUser.Username.Eq(r.Username),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	//Compare the login password with user password
	if err := auth.Compare(userModel.Password, r.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	//user locked
	if userModel.Status != 1 {
		handler.SendResponse(c, errno.ErrDisabledUser, nil)
		return
	}
	// Sign the json web token.
	t, e, re, err := token.Sign(c, token.Context{ID: userModel.ID, Username: userModel.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}
	go func() {
		expireTime, _ := time.ParseInLocation("2006-01-02 15:04:05", e, time.Local)
		RefreshTime, _ := time.ParseInLocation("2006-01-02 15:04:05", re, time.Local)
		userTokenService := usertokenservice.NewUserTokenService(c)
		userTokenService.UserID = userModel.ID
		userTokenService.Token = t
		userTokenService.ExpireTime = expireTime
		userTokenService.RefreshTime = RefreshTime
		_, _ = userTokenService.RecordToken()
	}()
	handler.SendResponse(c, nil, CreateResponse{
		Username:         userModel.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	})
}

// Logout
// @Description: 用户登出
// @param c
func Logout(c *gin.Context) {
	userID := c.GetUint64("userID")
	userTokenService := usertokenservice.NewUserTokenService(c)
	userTokenModel, errNo := userTokenService.Get([]field.Expr{
		cladminquery.Q.SysUserToken.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUserToken.UserID.Eq(userID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if userTokenModel == nil || userTokenModel.UserID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	go func() {
		tokenCtx := &token.Context{
			ID:         userID,
			ExpiredAt:  userTokenModel.ExpireTime.Unix(),
			RefreshExp: userTokenModel.RefreshTime.Unix(),
		}
		token.BLackListToken(userTokenModel.Token, tokenCtx)
	}()
	handler.SendResponse(c, nil, nil)
}

// LogoutLogin
// @Description: 注销管理员登录
// @param c
func LogoutLogin(c *gin.Context) {
	var r LogoutLoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	errGroup := &errgroup.Group{}
	for _, _id := range r.Ids {
		id := _id
		errGroup.Go(func() error {
			userTokenService := usertokenservice.NewUserTokenService(c)
			userTokenModel, _ := userTokenService.Get([]field.Expr{
				cladminquery.Q.SysUserToken.ALL,
			}, []gen.Condition{
				cladminquery.Q.SysUserToken.UserID.Eq(id),
			})
			if userTokenModel != nil && userTokenModel.UserID > 0 {
				tokenCtx := &token.Context{
					ID:         id,
					ExpiredAt:  userTokenModel.ExpireTime.Unix(),
					RefreshExp: userTokenModel.RefreshTime.Unix(),
				}
				token.BLackListToken(userTokenModel.Token, tokenCtx)
			}
			return nil
		})
	}
	if errNo := errGroup.Wait(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

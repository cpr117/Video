// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/dto"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type User struct{}

// 登录
func (*User) Login(c *gin.Context, username, password string) (resp.LoginVo, int) {
	// 检查用户是否存在
	userAuth := dao.GetOne[model.UserAuth]("username", username)
	if userAuth.Id == 0 {
		return *new(resp.LoginVo), r.ERROR_USER_NOT_EXIST
	}
	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		return *new(resp.LoginVo), r.ERROR_PASSWORD_WRONG
	}
	// 获取用户详细信息 DTO
	userDetailDTO := convertUserDetailDTO(c, userAuth)

	// 登录信息正确， 生成Token
	// TODO: 目前只给用户设定一个角色, 获取第一个值就行, 后期优化: 给用户设置多个角色
	// UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	uuid := utils.Encryptor.MD5(userDetailDTO.IpAddress + userDetailDTO.Browser + userDetailDTO.OS)
	// 生成Token
	token, err := utils.GetJWT().GenToken(userAuth.Id, userDetailDTO.RoleLabels[0], uuid)
	if err != nil {
		utils.Logger.Info("登录时生成Token失败", zap.Error(err))
		return *new(resp.LoginVo), r.ERROR_TOKEN_CREATE
	}
	userDetailDTO.Token = token
	// todo: 个人感觉这个应该在用户退出登录时记录
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{Id: userAuth.Id},
		IpAddress:     userDetailDTO.IpAddress,
		IpSource:      userDetailDTO.IpSource,
		LastLoginTime: userDetailDTO.LastLoginTime,
	})

	// 保存用户信息到Session和Redis中
	session := sessions.Default(c)
	// ! session 中只能存储字符串
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	session.Set(KEY_USER+uuid, sessionInfoStr)
	session.Save()
	return userDetailDTO.LoginVo, r.OK
}

func convertUserDetailDTO(c *gin.Context, userAuth model.UserAuth) dto.UserDetailDTO {
	// 获取用户IP相关信息
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	browser, os := "unknown", "unknown"

	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}

	// 获取用户详细信息
	userInfo := dao.GetOne[model.User]("id", userAuth.Id)
	// FIXME: 获取该用户对应的角色, 没有角色默认是 "test"
	roleLabels := roleDao.GetLabelsByUserInfoId(userInfo.Id)
	if len(roleLabels) == 0 {
		roleLabels = append(roleLabels, "test")
	}

	return dto.UserDetailDTO{
		LoginVo: resp.LoginVo{
			Id:           userAuth.Id,
			UserInfoId:   userInfo.Id,
			Email:        userAuth.Email,
			NickName:     userInfo.NickName,
			Coins:        userInfo.Coins,
			IsMember:     userInfo.IsMember,
			Avatar:       userInfo.Avatar,
			PersonalSign: userInfo.PersonalSign,
			//FollowNum:    len(userInfo.Followers),
			//SubNum:        len(userInfo.SubList),
			//UnreadMsg:     len(userInfo.Message),
			IpAddress:     ipAddress,
			IpSource:      ipSource,
			LastLoginTime: time.Now(),
		},
		Password:   userAuth.Password,
		RoleLabels: roleLabels,
		Browser:    browser,
		OS:         os,
		IsDisable:  userInfo.IsDisable,
	}
}

func (*User) Logout(c *gin.Context, username string) {
	userAuth := dao.GetOne[model.UserAuth]("username", username)
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{Id: userAuth.Id},
		LastLoginTime: time.Now(),
	})

	uuid := utils.GetFromContext[string](c, "uuid")
	session := sessions.Default(c)
	session.Delete(KEY_USER + uuid)
	session.Save()
	utils.Redis.Del(KEY_USER + uuid)
}

// 用户注册
func (*User) Register(req req.Register, roleId ...int) int {
	//RoleId := 3
	//if len(level) > 0 {
	//	RoleId = level[0]
	//}
	RoleId := 2
	if len(roleId) != 0 {
		RoleId = roleId[0]
	}
	// 检查用户名已存在 则该账号已经注册过了
	if exist := checkUserExistByName(req.UserName); exist {
		return r.ERROR_USER_NAME_USED
	}
	// VerifyCode 验证码
	//if !utils.VerifyCode(req.Code, req.Email, req.UserName) {
	//	return r.ERROR_VERIFY_CODE
	//}
	userInfo := &model.User{
		NickName: "user:" + req.UserName,
		Coins:    0,
		IsMember: false,
		Avatar:   BiliInfoService.GetBiliConfig().UserAvatar,
	}
	dao.Create(&userInfo)
	// 设置用户默认角色
	dao.Create(&model.UserRole{
		UserId: userInfo.Id,
		RoleId: RoleId, // 默认角色是 "test"
	})
	dao.Create(&model.UserAuth{
		UserInfoId:    userInfo.Id,
		Username:      req.UserName,
		Email:         req.Email,
		Password:      utils.Encryptor.BcryptHash(req.Password),
		LoginType:     1,
		LastLoginTime: time.Now(), // 注册时会更新 "上次登录时间"
	})
	return r.OK
}

// 检查用户名是否已经存在
func checkUserExistByName(username string) bool {
	existUser := dao.Count[model.UserAuth]("username = ?", username)
	return existUser != 0
}

// GetCaptcha 获取验证码
func (*User) GetCaptcha(c *gin.Context, req req.RegisterCode) (code int) {
	// 验证邮箱是否唯一
	if exist := checkUserExistByEmail(req.Email); exist {
		return r.ERROR_EMAIL_USED
	}
	// 生成并发送验证码
	if !utils.SendRegisterCode(req.UserName, req.Email) {
		return r.ERROR_EMAIL_SEND_CODE
	}
	return r.OK
}

func checkUserExistByEmail(email string) bool {
	existUser := dao.GetOne[model.UserAuth]("email = ?", email)
	return existUser.Id != 0
}

// 后台服务
func (*User) GetInfo(userId int) resp.UserDetailVo {
	ret := utils.CopyProperties[resp.UserDetailVo](dao.GetOne[model.User]("id", userId))
	ret.FollowList = userDao.GetFollows(userId)
	ret.FansList = userDao.GetFans(userId)
	ret.FavorList = userDao.GetFavor(userId)
	return ret
}

func (*User) GetList(req req.GetUserList) (ret resp.PageResult[[]resp.UserDetailVo]) {
	data, total := userDao.GetUserList(req)
	ret.List = data
	ret.Total = total
	ret.PageSize = req.PageSize
	ret.PageNum = req.PageNum
	return
}

func (*User) UpdateDisable(req req.UpdateDisable) {
	dao.Update(&model.User{
		Universal: model.Universal{Id: req.ID},
		IsDisable: req.IsDisable,
	})
}

func (*User) DeleteUser(id int) {
	dao.Delete[model.User]("id", id)
}

// 用户权限修改
func (*User) UpdateRole(req req.UpdateRole) {
	admin := dao.GetOne[model.UserRole]("user_id = ?", req.AdminID)
	user := dao.GetOne[model.UserRole]("user_id = ?", req.ID)
	if admin.RoleId > user.RoleId && admin.RoleId > req.RoleID {
		dao.Update(&model.UserRole{
			UserId: req.ID,
			RoleId: req.RoleID,
		})
	}
}

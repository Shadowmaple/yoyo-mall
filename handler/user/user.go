package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/token"
	"yoyo-mall/service/user"

	"github.com/gin-gonic/gin"
)

const (
	AppID  = "wx2a28193a6082cbb3"
	Secret = ""

	LoginURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type WechatLoginResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type LoginReq struct {
	Code string `json:"code"`
}

type LoginResp struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"`
}

func Login(c *gin.Context) {
	req := new(LoginReq)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}
	code := req.Code
	fmt.Println("code: ", code)

	resp, err := requestWechatLogin(code)
	if err != nil {
		handler.SendError(c, errno.ErrWechatServer, nil, err.Error())
		return
	}

	// 数据库用户表操作
	userID, isNew, err := user.Login(resp.OpenID)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	t, err := token.Sign(c, token.Context{ID: uint32(userID), Role: 0, OpenID: resp.OpenID}, "")
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, LoginResp{
		Token: t,
		IsNew: isNew,
	})
}

func requestWechatLogin(code string) (*WechatLoginResp, error) {
	requestURL := fmt.Sprintf(LoginURL, AppID, Secret, code)

	resp, err := http.Get(requestURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dataMp := new(WechatLoginResp)
	if err := json.Unmarshal(body, dataMp); err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("data: %+v\n", dataMp)
	if len(dataMp.OpenID) == 0 {
		err = errors.New("openID is empty, failed")
		return nil, err
	}

	return dataMp, nil
}

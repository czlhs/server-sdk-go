package rcserversdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// UserReslut User 返回信息
type UserReslut struct {
	Token        string `json:"token"`
	UserID       string `json:"userId"`
	BlockEndTime string `json:"blockEndTime"`
}

// BlockListReslut BlockListReslut
type BlockListReslut struct {
	Users []UserInfo `json:"users"`
}

// BlacklistReslut BlacklistReslut
type BlacklistReslut struct {
	Users []string `json:"users"`
}

// UserInfo UserInfo
type UserInfo struct {
	ID           string `json:"userId"`
	BlockEndTime string `json:"blockEndTime,omitempty"`
}

// UserRegister 注册用户，生成用户在融云的唯一身份标识 Token
/*
*@param  userId:用户 Id，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称.用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称。
*@param  portraitUri:用户头像 URI，最大长度 1024 字节.用来在 Push 推送时显示用户的头像。
*
*@return UserReslut, RCError
 */
func (rc *RongCloud) UserRegister(userID, name, portraitURI string) (UserReslut, error) {
	if userID == "" {
		return UserReslut{}, RCErrorNew(20005, "Paramer 'userID' is required")
	}
	if name == "" {
		return UserReslut{}, RCErrorNew(20005, "Paramer 'name' is required")
	}
	if portraitURI == "" {
		return UserReslut{}, RCErrorNew(20005, "Paramer 'portraitUri' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/getToken." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", userID)
	req.Param("name", name)
	req.Param("portraitUri", portraitURI)

	rep, err := req.Bytes()
	if err != nil {
		return UserReslut{}, err
	}

	var code CodeReslut
	var userReslut UserReslut

	if err := json.Unmarshal(rep, &struct {
		*CodeReslut
		*UserReslut
	}{&code, &userReslut}); err != nil {
		return UserReslut{}, err
	}
	if code.Code != 200 {
		return UserReslut{}, RCErrorNew(code.Code, code.ErrorMessage)
	}

	return userReslut, nil
}

// UserUpdate 修改用户信息
/*
*@param  userId:用户 Id，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节。用来在 Push 推送时，显示用户的名称，刷新用户名称后 5 分钟内生效。（可选，提供即刷新，不提供忽略）
*@param  portraitUri:用户头像 URI，最大长度 1024 字节。用来在 Push 推送时显示。（可选，提供即刷新，不提供忽略）
*
*@return error
 */
func (rc RongCloud) UserUpdate(userID, name, portraitURI string) error {
	if userID == "" {
		return RCErrorNew(20005, "Paramer 'userId' is required")
	}
	if name == "" {
		return RCErrorNew(20005, "Paramer 'name' is required")
	}
	if portraitURI == "" {
		return RCErrorNew(20005, "Paramer 'portraitURI' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/refresh." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", userID)
	req.Param("name", name)
	req.Param("portraitUri", portraitURI)

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeReslut
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}

	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// BlockAdd 添加用户到黑名单
/*
*@param  userId:用户 Id。
*@param  blackUserId:被加到黑名单的用户Id。
*
*@return error
 */
func (rc *RongCloud) BlockAdd(id string, minute uint64) error {

	if id == "" {

		return RCErrorNew(20005, "Paramer 'id' is required")
	}

	if minute > 43200 {
		return RCErrorNew(20004, "封禁时间不正确, 当前传入为 , 正确范围 1 - 1 * 30 * 24 * 60 分钟")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/block." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", id)
	req.Param("minute", strconv.FormatUint(minute, 10))

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeReslut
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}

	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// BlockRemove 从黑名单中移除用户
/*
*@param  id:用户 ID。
*
*@return error
 */
func (rc *RongCloud) BlockRemove(id string) error {
	if id == "" {
		return RCErrorNew(20005, "Paramer 'id' is required")
	}
	req := httplib.Post(rc.RongCloudURI + "/user/unblock." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", id)
	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeReslut
	if err := json.Unmarshal(rep, &code); err != nil {
		return RCErrorNew(20100, err.Error())
	}

	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// BlockGetList 获取某用户的黑名单列表
/*
*@return QueryBlockUserReslut
 */
func (rc *RongCloud) BlockGetList() (BlockListReslut, error) {
	req := httplib.Post(rc.RongCloudURI + "/user/block/query." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)

	rep, err := req.Bytes()

	if err != nil {
		return BlockListReslut{}, err
	}

	var dat BlockListReslut
	var code CodeReslut
	if err := json.Unmarshal(rep, &dat); err != nil {
		return BlockListReslut{}, err
	}
	if err := json.Unmarshal(rep, &code); err != nil {
		return BlockListReslut{}, err
	}
	if code.Code != 200 {
		return BlockListReslut{}, RCErrorNew(code.Code, code.ErrorMessage)
	}

	return dat, nil
}

// BlacklistAdd 添加用户到黑名单方法（每秒钟限 100 次）
/*
*@param  userId:用户 Id。
*@param  blackUserId:被加到黑名单的用户Id。
*
*@return error
 */
func (rc *RongCloud) BlacklistAdd(id string, blacklist []string) error {
	if id == "" {
		return RCErrorNew(20005, "Paramer 'id' is required")
	}

	if len(blacklist) == 0 {
		return RCErrorNew(20005, "Paramer 'blacklist' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/blacklist/add." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", id)
	for _, v := range blacklist {
		req.Param("blackUserId", v)
	}

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeReslut
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}

	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}

	return nil
}

// BlacklistRemove 从黑名单中移除用户方法（每秒钟限 100 次）
/*
*@param  userId:用户 Id。
*@param  blackUserId:被移除的用户Id。
*
*@return error
 */
func (rc *RongCloud) BlacklistRemove(id string, blacklist []string) error {
	if id == "" {
		return RCErrorNew(20005, "Paramer 'id' is required")
	}
	if len(blacklist) == 0 {
		return RCErrorNew(20005, "Paramer 'blacklist' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/blacklist/remove." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", id)
	for _, v := range blacklist {
		req.Param("blackUserId", v)
	}

	rep, err := req.Bytes()
	if err != nil {
		return err
	}

	var code CodeReslut
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}

	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// BlacklistGet 获取某用户的黑名单列表方法（每秒钟限 100 次）
/*
*@param  userId:用户 Id。
*
*@return QueryBlacklistUserReslut error
 */
func (rc *RongCloud) BlacklistGet(id string) (BlacklistReslut, error) {
	if id == "" {
		return BlacklistReslut{}, RCErrorNew(20005, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/user/blacklist/query." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("userId", id)

	rep, err := req.Bytes()
	if err != nil {
		return BlacklistReslut{}, err
	}

	var listReslut BlacklistReslut
	var code CodeReslut
	if err := json.Unmarshal(rep, &listReslut); err != nil {
		return BlacklistReslut{}, err
	}
	if err := json.Unmarshal(rep, &code); err != nil {
		return BlacklistReslut{}, err
	}

	if code.Code != 200 {
		return BlacklistReslut{}, RCErrorNew(code.Code, code.ErrorMessage)
	}
	fmt.Println(string(rep))
	return listReslut, nil
}
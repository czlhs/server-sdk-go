package sdk

import (
	"os"
	"testing"
)

func TestRongCloud_ChatRoomCreate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomCreate(
		"chrm01",
		"rcchrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomGet(
		"chrm01",
		500,
		1,
	)
	t.Log(err)
	t.Log(rep)

}

func TestRongCloud_ChatRoomIsExist(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomIsExist(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomDestroy(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomDestroy(
		"chrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomBanAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomBanAdd(
		[]string{"u01"},
		30)
	t.Log(err)
}

func TestRongCloud_ChatRoomBanGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomBanGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomBanRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomBanRemove(
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomBlockAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomBlockAdd(
		"chrm01",
		[]string{"u01"},
		30,
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomBlockGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomBlockGetList(
		"chrm01",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomBlockRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomBlockRemove(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomDemotionAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	txtmsg := "RC:TxtMsg"
	err := rc.ChatRoomDemotionAdd(
		[]string{
			txtmsg,
		},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomDemotionGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomDemotionGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomDemotionRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	txtmsg := "RC:TxtMsg"
	err := rc.ChatRoomDemotionRemove(
		[]string{
			txtmsg,
		},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomDistributionStop(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomDistributionStop(
		"chrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomDistributionResume(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomDistributionResume(
		"chrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomGagAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomGagAdd(
		"chrm01",
		[]string{"u01"},
		30,
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomGagGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomGagGetList(
		"chrm01",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomGagRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomGagRemove(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomKeepAliveAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomKeepAliveAdd(
		"chrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomKeepAliveGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomKeepAliveGetList(
		"chrm01",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomKeepAliveRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomKeepAliveRemove(
		"chrm01",
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomUserWhitelistAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)

	err := rc.ChatRoomUserWhitelistAdd(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomUserWhitelistGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomUserWhitelistGetList(
		"chrm01",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomUserWhitelistRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomUserWhitelistRemove(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomWhitelistAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	txtmsg := "RC:TxtMsg"

	err := rc.ChatRoomWhitelistAdd(
		[]string{txtmsg},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomWhitelistGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomWhitelistGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomWhitelistRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	txtmsg := "RC:TxtMsg"
	err := rc.ChatRoomWhitelistRemove(
		[]string{txtmsg},
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomMuteMembersAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomMuteMembersAdd(
		"chrm01",
		[]string{"u01"},
		30,
	)
	t.Log(err)
}

func TestRongCloud_ChatRoomMuteMembersGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	rep, err := rc.ChatRoomMuteMembersGetList(
		"chrm01",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_ChatRoomMuteMembersRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("RC_APP_ID"),
		os.Getenv("RC_APP_SECRET"),
	)
	err := rc.ChatRoomGagRemove(
		"chrm01",
		[]string{"u01"},
	)
	t.Log(err)
}

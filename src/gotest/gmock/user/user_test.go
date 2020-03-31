package user

import (
	"testing"
	"goms/gotest/gmock/mock"
	"github.com/golang/mock/gomock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)
	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo Error:%v", err)
	}
}
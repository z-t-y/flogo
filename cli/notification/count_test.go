package notification

import (
	"testing"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestNotificationCount(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	_, err = GetNotificationCount(accessToken)
	if err != nil {
		t.Error(err)
	}
}

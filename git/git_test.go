package git

import (
	"fmt"
	"testing"
)

const (
	_REPO = "ssh://git@gitlab.regium.com:33222/xr/configs.git"
	_FILE = "nginx/nginx.conf"
)

func TestGetFile(t *testing.T) {
	file, err := GetFile(_REPO, _FILE)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(file))
}

package whitelist

import (
	"os"

	"github.com/bytedance/sonic"
)

var WhiteListMap []int64

func init() {
	WhiteListMap = make([]int64, 0)
}

type Whitelist struct {
	Or []struct {
		GroupId struct {
			In []int64 `json:".in"`
		} `json:"group_id"`
		UserId interface{} `json:"user_id"`
	} `json:".or"`
}

// WhiteListHandler read whitelist
func WhiteListHandler() []int64 {
	loader, err := os.ReadFile("filter.json")
	if err != nil {
		panic(err)
	}
	var data Whitelist
	sonic.Unmarshal(loader, &data)
	return data.Or[0].GroupId.In
}

func WhiteListUpdater() {
	WhiteListMap = WhiteListHandler()
}

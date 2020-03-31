package setup

import (
	"fmt"
	"goms/pkg/setting"
	"goms/pkg/logging"
	_"goms/pkg/gredis"
	"goms/pkg/gmysql"
	_"goms/pkg/gmgo"
)

func init() {
	fmt.Println("Initing...")
	setting.Setup()
	logging.Setup()
	// gredis.SetUp()
	gmysql.SetUp()
	// gmgo.SetUp()
}
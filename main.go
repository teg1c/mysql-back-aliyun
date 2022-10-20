package main

import (
	"fmt"
	"mysql-back/util"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		panic("must pass into `config.yml` configurations")
	}

	configPath := os.Args[1]

	if !path.IsAbs(configPath) {
		wd, err := os.Getwd()
		if err != nil {
			panic("working directory is not existing")
		}
		configPath = path.Join(wd, configPath)
	}

	// Load config
	conf, err := util.LoadConfig(configPath)
	if err != nil {
		panic("load configuration error: " + err.Error())
	}
	// Init context
	ctx, err := util.NewContext(conf)
	if err != nil {
		panic("init context error: " + err.Error())
	}
	//check dir
	err = util.CheckDir(ctx)
	if err != nil {
		panic("check dir error: " + err.Error())
	}
	//begin back up
	err = util.Back(ctx)
	if err != nil {
		panic("back up error " + err.Error())
	}
	//备份
	err, success := util.PutFileToAliYunOss(ctx)
	if err != nil {
		panic("PutFile error " + err.Error())
	}
	//删除3天前的数据
	util.DelThreeDaysAliYunOssFile(ctx)
	fmt.Println(success)
}

package util

import (
	"fmt"
	"os"
	"time"
)

// CheckDir  判断是否存在目录
func CheckDir(ctx *Context) error {
	_exist, _err := HasDir(ctx.cnf.BackDir)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
	}
	if _exist {
		return _err
	}
	err := os.Mkdir(ctx.cnf.BackDir, os.ModePerm)
	if err != nil {
		fmt.Printf("创建目录异常 -> %v\n", err)
	}
	return err

}

// PutFileToAliYunOss 备份至阿里云oss
func PutFileToAliYunOss(ctx *Context) (error, string) {
	backUpFileName := fmt.Sprintf("%s-%s.sql", ctx.cnf.Database, time.Now().Format("20060102"))
	err := ctx.ossBucket.PutObjectFromFile(fmt.Sprintf("%s/%s", ctx.cnf.Database, backUpFileName), fmt.Sprintf("%s%s", ctx.cnf.BackDir, backUpFileName))
	if err != nil {
		return err, ""
	}
	//删除本地文件
	err = os.Remove(fmt.Sprintf("%s%s", ctx.cnf.BackDir, backUpFileName))
	if err != nil {
		fmt.Println(fmt.Sprintf("删除本地文件失败 %s", fmt.Sprintf("%s%s", ctx.cnf.BackDir, backUpFileName)))
	}
	return err, "备份成功:" + fmt.Sprintf("%s/%s", ctx.cnf.Database, backUpFileName)
}

// DelThreeDaysAliYunOssFile 删除三天前的备份文件
func DelThreeDaysAliYunOssFile(ctx *Context) bool {

	threeDays := []int{-3, -4, -5}

	var files []string

	for _, cnt := range threeDays {
		backUpFileName := fmt.Sprintf("%s-%s.sql", ctx.cnf.Database, time.Now().AddDate(0, 0, cnt).Format("20060102"))
		files = append(files, fmt.Sprintf("%s/%s", ctx.cnf.Database, backUpFileName))
	}

	delRes, err := ctx.ossBucket.DeleteObjects(files)
	if err != nil {
		return false
	}
	fmt.Println("Deleted Objects:", delRes.DeletedObjects)

	return true

}

// HasDir  判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

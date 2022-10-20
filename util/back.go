package util

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"os/exec"
	"time"
)

func Back(c *Context) error {
	timeout := 20
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout+5)*time.Second)
	defer cancel()
	call := spew.Sprintf("docker exec %s mysqldump -h%s -u%s -p%s  --databases %s >%s", c.cnf.ContainerName, c.cnf.MysqlHost, c.cnf.MysqlUsername, c.cnf.MysqlPassword, c.cnf.Database, c.cnf.FullPath)
	fmt.Println("exec command => ", call)
	cmdArr := []string{"-c", call}
	cmd := exec.CommandContext(ctx, "bash", cmdArr...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("> 1 执行命令失败", err.Error())
		os.Exit(0)
		return err
	}
	if ctx.Err() != nil {
		fmt.Println("> 2 执行命令失败", ctx.Err().Error())
		return ctx.Err()
	}
	fmt.Println("out", string(out))
	return err
}

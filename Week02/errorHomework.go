//作业
//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//我的做法是在dao层wrap error抛给上一层，如果按照业务逻辑需要吞掉ErrorNoRows这个错误则在service层做处理
//否则在service层直接返回错误，抛给调用service的地方记录日志或则做其他处理
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func main() {
	err := service()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("success")
}

func service() error {
	err := dao()
	if err != nil {
		return err
	}
	fmt.Println("suncess")
	return nil
}

func dao() error {
	err := sql.ErrNoRows
	return errors.Wrap(err, "error in dao,failed to find")
}

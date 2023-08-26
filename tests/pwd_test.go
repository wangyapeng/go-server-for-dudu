// pwd_test.go
package test

import (
	"dudu/util"
	"testing"
)

// go test -v utils/pwd_test.go utils/pwd.go
func TestA(t *testing.T) {
	t.Log("TestA")
}

func TestPasswordHash(t *testing.T) {
	t.Log("--> TestHelloWorld ")

	pwd := "admin123"
	hash, _ := util.GetPwd(pwd)

	t.Log("--> 输入密码:", pwd)
	t.Log("--> 生成hash:", hash)
	// $2a$10$lRewHvjtPrYZK4TQy.TWDemBMqwIEy/.IVoB7x/2KfqrjzZJNP2ia
	// $2a$10$KEl9ZHfD4gAPu/hgXAjAm.TLgWi5ce7EzBgTdhBfW5IOimtOSfin2

	match := util.ComparePwd(pwd, string(hash))
	t.Log("--> 验证结果:", match)
}

func TestPasswordVerify(t *testing.T) {
	t.Log("--> TestpwdVerify ")

	pwd := "admin2"
	// PHP 生成的密码为 $2y$ 开头 (PHP实现 $2a$ 时有bug，修复时改成了 $2y$)
	hash := "$2y$10$f7ZKW1ZOR4UzGM37.GTmTuY6RmJHknfSwhBacki.Yro1I1kIddEiu"

	match := util.ComparePwd(pwd, hash)
	t.Log("--> TestpwdVerify 验证结果:", match)
	if match == false {
		t.Errorf("TestpwdVerify failed. Got false, expected true.")
	}
}

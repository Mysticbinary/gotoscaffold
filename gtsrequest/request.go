package gtsrequest

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

// 发送邮件
/**
使用例子:
	mailConn := map[string]string{
		"from": "security-xx@xx.com",
		"user": "security-xx",
		"pass": "xxx",
		"host": "mail.xx.com",
		"port": "25",
	}
	To := "xx.xx@xx.com"
	Cc := []string{
		"xx.xx@xx.com",
	}
	subject := "Hello subject"
	body := "Hello, there is body, by gomail sent"
	err := gtsrequest.SendEmail(subject, body, mailConn, To, Cc...)
	if err != nil {
		log.Println(err)
		fmt.Println("email send fail")
	}
*/
func SendEmail(subject string, body string, emailConn map[string]string, To string, Cc ...string) error {
	port, _ := strconv.Atoi(emailConn["port"]) // 转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(emailConn["from"], "XX官方")) //添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+emailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",emailConn["user"])</code> 使用者可以自行实验看效果

	m.SetHeader("To", To)    //只能有一个收件人
	m.SetHeader("Cc", Cc...) //可以有多个抄送人
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(emailConn["host"], port, emailConn["user"], emailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

// No thing.
func PrintTheAuthorHeart() {
	fmt.Println(`
			欲买桂花同载酒，
			终不似，少年游。
`)
}

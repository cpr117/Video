// @User CPR
package utils

import (
	"VideoWeb/config"
	"crypto/tls"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/smtp"
)

const (
	RegisterCode = "注册验证"
)

// 参数： 目标邮箱， 邮件标题， 邮件内容
func SendEmail(email, title, code string) bool {

	// 设置邮件头
	header := make(map[string]string)
	header["From"] = config.Cfg.Email.User
	header["To"] = email
	header["Subject"] = title
	header["Content-Type"] = "text/html; charset=UTF-8"
	// 设置邮件内容
	body := "<h3>尊敬的用户：</h3><p>您好! 您的验证码是 <span style='color:red'> " + code + "</span>，五分钟内有效，祝您生活愉快！</p>"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 设置邮件服务器信息
	auth := smtp.PlainAuth("", config.Cfg.Email.User, config.Cfg.Email.Password, config.Cfg.Email.Host)

	// 发送邮件
	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", config.Cfg.Email.Host, config.Cfg.Email.Port),
		auth,
		config.Cfg.Email.User,
		[]string{email},
		[]byte(message),
	)
	if err != nil {
		Logger.Error("Send email error: ", zap.Error(err))
		return false
	}
	return true
}

func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	c, err := dial(addr)
	if err != nil {
		Logger.Error("Create smpt client error:" + err.Error())
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				Logger.Error("Error during AUTH:" + err.Error())
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		Logger.Error("Dialing Error:" + err.Error())
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

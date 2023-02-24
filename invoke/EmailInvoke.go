package invoke

import (
	"im/global"
	"net/smtp"
)

func SendEmail(subject string, htmlStr string, toAddrList ...string) (err error) {
	em := global.EmailClient
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = global.Config.Email.FromAccount

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = toAddrList

	// 设置主题
	em.Subject = subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.HTML = []byte(htmlStr)

	// 设置服务器相关的配置
	return em.Send(global.Config.Email.EmailServerAddr, smtp.PlainAuth("", global.Config.Email.Username, global.Config.Email.AuthorizationCode, global.Config.Email.Host))
}

package email

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "email"
	Comment       = "send email using http"
	RoutePrefix   = "/"
	DefaultEnable = false
)

var Module = module.Module{
	ID:      ID,
	Comment: Comment,
	// gin route
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	// cobra flag
	FlagFunc: ServeFlag,
}

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)
	serveCmd.Flags().String("mail.username", "", "smtp (发件人)用户名")
	serveCmd.Flags().String("mail.password", "", "smtp (发件人)密码")
	serveCmd.Flags().String("mail.smtpserver", "", "smtp 服务器地址")
	serveCmd.Flags().String("mail.smtpserverport", "", "smtp 服务器地址")
	serveCmd.Flags().Bool("mail.insecureskipverify", true, "是否跳过证书验证")
}

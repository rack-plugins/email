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
	serveCmd.Flags().String("email.prefix", "mailto", "自定义email 路由地址")
	serveCmd.Flags().String("email.username", "", "smtp (发件人)用户名")
	serveCmd.Flags().String("email.password", "", "smtp (发件人)密码")
	serveCmd.Flags().String("email.smtpserver", "", "smtp 服务器地址")
	serveCmd.Flags().String("email.smtpserverport", "", "smtp 服务器地址")
	serveCmd.Flags().Bool("email.insecureskipverify", true, "是否跳过证书验证")
}

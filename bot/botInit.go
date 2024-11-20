package bot

import (
	"context"
	"esctasy-bot-gin/config/configuration"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/token"
	"time"
)

func Init() {
	ctx := context.Background()
	t := token.BotToken(configuration.Config.Bot.AppId, configuration.Config.Bot.AccessToken)
	api := botgo.NewOpenAPI(t).WithTimeout(3 * time.Second) // 使用NewSandboxOpenAPI创建沙箱环境的实例

	user, err := api.Me(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(user.Username)

	pager := new(dto.GuildPager)
	pager.Limit = "10"
	meGuilds, _ := api.MeGuilds(ctx, pager)
	fmt.Println(meGuilds[0].Name)
}

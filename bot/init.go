package bot

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/token"
	"time"
)

func Init() {
	ctx := context.Background()
	t := token.BotToken(102152527, "Qv1oFY0h2Grl6d2zdhKq41ZOjVt2zMXR")
	api := botgo.NewOpenAPI(t).WithTimeout(3 * time.Second) // 使用NewSandboxOpenAPI创建沙箱环境的实例

	user, err := api.Me(ctx)
	if err != nil {
		//log.Error(err)
		logrus.Debug(err)
		return
	}

	fmt.Println(user.Username)

	pager := new(dto.GuildPager)
	pager.Limit = "10"
	meGuilds, _ := api.MeGuilds(ctx, pager)
	fmt.Println(meGuilds[0].Name)
}

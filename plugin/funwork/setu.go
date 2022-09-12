// 使用国内源的iw233 + 新背景实现
package funwork

import (
	"time"

	"github.com/FloatTech/floatbox/web"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	limit = rate.NewManager[int64](time.Minute*3, 8)
)

func init() {
	engine.OnFullMatch("来份二次元").SetBlock(true).Limit(ctxext.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.RequestDataWith(web.NewDefaultClient(), "https://mirlkoi.ifast3.vipnps.vip/api.php?sort=random&type=json", "GET", Referer, ua)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		picURL := gjson.Get(string(data), "pic.0").String()

		messageID := ctx.SendChain(message.Image(picURL))
		time.Sleep(time.Second * 20)
		ctx.DeleteMessage(messageID)
	})

	engine.OnFullMatch("来份星空").SetBlock(true).Limit(ctxext.LimitByUser).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.RequestDataWith(web.NewDefaultClient(), "https://mirlkoi.ifast3.vipnps.vip/api.php?sort=xing&type=json", "GET", Referer, ua)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		picURL := gjson.Get(string(data), "pic.0").String()
		ctx.SendChain(message.Image(picURL))
	})

	engine.OnFullMatch("来份兽耳").SetBlock(true).Limit(ctxext.LimitByUser).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.RequestDataWith(web.NewDefaultClient(), "https://mirlkoi.ifast3.vipnps.vip/api.php?sort=cat&type=json", "GET", Referer, ua)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		picURL := gjson.Get(string(data), "pic.0").String()
		messageID := ctx.SendChain(message.Image(picURL))
		time.Sleep(time.Second * 20)
		ctx.DeleteMessage(messageID)
	})

	engine.OnFullMatch("来份白毛").SetBlock(true).Limit(ctxext.LimitByUser).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.RequestDataWith(web.NewDefaultClient(), "https://mirlkoi.ifast3.vipnps.vip/api.php?sort=yin&type=json", "GET", Referer, ua)
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		picURL := gjson.Get(string(data), "pic.0").String()
		messageID := ctx.SendChain(message.Image(picURL))
		time.Sleep(time.Second * 20)
		ctx.DeleteMessage(messageID)
	})
	engine.OnFullMatch("来份猫猫表情包").SetBlock(true).Limit(ctxext.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.GetData("https://img.moehu.org/pic.php?id=miao&return=json")
		if err != nil {
			ctx.SendChain(message.Text("ERR:", err))
			return
		}
		picURL := gjson.Get(string(data), "acgurl").String()
		ctx.Send(message.Image(picURL))
	})
	engine.OnFullMatch("来份兽耳酱表情包").SetBlock(true).Limit(ctxext.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return
		}
		data, err := web.GetData("https://img.moehu.org/pic.php?id=kemomimi&return=json")
		if err != nil {
			ctx.SendChain(message.Text("ERR:", err))
			return
		}
		picURL := gjson.Get(string(data), "acgurl").String()
		ctx.Send(message.Image(picURL))
	})
}

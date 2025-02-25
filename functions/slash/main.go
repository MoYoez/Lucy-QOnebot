// Package slash https://github.com/Rongronggg9/SlashBot
package slash

import (
	"regexp"
	"strconv"
	"strings"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	// so noisy and try not to use this.
	engine = control.Register("slash", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Help:             "Hi NekoPachi!\n",
	})
)

func init() {
	engine.OnRegex(`^/(.*)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		getPatternInfo := ctx.State["regex_matched"].([]string)[1]
		re := regexp.MustCompile(`\[CQ:image.*?]`)
		getPatternInfo = re.ReplaceAllString(getPatternInfo, "")
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(ctx.CardOrNickName(ctx.Event.UserID)+getPatternInfo+"了他自己~"))
	})
	/*
		Params:
			/rua [CQ:at,qq=123123] || match1 = /rua | match2 = cq... | match3 = id
			match4 match 5 match 6
	*/
	engine.OnRegex(`^(/.*)(\[CQ:at,qq=(.*)\])|^(\[CQ:at,qq=(.*)\])\s(/.*)`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		getMatchedQID := ctx.State["regex_matched"].([]string)[3]
		var getMatchedInfo string
		if getMatchedQID == "" {
			getMatchedQID = ctx.State["regex_matched"].([]string)[5]
			getMatchedInfo = ctx.State["regex_matched"].([]string)[6]
		} else {
			getMatchedInfo = ctx.State["regex_matched"].([]string)[1]
		}
		// use matchedinfo
		qidToInt64, _ := strconv.ParseInt(getMatchedQID, 10, 64)
		getUserInfo := ctx.CardOrNickName(qidToInt64)
		getPersentUserinfo := ctx.CardOrNickName(ctx.Event.UserID)
		// split info
		modifyInfo := strings.ReplaceAll(getMatchedInfo, "/", "")
		splitInfo := strings.Split(modifyInfo, " ")
		re := regexp.MustCompile(`\[CQ:image.*?]`)
		if len(splitInfo) == 2 {
			splitInfo[0] = re.ReplaceAllString(splitInfo[0], "")
			splitInfo[1] = re.ReplaceAllString(splitInfo[1], "")
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(getPersentUserinfo+" "+splitInfo[0]+"了"+getUserInfo+splitInfo[1]))
		} else {
			splitInfo[0] = re.ReplaceAllString(splitInfo[0], "")
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(getPersentUserinfo+" "+splitInfo[0]+"了"+getUserInfo))
		}
	})

	engine.OnRegex(`^(\[CQ:reply,id=(.*)\])\s/(.*)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		getPatternUserMessageID := ctx.State["regex_matched"].([]string)[2]
		getPatternInfo := ctx.State["regex_matched"].([]string)[3]
		getSplit := strings.Split(getPatternInfo, " ")
		rsp := ctx.CallAction("get_msg", zero.Params{
			"message_id": getPatternUserMessageID,
		}).Data.String()
		sender := gjson.Get(rsp, "sender.user_id").Int()
		re := regexp.MustCompile(`\[CQ:image.*?]`)
		modifiedReplyHeader := regexp.MustCompile(`\[CQ:reply.*?]`)
		if len(getSplit) == 2 {
			getSplit[0] = re.ReplaceAllString(getSplit[0], "")
			getSplit[0] = modifiedReplyHeader.ReplaceAllString(getSplit[0], "")
			getSplit[1] = re.ReplaceAllString(getSplit[1], "")
			getSplit[1] = modifiedReplyHeader.ReplaceAllString(getSplit[1], "")
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(ctx.CardOrNickName(ctx.Event.UserID)+" "+getSplit[0]+"了 "+ctx.CardOrNickName(sender)+" "+getSplit[1]))
		} else {
			getSplit[0] = re.ReplaceAllString(getSplit[0], "")
			getSplit[0] = modifiedReplyHeader.ReplaceAllString(getSplit[0], "")
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(ctx.CardOrNickName(ctx.Event.UserID)+" "+getPatternInfo+"了 "+ctx.CardOrNickName(sender)))
		}
	})

}

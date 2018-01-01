package topsdk_test

import (
	"fmt"
	"testing"
	"topsdk"
)

func TestWirelessShareTpwdCreateParam(t *testing.T) {
	topClient := topsdk.NewTopClient(AppKey, AppSecrect)
	var p = topsdk.WirelessShareTpwdCreateParam{}
	p.URL = "https://uland.taobao.com/coupon/edetail?e=B1FFBiV9wvQGQASttHIRqcVGbV2UvfDWZ9DeC3ao2EwQzxWZZqjo%2BUc8aMSBvzz383CW3vZnaBuTMjBa%2FrKbTr9fwBwwUiql6DF%2F2MLpJMq9GaVAvC2u%2FW7PVn13QcLNgPRfTgnhrZM%3D&traceId=0ab2511b15148014109475395e&activityId=cd5d551e3ab74071aa2de096981a4501"
	p.Text = "S925纯银蝴蝶结耳钉女气质韩国个性简约百搭网红耳坠甜美耳环耳饰"
	fmt.Println("===== WirelessShareTpwdCreateParam =====")
	resp, err := topClient.Request(p)
	fmt.Println("resp:", resp, "err:", err)
	t.Log(resp, err)
}

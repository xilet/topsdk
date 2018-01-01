package topsdk_test

import (
	"fmt"
	"testing"
	"topsdk"
)

func TestTbkCouponGetParam(t *testing.T) {
	topClient := topsdk.NewTopClient(AppKey, AppSecrect)
	var p = topsdk.TbkCouponGetParam{}
	p.Me = "nfr%2BYTo2k1PX18gaNN%2BIPkIG2PadNYbBnwEsv6mRavWieOoOE3L9OdmbDSSyHbGxBAXjHpLKvZbL1320ML%2BCF5FRtW7N7yJ056Lgym4X01A%3D"
	fmt.Println("===== TbkCouponGetParam =====")
	resp, err := topClient.Request(p)
	fmt.Println("resp:", resp, "err:", err)
	t.Log(resp, err)
}

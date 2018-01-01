package topsdk_test

import (
	"fmt"
	"testing"
	"topsdk"
)

func TestTbkItemInfoGetParam(t *testing.T) {
	topClient := topsdk.NewTopClient(AppKey, AppSecrect)
	var p = topsdk.TbkItemInfoGetParam{}
	p.Fields = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick"
	p.NumIIDs = "558737509620"
	fmt.Println("===== TbkItemInfoGetParam =====")
	resp, err := topClient.Request(p)
	fmt.Println("resp:", resp, "err:", err)
	t.Log(resp, err)
}

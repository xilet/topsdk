代码来源于 https://github.com/smartwalle/taobao,感谢原作者提供的代码

这里主要就是做了一些体力劳动

淘宝客 Golang SDK
目前已经实现API：

淘宝客基础API
taobao.tbk.item.get    淘宝客商品查询
taobao.tbk.item.recommend.get    淘宝客商品关联推荐查询   已完成
taobao.tbk.item.info.get    淘宝客商品详情（简版）                已完成
taobao.tbk.shop.get    淘宝客店铺查询
taobao.tbk.shop.recommend.get    淘宝客店铺关联推荐查询
taobao.tbk.uatm.event.get    枚举正在进行中的定向招商的活动列表
taobao.tbk.uatm.event.item.get    获取淘宝联盟定向招商的宝贝信息
taobao.tbk.uatm.favorites.item.get    获取淘宝联盟选品库的宝贝信息
taobao.tbk.uatm.favorites.get    获取淘宝联盟选品库列表
taobao.tbk.ju.tqg.get    淘抢购api
taobao.tbk.spread.get    物料传播方式获取
taobao.ju.items.search    聚划算商品搜索接口
taobao.tbk.dg.item.coupon.get    好券清单API【导购】
taobao.tbk.coupon.get    阿里妈妈推广券信息查询
taobao.tbk.tpwd.create    淘宝客淘口令                       已完成

淘口令基础包    淘口令查询、解析、生成API
taobao.wireless.share.tpwd.create    生成淘口令             已完成
taobao.wireless.share.tpwd.query    查询解析淘口令       已完成

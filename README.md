
#有赞 API golang SDK

目前已开放：商品接口，商品类目接口，交易接口，用户接口，工具接口，物流接口，店铺接口


已完成项目前面标记了√



## 联系方式

邮箱:        weibaohui@gmail.com
微信群聊:    ![群聊二维码]()


##通讯协议

调用 API 可通过以下两种通讯协议：

1、AccessToken 免签名通讯协议 适用于：ERP开发商、ISV开发商

√ 2、AppId / AppSecret 签名通讯协议 适用于：个人开发者、单店铺开发者

全局错误返回码说明

##商品接口

kdt.item.add 新增一个商品

kdt.item.delete 删除一个商品

kdt.item.get 得到单个商品信息

kdt.item.sku.update 更新SKU信息

kdt.item.update 更新单个商品信息

kdt.item.update.delisting 商品下架

kdt.item.update.listing 商品上架

kdt.items.custom.get 根据商品货号获取商品

kdt.items.inventory.get 获取仓库中的商品列表

kdt.items.onsale.get 获取出售中的商品列表

kdt.skus.custom.get 根据外部编号取商品Sku

##商品类目接口

kdt.itemcategories.get 获取商品分类二维列表

kdt.itemcategories.promotions.get 获取商品推广栏目列表

kdt.itemcategories.tags.get 获取商品自定义标签列表

kdt.itemcategories.tags.getpage 分页获取商品自定义标签列表

##物流接口

kdt.logistics.online.confirm 卖家确认发货

kdt.logistics.online.marksign 卖家标记签收

kdt.logistics.trace.search 物流流转信息查询

##店铺接口

√ kdt.shop.basic.get 获取店铺基本信息

##交易接口

kdt.trade.close 卖家关闭一笔交易

√ kdt.trade.get 获取单笔交易的信息

kdt.trade.memo.update 修改一笔交易备注

√ kdt.trades.sold.get 查询卖家已卖出的交易列表

##用户接口

kdt.users.weixin.follower.get 根据微信粉丝用户的 openid 或 user_id 获取用户信息

kdt.users.weixin.follower.gets 根据多个微信粉丝用户的 openid 或 user_id 获取用户信息

kdt.users.weixin.followers.get 查询微信粉丝用户信息

##工具接口

√ kdt.regions.get 获取区域地名列表信息



#使用方法
###获取待发货订单
```
	appId := "xxxxxxxxx"
	appSecret := "xxxxxxxxxxxxxxxxxxxxxxxx"


	clt := trade.NewClient(appId, appSecret, nil)
	sold := new(request.Sold)
	sold.UseHasNext = false
	sold.Status = request.Status_WAIT_SELLER_SEND_GOODS
	sold.PageNo = 2
	sold.PageSize = 10
	list, err := clt.GetSold(sold)

```

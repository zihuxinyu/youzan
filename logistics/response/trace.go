package response
//物流流转信息查询结果
type TraceSearchResult struct {
	OutSid      string `json:"out_sid"`                    //快递单号（具体一个物流公司的真实快递单号）
	CompanyName string `json:"company_name"`               //物流公司名称
	Tid         string `json:"tid"`                        //交易编号
	Status      string `json:"status"`                     //订单的物流状态：
                                                           //在途，即货物处于运输过程中；
                                                           //揽件，货物已由快递公司揽收并且产生了第一条跟踪信息；
                                                           //疑难，货物寄送过程出了问题；
                                                           //签收，收件人已签收；
                                                           //退签，即货物由于用户拒签、超区等原因退回，而且发件人已经签收；
                                                           //派件，即快递正在进行同城派件；
                                                           //退回，货物正处于退回发件人的途中；
                                                           //转单
	TraceList   []TradeLogisticsTrack  `json:"trace_list"` //流转信息列表，按时间倒序排序

}

//物流跟踪信息数据结构
type TradeLogisticsTrack struct {
	StatusTime string `json:"status_time"` //状态发生的时间
	StatusDesc string `json:"status_desc"` //状态描述
}

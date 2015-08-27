package request
import "github.com/zihuxinyu/youzan"

const (
	OutStype_shentongEwuliu    int = 1           //      申通E物流
	OutStype_yuantongsudi    int = 2            //      圆通速递
	OutStype_zhongtongsudi    int = 3            //      中通速递
	OutStype_yundakuaiyun    int = 4         //      韵达快运
	OutStype_tiantiankuaidi    int = 5         //      天天快递
	OutStype_baishihuitong    int = 6         //      百世汇通
	OutStype_shunfengsuyun    int = 7         //      顺丰速运
	OutStype_youzhengguonaxiaobao    int = 8         //      邮政国内小包
	OutStype_EMSjingjikuaidi    int = 10        //       EMS经济快递
	OutStype_EMS    int = 11                    //       EMS
	OutStype_youzhengpingyou    int = 12        //       邮政平邮
	OutStype_debangkuaidi    int = 13        //       德邦快递
	OutStype_lianhaotong    int = 16        //       联昊通
	OutStype_quanfengkuaidi    int = 17        //       全峰快递
	OutStype_quanyikuaidi    int = 18        //       全一快递
	OutStype_chengshi100    int = 19        //       城市100
	OutStype_huiqiangkuaidi    int = 20        //       汇强快递
	OutStype_guangdongEMS    int = 21        //       广东EMS
	OutStype_suer    int = 22                   //       速尔
	OutStype_feikangdasuyun    int = 23        //       飞康达速运
	OutStype_zhaijisong    int = 25        //       宅急送
	OutStype_lianbangkuaidi    int = 27        //       联邦快递
	OutStype_debangwuliu    int = 28        //       德邦物流
	OutStype_zhongtiekuaiyun    int = 30        //       中铁快运
	OutStype_xinfengwuliu    int = 31        //       信丰物流
	OutStype_longbangsudi    int = 32        //       龙邦速递
	OutStype_tiandihuayu    int = 33        //       天地华宇
	OutStype_kuaijiesudi    int = 34        //       快捷速递
	OutStype_xinbangwuliu    int = 36        //       新邦物流
	OutStype_nengdasudi    int = 37        //       能达速递
	OutStype_yousukuaidi    int = 38        //       优速快递
	OutStype_guotongkuaidi    int = 40        //       国通快递
	OutStype_qita    int = 41        //       其他
	OutStype_shunfengkuaidi    int = 42        //       顺丰快递
	OutStype_AAE    int = 43        //       AAE
	OutStype_anxinda    int = 44        //       安信达
	OutStype_baifudongfang    int = 45        //       百福东方
	OutStype_BHT    int = 46        //       BHT
	OutStype_bangsongwuliu    int = 47        //       邦送物流
	OutStype_chuanxiwuliu    int = 48        //       传喜物流
	OutStype_datianwuliu    int = 49        //       大田物流
	OutStype_Dsukuaidi    int = 50        //       D速快递
	OutStype_disifang    int = 51        //       递四方
	OutStype_feikangdawuliu    int = 52        //       飞康达物流
	OutStype_feikuaida    int = 53        //       飞快达
	OutStype_fankerufengda    int = 54        //       凡客如风达
	OutStype_fengxingtianxia    int = 55        //       风行天下
	OutStype_feibaokuaidi    int = 56        //       飞豹快递
	OutStype_gangzhongnengda    int = 57        //       港中能达
	OutStype_guangdongyouzheng    int = 58        //       广东邮政
	OutStype_gongsuda    int = 59        //       共速达
	OutStype_huitongkuaiyun    int = 60        //       汇通快运
	OutStype_huayuwuliu    int = 61        //       华宇物流
	OutStype_hengluwuliu    int = 62        //       恒路物流
	OutStype_huaxialong    int = 63        //       华夏龙
	OutStype_haihangtiantian    int = 64        //       海航天天
	OutStype_haimengsudi    int = 65        //       海盟速递
	OutStype_huaqikuaiyun    int = 66        //       华企快运
	OutStype_shandonghaihong    int = 67        //       山东海红
	OutStype_jiajiwuliu    int = 68        //       佳吉物流
	OutStype_jiayiwuliu    int = 69        //       佳怡物流
	OutStype_jiayunmei    int = 70        //       加运美
	OutStype_jingguangsudi    int = 71        //       京广速递
	OutStype_jixianda    int = 72               //       急先达
	OutStype_jinyuekuaidi    int = 73            //       晋越快递
	OutStype_jietekuaidi    int = 74            //       捷特快递
	OutStype_jindawuliu    int = 75             //       金大物流
	OutStype_jialidatong    int = 76        //       嘉里大通
	OutStype_kangliwuliu    int = 77        //       康力物流
	OutStype_kuayuewuliu    int = 78        //       跨越物流
	OutStype_longbangwuliu    int = 79        //       龙邦物流
	OutStype_lanbiaokuaidi    int = 80        //       蓝镖快递
	OutStype_longlangkuaidi    int = 81        //       隆浪快递
	OutStype_menduimen    int = 82              //       门对门
	OutStype_mingliangwuliu    int = 83        //       明亮物流
	OutStype_quanchenkuaidi    int = 84        //       全晨快递
	OutStype_quanjitong    int = 85             //       全际通
	OutStype_quanritong    int = 86             //       全日通
	OutStype_rufengdakuaidi    int = 87        //       如风达快递
	OutStype_santaisudi    int = 88             //       三态速递
	OutStype_shenghuiwuliu    int = 89        //       盛辉物流
	OutStype_suerwuliu    int = 90              //       速尔物流
	OutStype_shengfengwuliu    int = 91        //       盛丰物流
	OutStype_shangdawuliu    int = 92                //       上大物流
	OutStype_saiaodi    int = 94                    //       赛澳递
	OutStype_shenganwuliu    int = 95                //       圣安物流
	OutStype_suijiawuliu    int = 96                //       穗佳物流
	OutStype_yousuwuliu    int = 97                 //       优速物流
	OutStype_wanjiawuliu    int = 98                //       万家物流
	OutStype_wanxiangwuliu    int = 99              //       万象物流
	OutStype_xindanaoshuowuliu    int = 100         //        新蛋奥硕物流
	OutStype_xianggangyouzheng    int = 101         //        香港邮政
	OutStype_yuntongkuaidi    int = 102             //        运通快递
	OutStype_yuanchengwuliu    int = 103             //        远成物流
	OutStype_yafengsudi    int = 104                //        亚风速递
	OutStype_yibangsudi    int = 105            //        一邦速递
	OutStype_yuanweifengkuaidi    int = 106       //        源伟丰快递
	OutStype_yuanzhijiecheng    int = 107            //        元智捷诚
	OutStype_yuefengwuliu    int = 108              //        越丰物流
	OutStype_yuananda    int = 109                  //        源安达
	OutStype_yuanfeihang    int = 110       //        原飞航
	OutStype_zhongxindakuaidi    int = 111       //        忠信达快递
	OutStype_zhimakaimen    int = 112       //        芝麻开门
	OutStype_yinjiesudi    int = 113       //        银捷速递
	OutStype_zhongyouwuliu    int = 114       //        中邮物流
	OutStype_zhongsukuaijian    int = 115       //        中速快件
	OutStype_zhongtianwanyun    int = 116       //        中天万运
	OutStype_hebeijianhua    int = 117       //        河北建华
	OutStype_lejiedi    int = 118       //        乐捷递
	OutStype_lijisong    int = 119       //        立即送
	OutStype_tonghetianxia    int = 120       //        通和天下
	OutStype_weitepai    int = 121       //        微特派
	OutStype_yitongfeihong    int = 122       //        一统飞鸿
	OutStype_zhengzhoujianhua    int = 123       //        郑州建华
	OutStype_shanxihongmajia    int = 125       //        山西红马甲
	OutStype_shanxihuangmajia    int = 126       //        陕西黄马甲
)

type OnlineConfirm struct {
	youzan.CommonHeader

	Tid         string `json:"tid,omitempty"`       //交易编号
	OuterTid    string `json:"outer_tid,omitempty"` //外部交易编号
                                                    //也可以根据外部交易编号发货，tid、outer_tid 必须选其一
	Oids        string `json:"oids,omitempty"`      //如果需要拆单发货，使用该字段指定要发货的交易明细的编号，多个明细编号用半角逗号“,”分隔；
                                                    //不需要拆单发货，则改字段不传或值为空；
	IsNoExpress int `json:"is_no_express"`          //发货是否无需物流
                                                    //如果为 0 则必须传递物流参数，
                                                    // 如果为 1 则无需传递物流参数（out_stype和out_sid），默认为 0

	OutStype    int `json:"out_stype"`              //快递公司类型（编号 =>               名称）：
	OutSid      string `json:"out_sid"`             //快递单号（具体一个物流公司的真实快递单号）
}

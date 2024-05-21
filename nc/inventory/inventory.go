package inventory

import "github.com/seerwo/yonyou/nc/base"

type Inventory struct {
}

type ReqOrderProduct struct {
	base.Base
	Bill struct {
		Billhead struct {
			Cgeneralhid   string `json:"cgeneralhid" xml:"cgeneralhid"`
			PkGroup       string `json:"pk_group" xml:"pk_group"`
			Corpoid       string `json:"corpoid" xml:"corpoid"`
			Corpvid       string `json:"corpvid" xml:"corpvid"`
			Cbiztype      string `json:"cbiztype" xml:"cbiztype"`
			Vbillcode     string `json:"vbillcode" xml:"vbillcode"`
			Dbilldate     string `json:"dbilldate" xml:"dbilldate"`
			Vtrantypecode string `json:"vtrantypecode" xml:"vtrantypecode"`
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type ReqInWare struct {
	base.Base
	Bill struct {
		Billhead struct {
			Cgeneralhid string `json:"cgeneralhid" xml:"cgeneralhid"`
			PkGroup     string `json:"pk_group" xml:"pk_group"`   //集团
			Corpoid     string `json:"corpoid" xml:"corpoid"`     //公司最新版本
			Corpvid     string `json:"corpvid" xml:"corpvid"`     //公司
			Vbillcode   string `json:"vbillcode" xml:"vbillcode"` //单据号
			Dbilldate   string `json:"dbilldate" xml:"dbilldate"` //单据日期
			PkOrg       string `json:"pk_org" xml:"pk_org"`       //销售组织
			PkOrgV      string `json:"pk_org_v" xml:"pk_org_v"`   //销售组织版本
			Cgeneralbid struct {
				Cgeneralibid string                      `json:"cgeneralibid" xml:"cgeneralibid"`
				Items        []*ReqInWareCgeneralbidItem `json:"item" xml:"item"`
			} `json:"cgeneralbid" xml:"cgeneralbid"`
			Cbiztype         string  `json:"cbiztype" xml:"cbiztype"`                 //业务流程
			Cwarehouseid     string  `json:"cwarehouseid" xml:"cwarehouseid"`         //仓库
			Vtrantypecode    string  `json:"vtrantypecode" xml:"vtrantypecode"`       //出入库类型编码
			Cothercalbodyoid string  `json:"cothercalbodyoid" xml:"cothercalbodyoid"` //出库库存组织
			Cothercalbodyvid string  `json:"cothercalbodyvid" xml:"cothercalbodyvid"` //出库库存组织版本
			Cotherwhid       string  `json:"cotherwhid" xml:"cotherwhid"`             //出库仓库
			Cwhsmanagerid    string  `json:"cwhsmanagerid" xml:"cwhsmanagerid"`       //库管员
			Cdptid           string  `json:"cdptid" xml:"cdptid"`                     //部门OID
			Cdptvid          string  `json:"cdptvid" xml:"cdptvid"`                   //部门VID
			Cbizid           string  `json:"cbizid" xml:"cbizid"`                     //业务员
			Ccustomerid      string  `json:"ccustomerid" xml:"ccustomerid"`           //客户
			PkMeasware       string  `json:"pk_measware" xml:"pk_measware"`           //计量器具
			Vnote            string  `json:"vnote" xml:"vnote"`                       //备注
			Fbillflag        string  `json:"fbillflag" xml:"fbillflag"`               //单据状态
			Iprintcount      int     `json:"iprintcount" xml:"iprintcount"`           //打印次数
			Vdef1            string  `json:"vdef1" xml:"vdef1"`
			Vdef2            string  `json:"vdef2" xml:"vdef2"`
			Vdef3            string  `json:"vdef3" xml:"vdef3"`
			Vdef4            string  `json:"vdef4" xml:"vdef4"`
			Vdef5            string  `json:"vdef5" xml:"vdef5"`
			Vdef6            string  `json:"vdef6" xml:"vdef6"`
			Vdef7            string  `json:"vdef7" xml:"vdef7"`
			Vdef8            string  `json:"vdef8" xml:"vdef8"`
			Vdef9            string  `json:"vdef9" xml:"vdef9"`
			Vdef10           string  `json:"vdef10" xml:"vdef10"`
			Vdef11           string  `json:"vdef11" xml:"vdef11"`
			Vdef12           string  `json:"vdef12" xml:"vdef12"`
			Vdef13           string  `json:"vdef13" xml:"vdef13"`
			Vdef14           string  `json:"vdef14" xml:"vdef14"`
			Vdef15           string  `json:"vdef15" xml:"vdef15"`
			Vdef16           string  `json:"vdef16" xml:"vdef16"`
			Vdef17           string  `json:"vdef17" xml:"vdef17"`
			Vdef18           string  `json:"vdef18" xml:"vdef18"`
			Vdef19           string  `json:"vdef19" xml:"vdef19"`
			Vdef20           string  `json:"vdef20" xml:"vdef20"`
			Ntotalnum        float64 `json:"ntotalnum" xml:"ntotalnum"`       //总数量
			Ntotalweight     float64 `json:"ntotalweight" xml:"ntotalweight"` //总重量
			Ntotalvolume     float64 `json:"ntotalvolume" xml:"ntotalvolume"` //总体积
			Ntotalpiece      float64 `json:"ntotalpiece" xml:"ntotalpiece"`   //总件数
			Creator          string  `json:"creator" xml:"creator"`           //创建人
			Billmaker        string  `json:"billmaker" xml:"billmaker"`       //制单人
			Creationtime     string  `json:"creationtime" xml:"creationtime"` //制单时间
			Approver         string  `json:"approver" xml:"approver"`         //签字人
			Taudittime       string  `json:"taudittime" xml:"taudittime"`     //签字时间
			Modifier         string  `json:"modifier" xml:"modifier"`         //最后修改人
			Modifiedtime     string  `json:"modifiedtime" xml:"modifiedtime"` //最后修改时间
			Ctrantypeid      string  `json:"ctrantypeid" xml:"ctrantypeid"`   //出入库类型
			//                 cpayfinorgoid           //应付财务组织最新版本
			//                 cpayfinorgvid           //应付财务组织
			//                 cfanaceorgoid           //结算财务组织最新版本
			//                 cfanaceorgvid           //结算财务组织
			//                 cpurorgoid              //采购组织
			//                 cpurorgvid              //采购组织最新版本
			//                 cvendorid               //供应商
			//                 ctrantypeid             //出入库类型
			//                 csendcountryid          //发货国
			//                 crececountryid          //收货国
			//                 ctaxcountryid           //报税国
			//                 fbuysellflag            //购销类型
			//                 btriatradeflag          //三角贸易
			//                 csendtypeid             //运输方式
			//                 ctradewordid            //贸易术语
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type ReqInWareCgeneralbidItem struct {
	Cgeneralbid string `json:"cgeneralbid" xml:"cgeneralbid"`
	Cgenerallid struct {
		Items []*ReqInWareCgeneralbidItemItem `json:"item" xml:"item"`
	} `json:"cgenerallid" xml:"cgenerallid"`
	Crowno           string  `json:"crowno" xml:"crowno"`             //行号
	PkGroup          string  `json:"pk_group" xml:"pk_group"`         //集团
	Corpoid          string  `json:"corpoid" xml:"corpoid"`           //公司最新版本
	Corpvid          string  `json:"corpvid" xml:"corpvid"`           //公司
	Cstateid         string  `json:"cstateid" xml:"cstateid"`         //库存状态
	Cmaterialoid     string  `json:"cmaterialoid" xml:"cmaterialoid"` //物料编码
	Cmaterialvid     string  `json:"cmaterialvid" xml:"cmaterialvid"` //物料版本
	Cunitid          string  `json:"cunitid" xml:"cunitid"`           //主单位
	Vfree1           string  `json:"vfree1" xml:"vfree1"`
	Vfree2           string  `json:"vfree2" xml:"vfree2"`
	Vfree3           string  `json:"vfree3" xml:"vfree3"`
	Vfree4           string  `json:"vfree4" xml:"vfree4"`
	Vfree5           string  `json:"vfree5" xml:"vfree5"`
	Vfree6           string  `json:"vfree6" xml:"vfree6"`
	Vfree7           string  `json:"vfree7" xml:"vfree7"`
	Vfree8           string  `json:"vfree8" xml:"vfree8"`
	Vfree9           string  `json:"vfree9" xml:"vfree9"`
	Vfree10          string  `json:"vfree10" xml:"vfree10"`
	Csourcetype      string  `json:"csourcetype" xml:"csourcetype"`
	Csourcetranstype string  `json:"csourcetranstype" xml:"csourcetranstype"`
	Castunitid       string  `json:"castunitid" xml:"castunitid"`             //单位
	Vchangerate      string  `json:"vchangerate" xml:"vchangerate"`           //换算率
	Vbatchcode       string  `json:"vbatchcode" xml:"vbatchcode"`             //批次号
	Nshouldnum       float64 `json:"nshouldnum" xml:"nshouldnum"`             //应收主数量
	Nshouldassistnum float64 `json:"nshouldassistnum" xml:"nshouldassistnum"` //应收数量
	Nnum             float64 `json:"nnum" xml:"nnum"`                         //实收主数量
	Nassistnum       float64 `json:"nassistnum" xml:"nassistnum"`             //实收数量
	Ncostprice       float64 `json:"ncostprice" xml:"ncostprice"`             //单价
	Ncostmny         float64 `json:"ncostmny" xml:"ncostmny"`                 //金额
	Nplannedprice    float64 `json:"nplannedprice" xml:"nplannedprice"`       //计划单价
	Nplannedmny      float64 `json:"nplannedmny" xml:"nplannedmny"`           //计划金额
	Dbizdate         string  `json:"dbizdate" xml:"dbizdate"`                 //入库日期
	Clocationid      string  `json:"clocationid" xml:"clocationid"`           //货位
	//
	//                        private int fchecked                //待检标志
	//                         ccheckstateid           //
	//                         cprojectid              //项目
	Csourcebillhid       string `json:"csourcebillhid" xml:"csourcebillhid"`             //
	Csourcebillbid       string `json:"csourcebillbid" xml:"csourcebillbid"`             //
	Ccorrespondtype      string `json:"ccorrespondtype" xml:"ccorrespondtype"`           //对应入库单类型
	Ccorrespondtranstype string `json:"ccorrespondtranstype" xml:"ccorrespondtranstype"` //对应入库单交易类型
	Vsourcebillcode      string `json:"vsourcebillcode" xml:"vsourcebillcode"`           //来源单据号
	Vsourcerowno         string `json:"vsourcerowno" xml:"vsourcerowno"`                 //来源单据行号
	Cprojectid           string `json:"cprojectid" xml:"cprojectid"`                     //项目
	Casscustid           string `json:"casscustid" xml:"casscustid"`                     //客户
	Ccorrespondbid       string `json:"ccorrespondbid" xml:"ccorrespondbid"`             //
	Ccorrespondhid       string `json:"ccorrespondhid" xml:"ccorrespondhid"`             //
	Ccorrespondcode      string `json:"ccorrespondcode" xml:"ccorrespondcode"`           //对应入库单单据号
	Ccorrespondrowno     string `json:"ccorrespondrowno" xml:"ccorrespondrowno"`         //对应入库单行号
	Cfirsttype           string `json:"cfirsttype" xml:"cfirsttype"`                     //源头单据类型
	Cfirsttranstype      string `json:"cfirsttranstype" xml:"cfirsttranstype"`           //源头单据交易类型
	Cfirstbillhid        string `json:"cfirstbillhid" xml:"cfirstbillhid"`               //
	Vfirstbillcode       string `json:"vfirstbillcode" xml:"vfirstbillcode"`             //源头单据号
	Vfirstrowno          string `json:"vfirstrowno" xml:"vfirstrowno"`                   //源头单据行号
	Cfirstbillbid        string `json:"cfirstbillbid" xml:"cfirstbillbid"`               //
	Cvendorid            string `json:"cvendorid" xml:"cvendorid"`                       //供应商
	Cffileid             string `json:"cffileid" xml:"cffileid"`                         //特征码
	Vnotebody            string `json:"vnotebody" xml:"vnotebody"`                       //行备注
	Ncountnum            string `json:"ncountnum" xml:"ncountnum"`                       //箱数
	PkPacksort           string `json:"pk_packsort" xml:"pk_packsort"`                   //包装类型
	Nbarcodenum          string `json:"nbarcodenum" xml:"nbarcodenum"`                   //条码主数量
	Bbarcodeclose        string `json:"bbarcodeclose" xml:"bbarcodeclose"`               //单据行是否条码关闭
	Bonroadflag          string `json:"bonroadflag" xml:"bonroadflag"`                   //是否在途
	Vvehiclecode         string `json:"vvehiclecode" xml:"vvehiclecode"`                 //运输工具号
	Vtransfercode        string `json:"vtransfercode" xml:"vtransfercode"`               //收货车号
	//                         flargess                //赠品
	//                         bsourcelargess          //上游赠品行
	//                         carriveorder_bbid       //
	//                         csourcetype             //来源单据类型
	//                         csourcetranstype        //来源单据交易类型
	//                         corder_bb1id            //
	//                         csrc2billtype           //其他来源单据类型编码
	//                         csrc2transtype          //其他来源交易类型编码
	//                         csrc2billhid            //
	//                         vsrc2billcode           //其他来源单据号
	//                         csrc2billbid            //
	//                         vsrc2billrowno          //其他来源单行号
	Vbdef1             string  `json:"vbdef1" xml:"vbdef1"`
	Vbdef2             string  `json:"vbdef2" xml:"vbdef2"`
	Vbdef3             string  `json:"vbdef3" xml:"vbdef3"`
	Vbdef4             string  `json:"vbdef4" xml:"vbdef4"`
	Vbdef5             string  `json:"vbdef5" xml:"vbdef5"`
	Vbdef6             string  `json:"vbdef6" xml:"vbdef6"`
	Vbdef7             string  `json:"vbdef7" xml:"vbdef7"`
	Vbdef8             string  `json:"vbdef8" xml:"vbdef8"`
	Vbdef9             string  `json:"vbdef9" xml:"vbdef9"`
	Vbdef10            string  `json:"vbdef10" xml:"vbdef10"`
	Vbdef11            string  `json:"vbdef11" xml:"vbdef11"`
	Vbdef12            string  `json:"vbdef12" xml:"vbdef12"`
	Vbdef13            string  `json:"vbdef13" xml:"vbdef13"`
	Vbdef14            string  `json:"vbdef14" xml:"vbdef14"`
	Vbdef15            string  `json:"vbdef15" xml:"vbdef15"`
	Vbdef16            string  `json:"vbdef16" xml:"vbdef16"`
	Vbdef17            string  `json:"vbdef17" xml:"vbdef17"`
	Vbdef18            string  `json:"vbdef18" xml:"vbdef18"`
	Vbdef19            string  `json:"vbdef19" xml:"vbdef19"`
	Vbdef20            string  `json:"vbdef20" xml:"vbdef20"`
	PkMeasware         string  `json:"pk_measware" xml:"pk_measware"`       //计量器具
	Ngrossnum          float64 `json:"ngrossnum" xml:"ngrossnum"`           //毛重主数量
	Ntarenum           float64 `json:"ntarenum" xml:"ntarenum"`             //皮重主数量
	Csrc2billtype      string  `json:"csrc2billtype" xml:"csrc2billtype"`   //其他来源单据类型编码
	Csrc2transtype     string  `json:"csrc2transtype" xml:"csrc2transtype"` //其他来源交易类型编码
	Csrc2billhid       string  `json:"csrc2billhid" xml:"csrc2billhid"`     //
	Vsrc2billcode      string  `json:"vsrc2billcode" xml:"vsrc2billcode"`   //其他来源单据号
	Csrc2billbid       string  `json:"csrc2billbid" xml:"csrc2billbid"`     //
	Vsrc2billrowno     string  `json:"vsrc2billrowno" xml:"vsrc2billrowno"` //其他来源单行号
	Idesatype          int     `json:"idesatype" xml:"idesatype"`           //拆解类型
	Cselastunitid      string  `json:"cselastunitid" xml:"cselastunitid"`
	Dproducedate       string  `json:"dproducedate" xml:"dproducedate"`
	Dvalidate          string  `json:"dvalidate" xml:"dvalidate"`
	Vbilltypeu8rm      string  `json:"vbilltypeu8rm" xml:"vbilltypeu8rm"`           //来自于零售之单据类型
	Vtranstypeu8rm     string  `json:"vtranstypeu8rm" xml:"vtranstypeu8rm"`         //来自于零售之交易类型
	Cbodytranstypecode string  `json:"cbodytranstypecode" xml:"cbodytranstypecode"` //出入库类型
	PkOrg              string  `json:"pk_org" xml:"pk_org"`                         //库存组织最新版本
	PkOrgv             string  `json:"pk_orgv" xml:"pk_orgv"`                       //库存组织
	Cbodywarehouseid   string  `json:"cbodywarehouseid" xml:"cbodywarehouseid"`     //库存仓库
	Cproductorid       string  `json:"cproductorid" xml:"cproductorid"`             //生产厂商
	Cvmivenderid       string  `json:"cvmivenderid" xml:"cvmivenderid"`             //寄存供应商
	Cliabilityoid      string  `json:"cliabilityoid" xml:"cliabilityoid"`           //入库利润中心
	Cliabilityvid      string  `json:"cliabilityvid" xml:"cliabilityvid"`           //入库利润中心版本
	//                         pk_creqwareid           //需求仓库
	//                         creqstoorgoid           //需求库存组织最新版本
	//                         creqstoorgvid           //需求库存组织
	//                         ctplcustomerid          //货主客户
	//                        private double nkdnum                  //扣吨主数量
	//                         vexigencycode           //紧急放行申请单号
	//                         vexigencyhid            //
	//                         vexigencyrowno          //紧急放行申请单行号
	//                         vexigencybid            //
	//                         cioliabilityoid         //收货利润中心
	//                         cioliabilityvid         //收货利润中心
	Nweight float64 `json:"nweight" xml:"nweight"` //重量
	Nvolume float64 `json:"nvolume" xml:"nvolume"` //体积
	Npiece  float64 `json:"npiece" xml:"npiece"`   //件数
	//                         borrowinflag            //借入转采购
	//                         corigcurrencyid         //币种
	//                        private double nchangestdrate          //折本汇率
	//                         ccurrencyid             //本位币
	//                        private double norignetprice           //主无税净价
	//                        private double norigtaxnetprice        //主含税净价
	//                        private double nnetprice               //主本币无税净价
	//                         ntaxnetprice            //主本币含税净价
	//                         cqtunitid               //报价单位
	//                        private double nqtunitnum              //报价数量
	//                         vqtunitrate             //报价换算率
	//                        private double nqtorignetprice         //无税净价
	//                        private double nqtorigtaxnetprice      //含税净价
	//                        private double nqtnetprice             //本币无税净价
	//                        private double nqttaxnetprice          //本币含税净价
	//                        private double norigprice              //主无税单价
	//                        private double norigtaxprice           //主含税单价
	//                        private double nprice                  //主本币无税单价
	//                        private double ntaxprice               //主本币含税单价
	//                        private double nqtorigprice            //无税单价
	//                        private double nqtorigtaxprice         //含税单价
	//                        private double nqtprice                //本币无税单价
	//                        private double nqttaxprice             //本币含税单价
	//                        private double nitemdiscountrate       //折扣
	//                        private double norigmny                //无税金额
	//                        private double norigtaxmny             //价税合计
	//                        private double norigtax                //税额
	//                        private double nmny                    //本币无税金额
	//                        private double ntaxmny                 //本币价税合计
	//                        private double ntax                    //本币税额
	//                         vreturnreason           //退库理由
	//                        private double ntaxrate                //税率
	//                         ftaxtypeflag            //扣税类别
	//                         tsourcebodyts           //来源表体时间戳
	//                        private double ncorrespondnum          //累计出库主数量
	Ncorrespondastnum float64 `json:"ncorrespondastnum" xml:"ncorrespondastnum"` //累计出库数量
	Ncorrespondgrsnum float64 `json:"ncorrespondgrsnum" xml:"ncorrespondgrsnum"` //累计出库毛重主数量
	//                        private double naccumtonum             //累计转调拨数量
	//                        private double nreplenishednum         //累计退库（补货）主数量
	//                        private double nreplenishedastnum      //累计退库（补货）数量
	//                        private double nsignnum                //累计开票主数量
	//                        private double naccumsettlenum         //累计结算数量
	//                        private double naccumvminum            //累计汇总匹配主数量
	//                         cglobalcurrencyid       //全局本位币
	//                        private double nglobalmny              //全局本币无税金额
	//                        private double nglobaltaxmny           //全局本币价税合计
	//                        private double nglobalexchgrate        //全局本位币汇率
	//                         cgroupcurrencyid        //集团本位币
	//                        private double ngroupmny               //集团本币无税金额
	//                        private double ngrouptaxmny            //集团本币价税合计
	//                        private double ngroupexchgrate         //集团本位币汇率
	//                         pk_batchcode            //批次主键
	//                         bassetcard              //已生成设备卡片
	//                         bfixedasset             //已转固
	//                         csrcmaterialoid         //来源物料编码
	//                         csrcmaterialvid         //来源物料版本
	//                         cprojecttaskid          //项目任务
	//                         csettlecurrencyid       //币种
	//                         ctaxcodeid              //税码
	//                        private double nnosubtaxrate           //不可抵扣税率
	//                        private double nnosubtax               //不可抵扣税额
	//                        private double ncaltaxmny              //计税税额
	//                         corigcountryid          //原产国
	//                         corigareaid             //原产地区
	//                         cdesticountryid         //目的国
	//                         cdestiareaid            //目的地区
	//                         bopptaxflag             //逆向征税标志
	//                        private double ncalcostmny             //计成本金额
	Ctplcustomerid string  `json:"ctplcustomerid" xml:"ctplcustomerid"` //货主客户
	Ncorrespondnum float64 `json:"ncorrespondnum" xml:"ncorrespondnum"` //累计出库主数量
	//                         ncorrespondastnum           //累计出库数量
	//                         ncorrespondgrsnum           //累计出库毛重主数量
	Nfeesettlecount int    `json:"nfeesettlecount" xml:"nfeesettlecount"` //费用结算次数
	Bhasiabill      string `json:"bhasiabill" xml:"bhasiabill"`           //是否已传存货核算
	PkBatchcode     string `json:"pk_batchcode" xml:"pk_batchcode"`       //批次主键
}

type ReqInWareCgeneralbidItemItem struct {
	Cgenerallid string  `json:"cgenerallid" xml:"cgenerallid"` //
	PkGroup     string  `json:"pk_group" xml:"pk_group"`       //集团
	Corpoid     string  `json:"corpoid" xml:"corpoid"`         //公司
	Corpvid     string  `json:"corpvid" xml:"corpvid"`         //公司版本
	Clocationid string  `json:"clocationid" xml:"clocationid"` //货位
	Vbarcode    string  `json:"vbarcode" xml:"vbarcode"`       //主条码
	Vboxbarcode string  `json:"vboxbarcode" xml:"vboxbarcode"` //箱条码
	Vbarcodesub string  `json:"vbarcodesub" xml:"vbarcodesub"` //次条码
	Vserialcode string  `json:"vserialcode" xml:"vserialcode"` //序列号
	Nnum        float64 `json:"nnum" xml:"nnum"`               //入库主数量
	Nassistnum  float64 `json:"nassistnum" xml:"nassistnum"`   //入库数量
	Ngrossnum   float64 `json:"ngrossnum" xml:"ngrossnum"`     //毛重主数量
	Cgeneralhid string  `json:"cgeneralhid" xml:"cgeneralhid"` //
}

type ReqOtherOutWare struct {
	base.Base
	Bill struct {
		Billhead struct {
			Cgeneralhid string `json:"cgeneralhid" xml:"cgeneralhid"`
			PkGroup     string `json:"pk_group" xml:"pk_group"`   //集团
			Corpoid     string `json:"corpoid" xml:"corpoid"`     //公司最新版本
			Corpvid     string `json:"corpvid" xml:"corpvid"`     //公司
			Vbillcode   string `json:"vbillcode" xml:"vbillcode"` //单据号
			Dbilldate   string `json:"dbilldate" xml:"dbilldate"` //单据日期
			PkOrg       string `json:"pk_org" xml:"pk_org"`       //销售组织
			PkOrgV      string `json:"pk_org_v" xml:"pk_org_v"`   //销售组织版本
			Cgeneralbid struct {
				Items []*ReqOtherOutItem `json:"item" xml:"item"`
			} `json:"cgeneralbid" xml:"cgeneralbid"`
			//                 cbiztype        //业务流程
			Cwarehouseid     string `json:"cwarehouseid" xml:"cwarehouseid"`         //仓库
			Vtrantypecode    string `json:"vtrantypecode" xml:"vtrantypecode"`       //出入库类型编码
			Cothercalbodyoid string `json:"cothercalbodyoid" xml:"cothercalbodyoid"` //入库库存组织
			Cothercalbodyvid string `json:"cothercalbodyvid" xml:"cothercalbodyvid"` //入库库存组织版本
			Cotherwhid       string `json:"cotherwhid" xml:"cotherwhid"`             //入库仓库
			Cwhsmanagerid    string `json:"cwhsmanagerid" xml:"cwhsmanagerid"`       //库管员
			Cdptid           string `json:"cdptid" xml:"cdptid"`                     //部门OID
			Cdptvid          string `json:"cdptvid" xml:"cdptvid"`                   //部门VID
			Cbizid           string `json:"cbizid" xml:"cbizid"`                     //业务员
			Ccustomerid      string `json:"ccustomerid" xml:"ccustomerid"`           //客户
			Vdiliveraddress  string `json:"vdiliveraddress" xml:"vdiliveraddress"`   //运输地址
			Cdilivertypeid   string `json:"cdilivertypeid" xml:"cdilivertypeid"`     //运输方式
			PkMeasware       string `json:"pk_measware" xml:"pk_measware"`           //计量器具
			Fbillflag        string `json:"fbillflag" xml:"fbillflag"`               //单据状态
			Vnote            string `json:"vnote" xml:"vnote"`                       //备注
			Iprintcount      string `json:"iprintcount" xml:"iprintcount"`           //打印次数
			Creator          string `json:"creator" xml:"creator"`                   //创建人
			Billmaker        string `json:"billmaker" xml:"billmaker"`               //制单人
			Creationtime     string `json:"creationtime" xml:"creationtime"`         //制单时间
			Approver         string `json:"approver" xml:"approver"`                 //签字人
			Taudittime       string `json:"taudittime" xml:"taudittime"`             //签字时间
			Modifier         string `json:"modifier" xml:"modifier"`                 //最后修改人
			Modifiedtime     string `json:"modifiedtime" xml:"modifiedtime"`         //最后修改时间
			Vdef1            string `json:"vdef1" xml:"vdef1"`
			Vdef2            string `json:"vdef2" xml:"vdef2"`
			Vdef3            string `json:"vdef3" xml:"vdef3"`
			Vdef4            string `json:"vdef4" xml:"vdef4"`
			Vdef5            string `json:"vdef5" xml:"vdef5"`
			Vdef6            string `json:"vdef6" xml:"vdef6"`
			Vdef7            string `json:"vdef7" xml:"vdef7"`
			Vdef8            string `json:"vdef8" xml:"vdef8"`
			Vdef9            string `json:"vdef9" xml:"vdef9"`
			Vdef10           string `json:"vdef10" xml:"vdef10"`
			Vdef11           string `json:"vdef11" xml:"vdef11"`
			Vdef12           string `json:"vdef12" xml:"vdef12"`
			Vdef13           string `json:"vdef13" xml:"vdef13"`
			Vdef14           string `json:"vdef14" xml:"vdef14"`
			Vdef15           string `json:"vdef15" xml:"vdef15"`
			Vdef16           string `json:"vdef16" xml:"vdef16"`
			Vdef17           string `json:"vdef17" xml:"vdef17"`
			Vdef18           string `json:"vdef18" xml:"vdef18"`
			Vdef19           string `json:"vdef19" xml:"vdef19"`
			Vdef20           string `json:"vdef20" xml:"vdef20"`
			Ntotalnum        string `json:"ntotalnum" xml:"ntotalnum"`       //总数量
			Ntotalweight     string `json:"ntotalweight" xml:"ntotalweight"` //总重量
			Ntotalvolume     string `json:"ntotalvolume" xml:"ntotalvolume"` //总体积
			Ntotalpiece      string `json:"ntotalpiece" xml:"ntotalpiece"`   //总件数
			Ctrantypeid      string `json:"ctrantypeid" xml:"ctrantypeid"`   //出入库类型
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type ReqOtherOutItem struct {
	Cgeneralbid string `json:"cgeneralbid" xml:"cgeneralbid"`
	Cgenerallid struct {
		Items []*CgeneralbidItem `json:"item" xml:"item"`
	} `json:"cgenerallid" xml:"cgenerallid"`
	Crowno               string  `json:"crowno" xml:"crowno"`             //行号
	PkGroup              string  `json:"pk_group" xml:"pk_group"`         //集团
	Corpoid              string  `json:"corpoid" xml:"corpoid"`           //公司最新版本
	Corpvid              string  `json:"corpvid" xml:"corpvid"`           //公司版本
	Cstateid             string  `json:"cstateid" xml:"cstateid"`         //库存状态
	Cmaterialoid         string  `json:"cmaterialoid" xml:"cmaterialoid"` //物料编码
	Cmaterialvid         string  `json:"cmaterialvid" xml:"cmaterialvid"` //物料版本
	Cunitid              string  `json:"cunitid" xml:"cunitid"`           //主单位
	Vfree1               string  `json:"vfree1" xml:"vfree1"`
	Vfree2               string  `json:"vfree2" xml:"vfree2"`
	Vfree3               string  `json:"vfree3" xml:"vfree3"`
	Vfree4               string  `json:"vfree4" xml:"vfree4"`
	Vfree5               string  `json:"vfree5" xml:"vfree5"`
	Vfree6               string  `json:"vfree6" xml:"vfree6"`
	Vfree7               string  `json:"vfree7" xml:"vfree7"`
	Vfree8               string  `json:"vfree8" xml:"vfree8"`
	Vfree9               string  `json:"vfree9" xml:"vfree9"`
	Vfree10              string  `json:"vfree10" xml:"vfree10"`
	Castunitid           string  `json:"castunitid" xml:"castunitid"`                     //单位
	Vchangerate          string  `json:"vchangerate" xml:"vchangerate"`                   //换算率
	Vbatchcode           string  `json:"vbatchcode" xml:"vbatchcode"`                     //批次号
	Nshouldnum           float64 `json:"nshouldnum" xml:"nshouldnum"`                     //应发主数量
	Nshouldassistnum     float64 `json:"nshouldassistnum" xml:"nshouldassistnum"`         //应收数量
	Nnum                 float64 `json:"nnum" xml:"nnum"`                                 //实收主数量
	Nassistnum           float64 `json:"nassistnum" xml:"nassistnum"`                     //实收数量
	Ntarenum             float64 `json:"ntarenum" xml:"ntarenum"`                         //皮重主数量
	Ngrossnum            float64 `json:"ngrossnum" xml:"ngrossnum"`                       //毛重主数量
	Vproductbatch        string  `json:"vproductbatch" xml:"vproductbatch"`               //生产订单号
	Ncostprice           float64 `json:"ncostprice" xml:"ncostprice"`                     //单价
	Ncostmny             float64 `json:"ncostmny" xml:"ncostmny"`                         //金额
	Nplannedprice        float64 `json:"nplannedprice" xml:"nplannedprice"`               //计划单价
	Nplannedmny          float64 `json:"nplannedmny" xml:"nplannedmny"`                   //计划金额
	Dbizdate             string  `json:"dbizdate" xml:"dbizdate"`                         //出库日期
	Clocationid          string  `json:"clocationid" xml:"clocationid"`                   //货位
	Csourcebillhid       string  `json:"csourcebillhid" xml:"csourcebillhid"`             //
	Csourcebillbid       string  `json:"csourcebillbid" xml:"csourcebillbid"`             //
	Csourcetype          string  `json:"csourcetype" xml:"csourcetype"`                   //来源单据类型
	Csourcetranstype     string  `json:"csourcetranstype" xml:"csourcetranstype"`         //来源单据交易类型
	Vsourcebillcode      string  `json:"vsourcebillcode" xml:"vsourcebillcode"`           //来源单据号
	Vsourcerowno         string  `json:"vsourcerowno" xml:"vsourcerowno"`                 //来源单据行号
	Cprojectid           string  `json:"cprojectid" xml:"cprojectid"`                     //项目
	Casscustid           string  `json:"casscustid" xml:"casscustid"`                     //客户
	Ccostobject          string  `json:"ccostobject" xml:"ccostobject"`                   //成本对象
	Ccorrespondbid       string  `json:"ccorrespondbid" xml:"ccorrespondbid"`             //
	Ccorrespondhid       string  `json:"ccorrespondhid" xml:"ccorrespondhid"`             //
	Ccorrespondtype      string  `json:"ccorrespondtype" xml:"ccorrespondtype"`           //对应入库单类型
	Ccorrespondtranstype string  `json:"ccorrespondtranstype" xml:"ccorrespondtranstype"` //对应入库单交易类型
	Ccorrespondcode      string  `json:"ccorrespondcode" xml:"ccorrespondcode"`           //对应入库单号
	Ccorrespondrowno     string  `json:"ccorrespondrowno" xml:"ccorrespondrowno"`         //对应入库单行号
	Cfirsttype           string  `json:"cfirsttype" xml:"cfirsttype"`                     //源头单据类型
	Cfirsttranstype      string  `json:"cfirsttranstype" xml:"cfirsttranstype"`           //源头单据交易类型
	Cfirstbillhid        string  `json:"cfirstbillhid" xml:"cfirstbillhid"`               //
	Vfirstbillcode       string  `json:"vfirstbillcode" xml:"vfirstbillcode"`             //源头单据号
	Cfirstbillbid        string  `json:"cfirstbillbid" xml:"cfirstbillbid"`               //
	Vfirstrowno          string  `json:"vfirstrowno" xml:"vfirstrowno"`                   //
	Cvendorid            string  `json:"cvendorid" xml:"cvendorid"`                       //供应商
	Cffileid             string  `json:"cffileid" xml:"cffileid"`                         //特征码
	Vnotebody            string  `json:"vnotebody" xml:"vnotebody"`                       //行备注
	Ncountnum            string  `json:"ncountnum" xml:"ncountnum"`                       //箱数
	PkPacksort           string  `json:"pk_packsort" xml:"pk_packsort"`                   //包装类型
	Nbarcodenum          string  `json:"nbarcodenum" xml:"nbarcodenum"`                   //条码数量
	Bbarcodeclose        string  `json:"bbarcodeclose" xml:"bbarcodeclose"`               //单据行是否条码关闭
	Bonroadflag          string  `json:"bonroadflag" xml:"bonroadflag"`                   //是否在途
	Vvehiclecode         string  `json:"vvehiclecode" xml:"vvehiclecode"`                 //运输工具号
	Vtransfercode        string  `json:"vtransfercode" xml:"vtransfercode"`               //收货车号
	Vbdef1               string  `json:"vbdef1" xml:"vbdef1"`
	Vbdef2               string  `json:"vbdef2" xml:"vbdef2"`
	Vbdef3               string  `json:"vbdef3" xml:"vbdef3"`
	Vbdef4               string  `json:"vbdef4" xml:"vbdef4"`
	Vbdef5               string  `json:"vbdef5" xml:"vbdef5"`
	Vbdef6               string  `json:"vbdef6" xml:"vbdef6"`
	Vbdef7               string  `json:"vbdef7" xml:"vbdef7"`
	Vbdef8               string  `json:"vbdef8" xml:"vbdef8"`
	Vbdef9               string  `json:"vbdef9" xml:"vbdef9"`
	Vbdef10              string  `json:"vbdef10" xml:"vbdef10"`
	Vbdef11              string  `json:"vbdef11" xml:"vbdef11"`
	Vbdef12              string  `json:"vbdef12" xml:"vbdef12"`
	Vbdef13              string  `json:"vbdef13" xml:"vbdef13"`
	Vbdef14              string  `json:"vbdef14" xml:"vbdef14"`
	Vbdef15              string  `json:"vbdef15" xml:"vbdef15"`
	Vbdef16              string  `json:"vbdef16" xml:"vbdef16"`
	Vbdef17              string  `json:"vbdef17" xml:"vbdef17"`
	Vbdef18              string  `json:"vbdef18" xml:"vbdef18"`
	Vbdef19              string  `json:"vbdef19" xml:"vbdef19"`
	Vbdef20              string  `json:"vbdef20" xml:"vbdef20"`
	PkMeasware           string  `json:"pk_measware" xml:"pk_measware"`       //计量器具
	Csrc2billtype        string  `json:"csrc2billtype" xml:"csrc2billtype"`   //其他来源单据类型编码
	Csrc2transtype       string  `json:"csrc2transtype" xml:"csrc2transtype"` //其他来源交易类型编码
	Csrc2billhid         string  `json:"csrc2billhid" xml:"csrc2billhid"`
	Vsrc2billcode        string  `json:"vsrc2billcode" xml:"vsrc2billcode"` //其他来源单据号
	Csrc2billbid         string  `json:"csrc2billbid" xml:"csrc2billbid"`
	Vsrc2billrowno       string  `json:"vsrc2billrowno" xml:"vsrc2billrowno"` //其他来源单行号
	Idesatype            int     `json:"idesatype" xml:"idesatype"`           //拆解类型
	Cselastunitid        string  `json:"cselastunitid" xml:"cselastunitid"`
	Dproducedate         string  `json:"dproducedate" xml:"dproducedate"`
	Dvalidate            string  `json:"dvalidate" xml:"dvalidate"`
	Vbilltypeu8rm        string  `json:"vbilltypeu8rm" xml:"vbilltypeu8rm"`           //来自于零售之单据类型
	Vtranstypeu8rm       string  `json:"vtranstypeu8rm" xml:"vtranstypeu8rm"`         //来自于零售之交易类型
	PkOrg                string  `json:"pk_org" xml:"pk_org"`                         //库存组织最新版本
	PkOrgV               string  `json:"pk_org_v" xml:"pk_org_v"`                     //库存组织
	Cbodytranstypecode   string  `json:"cbodytranstypecode" xml:"cbodytranstypecode"` //出入库类型
	Cproductorid         string  `json:"cproductorid" xml:"cproductorid"`             //生产厂商
	Cvmivenderid         string  `json:"cvmivenderid" xml:"cvmivenderid"`             //寄存供应商
	Vexigencyhid         string  `json:"vexigencyhid" xml:"vexigencyhid"`             //
	Vexigencycode        string  `json:"vexigencycode" xml:"vexigencycode"`           //紧急放行申请单号
	Vexigencyrowno       string  `json:"vexigencyrowno" xml:"vexigencyrowno"`         //紧急放行申请单行号
	Vexigencybid         string  `json:"vexigencybid" xml:"vexigencybid"`
	Cliabilityoid        string  `json:"cliabilityoid" xml:"cliabilityoid"`         //出库利润中心
	Cliabilityvid        string  `json:"cliabilityvid" xml:"cliabilityvid"`         //出库利润中心版本
	Cprojecttaskid       string  `json:"cprojecttaskid" xml:"cprojecttaskid"`       //项目任务
	Costclsidreason      string  `json:"costclsidreason" xml:"costclsidreason"`     //成本要素
	Nweight              float64 `json:"nweight" xml:"nweight"`                     //重量
	Nvolume              float64 `json:"nvolume" xml:"nvolume"`                     //体积
	Npiece               float64 `json:"npiece" xml:"npiece"`                       //件数
	Ctplcustomerid       string  `json:"ctplcustomerid" xml:"ctplcustomerid"`       //货主客户
	Csumid               string  `json:"csumid" xml:"csumid"`                       //
	Ncorrespondnum       float64 `json:"ncorrespondnum" xml:"ncorrespondnum"`       //累计出库主数量
	Ncorrespondastnum    float64 `json:"ncorrespondastnum" xml:"ncorrespondastnum"` //累计出库数量
	Ncorrespondgrsnum    float64 `json:"ncorrespondgrsnum" xml:"ncorrespondgrsnum"` //累计出库毛重主数量
	PkBatchcode          string  `json:"pk_batchcode" xml:"pk_batchcode"`           //批次主键
}

type CgeneralbidItem struct {
	Cgenerallid string  `json:"cgenerallid" xml:"cgenerallid"` //
	PkGroup     string  `json:"pk_group" xml:"pk_group"`       //集团
	Corpoid     string  `json:"corpoid" xml:"corpoid"`         //公司
	Corpvid     string  `json:"corpvid" xml:"corpvid"`         //公司版本
	Clocationid string  `json:"clocationid" xml:"clocationid"` //货位
	Vbarcode    string  `json:"vbarcode" xml:"vbarcode"`       //主条码
	Vboxbarcode string  `json:"vboxbarcode" xml:"vboxbarcode"` //箱条码
	Vbarcodesub string  `json:"vbarcodesub" xml:"vbarcodesub"` //次条码
	Vserialcode string  `json:"vserialcode" xml:"vserialcode"` //序列号
	Nnum        float64 `json:"nnum" xml:"nnum"`               //入库主数量
	Nassistnum  float64 `json:"nassistnum" xml:"nassistnum"`   //入库数量
	Ngrossnum   string  `json:"ngrossnum" xml:"ngrossnum"`     //毛重主数量
	Cgeneralhid string  `json:"cgeneralhid" xml:"cgeneralhid"` //
}

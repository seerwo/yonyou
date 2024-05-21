package purchase

import "github.com/seerwo/yonyou/nc/base"

//Purchase struct
type Purchase struct {
	//*context.Context
}

type ReqPurchaseOrder struct {
	base.Base
	Bill struct {
		BillHead struct {
			PkOrder           string  `json:"pk_order" xml:"pk_order"`
			Ctradewordid      string  `json:"ctradewordid" xml:"ctradewordid"`           //贸易术语
			PkGroup           string  `json:"pk_group" xml:"pk_group"`                   //集团
			PkOrg             string  `json:"pk_org" xml:"pk_org"`                       //销售组织
			PkOrgV            string  `json:"pk_org_v" xml:"pk_org_v"`                   //销售组织版本
			Vbillcode         string  `json:"vbillcode" xml:"vbillcode"`                 //单据号
			Dbilldate         string  `json:"dbilldate" xml:"dbilldate"`                 //单据日期
			PkFreecust        string  `json:"pk_freecust" xml:"pk_freecust"`             //散户
			PkSupplier        string  `json:"pk_supplier" xml:"pk_supplier"`             //供应商
			PkBankdoc         string  `json:"pk_bankdoc" xml:"pk_bankdoc"`               //开户银行
			PkDept            string  `json:"pk_dept" xml:"pk_dept"`                     //采购部门最新版本
			Cemployeeid       string  `json:"cemployeeid" xml:"cemployeeid"`             //采购员
			Vtrantypecode     string  `json:"vtrantypecode" xml:"vtrantypecode"`         //订单类型编码
			PkRecvcustomer    string  `json:"pk_recvcustomer" xml:"pk_recvcustomer"`     //收货客户
			PkInvcsupllier    string  `json:"pk_invcsupllier" xml:"pk_invcsupllier"`     //开票供应商
			PkDeliveradd      string  `json:"pk_deliveradd" xml:"pk_deliveradd"`         //供应商发货地址
			PkTransporttype   string  `json:"pk_transporttype" xml:"pk_transporttype"`   //运输方式
			PkPayterm         string  `json:"pk_payterm" xml:"pk_payterm"`               //付款协议
			Billmaker         string  `json:"billmaker" xml:"billmaker"`                 //制单人
			Forderstatus      string  `json:"forderstatus" xml:"forderstatus"`           //单据状态
			Approver          string  `json:"approver" xml:"approver"`                   //审批人
			Vmemo             string  `json:"vmemo" xml:"vmemo"`                         //备注
			Ccontracttextpath string  `json:"ccontracttextpath" xml:"ccontracttextpath"` //合同附件
			Norgprepaylimit   float64 `json:"norgprepaylimit" xml:"norgprepaylimit"`     //预付款限额
			Nversion          int     `json:"nversion" xml:"nversion"`                   //版本号
			Bislatest         string  `json:"bislatest" xml:"bislatest"`                 //最新版本
			Bisreplenish      string  `json:"bisreplenish" xml:"bisreplenish"`           //补货
			Breturn           string  `json:"breturn" xml:"breturn"`                     //退货
			Iprintcount       string  `json:"iprintcount" xml:"iprintcount"`             //打印次数
			Vnote             string  `json:"vnote" xml:"vnote"`
			Vdef1             string  `json:"vdef1" xml:"vdef1"`
			Vdef2             string  `json:"vdef2" xml:"vdef2"`
			Vdef3             string  `json:"vdef3" xml:"vdef3"`
			Vdef4             string  `json:"vdef4" xml:"vdef4"`
			Vdef5             string  `json:"vdef5" xml:"vdef5"`
			Vdef6             string  `json:"vdef6" xml:"vdef6"`
			Vdef7             string  `json:"vdef7" xml:"vdef7"`
			Vdef8             string  `json:"vdef8" xml:"vdef8"`
			Vdef9             string  `json:"vdef9" xml:"vdef9"`
			Vdef10            string  `json:"vdef10" xml:"vdef10"`
			Vdef11            string  `json:"vdef11" xml:"vdef11"`
			Vdef12            string  `json:"vdef12" xml:"vdef12"`
			Vdef13            string  `json:"vdef13" xml:"vdef13"`
			Vdef14            string  `json:"vdef14" xml:"vdef14"`
			Vdef15            string  `json:"vdef15" xml:"vdef15"`
			Vdef16            string  `json:"vdef16" xml:"vdef16"`
			Vdef17            string  `json:"vdef17" xml:"vdef17"`
			Vdef18            string  `json:"vdef18" xml:"vdef18"`
			Vdef19            string  `json:"vdef19" xml:"vdef19"`
			Vdef20            string  `json:"vdef20" xml:"vdef20"`
			Creationtime      string  `json:"creationtime" xml:"creationtime"`       //制单时间
			Taudittime        string  `json:"taudittime" xml:"taudittime"`           //审批时间
			Modifiedtime      string  `json:"modifiedtime" xml:"modifiedtime"`       //最后修改时间
			Crevisepsn        string  `json:"crevisepsn" xml:"crevisepsn"`           //修订人
			Trevisiontime     string  `json:"trevisiontime" xml:"trevisiontime"`     //修订时间
			Bcooptoso         string  `json:"bcooptoso" xml:"bcooptoso"`             //已协同生成销售订单
			Bsocooptome       string  `json:"bsocooptome" xml:"bsocooptome"`         //由销售订单协同生成
			Vcoopordercode    string  `json:"vcoopordercode" xml:"vcoopordercode"`   //对方订单号
			Ntotalastnum      float64 `json:"ntotalastnum" xml:"ntotalastnum"`       //总数量
			Ntotalorigmny     float64 `json:"ntotalorigmny" xml:"ntotalorigmny"`     //价税合计
			Modifier          string  `json:"modifier" xml:"modifier"`               //最后修改人
			PkFreezepsndoc    string  `json:"pk_freezepsndoc" xml:"pk_freezepsndoc"` //冻结人
			Tfreezetime       string  `json:"tfreezetime" xml:"tfreezetime"`         //冻结时间
			Bfrozen           string  `json:"bfrozen" xml:"bfrozen"`                 //冻结
			Vfrozenreason     string  `json:"vfrozenreason" xml:"vfrozenreason"`     //冻结原因
			PkBusitype        string  `json:"pk_busitype" xml:"pk_busitype"`         //业务流程
			Fhtaxtypeflag     int     `json:"fhtaxtypeflag" xml:"fhtaxtypeflag"`     //整单扣税类别
			Nhtaxrate         float64 `json:"nhtaxrate" xml:"nhtaxrate"`             //整单税率
			Corigcurrencyid   string  `json:"corigcurrencyid" xml:"corigcurrencyid"` //币种
			Brefwhenreturn    string  `json:"brefwhenreturn" xml:"brefwhenreturn"`   //退货/库基于原订单补货
			Ntotalweight      float64 `json:"ntotalweight" xml:"ntotalweight"`       //总重量
			Ntotalvolume      float64 `json:"ntotalvolume" xml:"ntotalvolume"`       //总体积
			Ntotalpiece       float64 `json:"ntotalpiece" xml:"ntotalpiece"`         //总件数
			Bfinalclose       string  `json:"bfinalclose" xml:"bfinalclose"`         //最终关闭
			Creator           string  `json:"creator" xml:"creator"`                 //创建人
			PkProject         string  `json:"pk_project" xml:"pk_project"`           //项目
			Dclosedate        string  `json:"dclosedate" xml:"dclosedate"`           //最终关闭日期
			PkDeptV           string  `json:"pk_dept_v" xml:"pk_dept_v"`             //采购部门
			Ctrantypeid       string  `json:"ctrantypeid" xml:"ctrantypeid"`         //订单类型
			Irespstatus       string  `json:"irespstatus" xml:"irespstatus"`         //响应状态
			Vreason           string  `json:"vreason" xml:"vreason"`                 //拒绝/变更理由
			Bpublish          string  `json:"bpublish" xml:"bpublish"`               //发布
			PkPubpsn          string  `json:"pk_pubpsn" xml:"pk_pubpsn"`             //发布人
			Tpubtime          string  `json:"tpubtime" xml:"tpubtime"`               //发布时间
			PkResppsn         string  `json:"pk_resppsn" xml:"pk_resppsn"`           //响应人
			Tresptime         string  `json:"tresptime" xml:"tresptime"`             //响应时间
			PkOrderB          struct {
				Items []PurchaseItem `json:"item" xml:"item"`
			} `json:"pk_order_b" xml:"pk_order_b"`
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type PurchaseItem struct {
	PkOrderB           string  `json:"pk_order_b" xml:"pk_order_b"`
	PkGroup            string  `json:"pk_group" xml:"pk_group"`                   //所属集团
	PkOrg              string  `json:"pk_org" xml:"pk_org"`                       //采购组织版本信息
	PkOrgV             string  `json:"pk_org_v" xml:"pk_org_v"`                   //采购组织
	PkReqcorp          string  `json:"pk_reqcorp"`                                //需求公司
	PkReqstoorg        string  `json:"pk_reqstoorg" xml:"pk_reqstoorg"`           //需求库存组织最新版本
	PkReqstoorgV       string  `json:"pk_reqstoorg_v" xml:"pk_reqstoorg_v"`       //需求库存组织
	PkReqstordoc       string  `json:"pk_reqstordoc" xml:"pk_reqstordoc"`         //需求仓库
	PkArrvstoorg       string  `json:"pk_arrvstoorg" xml:"pk_arrvstoorg"`         //收货库存组织最新版本
	PkArrvstoorgV      string  `json:"pk_arrvstoorg_v" xml:"pk_arrvstoorg_v"`     //收货库存组织
	PkReqdept          string  `json:"pk_reqdept" xml:"pk_reqdept"`               //需求部门最新版本
	PkFlowstockorg     string  `json:"pk_flowstockorg" xml:"pk_flowstockorg"`     //物流组织最新版本
	PkFlowstockorgV    string  `json:"pk_flowstockorg_v" xml:"pk_flowstockorg_v"` //物流组织
	Crowno             string  `json:"crowno" xml:"crowno"`                       //行号
	PkMaterial         string  `json:"pk_material" xml:"pk_material"`             //物料版本信息
	PkSrcmaterial      string  `json:"pk_srcmaterial" xml:"pk_srcmaterial"`       //物料信息
	Cunitid            string  `json:"cunitid" xml:"cunitid"`                     //主单位
	Nnum               float64 `json:"nnum" xml:"nnum"`                           //主数量
	Castunitid         string  `json:"castunitid" xml:"castunitid"`               //单位
	Nastnum            string  `json:"nastnum" xml:"nastnum"`                     //数量
	Vchangerate        string  `json:"vchangerate" xml:"vchangerate"`             //换算率
	Cqtunitid          string  `json:"cqtunitid" xml:"cqtunitid"`                 //报价单位
	Nqtunitnum         string  `json:"nqtunitnum" xml:"nqtunitnum"`               //报价数量
	Vqtunitrate        string  `json:"vqtunitrate" xml:"vqtunitrate"`             //报价换算率
	Nqtorigprice       float64 `json:"nqtorigprice" xml:"nqtorigprice"`           //无税单价
	Nqtorigtaxprice    float64 `json:"nqtorigtaxprice" xml:"nqtorigtaxprice"`     //含税单价
	Nqtorignetprice    float64 `json:"nqtorignetprice" xml:"nqtorignetprice"`     //无税净价
	Nqtorigtaxnetprc   float64 `json:"nqtorigtaxnetprc" xml:"nqtorigtaxnetprc"`   //含税净价
	Nqtnetprice        float64 `json:"nqtnetprice" xml:"nqtnetprice"`             //本币无税净价
	Nqttaxnetprice     float64 `json:"nqttaxnetprice" xml:"nqttaxnetprice"`       //本币含税净价
	Nitemdiscountrate  float64 `json:"nitemdiscountrate" xml:"nitemdiscountrate"` //折扣
	Norigmny           float64 `json:"norigmny" xml:"norigmny"`                   //无税金额
	Norigtaxmny        float64 `json:"norigtaxmny" xml:"norigtaxmny"`             //价税合计
	Nmny               float64 `json:"nmny" xml:"nmny"`                           //本币无税金额
	Ntaxmny            float64 `json:"ntaxmny" xml:"ntaxmny"`                     //本币价税合计
	Ntax               float64 `json:"ntax" xml:"ntax"`                           //本币税额
	Norigprice         float64 `json:"norigprice" xml:"norigprice"`               //主无税单价
	Norigtaxprice      float64 `json:"norigtaxprice" xml:"norigtaxprice"`         //主含税单价
	Norignetprice      float64 `json:"norignetprice" xml:"norignetprice"`         //主无税净价
	Norigtaxnetprice   float64 `json:"norigtaxnetprice" xml:"norigtaxnetprice"`   //主含税净价
	Nnetprice          float64 `json:"nnetprice" xml:"nnetprice"`                 //主本币无税净价
	Ntaxnetprice       float64 `json:"ntaxnetprice" xml:"ntaxnetprice"`           //主本币含税净价
	Naccumarrvnum      float64 `json:"naccumarrvnum" xml:"naccumarrvnum"`         //累计到货主数量
	Naccumstorenum     float64 `json:"naccumstorenum" xml:"naccumstorenum"`       //累计入库主数量
	Naccuminvoicenum   float64 `json:"naccuminvoicenum" xml:"naccuminvoicenum"`   //累计开票主数量
	Naccumwastnum      float64 `json:"naccumwastnum" xml:"naccumwastnum"`         //累计途耗主数量
	Dplanarrvdate      string  `json:"dplanarrvdate" xml:"dplanarrvdate"`         //
	PkRecvstordoc      string  `json:"pk_recvstordoc" xml:"pk_recvstordoc"`       //收货仓库
	PkReceiveaddress   string  `json:"pk_receiveaddress" xml:"pk_receiveaddress"` //收货地址
	Dcorrectdate       string  `json:"dcorrectdate" xml:"dcorrectdate"`           //修正日期
	Chandler           string  `json:"chandler" xml:"chandler"`                   //操作员
	Fisactive          int     `json:"fisactive" xml:"fisactive"`                 //激活
	Csourcetypecode    string  `json:"csourcetypecode" xml:"csourcetypecode"`     //来源单据类型
	Csourceid          string  `json:"csourceid" xml:"csourceid"`                 //
	Csourcebid         string  `json:"csourcebid" xml:"csourcebid"`
	Cfirsttypecode     string  `json:"cfirsttypecode" xml:"cfirsttypecode"` //源头单据类型
	Cfirstid           string  `json:"cfirstid" xml:"cfirstid"`
	Cfirstbid          string  `json:"cfirstbid" xml:"cfirstbid"`
	Vbmemo             string  `json:"vbmemo" xml:"vbmemo"`                   //备注
	PkBatchcode        string  `json:"pk_batchcode" xml:"pk_batchcode"`       //
	Vbatchcode         string  `json:"vbatchcode" xml:"vbatchcode"`           //批次号
	Nbackarrvnum       string  `json:"nbackarrvnum" xml:"nbackarrvnum"`       //累计退货主数量
	Nbackstorenum      int     `json:"nbackstorenum" xml:"nbackstorenum"`     //累计退库主数量
	Cqpbaseschemeid    string  `json:"cqpbaseschemeid" xml:"cqpbaseschemeid"` //优质优价方案
	Ccontractrowid     string  `json:"ccontractrowid" xml:"ccontractrowid"`
	Ccontractid        string  `json:"ccontractid" xml:"ccontractid"`             //合同信息
	Vcontractcode      string  `json:"vcontractcode" xml:"vcontractcode"`         //合同号
	Vpriceauditcode    string  `json:"vpriceauditcode" xml:"vpriceauditcode"`     //价格审批单号
	Cpriceauditid      string  `json:"cpriceauditid" xml:"cpriceauditid"`         //价格审批单
	CpriceauditBid     string  `json:"cpriceaudit_bid" xml:"cpriceaudit_bid"`     //
	CpriceauditBb1id   string  `json:"cpriceaudit_bb1id" xml:"cpriceaudit_bb1id"` //
	Breceiveplan       string  `json:"breceiveplan" xml:"breceiveplan"`           //存在到货计划
	Naccumrpnum        float64 `json:"naccumrpnum" xml:"naccumrpnum"`             //累计到货计划主数量
	Blargess           string  `json:"blargess" xml:"blargess"`                   //是否赠品
	Cvenddevareaid     string  `json:"cvenddevareaid" xml:"cvenddevareaid"`       //供应商发货地区
	Cvenddevaddrid     string  `json:"cvenddevaddrid" xml:"cvenddevaddrid"`       //供应商发货地点
	Vvenddevaddr       string  `json:"vvenddevaddr" xml:"vvenddevaddr"`           //供应商发货地址
	Cdevareaid         string  `json:"cdevareaid" xml:"cdevareaid"`               //收货地区
	Cdevaddrid         string  `json:"cdevaddrid" xml:"cdevaddrid"`               //收货地点
	Naccumdevnum       float64 `json:"naccumdevnum" xml:"naccumdevnum"`           //累计运输主数量
	Ncaninnum          float64 `json:"ncaninnum" xml:"ncaninnum"`
	Vfree1             string  `json:"vfree1" xml:"vfree1"`
	Vfree2             string  `json:"vfree2" xml:"vfree2"`
	Vfree3             string  `json:"vfree3" xml:"vfree3"`
	Vfree4             string  `json:"vfree4" xml:"vfree4"`
	Vfree5             string  `json:"vfree5" xml:"vfree5"`
	Vfree6             string  `json:"vfree6" xml:"vfree6"`
	Vfree7             string  `json:"vfree7" xml:"vfree7"`
	Vfree8             string  `json:"vfree8" xml:"vfree8"`
	Vfree9             string  `json:"vfree9" xml:"vfree9"`
	Vfree10            string  `json:"vfree10" xml:"vfree10"`
	Vfree11            string  `json:"vfree11" xml:"vfree11"`
	Vfree12            string  `json:"vfree12" xml:"vfree12"`
	Vfree13            string  `json:"vfree13" xml:"vfree13"`
	Vfree14            string  `json:"vfree14" xml:"vfree14"`
	Vfree15            string  `json:"vfree15" xml:"vfree15"`
	Vfree16            string  `json:"vfree16" xml:"vfree16"`
	Vfree17            string  `json:"vfree17" xml:"vfree17"`
	Vfree18            string  `json:"vfree18" xml:"vfree18"`
	Vfree19            string  `json:"vfree19" xml:"vfree19"`
	Vfree20            string  `json:"vfree20" xml:"vfree20"`
	Vbdef1             string  `json:"vbdef1" xml:"vbdef1"`
	Vbdef2             string  `json:"vbdef2" xml:"vbdef2"`
	Vbdef3             string  `json:"vbdef3" xml:"vbdef3"`
	Vbdef4             string  `json:"vbdef4" xml:"vbdef4"`
	Vbdef5             string  `json:"vbdef5" xml:"vbdef5"`
	Vvendinventorycode string  `json:"vvendinventorycode" xml:"vvendinventorycode"` //对应物料编码
	Vvendinventoryname string  `json:"vvendinventoryname" xml:"vvendinventoryname"` //对应物料名称
	Btransclosed       string  `json:"btransclosed" xml:"btransclosed"`             //是否运输关闭
	Nfeemny            float64 `json:"nfeemny" xml:"nfeemny"`                       //费用累计开票金额
	PkPsfinanceorg     string  `json:"pk_psfinanceorg" xml:"pk_psfinanceorg"`       //结算财务组织最新版本
	PkPsfinanceorgV    string  `json:"pk_psfinanceorg_v" xml:"pk_psfinanceorg_v"`   //结算财务组织
	PkApfinanceorg     string  `json:"pk_apfinanceorg" xml:"pk_apfinanceorg"`       //应付组织最新版本
	PkApfinanceorgV    string  `json:"pk_apfinanceorg_v" xml:"pk_apfinanceorg_v"`   //应付组织
	PkApliabcenter     string  `json:"pk_apliabcenter" xml:"pk_apliabcenter"`       //利润中心最新版本
	PkApliabcenterV    string  `json:"pk_apliabcenter_v"`                           //利润中心
	Nacccancelpaymny   float64 `json:"nacccancelpaymny" xml:"nacccancelpaymny"`     //累计已核销本币付款金额
	Bborrowpur         string  `json:"bborrowpur" xml:"bborrowpur"`                 //借入转采购
	Vfirsttrantype     string  `json:"vfirsttrantype" xml:"vfirsttrantype"`         //源头交易类型
	Vfirstcode         string  `json:"vfirstcode" xml:"vfirstcode"`                 //源头单据号
	Vfirstrowno        string  `json:"vfirstrowno" xml:"vfirstrowno"`               //源头单据行号
	Vsourcetrantype    string  `json:"vsourcetrantype" xml:"vsourcetrantype"`       //来源交易类型
	Vsourcecode        string  `json:"vsourcecode" xml:"vsourcecode"`               //来源单据号
	Vsourcerowno       string  `json:"vsourcerowno" xml:"vsourcerowno"`             //来源单据行号
	Naccuminvoicemny   float64 `json:"naccuminvoicemny" xml:"naccuminvoicemny"`     //累计本币开票金额
	Nacccancelinvmny   float64 `json:"nacccancelinvmny" xml:"nacccancelinvmny"`     //累计已核销本币开票金额
	Nweight            float64 `json:"nweight" xml:"nweight"`                       //重量
	Nvolumn            float64 `json:"nvolumn" xml:"nvolumn"`                       //体积
	Npacknum           float64 `json:"npacknum" xml:"npacknum"`                     //件数
	Bstockclose        string  `json:"bstockclose" xml:"bstockclose"`               //入库关闭
	Binvoiceclose      string  `json:"binvoiceclose" xml:"binvoiceclose"`           //开票关闭
	Barriveclose       string  `json:"barriveclose" xml:"barriveclose"`             //到货关闭
	Bpayclose          string  `json:"bpayclose" xml:"bpayclose"`                   //付款关闭
	Ftaxtypeflag       int     `json:"ftaxtypeflag" xml:"ftaxtypeflag"`             //扣税类别
	Ntaxrate           float64 `json:"ntaxrate" xml:"ntaxrate"`                     //税率
	Ccurrencyid        string  `json:"ccurrencyid" xml:"ccurrencyid"`               //本币币种
	Nexchangerate      float64 `json:"nexchangerate" xml:"nexchangerate"`           //折本汇率
	Ngroupexchgrate    float64 `json:"ngroupexchgrate" xml:"ngroupexchgrate"`       //集团本位币汇率
	Nglobalexchgrate   float64 `json:"nglobalexchgrate" xml:"nglobalexchgrate"`     //全局本位币汇率
	Ngroupmny          float64 `json:"ngroupmny" xml:"ngroupmny"`                   //集团本币无税金额
	Ngrouptaxmny       float64 `json:"ngrouptaxmny" xml:"ngrouptaxmny"`             //集团本币价税合计
	Nglobalmny         float64 `json:"nglobalmny" xml:"nglobalmny"`                 //全局本币无税金额
	Nglobaltaxmny      float64 `json:"nglobaltaxmny" xml:"nglobaltaxmny"`           //全局本币价税合计
	Cprojectid         string  `json:"cprojectid" xml:"cprojectid"`                 //项目
	Cproductorid       string  `json:"cproductorid" xml:"cproductorid"`             //生产厂商
	Dbilldate          string  `json:"dbilldate" xml:"dbilldate"`                   //订单日期
	PkSupplier         string  `json:"pk_supplier" xml:"pk_supplier"`               //供应商
	Corigcurrencyid    string  `json:"corigcurrencyid" xml:"corigcurrencyid"`       //币种
	Cqualitylevelid    string  `json:"cqualitylevelid" xml:"cqualitylevelid"`       //质量等级
	Nsuprsnum          float64 `json:"nsuprsnum" xml:"nsuprsnum"`                   //被预留数量
	PkSrcorderB        string  `json:"pk_srcorder_b" xml:"pk_srcorder_b"`           //
	PkReqdeptV         string  `json:"pk_reqdept_v" xml:"pk_reqdept_v"`             //需求部门
	Nqtprice           float64 `json:"nqtprice" xml:"nqtprice"`                     //报价本币无税单价
	Nqttaxprice        float64 `json:"nqttaxprice" xml:"nqttaxprice"`               //报价本币含税单价
	Nprice             float64 `json:"nprice" xml:"nprice"`                         //主本币无税单价
	Ntaxprice          float64 `json:"ntaxprice" xml:"ntaxprice"`                   //主本币含税单价
	Cecbillid          string  `json:"cecbillid" xml:"cecbillid"`                   //
	Cecbillbid         string  `json:"cecbillbid" xml:"cecbillbid"`                 //
	Vecbillcode        string  `json:"vecbillcode" xml:"vecbillcode"`               //电子商务单据号
	Cectypecode        string  `json:"cectypecode" xml:"cectypecode"`               //电子商务单据类型
	Norigordernum      float64 `json:"norigordernum" xml:"norigordernum"`           //原始订单数量
	Norigorderprice    float64 `json:"norigorderprice" xml:"norigorderprice"`       //原始订单净含税单价
	Dorigplanarrvdate  string  `json:"dorigplanarrvdate" xml:"dorigplanarrvdate"`   //原始计划到货日期
	Istorestatus       string  `json:"istorestatus" xml:"istorestatus"`             //供应商交货状态
	Casscustid         string  `json:"casscustid" xml:"casscustid"`                 //客户
	Cprojecttaskid     string  `json:"cprojecttaskid" xml:"cprojecttaskid"`         //项目任务
	Csendcountryid     string  `json:"csendcountryid" xml:"csendcountryid"`         //发货国家/地区
	Crececountryid     string  `json:"crececountryid" xml:"crececountryid"`         //收货国家/地区
	Ctaxcountryid      string  `json:"ctaxcountryid" xml:"ctaxcountryid"`           //报税国家/地区
	Fbuysellflag       int     `json:"fbuysellflag" xml:"fbuysellflag"`             //购销类型
	Btriatradeflag     string  `json:"btriatradeflag" xml:"btriatradeflag"`         //三角贸易
	Ctaxcodeid         string  `json:"ctaxcodeid" xml:"ctaxcodeid"`                 //税码
	Nnosubtaxrate      float64 `json:"nnosubtaxrate" xml:"nnosubtaxrate"`           //不可抵扣税率
	Nnosubtax          float64 `json:"nnosubtax" xml:"nnosubtax"`                   //不可抵扣税额
	Ncaltaxmny         float64 `json:"ncaltaxmny" xml:"ncaltaxmny"`                 //计税金额
	Ncalcostmny        float64 `json:"ncalcostmny" xml:"ncalcostmny"`               //计成本金额
	Corigcountryid     string  `json:"corigcountryid" xml:"corigcountryid"`         //原产国
	Corigareaid        string  `json:"corigareaid" xml:"corigareaid"`               //原产地区
	Cdesticountryid    string  `json:"cdesticountryid" xml:"cdesticountryid"`       //目的地国
	Cdestiareaid       string  `json:"cdestiareaid" xml:"cdestiareaid"`             //目的地地区
	Nsourcenum         float64 `json:"nsourcenum" xml:"nsourcenum"`                 //来源单据主数量
}

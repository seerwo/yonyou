package sale

import (
	"encoding/xml"
	"fmt"
	"github.com/seerwo/yonyou/nc/base"
	"github.com/seerwo/yonyou/nc/context"
	"github.com/seerwo/yonyou/util"
)

//Sale struct
type Sale struct {
	*context.Context
}

//NewSale instance
func NewSale(context *context.Context) *Sale {
	sale := new(Sale)
	sale.Context = context
	return sale
}

type ReqSaleOrder struct {
	XMLName xml.Name `xml:"ufinterface"`
	base.Base
	Bill struct {
		Billhead struct {
			Csaleorderid     string  `json:"csaleorderid" xml:"csaleorderid"`
			PkGroup          string  `json:"pk_group" xml:"pk_group"`                 //集团
			PkOrg            string  `json:"pk_org" xml:"pk_org"`                     //销售组织
			PkOrgV           string  `json:"pk_org_v" xml:"pk_org_v"`                 //销售组织版本
			Ctrantypeid      string  `json:"ctrantypeid" xml:"ctrantypeid"`           //订单类型
			Vtrantypecode    string  `json:"vtrantypecode" xml:"vtrantypecode"`       //订单类型编码
			Cbiztypeid       string  `json:"cbiztypeid" xml:"cbiztypeid"`             //业务流程
			Vbillcode        string  `json:"vbillcode" xml:"vbillcode"`               //单据号
			Dbilldate        string  `json:"dbilldate" xml:"dbilldate"`               //单据日期
			Ccustomerid      string  `json:"ccustomerid" xml:"ccustomerid"`           //客户
			Bfreecustflag    string  `json:"bfreecustflag" xml:"bfreecustflag"`       //是否散户
			Cfreecustid      string  `json:"cfreecustid" xml:"cfreecustid"`           //散户
			Cdeptvid         string  `json:"cdeptvid" xml:"cdeptvid"`                 //部门
			Cdeptid          string  `json:"cdeptid" xml:"cdeptid"`                   //部门
			Cemployeeid      string  `json:"cemployeeid" xml:"cemployeeid"`           //业务员
			Corigcurrencyid  string  `json:"corigcurrencyid" xml:"corigcurrencyid"`   //原币
			Cinvoicecustid   string  `json:"cinvoicecustid" xml:"cinvoicecustid"`     //开票客户
			Ccustbankid      string  `json:"ccustbankid" xml:"ccustbankid"`           //开户银行
			Ccustbankaccid   string  `json:"ccustbankaccid" xml:"ccustbankaccid"`     //开户银行账户
			Cchanneltypeid   string  `json:"cchanneltypeid" xml:"cchanneltypeid"`     //销售渠道类型
			Cpaytermid       string  `json:"cpaytermid" xml:"cpaytermid"`             //收款协议
			Ctransporttypeid string  `json:"ctransporttypeid" xml:"ctransporttypeid"` //运输方式
			Ctradewordid     string  `json:"ctradewordid" xml:"ctradewordid"`         //贸易术语
			Ndiscountrate    float64 `json:"ndiscountrate" xml:"ndiscountrate"`       //整单折扣
			Vcreditnum       string  `json:"vcreditnum" xml:"vcreditnum"`             //信用证号
			Cbalancetypeid   string  `json:"cbalancetypeid"`                          //结算方式
			Badvfeeflag      string  `json:"badvfeeflag" xml:"badvfeeflag"`           //代垫运费
			Npreceiverate    float64 `json:"npreceiverate" xml:"npreceiverate"`       //订单收款比例
			Npreceivequota   float64 `json:"npreceivequota" xml:"npreceivequota"`     //订单收款限额
			Bpreceiveflag    string  `json:"bpreceiveflag" xml:"bpreceiveflag"`       //收款限额控制预收
			Npreceivemny     float64 `json:"npreceivemny" xml:"npreceivemny"`         //实际预收款金额
			Nreceivedmny     float64 `json:"nreceivedmny" xml:"nreceivedmny"`         //实际收款金额
			Nthisreceivemny  float64 `json:"nthisreceivemny" xml:"nthisreceivemny"`   //本次收款金额
			Ntotalnum        float64 `json:"ntotalnum" xml:"ntotalnum"`               //总数量
			Ntotalweight     float64 `json:"ntotalweight" xml:"ntotalweight"`         //合计重量
			Ntotalvolume     float64 `json:"ntotalvolume" xml:"ntotalvolume"`         //合计体积
			Ntotalpiece      float64 `json:"ntotalpiece" xml:"ntotalpiece"`           //总件数
			Ntotalorigmny    float64 `json:"ntotalorigmny" xml:"ntotalorigmny"`       //价税合计
			Ntotalmny        float64 `json:"ntotalmny" xml:"ntotalmny"`               //冲抵前金额
			Ntotalorigsubmny float64 `json:"ntotalorigsubmny" xml:"ntotalorigsubmny"` //费用冲抵金额
			Boffsetflag      string  `json:"boffsetflag" xml:"boffsetflag"`           //是否冲抵
			Vcooppohcode     string  `json:"vcooppohcode" xml:"vcooppohcode"`         //对方订单号
			Bpocooptomeflag  string  `json:"bpocooptomeflag" xml:"bpocooptomeflag"`   //由采购订单协同生成
			DestPkOrg        string  `json:"dest_pk_org" xml:"dest_pk_org"`
			Bcooptopoflag    string  `json:"bcooptopoflag" xml:"bcooptopoflag"` //已协同生成采购订单
			Fstatusflag      int     `json:"fstatusflag" xml:"fstatusflag"`     //单据状态
			Fpfstatusflag    string  `json:"fpfstatusflag" xml:"fpfstatusflag"` //审批流状态
			Vnote            string  `json:"vnote" xml:"vnote"`
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
			Vdef14           string  `json:"vdef14" xml:"vdef14"` //销售分类 b2b,b2c
			Vdef15           string  `json:"vdef15" xml:"vdef15"`
			Vdef16           string  `json:"vdef16" xml:"vdef16"`
			Vdef17           string  `json:"vdef17" xml:"vdef17"`
			Vdef18           string  `json:"vdef18" xml:"vdef18"`
			Vdef19           string  `json:"vdef19" xml:"vdef19"`
			Vdef20           string  `json:"vdef20" xml:"vdef20"`
			Billmaker        string  `json:"billmaker" xml:"billmaker"`             //制单人
			Dmakedate        string  `json:"dmakedate" xml:"dmakedate"`             //制单日期
			Creator          string  `json:"creator" xml:"creator"`                 //创建人
			Creationtime     string  `json:"creationtime" xml:"creationtime"`       //创建时间
			Modifier         string  `json:"modifier" xml:"modifier"`               //修改人
			Modifiedtime     string  `json:"modifiedtime" xml:"modifiedtime"`       //修改时间
			Approver         string  `json:"approver" xml:"approver"`               //审批人
			Taudittime       string  `json:"taudittime" xml:"taudittime"`           //审核日期
			Creviserid       string  `json:"creviserid" xml:"creviserid"`           //修订人
			Iversion         int     `json:"iversion" xml:"iversion"`               //修订版本号
			Vrevisereason    string  `json:"vrevisereason" xml:"vrevisereason"`     //修订理由
			Trevisetime      string  `json:"trevisetime" xml:"trevisetime"`         //修订时间
			Bsendendflag     string  `json:"bsendendflag" xml:"bsendendflag"`       //发货关闭
			Boutendflag      string  `json:"boutendflag" xml:"boutendflag"`         //出库关闭
			Binvoicendflag   string  `json:"binvoicendflag" xml:"binvoicendflag"`   //开票关闭
			Bcostsettleflag  string  `json:"bcostsettleflag" xml:"bcostsettleflag"` //成本结算关闭
			Barsettleflag    string  `json:"barsettleflag" xml:"barsettleflag"`     //收入结算关闭
			Iprintcount      int     `json:"iprintcount" xml:"iprintcount"`         //打印次数
			SoSaleorderB     struct {
				Items []SaleItem `json:"item" xml:"item"`
			} `json:"so_saleorder_b" xml:"so_saleorder_b"`
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type ResSaleOrder struct {
	XMLName    xml.Name `xml:"ufinterface"`
	Successful string   `xml:"successful,attr"`
	base.Base
	Sendresult struct {
		Billpk            string     `xml:"billpk"`
		Bdocid            string     `xml:"bdocid"`
		Filename          string     `xml:"filename"`
		Resultcode        int        `xml:"resultcode"`
		Resultdescription string     `xml:"resultdescription"`
		Invaliddoc        []*Docitem `xml:"invaliddoc"`
		Content           string     `xml:"content"`
	} `xml:"sendresult"`
}

type Docitem struct {
	Name string `xml:"name,attr"`
}

type SaleItem struct {
	Csaleorderbid      string  `json:"csaleorderbid" xml:"csaleorderbid"`
	PkGroup            string  `json:"pk_group" xml:"pk_group"`                     //集团
	PkOrg              string  `json:"pk_org" xml:"pk_org"`                         //销售组织
	Dbilldate          string  `json:"dbilldate" xml:"dbilldate"`                   //单据日期
	Crowno             string  `json:"crowno" xml:"crowno"`                         //行号
	Cmaterialvid       string  `json:"cmaterialvid" xml:"cmaterialvid"`             //物料编码
	Cmaterialid        string  `json:"cmaterialid" xml:"cmaterialid"`               //物料
	Cvendorid          string  `json:"cvendorid" xml:"cvendorid"`                   //供应商
	Cprojectid         string  `json:"cprojectid" xml:"cprojectid"`                 //项目
	Cproductorid       string  `json:"cproductorid" xml:"cproductorid"`             //生产厂商
	Cfactoryid         string  `json:"cfactoryid" xml:"cfactoryid"`                 //工厂
	Cqualitylevelid    string  `json:"cqualitylevelid" xml:"cqualitylevelid"`       //质量等级
	Corigcountryid     string  `json:"corigcountryid" xml:"corigcountryid"`         //原产国
	Corigareaid        string  `json:"corigareaid" xml:"corigareaid"`               //原产地区
	Cunitid            string  `json:"cunitid" xml:"cunitid"`                       //主单位
	Castunitid         string  `json:"castunitid" xml:"castunitid"`                 //单位
	Nnum               int     `json:"nnum" xml:"nnum"`                             //主数量
	Nastnum            int     `json:"nastnum" xml:"nastnum"`                       //数量
	Vchangerate        string  `json:"vchangerate" xml:"vchangerate"`               //换算率
	Cqtunitid          string  `json:"cqtunitid" xml:"cqtunitid"`                   //报价单位
	Nqtunitnum         float64 `json:"nqtunitnum" xml:"nqtunitnum"`                 //报价单位数量
	Vqtunitrate        string  `json:"vqtunitrate" xml:"vqtunitrate"`               //报价换算率
	Ndiscountrate      float64 `json:"ndiscountrate" xml:"ndiscountrate"`           //整单折扣
	Nitemdiscountrate  float64 `json:"nitemdiscountrate" xml:"nitemdiscountrate"`   //单品折扣
	Ctaxcodeid         string  `json:"ctaxcodeid" xml:"ctaxcodeid"`                 //税码
	Ntaxrate           float64 `json:"ntaxrate" xml:"ntaxrate"`                     //税率
	Ftaxtypeflag       int     `json:"ftaxtypeflag" xml:"ftaxtypeflag"`             //扣税类别
	Ccurrencyid        string  `json:"ccurrencyid" xml:"ccurrencyid"`               //本位币
	Nexchangerate      float64 `json:"nexchangerate" xml:"nexchangerate"`           //折本汇率
	Nqtorigtaxprice    float64 `json:"nqtorigtaxprice" xml:"nqtorigtaxprice"`       //含税单价
	Nqtorigprice       float64 `json:"nqtorigprice" xml:"nqtorigprice"`             //无税单价
	Nqtorigtaxnetprc   float64 `json:"nqtorigtaxnetprc" xml:"nqtorigtaxnetprc"`     //含税净价
	Nqtorignetprice    float64 `json:"nqtorignetprice" xml:"nqtorignetprice"`       //无税净价
	Norigtaxprice      float64 `json:"norigtaxprice" xml:"norigtaxprice"`           //主含税单价
	Norigprice         float64 `json:"norigprice" xml:"norigprice"`                 //主无税单价
	Norigtaxnetprice   float64 `json:"norigtaxnetprice" xml:"norigtaxnetprice"`     //主含税净价
	Norignetprice      float64 `json:"norignetprice" xml:"norignetprice"`           //主无税净价
	Ntax               float64 `json:"ntax" xml:"ntax"`                             //税额
	Ncaltaxmny         float64 `json:"ncaltaxmny" xml:"ncaltaxmny"`                 //计税金额
	Norigmny           float64 `json:"norigmny" xml:"norigmny"`                     //无税金额
	Norigtaxmny        float64 `json:"norigtaxmny" xml:"norigtaxmny"`               //价税合计
	Norigdiscount      float64 `json:"norigdiscount" xml:"norigdiscount"`           //折扣额
	Nbforigsubmny      float64 `json:"nbforigsubmny" xml:"nbforigsubmny"`           //冲抵前金额
	Nqttaxprice        float64 `json:"nqttaxprice" xml:"nqttaxprice"`               //本币含税单价
	Nqtprice           float64 `json:"nqtprice" xml:"nqtprice"`                     //本币无税单价
	Nqttaxnetprice     float64 `json:"nqttaxnetprice" xml:"nqttaxnetprice"`         //本币含税净价
	Nqtnetprice        float64 `json:"nqtnetprice" xml:"nqtnetprice"`               //本币无税净价
	Ntaxprice          float64 `json:"ntaxprice" xml:"ntaxprice"`                   //主本币含税单价
	Nprice             float64 `json:"nprice" xml:"nprice"`                         //主本币无税单价
	Ntaxnetprice       float64 `json:"ntaxnetprice" xml:"ntaxnetprice"`             //主本币含税净价
	Nnetprice          float64 `json:"nnetprice" xml:"nnetprice"`                   //主本币无税净价
	Nmny               float64 `json:"nmny" xml:"nmny"`                             //本币无税金额
	Ntaxmny            float64 `json:"ntaxmny" xml:"ntaxmny"`                       //本币价税合计
	Ndiscount          float64 `json:"ndiscount" xml:"ndiscount"`                   //本币折扣额
	Ngroupexchgrate    float64 `json:"ngroupexchgrate" xml:"ngroupexchgrate"`       //集团本位币汇率
	Ngroupmny          float64 `json:"ngroupmny" xml:"ngroupmny"`                   //集团本币无税金额
	Ngrouptaxmny       float64 `json:"ngrouptaxmny" xml:"ngrouptaxmny"`             //集团本币价税合计
	Nglobalexchgrate   float64 `json:"nglobalexchgrate" xml:"nglobalexchgrate"`     //全局本位币汇率
	Nglobalmny         float64 `json:"nglobalmny" xml:"nglobalmny"`                 //全局本币无税金额
	Nglobaltaxmny      float64 `json:"nglobaltaxmny" xml:"nglobaltaxmny"`           //全局本币价税合计
	Naskqtorigtaxprc   float64 `json:"naskqtorigtaxprc" xml:"naskqtorigtaxprc"`     //询价原币含税单价
	Naskqtorigprice    float64 `json:"naskqtorigprice" xml:"naskqtorigprice"`       //询价原币无税单价
	Naskqtorigtxntprc  float64 `json:"naskqtorigtxntprc" xml:"naskqtorigtxntprc"`   //询价原币含税净价
	Naskqtorignetprice float64 `json:"naskqtorignetprice" xml:"naskqtorignetprice"` //询价原币无税净价
	Nweight            float64 `json:"nweight" xml:"nweight"`                       //重量
	Nvolume            float64 `json:"nvolume" xml:"nvolume"`                       //体积
	Npiece             float64 `json:"npiece" xml:"npiece"`                         //件数
	Cprodlineid        string  `json:"cprodlineid" xml:"cprodlineid"`               //产品线
	Vbatchcode         string  `json:"vbatchcode" xml:"vbatchcode"`                 //批次号
	PkBatchcode        string  `json:"pk_batchcode" xml:"pk_batchcode"`
	Blaborflag         string  `json:"blaborflag" xml:"blaborflag"`               //服务类
	Bdiscountflag      string  `json:"bdiscountflag" xml:"bdiscountflag"`         //折扣类
	Blargessflag       string  `json:"blargessflag" xml:"blargessflag"`           //赠品
	Bbindflag          string  `json:"bbindflag" xml:"bbindflag"`                 //捆绑存货
	Dsenddate          string  `json:"dsenddate" xml:"dsenddate"`                 //要求发货日期
	Dreceivedate       string  `json:"dreceivedate" xml:"dreceivedate"`           //计划到货日期
	Creceivecustid     string  `json:"creceivecustid" xml:"creceivecustid"`       //收货客户
	Creceiveareaid     string  `json:"creceiveareaid" xml:"creceiveareaid"`       //收货地区
	Creceiveaddrid     string  `json:"creceiveaddrid" xml:"creceiveaddrid"`       //收货地址
	Creceiveadddocid   string  `json:"creceiveadddocid" xml:"creceiveadddocid"`   //收货地点
	Csendstockorgvid   string  `json:"csendstockorgvid" xml:"csendstockorgvid"`   //发货库存组织
	Csendstockorgid    string  `json:"csendstockorgid" xml:"csendstockorgid"`     //发货库存组织最新版本
	Csendstordocid     string  `json:"csendstordocid" xml:"csendstordocid"`       //发货仓库
	Ctrafficorgvid     string  `json:"ctrafficorgvid" xml:"ctrafficorgvid"`       //物流组织
	Ctrafficorgid      string  `json:"ctrafficorgid" xml:"ctrafficorgid"`         //物流组织
	Csettleorgvid      string  `json:"csettleorgvid" xml:"csettleorgvid"`         //结算财务组织
	Csettleorgid       string  `json:"csettleorgid" xml:"csettleorgid"`           //结算财务组织
	Crececountryid     string  `json:"crececountryid" xml:"crececountryid"`       //收货国家/地区
	Csendcountryid     string  `json:"csendcountryid" xml:"csendcountryid"`       //发货国家/地区
	Ctaxcountryid      string  `json:"ctaxcountryid" xml:"ctaxcountryid"`         //报税国家/地区
	Fbuysellflag       int     `json:"fbuysellflag" xml:"fbuysellflag"`           //购销类型
	Btriatradeflag     string  `json:"btriatradeflag" xml:"btriatradeflag"`       //三角贸易
	Carorgvid          string  `json:"carorgvid" xml:"carorgvid"`                 //应收组织
	Carorgid           string  `json:"carorgid" xml:"carorgid"`                   //应收组织最新版本
	Cprofitcentervid   string  `json:"cprofitcentervid" xml:"cprofitcentervid"`   //利润中心
	Cprofitcenterid    string  `json:"cprofitcenterid" xml:"cprofitcenterid"`     //利润中心
	Vbrevisereason     string  `json:"vbrevisereason" xml:"vbrevisereason"`       //修订理由
	Frowstatus         int     `json:"frowstatus" xml:"frowstatus"`               //行状态
	Vrownote           string  `json:"vrownote" xml:"vrownote"`                   //行备注
	Cpriceformid       string  `json:"cpriceformid" xml:"cpriceformid"`           //价格组成
	Cpricepolicyid     string  `json:"cpricepolicyid" xml:"cpricepolicyid"`       //价格政策
	Cpriceitemid       string  `json:"cpriceitemid" xml:"cpriceitemid"`           //价格项目
	Cpriceitemtableid  string  `json:"cpriceitemtableid" xml:"cpriceitemtableid"` //价目表
	Bjczxsflag         string  `json:"bjczxsflag" xml:"bjczxsflag"`               //借出转销售
	Vcttype            string  `json:"vcttype" xml:"vcttype"`                     //合同类型
	Vctcode            string  `json:"vctcode" xml:"vctcode"`                     //销售合同号
	Cctmanageid        string  `json:"cctmanageid" xml:"cctmanageid"`             //合同主表
	Ctmanagebid        string  `json:"ctmanagebid" xml:"ctmanagebid"`             //
	Vsrctype           string  `json:"vsrctype" xml:"vsrctype"`                   //来源单据类型
	Vsrctrantype       string  `json:"vsrctrantype" xml:"vsrctrantype"`           //来源交易类型
	Vsrccode           string  `json:"vsrccode" xml:"vsrccode"`                   //来源单据号
	Vsrcrowno          string  `json:"vsrcrowno" xml:"vsrcrowno"`                 //来源单据行号
	Csrcid             string  `json:"csrcid" xml:"csrcid"`                       //
	Csrcbid            string  `json:"csrcbid" xml:"csrcbid"`                     //
	Vfirsttype         string  `json:"vfirsttype" xml:"vfirsttype"`               //源头单据类型
	Vfirsttrantype     string  `json:"vfirsttrantype" xml:"vfirsttrantype"`       //源头交易类型
	Vfirstcode         string  `json:"vfirstcode" xml:"vfirstcode"`               //源头单据号
	Cfirstid           string  `json:"cfirstid" xml:"cfirstid"`                   //
	Cfirstbid          string  `json:"cfirstbid" xml:"cfirstbid"`                 //
	Vfirstrowno        string  `json:"vfirstrowno" xml:"vfirstrowno"`             //源头单据行号
	Cretreasonid       string  `json:"cretreasonid" xml:"cretreasonid"`           //退货原因
	Cretpolicyid       string  `json:"cretpolicyid" xml:"cretpolicyid"`           //退货政策
	Vreturnmode        string  `json:"vreturnmode" xml:"vreturnmode"`             //退货责任处理方式
	Fretexchange       int     `json:"fretexchange" xml:"fretexchange"`           //退换货标记
	Cexchangesrcretid  string  `json:"cexchangesrcretid" xml:"cexchangesrcretid"`
	Clargesssrcid      string  `json:"clargesssrcid" xml:"clargesssrcid"`       //
	Cbindsrcid         string  `json:"cbindsrcid" xml:"cbindsrcid"`             //
	Flargesstypeflag   string  `json:"flargesstypeflag" xml:"flargesstypeflag"` //赠品价格分摊方式
	Nlargessmny        float64 `json:"nlargessmny" xml:"nlargessmny"`           //赠品价格分摊前无税金额
	Nlargesstaxmny     string  `json:"nlargesstaxmny" xml:"nlargesstaxmny"`     //赠品价格分摊前价税合计
	Bprerowcloseflag   string  `json:"bprerowcloseflag" xml:"bprerowcloseflag"` //预订单行关闭
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
	Bbsendendflag      string  `json:"bbsendendflag" xml:"bbsendendflag"`         //发货关闭
	Bbinvoicendflag    string  `json:"bbinvoicendflag" xml:"bbinvoicendflag"`     //开票关闭
	Bboutendflag       string  `json:"bboutendflag" xml:"bboutendflag"`           //出库关闭
	Bbcostsettleflag   string  `json:"bbcostsettleflag" xml:"bbcostsettleflag"`   //成本结算关闭
	Bbarsettleflag     string  `json:"bbarsettleflag" xml:"bbarsettleflag"`       //收入结算关闭
	Vclosereason       string  `json:"vclosereason" xml:"vclosereason"`           //关闭/打开原因
	Nsendauditnum      float64 `json:"nsendauditnum" xml:"nsendauditnum"`         //发货审批主数量
	Noutauditnum       float64 `json:"noutauditnum" xml:"noutauditnum"`           //出库审批主数量
	Ninvoiceauditnum   float64 `json:"ninvoiceauditnum" xml:"ninvoiceauditnum"`   //发票审批主数量
	Noutnotauditnum    float64 `json:"noutnotauditnum" xml:"noutnotauditnum"`     //出库未签字主数量
	Nlossnotauditnum   float64 `json:"nlossnotauditnum" xml:"nlossnotauditnum"`   //途损单未审核主数量
	Ntotalsendnum      float64 `json:"ntotalsendnum" xml:"ntotalsendnum"`         //累计发货主数量
	Ntotalinvoicenum   float64 `json:"ntotalinvoicenum" xml:"ntotalinvoicenum"`   //累计开票主数量
	Ntotaloutnum       float64 `json:"ntotaloutnum" xml:"ntotaloutnum"`           //累计出库主数量
	Ntotalnotoutnum    float64 `json:"ntotalnotoutnum" xml:"ntotalnotoutnum"`     //累计应发未出库主数量
	Ntotalsignnum      float64 `json:"ntotalsignnum" xml:"ntotalsignnum"`         //累计签收主数量
	Ntranslossnum      float64 `json:"ntranslossnum" xml:"ntranslossnum"`         //累计途损主数量
	Ntotalrushnum      float64 `json:"ntotalrushnum" xml:"ntotalrushnum"`         //累计出库对冲主数量
	Ntotalestarnum     float64 `json:"ntotalestarnum" xml:"ntotalestarnum"`       //累计暂估应收主数量
	Ntotalarnum        float64 `json:"ntotalarnum" xml:"ntotalarnum"`             //累计确认应收主数量
	Ntotalcostnum      float64 `json:"ntotalcostnum" xml:"ntotalcostnum"`         //累计成本结算主数量
	Ntotalestarmny     float64 `json:"ntotalestarmny" xml:"ntotalestarmny"`       //累计暂估应收金额
	Ntotalarmny        float64 `json:"ntotalarmny" xml:"ntotalarmny"`             //累计确认应收金额
	Ntotalpaymny       float64 `json:"ntotalpaymny" xml:"ntotalpaymny"`           //累计财务核销金额
	Norigsubmny        float64 `json:"norigsubmny" xml:"norigsubmny"`             //累计冲抵金额
	Narrangescornum    float64 `json:"narrangescornum" xml:"narrangescornum"`     //累计安排委外订单主数量
	Narrangepoappnum   float64 `json:"narrangepoappnum" xml:"narrangepoappnum"`   //累计安排请购单主数量
	Narrangetoornum    float64 `json:"narrangetoornum" xml:"narrangetoornum"`     //累计安排调拨订单主数量
	Narrangetoappnum   float64 `json:"narrangetoappnum" xml:"narrangetoappnum"`   //累计安排调拨申请主数量
	Narrangemonum      float64 `json:"narrangemonum" xml:"narrangemonum"`         //累计安排生产订单主数量
	Narrangeponum      float64 `json:"narrangeponum" xml:"narrangeponum"`         //累计安排采购订单主数量
	Ntotalplonum       float64 `json:"ntotalplonum" xml:"ntotalplonum"`           //累计生成计划订单主数量
	Nreqrsnum          float64 `json:"nreqrsnum" xml:"nreqrsnum"`                 //预留主数量
	Ntotalreturnnum    float64 `json:"ntotalreturnnum" xml:"ntotalreturnnum"`     //累计退货主数量
	Ntotaltradenum     float64 `json:"ntotaltradenum" xml:"ntotaltradenum"`       //累计发出商品主数量
	Barrangedflag      string  `json:"barrangedflag" xml:"barrangedflag"`         //是否货源安排完毕
	Carrangepersonid   string  `json:"carrangepersonid" xml:"carrangepersonid"`   //
	Tlastarrangetime   string  `json:"tlastarrangetime" xml:"tlastarrangetime"`   //
	Nsendunfinisednum  float64 `json:"nsendunfinisednum" xml:"nsendunfinisednum"` //发货未完成量
	Noutunfinisednum   float64 `json:"noutunfinisednum" xml:"noutunfinisednum"`   //出库未完成量
	Ninvunfinisednum   float64 `json:"ninvunfinisednum" xml:"ninvunfinisednum"`   //发票未完成量
	Nnotarnum          float64 `json:"nnotarnum" xml:"nnotarnum"`                 //未传确认应收主数量
	Nnotcostnum        float64 `json:"nnotcostnum" xml:"nnotcostnum"`             //未传存货核算主数量
	Bbsettleendflag    string  `json:"bbsettleendflag" xml:"bbsettleendflag"`     //结算关闭
	Srcts              string  `json:"srcts" xml:"srcts"`                         //来源单据表头时间戳
	Srcbts             string  `json:"srcbts" xml:"srcbts"`                       //来源单据表体时间戳
	Srcorgid           string  `json:"srcorgid" xml:"srcorgid"`                   //
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
}

type ReqSaleOutWare struct {
	base.Base
	Bill struct {
		Billhead struct {
			Cgeneralhid string `json:"cgeneralhid" xml:"cgeneralhid"`
			PkGroup     string `json:"pk_group" xml:"pk_group"` //集团
			PkOrg       string `json:"pk_org" xml:"pk_org"`     //销售组织
			PkOrgV      string `json:"pk_org_v" xml:"pk_org_v"` //销售组织版本
			Cgeneralbid struct {
				Cgeneralibid string         `json:"cgeneralibid" xml:"cgeneralibid"`
				Items        []*SaleOutItem `json:"item" xml:"item"`
			} `json:"cgeneralbid" xml:"cgeneralbid"`
			Corpoid          string  `json:"corpoid" xml:"corpoid"`                   //公司最新版本
			Corpvid          string  `json:"corpvid" xml:"corpvid"`                   //公司
			Cbiztype         string  `json:"cbiztype" xml:"cbiztype"`                 //业务流程
			Vbillcode        string  `json:"vbillcode" xml:"vbillcode"`               //单据号
			Cwarehouseid     string  `json:"cwarehouseid" xml:"cwarehouseid"`         //仓库
			Dbilldate        string  `json:"dbilldate" xml:"dbilldate"`               //单据日期
			Vtrantypecode    string  `json:"vtrantypecode" xml:"vtrantypecode"`       //出入库类型编码
			Trafficorgvid    string  `json:"trafficorgvid" xml:"trafficorgvid"`       //物流组织
			Cothercalbodyoid string  `json:"cothercalbodyoid" xml:"cothercalbodyoid"` //入库库存组织最新版本
			Cothercalbodyvid string  `json:"cothercalbodyvid" xml:"cothercalbodyvid"` //入库库存组织
			Cotherwhid       string  `json:"cotherwhid" xml:"cotherwhid"`             //入库仓库
			Cwhsmanagerid    string  `json:"cwhsmanagerid" xml:"cwhsmanagerid"`       //库管员
			Ntotalnum        float64 `json:"ntotalnum" xml:"ntotalnum"`               //总数量
			Ntotalweight     float64 `json:"ntotalweight" xml:"ntotalweight"`         //总重量
			Ntotalvolume     float64 `json:"ntotalvolume" xml:"ntotalvolume"`         //总体积
			Ntotalpiece      float64 `json:"ntotalpiece" xml:"ntotalpiece"`           //总件数
			Freplenishflag   string  `json:"freplenishflag" xml:"freplenishflag"`     //采购退库
			Boutretflag      string  `json:"boutretflag" xml:"boutretflag"`           //销售退回
			Vdiliveraddress  string  `json:"vdiliveraddress" xml:"vdiliveraddress"`   //运输地址
			Cdilivertypeid   string  `json:"cdilivertypeid" xml:"cdilivertypeid"`     //运输方式
			Ccustomerid      string  `json:"ccustomerid" xml:"ccustomerid"`           //客户
			PkMeasware       string  `json:"pk_measware" xml:"pk_measware"`           //计量器具
			Vnote            string  `json:"vnote" xml:"vnote"`                       //备注
			Fbillflag        int     `json:"fbillflag" xml:"fbillflag"`               //单据状态
			Iprintcount      int     `json:"iprintcount" xml:"iprintcount"`           //打印次数
			Cdptid           string  `json:"cdptid" xml:"cdptid"`
			//                 vdef1
			//                 vdef2
			//                 vdef3
			//                 vdef4
			//                 vdef5
			//                 vdef6
			//                 vdef7
			//                 vdef8
			//                 vdef9
			//                 vdef10
			//                 vdef11
			//                 vdef12
			//                 vdef13
			//                 vdef14
			//                 vdef15
			//                 vdef16
			//                 vdef17
			//                 vdef18
			//                 vdef19
			//                 vdef20
			Ctrancustid  string `json:"ctrancustid" xml:"ctrancustid"`   //承运商
			Bsalecooppur string `json:"bsalecooppur" xml:"bsalecooppur"` //已协同生成采购入库单
			// vreturnreason           //退库理由
			Creator      string `json:"creator" xml:"creator"`           //创建人
			Billmaker    string `json:"billmaker" xml:"billmaker"`       //制单人
			Creationtime string `json:"creationtime" xml:"creationtime"` //制单时间
			Approver     string `json:"approver" xml:"approver"`         //签字人
			Taudittime   string `json:"taudittime" xml:"taudittime"`     //签字时间
			Modifier     string `json:"modifier" xml:"modifier"`         //最后修改人
			Modifiedtime string `json:"modifiedtime" xml:"modifiedtime"` //最后修改时间
			//                 cpayfinorgoid           //应付财务组织最新版本
			//                 cpayfinorgvid           //应付财务组织
			//                 cfanaceorgoid           //结算财务组织最新版本
			//                 cfanaceorgvid           //结算财务组织
			//                 cpurorgoid              //采购组织
			//                 cpurorgvid              //采购组织最新版本
			// cdptid                  //采购部门OID
			Cdptvid          string `json:"cdptvid" xml:"cdptvid"`                   //采购部门VID
			Csaleorgoid      string `json:"csaleorgoid" xml:"csaleorgoid"`           //销售组织最新版本
			Csaleorgvid      string `json:"csaleorgvid" xml:"csaleorgvid"`           //销售组织
			Cfanaceorgoid    string `json:"cfanaceorgoid" xml:"cfanaceorgoid"`       //结算财务组织最新版本
			Cfanaceorgvid    string `json:"cfanaceorgvid" xml:"cfanaceorgvid"`       //结算财务组织
			Creceivfinorgoid string `json:"creceivfinorgoid" xml:"creceivfinorgoid"` //应收财务组织最新版本
			Creceivfinorgvid string `json:"creceivfinorgvid" xml:"creceivfinorgvid"` //应收财务组织
			Ctrantypeid      string `json:"ctrantypeid" xml:"ctrantypeid"`           //出入库类型
			Csendcountryid   string `json:"csendcountryid" xml:"csendcountryid"`     //发货国
			Crececountryid   string `json:"crececountryid" xml:"crececountryid"`     //收货国
			Ctaxcountryid    string `json:"ctaxcountryid" xml:"ctaxcountryid"`       //报税国
			Fbuysellflag     string `json:"fbuysellflag" xml:"fbuysellflag"`         //购销类型
			Btriatradeflag   string `json:"btriatradeflag" xml:"btriatradeflag"`     //三角贸易
			Ctradewordid     string `json:"ctradewordid" xml:"ctradewordid"`         //贸易术语
			Cbizid           string `json:"cbizid" xml:"cbizid"`                     //采购员
			Trafficorgoid    string `json:"trafficorgoid" xml:"trafficorgoid"`       //物流组织最新版本
			Cvendorid        string `json:"cvendorid" xml:"cvendorid"`               //供应商
			// csendtypeid             //运输方式
			// ctradewordid            //贸易术语
		} `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type SaleOutItem struct {
	Cgeneralbid string `json:"cgeneralbid" xml:"cgeneralbid"`
	Cgenerallid struct {
		Items []*CgenerallidItem `json:"item" xml:"item"`
	} `json:"cgenerallid" xml:"cgenerallid"`
	PkGroup string `json:"pk_group" xml:"pk_group"` //集团
	Corpoid string `json:"corpoid" xml:"corpoid"`   //公司最新版本
	Corpvid string `json:"corpvid" xml:"corpvid"`   //公司
	Crowno  string `json:"crowno" xml:"crowno"`     //行号
	// cstateid            //库存状态
	Cmaterialoid string `json:"cmaterialoid" xml:"cmaterialoid"` //物料编码
	Cmaterialvid string `json:"cmaterialvid" xml:"cmaterialvid"` //物料版本
	Cunitid      string `json:"cunitid" xml:"cunitid"`           //主单位
	//                         vfree1
	//                         vfree2
	//                         vfree3
	//                         vfree4
	//                         vfree5
	//                         vfree6
	//                         vfree7
	//                         vfree8
	//                         vfree9
	//                         vfree10
	Castunitid       string  `json:"castunitid" xml:"castunitid"`             //单位
	Vchangerate      string  `json:"vchangerate" xml:"vchangerate"`           //换算率
	Vbatchcode       string  `json:"vbatchcode" xml:"vbatchcode"`             //批次号
	Nshouldassistnum float64 `json:"nshouldassistnum" xml:"nshouldassistnum"` //应收数量
	Nshouldnum       float64 `json:"nshouldnum" xml:"nshouldnum"`             //应收主数量
	Nassistnum       float64 `json:"nassistnum" xml:"nassistnum"`             //实收数量
	Nnum             float64 `json:"nnum" xml:"nnum"`                         //实收主数量
	Ncostprice       float64 `json:"ncostprice" xml:"ncostprice"`             //单价
	Ncostmny         float64 `json:"ncostmny" xml:"ncostmny"`                 //金额
	Nplannedprice    float64 `json:"nplannedprice" xml:"nplannedprice"`       //计划单价
	Nplannedmny      float64 `json:"nplannedmny" json:"nplannedmny"`          //计划金额
	Clocationid      string  `json:"clocationid" xml:"clocationid"`           //货位
	Dbizdate         string  `json:"dbizdate" xml:"dbizdate"`                 //入库日期
	Cinvoicecustid   string  `json:"cinvoicecustid" xml:"cinvoicecustid"`     //开票客户
	//private int fchecked                //待检标志
	// ccheckstateid           //开票客户
	Cprojectid     string `json:"cprojectid" xml:"cprojectid"`         //项目
	Ccorrespondhid string `json:"ccorrespondhid" xml:"ccorrespondhid"` //
	// ccorrespondbid          //
	Ccorrespondtype      string `json:"ccorrespondtype" xml:"ccorrespondtype"`           //对应入库单类型
	Ccorrespondtranstype string `json:"ccorrespondtranstype" xml:"ccorrespondtranstype"` //对应入库单交易类型
	Ccorrespondcode      string `json:"ccorrespondcode" xml:"ccorrespondcode"`           //对应入库单单据号
	Ccorrespondrowno     string `json:"ccorrespondrowno" xml:"ccorrespondrowno"`         //对应入库单行号
	Flargess             string `json:"flargess" xml:"flargess"`                         //赠品
	Bsourcelargess       string `json:"bsourcelargess" xml:"bsourcelargess"`             //上游赠品行
	Csourcebillhid       string `json:"csourcebillhid" xml:"csourcebillhid"`             //
	Csourcebillbid       string `json:"csourcebillbid" xml:"csourcebillbid"`             //
	// carriveorder_bbid       //
	Csourcetype        string `json:"csourcetype" xml:"csourcetype"`               //来源单据类型
	Csourcetranstype   string `json:"csourcetranstype" xml:"csourcetranstype"`     //来源单据交易类型
	Vsourcebillcode    string `json:"vsourcebillcode" xml:"vsourcebillcode"`       //来源单据号
	Vsourcerowno       string `json:"vsourcerowno" xml:"vsourcerowno"`             //来源单据行号
	Csourcematerialoid string `json:"csourcematerialoid" xml:"csourcematerialoid"` //
	Cfirsttype         string `json:"cfirsttype" xml:"cfirsttype"`                 //源头单据类型
	Cfirsttranstype    string `json:"cfirsttranstype" xml:"cfirsttranstype"`       //源头单据交易类型
	Cfirstbillhid      string `json:"cfirstbillhid" xml:"cfirstbillhid"`           //
	Cfirstbillbid      string `json:"cfirstbillbid" xml:"cfirstbillbid"`
	// corder_bb1id            //
	Vfirstbillcode string `json:"vfirstbillcode" xml:"vfirstbillcode"` //源头单据号
	Vfirstrowno    string `json:"vfirstrowno" xml:"vfirstrowno"`       //源头单据行号
	Csrc2billtype  string `json:"csrc2billtype" xml:"csrc2billtype"`   //其他来源单据类型编码
	Csrc2transtype string `json:"csrc2transtype" xml:"csrc2transtype"` //其他来源交易类型编码
	Csrc2billhid   string `json:"csrc2billhid" xml:"csrc2billhid"`     //
	Vsrc2billcode  string `json:"vsrc2billcode" xml:"vsrc2billcode"`   //其他来源单据号
	Csrc2billbid   string `json:"csrc2billbid" xml:"csrc2billbid"`     //
	Vsrc2billrowno string `json:"vsrc2billrowno" xml:"vsrc2billrowno"` //其他来源单行号
	Cvendorid      string `json:"cvendorid" xml:"cvendorid"`           //供应商
	Casscustid     string `json:"casscustid" xml:"casscustid"`         //客户
	Ccostobject    string `json:"ccostobject" xml:"ccostobject"`       //成本对象
	// cffileid                //特征码
	Fbillrowflag  string  `json:"fbillrowflag" xml:"fbillrowflag"`   //配套标志
	Vnotebody     string  `json:"vnotebody" xml:"vnotebody"`         //行备注
	Creceieveid   string  `json:"creceieveid" xml:"creceieveid"`     //
	Ncountnum     string  `json:"ncountnum" xml:"ncountnum"`         //箱数
	PkPacksort    string  `json:"pk_packsort" xml:"pk_packsort"`     //包装类型
	Nbarcodenum   float64 `json:"nbarcodenum" xml:"nbarcodenum"`     //条码主数量
	Bbarcodeclose string  `json:"bbarcodeclose" xml:"bbarcodeclose"` //单据行是否条码关闭
	Bonroadflag   string  `json:"bonroadflag" xml:"bonroadflag"`     //是否在途
	Vvehiclecode  string  `json:"vvehiclecode" xml:"vvehiclecode"`   //运输工具号
	Vtransfercode string  `json:"vtransfercode" xml:"vtransfercode"` //收货车号
	//                         vbdef1
	//                         vbdef2
	//                         vbdef3
	//                         vbdef4
	//                         vbdef5
	//                         vbdef6
	//                         vbdef7
	//                         vbdef8
	//                         vbdef9
	//                         vbdef10
	//                         vbdef11
	//                         vbdef12
	//                         vbdef13
	//                         vbdef14
	//                         vbdef15
	//                         vbdef16
	//                         vbdef17
	//                         vbdef18
	//                         vbdef19
	//                         vbdef20
	PkMeasware           string  `json:"pk_measware" xml:"pk_measware"` //计量器具
	Ngrossnum            float64 `json:"ngrossnum" xml:"ngrossnum"`     //毛重主数量
	Ntarenum             float64 `json:"ntarenum" xml:"ntarenum"`       //皮重主数量
	Irdesatype           int     `json:"irdesatype" xml:"irdesatype"`   //拆解类型
	Cselastunitid        string  `json:"cselastunitid" xml:"cselastunitid"`
	Dproducedate         string  `json:"dproducedate" xml:"dproducedate"`
	Dvalidate            string  `json:"dvalidate" xml:"dvalidate"`
	Cbodytranstypecode   string  `json:"cbodytranstypecode" xml:"cbodytranstypecode"` //出入库类型
	PkOrg                string  `json:"pk_org" xml:"pk_org"`                         //库存组织最新版本
	PkOrgV               string  `json:"pk_org_v" xml:"pk_org_v"`                     //库存组织
	Cbodywarehouseid     string  `json:"cbodywarehouseid" xml:"cbodywarehouseid"`     //库存仓库
	Cproductorid         string  `json:"cproductorid" xml:"cproductorid"`             //生产厂商
	Cvmivenderid         string  `json:"cvmivenderid" xml:"cvmivenderid"`             //寄存供应商
	Bsafeprice           string  `json:"bsafeprice" xml:"bsafeprice"`                 //价保
	Breturnprofit        string  `json:"breturnprofit" xml:"breturnprofit"`           //返利
	PkReturnreason       string  `json:"pk_returnreason" xml:"pk_returnreason"`       //退货原因
	Btou8rm              string  `json:"btou8rm" xml:"btou8rm"`                       //已下发零售
	Vbilltypeu8rm        string  `json:"vbilltypeu8rm" xml:"vbilltypeu8rm"`           //来自于零售之单据类型
	Vtranstypeu8rm       string  `json:"vtranstypeu8rm" xml:"vtranstypeu8rm"`         //来自于零售之交易类型
	Csignwasttype        string  `json:"csignwasttype" xml:"csignwasttype"`           //出入库单其它来源类型
	Csignwasthid         string  `json:"csignwasthid" xml:"csignwasthid"`             //
	Vsignwastcode        string  `json:"vsignwastcode" xml:"vsignwastcode"`           //出入库单其它来源单据号
	Csignwastbid         string  `json:"csignwastbid" xml:"csignwastbid"`
	Vsignwastrowno       string  `json:"vsignwastrowno" xml:"vsignwastrowno"`             //出入库单其它来源单据行标识
	Creceiveareaid       string  `json:"creceiveareaid" xml:"creceiveareaid"`             //收货地区
	Creceivepointid      string  `json:"creceivepointid" xml:"creceivepointid"`           //收货地点
	Vreceiveaddress      string  `json:"vreceiveaddress" xml:"vreceiveaddress"`           //收货地址
	Ddeliverdate         string  `json:"ddeliverdate" xml:"ddeliverdate"`                 //要求收货日期
	Csourcewasttype      string  `json:"csourcewasttype" xml:"csourcewasttype"`           //来源途损单据类型
	Csourcewasttranstype string  `json:"csourcewasttranstype" xml:"csourcewasttranstype"` //来源途损交易类型
	Csourcewasthid       string  `json:"csourcewasthid" xml:"csourcewasthid"`             //
	Vsourcewastcode      string  `json:"vsourcewastcode" xml:"vsourcewastcode"`           //来源途损单据号
	Csourcewastbid       string  `json:"csourcewastbid" xml:"csourcewastbid"`             //
	Vsourcewastrowno     string  `json:"vsourcewastrowno" xml:"vsourcewastrowno"`         //来源途损单据行标识

	// pk_creqwareid           //需求仓库
	Creqstoorgoid  string `json:"creqstoorgoid" xml:"creqstoorgoid"`   //需求库存组织最新版本
	Creqstoorgvid  string `json:"creqstoorgvid" xml:"creqstoorgvid"`   //需求库存组织
	Ctplcustomerid string `json:"ctplcustomerid" xml:"ctplcustomerid"` //货主客户
	//private double nkdnum                  //扣吨主数量
	Vexigencycode      string  `json:"vexigencycode" xml:"vexigencycode"`           //紧急放行申请单号
	Vexigencyhid       string  `json:"vexigencyhid" xml:"vexigencyhid"`             //
	Vexigencyrowno     string  `json:"vexigencyrowno" xml:"vexigencyrowno"`         //紧急放行申请单行号
	Vexigencybid       string  `json:"vexigencybid" xml:"vexigencybid"`             //
	Cliabilityoid      string  `json:"cliabilityoid" xml:"cliabilityoid"`           //利润中心最新版本
	Cliabilityvid      string  `json:"cliabilityvid" xml:"cliabilityvid"`           //利润中心
	Cprodlineid        string  `json:"cprodlineid" xml:"cprodlineid"`               //产品线
	Cioliabilityoid    string  `json:"cioliabilityoid" xml:"cioliabilityoid"`       //收货利润中心
	Cioliabilityvid    string  `json:"cioliabilityvid" xml:"cioliabilityvid"`       //收货利润中心
	Nweight            float64 `json:"nweight" xml:"nweight"`                       //重量
	Nvolume            float64 `json:"nvolume" xml:"nvolume"`                       //体积
	Npiece             float64 `json:"npiece" xml:"npiece"`                         //件数
	Borrowinflag       string  `json:"borrowinflag" xml:"borrowinflag"`             //借入转采购
	Corigcurrencyid    string  `json:"corigcurrencyid" xml:"corigcurrencyid"`       //币种
	Nbdiscountrate     string  `json:"nbdiscountrate" xml:"nbdiscountrate"`         //整单折扣
	Nchangestdrate     string  `json:"nchangestdrate" xml:"nchangestdrate"`         //折本汇率
	Ccurrencyid        string  `json:"ccurrencyid" xml:"ccurrencyid"`               //本位币
	Norignetprice      float64 `json:"norignetprice" xml:"norignetprice"`           //主无税净价
	Norigtaxnetprice   float64 `json:"norigtaxnetprice" xml:"norigtaxnetprice"`     //主含税净价
	Nnetprice          float64 `json:"nnetprice" xml:"nnetprice"`                   //主本币无税净价
	Ntaxnetprice       float64 `json:"ntaxnetprice" xml:"ntaxnetprice"`             //主本币含税净价
	Cqtunitid          string  `json:"cqtunitid" xml:"cqtunitid"`                   //报价单位
	Nqtunitnum         float64 `json:"nqtunitnum" xml:"nqtunitnum"`                 //报价数量
	Vqtunitrate        string  `json:"vqtunitrate" xml:"vqtunitrate"`               //报价换算率
	Nqtorignetprice    float64 `json:"nqtorignetprice" xml:"nqtorignetprice"`       //无税净价
	Nqtorigtaxnetprice float64 `json:"nqtorigtaxnetprice" xml:"nqtorigtaxnetprice"` //含税净价
	Nqtnetprice        float64 `json:"nqtnetprice" xml:"nqtnetprice"`               //本币无税净价
	Nqttaxnetprice     float64 `json:"nqttaxnetprice" xml:"nqttaxnetprice"`         //本币含税净价
	Norigprice         float64 `json:"norigprice" xml:"norigprice"`                 //主无税单价
	Norigtaxprice      float64 `json:"norigtaxprice" xml:"norigtaxprice"`           //主含税单价
	Nprice             float64 `json:"nprice" xml:"nprice"`                         //主本币无税单价
	Ntaxprice          float64 `json:"ntaxprice" xml:"ntaxprice"`                   //主本币含税单价
	Nqtorigprice       float64 `json:"nqtorigprice" xml:"nqtorigprice"`             //无税单价
	Nqtorigtaxprice    float64 `json:"nqtorigtaxprice" xml:"nqtorigtaxprice"`       //含税单价
	Nqtprice           float64 `json:"nqtprice" xml:"nqtprice"`                     //本币无税单价
	Nqttaxprice        float64 `json:"nqttaxprice" xml:"nqttaxprice"`               //本币含税单价
	Nitemdiscountrate  float64 `json:"nitemdiscountrate" xml:"nitemdiscountrate"`   //折扣
	Norigmny           float64 `json:"norigmny" xml:"norigmny"`                     //无税金额
	Norigtaxmny        float64 `json:"norigtaxmny" xml:"norigtaxmny"`               //价税合计
	Norigtax           float64 `json:"norigtax" xml:"norigtax"`                     //税额
	Nmny               float64 `json:"nmny" xml:"nmny"`                             //本币无税金额
	Ntaxmny            float64 `json:"ntaxmny" xml:"ntaxmny"`                       //本币价税合计
	Ntax               float64 `json:"ntax" xml:"ntax"`                             //本币税额
	Blendoutflag       string  `json:"blendoutflag" xml:"blendoutflag"`             //借出转销售

	// vreturnreason           //退库理由
	Ntaxrate         float64 `json:"ntaxrate" xml:"ntaxrate"`         //税率
	Ftaxtypeflag     int     `json:"ftaxtypeflag" xml:"ftaxtypeflag"` //扣税类别
	Csumid           string  `json:"csumid" xml:"csumid"`
	Naccumwastnum    float64 `json:"naccumwastnum" xml:"naccumwastnum"`       //累计途损主数量
	Btranendflag     string  `json:"btranendflag" xml:"btranendflag"`         //是否运输关闭
	Ntotaltrannum    float64 `json:"ntotaltrannum" xml:"ntotaltrannum"`       //累计运输主数量
	Naccumoutsignnum float64 `json:"naccumoutsignnum" xml:"naccumoutsignnum"` //累计签收主数量
	Naccumoutbacknum float64 `json:"naccumoutbacknum" xml:"naccumoutbacknum"` //累计出库退回主数量
	// tsourcebodyts           //来源表体时间戳
	Ncorrespondnum    float64 `json:"ncorrespondnum" xml:"ncorrespondnum"`       //累计出库主数量
	Ncorrespondastnum float64 `json:"ncorrespondastnum" xml:"ncorrespondastnum"` //累计出库数量
	Ncorrespondgrsnum float64 `json:"ncorrespondgrsnum" xml:"ncorrespondgrsnum"` //累计出库毛重主数量
	//private double naccumtonum             //累计转调拨数量
	Nreplenishednum float64 `json:"nreplenishednum" xml:"nreplenishednum"` //累计退库（补货）主数量
	CdeliveryBbid   string  `json:"cdelivery_bbid" xml:"cdelivery_bbid"`   //
	//private double nreplenishedastnum      //累计退库（补货）数量
	Nsignnum float64 `json:"nsignnum" xml:"nsignnum"` //累计开票主数量
	Nrushnum float64 `json:"nrushnum" xml:"nrushnum"` //累计对冲主数量
	//private double naccumsettlenum         //累计结算数量
	//private double naccumvminum            //累计汇总匹配主数量
	Cglobalcurrencyid string  `json:"cglobalcurrencyid" xml:"cglobalcurrencyid"` //全局本位币
	Nglobalmny        float64 `json:"nglobalmny" xml:"nglobalmny"`               //全局本币无税金额
	Nglobaltaxmny     float64 `json:"nglobaltaxmny" xml:"nglobaltaxmny"`         //全局本币价税合计
	Nglobalexchgrate  float64 `json:"nglobalexchgrate" xml:"nglobalexchgrate"`   //全局本位币汇率
	Cgroupcurrencyid  string  `json:"cgroupcurrencyid" xml:"cgroupcurrencyid"`   //集团本位币
	Ngroupmny         float64 `json:"ngroupmny" xml:"ngroupmny"`                 //集团本币无税金额
	Ngrouptaxmny      float64 `json:"ngrouptaxmny" xml:"ngrouptaxmny"`           //集团本币价税合计
	Ngroupexchgrate   float64 `json:"ngroupexchgrate" xml:"ngroupexchgrate"`     //集团本位币汇率
	PkBatchcode       string  `json:"pk_batchcode" xml:"pk_batchcode"`           //批次主键
	Cfirstbilldate    string  `json:"cfirstbilldate" xml:"cfirstbilldate"`       //源头单据日期
	Nqtorigtaxnetprc  float64 `json:"nqtorigtaxnetprc" xml:"nqtorigtaxnetprc"`   //报价单位含税净价
	Bwastageflag      string  `json:"bwastageflag" xml:"bwastageflag"`           //签收
	Badvfeeflag       string  `json:"badvfeeflag" xml:"badvfeeflag"`             //代垫运费
	Csourcebilldate   string  `json:"csourcebilldate" xml:"csourcebilldate"`     //来源单据日期
	Bassetcard        string  `json:"bassetcard" xml:"bassetcard"`               //已生成设备卡片
	Bfixedasset       string  `json:"bfixedasset" xml:"bfixedasset"`             //已转固
	Csrcmaterialoid   string  `json:"csrcmaterialoid" xml:"csrcmaterialoid"`     //来源物料编码
	Csrcmaterialvid   string  `json:"csrcmaterialvid" xml:"csrcmaterialvid"`     //来源物料版本
	// cprojecttaskid          //项目任务
	// csettlecurrencyid       //币种
	Ctaxcodeid string `json:"ctaxcodeid" xml:"ctaxcodeid"` //税码
	//private double nnosubtaxrate           //不可抵扣税率
	//private double nnosubtax               //不可抵扣税额
	Ncaltaxmny     float64 `json:"ncaltaxmny" xml:"ncaltaxmny"`         //计税税额
	Corigcountryid string  `json:"corigcountryid" xml:"corigcountryid"` //原产国
	Corigareaid    string  `json:"corigareaid" xml:"corigareaid"`       //原产地区
	//  cdesticountryid         //目的国
	// cdestiareaid            //目的地区
	Bopptaxflag string  `json:"bopptaxflag" xml:"bopptaxflag"` //逆向征税标志
	Ncalcostmny float64 `json:"ncalcostmny" xml:"ncalcostmny"` //计成本金额
}

type CgenerallidItem struct {
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

func (sale *Sale) GetSaleOrder(req ReqSaleOrder) (res ResSaleOrder, err error) {
	accessParam := ""
	if accessParam, err = sale.GetUrlParam(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s", accessParam, util.BASE_URL)
	var response []byte
	if response, err = util.PostXML(uri, req); err != nil {
		return
	}

	if err = xml.Unmarshal(response, &res); err != nil {
		return
	}
	if ! (res.Successful=="Y" && res.Sendresult.Resultcode == 1) {
		err = fmt.Errorf("GetSaleOrder Error , errcode=%d , errmsg=%s", res.Sendresult.Resultcode, res.Sendresult.Resultdescription)
		return
	}
	return
}

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

type ReqSaleOrderBillHead struct{
			Csaleorderid     string  `json:"csaleorderid,omitempty" xml:"csaleorderid,omitempty" comment:""`
			PkGroup          string  `json:"pk_group,omitempty" xml:"pk_group,omitempty" comment:"集团"`                 
			PkOrg            string  `json:"pk_org,omitempty" xml:"pk_org,omitempty" comment:"销售组织"`                     
			PkOrgV           string  `json:"pk_org_v,omitempty" xml:"pk_org_v,omitempty" comment:"销售组织版本"`               
			Ctrantypeid      string  `json:"ctrantypeid,omitempty" xml:"ctrantypeid,omitempty" comment:"订单类型"`           
			Vtrantypecode    string  `json:"vtrantypecode,omitempty" xml:"vtrantypecode,omitempty" comment:"订单类型编码"`       
			Cbiztypeid       string  `json:"cbiztypeid,omitempty" xml:"cbiztypeid,omitempty" comment:"业务流程"`             
			Vbillcode        string  `json:"vbillcode,omitempty" xml:"vbillcode,omitempty" comment:"单据号"`               
			Dbilldate        string  `json:"dbilldate,omitempty" xml:"dbilldate,omitempty" comment:"单据日期"`               
			Ccustomerid      string  `json:"ccustomerid,omitempty" xml:"ccustomerid,omitempty" comment:"客户"`           
			Bfreecustflag    string  `json:"bfreecustflag,omitempty" xml:"bfreecustflag,omitempty" comment:"是否散户"`       
			Cfreecustid      string  `json:"cfreecustid,omitempty" xml:"cfreecustid,omitempty" comment:"散户"`           
			Cdeptvid         string  `json:"cdeptvid,omitempty" xml:"cdeptvid,omitempty" comment:"部门"`                 
			Cdeptid          string  `json:"cdeptid,omitempty" xml:"cdeptid,omitempty" comment:"部门"`                   
			Cemployeeid      string  `json:"cemployeeid,omitempty" xml:"cemployeeid,omitempty" comment:"业务员"`           
			Corigcurrencyid  string  `json:"corigcurrencyid,omitempty" xml:"corigcurrencyid,omitempty" comment:"原币"`   
			Cinvoicecustid   string  `json:"cinvoicecustid,omitempty" xml:"cinvoicecustid,omitempty" comment:"开票客户"`     
			Ccustbankid      string  `json:"ccustbankid,omitempty" xml:"ccustbankid,omitempty" comment:"开户银行"`           
			Ccustbankaccid   string  `json:"ccustbankaccid,omitempty" xml:"ccustbankaccid,omitempty" comment:"开户银行账户"`     
			Cchanneltypeid   string  `json:"cchanneltypeid,omitempty" xml:"cchanneltypeid,omitempty" comment:"销售渠道类型"`     
			Cpaytermid       string  `json:"cpaytermid,omitempty" xml:"cpaytermid,omitempty" comment:"收款协议"`             
			Ctransporttypeid string  `json:"ctransporttypeid,omitempty" xml:"ctransporttypeid,omitempty" comment:"运输方式"` 
			Ctradewordid     string  `json:"ctradewordid,omitempty" xml:"ctradewordid,omitempty" comment:"贸易术语"`         
			Ndiscountrate    float64 `json:"ndiscountrate,omitempty" xml:"ndiscountrate,omitempty" comment:"整单折扣"`       
			Vcreditnum       string  `json:"vcreditnum,omitempty" xml:"vcreditnum,omitempty" comment:"信用证号"`             
			Cbalancetypeid   string  `json:"cbalancetypeid,omitempty" xml:"cbalancetypeid,omitempty" comment:"结算方式"`                          
			Badvfeeflag      string  `json:"badvfeeflag,omitempty" xml:"badvfeeflag,omitempty" comment:"代垫运费"`          
			Npreceiverate    float64 `json:"npreceiverate,omitempty" xml:"npreceiverate,omitempty" comment:"订单收款比例"`       
			Npreceivequota   float64 `json:"npreceivequota,omitempty" xml:"npreceivequota,omitempty" comment:"订单收款限额"`     
			Bpreceiveflag    string  `json:"bpreceiveflag,omitempty" xml:"bpreceiveflag,omitempty" comment:"收款限额控制预收"`       
			Npreceivemny     float64 `json:"npreceivemny,omitempty" xml:"npreceivemny,omitempty" comment:"实际预收款金额"`         
			Nreceivedmny     float64 `json:"nreceivedmny,omitempty" xml:"nreceivedmny,omitempty" comment:"实际收款金额"`         
			Nthisreceivemny  float64 `json:"nthisreceivemny,omitempty" xml:"nthisreceivemny,omitempty" comment:"本次收款金额"`   
			Ntotalnum        float64 `json:"ntotalnum,omitempty" xml:"ntotalnum,omitempty" comment:"总数量"`               
			Ntotalweight     float64 `json:"ntotalweight,omitempty" xml:"ntotalweight,omitempty" comment:"合计重量"`         
			Ntotalvolume     float64 `json:"ntotalvolume,omitempty" xml:"ntotalvolume,omitempty" comment:"合计体积"`         
			Ntotalpiece      float64 `json:"ntotalpiece,omitempty" xml:"ntotalpiece,omitempty" comment:"总件数"`           
			Ntotalorigmny    float64 `json:"ntotalorigmny,omitempty" xml:"ntotalorigmny,omitempty" comment:"价税合计"`       
			Ntotalmny        float64 `json:"ntotalmny,omitempty" xml:"ntotalmny,omitempty" comment:"冲抵前金额"`               
			Ntotalorigsubmny float64 `json:"ntotalorigsubmny,omitempty" xml:"ntotalorigsubmny,omitempty" comment:"费用冲抵金额"` 
			Boffsetflag      string  `json:"boffsetflag,omitempty" xml:"boffsetflag,omitempty" comment:"是否冲抵"`           
			Vcooppohcode     string  `json:"vcooppohcode,omitempty" xml:"vcooppohcode,omitempty" comment:"对方订单号"`         
			Bpocooptomeflag  string  `json:"bpocooptomeflag,omitempty" xml:"bpocooptomeflag,omitempty" comment:"由采购订单协同生成"`   
			DestPkOrg        string  `json:"dest_pk_org,omitempty" xml:"dest_pk_org,omitempty" comment:""`
			Bcooptopoflag    string  `json:"bcooptopoflag,omitempty" xml:"bcooptopoflag,omitempty" comment:"已协同生成采购订单"` 
			Fstatusflag      int     `json:"fstatusflag,omitempty" xml:"fstatusflag,omitempty" comment:"单据状态"`     
			Fpfstatusflag    string  `json:"fpfstatusflag,omitempty" xml:"fpfstatusflag,omitempty" comment:"审批流状态"` 
			Vnote            string  `json:"vnote,omitempty" xml:"vnote,omitempty" comment:""`
			Vdef1            string  `json:"vdef1,omitempty" xml:"vdef1,omitempty" comment:""`
			Vdef2            string  `json:"vdef2,omitempty" xml:"vdef2,omitempty" comment:""`
			Vdef3            string  `json:"vdef3,omitempty" xml:"vdef3,omitempty" comment:""`
			Vdef4            string  `json:"vdef4,omitempty" xml:"vdef4,omitempty" comment:""`
			Vdef5            string  `json:"vdef5,omitempty" xml:"vdef5,omitempty" comment:""`
			Vdef6            string  `json:"vdef6,omitempty" xml:"vdef6,omitempty" comment:""`
			Vdef7            string  `json:"vdef7,omitempty" xml:"vdef7,omitempty" comment:""`
			Vdef8            string  `json:"vdef8,omitempty" xml:"vdef8,omitempty" comment:""`
			Vdef9            string  `json:"vdef9,omitempty" xml:"vdef9,omitempty" comment:""`
			Vdef10           string  `json:"vdef10,omitempty" xml:"vdef10,omitempty" comment:""`
			Vdef11           string  `json:"vdef11,omitempty" xml:"vdef11,omitempty" comment:""`
			Vdef12           string  `json:"vdef12,omitempty" xml:"vdef12,omitempty" comment:""`
			Vdef13           string  `json:"vdef13,omitempty" xml:"vdef13,omitempty" comment:""`
			Vdef14           string  `json:"vdef14,omitempty" xml:"vdef14,omitempty" comment:"销售分类 b2b,b2c"` 
			Vdef15           string  `json:"vdef15,omitempty" xml:"vdef15,omitempty" comment:""`
			Vdef16           string  `json:"vdef16,omitempty" xml:"vdef16,omitempty" comment:""`
			Vdef17           string  `json:"vdef17,omitempty" xml:"vdef17,omitempty" comment:""`
			Vdef18           string  `json:"vdef18,omitempty" xml:"vdef18,omitempty" comment:""`
			Vdef19           string  `json:"vdef19,omitempty" xml:"vdef19,omitempty" comment:""`
			Vdef20           string  `json:"vdef20,omitempty" xml:"vdef20,omitempty" comment:""`
			Billmaker        string  `json:"billmaker,omitempty" xml:"billmaker,omitempty" comment:"制单人"`             
			Dmakedate        string  `json:"dmakedate,omitempty" xml:"dmakedate,omitempty" comment:"制单日期"`            
			Creator          string  `json:"creator,omitempty" xml:"creator,omitempty" comment:"创建人"`                 
			Creationtime     string  `json:"creationtime,omitempty" xml:"creationtime,omitempty" comment:"创建时间"`       
			Modifier         string  `json:"modifier,omitempty" xml:"modifier,omitempty" comment:"修改人"`               
			Modifiedtime     string  `json:"modifiedtime,omitempty" xml:"modifiedtime,omitempty" comment:"修改时间"`        
			Approver         string  `json:"approver,omitempty" xml:"approver,omitempty" comment:"审批人"`               
			Taudittime       string  `json:"taudittime,omitempty" xml:"taudittime,omitempty" comment:"审核日期"`           
			Creviserid       string  `json:"creviserid,omitempty" xml:"creviserid,omitempty" comment:"修订人"`           
			Iversion         int     `json:"iversion,omitempty" xml:"iversion,omitempty" comment:"修订版本号"`               
			Vrevisereason    string  `json:"vrevisereason,omitempty" xml:"vrevisereason,omitempty" comment:"修订理由"`     
			Trevisetime      string  `json:"trevisetime,omitempty" xml:"trevisetime,omitempty" comment:"修订时间"`         
			Bsendendflag     string  `json:"bsendendflag,omitempty" xml:"bsendendflag,omitempty" comment:"发货关闭"`       
			Boutendflag      string  `json:"boutendflag,omitempty" xml:"boutendflag,omitempty" comment:"出库关闭"`         
			Binvoicendflag   string  `json:"binvoicendflag,omitempty" xml:"binvoicendflag,omitempty" comment:"开票关闭"`   
			Bcostsettleflag  string  `json:"bcostsettleflag,omitempty" xml:"bcostsettleflag,omitempty" comment:"成本结算关闭"` 
			Barsettleflag    string  `json:"barsettleflag,omitempty" xml:"barsettleflag,omitempty" comment:"收入结算关闭"`     
			Iprintcount      int     `json:"iprintcount,omitempty" xml:"iprintcount,omitempty" comment:"打印次数"`  
			Nlrgtotalorigmny float64 `json:"nlrgtotalorigmny,omitempty" xml:"nlrgtotalorigmny,omitempty" comment:""`       
			SoSaleorderB     struct {
				Items []SaleItem `json:"item" xml:"item"`
			} `json:"so_saleorder_b" xml:"so_saleorder_b"`
}

type ReqSaleOrder struct {
	XMLName xml.Name `xml:"ufinterface"`
	base.Base
	Bill struct {
		Billhead ReqSaleOrderBillHead  `json:"billhead" xml:"billhead"`
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
	Csaleorderbid      string  `json:"csaleorderbid,omitempty" xml:"csaleorderbid,omitempty" comment:""`
	PkGroup            string  `json:"pk_group,omitempty" xml:"pk_group,omitempty" comment:"集团"`                     
	PkOrg              string  `json:"pk_org,omitempty" xml:"pk_org,omitempty" comment:"销售组织"`                         
	Dbilldate          string  `json:"dbilldate,omitempty" xml:"dbilldate,omitempty" comment:"单据日期"`                  
	Crowno             string  `json:"crowno,omitempty" xml:"crowno,omitempty" comment:"行号"`                         
	Cmaterialvid       string  `json:"cmaterialvid,omitempty" xml:"cmaterialvid,omitempty" comment:"物料编码"`             
	Cmaterialid        string  `json:"cmaterialid,omitempty" xml:"cmaterialid,omitempty" comment:"物料"`               
	Cvendorid          string  `json:"cvendorid,omitempty" xml:"cvendorid,omitempty" comment:"供应商"`                   
	Cprojectid         string  `json:"cprojectid,omitempty" xml:"cprojectid,omitempty" comment:"项目"`                 
	Cproductorid       string  `json:"cproductorid,omitempty" xml:"cproductorid,omitempty" comment:"生产厂商"`             
	Cfactoryid         string  `json:"cfactoryid,omitempty" xml:"cfactoryid,omitempty" comment:"工厂"`                 
	Cqualitylevelid    string  `json:"cqualitylevelid,omitempty" xml:"cqualitylevelid,omitempty" comment:"质量等级"`      
	Corigcountryid     string  `json:"corigcountryid,omitempty" xml:"corigcountryid,omitempty" comment:"原产国"`         
	Corigareaid        string  `json:"corigareaid,omitempty" xml:"corigareaid,omitempty" comment:"原产地区"`               
	Cunitid            string  `json:"cunitid,omitempty" xml:"cunitid,omitempty" comment:"主单位"`                       
	Castunitid         string  `json:"castunitid,omitempty" xml:"castunitid,omitempty" comment:"单位"`                 
	Nnum               int     `json:"nnum,omitempty" xml:"nnum,omitempty" comment:"主数量"`                             
	Nastnum            int     `json:"nastnum,omitempty" xml:"nastnum,omitempty" comment:"数量"`                       
	Vchangerate        string  `json:"vchangerate,omitempty" xml:"vchangerate,omitempty" comment:"换算率"`               
	Cqtunitid          string  `json:"cqtunitid,omitempty" xml:"cqtunitid,omitempty" comment:"报价单位"`                   
	Nqtunitnum         float64 `json:"nqtunitnum,omitempty" xml:"nqtunitnum,omitempty" comment:"报价单位数量"`                 
	Vqtunitrate        string  `json:"vqtunitrate,omitempty" xml:"vqtunitrate,omitempty" comment:"报价换算率"`               
	Ndiscountrate      float64 `json:"ndiscountrate,omitempty" xml:"ndiscountrate,omitempty" comment:"整单折扣"`           
	Nitemdiscountrate  float64 `json:"nitemdiscountrate,omitempty" xml:"nitemdiscountrate,omitempty" comment:"单品折扣"`   
	Ctaxcodeid         string  `json:"ctaxcodeid,omitempty" xml:"ctaxcodeid,omitempty" comment:"税码"`                 
	Ntaxrate           float64 `json:"ntaxrate,omitempty" xml:"ntaxrate" comment:"税率"`                     
	Ftaxtypeflag       int     `json:"ftaxtypeflag,omitempty" xml:"ftaxtypeflag,omitempty" comment:"扣税类别"`             
	Ccurrencyid        string  `json:"ccurrencyid,omitempty" xml:"ccurrencyid,omitempty" comment:"本位币"`              
	Nexchangerate      float64 `json:"nexchangerate,omitempty" xml:"nexchangerate,omitempty" comment:"折本汇率"`           
	Nqtorigtaxprice    float64 `json:"nqtorigtaxprice,omitempty" xml:"nqtorigtaxprice,omitempty" comment:"含税单价"`     
	Nqtorigprice       float64 `json:"nqtorigprice,omitempty" xml:"nqtorigprice,omitempty" comment:"无税单价"`             
	Nqtorigtaxnetprc   float64 `json:"nqtorigtaxnetprc,omitempty" xml:"nqtorigtaxnetprc,omitempty" comment:"含税净价"`  
	Nqtorignetprice    float64 `json:"nqtorignetprice,omitempty" xml:"nqtorignetprice,omitempty" comment:"无税净价"`       
	Norigtaxprice      float64 `json:"norigtaxprice,omitempty" xml:"norigtaxprice,omitempty" comment:"主含税单价"`          
	Norigprice         float64 `json:"norigprice,omitempty" xml:"norigprice,omitempty" comment:"主无税单价"`                 
	Norigtaxnetprice   float64 `json:"norigtaxnetprice,omitempty" xml:"norigtaxnetprice,omitempty" comment:"主含税净价"`     
	Norignetprice      float64 `json:"norignetprice,omitempty" xml:"norignetprice,omitempty" comment:"主无税净价"`           
	Ntax               float64 `json:"ntax,omitempty" xml:"ntax,omitempty" comment:"税额"`                             
	Ncaltaxmny         float64 `json:"ncaltaxmny" xml:"ncaltaxmny,omitempty" comment:"计税金额"`                
	Norigmny           float64 `json:"norigmny,omitempty" xml:"norigmny,omitempty" comment:"无税金额"`                     
	Norigtaxmny        float64 `json:"norigtaxmny,omitempty" xml:"norigtaxmny,omitempty" comment:"价税合计"`               
	Norigdiscount      float64 `json:"norigdiscount,omitempty" xml:"norigdiscount,omitempty" comment:"折扣额"`           
	Nbforigsubmny      float64 `json:"nbforigsubmny,omitempty" xml:"nbforigsubmny,omitempty" comment:"冲抵前金额"`           
	Nqttaxprice        float64 `json:"nqttaxprice,omitempty" xml:"nqttaxprice,omitempty" comment:"本币含税单价"`               
	Nqtprice           float64 `json:"nqtprice,omitempty" xml:"nqtprice,omitempty" comment:"本币无税单价"`                     
	Nqttaxnetprice     float64 `json:"nqttaxnetprice,omitempty" xml:"nqttaxnetprice,omitempty" comment:"本币含税净价"`         
	Nqtnetprice        float64 `json:"nqtnetprice,omitempty" xml:"nqtnetprice,omitempty" comment:"本币无税净价"`               
	Ntaxprice          float64 `json:"ntaxprice,omitempty" xml:"ntaxprice,omitempty" comment:"主本币含税单价"`                   
	Nprice             float64 `json:"nprice,omitempty" xml:"nprice,omitempty" comment:"主本币无税单价"`                         
	Ntaxnetprice       float64 `json:"ntaxnetprice,omitempty" xml:"ntaxnetprice,omitempty" comment:"主本币含税净价"`             
	Nnetprice          float64 `json:"nnetprice,omitempty" xml:"nnetprice,omitempty" comment:"主本币无税净价"`                  
	Nmny               float64 `json:"nmny,omitempty" xml:"nmny,omitempty" comment:"本币无税金额"`                             
	Ntaxmny            float64 `json:"ntaxmny,omitempty" xml:"ntaxmny,omitempty" comment:"本币价税合计"`                       
	Ndiscount          float64 `json:"ndiscount,omitempty" xml:"ndiscount,omitempty" comment:"本币折扣额"`                   
	Ngroupexchgrate    float64 `json:"ngroupexchgrate,omitempty" xml:"ngroupexchgrate,omitempty" comment:"集团本位币汇率"`     
	Ngroupmny          float64 `json:"ngroupmny,omitempty" xml:"ngroupmny,omitempty" comment:"集团本币无税金额"`                   
	Ngrouptaxmny       float64 `json:"ngrouptaxmny,omitempty" xml:"ngrouptaxmny,omitempty" comment:"集团本币价税合计"`             
	Nglobalexchgrate   float64 `json:"nglobalexchgrate,omitempty" xml:"nglobalexchgrate,omitempty" comment:"全局本位币汇率"`     
	Nglobalmny         float64 `json:"nglobalmny,omitempty" xml:"nglobalmny,omitempty" comment:"全局本币无税金额"`              
	Nglobaltaxmny      float64 `json:"nglobaltaxmny,omitempty" xml:"nglobaltaxmny,omitempty" comment:"全局本币价税合计"`          
	Naskqtorigtaxprc   float64 `json:"naskqtorigtaxprc,omitempty" xml:"naskqtorigtaxprc,omitempty" comment:"询价原币含税单价"`     
	Naskqtorigprice    float64 `json:"naskqtorigprice,omitempty" xml:"naskqtorigprice,omitempty" comment:"询价原币无税单价"`       
	Naskqtorigtxntprc  float64 `json:"naskqtorigtxntprc,omitempty" xml:"naskqtorigtxntprc,omitempty" comment:"询价原币含税净价"`   
	Naskqtorignetprice float64 `json:"naskqtorignetprice,omitempty" xml:"naskqtorignetprice,omitempty" comment:"询价原币无税净价"` 
	Nweight            float64 `json:"nweight,omitempty" xml:"nweight,omitempty" comment:"重量"`                       
	Nvolume            float64 `json:"nvolume,omitempty" xml:"nvolume,omitempty" comment:"体积"`                        
	Npiece             float64 `json:"npiece,omitempty" xml:"npiece,omitempty" comment:"件数"`                         
	Cprodlineid        string  `json:"cprodlineid,omitempty" xml:"cprodlineid,omitempty" comment:"产品线"`               
	Vbatchcode         string  `json:"vbatchcode,omitempty" xml:"vbatchcode,omitempty" comment:"批次号"`                 
	PkBatchcode        string  `json:"pk_batchcode,omitempty" xml:"pk_batchcode,omitempty" comment:""`
	Blaborflag         string  `json:"blaborflag,omitempty" xml:"blaborflag,omitempty" comment:"服务类"`               
	Bdiscountflag      string  `json:"bdiscountflag,omitempty" xml:"bdiscountflag,omitempty" comment:"折扣类"`         
	Blargessflag       string  `json:"blargessflag,omitempty" xml:"blargessflag" comment:"赠品"`           
	Bbindflag          string  `json:"bbindflag,omitempty" xml:"bbindflag,omitempty" comment:"捆绑存货"`                 
	Dsenddate          string  `json:"dsenddate,omitempty" xml:"dsenddate,omitempty" comment:"要求发货日期"`                 
	Dreceivedate       string  `json:"dreceivedate,omitempty" xml:"dreceivedate,omitempty" comment:"计划到货日期"`           
	Creceivecustid     string  `json:"creceivecustid,omitempty" xml:"creceivecustid,omitempty" comment:"收货客户"`     
	Creceiveareaid     string  `json:"creceiveareaid,omitempty" xml:"creceiveareaid,omitempty" comment:"收货地区"`       
	Creceiveaddrid     string  `json:"creceiveaddrid,omitempty" xml:"creceiveaddrid,omitempty" comment:"收货地址"`       
	Creceiveadddocid   string  `json:"creceiveadddocid,omitempty" xml:"creceiveadddocid,omitempty" comment:"收货地点"`  
	Csendstockorgvid   string  `json:"csendstockorgvid,omitempty" xml:"csendstockorgvid,omitempty" comment:"发货库存组织"`   
	Csendstockorgid    string  `json:"csendstockorgid,omitempty" xml:"csendstockorgid,omitempty" comment:"发货库存组织最新版本"`     
	Csendstordocid     string  `json:"csendstordocid,omitempty" xml:"csendstordocid,omitempty" comment:"发货仓库"`       
	Ctrafficorgvid     string  `json:"ctrafficorgvid,omitempty" xml:"ctrafficorgvid,omitempty" comment:"物流组织"`       
	Ctrafficorgid      string  `json:"ctrafficorgid,omitempty" xml:"ctrafficorgid,omitempty" comment:"物流组织"`         
	Csettleorgvid      string  `json:"csettleorgvid,omitempty" xml:"csettleorgvid,omitempty" comment:"结算财务组织"`         
	Csettleorgid       string  `json:"csettleorgid,omitempty" xml:"csettleorgid,omitempty" comment:"结算财务组织"`           
	Crececountryid     string  `json:"crececountryid,omitempty" xml:"crececountryid,omitempty" comment:"收货国家/地区"`       
	Csendcountryid     string  `json:"csendcountryid,omitempty" xml:"csendcountryid,omitempty" comment:"发货国家/地区"`       
	Ctaxcountryid      string  `json:"ctaxcountryid,omitempty" xml:"ctaxcountryid,omitempty" comment:"报税国家/地区"`         
	Fbuysellflag       int     `json:"fbuysellflag,omitempty" xml:"fbuysellflag,omitempty" comment:"购销类型"`           
	Btriatradeflag     string  `json:"btriatradeflag,omitempty" xml:"btriatradeflag,omitempty" comment:"三角贸易"`       
	Carorgvid          string  `json:"carorgvid,omitempty" xml:"carorgvid,omitempty" comment:"应收组织"`               
	Carorgid           string  `json:"carorgid,omitempty" xml:"carorgid,omitempty" comment:"应收组织最新版本"`                   
	Cprofitcentervid   string  `json:"cprofitcentervid,omitempty" xml:"cprofitcentervid,omitempty" comment:"利润中心"`   
	Cprofitcenterid    string  `json:"cprofitcenterid,omitempty" xml:"cprofitcenterid,omitempty" comment:"利润中心"`     
	Vbrevisereason     string  `json:"vbrevisereason,omitempty" xml:"vbrevisereason,omitempty" comment:"修订理由"`       
	Frowstatus         int     `json:"frowstatus,omitempty" xml:"frowstatus,omitempty" comment:"行状态"`               
	Vrownote           string  `json:"vrownote,omitempty" xml:"vrownote,omitempty" comment:"行备注"`                   
	Cpriceformid       string  `json:"cpriceformid,omitempty" xml:"cpriceformid,omitempty" comment:"价格组成"`           
	Cpricepolicyid     string  `json:"cpricepolicyid,omitempty" xml:"cpricepolicyid,omitempty" comment:"价格政策"`       
	Cpriceitemid       string  `json:"cpriceitemid,omitempty" xml:"cpriceitemid,omitempty" comment:"价格项目"`           
	Cpriceitemtableid  string  `json:"cpriceitemtableid,omitempty" xml:"cpriceitemtableid,omitempty" comment:"价目表"` 
	Bjczxsflag         string  `json:"bjczxsflag,omitempty" xml:"bjczxsflag,omitempty" comment:"借出转销售"`               
	Vcttype            string  `json:"vcttype,omitempty" xml:"vcttype,omitempty" comment:"合同类型"`                     
	Vctcode            string  `json:"vctcode,omitempty" xml:"vctcode,omitempty" comment:"销售合同号"`                     
	Cctmanageid        string  `json:"cctmanageid,omitempty" xml:"cctmanageid,omitempty" comment:"合同主表"`             
	Ctmanagebid        string  `json:"ctmanagebid,omitempty" xml:"ctmanagebid,omitempty" comment:""`             
	Vsrctype           string  `json:"vsrctype,omitempty" xml:"vsrctype,omitempty" comment:"来源单据类型"`                   
	Vsrctrantype       string  `json:"vsrctrantype,omitempty" xml:"vsrctrantype,omitempty" comment:"来源交易类型"`           
	Vsrccode           string  `json:"vsrccode,omitempty" xml:"vsrccode,omitempty" comment:"来源单据号"`                   
	Vsrcrowno          string  `json:"vsrcrowno,omitempty" xml:"vsrcrowno,omitempty" comment:"来源单据行号"`                
	Csrcid             string  `json:"csrcid,omitempty" xml:"csrcid,omitempty" comment:""`                      
	Csrcbid            string  `json:"csrcbid,omitempty" xml:"csrcbid,omitempty" comment:""`                     
	Vfirsttype         string  `json:"vfirsttype,omitempty" xml:"vfirsttype,omitempty" comment:"源头单据类型"`               
	Vfirsttrantype     string  `json:"vfirsttrantype,omitempty" xml:"vfirsttrantype,omitempty" comment:"源头交易类型"`       
	Vfirstcode         string  `json:"vfirstcode,omitempty" xml:"vfirstcode,omitempty" comment:"源头单据号"`               
	Cfirstid           string  `json:"cfirstid,omitempty" xml:"cfirstid,omitempty" comment:""`                   
	Cfirstbid          string  `json:"cfirstbid,omitempty" xml:"cfirstbid,omitempty" comment:""`                
	Vfirstrowno        string  `json:"vfirstrowno,omitempty" xml:"vfirstrowno,omitempty" comment:"源头单据行号"`             
	Cretreasonid       string  `json:"cretreasonid,omitempty" xml:"cretreasonid,omitempty" comment:"退货原因"`           
	Cretpolicyid       string  `json:"cretpolicyid,omitempty" xml:"cretpolicyid,omitempty" comment:"退货政策"`           
	Vreturnmode        string  `json:"vreturnmode,omitempty" xml:"vreturnmode,omitempty" comment:"退货责任处理方式"`            
	Fretexchange       int     `json:"fretexchange,omitempty" xml:"fretexchange,omitempty" comment:"退换货标记"`           
	Cexchangesrcretid  string  `json:"cexchangesrcretid,omitempty" xml:"cexchangesrcretid,omitempty" comment:""`
	Clargesssrcid      string  `json:"clargesssrcid,omitempty" xml:"clargesssrcid,omitempty" comment:""`       
	Cbindsrcid         string  `json:"cbindsrcid,omitempty" xml:"cbindsrcid,omitempty" comment:""`             
	Flargesstypeflag   string  `json:"flargesstypeflag,omitempty" xml:"flargesstypeflag,omitempty" comment:"赠品价格分摊方式"` 
	Nlargessmny        float64 `json:"nlargessmny,omitempty" xml:"nlargessmny,omitempty" comment:"赠品价格分摊前无税金额"`           
	Nlargesstaxmny     string  `json:"nlargesstaxmny,omitempty" xml:"nlargesstaxmny,omitempty" comment:"赠品价格分摊前价税合计"`     
	Bprerowcloseflag   string  `json:"bprerowcloseflag" xml:"bprerowcloseflag,omitempty" comment:"预订单行关闭"` 
	Vfree1             string  `json:"vfree1,omitempty" xml:"vfree1,omitempty" comment:""`
	Vfree2             string  `json:"vfree2,omitempty" xml:"vfree2,omitempty" comment:""`
	Vfree3             string  `json:"vfree3,omitempty" xml:"vfree3,omitempty" comment:""`
	Vfree4             string  `json:"vfree4,omitempty" xml:"vfree4,omitempty" comment:""`
	Vfree5             string  `json:"vfree5,omitempty" xml:"vfree5,omitempty" comment:""`
	Vfree6             string  `json:"vfree6,omitempty" xml:"vfree6,omitempty" comment:""`
	Vfree7             string  `json:"vfree7,omitempty" xml:"vfree7,omitempty" comment:""`
	Vfree8             string  `json:"vfree8,omitempty" xml:"vfree8,omitempty" comment:""`
	Vfree9             string  `json:"vfree9,omitempty" xml:"vfree9,omitempty" comment:""`
	Vfree10            string  `json:"vfree10,omitempty" xml:"vfree10,omitempty" comment:""`
	Vfree11            string  `json:"vfree11,omitempty" xml:"vfree11,omitempty" comment:""`
	Vfree12            string  `json:"vfree12,omitempty" xml:"vfree12,omitempty" comment:""`
	Vfree13            string  `json:"vfree13,omitempty" xml:"vfree13,omitempty" comment:""`
	Vfree14            string  `json:"vfree14,omitempty" xml:"vfree14,omitempty" comment:""`
	Vfree15            string  `json:"vfree15,omitempty" xml:"vfree15,omitempty" comment:""`
	Vfree16            string  `json:"vfree16,omitempty" xml:"vfree16,omitempty" comment:""`
	Vfree17            string  `json:"vfree17,omitempty" xml:"vfree17,omitempty" comment:""`
	Vfree18            string  `json:"vfree18,omitempty" xml:"vfree18,omitempty" comment:""`
	Vfree19            string  `json:"vfree19,omitempty" xml:"vfree19,omitempty" comment:""`
	Vfree20            string  `json:"vfree20,omitempty" xml:"vfree20,omitempty" comment:""`
	Bbsendendflag      string  `json:"bbsendendflag,omitempty" xml:"bbsendendflag,omitempty" comment:"发货关闭"`         
	Bbinvoicendflag    string  `json:"bbinvoicendflag,omitempty" xml:"bbinvoicendflag,omitempty" comment:"开票关闭"`     
	Bboutendflag       string  `json:"bboutendflag,omitempty" xml:"bboutendflag,omitempty" comment:"出库关闭"`           
	Bbcostsettleflag   string  `json:"bbcostsettleflag,omitempty" xml:"bbcostsettleflag,omitempty" comment:"成本结算关闭"`   
	Bbarsettleflag     string  `json:"bbarsettleflag,omitempty" xml:"bbarsettleflag,omitempty" comment:"收入结算关闭"`       
	Vclosereason       string  `json:"vclosereason,omitempty" xml:"vclosereason,omitempty" comment:"关闭/打开原因"`           
	Nsendauditnum      float64 `json:"nsendauditnum,omitempty" xml:"nsendauditnum,omitempty" comment:"发货审批主数量"`         
	Noutauditnum       float64 `json:"noutauditnum,omitempty" xml:"noutauditnum,omitempty" comment:"出库审批主数量"`           
	Ninvoiceauditnum   float64 `json:"ninvoiceauditnum,omitempty" xml:"ninvoiceauditnum,omitempty" comment:"发票审批主数量"`   
	Noutnotauditnum    float64 `json:"noutnotauditnum,omitempty" xml:"noutnotauditnum,omitempty" comment:"出库未签字主数量"`     
	Nlossnotauditnum   float64 `json:"nlossnotauditnum,omitempty" xml:"nlossnotauditnum,omitempty" comment:"途损单未审核主数量"`   
	Ntotalsendnum      float64 `json:"ntotalsendnum,omitempty" xml:"ntotalsendnum,omitempty" comment:"累计发货主数量"`         
	Ntotalinvoicenum   float64 `json:"ntotalinvoicenum,omitempty" xml:"ntotalinvoicenum,omitempty" comment:"累计开票主数量"`   
	Ntotaloutnum       float64 `json:"ntotaloutnum,omitempty" xml:"ntotaloutnum,omitempty" comment:"累计出库主数量"`           
	Ntotalnotoutnum    float64 `json:"ntotalnotoutnum,omitempty" xml:"ntotalnotoutnum,omitempty" comment:"累计应发未出库主数量"`     
	Ntotalsignnum      float64 `json:"ntotalsignnum,omitempty" xml:"ntotalsignnum,omitempty" comment:"累计签收主数量"`         
	Ntranslossnum      float64 `json:"ntranslossnum,omitempty" xml:"ntranslossnum,omitempty" comment:"累计途损主数量"`         
	Ntotalrushnum      float64 `json:"ntotalrushnum,omitempty" xml:"ntotalrushnum,omitempty" comment:"累计出库对冲主数量"`      
	Ntotalestarnum     float64 `json:"ntotalestarnum,omitempty" xml:"ntotalestarnum,omitempty" comment:"累计暂估应收主数量"`       
	Ntotalarnum        float64 `json:"ntotalarnum,omitempty" xml:"ntotalarnum,omitempty" comment:"累计确认应收主数量"`             
	Ntotalcostnum      float64 `json:"ntotalcostnum,omitempty" xml:"ntotalcostnum,omitempty" comment:"累计成本结算主数量"`         
	Ntotalestarmny     float64 `json:"ntotalestarmny,omitempty" xml:"ntotalestarmny,omitempty" comment:"累计暂估应收金额"`       
	Ntotalarmny        float64 `json:"ntotalarmny,omitempty" xml:"ntotalarmny,omitempty" comment:"累计确认应收金额"`             
	Ntotalpaymny       float64 `json:"ntotalpaymny,omitempty" xml:"ntotalpaymny,omitempty" comment:"累计财务核销金额"`           
	Norigsubmny        float64 `json:"norigsubmny,omitempty" xml:"norigsubmny,omitempty" comment:"累计冲抵金额"`             
	Narrangescornum    float64 `json:"narrangescornum,omitempty" xml:"narrangescornum,omitempty" comment:"累计安排委外订单主数量"`     
	Narrangepoappnum   float64 `json:"narrangepoappnum,omitempty" xml:"narrangepoappnum,omitempty" comment:"累计安排请购单主数量"`   
	Narrangetoornum    float64 `json:"narrangetoornum,omitempty" xml:"narrangetoornum,omitempty" comment:"累计安排调拨订单主数量"`    
	Narrangetoappnum   float64 `json:"narrangetoappnum,omitempty" xml:"narrangetoappnum,omitempty" comment:"累计安排调拨申请主数量"`   
	Narrangemonum      float64 `json:"narrangemonum,omitempty" xml:"narrangemonum,omitempty" comment:"累计安排生产订单主数量"`         
	Narrangeponum      float64 `json:"narrangeponum,omitempty" xml:"narrangeponum,omitempty" comment:"累计安排采购订单主数量"`       
	Ntotalplonum       float64 `json:"ntotalplonum,omitempty" xml:"ntotalplonum,omitempty" comment:"累计生成计划订单主数量"`           
	Nreqrsnum          float64 `json:"nreqrsnum,omitempty" xml:"nreqrsnum,omitempty" comment:"预留主数量"`                 
	Ntotalreturnnum    float64 `json:"ntotalreturnnum,omitempty" xml:"ntotalreturnnum,omitempty" comment:"累计退货主数量"`     
	Ntotaltradenum     float64 `json:"ntotaltradenum,omitempty" xml:"ntotaltradenum,omitempty" comment:"累计发出商品主数量"`       
	Barrangedflag      string  `json:"barrangedflag,omitempty" xml:"barrangedflag,omitempty" comment:"是否货源安排完毕"`         
	Carrangepersonid   string  `json:"carrangepersonid,omitempty" xml:"carrangepersonid,omitempty" comment:""`   
	Tlastarrangetime   string  `json:"tlastarrangetime,omitempty" xml:"tlastarrangetime,omitempty" comment:""`   
	Nsendunfinisednum  float64 `json:"nsendunfinisednum,omitempty" xml:"nsendunfinisednum,omitempty" comment:"发货未完成量"` 
	Noutunfinisednum   float64 `json:"noutunfinisednum,omitempty" xml:"noutunfinisednum,omitempty" comment:"出库未完成量"`  
	Ninvunfinisednum   float64 `json:"ninvunfinisednum,omitempty" xml:"ninvunfinisednum,omitempty" comment:"发票未完成量"`   
	Nnotarnum          float64 `json:"nnotarnum,omitempty" xml:"nnotarnum,omitempty" comment:"未传确认应收主数量"`                
	Nnotcostnum        float64 `json:"nnotcostnum,omitempty" xml:"nnotcostnum,omitempty" comment:"未传存货核算主数量"`             
	Bbsettleendflag    string  `json:"bbsettleendflag,omitempty" xml:"bbsettleendflag,omitempty" comment:"结算关闭"`     
	Srcts              string  `json:"srcts,omitempty" xml:"srcts,omitempty" comment:"来源单据表头时间戳"`                         
	Srcbts             string  `json:"srcbts,omitempty" xml:"srcbts,omitempty" comment:"来源单据表体时间戳"`                       
	Srcorgid           string  `json:"srcorgid,omitempty" xml:"srcorgid,omitempty" comment:""`                 
	Vbdef1             string  `json:"vbdef1,omitempty" xml:"vbdef1,omitempty" comment:""`
	Vbdef2             string  `json:"vbdef2,omitempty" xml:"vbdef2,omitempty" comment:""`
	Vbdef3             string  `json:"vbdef3,omitempty" xml:"vbdef3,omitempty" comment:""`
	Vbdef4             string  `json:"vbdef4,omitempty" xml:"vbdef4,omitempty" comment:""`
	Vbdef5             string  `json:"vbdef5,omitempty" xml:"vbdef5,omitempty" comment:""`
	Vbdef6             string  `json:"vbdef6,omitempty" xml:"vbdef6,omitempty" comment:""`
	Vbdef7             string  `json:"vbdef7,omitempty" xml:"vbdef7,omitempty" comment:""`
	Vbdef8             string  `json:"vbdef8,omitempty" xml:"vbdef8,omitempty" comment:""`
	Vbdef9             string  `json:"vbdef9,omitempty" xml:"vbdef9,omitempty" comment:""`
	Vbdef10            string  `json:"vbdef10,omitempty" xml:"vbdef10,omitempty" comment:""`
	Vbdef11            string  `json:"vbdef11,omitempty" xml:"vbdef11,omitempty" comment:""`
	Vbdef12            string  `json:"vbdef12,omitempty" xml:"vbdef12,omitempty" comment:""`
	Vbdef13            string  `json:"vbdef13,omitempty" xml:"vbdef13,omitempty" comment:""`
	Vbdef14            string  `json:"vbdef14,omitempty" xml:"vbdef14,omitempty" comment:""`
	Vbdef15            string  `json:"vbdef15,omitempty" xml:"vbdef15,omitempty" comment:""`
	Vbdef16            string  `json:"vbdef16,omitempty" xml:"vbdef16,omitempty" comment:""`
	Vbdef17            string  `json:"vbdef17,omitempty" xml:"vbdef17,omitempty" comment:""`
	Vbdef18            string  `json:"vbdef18,omitempty" xml:"vbdef18,omitempty" comment:""`
	Vbdef19            string  `json:"vbdef19,omitempty" xml:"vbdef19,omitempty" comment:""`
	Vbdef20            string  `json:"vbdef20,omitempty" xml:"vbdef20,omitempty" comment:""`
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

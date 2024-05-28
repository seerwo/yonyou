package purchase

import (
	"encoding/xml"
	"fmt"
	"github.com/seerwo/yonyou/nc/base"
	"github.com/seerwo/yonyou/nc/context"
	"github.com/seerwo/yonyou/util"
)

type ReqPurchaseOrderBillHead struct {
			PkOrder           string  `json:"pk_order,omitempty" xml:"pk_order,omitempty" comment:""`
			Ctradewordid      string  `json:"ctradewordid,omitempty" xml:"ctradewordid,omitempty" comment:"贸易术语"`           
			PkGroup           string  `json:"pk_group,omitempty" xml:"pk_group,omitempty" comment:"集团"`                   
			PkOrg             string  `json:"pk_org,omitempty" xml:"pk_org,omitempty" comment:"销售组织"`                      
			PkOrgV            string  `json:"pk_org_v,omitempty" xml:"pk_org_v,omitempty" comment:"销售组织版本"`                   
			Vbillcode         string  `json:"vbillcode,omitempty" xml:"vbillcode,omitempty" comment:"单据号"`                 
			Dbilldate         string  `json:"dbilldate,omitempty" xml:"dbilldate,omitempty" comment:"单据日期"`                 
			PkFreecust        string  `json:"pk_freecust,omitempty" xml:"pk_freecust,omitempty" comment:"散户"`             
			PkSupplier        string  `json:"pk_supplier,omitempty" xml:"pk_supplier,omitempty" comment:"供应商"`             
			PkBankdoc         string  `json:"pk_bankdoc,omitempty" xml:"pk_bankdoc,omitempty" comment:"开户银行"`               
			PkDept            string  `json:"pk_dept,omitempty" xml:"pk_dept,omitempty" comment:"采购部门最新版本"`                     
			Cemployeeid       string  `json:"cemployeeid,omitempty" xml:"cemployeeid,omitempty" comment:"采购员"`             
			Vtrantypecode     string  `json:"vtrantypecode,omitempty" xml:"vtrantypecode,omitempty" comment:"订单类型编码"`         
			PkRecvcustomer    string  `json:"pk_recvcustomer,omitempty" xml:"pk_recvcustomer,omitempty" comment:"收货客户"`     
			PkInvcsupllier    string  `json:"pk_invcsupllier,omitempty" xml:"pk_invcsupllier,omitempty" comment:"开票供应商"`     
			PkDeliveradd      string  `json:"pk_deliveradd,omitempty" xml:"pk_deliveradd,omitempty" comment:"供应商发货地址"`         
			PkTransporttype   string  `json:"pk_transporttype,omitempty" xml:"pk_transporttype,omitempty" comment:"运输方式"`   
			PkPayterm         string  `json:"pk_payterm,omitempty" xml:"pk_payterm,omitempty" comment:"付款协议"`             
			Billmaker         string  `json:"billmaker,omitempty" xml:"billmaker,omitempty" comment:"制单人"`                 
			Forderstatus      string  `json:"forderstatus,omitempty" xml:"forderstatus,omitempty" comment:"单据状态"`           
			Approver          string  `json:"approver,omitempty" xml:"approver,omitempty" comment:"审批人"`                   
			Vmemo             string  `json:"vmemo,omitempty" xml:"vmemo,omitempty" comment:"备注"`                         
			Ccontracttextpath string  `json:"ccontracttextpath,omitempty" xml:"ccontracttextpath,omitempty" comment:"合同附件"`
			Norgprepaylimit   float64 `json:"norgprepaylimit,omitempty" xml:"norgprepaylimit,omitempty" comment:"预付款限额"`     
			Nversion          int     `json:"nversion,omitempty" xml:"nversion,omitempty" comment:"版本号"`                   
			Bislatest         string  `json:"bislatest,omitempty" xml:"bislatest,omitempty" comment:"最新版本"`                 
			Bisreplenish      string  `json:"bisreplenish,omitempty" xml:"bisreplenish,omitempty" comment:"补货"`           
			Breturn           string  `json:"breturn,omitempty" xml:"breturn,omitempty" comment:"退货"`                     
			Iprintcount       string  `json:"iprintcount,omitempty" xml:"iprintcount,omitempty" comment:"打印次数"`             
			Vnote             string  `json:"vnote,omitempty" xml:"vnote,omitempty" comment:""`
			Vdef1             string  `json:"vdef1,omitempty" xml:"vdef1,omitempty" comment:""`
			Vdef2             string  `json:"vdef2,omitempty" xml:"vdef2,omitempty" comment:""`
			Vdef3             string  `json:"vdef3,omitempty" xml:"vdef3,omitempty" comment:""`
			Vdef4             string  `json:"vdef4,omitempty" xml:"vdef4,omitempty" comment:""`
			Vdef5             string  `json:"vdef5,omitempty" xml:"vdef5,omitempty" comment:""`
			Vdef6             string  `json:"vdef6,omitempty" xml:"vdef6,omitempty" comment:""`
			Vdef7             string  `json:"vdef7,omitempty" xml:"vdef7,omitempty" comment:""`
			Vdef8             string  `json:"vdef8,omitempty" xml:"vdef8,omitempty" comment:""`
			Vdef9             string  `json:"vdef9,omitempty" xml:"vdef9,omitempty" comment:""`
			Vdef10            string  `json:"vdef10,omitempty" xml:"vdef10,omitempty" comment:""`
			Vdef11            string  `json:"vdef11,omitempty" xml:"vdef11,omitempty" comment:""`
			Vdef12            string  `json:"vdef12,omitempty" xml:"vdef12,omitempty" comment:""`
			Vdef13            string  `json:"vdef13,omitempty" xml:"vdef13,omitempty" comment:""`
			Vdef14            string  `json:"vdef14,omitempty" xml:"vdef14,omitempty" comment:""`
			Vdef15            string  `json:"vdef15,omitempty" xml:"vdef15,omitempty" comment:""`
			Vdef16            string  `json:"vdef16,omitempty" xml:"vdef16,omitempty" comment:""`
			Vdef17            string  `json:"vdef17,omitempty" xml:"vdef17,omitempty" comment:""`
			Vdef18            string  `json:"vdef18,omitempty" xml:"vdef18,omitempty" comment:""`
			Vdef19            string  `json:"vdef19,omitempty" xml:"vdef19,omitempty" comment:""`
			Vdef20            string  `json:"vdef20,omitempty" xml:"vdef20,omitempty" comment:""`
			Creationtime      string  `json:"creationtime,omitempty" xml:"creationtime,omitempty" comment:"制单时间"`       
			Taudittime        string  `json:"taudittime,omitempty" xml:"taudittime,omitempty" comment:"审批时间"`          
			Modifiedtime      string  `json:"modifiedtime,omitempty" xml:"modifiedtime,omitempty" comment:"最后修改时间"`       
			Crevisepsn        string  `json:"crevisepsn,omitempty" xml:"crevisepsn,omitempty" comment:"修订人"`           
			Trevisiontime     string  `json:"trevisiontime,omitempty" xml:"trevisiontime,omitempty" comment:"修订时间"`     
			Bcooptoso         string  `json:"bcooptoso,omitempty" xml:"bcooptoso,omitempty" comment:"已协同生成销售订单"`              
			Bsocooptome       string  `json:"bsocooptome,omitempty" xml:"bsocooptome,omitempty" comment:"由销售订单协同生成"`         
			Vcoopordercode    string  `json:"vcoopordercode,omitempty" xml:"vcoopordercode,omitempty" comment:"对方订单号"`   
			Ntotalastnum      float64 `json:"ntotalastnum,omitempty" xml:"ntotalastnum,omitempty" comment:"总数量"`       
			Ntotalorigmny     float64 `json:"ntotalorigmny,omitempty" xml:"ntotalorigmny,omitempty" comment:"价税合计"`     
			Modifier          string  `json:"modifier,omitempty" xml:"modifier,omitempty" comment:"最后修改人"`               
			PkFreezepsndoc    string  `json:"pk_freezepsndoc,omitempty" xml:"pk_freezepsndoc,omitempty" comment:"冻结人"` 
			Tfreezetime       string  `json:"tfreezetime,omitempty" xml:"tfreezetime,omitempty" comment:"冻结时间"`         
			Bfrozen           string  `json:"bfrozen,omitempty" xml:"bfrozen,omitempty" comment:"冻结"`                 
			Vfrozenreason     string  `json:"vfrozenreason,omitempty" xml:"vfrozenreason,omitempty" comment:"冻结原因"`     
			PkBusitype        string  `json:"pk_busitype,omitempty" xml:"pk_busitype,omitempty" comment:"业务流程"`         
			Fhtaxtypeflag     int     `json:"fhtaxtypeflag,omitempty" xml:"fhtaxtypeflag,omitempty" comment:"整单扣税类别"`     
			Nhtaxrate         float64 `json:"nhtaxrate,omitempty" xml:"nhtaxrate,omitempty" comment:"整单税率"`             
			Corigcurrencyid   string  `json:"corigcurrencyid,omitempty" xml:"corigcurrencyid,omitempty" comment:"币种"` 
			Brefwhenreturn    string  `json:"brefwhenreturn,omitempty" xml:"brefwhenreturn,omitempty" comment:"退货/库基于原订单补货"`   
			Ntotalweight      float64 `json:"ntotalweight,omitempty" xml:"ntotalweight,omitempty" comment:"总重量"`       
			Ntotalvolume      float64 `json:"ntotalvolume,omitempty" xml:"ntotalvolume,omitempty" comment:"总体积"`       
			Ntotalpiece       float64 `json:"ntotalpiece,omitempty" xml:"ntotalpiece,omitempty" comment:"总件数"`         
			Bfinalclose       string  `json:"bfinalclose,omitempty" xml:"bfinalclose,omitempty" comment:"最终关闭"`         
			Creator           string  `json:"creator,omitempty" xml:"creator,omitempty" comment:"创建人"`                 
			PkProject         string  `json:"pk_project,omitempty" xml:"pk_project,omitempty" comment:"项目"`           
			Dclosedate        string  `json:"dclosedate,omitempty" xml:"dclosedate,omitempty" comment:"最终关闭日期"`           
			PkDeptV           string  `json:"pk_dept_v,omitempty" xml:"pk_dept_v,omitempty" comment:"采购部门"`            
			Ctrantypeid       string  `json:"ctrantypeid,omitempty" xml:"ctrantypeid,omitempty" comment:"订单类型"`         
			Irespstatus       string  `json:"irespstatus,omitempty" xml:"irespstatus,omitempty" comment:"响应状态"`         
			Vreason           string  `json:"vreason,omitempty" xml:"vreason,omitempty" comment:"拒绝/变更理由"`                
			Bpublish          string  `json:"bpublish,omitempty" xml:"bpublish,omitempty" comment:"发布"`               
			PkPubpsn          string  `json:"pk_pubpsn,omitempty" xml:"pk_pubpsn,omitempty" comment:"发布人"`             
			Tpubtime          string  `json:"tpubtime,omitempty" xml:"tpubtime,omitempty" comment:"发布时间"`               
			PkResppsn         string  `json:"pk_resppsn,omitempty" xml:"pk_resppsn,omitempty" comment:"响应人"`          
			Tresptime         string  `json:"tresptime,omitempty" xml:"tresptime,omitempty" comment:"响应时间"`             
			PkOrderB          struct {
				Items []PurchaseItem `json:"item" xml:"item"`
			} `json:"pk_order_b" xml:"pk_order_b"`
}

type ReqPurchaseOrder struct {
	base.Base
	Bill struct {
		BillHead ReqPurchaseOrderBillHead `json:"billhead" xml:"billhead"`
	} `json:"bill" xml:"bill"`
}

type PurchaseItem struct {
	PkOrderB           string  `json:"pk_order_b,omitempty" xml:"pk_order_b,omitempty" comment:""`
	PkGroup            string  `json:"pk_group,omitempty" xml:"pk_group,omitempty" comment:"所属集团"`                   
	PkOrg              string  `json:"pk_org,omitempty" xml:"pk_org,omitempty" comment:"采购组织版本信息"`                       
	PkOrgV             string  `json:"pk_org_v,omitempty" xml:"pk_org_v,omitempty" comment:"采购组织"`                   
	PkReqcorp          string  `json:"pk_reqcorp,omitempty" comment:"需求公司"`                                
	PkReqstoorg        string  `json:"pk_reqstoorg,omitempty" xml:"pk_reqstoorg,omitempty" comment:"需求库存组织最新版本"`           
	PkReqstoorgV       string  `json:"pk_reqstoorg_v,omitempty" xml:"pk_reqstoorg_v,omitempty" comment:"需求库存组织"`       
	PkReqstordoc       string  `json:"pk_reqstordoc,omitempty" xml:"pk_reqstordoc,omitempty" comment:"需求仓库"`         
	PkArrvstoorg       string  `json:"pk_arrvstoorg,omitempty" xml:"pk_arrvstoorg,omitempty" comment:"收货库存组织最新版本"`         
	PkArrvstoorgV      string  `json:"pk_arrvstoorg_v,omitempty" xml:"pk_arrvstoorg_v,omitempty" comment:"收货库存组织"`     
	PkReqdept          string  `json:"pk_reqdept,omitempty" xml:"pk_reqdept,omitempty" comment:"需求部门最新版本"`               
	PkFlowstockorg     string  `json:"pk_flowstockorg,omitempty" xml:"pk_flowstockorg,omitempty" comment:"物流组织最新版本"`     
	PkFlowstockorgV    string  `json:"pk_flowstockorg_v,omitempty" xml:"pk_flowstockorg_v,omitempty" comment:"物流组织"` 
	Crowno             string  `json:"crowno,omitempty" xml:"crowno,omitempty" comment:"行号"`                       
	PkMaterial         string  `json:"pk_material,omitempty" xml:"pk_material,omitempty" comment:"物料版本信息"`             
	PkSrcmaterial      string  `json:"pk_srcmaterial,omitempty" xml:"pk_srcmaterial,omitempty" comment:"物料信息"`     
	Cunitid            string  `json:"cunitid,omitempty" xml:"cunitid,omitempty" comment:"主单位"`                     
	Nnum               float64 `json:"nnum,omitempty" xml:"nnum,omitempty" comment:"主数量"`                           
	Castunitid         string  `json:"castunitid,omitempty" xml:"castunitid,omitempty" comment:"单位"`               
	Nastnum            string  `json:"nastnum,omitempty" xml:"nastnum,omitempty" comment:"数量"`                     
	Vchangerate        string  `json:"vchangerate,omitempty" xml:"vchangerate,omitempty" comment:"换算率"`             
	Cqtunitid          string  `json:"cqtunitid,omitempty" xml:"cqtunitid,omitempty" comment:"报价单位"`                 
	Nqtunitnum         string  `json:"nqtunitnum,omitempty" xml:"nqtunitnum,omitempty" comment:"报价数量"`               
	Vqtunitrate        string  `json:"vqtunitrate,omitempty" xml:"vqtunitrate,omitempty" comment:"报价换算率"`           
	Nqtorigprice       float64 `json:"nqtorigprice,omitempty" xml:"nqtorigprice,omitempty" comment:"无税单价"`           
	Nqtorigtaxprice    float64 `json:"nqtorigtaxprice,omitempty" xml:"nqtorigtaxprice,omitempty" comment:"含税单价"`     
	Nqtorignetprice    float64 `json:"nqtorignetprice,omitempty" xml:"nqtorignetprice,omitempty" comment:"无税净价"`     
	Nqtorigtaxnetprc   float64 `json:"nqtorigtaxnetprc,omitempty" xml:"nqtorigtaxnetprc,omitempty" comment:"含税净价"`   
	Nqtnetprice        float64 `json:"nqtnetprice,omitempty" xml:"nqtnetprice,omitempty" comment:"本币无税净价"`             
	Nqttaxnetprice     float64 `json:"nqttaxnetprice,omitempty" xml:"nqttaxnetprice,omitempty" comment:"本币含税净价"`       
	Nitemdiscountrate  float64 `json:"nitemdiscountrate,omitempty" xml:"nitemdiscountrate,omitempty" comment:"折扣"` 
	Norigmny           float64 `json:"norigmny,omitempty" xml:"norigmny,omitempty" comment:"无税金额"`                   
	Norigtaxmny        float64 `json:"norigtaxmny,omitempty" xml:"norigtaxmny,omitempty" comment:"价税合计"`             
	Nmny               float64 `json:"nmny,omitempty" xml:"nmny,omitempty" comment:"本币无税金额"`                           
	Ntaxmny            float64 `json:"ntaxmny,omitempty" xml:"ntaxmny,omitempty" comment:"本币价税合计"`                     
	Ntax               float64 `json:"ntax,omitempty" xml:"ntax,omitempty" comment:"本币税额"`                           
	Norigprice         float64 `json:"norigprice,omitempty" xml:"norigprice,omitempty" comment:"主无税单价"`               
	Norigtaxprice      float64 `json:"norigtaxprice,omitempty" xml:"norigtaxprice,omitempty" comment:"主含税单价"`         
	Norignetprice      float64 `json:"norignetprice,omitempty" xml:"norignetprice,omitempty" comment:"主无税净价"`         
	Norigtaxnetprice   float64 `json:"norigtaxnetprice,omitempty" xml:"norigtaxnetprice,omitempty" comment:"主含税净价"` 
	Nnetprice          float64 `json:"nnetprice,omitempty" xml:"nnetprice,omitempty" comment:"主本币无税净价"`                 
	Ntaxnetprice       float64 `json:"ntaxnetprice,omitempty" xml:"ntaxnetprice,omitempty" comment:"主本币含税净价"`          
	Naccumarrvnum      float64 `json:"naccumarrvnum,omitempty" xml:"naccumarrvnum,omitempty" comment:"累计到货主数量"`         
	Naccumstorenum     float64 `json:"naccumstorenum,omitempty" xml:"naccumstorenum,omitempty" comment:"累计入库主数量"`       
	Naccuminvoicenum   float64 `json:"naccuminvoicenum,omitempty" xml:"naccuminvoicenum,omitempty" comment:"累计开票主数量"`  
	Naccumwastnum      float64 `json:"naccumwastnum,omitempty" xml:"naccumwastnum,omitempty" comment:"累计途耗主数量"`         
	Dplanarrvdate      string  `json:"dplanarrvdate,omitempty" xml:"dplanarrvdate,omitempty" comment:""`         
	PkRecvstordoc      string  `json:"pk_recvstordoc,omitempty" xml:"pk_recvstordoc,omitempty" comment:"收货仓库"`       
	PkReceiveaddress   string  `json:"pk_receiveaddress,omitempty" xml:"pk_receiveaddress,omitempty" comment:"收货地址"` 
	Dcorrectdate       string  `json:"dcorrectdate,omitempty" xml:"dcorrectdate,omitempty" comment:"修正日期"`           
	Chandler           string  `json:"chandler,omitempty" xml:"chandler,omitempty" comment:"操作员"`                 
	Fisactive          int     `json:"fisactive,omitempty" xml:"fisactive,omitempty" comment:"激活"`                
	Csourcetypecode    string  `json:"csourcetypecode,omitempty" xml:"csourcetypecode" comment:"来源单据类型"`     
	Csourceid          string  `json:"csourceid,omitempty" xml:"csourceid,omitempty" comment:""`                 
	Csourcebid         string  `json:"csourcebid,omitempty" xml:"csourcebid,omitempty" comment:""`
	Cfirsttypecode     string  `json:"cfirsttypecode,omitempty" xml:"cfirsttypecode" comment:"源头单据类型"` 
	Cfirstid           string  `json:"cfirstid,omitempty" xml:"cfirstid,omitempty" comment:""`
	Cfirstbid          string  `json:"cfirstbid,omitempty" xml:"cfirstbid,omitempty" comment:""`
	Vbmemo             string  `json:"vbmemo,omitempty" xml:"vbmemo,omitempty" comment:"备注"`                   
	PkBatchcode        string  `json:"pk_batchcode,omitempty" xml:"pk_batchcode,omitempty" comment:""`       
	Vbatchcode         string  `json:"vbatchcode,omitempty" xml:"vbatchcode,omitempty" comment:"批次号"`           
	Nbackarrvnum       string  `json:"nbackarrvnum,omitempty" xml:"nbackarrvnum,omitempty" comment:"累计退货主数量"`       
	Nbackstorenum      int     `json:"nbackstorenum,omitempty" xml:"nbackstorenum,omitempty" comment:"累计退库主数量"`     
	Cqpbaseschemeid    string  `json:"cqpbaseschemeid,omitempty" xml:"cqpbaseschemeid,omitempty" comment:"优质优价方案"` 
	Ccontractrowid     string  `json:"ccontractrowid,omitempty" xml:"ccontractrowid,omitempty" comment:""`
	Ccontractid        string  `json:"ccontractid,omitempty" xml:"ccontractid,omitempty" comment:"合同信息"`             
	Vcontractcode      string  `json:"vcontractcode,omitempty" xml:"vcontractcode,omitempty" comment:"合同号"`         
	Vpriceauditcode    string  `json:"vpriceauditcode,omitempty" xml:"vpriceauditcode,omitempty" comment:"价格审批单号"`  
	Cpriceauditid      string  `json:"cpriceauditid,omitempty" xml:"cpriceauditid,omitempty" comment:"价格审批单"`         
	CpriceauditBid     string  `json:"cpriceaudit_bid,omitempty" xml:"cpriceaudit_bid,omitempty" comment:""`     
	CpriceauditBb1id   string  `json:"cpriceaudit_bb1id,omitempty" xml:"cpriceaudit_bb1id,omitempty" comment:""` 
	Breceiveplan       string  `json:"breceiveplan,omitempty" xml:"breceiveplan,omitempty" comment:"存在到货计划"`           
	Naccumrpnum        float64 `json:"naccumrpnum,omitempty" xml:"naccumrpnum,omitempty" comment:"累计到货计划主数量"`             
	Blargess           string  `json:"blargess,omitempty" xml:"blargess,omitempty" comment:"是否赠品"`                   
	Cvenddevareaid     string  `json:"cvenddevareaid,omitempty" xml:"cvenddevareaid,omitempty" comment:"供应商发货地区"`       
	Cvenddevaddrid     string  `json:"cvenddevaddrid,omitempty" xml:"cvenddevaddrid,omitempty" comment:"供应商发货地点"`       
	Vvenddevaddr       string  `json:"vvenddevaddr,omitempty" xml:"vvenddevaddr,omitempty" comment:"供应商发货地址"`           
	Cdevareaid         string  `json:"cdevareaid,omitempty" xml:"cdevareaid,omitempty" comment:"收货地区"`               
	Cdevaddrid         string  `json:"cdevaddrid,omitempty" xml:"cdevaddrid,omitempty" comment:"收货地点"`               
	Naccumdevnum       float64 `json:"naccumdevnum,omitempty" xml:"naccumdevnum,omitempty" comment:"累计运输主数量"`           
	Ncaninnum          float64 `json:"ncaninnum,omitempty" xml:"ncaninnum,omitempty" comment:""`
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
	Vbdef1             string  `json:"vbdef1,omitempty" xml:"vbdef1,omitempty" comment:""`
	Vbdef2             string  `json:"vbdef2,omitempty" xml:"vbdef2,omitempty" comment:""`
	Vbdef3             string  `json:"vbdef3,omitempty" xml:"vbdef3,omitempty" comment:""`
	Vbdef4             string  `json:"vbdef4,omitempty" xml:"vbdef4,omitempty" comment:""`
	Vbdef5             string  `json:"vbdef5,omitempty" xml:"vbdef5,omitempty" comment:""`
	Vvendinventorycode string  `json:"vvendinventorycode,omitempty" xml:"vvendinventorycode,omitempty" comment:"对应物料编码"` 
	Vvendinventoryname string  `json:"vvendinventoryname,omitempty" xml:"vvendinventoryname,omitempty" comment:"对应物料名称"` 
	Btransclosed       string  `json:"btransclosed,omitempty" xml:"btransclosed,omitempty" comment:"是否运输关闭"`              
	Nfeemny            float64 `json:"nfeemny,omitempty" xml:"nfeemny,omitempty" comment:"费用累计开票金额"`                       
	PkPsfinanceorg     string  `json:"pk_psfinanceorg,omitempty" xml:"pk_psfinanceorg,omitempty" comment:"结算财务组织最新版本"`       
	PkPsfinanceorgV    string  `json:"pk_psfinanceorg_v,omitempty" xml:"pk_psfinanceorg_v,omitempty" comment:"结算财务组织"`  
	PkApfinanceorg     string  `json:"pk_apfinanceorg,omitempty" xml:"pk_apfinanceorg,omitempty" comment:"应付组织最新版本"`       
	PkApfinanceorgV    string  `json:"pk_apfinanceorg_v,omitempty" xml:"pk_apfinanceorg_v,omitempty" comment:"应付组织"`   
	PkApliabcenter     string  `json:"pk_apliabcenter,omitempty" xml:"pk_apliabcenter,omitempty" comment:"利润中心最新版本"`       
	PkApliabcenterV    string  `json:"pk_apliabcenter_v,omitempty"  comment:"利润中心"`                           
	Nacccancelpaymny   float64 `json:"nacccancelpaymny,omitempty" xml:"nacccancelpaymny,omitempty" comment:"累计已核销本币付款金额"`     
	Bborrowpur         string  `json:"bborrowpur,omitempty" xml:"bborrowpur,omitempty" comment:"借入转采购"`               
	Vfirsttrantype     string  `json:"vfirsttrantype,omitempty" xml:"vfirsttrantype,omitempty" comment:"源头交易类型"`         
	Vfirstcode         string  `json:"vfirstcode,omitempty" xml:"vfirstcode,omitempty" comment:"源头单据号"`               
	Vfirstrowno        string  `json:"vfirstrowno,omitempty" xml:"vfirstrowno,omitempty" comment:"源头单据行号"`               
	Vsourcetrantype    string  `json:"vsourcetrantype,omitempty" xml:"vsourcetrantype,omitempty" comment:"来源交易类型"`       
	Vsourcecode        string  `json:"vsourcecode,omitempty" xml:"vsourcecode,omitempty" comment:"来源单据号"`               
	Vsourcerowno       string  `json:"vsourcerowno,omitempty" xml:"vsourcerowno,omitempty" comment:"来源单据行号"`             
	Naccuminvoicemny   float64 `json:"naccuminvoicemny,omitempty" xml:"naccuminvoicemny,omitempty" comment:"累计本币开票金额"`     
	Nacccancelinvmny   float64 `json:"nacccancelinvmny,omitempty" xml:"nacccancelinvmny,omitempty" comment:"累计已核销本币开票金额"`     
	Nweight            float64 `json:"nweight,omitempty" xml:"nweight,omitempty" comment:"重量"`                       
	Nvolumn            float64 `json:"nvolumn,omitempty" xml:"nvolumn,omitempty" comment:"体积"`                       
	Npacknum           float64 `json:"npacknum,omitempty" xml:"npacknum,omitempty" comment:"件数"`                   
	Bstockclose        string  `json:"bstockclose,omitempty" xml:"bstockclose,omitempty" comment:"入库关闭"`               
	Binvoiceclose      string  `json:"binvoiceclose,omitempty" xml:"binvoiceclose,omitempty" comment:"开票关闭"`           
	Barriveclose       string  `json:"barriveclose,omitempty" xml:"barriveclose,omitempty" comment:"到货关闭"`             
	Bpayclose          string  `json:"bpayclose,omitempty" xml:"bpayclose,omitempty" comment:"付款关闭"`                   
	Ftaxtypeflag       int     `json:"ftaxtypeflag,omitempty" xml:"ftaxtypeflag,omitempty" comment:"扣税类别"`           
	Ntaxrate           float64 `json:"ntaxrate,omitempty" xml:"ntaxrate,omitempty" comment:"税率"`                     
	Ccurrencyid        string  `json:"ccurrencyid,omitempty" xml:"ccurrencyid,omitempty" comment:"本币币种"`               
	Nexchangerate      float64 `json:"nexchangerate,omitempty" xml:"nexchangerate,omitempty" comment:"折本汇率"`           
	Ngroupexchgrate    float64 `json:"ngroupexchgrate,omitempty" xml:"ngroupexchgrate,omitempty" comment:"集团本位币汇率"`       
	Nglobalexchgrate   float64 `json:"nglobalexchgrate,omitempty" xml:"nglobalexchgrate,omitempty" comment:"全局本位币汇率"`     
	Ngroupmny          float64 `json:"ngroupmny,omitempty" xml:"ngroupmny,omitempty" comment:"集团本币无税金额"`                   
	Ngrouptaxmny       float64 `json:"ngrouptaxmny,omitempty" xml:"ngrouptaxmny,omitempty" comment:"集团本币价税合计"`             
	Nglobalmny         float64 `json:"nglobalmny,omitempty" xml:"nglobalmny,omitempty" comment:"全局本币无税金额"`                 
	Nglobaltaxmny      float64 `json:"nglobaltaxmny,omitempty" xml:"nglobaltaxmny,omitempty" comment:"全局本币价税合计"`           
	Cprojectid         string  `json:"cprojectid,omitempty" xml:"cprojectid,omitempty" comment:"项目"`                
	Cproductorid       string  `json:"cproductorid,omitempty" xml:"cproductorid,omitempty" comment:"生产厂商"`             
	Dbilldate          string  `json:"dbilldate,omitempty" xml:"dbilldate,omitempty" comment:"订单日期"`                   
	PkSupplier         string  `json:"pk_supplier,omitempty" xml:"pk_supplier,omitempty" comment:"供应商"`               
	Corigcurrencyid    string  `json:"corigcurrencyid,omitempty" xml:"corigcurrencyid,omitempty" comment:"币种"`      
	Cqualitylevelid    string  `json:"cqualitylevelid,omitempty" xml:"cqualitylevelid,omitempty" comment:"质量等级"`       
	Nsuprsnum          float64 `json:"nsuprsnum,omitempty" xml:"nsuprsnum,omitempty" comment:"被预留数量"`                   
	PkSrcorderB        string  `json:"pk_srcorder_b,omitempty" xml:"pk_srcorder_b,omitempty" comment:""`            
	PkReqdeptV         string  `json:"pk_reqdept_v,omitempty" xml:"pk_reqdept_v,omitempty" comment:"需求部门"`             
	Nqtprice           float64 `json:"nqtprice,omitempty" xml:"nqtprice,omitempty" comment:"报价本币无税单价"`                     
	Nqttaxprice        float64 `json:"nqttaxprice,omitempty" xml:"nqttaxprice" comment:"报价本币含税单价"`               
	Nprice             float64 `json:"nprice,omitempty" xml:"nprice,omitempty" comment:"主本币无税单价"`                         
	Ntaxprice          float64 `json:"ntaxprice,omitempty" xml:"ntaxprice,omitempty" comment:"主本币含税单价"`                   
	Cecbillid          string  `json:"cecbillid,omitempty" xml:"cecbillid,omitempty" comment:""`                   
	Cecbillbid         string  `json:"cecbillbid,omitempty" xml:"cecbillbid,omitempty" comment:""`                 
	Vecbillcode        string  `json:"vecbillcode,omitempty" xml:"vecbillcode,omitempty" comment:"电子商务单据号"`               
	Cectypecode        string  `json:"cectypecode,omitempty" xml:"cectypecode,omitempty" comment:"电子商务单据类型"`               
	Norigordernum      float64 `json:"norigordernum,omitempty" xml:"norigordernum,omitempty" comment:"原始订单数量"`           
	Norigorderprice    float64 `json:"norigorderprice,omitempty" xml:"norigorderprice,omitempty" comment:"原始订单净含税单价"`       
	Dorigplanarrvdate  string  `json:"dorigplanarrvdate,omitempty" xml:"dorigplanarrvdate,omitempty" comment:"原始计划到货日期"`   
	Istorestatus       string  `json:"istorestatus,omitempty" xml:"istorestatus,omitempty" comment:"供应商交货状态"`             
	Casscustid         string  `json:"casscustid,omitempty" xml:"casscustid,omitempty" comment:"客户"`                 
	Cprojecttaskid     string  `json:"cprojecttaskid,omitempty" xml:"cprojecttaskid,omitempty" comment:"项目任务"`        
	Csendcountryid     string  `json:"csendcountryid,omitempty" xml:"csendcountryid,omitempty" comment:"发货国家/地区"`         
	Crececountryid     string  `json:"crececountryid,omitempty" xml:"crececountryid,omitempty" comment:"收货国家/地区"`         
	Ctaxcountryid      string  `json:"ctaxcountryid,omitempty" xml:"ctaxcountryid,omitempty" comment:"报税国家/地区"`           
	Fbuysellflag       int     `json:"fbuysellflag,omitempty" xml:"fbuysellflag,omitempty" comment:"购销类型"`             
	Btriatradeflag     string  `json:"btriatradeflag,omitempty" xml:"btriatradeflag,omitempty" comment:"三角贸易"`         
	Ctaxcodeid         string  `json:"ctaxcodeid,omitempty" xml:"ctaxcodeid,omitempty" comment:"税码"`                 
	Nnosubtaxrate      float64 `json:"nnosubtaxrate,omitempty" xml:"nnosubtaxrate,omitempty" comment:"不可抵扣税率"`         
	Nnosubtax          float64 `json:"nnosubtax,omitempty" xml:"nnosubtax,omitempty" comment:"不可抵扣税额"`                   
	Ncaltaxmny         float64 `json:"ncaltaxmny,omitempty" xml:"ncaltaxmny,omitempty" comment:"计税金额"`                 
	Ncalcostmny        float64 `json:"ncalcostmny,omitempty" xml:"ncalcostmny,omitempty" comment:"计成本金额"`               
	Corigcountryid     string  `json:"corigcountryid,omitempty" xml:"corigcountryid,omitempty" comment:"原产国"`         
	Corigareaid        string  `json:"corigareaid,omitempty" xml:"corigareaid,omitempty" comment:"原产地区"`               
	Cdesticountryid    string  `json:"cdesticountryid,omitempty" xml:"cdesticountryid,omitempty" comment:"目的地国"`      
	Cdestiareaid       string  `json:"cdestiareaid,omitempty" xml:"cdestiareaid,omitempty" comment:"目的地地区"`             
	Nsourcenum         float64 `json:"nsourcenum,omitempty" xml:"nsourcenum,omitempty" comment:"来源单据主数量"`                 
}

type ResPurchaseOrder struct {
	XMLName    xml.Name `xml:"ufinterface"`
	Successful string   `xml:"successful,attr"`
	base.Base
	Sendresult struct {
		Billpk            string     `xml:"billpk"`
		Bdocid            string     `xml:"bdocid"`
		Filename          string     `xml:"filename"`
		Resultcode        int        `xml:"resultcode"`
		Resultdescription string     `xml:"resultdescription"`
		//Invaliddoc        []Docitem `xml:"invaliddoc"`
		Content           string     `xml:"content"`
	} `xml:"sendresult"`
}

//Purchase struct
type Purchase struct {
	*context.Context
}

//NewSale instance
func NewPurchase(context *context.Context) *Purchase {
	purchase := new(Purchase)
	purchase.Context = context
	return purchase
}

func (purchase *Purchase) GetPurchaseOrder(req ReqPurchaseOrder) (res ResPurchaseOrder, err error) {
	accessParam := ""
	if accessParam, err = purchase.GetUrlParam(); err != nil {
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
		err = fmt.Errorf("GetPurchaseOrder Error , errcode=%d , errmsg=%s", res.Sendresult.Resultcode, res.Sendresult.Resultdescription)
		return
	}
	return
}

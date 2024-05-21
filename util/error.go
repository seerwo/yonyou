package util

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/yonyou/nc/base"
	"reflect"
)

// CommonError nc return common json
type CommonError struct {
	base.Base
	SendResult struct {
		Billpk string `json:"billpk" xml:"billpk"`
		Bdocid string `json:"bdocid" xml:"bdocid"`
		Filename string `json:"filename" xml:"filename"`
		Resultcode string `json:"resultcode" xml:"resultcode"`
		Resultdescription string `json:"resultdescription" xml:"resultdescription"`
		Content string `json:"content" xml:"content"`
	} `json:"sendresult" xml:"sendresult"`
 	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DecodeWithCommonError return CommonError json
func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commError CommonError
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return
	}
	if commError.ErrCode != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, commError.ErrCode, commError.ErrMsg)
	}
	return nil
}

// DecodeWithError return json
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, errCode.Int(), errMsg.String())
	}
	return nil
}
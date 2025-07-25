// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCardOcrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CardOcrResponseBody
	GetCode() *string
	SetMessage(v string) *CardOcrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CardOcrResponseBody
	GetRequestId() *string
	SetResult(v *CardOcrResponseBodyResult) *CardOcrResponseBody
	GetResult() *CardOcrResponseBodyResult
}

type CardOcrResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CardOcrResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CardOcrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CardOcrResponseBody) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CardOcrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CardOcrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CardOcrResponseBody) GetResult() *CardOcrResponseBodyResult {
	return s.Result
}

func (s *CardOcrResponseBody) SetCode(v string) *CardOcrResponseBody {
	s.Code = &v
	return s
}

func (s *CardOcrResponseBody) SetMessage(v string) *CardOcrResponseBody {
	s.Message = &v
	return s
}

func (s *CardOcrResponseBody) SetRequestId(v string) *CardOcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CardOcrResponseBody) SetResult(v *CardOcrResponseBodyResult) *CardOcrResponseBody {
	s.Result = v
	return s
}

func (s *CardOcrResponseBody) Validate() error {
	return dara.Validate(s)
}

type CardOcrResponseBodyResult struct {
	ExtCardInfo *string `json:"ExtCardInfo,omitempty" xml:"ExtCardInfo,omitempty"`
	ExtIdInfo   *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CardOcrResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CardOcrResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardOcrResponseBodyResult) GetExtCardInfo() *string {
	return s.ExtCardInfo
}

func (s *CardOcrResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *CardOcrResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *CardOcrResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CardOcrResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CardOcrResponseBodyResult) SetExtCardInfo(v string) *CardOcrResponseBodyResult {
	s.ExtCardInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetExtIdInfo(v string) *CardOcrResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetPassed(v string) *CardOcrResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetSubCode(v string) *CardOcrResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CardOcrResponseBodyResult) SetTransactionId(v string) *CardOcrResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *CardOcrResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaPeriodVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaPeriodVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaPeriodVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *Id2MetaPeriodVerifyIntlResponseBodyResult) *Id2MetaPeriodVerifyIntlResponseBody
	GetResult() *Id2MetaPeriodVerifyIntlResponseBodyResult
}

type Id2MetaPeriodVerifyIntlResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7B97D932-7FF5-517D-BF39-7CA1BEE3CDD9
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Id2MetaPeriodVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s Id2MetaPeriodVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) GetResult() *Id2MetaPeriodVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetCode(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetMessage(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetRequestId(v string) *Id2MetaPeriodVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) SetResult(v *Id2MetaPeriodVerifyIntlResponseBodyResult) *Id2MetaPeriodVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type Id2MetaPeriodVerifyIntlResponseBodyResult struct {
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Id2MetaPeriodVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) SetPassed(v string) *Id2MetaPeriodVerifyIntlResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) SetSubCode(v string) *Id2MetaPeriodVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *Id2MetaPeriodVerifyIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

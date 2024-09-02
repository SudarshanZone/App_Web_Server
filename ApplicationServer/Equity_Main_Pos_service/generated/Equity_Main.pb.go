// protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal/positions/Equity_Main.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: positions/Equity_Main.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpClmMtchAcct string `protobuf:"bytes,1,opt,name=otp_clm_mtch_acct,json=otpClmMtchAcct,proto3" json:"otp_clm_mtch_acct,omitempty"`
}

func (x *PositionRequest) Reset() {
	*x = PositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positions_Equity_Main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionRequest) ProtoMessage() {}

func (x *PositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_positions_Equity_Main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionRequest.ProtoReflect.Descriptor instead.
func (*PositionRequest) Descriptor() ([]byte, []int) {
	return file_positions_Equity_Main_proto_rawDescGZIP(), []int{0}
}

func (x *PositionRequest) GetOtpClmMtchAcct() string {
	if x != nil {
		return x.OtpClmMtchAcct
	}
	return ""
}

type EquityPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpClmMtchAcct       string  `protobuf:"bytes,1,opt,name=otp_clm_mtch_acct,json=otpClmMtchAcct,proto3" json:"otp_clm_mtch_acct,omitempty"`                    // character(10)
	OtpXchngCd           string  `protobuf:"bytes,2,opt,name=otp_xchng_cd,json=otpXchngCd,proto3" json:"otp_xchng_cd,omitempty"`                                  // character(3)
	OtpXchngSgmntCd      string  `protobuf:"bytes,3,opt,name=otp_xchng_sgmnt_cd,json=otpXchngSgmntCd,proto3" json:"otp_xchng_sgmnt_cd,omitempty"`                 // character(2)
	OtpXchngSgmntSttlmnt int32   `protobuf:"varint,4,opt,name=otp_xchng_sgmnt_sttlmnt,json=otpXchngSgmntSttlmnt,proto3" json:"otp_xchng_sgmnt_sttlmnt,omitempty"` // integer
	OtpStckCd            string  `protobuf:"bytes,5,opt,name=otp_stck_cd,json=otpStckCd,proto3" json:"otp_stck_cd,omitempty"`                                     // character(6)
	OtpFlw               string  `protobuf:"bytes,6,opt,name=otp_flw,json=otpFlw,proto3" json:"otp_flw,omitempty"`                                                // character(1)
	OtpQty               int64   `protobuf:"varint,7,opt,name=otp_qty,json=otpQty,proto3" json:"otp_qty,omitempty"`                                               // bigint
	OtpCnvrtDlvryQty     int64   `protobuf:"varint,8,opt,name=otp_cnvrt_dlvry_qty,json=otpCnvrtDlvryQty,proto3" json:"otp_cnvrt_dlvry_qty,omitempty"`             // bigint
	OtpCvrdQty           int64   `protobuf:"varint,9,opt,name=otp_cvrd_qty,json=otpCvrdQty,proto3" json:"otp_cvrd_qty,omitempty"`                                 // bigint
	OtpRt                string  `protobuf:"bytes,10,opt,name=otp_rt,json=otpRt,proto3" json:"otp_rt,omitempty"`                                                  // numeric(20,10) - represented as string in proto
	OtpMrgnAmt           float64 `protobuf:"fixed64,11,opt,name=otp_mrgn_amt,json=otpMrgnAmt,proto3" json:"otp_mrgn_amt,omitempty"`                               // double precision
	OtpTrdVal            float64 `protobuf:"fixed64,12,opt,name=otp_trd_val,json=otpTrdVal,proto3" json:"otp_trd_val,omitempty"`                                  // double precision
	OtpRmrks             string  `protobuf:"bytes,13,opt,name=otp_rmrks,json=otpRmrks,proto3" json:"otp_rmrks,omitempty"`                                         // character varying(50)
	OtpXferMrgnStts      string  `protobuf:"bytes,14,opt,name=otp_xfer_mrgn_stts,json=otpXferMrgnStts,proto3" json:"otp_xfer_mrgn_stts,omitempty"`                // character(1)
	OtpSellOpnPrccsd     string  `protobuf:"bytes,15,opt,name=otp_sell_opn_prccsd,json=otpSellOpnPrccsd,proto3" json:"otp_sell_opn_prccsd,omitempty"`             // character(1)
	OtpBuyOpnPrccsd      string  `protobuf:"bytes,16,opt,name=otp_buy_opn_prccsd,json=otpBuyOpnPrccsd,proto3" json:"otp_buy_opn_prccsd,omitempty"`                // character(1)
	OtpMrgnSqroffMode    string  `protobuf:"bytes,17,opt,name=otp_mrgn_sqroff_mode,json=otpMrgnSqroffMode,proto3" json:"otp_mrgn_sqroff_mode,omitempty"`          // character(1)
	OtpEmTrdspltPrcsFlg  string  `protobuf:"bytes,18,opt,name=otp_em_trdsplt_prcs_flg,json=otpEmTrdspltPrcsFlg,proto3" json:"otp_em_trdsplt_prcs_flg,omitempty"`  // character(1)
	OtpMtmFlg            string  `protobuf:"bytes,19,opt,name=otp_mtm_flg,json=otpMtmFlg,proto3" json:"otp_mtm_flg,omitempty"`                                    // character(1)
	OtpMtmCansq          string  `protobuf:"bytes,20,opt,name=otp_mtm_cansq,json=otpMtmCansq,proto3" json:"otp_mtm_cansq,omitempty"`                              // character(1)
	OtpEosCan            string  `protobuf:"bytes,21,opt,name=otp_eos_can,json=otpEosCan,proto3" json:"otp_eos_can,omitempty"`                                    // character(1)
	OtpTrgrPrc           float64 `protobuf:"fixed64,22,opt,name=otp_trgr_prc,json=otpTrgrPrc,proto3" json:"otp_trgr_prc,omitempty"`                               // double precision
	Otp_16TrgrPrc        float32 `protobuf:"fixed32,23,opt,name=otp_16_trgr_prc,json=otp16TrgrPrc,proto3" json:"otp_16_trgr_prc,omitempty"`                       // double precision
	OtpMinMrgn           float64 `protobuf:"fixed64,24,opt,name=otp_min_mrgn,json=otpMinMrgn,proto3" json:"otp_min_mrgn,omitempty"`                               // double precision
	OtpT5TrdspltPrcsFlg  string  `protobuf:"bytes,25,opt,name=otp_t5_trdsplt_prcs_flg,json=otpT5TrdspltPrcsFlg,proto3" json:"otp_t5_trdsplt_prcs_flg,omitempty"`
}

func (x *EquityPosition) Reset() {
	*x = EquityPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positions_Equity_Main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquityPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquityPosition) ProtoMessage() {}

func (x *EquityPosition) ProtoReflect() protoreflect.Message {
	mi := &file_positions_Equity_Main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquityPosition.ProtoReflect.Descriptor instead.
func (*EquityPosition) Descriptor() ([]byte, []int) {
	return file_positions_Equity_Main_proto_rawDescGZIP(), []int{1}
}

func (x *EquityPosition) GetOtpClmMtchAcct() string {
	if x != nil {
		return x.OtpClmMtchAcct
	}
	return ""
}

func (x *EquityPosition) GetOtpXchngCd() string {
	if x != nil {
		return x.OtpXchngCd
	}
	return ""
}

func (x *EquityPosition) GetOtpXchngSgmntCd() string {
	if x != nil {
		return x.OtpXchngSgmntCd
	}
	return ""
}

func (x *EquityPosition) GetOtpXchngSgmntSttlmnt() int32 {
	if x != nil {
		return x.OtpXchngSgmntSttlmnt
	}
	return 0
}

func (x *EquityPosition) GetOtpStckCd() string {
	if x != nil {
		return x.OtpStckCd
	}
	return ""
}

func (x *EquityPosition) GetOtpFlw() string {
	if x != nil {
		return x.OtpFlw
	}
	return ""
}

func (x *EquityPosition) GetOtpQty() int64 {
	if x != nil {
		return x.OtpQty
	}
	return 0
}

func (x *EquityPosition) GetOtpCnvrtDlvryQty() int64 {
	if x != nil {
		return x.OtpCnvrtDlvryQty
	}
	return 0
}

func (x *EquityPosition) GetOtpCvrdQty() int64 {
	if x != nil {
		return x.OtpCvrdQty
	}
	return 0
}

func (x *EquityPosition) GetOtpRt() string {
	if x != nil {
		return x.OtpRt
	}
	return ""
}

func (x *EquityPosition) GetOtpMrgnAmt() float64 {
	if x != nil {
		return x.OtpMrgnAmt
	}
	return 0
}

func (x *EquityPosition) GetOtpTrdVal() float64 {
	if x != nil {
		return x.OtpTrdVal
	}
	return 0
}

func (x *EquityPosition) GetOtpRmrks() string {
	if x != nil {
		return x.OtpRmrks
	}
	return ""
}

func (x *EquityPosition) GetOtpXferMrgnStts() string {
	if x != nil {
		return x.OtpXferMrgnStts
	}
	return ""
}

func (x *EquityPosition) GetOtpSellOpnPrccsd() string {
	if x != nil {
		return x.OtpSellOpnPrccsd
	}
	return ""
}

func (x *EquityPosition) GetOtpBuyOpnPrccsd() string {
	if x != nil {
		return x.OtpBuyOpnPrccsd
	}
	return ""
}

func (x *EquityPosition) GetOtpMrgnSqroffMode() string {
	if x != nil {
		return x.OtpMrgnSqroffMode
	}
	return ""
}

func (x *EquityPosition) GetOtpEmTrdspltPrcsFlg() string {
	if x != nil {
		return x.OtpEmTrdspltPrcsFlg
	}
	return ""
}

func (x *EquityPosition) GetOtpMtmFlg() string {
	if x != nil {
		return x.OtpMtmFlg
	}
	return ""
}

func (x *EquityPosition) GetOtpMtmCansq() string {
	if x != nil {
		return x.OtpMtmCansq
	}
	return ""
}

func (x *EquityPosition) GetOtpEosCan() string {
	if x != nil {
		return x.OtpEosCan
	}
	return ""
}

func (x *EquityPosition) GetOtpTrgrPrc() float64 {
	if x != nil {
		return x.OtpTrgrPrc
	}
	return 0
}

func (x *EquityPosition) GetOtp_16TrgrPrc() float32 {
	if x != nil {
		return x.Otp_16TrgrPrc
	}
	return 0
}

func (x *EquityPosition) GetOtpMinMrgn() float64 {
	if x != nil {
		return x.OtpMinMrgn
	}
	return 0
}

func (x *EquityPosition) GetOtpT5TrdspltPrcsFlg() string {
	if x != nil {
		return x.OtpT5TrdspltPrcsFlg
	}
	return ""
}

type PositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Equity []*EquityPosition `protobuf:"bytes,1,rep,name=equity,proto3" json:"equity,omitempty"`
}

func (x *PositionResponse) Reset() {
	*x = PositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positions_Equity_Main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionResponse) ProtoMessage() {}

func (x *PositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_positions_Equity_Main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionResponse.ProtoReflect.Descriptor instead.
func (*PositionResponse) Descriptor() ([]byte, []int) {
	return file_positions_Equity_Main_proto_rawDescGZIP(), []int{2}
}

func (x *PositionResponse) GetEquity() []*EquityPosition {
	if x != nil {
		return x.Equity
	}
	return nil
}

var File_positions_Equity_Main_proto protoreflect.FileDescriptor

var file_positions_Equity_Main_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x45, 0x71, 0x75, 0x69,
	0x74, 0x79, 0x5f, 0x4d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x6f, 0x74,
	0x70, 0x5f, 0x63, 0x6c, 0x6d, 0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x74, 0x70, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63,
	0x68, 0x41, 0x63, 0x63, 0x74, 0x22, 0xcf, 0x07, 0x0a, 0x0e, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x11, 0x6f, 0x74, 0x70, 0x5f,
	0x63, 0x6c, 0x6d, 0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x74, 0x70, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63, 0x68, 0x41,
	0x63, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x74, 0x70, 0x5f, 0x78, 0x63, 0x68, 0x6e, 0x67,
	0x5f, 0x63, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x70, 0x58, 0x63,
	0x68, 0x6e, 0x67, 0x43, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x74, 0x70, 0x5f, 0x78, 0x63, 0x68,
	0x6e, 0x67, 0x5f, 0x73, 0x67, 0x6d, 0x6e, 0x74, 0x5f, 0x63, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x74, 0x70, 0x58, 0x63, 0x68, 0x6e, 0x67, 0x53, 0x67, 0x6d, 0x6e, 0x74,
	0x43, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x6f, 0x74, 0x70, 0x5f, 0x78, 0x63, 0x68, 0x6e, 0x67, 0x5f,
	0x73, 0x67, 0x6d, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x74, 0x6c, 0x6d, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x14, 0x6f, 0x74, 0x70, 0x58, 0x63, 0x68, 0x6e, 0x67, 0x53, 0x67, 0x6d,
	0x6e, 0x74, 0x53, 0x74, 0x74, 0x6c, 0x6d, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x74, 0x70,
	0x5f, 0x73, 0x74, 0x63, 0x6b, 0x5f, 0x63, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x74, 0x70, 0x53, 0x74, 0x63, 0x6b, 0x43, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x74, 0x70,
	0x5f, 0x66, 0x6c, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x74, 0x70, 0x46,
	0x6c, 0x77, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x74, 0x70, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x74, 0x70, 0x51, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x6f,
	0x74, 0x70, 0x5f, 0x63, 0x6e, 0x76, 0x72, 0x74, 0x5f, 0x64, 0x6c, 0x76, 0x72, 0x79, 0x5f, 0x71,
	0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6f, 0x74, 0x70, 0x43, 0x6e, 0x76,
	0x72, 0x74, 0x44, 0x6c, 0x76, 0x72, 0x79, 0x51, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x74,
	0x70, 0x5f, 0x63, 0x76, 0x72, 0x64, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6f, 0x74, 0x70, 0x43, 0x76, 0x72, 0x64, 0x51, 0x74, 0x79, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x74, 0x70, 0x5f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74,
	0x70, 0x52, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x74, 0x70, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x5f,
	0x61, 0x6d, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6f, 0x74, 0x70, 0x4d, 0x72,
	0x67, 0x6e, 0x41, 0x6d, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x74, 0x70, 0x5f, 0x74, 0x72, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f, 0x74, 0x70, 0x54,
	0x72, 0x64, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x74, 0x70, 0x5f, 0x72, 0x6d, 0x72,
	0x6b, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x74, 0x70, 0x52, 0x6d, 0x72,
	0x6b, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x74, 0x70, 0x5f, 0x78, 0x66, 0x65, 0x72, 0x5f, 0x6d,
	0x72, 0x67, 0x6e, 0x5f, 0x73, 0x74, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6f, 0x74, 0x70, 0x58, 0x66, 0x65, 0x72, 0x4d, 0x72, 0x67, 0x6e, 0x53, 0x74, 0x74, 0x73, 0x12,
	0x2d, 0x0a, 0x13, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x6e, 0x5f,
	0x70, 0x72, 0x63, 0x63, 0x73, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x74,
	0x70, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x70, 0x6e, 0x50, 0x72, 0x63, 0x63, 0x73, 0x64, 0x12, 0x2b,
	0x0a, 0x12, 0x6f, 0x74, 0x70, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x70, 0x6e, 0x5f, 0x70, 0x72,
	0x63, 0x63, 0x73, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x74, 0x70, 0x42,
	0x75, 0x79, 0x4f, 0x70, 0x6e, 0x50, 0x72, 0x63, 0x63, 0x73, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x6f,
	0x74, 0x70, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x5f, 0x73, 0x71, 0x72, 0x6f, 0x66, 0x66, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x74, 0x70, 0x4d, 0x72,
	0x67, 0x6e, 0x53, 0x71, 0x72, 0x6f, 0x66, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x17,
	0x6f, 0x74, 0x70, 0x5f, 0x65, 0x6d, 0x5f, 0x74, 0x72, 0x64, 0x73, 0x70, 0x6c, 0x74, 0x5f, 0x70,
	0x72, 0x63, 0x73, 0x5f, 0x66, 0x6c, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f,
	0x74, 0x70, 0x45, 0x6d, 0x54, 0x72, 0x64, 0x73, 0x70, 0x6c, 0x74, 0x50, 0x72, 0x63, 0x73, 0x46,
	0x6c, 0x67, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x74, 0x70, 0x5f, 0x6d, 0x74, 0x6d, 0x5f, 0x66, 0x6c,
	0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x70, 0x4d, 0x74, 0x6d, 0x46,
	0x6c, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x74, 0x70, 0x5f, 0x6d, 0x74, 0x6d, 0x5f, 0x63, 0x61,
	0x6e, 0x73, 0x71, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x70, 0x4d, 0x74,
	0x6d, 0x43, 0x61, 0x6e, 0x73, 0x71, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x74, 0x70, 0x5f, 0x65, 0x6f,
	0x73, 0x5f, 0x63, 0x61, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x70,
	0x45, 0x6f, 0x73, 0x43, 0x61, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x74, 0x70, 0x5f, 0x74, 0x72,
	0x67, 0x72, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6f, 0x74,
	0x70, 0x54, 0x72, 0x67, 0x72, 0x50, 0x72, 0x63, 0x12, 0x25, 0x0a, 0x0f, 0x6f, 0x74, 0x70, 0x5f,
	0x31, 0x36, 0x5f, 0x74, 0x72, 0x67, 0x72, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x6f, 0x74, 0x70, 0x31, 0x36, 0x54, 0x72, 0x67, 0x72, 0x50, 0x72, 0x63, 0x12,
	0x20, 0x0a, 0x0c, 0x6f, 0x74, 0x70, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6f, 0x74, 0x70, 0x4d, 0x69, 0x6e, 0x4d, 0x72, 0x67,
	0x6e, 0x12, 0x34, 0x0a, 0x17, 0x6f, 0x74, 0x70, 0x5f, 0x74, 0x35, 0x5f, 0x74, 0x72, 0x64, 0x73,
	0x70, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x63, 0x73, 0x5f, 0x66, 0x6c, 0x67, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x6f, 0x74, 0x70, 0x54, 0x35, 0x54, 0x72, 0x64, 0x73, 0x70, 0x6c, 0x74,
	0x50, 0x72, 0x63, 0x73, 0x46, 0x6c, 0x67, 0x22, 0x44, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65,
	0x71, 0x75, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x71, 0x75, 0x69, 0x74, 0x79, 0x32, 0x57, 0x0a,
	0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_positions_Equity_Main_proto_rawDescOnce sync.Once
	file_positions_Equity_Main_proto_rawDescData = file_positions_Equity_Main_proto_rawDesc
)

func file_positions_Equity_Main_proto_rawDescGZIP() []byte {
	file_positions_Equity_Main_proto_rawDescOnce.Do(func() {
		file_positions_Equity_Main_proto_rawDescData = protoimpl.X.CompressGZIP(file_positions_Equity_Main_proto_rawDescData)
	})
	return file_positions_Equity_Main_proto_rawDescData
}

var file_positions_Equity_Main_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_positions_Equity_Main_proto_goTypes = []any{
	(*PositionRequest)(nil),  // 0: position.PositionRequest
	(*EquityPosition)(nil),   // 1: position.EquityPosition
	(*PositionResponse)(nil), // 2: position.PositionResponse
}
var file_positions_Equity_Main_proto_depIdxs = []int32{
	1, // 0: position.PositionResponse.equity:type_name -> position.EquityPosition
	0, // 1: position.PositionService.GetPosition:input_type -> position.PositionRequest
	2, // 2: position.PositionService.GetPosition:output_type -> position.PositionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_positions_Equity_Main_proto_init() }
func file_positions_Equity_Main_proto_init() {
	if File_positions_Equity_Main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_positions_Equity_Main_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PositionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_positions_Equity_Main_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EquityPosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_positions_Equity_Main_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PositionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_positions_Equity_Main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_positions_Equity_Main_proto_goTypes,
		DependencyIndexes: file_positions_Equity_Main_proto_depIdxs,
		MessageInfos:      file_positions_Equity_Main_proto_msgTypes,
	}.Build()
	File_positions_Equity_Main_proto = out.File
	file_positions_Equity_Main_proto_rawDesc = nil
	file_positions_Equity_Main_proto_goTypes = nil
	file_positions_Equity_Main_proto_depIdxs = nil
}

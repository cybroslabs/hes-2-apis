package acquisition

import (
	"github.com/cybroslabs/hes-2-apis/gen/go/common"
	"k8s.io/utils/ptr"
)

func GetDataLinkFields(dataLinkProtocol DataLinkProtocol) []*common.FieldDescriptor {
	switch dataLinkProtocol {
	case DataLinkProtocol_LINKPROTO_HDLC:
		// HDLC specific fields
		return []*common.FieldDescriptor{
			common.FieldDescriptor_builder{
				FieldId:  ptr.To("negotiate"),
				Label:    ptr.To("Auto HDLC frame size"),
				Tooltip:  ptr.To("Enable auto negotiation of HDLC frame size."),
				DataType: common.FieldDataType_BOOLEAN.Enum(),
				Required: ptr.To(true),
				DefaultValue: common.FieldValue_builder{
					BoolValue: ptr.To(false),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("retransmits"),
				Label:        ptr.To("HDLC retransmissions"),
				Tooltip:      ptr.To("Number of HDLC frame retransmissions to perform. It can be set to 0 to disable retransmissions."),
				DataType:     common.FieldDataType_INTEGER.Enum(),
				DefaultValue: common.FieldValue_builder{IntegerValue: ptr.To(int64(3))}.Build(),
				Required:     ptr.To(true),
				Validation: common.FieldValidation_builder{
					MinInteger: ptr.To(int64(0)),
					MaxInteger: ptr.To(int64(10)),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("window_size_transmit"),
				Label:        ptr.To("Window size (transmit)"),
				Tooltip:      ptr.To("Number of HDLC frame retransmissions to perform. It can be set to 0 to disable retransmissions."),
				DataType:     common.FieldDataType_INTEGER.Enum(),
				DefaultValue: common.FieldValue_builder{IntegerValue: ptr.To(int64(1))}.Build(),
				Required:     ptr.To(true),
				Validation: common.FieldValidation_builder{
					MinInteger: ptr.To(int64(1)),
					MaxInteger: ptr.To(int64(7)),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("window_size_receive"),
				Label:        ptr.To("Window size (receive)"),
				Tooltip:      ptr.To("Number of HDLC frame retransmissions to perform. It can be set to 0 to disable retransmissions."),
				DataType:     common.FieldDataType_INTEGER.Enum(),
				DefaultValue: common.FieldValue_builder{IntegerValue: ptr.To(int64(1))}.Build(),
				Required:     ptr.To(true),
				Validation: common.FieldValidation_builder{
					MinInteger: ptr.To(int64(1)),
					MaxInteger: ptr.To(int64(7)),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("max_info_field_length_transmit"),
				Label:        ptr.To("Max info-field length (transmit)"),
				Tooltip:      ptr.To("Number of HDLC frame retransmissions to perform. It can be set to 0 to disable retransmissions."),
				DataType:     common.FieldDataType_INTEGER.Enum(),
				DefaultValue: common.FieldValue_builder{IntegerValue: ptr.To(int64(2030))}.Build(),
				Required:     ptr.To(true),
				Validation: common.FieldValidation_builder{
					MinInteger: ptr.To(int64(32)),
					MaxInteger: ptr.To(int64(2030)),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("max_info_field_length_receive"),
				Label:        ptr.To("Max info-field length (receive)"),
				Tooltip:      ptr.To("Number of HDLC frame retransmissions to perform. It can be set to 0 to disable retransmissions."),
				DataType:     common.FieldDataType_INTEGER.Enum(),
				DefaultValue: common.FieldValue_builder{IntegerValue: ptr.To(int64(2030))}.Build(),
				Required:     ptr.To(true),
				Validation: common.FieldValidation_builder{
					MinInteger: ptr.To(int64(32)),
					MaxInteger: ptr.To(int64(2030)),
				}.Build(),
			}.Build(),
			common.FieldDescriptor_builder{
				FieldId:      ptr.To("adaptive_addressing"),
				Label:        ptr.To("Adaptive HDLC addressing"),
				Tooltip:      ptr.To("Disables adaptive HDLC addressing. When disabled the driver always sends 4-byte HDLC addresses."),
				DataType:     common.FieldDataType_BOOLEAN.Enum(),
				DefaultValue: common.FieldValue_builder{BoolValue: ptr.To(true)}.Build(),
				Required:     ptr.To(true),
			}.Build(),
		}

	default:
		return []*common.FieldDescriptor{}
	}
}

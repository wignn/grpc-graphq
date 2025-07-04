// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: review.proto

package genproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Review struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	AccountId     string                 `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Rating        int32                  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt     []byte                 `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Review) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Review) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetCreatedAt() []byte {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ReviewIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewIdRequest) Reset() {
	*x = ReviewIdRequest{}
	mi := &file_review_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewIdRequest) ProtoMessage() {}

func (x *ReviewIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewIdRequest.ProtoReflect.Descriptor instead.
func (*ReviewIdRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *ReviewIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProductIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Skip          uint64                 `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	Take          uint64                 `protobuf:"varint,3,opt,name=take,proto3" json:"take,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductIdRequest) Reset() {
	*x = ProductIdRequest{}
	mi := &file_review_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductIdRequest) ProtoMessage() {}

func (x *ProductIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductIdRequest.ProtoReflect.Descriptor instead.
func (*ProductIdRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *ProductIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductIdRequest) GetSkip() uint64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *ProductIdRequest) GetTake() uint64 {
	if x != nil {
		return x.Take
	}
	return 0
}

type PostReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	AccountId     string                 `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Rating        int32                  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostReviewRequest) Reset() {
	*x = PostReviewRequest{}
	mi := &file_review_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostReviewRequest) ProtoMessage() {}

func (x *PostReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostReviewRequest.ProtoReflect.Descriptor instead.
func (*PostReviewRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *PostReviewRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PostReviewRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *PostReviewRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostReviewRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type ReviewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewResponse) Reset() {
	*x = ReviewResponse{}
	mi := &file_review_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewResponse) ProtoMessage() {}

func (x *ReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewResponse.ProtoReflect.Descriptor instead.
func (*ReviewResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewResponse) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type ReviewListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reviews       []*Review              `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewListResponse) Reset() {
	*x = ReviewListResponse{}
	mi := &file_review_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewListResponse) ProtoMessage() {}

func (x *ReviewListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewListResponse.ProtoReflect.Descriptor instead.
func (*ReviewListResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{5}
}

func (x *ReviewListResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

var File_review_proto protoreflect.FileDescriptor

const file_review_proto_rawDesc = "" +
	"\n" +
	"\freview.proto\x12\bgenproto\"\xa4\x01\n" +
	"\x06Review\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1c\n" +
	"\tproductId\x18\x02 \x01(\tR\tproductId\x12\x1c\n" +
	"\taccountId\x18\x03 \x01(\tR\taccountId\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12\x16\n" +
	"\x06rating\x18\x05 \x01(\x05R\x06rating\x12\x1c\n" +
	"\tcreatedAt\x18\x06 \x01(\fR\tcreatedAt\"!\n" +
	"\x0fReviewIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"J\n" +
	"\x10ProductIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04skip\x18\x02 \x01(\x04R\x04skip\x12\x12\n" +
	"\x04take\x18\x03 \x01(\x04R\x04take\"\x81\x01\n" +
	"\x11PostReviewRequest\x12\x1c\n" +
	"\tproductId\x18\x01 \x01(\tR\tproductId\x12\x1c\n" +
	"\taccountId\x18\x02 \x01(\tR\taccountId\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x16\n" +
	"\x06rating\x18\x04 \x01(\x05R\x06rating\":\n" +
	"\x0eReviewResponse\x12(\n" +
	"\x06review\x18\x01 \x01(\v2\x10.genproto.ReviewR\x06review\"@\n" +
	"\x12ReviewListResponse\x12*\n" +
	"\areviews\x18\x01 \x03(\v2\x10.genproto.ReviewR\areviews2\xe5\x01\n" +
	"\rReviewService\x12C\n" +
	"\n" +
	"PostReview\x12\x1b.genproto.PostReviewRequest\x1a\x18.genproto.ReviewResponse\x128\n" +
	"\tGetReview\x12\x19.genproto.ReviewIdRequest\x1a\x10.genproto.Review\x12U\n" +
	"\x19GetReviewByProductAndUser\x12\x1a.genproto.ProductIdRequest\x1a\x1c.genproto.ReviewListResponseB*Z(github.com/wignn/micro-3/review/genprotob\x06proto3"

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData []byte
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_review_proto_rawDesc), len(file_review_proto_rawDesc)))
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_review_proto_goTypes = []any{
	(*Review)(nil),             // 0: genproto.Review
	(*ReviewIdRequest)(nil),    // 1: genproto.ReviewIdRequest
	(*ProductIdRequest)(nil),   // 2: genproto.ProductIdRequest
	(*PostReviewRequest)(nil),  // 3: genproto.PostReviewRequest
	(*ReviewResponse)(nil),     // 4: genproto.ReviewResponse
	(*ReviewListResponse)(nil), // 5: genproto.ReviewListResponse
}
var file_review_proto_depIdxs = []int32{
	0, // 0: genproto.ReviewResponse.review:type_name -> genproto.Review
	0, // 1: genproto.ReviewListResponse.reviews:type_name -> genproto.Review
	3, // 2: genproto.ReviewService.PostReview:input_type -> genproto.PostReviewRequest
	1, // 3: genproto.ReviewService.GetReview:input_type -> genproto.ReviewIdRequest
	2, // 4: genproto.ReviewService.GetReviewByProductAndUser:input_type -> genproto.ProductIdRequest
	4, // 5: genproto.ReviewService.PostReview:output_type -> genproto.ReviewResponse
	0, // 6: genproto.ReviewService.GetReview:output_type -> genproto.Review
	5, // 7: genproto.ReviewService.GetReviewByProductAndUser:output_type -> genproto.ReviewListResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_review_proto_rawDesc), len(file_review_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}

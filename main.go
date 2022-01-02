package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/tomill/openrtb-txtx/internal/openrtb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

var (
	input  = flag.String("in", "", "json or text")
	output = flag.String("out", "", "json or text or dump")
)

func main() {
	flag.Parse()
	msg := message(flag.Arg(0))

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	if *input == "" && strings.HasPrefix(string(b), "{") {
		*input = "json"
		if *output == "" {
			*output = "text"
		}
	} else if *input == "" {
		*input = "text"
		if *output == "" {
			*output = "json"
		}
	}

	switch *input {
	case "text":
		if err := prototext.Unmarshal(b, msg); err != nil {
			log.Fatal(err)
		}
	case "json":
		if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(b, msg); err != nil {
			log.Fatal(err)
		}
	case "dummy":
		msg = &openrtb.BidRequest_Imp{
			Id: proto.String("dummy"),
		}
	}

	var o []byte
	switch *output {
	case "text":
		o, err = prototext.MarshalOptions{Multiline: true}.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
	case "json":
		o, err = protojson.MarshalOptions{Multiline: true, UseProtoNames: true, UseEnumNumbers: true}.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
	case "dump":
		spew.Dump(msg)
		os.Exit(0)
	}

	fmt.Println(string(o))
}

func message(s string) proto.Message {
	s = strings.ToLower(s)
	var msg interface{}
	switch s {
	case "bidrequest", "req":
		msg = &openrtb.BidRequest{}
	case "source":
		msg = &openrtb.BidRequest_Source{}
	case "imp":
		msg = &openrtb.BidRequest_Imp{}
	case "metric":
		msg = &openrtb.BidRequest_Imp_Metric{}
	case "banner":
		msg = &openrtb.BidRequest_Imp_Banner{}
	case "banner.format":
		msg = &openrtb.BidRequest_Imp_Banner_Format{}
	case "video":
		msg = &openrtb.BidRequest_Imp_Video{}
	case "video.companionad":
		msg = &openrtb.BidRequest_Imp_Video_CompanionAd{}
	case "audio":
		msg = &openrtb.BidRequest_Imp_Audio{}
	case "native":
		msg = &openrtb.BidRequest_Imp_Native{}
	case "pmp":
		msg = &openrtb.BidRequest_Imp_Pmp{}
	case "deal":
		msg = &openrtb.BidRequest_Imp_Pmp_Deal{}
	case "site":
		msg = &openrtb.BidRequest_Site{}
	case "app":
		msg = &openrtb.BidRequest_App{}
	case "publisher":
		msg = &openrtb.BidRequest_Publisher{}
	case "content":
		msg = &openrtb.BidRequest_Content{}
	case "producer":
		msg = &openrtb.BidRequest_Producer{}
	case "device":
		msg = &openrtb.BidRequest_Device{}
	case "geo":
		msg = &openrtb.BidRequest_Geo{}
	case "user":
		msg = &openrtb.BidRequest_User{}
	case "data":
		msg = &openrtb.BidRequest_Data{}
	case "data.segment":
		msg = &openrtb.BidRequest_Data_Segment{}
	case "regs":
		msg = &openrtb.BidRequest_Regs{}
	case "bidresponse", "res":
		msg = &openrtb.BidResponse{}
	case "seatbid":
		msg = &openrtb.BidResponse_SeatBid{}
	case "bid":
		msg = &openrtb.BidResponse_SeatBid_Bid{}
	case "nativerequest":
		msg = &openrtb.NativeRequest{}
	case "nativerequest.asset":
		msg = &openrtb.NativeRequest_Asset{}
	case "nativerequest.asset.title":
		msg = &openrtb.NativeRequest_Asset_Title{}
	case "nativerequest.asset.image":
		msg = &openrtb.NativeRequest_Asset_Image{}
	case "nativerequest.asset.data":
		msg = &openrtb.NativeRequest_Asset_Data{}
	case "nativerequest.eventtrackers":
		msg = &openrtb.NativeRequest_EventTrackers{}
	case "nativeresponse":
		msg = &openrtb.NativeResponse{}
	case "nativeresponse.link":
		msg = &openrtb.NativeResponse_Link{}
	case "nativeresponse.asset":
		msg = &openrtb.NativeResponse_Asset{}
	case "nativeresponse.asset.title":
		msg = &openrtb.NativeResponse_Asset_Title{}
	case "nativeresponse.asset.image":
		msg = &openrtb.NativeResponse_Asset_Image{}
	case "nativeresponse.asset.data":
		msg = &openrtb.NativeResponse_Asset_Data{}
	case "nativeresponse.asset.video":
		msg = &openrtb.NativeResponse_Asset_Video{}
	case "nativeresponse.eventtracker":
		msg = &openrtb.NativeResponse_EventTracker{}
	default:
		return &openrtb.BidRequest{}
	}

	return (msg).(proto.Message)
}

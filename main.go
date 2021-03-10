package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/tomill/openrtb-txtx/internal/adx"
	"github.com/tomill/openrtb-txtx/internal/openrtb"
)

var (
	_      adx.BidRequestExt
	input  = flag.String("in", "", "json or text")
	output = flag.String("out", "", "json or text")
)

func main() {
	flag.Parse()
	msg := message(flag.Arg(0))

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)

	if *input == "" && strings.HasPrefix(s, "{") {
		*input = "json"
		*output = "text"
	} else if *input == "" {
		*input = "text"
		*output = "json"
	}

	switch *input {
	case "text":
		if err := proto.UnmarshalText(s, msg); err != nil {
			log.Fatal(err)
		}
	case "json":
		if err := jsonpb.UnmarshalString(s, msg); err != nil {
			log.Fatal(err)
		}
	}

	switch *output {
	case "text":
		if err := proto.MarshalText(os.Stdout, msg); err != nil {
			log.Fatal(err)
		}
	case "json":
		if err := (&jsonpb.Marshaler{Indent: "  "}).Marshal(os.Stdout, msg); err != nil {
			log.Fatal(err)
		}
	}
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

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/voyagegroup/fluct_pb/go/adx"
	"github.com/voyagegroup/fluct_pb/go/openrtb"
)

var (
	_ adx.BidRequestExt
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)

	var msg interface{}
	if strings.Contains(s, `,"imp":`) {
		msg = &openrtb.BidRequest{}
	} else if strings.Contains(s, `,"seatbid":`) {
		msg = &openrtb.BidResponse{}
	}

	var data = (msg).(proto.Message)
	if err := jsonpb.UnmarshalString(s, data); err != nil {
		log.Fatal(err)
	}
	if err := proto.MarshalText(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

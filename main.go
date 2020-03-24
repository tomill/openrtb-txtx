package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/voyagegroup/fluct_pb/go/adx"
	"github.com/voyagegroup/fluct_pb/go/openrtb"
)

var (
	_ adx.BidRequestExt
	_ adx.ImpExt
	_ adx.AppExt
	_ adx.SiteExt
	_ adx.UserExt
	_ adx.BidExt
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var msg openrtb.BidRequest
	if err := json.Unmarshal(b, &msg); err != nil {
		log.Fatal(err)
	}
	if msg.String() == "" {
		log.Fatal("Unmarshal failedddd")
	}
	//
	//x, err := proto.GetExtension(&msg, adx.E_BidRequest)
	//fmt.Printf("--\n%#v\n", x)
	//
	//if err := proto.MarshalText(os.Stdout, &msg); err != nil {
	//	log.Fatal(err)
	//}
}

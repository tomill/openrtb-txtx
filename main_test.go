package main

import (
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/voyagegroup/fluct_pb/go/adx"
	"github.com/voyagegroup/fluct_pb/go/openrtb"
)

var _ adx.ImpExt

func TestFoo(t *testing.T) {

	input := `{"id":"res-id","[com.google.doubleclick.imp]":{"dfp_ad_unit_code":"foo"}}`
	imp := &openrtb.BidRequest_Imp{}

	err := jsonpb.UnmarshalString(input, imp)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", imp)

	impExt := getImpExt(imp)

	t.Logf("%#v", impExt.GetDfpAdUnitCode())

}

func getImpExt(imp *openrtb.BidRequest_Imp) *adx.ImpExt {
	m, _ := proto.GetExtension(imp, adx.E_Imp)
	if ext, ok := m.(*adx.ImpExt); ok {
		return ext
	}

	return &adx.ImpExt{}
}

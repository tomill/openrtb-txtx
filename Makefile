
install: internal
	go install

deps:
	go get google.golang.org/protobuf/cmd/protoc-gen-go

internal: deps openrtb.proto openrtb-adx.proto
	mkdir -p internal/openrtb
	protoc --go_out=. --go_opt=Mopenrtb-adx.proto=internal/openrtb,Mopenrtb.proto=internal/openrtb openrtb.proto openrtb-adx.proto

test:
	bash -c "diff -u testdata/imp.json <(cat testdata/imp.text | go run main.go imp)"
	bash -c "diff -u testdata/imp.text <(cat testdata/imp.json | go run main.go imp)"

.PHONY: *.proto

openrtb.proto:
	echo "// @ref https://github.com/google/openrtb/" > $@
	curl -sL https://raw.githubusercontent.com/google/openrtb/master/openrtb-core/src/main/protobuf/openrtb.proto >> $@

openrtb-adx.proto:
	echo "// @ref https://developers.google.com/authorized-buyers/rtb/data" > $@
	curl -sL https://developers.google.com/authorized-buyers/rtb/downloads/openrtb-adx-proto.txt >> $@

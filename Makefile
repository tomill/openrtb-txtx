
build: gen
	go build

gen: openrtb.proto openrtb-adx.proto
	mkdir -p internal/openrtb internal/adx
	protoc --go_out=import_path=openrtb:internal/openrtb openrtb.proto
	protoc --go_out=import_path=adx,Mopenrtb.proto=github.com/tomill/openrtb-txtx/internal/openrtb:internal/adx openrtb-adx.proto

.PHONY: *.proto

openrtb.proto:
	echo "// @ref https://github.com/google/openrtb/" > $@
	curl -sL https://raw.githubusercontent.com/google/openrtb/master/openrtb-core/src/main/protobuf/openrtb.proto >> $@

openrtb-adx.proto:
	echo "// @ref https://developers.google.com/authorized-buyers/rtb/data" > $@
	curl -sL https://developers.google.com/authorized-buyers/rtb/downloads/openrtb-adx-proto.txt >> $@

test:
	diff -u <(cat testdata/imp.text | go run main.go -in text -out json imp) testdata/imp.json
	diff -u <(cat testdata/imp.json | go run main.go -in json -out text imp) testdata/imp.text

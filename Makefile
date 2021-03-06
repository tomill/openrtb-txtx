
install: internal
	go install

build: internal
	go build

internal: openrtb.proto openrtb-adx.proto
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
	bash -c "diff --strip-trailing-cr -u <(cat testdata/imp.text | go run main.go imp) testdata/imp.json"
	bash -c "diff --strip-trailing-cr -u <(cat testdata/imp.json | go run main.go imp) testdata/imp.text"

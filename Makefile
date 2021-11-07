
install: internal
	go install

build: internal
	go build

internal: openrtb.proto openrtb-adx.proto
	mkdir -p internal/openrtb internal/adx
	protoc --go_out=. --go_opt=Mopenrtb.proto=internal/openrtb openrtb.proto
	protoc --go_out=. --go_opt=Mopenrtb-adx.proto=internal/adx --go_opt=Mopenrtb.proto=github.com/tomill/openrtb-txtx/internal/openrtb openrtb-adx.proto

.PHONY: *.proto

openrtb.proto:
	echo "// @ref https://github.com/google/openrtb/" > $@
	curl -sL https://raw.githubusercontent.com/google/openrtb/master/openrtb-core/src/main/protobuf/openrtb.proto >> $@

openrtb-adx.proto:
	echo "// @ref https://developers.google.com/authorized-buyers/rtb/data" > $@
	curl -sL https://developers.google.com/authorized-buyers/rtb/downloads/openrtb-adx-proto.txt >> $@

test:
	bash -c "diff -u testdata/imp.json <(cat testdata/imp.text | go run main.go imp)"
	bash -c "diff -u testdata/imp.text <(cat testdata/imp.json | go run main.go imp)"

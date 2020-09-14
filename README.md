# openrtb-txtx

protobuf text format <=> json for OpenRTB + AdX Ext

## Install

```
make install
```

## Usage

```
cat req.json | openrtb-txtx --out text

cat req.bind | protoc -I pb/ --decode_raw | openrtb-txtx --out json
```
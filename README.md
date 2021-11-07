# openrtb-txtx

"protobuf text format" <=> json converter for [Google OpenRTB + AdX Ext](https://developers.google.com/authorized-buyers/rtb/data?hl=en#proto-)

## Install

```
make install
```

## Usage

```
openrtb-txtx [--in ...] [--out ...] [message name]

  --in string json or text
  --out string json or text or dump
  
  message name: default is "BidRequest"
```

### Examples

json to "protobuf text format" (message type is BidRequest (default))

```
$ cat req.json | openrtb-txtx req

id:  "xxxxxxxxxxxxxxxxxxxxxx"
imp:  {
  id:  "1"
  banner:  {
    w: 300
    h: 250
    format:  {
      w: 300
      h: 250
    }
```

"protobuf text format" to json

```
$ cat req.bin | protoc -I pb/ --decode_raw | openrtb-txtx

{
  "id": "xxxxxxxxxxxxxxxxxxxxxx",
  "imp": [
    {
      "id": "1",
      "banner": {
        "w": 300,
        "h": 250,
        "format": [
          {
            "w": 300,
            "h": 250
          },
```



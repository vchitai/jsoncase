# jsoncase
a go package for converting json case to various cases 

[![Godoc Reference](https://godoc.org/github.com/vchitai/jsoncase?status.svg)](http://godoc.org/github.com/vchitai/jsoncase)
[![Coverage](http://gocover.io/_badge/github.com/vchitai/jsoncase?0)](http://gocover.io/github.com/vchitai/jsoncase)
[![Go Report Card](https://goreportcard.com/badge/github.com/vchitai/jsoncase)](https://goreportcard.com/report/github.com/vchitai/jsoncase)

## Example

```go
s := `
{
  "a_b": {
    "c_d": "e"
  },
  "b_e": "x",
  "e_f": [
    {
      "c_f": "e"
    }
  ],
  "x_y": {
    "a_b": {
      "c_d": "e"
    },
    "b_e": "x",
    "e_f": [
      {
        "c_f": "e"
      }
    ]
  }
}
`
s2 := ToCamelJson(s)
fmt.Prilntln(s2)
=> `
{
  "aB": {
    "cD": "e"
  },
  "bE": "x",
  "eF": [
    {
      "cF": "e"
    }
  ],
  "xY": {
    "aB": {
      "cD": "e"
    },
    "bE": "x",
    "eF": [
      {
        "cF": "e"
      }
    ]
  }
}
`
```

| Function                                  | Usage               |
|-------------------------------------------|----------------------|
| `ToSnakeJson`                              | string-string |
| `ToCamelJson`                              | string-string |
| `ToSnakeJsonBytes`                              | []byte-[]byte |
| `ToCamelJsonBytes`                              | []byte-[]byte |
| `ToSnakeVal`                     | interface{} - interface{} |
| `ToCamelVal`                     | interface{} - interface{} |



## Install

```bash
go get -u https://github.com/vchitai/jsoncase
```
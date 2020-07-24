# jsoncase
a go package for converting json case to various cases 

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
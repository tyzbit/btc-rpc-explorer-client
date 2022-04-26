# go-btc-rpc-explorer-api
Go wrapper around the btc-rpc-explorer API

## Usage Example

```golang
func main() {
  address := "34rng4QwB5pHUbGDJw1JxjLwgEU8TQuEqv"
	btcrpc := Config{
		APIEndpoint: TestingEndpoint,
	}
	summary, err := btcrpc.AddressSummary(address)
	if err != nil {
		fmt.Println(err)
	}
  fmt.Println("the balance of ", address, " is ", summary.TXHistory.BalanceSat, " satoshis")
}
```
# btcapi
Go wrapper around the btc-rpc-explorer API

## Usage Example

```golang
import "github.com/tyzbit/btcapi"

func main() {
    address := "34rng4QwB5pHUbGDJw1JxjLwgEU8TQuEqv"
    btc := btcapi.Config{
        APIEndpoint: "https://btc-rpc-explorer.example.com",
    }
    summary, err := btc.AddressSummary(address)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("the balance of ", address, " is ", summary.TXHistory.BalanceSat, " satoshis")
}
```
# go-optional-setter
This pkg is setter/getter, each values.

## Usage
``` go
type Data struct {
    Attack optional.Int
    Magic optional.Int
}

func main() {
    data := &Data{}
    data.Attack.Set(10)

    fmt.Println(data.Attack.Value())
    // -> 10

    fmt.Println(data.Attack.IsSet())
    // -> true
    fmt.Println(data.Magic.IsSet())
    // -> false
}
```
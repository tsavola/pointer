```go
package motivationalexample

import (
	"encoding/json"

	"github.com/tsavola/pointer"
)

const (
	defaultName = "some name"
)

// Stuff contains optional fields.
type Stuff struct {
	Name    *string
	Comment *string
	Value   *int64
}

// SomeStuff makes some JSON-encoded stuff.
func SomeStuff() (data []byte, err error) {
	return json.Marshal(&Stuff{
		Name:    pointer.String(defaultName), // can't say &defaultName
		Comment: pointer.String("not yet"),   // can't say &"not yet"
		Value:   pointer.Int64(42),           // can't say &42 or &int64(42)
	})
}
```

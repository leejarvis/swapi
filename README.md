swapi
=====

API client for http://swapi.co/

Example
-------

```go
package main

import (
	"fmt"

	"github.com/leejarvis/swapi"
)

func main() {
	if s, err := swapi.GetSpecies(3); err == nil {
		fmt.Println(s.Name)
	}
}
```

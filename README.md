# Post to Google Forms

## Usage

From [example/](https://github.com/fawkesley/googleforms/blob/main/example/main.go):

```go
package main

import (
	"fmt"
	"time"

	"github.com/fawkesley/googleforms"
)

func main() {
	form := googleforms.Form{
		FormID: "1FAIpQLScKpe8JnLk5yZoBTRCMjFO7-ToKoIsXswiU05jIVz_iPTu52g",
		Questions: map[string]string{
			"email":   "entry.577692960",
			"comment": "entry.188106111",
		},
	}

	err := form.Post(map[string]string{
		"email":   "test@example.com",
		"comment": fmt.Sprintf("Hi, the time is %s", time.Now()),
	})
	if err != nil {
		panic(err)
	}
}
```



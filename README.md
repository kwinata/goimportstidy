# goimportstidy

This tool updates your Go import lines, grouping it into three groups: 
 - stdlib,
 - external libraries,
 - local libraries (optional).
 
Installation: 

     $ go install github.com/kwinata/goimportstidy@latest
     
Usage:

    $ goimportstidy -w -local github.com/shipwallet main.go -current github.com/shipwallet/core .

Difference from the original repo is we can also support `current` repo, to be 4 groups, example:

Input: 

```
import (
	"fmt"

	"github.com/krzysztofdrys/a"

	"github.com/krzysztofdrys/a/core"

	"github.com/krzysztofdrys/a/something/else"

	"github.com/krzysztofdrys/b"

	"github.com/krzysztofdrys/c"
	"gitlab.com/krzysztofdrys/a"

	"gitlab.com/krzysztofdrys/b"
	"gitlab.com/krzysztofdrys/b"
)
```

Output with `-local github.com/krzysztofdrys -current github.com/krzysztofdrys/a`:

```
import (
	"fmt"

	"gitlab.com/krzysztofdrys/a"
	"gitlab.com/krzysztofdrys/b"
	"gitlab.com/krzysztofdrys/b"

	"github.com/krzysztofdrys/b"
	"github.com/krzysztofdrys/c"

	"github.com/krzysztofdrys/a"
	"github.com/krzysztofdrys/a/core"
	"github.com/krzysztofdrys/a/something/else"
)

```
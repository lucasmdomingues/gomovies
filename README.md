# Installation

```go 
go get github.com/lucasmdomingues/gomovies
```

# Examples

### Search by title

```go
import (
	"log"

	"github.com/lucasmdomingues/gomovies"
)

func main() {
	apiKey := "..."

	movie, err := gomovies.SearchByTitle(apiKey, "Lord of the rings")
	if err != nil {
        log.Fatal(err)
	}
}
```

### Search by ID

```go
import (
	"log"

	"github.com/lucasmdomingues/gomovies"
)

func main() {
	apiKey := "..."

	movie, err := gomovies.SearchByID(apiKey, "tt0120737")
	if err != nil {
        log.Fatal(err)
	}
} 
```

### Omdb Movies
http://www.omdbapi.com/

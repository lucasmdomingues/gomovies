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
	service := gomovies.NewService("API_KEY")

	_, err := service.SearchByTitle("Lord of the rings")
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
	service := gomovies.NewService("API_KEY")

	_, err := service.SearchByID("tt0120737")
	if err != nil {
		log.Fatal(err)
	}
} 
```

### Omdb Movies
http://www.omdbapi.com/

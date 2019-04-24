# Instalação

```go 
go get github.com/lucasmdomingues/gomovies
```

# Exemplos

### Busca por ID ou Título

```go
import (
	"fmt"

	"github.com/lucasmdomingues/gomovies"
)

func main() {

	key := "KEY"

	movie, err := gomovies.GetMovie(key, "Lord of the rings", "tt0120737")
	if err != nil {
        fmt.Println(err)
        return
	}

	fmt.Println(movie)
}
```

### Omdb Movies
http://www.omdbapi.com/
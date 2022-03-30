# dotenv

Loads the env file in a Golang map

> Big O notation: O(n)

## Usage

```bash
# .env

NAME=SomeName wrong_value # some comment

KEY="--- Your key here ---" # secret key

# empty variables
EMPTY_VARIABLE=
```

```golang
// config/env.go
import "github.com/marlonmp/dotenv"

var env map[string]string

// loads the env file wherever you want
func init() {
    envPath := "path/to/.env/file"

    dotenv.LoadFile(envPath, &env)
}

// create your own `GetEnv`
func Env(idx string) string {
    return env[idx]
}

// main.go

import "packageName/config"

func main() {
    name := config.Env("NAME")

    println("Hello", name) // Hello SomeName
}

```

Happy hacking ;)
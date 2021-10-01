# go_envs

## Usage

- ```go
    package main

    import "github.com/m4rshmallow/go_envs"


    func main() {
        // by default load .env file
        go_envs.Init()
        fmt.Println(os.Getenv("SERVER_PORT"))
        fmt.Println(go_envs.Map.Get("API_PORT"))
    }
  ```

- ```go
    package main

    import (
            "fmt"
            "os"

            "github.com/m4rshmallow/go_envs"
        )

        func main() {
            go_envs.Init(".env", ".env_extra")
            fmt.Println(os.Getenv("API_PORT"))
            fmt.Println(go_envs.Map.Get("API_PORT"))
        }
  ```

- ```go
    package main
    
    import (
            "fmt"
            "os"

            "github.com/m4rshmallow/go_envs"
            _ "github.com/m4rshmallow/go_envs/autoload"
        )


        func main() {
            fmt.Println(go_envs.Map.Get("API_PORT"))
            fmt.Println(os.Getenv("API_PORT"))
        }
  ```

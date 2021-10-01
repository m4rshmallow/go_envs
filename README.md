# go_envs

## Usage

- ```go
   import "github.com/m4rshmallow/go_envs"


   func main() {
       // by default load .env file
       go_envs.Init()
       fmt.Println(os.Getenv("SERVER_PORT"))
   }
  ```

- ```go
  import "github.com/m4rshmallow/go_envs"


  func main() {
      go_envs.Init(".env", ".some_another_env")
      fmt.Println(os.Getenv("SERVER_PORT"))
  }
  ```

- ```go
   import (
       _ "github.com/m4rshmallow/go_envs/autoload"
       "fmt"
   )


   func main() {
       fmt.Println(go_envs.EnvMap.Get("SERVER_PORT"))
       fmt.Println(os.Getenv("SERVER_PORT"))
   }
  ```

# goenvs

## Usage

- ```go
   import "envs"


   func main() {
       // by default load .env file
       envs.Init()
       fmt.Println(os.Getenv("SERVER_PORT"))
   }
  ```

- ```go
  import "envs"


  func main() {
      envs.Init(".env", ".some_another_env")
      fmt.Println(os.Getenv("SERVER_PORT"))
  }
  ```

- ```go
   import (
       _ "envs/autoload"
       "fmt"
   )


   func main() {
       fmt.Println(envs.EnvMap.Get("SERVER_PORT"))
       fmt.Println(os.Getenv("SERVER_PORT"))
   }
  ```

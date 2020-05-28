# Compression

![Test](https://github.com/gofiber/compression/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/compression/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/compression/workflows/Linter/badge.svg)

### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/compression
```
### Example
```go
package main

import 
  "github.com/gofiber/fiber"
  "github.com/gofiber/compression"
)

func main() {
  app := fiber.New()

  app.Use(compression.New())

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Welcome!")
  })

  app.Listen(3000)
}
```

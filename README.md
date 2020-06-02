# Compression

![Release](https://img.shields.io/github/release/gofiber/compression.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
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
### Config
| Property | Type | Description | Default |
| :--- | :--- | :--- | :--- |
| Filter | `func(*fiber.Ctx) bool` | Defines a function to skip middleware | `nil` |
| Level | `int` | Level of compression | `0` (_Default GZIP compression_) |

### Levels
| Value | Description |
| :--- | :--- |
| `-1` | No compression |
| `0` | Default GZIP compression |
| `1` | Best GZIP speed |
| `2` | Best GZIP compression |
| `3` | GZIP HuffmanOnly |
| `4` | Default Brotli compression |
| `5` | Best Brotli Speed |
| `6` | Best Brotli Compression |
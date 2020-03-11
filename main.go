// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package compression

import (
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp"
)

// Supported compression levels
const (
	LevelNoCompression      = -1
	LevelDefaultCompression = 0
	LevelBestSpeed          = 1
	LevelBestCompression    = 2
	LevelHuffmanOnly        = 3
)

// Config ...
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// Level of compression
	// Optional. Default value 0.
	Level int
}

// New ...
func New(config ...Config) func(*fiber.Ctx) {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	// Convert compress levels to correct int
	// https://github.com/valyala/fasthttp/blob/master/compress.go#L17
	switch cfg.Level {
	case -1:
		cfg.Level = 0
	case 1:
		cfg.Level = 1
	case 2:
		cfg.Level = 9
	case 3:
		cfg.Level = -2
	default:
		cfg.Level = 6
	}
	compress := fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {
		return
	}, cfg.Level)
	// Middleware function
	return func(c *fiber.Ctx) {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
		c.Next()
		compress(c.Fasthttp)
	}
}

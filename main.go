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
	LevelNoCompression         = -1
	LevelDefaultCompression    = 0
	LevelBestSpeed             = 1
	LevelBestCompression       = 2
	LevelHuffmanOnly           = 3
	LevelDefaultBrotli         = 4
	LevelBrotliBestSpeed       = 5
	LevelBrotliBestCompression = 6
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
	var compress fasthttp.RequestHandler
	switch cfg.Level {
	case -1:
		compress = fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {}, 0)
	case 1:
		compress = fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {}, 1)
	case 2:
		compress = fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {}, 9)
	case 3:
		compress = fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {}, -2)
	case 4:
		compress = fasthttp.CompressHandlerBrotliLevel(func(c *fasthttp.RequestCtx) {}, 4, 6)
	case 5:
		compress = fasthttp.CompressHandlerBrotliLevel(func(c *fasthttp.RequestCtx) {}, 0, 6)
	case 6:
		compress = fasthttp.CompressHandlerBrotliLevel(func(c *fasthttp.RequestCtx) {}, 11, 6)
	default:
		compress = fasthttp.CompressHandlerLevel(func(c *fasthttp.RequestCtx) {}, 6)
	}
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

package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"time"
)

// usage:
// http request: GET /?name=example.com&type=A&encode=meme
// take a look at the Cloudflare DoH documentation: https://developers.cloudflare.com/1.1.1.1/encrypted-dns/dns-over-https/make-api-requests/dns-json
// encoders: plain, base64, meme (LSB steganography)
// set PORT environment variable to change to port (default: 3000)
func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// get query
		name := c.Query("name")
		queryType := c.Query("type")
		encoderName := c.Query("encoder", "plain")

		encoder, ok := encoders[encoderName]
		if !ok {
			return c.Status(http.StatusBadRequest).SendString("unknown encoder")
		}

		// resolve
		start := time.Now()
		data, status, err := resolve(name, queryType)
		if err != nil {
			log.Println("resolving failed", name, queryType, err)
			return c.Status(http.StatusInternalServerError).SendString("resolve failed")
		}
		resolveDuration := time.Since(start)

		// encode
		encoded, err := encoder.Encode(data)
		if err != nil {
			log.Println("failed to encode", encoderName, err)
			return c.Status(http.StatusInternalServerError).SendString("encoder failed")
		}

		log.Println("resolved", name, queryType, status, resolveDuration, len(data))

		return c.Status(status).Send(encoded)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("listening on port", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("listen failed", err)
	}
}

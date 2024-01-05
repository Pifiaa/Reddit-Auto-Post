package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	// Create a channel to handle shutdown signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Start the Fiber server in a goroutine
	go func() {
		// Initialize progress bar
		bar := progressbar.NewOptions(100, progressbar.OptionSetWriter(os.Stdout))

		// Start the Fiber server
		app := fiber.New()

		// Your Fiber routes and configurations go here
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		// Run Fiber server in a goroutine
		go func() {
			err := app.Listen(":3000")
			if err != nil {
				log.Fatal(err)
			}
		}()

		// Update progress bar until server is ready
		for i := 0; i < 100; i++ {
			time.Sleep(50 * time.Millisecond) // Simulate server startup delay
			bar.Add(1)
		}

		fmt.Println("\nFiber server is ready!")

		// Block until the server is shut down
		<-shutdown

		// Gracefully shut down the Fiber server
		err := app.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for shutdown signal
	<-shutdown
	fmt.Println("Shutting down...")
}

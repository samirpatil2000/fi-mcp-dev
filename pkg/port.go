package pkg

import (
	"log"
	"os"
)

func GetPort() string {
	port := os.Getenv("FI_MCP_PORT")
	if port != "" {
		return port
	}
	log.Println("Warning: FI_MCP_PORT environment variable not set, using default port 8080")
	return "8080"
}

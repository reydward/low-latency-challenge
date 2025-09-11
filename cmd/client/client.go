package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Single test
	latency := measureLatency()
	fmt.Printf("Latencia: %v\n", latency)

	// Multiple tests for statistical analysis
	fmt.Println("\n--- Pruebas múltiples ---")
	var total time.Duration
	var min, max time.Duration = time.Hour, 0

	for i := 0; i < 10; i++ {
		lat := measureLatency()
		total += lat
		if lat < min {
			min = lat
		}
		if lat > max {
			max = lat
		}
		fmt.Printf("Test %d: %v\n", i+1, lat)
	}

	avg := total / 10
	fmt.Printf("\nEstadísticas:\n")
	fmt.Printf("Promedio: %v\n", avg)
	fmt.Printf("Mínimo: %v\n", min)
	fmt.Printf("Máximo: %v\n", max)
	fmt.Printf("Target <1ms: %t\n", avg < time.Millisecond)
}

func measureLatency() time.Duration {
	start := time.Now()

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Error conectando: %v\n", err)
		return 0
	}
	defer conn.Close()

	// Send stimulus
	conn.Write([]byte("test"))

	// Read response
	buffer := make([]byte, 64)
	conn.Read(buffer)

	return time.Since(start)
}
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// Setup logging
	logFile, err := os.OpenFile("latency.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error abriendo archivo de log:", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	// Single test
	latency := measureLatency()
	fmt.Printf("Latencia: %v\n", latency)
	logger.Printf("Latencia individual: %v", latency)

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
		logger.Printf("Test %d: %v", i+1, lat)
	}

	avg := total / 10
	fmt.Printf("\nEstadísticas:\n")
	fmt.Printf("Promedio: %v\n", avg)
	fmt.Printf("Mínimo: %v\n", min)
	fmt.Printf("Máximo: %v\n", max)
	fmt.Printf("Target <1ms: %t\n", avg < time.Millisecond)

	// Log statistics
	logger.Printf("=== ESTADÍSTICAS ===")
	logger.Printf("Promedio: %v", avg)
	logger.Printf("Mínimo: %v", min)
	logger.Printf("Máximo: %v", max)
	logger.Printf("Target <1ms: %t", avg < time.Millisecond)
	logger.Printf("===================")
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

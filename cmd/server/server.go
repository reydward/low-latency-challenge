package main

import (
	"net"
	"runtime"
	"runtime/debug"
)

func main() {
	// Ultra-low latency optimizations
	runtime.GOMAXPROCS(1)  // Single thread for predictable performance
	runtime.GC()           // Initial GC cleanup
	debug.SetGCPercent(-1) // Disable GC during operation

	// Pre-allocated response to avoid allocations
	response := []byte("respuesta")
	buffer := make([]byte, 64) // Reusable buffer

	// Listen on localhost for minimal network latency
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		// Handle connection inline (no goroutine overhead)
		conn.Read(buffer)    // Read stimulus (ignore content)
		conn.Write(response) // Send fixed response
		conn.Close()
	}
}

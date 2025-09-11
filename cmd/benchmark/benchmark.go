package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚀 Ultra-Low Latency Benchmark")
	fmt.Println("==============================")

	// Warm up connection
	warmUp()

	// Single connection benchmark
	fmt.Println("\n📊 Single Connection Test:")
	singleConnectionBenchmark()

	// Concurrent connections benchmark
	fmt.Println("\n🔥 Concurrent Connections Test:")
	concurrentBenchmark(100) // 100 concurrent connections
}

func warmUp() {
	fmt.Println("🔧 Warming up connection...")
	for i := 0; i < 5; i++ {
		measureLatency()
	}
}

func singleConnectionBenchmark() {
	iterations := 1000
	var total time.Duration
	var min, max time.Duration = time.Hour, 0
	under1ms := 0

	for i := 0; i < iterations; i++ {
		lat := measureLatency()
		total += lat
		
		if lat < min {
			min = lat
		}
		if lat > max {
			max = lat
		}
		if lat < time.Millisecond {
			under1ms++
		}
	}

	avg := total / time.Duration(iterations)
	successRate := float64(under1ms) / float64(iterations) * 100

	fmt.Printf("Iteraciones: %d\n", iterations)
	fmt.Printf("Promedio: %v\n", avg)
	fmt.Printf("Mínimo: %v\n", min)
	fmt.Printf("Máximo: %v\n", max)
	fmt.Printf("< 1ms: %.1f%%\n", successRate)
	fmt.Printf("Target achieved: %t\n", avg < time.Millisecond)
}

func concurrentBenchmark(connections int) {
	var wg sync.WaitGroup
	results := make(chan time.Duration, connections)

	start := time.Now()

	for i := 0; i < connections; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lat := measureLatency()
			results <- lat
		}()
	}

	wg.Wait()
	close(results)

	totalDuration := time.Since(start)

	var total time.Duration
	var min, max time.Duration = time.Hour, 0
	under1ms := 0
	count := 0

	for lat := range results {
		total += lat
		count++
		
		if lat < min {
			min = lat
		}
		if lat > max {
			max = lat
		}
		if lat < time.Millisecond {
			under1ms++
		}
	}

	avg := total / time.Duration(count)
	successRate := float64(under1ms) / float64(count) * 100
	throughput := float64(count) / totalDuration.Seconds()

	fmt.Printf("Conexiones concurrentes: %d\n", connections)
	fmt.Printf("Tiempo total: %v\n", totalDuration)
	fmt.Printf("Throughput: %.0f req/s\n", throughput)
	fmt.Printf("Latencia promedio: %v\n", avg)
	fmt.Printf("< 1ms: %.1f%%\n", successRate)
}

func measureLatency() time.Duration {
	start := time.Now()

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return 0
	}
	defer conn.Close()

	conn.Write([]byte("test"))
	buffer := make([]byte, 64)
	conn.Read(buffer)

	return time.Since(start)
}
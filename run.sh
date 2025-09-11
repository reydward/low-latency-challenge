#!/bin/bash

echo "🚀 Ultra-Low Latency System"
echo "=========================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go no está instalado. Instala Go desde https://golang.org/"
    exit 1
fi

echo "✅ Go version: $(go version)"

# Initialize Go module if not exists
if [ ! -f "go.mod" ]; then
    echo "📦 Inicializando módulo Go..."
    go mod init low-latency-challenge
fi

# Build server
echo "🔨 Compilando servidor..."
go build -o bin/server ./cmd/server
if [ $? -ne 0 ]; then
    echo "❌ Error compilando servidor"
    exit 1
fi

# Build client
echo "🔨 Compilando cliente..."
go build -o bin/client ./cmd/client
if [ $? -ne 0 ]; then
    echo "❌ Error compilando cliente"
    exit 1
fi

# Build benchmark
echo "🔨 Compilando benchmark..."
go build -o bin/benchmark ./cmd/benchmark
if [ $? -ne 0 ]; then
    echo "❌ Error compilando benchmark"
    exit 1
fi

echo ""
echo "✅ Compilación exitosa!"
echo ""
echo "📋 Comandos disponibles:"
echo "  ./bin/server     - Ejecutar servidor"
echo "  ./bin/client     - Probar latencia"
echo "  ./bin/benchmark  - Benchmark completo"
echo ""
echo "🚀 Para empezar:"
echo "  1. Terminal 1: ./bin/server"
echo "  2. Terminal 2: ./bin/client"
echo ""

# Make files executable
chmod +x bin/server bin/client bin/benchmark
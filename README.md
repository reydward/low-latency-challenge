# ⚡ Ultra-Low Latency Challenge

Sistema de respuesta ultra-rápida diseñado para lograr latencias menores a 1 milisegundo.

## 🎯 Objetivo

Diseñar y crear un sistema en el que, ante un estímulo (por ejemplo, un mensaje cualquiera), el sistema responda con otro mensaje (por ejemplo, “respuesta”) en el menor tiempo posible. El desafío es lograr una latencia (tiempo entre enviar y recibir respuesta) mínima, preferiblemente menor a un milisegundo..

## 🏗️ Arquitectura

```
Cliente ──localhost──► Servidor Go ──► "respuesta"
         (127.0.0.1)    (puerto 8080)
```

### Estructura del Proyecto

```
/cmd/
├── server/server.go      # Servidor ultra-optimizado
├── client/client.go      # Cliente de prueba
└── benchmark/benchmark.go # Benchmark completo

/bin/
├── server               # Binario del servidor
├── client               # Binario del cliente  
└── benchmark            # Binario del benchmark
```

### Optimizaciones Aplicadas

- **Loopback Network**: Sin latencia de red física
- **Single Thread**: GOMAXPROCS=1 para predictibilidad
- **GC Disabled**: Sin pausas de garbage collection
- **Pre-allocated Buffers**: Sin allocaciones durante operación
- **Inline Processing**: Sin overhead de goroutines

## 🚀 Uso

```bash
# 1. Preparar sistema
chmod +x run.sh
./run.sh

# 2. Ejecutar servidor (Terminal 1)
./bin/server

# 3. Probar latencia (Terminal 2)
./bin/client

# 4. Benchmark completo
./bin/benchmark
```

## 📊 Resultados Esperados

- **Latencia típica**: 130-180μs  
- **Target**: <1ms (1000μs) ✅
- **Success rate**: >99%
- **Throughput**: ~20K req/s concurrente

## 📁 Estructura de Archivos

### Código Fuente
- `cmd/server/server.go` - Servidor ultra-optimizado
- `cmd/client/client.go` - Cliente de prueba con estadísticas  
- `cmd/benchmark/benchmark.go` - Benchmark completo con concurrencia

### Binarios (generados)
- `bin/server` - Servidor ejecutable
- `bin/client` - Cliente ejecutable
- `bin/benchmark` - Benchmark ejecutable

### Configuración
- `run.sh` - Script de compilación y ejecución
- `go.mod` - Configuración del módulo Go

## 📈 Monitoreo

El sistema incluye métricas automáticas:
- Latencia promedio, mínima y máxima
- Porcentaje de respuestas <1ms
- Throughput(peticiones por unidad de tiempo) en requests/segundo
- Pruebas de concurrencia

## 🔧 Optimizaciones Avanzadas

Para latencias aún menores, considera:

1. **Unix Domain Sockets** en lugar de TCP
2. **CPU Affinity** para cores dedicados  
3. **Real-time scheduling** con `SCHED_FIFO`
4. **Memory locking** con `mlock()`
5. **DPDK** para bypass del kernel de red

## 🏃‍♂️ Inicio Rápido

```bash
# Clonar y ejecutar en un comando
git clone <repo> && cd low-latency-challenge && ./run.sh

# En terminal separada
./bin/server &
./bin/client
```

## ⚠️ Consideraciones

Este sistema está optimizado para **latencia ultra-baja** sacrificando:
- Manejo de errores robusto
- Escalabilidad horizontal
- Funcionalidades avanzadas
- Seguridad completa


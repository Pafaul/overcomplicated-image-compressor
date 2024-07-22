package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"runtime"
	"time"
)

var start = time.Now()

const MB float64 = 1.0 * 1024 * 1024

func HealthStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(getHealthStats())
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

type HealthStats struct {
	Uptime               int64   `json:"uptime"`
	AllocatedMemory      float64 `json:"allocatedMemory"`
	TotalAllocatedMemory float64 `json:"totalAllocatedMemory"`
	Goroutines           int     `json:"goroutines"`
	GCCycles             uint32  `json:"completedGCCycles"`
	NumberOfCPUs         int     `json:"cpus"`
	HeapSys              float64 `json:"maxHeapUsage"`
	HeapAllocated        float64 `json:"heapInUse"`
	ObjectsInUse         uint64  `json:"objectsInUse"`
	OSMemoryObtained     float64 `json:"OSMemoryObtained"`
}

func getHealthStats() *HealthStats {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)

	return &HealthStats{
		Uptime:               getUptime(),
		AllocatedMemory:      toMegaBytes(mem.Alloc),
		TotalAllocatedMemory: toMegaBytes(mem.TotalAlloc),
		Goroutines:           runtime.NumGoroutine(),
		NumberOfCPUs:         runtime.NumCPU(),
		GCCycles:             mem.NumGC,
		HeapSys:              toMegaBytes(mem.HeapSys),
		HeapAllocated:        toMegaBytes(mem.HeapAlloc),
		ObjectsInUse:         mem.Mallocs - mem.Frees,
		OSMemoryObtained:     toMegaBytes(mem.Sys),
	}
}

func getUptime() int64 {
	return time.Now().Unix() - start.Unix()
}

func toMegaBytes(bytes uint64) float64 {
	return toFixed(float64(bytes)/MB, 2)
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

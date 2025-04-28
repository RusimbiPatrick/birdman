package utils

import (
	"fmt"
	"net/http"
	"time"
	"birdman/stats"
    "net/http/httptest"
    "encoding/json"
)

func HTTPWithRetry(f func(string) (*http.Response, error), url string) (*http.Response, error) {
	count := 10
	var resp *http.Response
	var err error
	for i := 0; i < count; i++ {
		resp, err = f(url)
		if err != nil {
			fmt.Printf("Error calling url %v\n", url)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	return resp, err
}

// CreateTestServer creates a mock HTTP server for testing.
func CreateTestServer(handler http.Handler) *httptest.Server {
    return httptest.NewServer(handler)
}
// GetStats creates a mock stats object for testing.
func GetStats(diskTotal, diskFree, memTotal, memUsed int, user, nice, sys, idle, iowait, irq, softirq, steal, guest, guestNice uint64) *stats.Stats {
    return &stats.Stats{
        DiskTotal: diskTotal,
        DiskFree:  diskFree,
        MemTotal:  memTotal,
        MemUsed:   memUsed, // Corrected field name
        CpuStats: stats.CPU{ // Corrected type name
            User:      user,
            Nice:      nice,
            System:    sys,
            Idle:      idle,
            IOWait:    iowait,
            IRQ:       irq,
            SoftIRQ:   softirq,
            Steal:     steal,
            Guest:     guest,
            GuestNice: guestNice,
        },
    }
}

// NewStatsHandler creates a handler for serving stats.
func NewStatsHandler(stats *stats.Stats) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(stats)
    })
}
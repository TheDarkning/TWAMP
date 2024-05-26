package models

type PacketConfig struct {
	IP       string
	Port     int
	Count    int
	Interval int
	Payload  int
}

type ChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

type TableRow struct {
	ID      string `json:"id"`
	Latency string `json:"latency"`
}

type SummaryData struct {
	PacketLoss string `json:"packetLoss"`
	MinLatency string `json:"minLatency"`
	MaxLatency string `json:"maxLatency"`
	AvgLatency string `json:"avgLatency"`
}

type TableData struct {
	TableRows []TableRow  `json:"tableRows"`
	Summary   SummaryData `json:"summary"`
}


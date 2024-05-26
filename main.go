package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

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

var rawdata []string

func main() {
	fs := http.FileServer(http.Dir("app"))
	hn := http.StripPrefix("/app/", fs)
	http.Handle("/app/", hn)

	http.HandleFunc("/", handler)
	http.HandleFunc("/submited", formHandler)

	port := 8080
	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/submited":
		formHandler(w, r)
	case "/":
		mainPage(w, r)
	case "/result":
		resultHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		filler(w, r, rawdata)
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	target_ip := r.FormValue("target_ip")
	target_port, err := strconv.Atoi(r.FormValue("target_port"))
	if err != nil {
		http.Error(w, "Invalid port number", http.StatusBadRequest)
		return
	}
	packet_number, err := strconv.Atoi(r.FormValue("packet_number"))
	if err != nil {
		http.Error(w, "Invalid packet number", http.StatusBadRequest)
		return
	}
	interval, err := strconv.Atoi(r.FormValue("interval"))
	if err != nil {
		http.Error(w, "Invalid interval", http.StatusBadRequest)
		return
	}
	packet_size, err := strconv.Atoi(r.FormValue("packet_size"))
	if err != nil {
		http.Error(w, "Invalid packet size", http.StatusBadRequest)
		return
	}

	config := PacketConfig{
		IP:       target_ip,
		Port:     target_port,
		Count:    packet_number,
		Interval: interval,
		Payload:  packet_size,
	}

	err = startServer(config)
	if err != nil {
		http.Error(w, "Failed to start binary file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rttData, err := readOutputFile("out.txt")
	if err != nil {
		http.Error(w, "Failed to read output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawdata = rttData;

	filler(w,r, rawdata)
}

func startServer(config PacketConfig) error {
	cmd := exec.Command("./twamp/twamp_test",
		fmt.Sprintf("%s:%d", config.IP, config.Port),
		fmt.Sprintf("--count=%d", config.Count),
		fmt.Sprintf("--interval=%d", config.Interval),
		fmt.Sprintf("--payload=%d", config.Payload),
		"--output=out.txt")
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func readOutputFile(filename string) ([]string, error) {
	for {
		if _, err := os.Stat(filename); err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = os.Remove(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func generateChartData(rttData []string) string {
	labels := []string{}
	values := []float64{}

	for _, line := range rttData {
		if strings.HasPrefix(line, "ID ") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				labels = append(labels, parts[0])
				value, err := strconv.ParseFloat(parts[1], 64)
				if err == nil {
					values = append(values, value)
				}
			}
		}
	}

	chartData := ChartData{Labels: labels, Values: values}
	chartDataJSON, _ := json.Marshal(chartData)
	return string(chartDataJSON)
}

func generateTableData(rttData []string) string {
	tableRows := []TableRow{}
	var summary SummaryData

	for _, line := range rttData {
		if strings.HasPrefix(line, "ID ") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				tableRows = append(tableRows, TableRow{ID: parts[0], Latency: parts[1]})
			}
		} else if strings.HasPrefix(line, "Packet Loss:") {
			summary.PacketLoss = strings.TrimSpace(line)
		} else if strings.HasPrefix(line, "Min Latency:") {
			summary.MinLatency = strings.TrimSpace(line)
		} else if strings.HasPrefix(line, "Max Latency:") {
			summary.MaxLatency = strings.TrimSpace(line)
		} else if strings.HasPrefix(line, "Avg Latency:") {
			summary.AvgLatency = strings.TrimSpace(line)
		}
	}

	tableData := TableData{TableRows: tableRows, Summary: summary}
	tableDataJSON, _ := json.Marshal(tableData)
	return string(tableDataJSON)
}


func resultHandler(w http.ResponseWriter, r *http.Request) {
	tableData := generateTableData(rawdata)

	data := struct {
		TableData template.JS
	}{
		TableData: template.JS(tableData),
	}

	t, err := template.ParseFiles("app/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error during executing of the template", http.StatusInternalServerError)
	}
}

func filler(w http.ResponseWriter, r *http.Request, rdata []string) {
	if len(rdata) == 0 {
		http.Redirect(w, r, "/", http.StatusNoContent)
		return
	}
	chartData := generateChartData(rdata)
	tableData := generateTableData(rdata)

	data := struct {
		ChartData template.JS
		TableData template.JS
	}{
		ChartData: template.JS(chartData),
		TableData: template.JS(tableData),
	}

	t, err := template.ParseFiles("app/submited.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error during executing of the template", http.StatusInternalServerError)
	}
}

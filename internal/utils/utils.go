package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"TWAMP/internal/models"
)

func ReadOutputFile(filename string) ([]string, error) {
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

func GenerateChartData(rttData []string) string {
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

	chartData := models.ChartData{Labels: labels, Values: values}
	chartDataJSON, _ := json.Marshal(chartData)
	return string(chartDataJSON)
}

func GenerateTableData(rttData []string) string {
	tableRows := []models.TableRow{}
	var summary models.SummaryData

	for _, line := range rttData {
		if strings.HasPrefix(line, "ID ") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				tableRows = append(tableRows, models.TableRow{ID: parts[0], Latency: parts[1]})
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

	tableData := models.TableData{TableRows: tableRows, Summary: summary}
	tableDataJSON, _ := json.Marshal(tableData)
	return string(tableDataJSON)
}


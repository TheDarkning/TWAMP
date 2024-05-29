package utils

import (
	"TWAMP/internal/models"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Test cases for ReadOutputFile
func TestReadOutputFile(t *testing.T) {
	// Test 1: Normal case
	t.Run("NormalCase", func(t *testing.T) {
		filename := "testfile1.txt"
		content := "ID 1: 23.5\nID 2: 30.1\n"
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}
		defer os.Remove(filename)

		lines, err := ReadOutputFile(filename)
		if err != nil {
			t.Fatalf("ReadOutputFile returned error: %v", err)
		}

		expectedLines := []string{"ID 1: 23.5", "ID 2: 30.1", ""}
		if !reflect.DeepEqual(lines, expectedLines) {
			t.Errorf("Expected %v, got %v", expectedLines, lines)
		}
	})

	// Test 2: File does not exist initially
	t.Run("FileDoesNotExistInitially", func(t *testing.T) {
		filename := "testfile2.txt"
		content := "ID 1: 23.5\nID 2: 30.1\n"

		go func() {
			time.Sleep(200 * time.Millisecond)
			ioutil.WriteFile(filename, []byte(content), 0644)
		}()

		lines, err := ReadOutputFile(filename)
		if err != nil {
			t.Fatalf("ReadOutputFile returned error: %v", err)
		}
		defer os.Remove(filename)

		expectedLines := []string{"ID 1: 23.5", "ID 2: 30.1", ""}
		if !reflect.DeepEqual(lines, expectedLines) {
			t.Errorf("Expected %v, got %v", expectedLines, lines)
		}
	})

	// Test 3: Empty file
	t.Run("EmptyFile", func(t *testing.T) {
		filename := "testfile3.txt"
		content := ""
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}
		defer os.Remove(filename)

		lines, err := ReadOutputFile(filename)
		if err != nil {
			t.Fatalf("ReadOutputFile returned error: %v", err)
		}

		expectedLines := []string{""}
		if !reflect.DeepEqual(lines, expectedLines) {
			t.Errorf("Expected %v, got %v", expectedLines, lines)
		}
	})

	// Test 4: File with no new line at end
	t.Run("FileWithoutNewLineAtEnd", func(t *testing.T) {
		filename := "testfile4.txt"
		content := "ID 1: 23.5\nID 2: 30.1"
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}
		defer os.Remove(filename)

		lines, err := ReadOutputFile(filename)
		if err != nil {
			t.Fatalf("ReadOutputFile returned error: %v", err)
		}

		expectedLines := []string{"ID 1: 23.5", "ID 2: 30.1"}
		if !reflect.DeepEqual(lines, expectedLines) {
			t.Errorf("Expected %v, got %v", expectedLines, lines)
		}
	})

	// Test 5: Large file
	t.Run("LargeFile", func(t *testing.T) {
		filename := "testfile5.txt"
		content := ""
		for i := 0; i < 1000; i++ {
			content += "ID " + strconv.Itoa(i) + ": " + strconv.FormatFloat(float64(i), 'f', 1, 64) + "\n"
		}
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}
		defer os.Remove(filename)

		lines, err := ReadOutputFile(filename)
		if err != nil {
			t.Fatalf("ReadOutputFile returned error: %v", err)
		}

		expectedLines := strings.Split(content, "\n")
		if !reflect.DeepEqual(lines, expectedLines) {
			t.Errorf("Expected %v, got %v", expectedLines, lines)
		}
	})
}

// Test cases for GenerateChartData
func TestGenerateChartData(t *testing.T) {
	// Test 1: Normal case
	t.Run("NormalCase", func(t *testing.T) {
		rttData := []string{"ID 1: 23.5", "ID 2: 30.1"}
		expectedLabels := []string{"ID 1", "ID 2"}
		expectedValues := []float64{23.5, 30.1}

		chartDataJSON := GenerateChartData(rttData)

		var chartData models.ChartData
		err := json.Unmarshal([]byte(chartDataJSON), &chartData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(chartData.Labels, expectedLabels) {
			t.Errorf("Expected labels %v, got %v", expectedLabels, chartData.Labels)
		}

		if !reflect.DeepEqual(chartData.Values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, chartData.Values)
		}
	})

	// Test 2: No data
	t.Run("NoData", func(t *testing.T) {
		rttData := []string{}
		expectedLabels := []string{}
		expectedValues := []float64{}

		chartDataJSON := GenerateChartData(rttData)

		var chartData models.ChartData
		err := json.Unmarshal([]byte(chartDataJSON), &chartData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(chartData.Labels, expectedLabels) {
			t.Errorf("Expected labels %v, got %v", expectedLabels, chartData.Labels)
		}

		if !reflect.DeepEqual(chartData.Values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, chartData.Values)
		}
	})

	// Test 3: Invalid data format
	t.Run("InvalidDataFormat", func(t *testing.T) {
		rttData := []string{"ID 1: 23.5", "Invalid Data"}
		expectedLabels := []string{"ID 1"}
		expectedValues := []float64{23.5}

		chartDataJSON := GenerateChartData(rttData)

		var chartData models.ChartData
		err := json.Unmarshal([]byte(chartDataJSON), &chartData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(chartData.Labels, expectedLabels) {
			t.Errorf("Expected labels %v, got %v", expectedLabels, chartData.Labels)
		}

		if !reflect.DeepEqual(chartData.Values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, chartData.Values)
		}
	})

	// Test 4: Data with non-numeric values
	t.Run("DataWithNonNumericValues", func(t *testing.T) {
		rttData := []string{"ID 1: abc", "ID 2: 30.1"}
		expectedLabels := []string{"ID 2"}
		expectedValues := []float64{30.1}

		chartDataJSON := GenerateChartData(rttData)

		var chartData models.ChartData
		err := json.Unmarshal([]byte(chartDataJSON), &chartData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(chartData.Labels, expectedLabels) {
			t.Errorf("Expected labels %v, got %v", expectedLabels, chartData.Labels)
		}

		if !reflect.DeepEqual(chartData.Values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, chartData.Values)
		}
	})

	// Test 5: Large data set
	t.Run("LargeDataSet", func(t *testing.T) {
		rttData := []string{}
		expectedLabels := []string{}
		expectedValues := []float64{}
		for i := 0; i < 1000; i++ {
			rttData = append(rttData, "ID "+strconv.Itoa(i)+": "+strconv.FormatFloat(float64(i), 'f', 1, 64))
			expectedLabels = append(expectedLabels, "ID "+strconv.Itoa(i))
			expectedValues = append(expectedValues, float64(i))
		}

		chartDataJSON := GenerateChartData(rttData)

		var chartData models.ChartData
		err := json.Unmarshal([]byte(chartDataJSON), &chartData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(chartData.Labels, expectedLabels) {
			t.Errorf("Expected labels %v, got %v", expectedLabels, chartData.Labels)
		}

		if !reflect.DeepEqual(chartData.Values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, chartData.Values)
		}
	})
}

// Test cases for GenerateTableData
func TestGenerateTableData(t *testing.T) {
	// Test 1: Normal case
	t.Run("NormalCase", func(t *testing.T) {
		rttData := []string{
			"ID 1: 23.5",
			"ID 2: 30.1",
			"Packet Loss: 0%",
			"Min Latency: 23.5ms",
			"Max Latency: 30.1ms",
			"Avg Latency: 26.8ms",
		}

		expectedRows := []models.TableRow{
			{ID: "ID 1", Latency: "23.5"},
			{ID: "ID 2", Latency: "30.1"},
		}

		expectedSummary := models.SummaryData{
			PacketLoss: "Packet Loss: 0%",
			MinLatency: "Min Latency: 23.5ms",
			MaxLatency: "Max Latency: 30.1ms",
			AvgLatency: "Avg Latency: 26.8ms",
		}

		tableDataJSON := GenerateTableData(rttData)

		var tableData models.TableData
		err := json.Unmarshal([]byte(tableDataJSON), &tableData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(tableData.TableRows, expectedRows) {
			t.Errorf("Expected rows %v, got %v", expectedRows, tableData.TableRows)
		}

		if !reflect.DeepEqual(tableData.Summary, expectedSummary) {
			t.Errorf("Expected summary %v, got %v", expectedSummary, tableData.Summary)
		}
	})

	// Test 2: No data
	t.Run("NoData", func(t *testing.T) {
		rttData := []string{}

		expectedRows := []models.TableRow{}
		expectedSummary := models.SummaryData{}

		tableDataJSON := GenerateTableData(rttData)

		var tableData models.TableData
		err := json.Unmarshal([]byte(tableDataJSON), &tableData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(tableData.TableRows, expectedRows) {
			t.Errorf("Expected rows %v, got %v", expectedRows, tableData.TableRows)
		}

		if !reflect.DeepEqual(tableData.Summary, expectedSummary) {
			t.Errorf("Expected summary %v, got %v", expectedSummary, tableData.Summary)
		}
	})

	// Test 3: Invalid data format
	t.Run("InvalidDataFormat", func(t *testing.T) {
		rttData := []string{
			"ID 1: 23.5",
			"Invalid Data",
			"Packet Loss: 0%",
			"Min Latency: 23.5ms",
			"Max Latency: 30.1ms",
			"Avg Latency: 26.8ms",
		}

		expectedRows := []models.TableRow{
			{ID: "ID 1", Latency: "23.5"},
		}

		expectedSummary := models.SummaryData{
			PacketLoss: "Packet Loss: 0%",
			MinLatency: "Min Latency: 23.5ms",
			MaxLatency: "Max Latency: 30.1ms",
			AvgLatency: "Avg Latency: 26.8ms",
		}

		tableDataJSON := GenerateTableData(rttData)

		var tableData models.TableData
		err := json.Unmarshal([]byte(tableDataJSON), &tableData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(tableData.TableRows, expectedRows) {
			t.Errorf("Expected rows %v, got %v", expectedRows, tableData.TableRows)
		}

		if !reflect.DeepEqual(tableData.Summary, expectedSummary) {
			t.Errorf("Expected summary %v, got %v", expectedSummary, tableData.Summary)
		}
	})

	// Test 4: Data with missing summary
	t.Run("DataWithMissingSummary", func(t *testing.T) {
		rttData := []string{
			"ID 1: 23.5",
			"ID 2: 30.1",
		}

		expectedRows := []models.TableRow{
			{ID: "ID 1", Latency: "23.5"},
			{ID: "ID 2", Latency: "30.1"},
		}

		expectedSummary := models.SummaryData{}

		tableDataJSON := GenerateTableData(rttData)

		var tableData models.TableData
		err := json.Unmarshal([]byte(tableDataJSON), &tableData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(tableData.TableRows, expectedRows) {
			t.Errorf("Expected rows %v, got %v", expectedRows, tableData.TableRows)
		}

		if !reflect.DeepEqual(tableData.Summary, expectedSummary) {
			t.Errorf("Expected summary %v, got %v", expectedSummary, tableData.Summary)
		}
	})

	// Test 5: Large data set
	t.Run("LargeDataSet", func(t *testing.T) {
		rttData := []string{}
		expectedRows := []models.TableRow{}
		for i := 0; i < 1000; i++ {
			rttData = append(rttData, "ID "+strconv.Itoa(i)+": "+strconv.FormatFloat(float64(i), 'f', 1, 64))
			expectedRows = append(expectedRows, models.TableRow{ID: "ID " + strconv.Itoa(i), Latency: strconv.FormatFloat(float64(i), 'f', 1, 64)})
		}
		rttData = append(rttData, "Packet Loss: 0%")
		rttData = append(rttData, "Min Latency: 0.0ms")
		rttData = append(rttData, "Max Latency: 999.0ms")
		rttData = append(rttData, "Avg Latency: 499.5ms")

		expectedSummary := models.SummaryData{
			PacketLoss: "Packet Loss: 0%",
			MinLatency: "Min Latency: 0.0ms",
			MaxLatency: "Max Latency: 999.0ms",
			AvgLatency: "Avg Latency: 499.5ms",
		}

		tableDataJSON := GenerateTableData(rttData)

		var tableData models.TableData
		err := json.Unmarshal([]byte(tableDataJSON), &tableData)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		if !reflect.DeepEqual(tableData.TableRows, expectedRows) {
			t.Errorf("Expected rows %v, got %v", expectedRows, tableData.TableRows)
		}

		if !reflect.DeepEqual(tableData.Summary, expectedSummary) {
			t.Errorf("Expected summary %v, got %v", expectedSummary, tableData.Summary)
		}
	})
}

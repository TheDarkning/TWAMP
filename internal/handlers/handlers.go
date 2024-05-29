package handlers

import (
	"fmt"
	"TWAMP/internal/commands"
	"TWAMP/internal/models"
	"TWAMP/internal/utils"
	"html/template"
	"net/http"
	"strconv"
)

var rawdata []string

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/submited":
		FormHandler(w, r)
	case "/":
		MainPage(w, r)
	case "/result":
		ResultHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Filler(w, r, rawdata)
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

	config := models.PacketConfig{
		IP:       target_ip,
		Port:     target_port,
		Count:    packet_number,
		Interval: interval,
		Payload:  packet_size,
	}

	fmt.Println(target_ip, target_port, packet_number, interval, packet_size)
	err = commands.StartServer(config)
	if err != nil {
		http.Error(w, "Failed to start binary file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rttData, err := utils.ReadOutputFile("out.txt")
	if err != nil {
		http.Error(w, "Failed to read output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawdata = rttData

	Filler(w, r, rawdata)
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	tableData := utils.GenerateTableData(rawdata)

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

func Filler(w http.ResponseWriter, r *http.Request, rdata []string) {
	if len(rdata) == 0 {
		http.Redirect(w, r, "/", http.StatusNoContent)
		return
	}
	chartData := utils.GenerateChartData(rdata)
	tableData := utils.GenerateTableData(rdata)

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

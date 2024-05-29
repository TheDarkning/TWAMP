package handlers

import (
	// "TWAMP/internal/commands"
	"TWAMP/internal/models"
	// "TWAMP/internal/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Handlers struct {
	StartServer       func(config models.PacketConfig) error
	ReadOutputFile    func(filename string) ([]string, error)
	GenerateChartData func(rttData []string) string
	GenerateTableData func(rttData []string) string
	ParseFiles        func(filenames ...string) (*template.Template, error)
}

var rawdata []string

func (h *Handlers) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/submitted":
		h.FormHandler(w, r)
	case "/":
		h.MainPage(w, r)
	case "/result":
		h.ResultHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handlers) MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := h.ParseFiles("app/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.Filler(w, r, rawdata)
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
	err = h.StartServer(config)
	if err != nil {
		http.Error(w, "Failed to start binary file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rttData, err := h.ReadOutputFile("out.txt")
	if err != nil {
		http.Error(w, "Failed to read output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawdata = rttData

	h.Filler(w, r, rawdata)
}

func (h *Handlers) ResultHandler(w http.ResponseWriter, r *http.Request) {
	tableData := h.GenerateTableData(rawdata)

	data := struct {
		TableData template.JS
	}{
		TableData: template.JS(tableData),
	}

	t, err := h.ParseFiles("app/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error during executing of the template", http.StatusInternalServerError)
	}
}

func (h *Handlers) Filler(w http.ResponseWriter, r *http.Request, rdata []string) {
	if len(rdata) == 0 {
		http.Redirect(w, r, "/", http.StatusNoContent)
		return
	}
	chartData := h.GenerateChartData(rdata)
	tableData := h.GenerateTableData(rdata)

	data := struct {
		ChartData template.JS
		TableData template.JS
	}{
		ChartData: template.JS(chartData),
		TableData: template.JS(tableData),
	}

	t, err := h.ParseFiles("app/submitted.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error during executing of the template", http.StatusInternalServerError)
	}
}

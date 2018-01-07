package utils

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/467754239/GolangNote/lh/common"
	"github.com/olekukonko/tablewriter"
)

func FormatTable(output []common.Hostinfo) {
	data := make([][]string, 10, 10)
	for i, v := range output {
		var m []string
		m = append(m, strconv.Itoa(i+1), v.Instance_id, v.Hostname, v.Public_ip, v.Private_ip)
		data = append(data, m)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "实例Id", "主机名", "公网IP", "内网IP"})
	/*
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	*/

	table.SetBorder(true) // Set Border to false
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.BgGreenColor, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgHiBlueColor, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

	table.AppendBulk(data)
	table.Render()

}

func FormatJson(data []common.Hostinfo) {
	jsIndent, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(jsIndent))
}

func FormatXml(data []common.Hostinfo) {
	output, _ := xml.MarshalIndent(data, "  ", "    ")
	os.Stdout.Write(output)
}

func FormatCsv(output []common.Hostinfo) {
	fd, _ := os.OpenFile("zhengyscn.csv", os.O_RDWR|os.O_CREATE, 0644)
	defer fd.Close()

	w := csv.NewWriter(fd)
	//w := csv.NewWriter(os.Stdout)

	data := make([][]string, 10, 10)
	for i, v := range output {
		var m []string
		m = append(m, strconv.Itoa(i+1), v.Instance_id, v.Hostname, v.Public_ip, v.Private_ip)
		data = append(data, m)
	}
	for _, record := range data {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}

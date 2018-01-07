package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/467754239/GolangNote/lh/audit"
	"github.com/467754239/GolangNote/lh/common"
	"github.com/467754239/GolangNote/lh/config"
	"github.com/467754239/GolangNote/lh/httpapi"
	"github.com/467754239/GolangNote/lh/utils"
	"github.com/fatih/color"
)

var (
	url      string
	response common.Response
	results  []common.Hostinfo
	host     = flag.String("host", "", "Match remote host")
	format   = flag.String("format", "default", "Format {default | table | json | xml}")
	version  = flag.Bool("version", false, "Show Version.")
)

func fibGet(_url string) {
	//v1 >>> /api/v1/assets&page=2
	//v2 >>> /api/v1/assets?default=*logstash*&page=2
	buf, err := httpapi.GetHttpApi(_url)
	buf = strings.Replace(buf, "\"next\":null", "\"next\":\"null\"", -1)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(buf), &response)
	if err != nil {
		log.Fatal(err)
	}
	results = append(results, response.Results...)
	if response.Next == "null" {
		return
	} else {
		if strings.Contains(response.Next, "default") {
			bytes := strings.Split(response.Next, "?")
			_url = fmt.Sprintf("%s?%s", config.VKLH_ASSETS_PROD_URL, bytes[1])
			fibGet(_url)
		}
	}

}

func PrintDefaults() {
	fmt.Fprintf(os.Stderr, "usage: lh-client -host=*monkey* -format=default\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Version: %d.%d\nOperation: Yansheng Zheng\nCopyright © 2017 - 2017 zhengyscn.com.cn . All Rights Reserved.\n", config.VKLH_CLIENT_VER_MAJOR, config.VKLH_CLIENT_VER_MINOR)
		os.Exit(0)
	} else if *host == "" {
		Usage()
	}

	if err := audit.Authentication(); err != nil {
		color.Red("[!]Illegal implementation, Disable execution in the current host.")
		//os.Exit(0)
	}

	url = fmt.Sprintf("%s?default=%s", config.VKLH_ASSETS_PROD_URL, *host)
	fibGet(url)

	// format output
	switch *format {
	case "table":
		utils.FormatTable(results)
	case "json":
		utils.FormatJson(results)
	case "xml":
		utils.FormatXml(results)
	default:
		fmt.Printf("%#v\n", results)
	}

	// send audit to server.
	sender := audit.NewSender(fmt.Sprintf("%s:%d", config.VKLH_AUDIT_PROD_ADDR, config.VKLH_AUDIT_PROD_PORT))
	sender.Start()
}

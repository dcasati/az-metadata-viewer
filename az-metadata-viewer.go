package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type AzMetadata struct {
	Compute struct {
		Location             string `json:"location"`
		Name                 string `json:"name"`
		Offer                string `json:"offer"`
		OsType               string `json:"osType"`
		PlatformFaultDomain  string `json:"platformFaultDomain"`
		PlatformUpdateDomain string `json:"platformUpdateDomain"`
		Publisher            string `json:"publisher"`
		Sku                  string `json:"sku"`
		Version              string `json:"version"`
		VMID                 string `json:"vmId"`
		VMSize               string `json:"vmSize"`
	} `json:"compute"`
	Network struct {
		Interface []struct {
			Ipv4 struct {
				IPAddress []struct {
					PrivateIPAddress string `json:"privateIpAddress"`
					PublicIPAddress  string `json:"publicIpAddress"`
				} `json:"ipAddress"`
				Subnet []struct {
					Address string `json:"address"`
					Prefix  string `json:"prefix"`
				} `json:"subnet"`
			} `json:"ipv4"`
			Ipv6 struct {
				IPAddress []interface{} `json:"ipAddress"`
			} `json:"ipv6"`
			MacAddress string `json:"macAddress"`
		} `json:"interface"`
	} `json:"network"`
}

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)

	req.Header.Add("Metadata", "True")
	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2017-04-02")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	var m AzMetadata
	err = json.Unmarshal(resp_body, &m)
	if err != nil {
		fmt.Println("error:", err)
	}

	t := template.Must(template.New("view").Parse(viewTemplate))
	if err = t.Execute(os.Stdout, m); err != nil {
		panic(err)
	}
}

const viewTemplate = `
Compute 
------------------------------------------------------------------------------
{{if .Compute.Location}}{{printf "\tLocation: %s\n" .Compute.Location}}{{end -}}
{{if .Compute.Name}}{{printf "\tName: %s\n" .Compute.Name}}{{end -}}
{{if .Compute.Offer}}{{printf "\tOffer: %s\n" .Compute.Offer}}{{end -}}
{{if .Compute.OsType}}{{printf "\tOS Type: %s\n" .Compute.OsType}}{{end -}}
{{if .Compute.PlatformFaultDomain}}{{printf "\tPlatform Fault Domain: %s\n" .Compute.PlatformFaultDomain}}{{end -}}
{{if .Compute.PlatformUpdateDomain}}{{printf "\tPlatform Update Domain: %s\n" .Compute.PlatformUpdateDomain}}{{end -}}
{{if .Compute.Publisher}}{{printf "\tPublisher: %s\n" .Compute.Publisher}}{{end -}}
{{if .Compute.Sku}}{{printf "\tSku: %s\n" .Compute.Sku}}{{end -}}
{{if .Compute.Version}}{{printf "\tVersion: %s\n" .Compute.Version}}{{end -}}
{{if .Compute.VMID}}{{printf "\tVMID: %s\n" .Compute.VMID}}{{end -}} 
{{if .Compute.VMSize}}{{printf "\tVM Size: %s" .Compute.VMSize}}{{end}} 

Network
------------------------------------------------------------------------------
{{range .Network.Interface -}}
{{range .Ipv4.IPAddress -}}
*
{{if .PrivateIPAddress}}{{printf "\tPrivate IP: %s" .PrivateIPAddress}}{{end}}
{{if .PublicIPAddress}}{{printf "\tPublic IP: %s" .PublicIPAddress}}{{end}}
{{end -}}
{{end -}}
`

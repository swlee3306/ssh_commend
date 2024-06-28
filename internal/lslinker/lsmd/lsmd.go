package lsmd

type OutputMetricData struct {
	Agent struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		LocId string `json:"loc_id"`
	} `json:"agent"`
	Service struct {
		Type    string `json:"type"`
		Adapter string `json:"adapter"`
	} `json:"service"`
	CollectorTime int64 `json:"collector_time"`
	ApiInfo       struct {
		Type       string            `json:"type"`
		Url        string            `json:"url"`
		PortPath   string            `json:"port_path"`
		Credential map[string]string `json:"credential"`
		Filter     struct {
			Type  string   `json:"type"`
			Value []string `json:"value"`
		} `json:"filter"`
	} `json:"api_info"`
	OutputInfo struct {
		Type string   `json:"type"`
		Urls []string `json:"urls"`
	} `json:"output_info"`
	Basic struct {
		Vm struct{
			VmId       string  `json:"vm_id"`
			Domain     string  `json:"domain"`
			PowerState string  `json:"power_state"`
			State      float64 `json:"state"`
		}`json:"vm"`
		System struct {
			Filesystem  struct {
				All struct{
					Available float64 `json:"available"`
					Total float64 `json:"total"`
				} `json:"all"`
				Each struct{
					Name string `json:"name"`
					Available float64 `json:"Available"`
					Total float64 `json:"total"`
				} `json:"each"`
			} `json:"filesystem"`
		} `json:"system"`
		HypervisorIp string `json:"hypervisor_ip"`
	} `json:"basic"`
}
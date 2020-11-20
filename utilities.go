package harperdb

import "time"

func (c *Client) DeleteFilesBefore(schema, table string, date time.Time) error {
	return c.opRequest(operation{
		Operation: OP_DELETE_FILES_BEFORE,
		Schema:    schema,
		Table:     table,
		Date:      date,
	}, nil)
}

const (
	SearchBySQL   = "sql"
	SearchByHash  = "search_by_hash"
	SearchByValue = "search_by_value"
)

type SearchOperation struct {
	Operation string `json:"operation"`
	SQL       string `json:"sql"`
}

type S3Credentials struct {
	AWSAccessKeyID     string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
	Bucket             string `json:"bucket"`
	Key                string `json:"filename"`
}

type SysInfo struct {
	System struct {
		Platform    string `json:"platform"`
		Distro      string `json:"distro"`
		Release     string `json:"release"`
		Codename    string `json:"codename"`
		Kernel      string `json:"kernel"`
		Arch        string `json:"arch"`
		Hostname    string `json:"hostname"`
		NodeVersion string `json:"node_version"`
		NPMVersion  string `json:"npm_version"`
	} `json:"system"`
	Time struct {
		Current      Timestamp `json:"current"`
		Uptime       int64     `json:"uptime"`
		Timezone     string    `json:"timezone"`
		TimezoneName string    `json:"timezoneName"`
	} `json:"time"`
	CPU struct {
		Manufacturer  string `json:"manufacturer"`
		Brand         string `json:"brand"`
		Vendor        string `json:"vendor"`
		Speed         string `json:"speed"`
		Cores         int    `json:"cores"`
		PhysicalCores int    `json:"physicalCores"`
		Processors    int    `json:"processors"`
		CPUSpeed      struct {
			Min   float64   `json:"min"`
			Max   float64   `json:"max"`
			Avg   float64   `json:"avg"`
			Cores []float64 `json:"cores"`
		} `json:"cpu_speed"`
		CurrentLoad struct {
			AvgLoag           float64 `json:"avgload"`
			CurrentLoad       float64 `json:"currentload"`
			CurrentLoadUser   float64 `json:"currentload_user"`
			CurrentLoadSystem float64 `json:"currentload_system"`
			CurrentLoadNice   float64 `json:"currentload_nice"`
			CurrentLoadIdle   float64 `json:"currentload_idle"`
			CurrentLoadIRQ    float64 `json:"currentload_irq"`
		} `json:"current_load"`
		CPUs []CPULoad `json:"cpus"`
	} `json:"cpu"`
	Memory struct {
		Total     int64 `json:"total"`
		Free      int64 `json:"free"`
		Used      int64 `json:"used"`
		Active    int64 `json:"active"`
		Available int64 `json:"available"`
		SwapTotal int64 `json:"swaptotal"`
		SwapUsed  int64 `json:"swapused"`
		SwapFree  int64 `json:"swapfree"`
	} `json:"memory"`
	Disk struct {
		IO struct {
			RIO int64 `json:"rIO"`
			WIO int64 `json:"wIO"`
			TIO int64 `json:"tIO"`
		} `json:"io"`
		ReadWrite struct {
			RX int64 `json:"rx"`
			WX int64 `json:"wx"`
			TX int64 `json:"tx"`
			MS int64 `json:"ms"`
		} `json:"read_write"`
		Size []DiskSize `json:"size"`
	} `json:"disk"`
	Network struct {
		DefaultInterface string `json:"default_interface"`
		Latency          struct {
			URL    string `json:"url"`
			Ok     bool   `json:"ok"`
			Status int64  `json:"status"`
			MS     int64  `json:"ms"`
		} `json:"latency"`
		Interfaces  []NetworkInterface  `json:"interfaces"`
		Stats       []NetworkStats      `json:"stats"`
		Connections []NetworkConnection `json:"connections"`
	} `json:"network"`
	HarperDBProcesses struct {
		Core       interface{} `json:"core"`       // TODO Unknown
		Clustering interface{} `json:"clustering"` // TODO Unknown
	} `json:"harperdb_processes"`
	TableSize []TableSize `json:"table_size"`
}

type CPULoad struct {
	Load       float64 `json:"load"`
	LoadUser   float64 `json:"load_user"`
	LoadSystem float64 `json:"load_system"`
	LoadNice   float64 `json:"load_nice"`
	LoadIdle   float64 `json:"load_idle"`
	LoadIRQ    float64 `json:"load_irq"`
}

type DiskSize struct {
	FS    string  `json:"fs"`
	Type  string  `json:"overlay"`
	Size  int64   `json:"size"`
	Used  int64   `json:"used"`
	Use   float64 `json:"use"`
	Mount string  `json:"mount"`
}

type NetworkInterface struct {
	Iface          string `json:"iface"`
	IfaceName      string `json:"ifaceName"`
	IP4            string `json:"ip4"`
	IP6            string `json:"ip6"`
	Mac            string `json:"mac"`
	OperState      string `json:"operstate"`
	Type           string `json:"virtual"`
	Duplex         string `json:"duplex"`
	Speed          int64  `json:"speed"`
	CarrierChanges int64  `json:"carrierChanges"`
}

type NetworkStats struct {
	Iface     string `json:"iface"`
	OperState string `json:"operstate"`
	RxBytes   int64  `json:"rx_bytes"`
	RxDropped int64  `json:"rx_dropped"`
	RxErrors  int64  `json:"rx_errors"`
	TxBytes   int64  `json:"tx_bytes"`
	TxDropped int64  `json:"tx_dropped"`
	TxErrors  int64  `json:"tx_errors"`
}

type NetworkConnection struct {
	Protocol     string `json:"protocol"`
	LocalAddress string `json:"localaddress"`
	LocalPort    string `json:"localport"`
	PeerAddress  string `json:"peeraddress"`
	PeerPort     string `json:"peerport"`
	State        string `json:"state"`
	PID          int64  `json:"pid"`
	Process      string `json:"node"`
}

type TableSize struct {
	Schema                    string `json:"schema"`
	Table                     string `json:"table"`
	TableSize                 int64  `json:"table_size"`
	RecordCount               int64  `json:"record_count"`
	TransactionLogSize        int64  `json:"transaction_log_size"`
	TransactionLogRecordCount int64  `json:"transaction_log_record_count"`
}

func (c *Client) ExportLocal(format, path string, searchOperation SearchOperation) error {
	return c.opRequest(operation{
		Operation:       OP_EXPORT_TO_S3,
		Format:          format,
		Path:            path,
		SearchOperation: searchOperation,
	}, nil)
}

func (c *Client) ExportToS3(format string, s3creds S3Credentials, searchOperation SearchOperation) error {
	return c.opRequest(operation{
		Operation:       OP_EXPORT_TO_S3,
		Format:          format,
		S3:              s3creds,
		SearchOperation: searchOperation,
	}, nil)
}

func (c *Client) SystemInformation() (*SysInfo, error) {
	var sysInfo SysInfo
	err := c.opRequest(operation{
		Operation: OP_SYSTEM_INFORMATION,
	}, &sysInfo)
	return &sysInfo, err
}

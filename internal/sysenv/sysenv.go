package sysenv

var (
	SSH_COMMEND = make(chan CommendReq)
	SSH_COMMEND_ALL = make(chan CommendAllReq)
	SSH_COMMEND_RES = make(chan CommendRes)
	SSH_COMMEND_ALL_RES = make(chan CommendAllRes)
)

var Mode struct {
	IsDebug bool
}
var Database struct {
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxLifetimeHour int
}

var ApiServer struct {
	Tls      bool
	Port     int
	KeyFile  string
	CertFile string
}

type CommendReq struct {
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	HostIp string `json:"host_ip"`
	Port   uint `json:"port"`
	Cmd    string `json:"cmd"`
}

type CommendAllReq struct {
	Status bool
	Cmd    string `json:"cmd"`
}

type CommendRes struct{
	User string `json:"user"`
	Output string `json:"output"`
	Err string `json:"error"`
}

type CommendAllRes struct{
	Vm []*CommendRes `json:"VM"`
}

type Configuration struct {
	Vmattr struct {
		Id  string
		Pwd string
		HostIp string
		Port uint
		State float64
	}
	Logstash struct {
		Urls string
	}
	Cmdline struct{
		Cmd string
	}
}

type VmEnv struct {
	ID	 int64
	PrvId int64
	UUID string
	Name string
	Cfg  *Configuration
}

type DiskData struct {
	FilesystemName string
	TotalSize float64
	UsedSize float64
	AvailableSize float64
	UsagePercent string
	Uuid string
}

type DiskDatas struct{
	DiskData []*DiskData
}
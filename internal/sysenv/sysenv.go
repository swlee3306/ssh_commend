package sysenv

var (
	SSH_COMMEND = make(chan CommendReq)
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
	Port   string `json:"port"`
	Cmd    string `json:"cmd"`
}

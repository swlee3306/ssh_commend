package sysenv

var (
	SSH_COMMEND = make(chan CommendReq)
	SSH_COMMEND_RES = make(chan CommendRes)
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

type CommendRes struct{
	Output string `json:"output"`
	Err string `json:"error"`
}

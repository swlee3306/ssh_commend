package cmdline

type cmd struct{
	DISK_USE string
}

var Cmdline = cmd {
	DISK_USE: "sudo dmidecode -s system-uuid ; df -B1 | awk 'NR > 1 && $1 ~ /^"+"\\"+"// {print $1, $2, $3, $4, $5}'",
}
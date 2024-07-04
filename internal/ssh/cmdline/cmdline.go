package cmdline

type cmd struct{
	DISK_USE string
	OS_KIND string
}

var Cmdline = cmd {
	DISK_USE: "sudo dmidecode -s system-uuid ; df -B1 | awk 'NR > 1 && $1 ~ /^"+"\\"+"// {print $1, $2, $3, $4, $5}'",
	OS_KIND: `awk -F= '/^(DISTRIB_ID|DISTRIB_RELEASE|DISTRIB_CODENAME|DISTRIB_DESCRIPTION)/ {print $2}' /etc/*-release | sed 's/"//g'`,
}
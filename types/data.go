package types

type LoginConfig struct {
	User     string `form:"username"`
	Password string `form:"password"`
	Addr     string `form:"address"`
}

type ServerInfo struct {
	Title  string
	Url    string
	Target string
}

const DomainsCmd string = "grep -r -Eo 'server_name .*\\.com' /etc/nginx/ | awk '{print $2}'"
const OsCmd string = "lsb_release -a 2>/dev/null | grep 'Description' | awk -F ':\t' '{print $2}'"
const HostCmd string = "hostname"
const IpCmd string = "hostname -I | awk '{print $1}'"
const RamCmd string = "sudo dmidecode -t 17 | grep Size"
const CpuCmd string = "nproc"

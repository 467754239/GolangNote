package common

type next_step interface{}

type Hostinfo struct {
	Instance_id string `json:"instance_id"`
	Hostname    string `json:"hostname"`
	Public_ip   string `json:"public_ip"`
	Private_ip  string `json:"private_ip"`
}

type Response struct {
	Count    int    `json:"count"`
	Previous string `json:"previous"`
	//Next     next_step `json:"next"`
	Next    string     `json:"next"`
	Results []Hostinfo `json:"results"`
}

type UserAudit struct {
	//Id        int    `xorm:"autoincr default(0) 'id'"`
	Hostname  string `xorm:"'hostname'"`
	Loginuser string `xorm:"'login_user'"`
	Localip   string `xorm:"'local_ip'"`
	Fromip    string `xorm:"'from_ip'"`
	Curuser   string `xorm:"'cur_user'"`
	Curpwd    string `xorm:"'cur_pwd'"`
	Curtime   string `xorm:"'cur_time'"`
	Sshtty    string `xorm:"'ssh_tty'"`
}

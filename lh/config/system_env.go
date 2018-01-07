package config

import (
	"net"
	"os"
	"os/user"
	"regexp"
	"strings"
	"time"
)

type SystemEnv struct {
	Hostname  string
	Loginuser string
	Localip   string
	Fromip    string
	Curuser   string
	Curpwd    string
	Curtime   string
	Sshtty    string
}

func (this *SystemEnv) Setenv() error {
	// use golang net env
	this.Hostname = getHostname()
	this.Loginuser = getCurrentUser()
	this.Curpwd = getCurPwd()
	this.Sshtty = getSshTty()
	this.Localip = getLocalIp()
	this.Fromip = getFromIp()
	this.Curtime = getCurTime()
	this.Curuser = getCurrentUser()
	return nil
}

func getCurrentUser() string {
	current_user_info, err := user.Current()
	if err != nil {
		exec_user := os.Getenv("USER")
		return exec_user
	}
	return current_user_info.Username

}

func getCurPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return "Unknow Cur Pwd."
	}
	return pwd

}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "Unknow Cur Hostname."
	}
	return name
}

func getSshTty() string {
	ssh_tty := os.Getenv("SSH_TTY")
	return ssh_tty
}

func getFromIp() string {
	client := os.Getenv("SSH_CLIENT")
	a := strings.Split(client, " ")
	return a[0]

}

func getLocalIp() string {
	interfaces, _ := net.Interfaces()
	for _, inter_face := range interfaces {
		switch inter_face.Name {
		case "eth0", "em1", "p1p1":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "eth1", "em2", "p1p2":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "eth2", "em3", "p1p3":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "eth3", "em4", "p1p4":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "br0":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "br1":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "br100":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		case "bond0":
			if ok := verifyLocalIp(getInterfaceAddrs(inter_face)); ok == true {
				return getInterfaceAddrs(inter_face)
			}
		}
	}
	/*
	   You can not modify return information
	   Please refer to GetMetricPrefix
	*/
	return "Get Interface Ip Fail"
}

func verifyLocalIp(interip string) bool {
	regular := "^(192|10|172).*$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(interip)
}

func getInterfaceAddrs(inter_face net.Interface) string {
	addrs, err := inter_face.Addrs()
	if err != nil {
		return ""
	}
	for _, ip := range addrs {
		local_ip := strings.Split(ip.String(), "/")[0]
		return local_ip
	}
	return ""
}

func getCurTime() string {
	cur_time := time.Now().Format("2006-01-02 15:04:05")
	return cur_time
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/melbahja/goph"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	time "time"
	"unicode"
)

type bgp struct {
	BgpNeighbors []struct {
		AddressFamily int      `json:"address_family"`
		CustomerAs    int      `json:"customer_as"`
		CustomerIp    string   `json:"customer_ip"`
		Md5Enabled    bool     `json:"md5_enabled"`
		Md5Password   string   `json:"md5_password"`
		Multihop      bool     `json:"multihop"`
		PeerAs        int      `json:"peer_as"`
		PeerIps       []string `json:"peer_ips"`
		RoutesIn      []struct {
			Route string `json:"route"`
			Exact bool   `json:"exact"`
		} `json:"routes_in"`
		RoutesOut []struct {
			Route string `json:"route"`
			Exact bool   `json:"exact"`
		} `json:"routes_out"`
	} `json:"bgp_neighbors"`
}

var ops = "blank"

type devis struct {
	Id              string    `json:"id"`
	ShortId         string    `json:"short_id"`
	Hostname        string    `json:"hostname"`
	Description     string    `json:"description"`
	State           string    `json:"state"`
	Tags            []string  `json:"tags"`
	ImageUrl        string    `json:"image_url"`
	BillingCycle    string    `json:"billing_cycle"`
	User            string    `json:"user"`
	Iqn             string    `json:"iqn"`
	Locked          bool      `json:"locked"`
	BondingMode     int       `json:"bonding_mode"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	SpotInstance    bool      `json:"spot_instance"`
	SpotPriceMax    int       `json:"spot_price_max"`
	TerminationTime time.Time `json:"termination_time"`
	Customdata      struct {
	} `json:"customdata"`
	ProvisioningPercentage int `json:"provisioning_percentage"`
	OperatingSystem        struct {
		Id              string   `json:"id"`
		Slug            string   `json:"slug"`
		Name            string   `json:"name"`
		Distro          string   `json:"distro"`
		Version         string   `json:"version"`
		Preinstallable  bool     `json:"preinstallable"`
		ProvisionableOn []string `json:"provisionable_on"`
		Pricing         struct {
		} `json:"pricing"`
		Licensed bool `json:"licensed"`
	} `json:"operating_system"`
	AlwaysPxe     bool   `json:"always_pxe"`
	IpxeScriptUrl string `json:"ipxe_script_url"`
	Facility      struct {
		Id       string   `json:"id"`
		Name     string   `json:"name"`
		Code     string   `json:"code"`
		Features []string `json:"features"`
		IpRanges []string `json:"ip_ranges"`
		Address  struct {
			Address     string `json:"address"`
			Address2    string `json:"address2"`
			City        string `json:"city"`
			State       string `json:"state"`
			ZipCode     string `json:"zip_code"`
			Country     string `json:"country"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
		} `json:"address"`
		Metro struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country string `json:"country"`
		} `json:"metro"`
	} `json:"facility"`
	Metro struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Code    string `json:"code"`
		Country string `json:"country"`
	} `json:"metro"`
	Plan struct {
		Id          string `json:"id"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Line        string `json:"line"`
		Specs       struct {
		} `json:"specs"`
		Pricing struct {
		} `json:"pricing"`
		Legacy      bool   `json:"legacy"`
		Class       string `json:"class"`
		AvailableIn []struct {
			Href string `json:"href"`
		} `json:"available_in"`
	} `json:"plan"`
	Userdata     string `json:"userdata"`
	RootPassword string `json:"root_password"`
	SwitchUuid   string `json:"switch_uuid"`
	NetworkPorts struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
		Data struct {
		} `json:"data"`
		DisbondOperationSupported bool `json:"disbond_operation_supported"`
		VirtualNetworks           []struct {
			Href string `json:"href"`
		} `json:"virtual_networks"`
		Href string `json:"href"`
	} `json:"network_ports"`
	Href    string `json:"href"`
	Project struct {
		Href string `json:"href"`
	} `json:"project"`
	ProjectLite struct {
		Href string `json:"href"`
	} `json:"project_lite"`
	Volumes []struct {
		Href string `json:"href"`
	} `json:"volumes"`
	HardwareReservation struct {
		Href string `json:"href"`
	} `json:"hardware_reservation"`
	SshKeys []struct {
		Href string `json:"href"`
	} `json:"ssh_keys"`
	IpAddresses []struct {
		Id            string `json:"id"`
		AddressFamily int    `json:"address_family"`
		Netmask       string `json:"netmask"`
		Public        bool   `json:"public"`
		Enabled       bool   `json:"enabled"`
		Cidr          int    `json:"cidr"`
		Management    bool   `json:"management"`
		Manageable    bool   `json:"manageable"`
		GlobalIp      bool   `json:"global_ip"`
		AssignedTo    struct {
			Href string `json:"href"`
		} `json:"assigned_to"`
		Network   string    `json:"network"`
		Address   string    `json:"address"`
		Gateway   string    `json:"gateway"`
		Href      string    `json:"href"`
		CreatedAt time.Time `json:"created_at"`
		Metro     struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country string `json:"country"`
		} `json:"metro"`
		ParentBlock struct {
			Network string `json:"network"`
			Netmask string `json:"netmask"`
			Cidr    int    `json:"cidr"`
			Href    string `json:"href"`
		} `json:"parent_block"`
	} `json:"ip_addresses"`
	ProvisioningEvents []struct {
		Id            string `json:"id"`
		State         string `json:"state"`
		Type          string `json:"type"`
		Body          string `json:"body"`
		Relationships []struct {
			Href string `json:"href"`
		} `json:"relationships"`
		Interpolated string    `json:"interpolated"`
		CreatedAt    time.Time `json:"created_at"`
		Href         string    `json:"href"`
	} `json:"provisioning_events"`
}
type VlanbyPort struct {
	VlanAssignments []struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Native    bool      `json:"native"`
		State     string    `json:"state"`
		Vlan      int       `json:"vlan"`
		Port      struct {
			Href string `json:"href"`
		} `json:"port"`
		VirtualNetwork struct {
			Href string `json:"href"`
		} `json:"virtual_network"`
	} `json:"vlan_assignments"`
}
type privateips struct {
	Available []string `json:"available"`
}
type iprespoce struct {
	IpAddresses []struct {
		Address       string `json:"address"`
		AddressFamily int    `json:"address_family"`
		AssignedTo    struct {
			Href string `json:"href"`
		} `json:"assigned_to"`
		Cidr       int       `json:"cidr"`
		CreatedAt  time.Time `json:"created_at"`
		Enabled    bool      `json:"enabled"`
		Gateway    string    `json:"gateway"`
		GlobalIp   bool      `json:"global_ip"`
		Href       string    `json:"href"`
		Id         string    `json:"id"`
		Manageable bool      `json:"manageable"`
		Management bool      `json:"management"`
		Metro      struct {
			Code    string `json:"code"`
			Country string `json:"country"`
			Id      string `json:"id"`
			Name    string `json:"name"`
		} `json:"metro"`
		Netmask     string `json:"netmask"`
		Network     string `json:"network"`
		ParentBlock struct {
			Cidr    int    `json:"cidr"`
			Href    string `json:"href"`
			Netmask string `json:"netmask"`
			Network string `json:"network"`
		} `json:"parent_block"`
		Public  bool   `json:"public"`
		State   string `json:"state"`
		NextHop string `json:"next_hop"`
	} `json:"ip_addresses"`
}
type vlan struct {
	include string `json:"include"`
	exclude string `json:"exclude"`
}

var ip4 = "blank"

func main() {

	portIfe("776262f3-efaa-4205-a68b-5ecb25640792", "vgcMF1FdSqqp5K6qCyusEXpbBttvSiMR")
}
func dockerinstall(os string) {
	auth, err := goph.Key("/Users/esimon/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, errc := goph.New("root", ip4, auth)
	if errc != nil {
		log.Fatal(errc)

	}
	res1 := strings.Split(os, " ")
	if strings.Compare("Centos", res1[0]) == 0 {
		print("this is centos")
		out9, err := client.Run("sudo yum remove docker \n docker-client \n docker-client-latest \n docker-common \n docker-latest \n docker-latest-logrotate \n docker-logrotate \n docker-engine \n podman \n runc")

		out10, err := client.Run("sudo yum install -y yum-utils")
		out11, err := client.Run("sudo apt-get install ca-certificates curl gnupg")
		out12, err := client.Run("sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo")

		out13, err := client.Run("sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin")
		out14, err := client.Run("sudo systemctl start docker")
		out15, err := client.Run("sudo docker run hello-world")
		fmt.Println(string(out9))
		fmt.Println(string(out10))
		fmt.Println(string(out11))
		fmt.Println(string(out12))
		fmt.Println(string(out13))
		fmt.Println(string(out14))
		fmt.Println(string(out15))

		if err != nil {
			log.Fatal(err)

		}
	} else if strings.Compare("Debian", res1[0]) == 0 {
		print("this debian")
		out9, err := client.Run("for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done")

		out10, err := client.Run("sudo apt-get update")
		out11, err := client.Run("sudo apt-get install ca-certificates curl gnupg")
		out12, err := client.Run("sudo install -m 0755 -d /etc/apt/keyrings")

		out13, err := client.Run("curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg")
		out14, err := client.Run("sudo chmod a+r /etc/apt/keyrings/docker.gpg")
		out15, err := client.Run("echo \\\n  \"deb [arch=\"$(dpkg --print-architecture)\" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \\\n  \"$(. /etc/os-release && echo \"$VERSION_CODENAME\")\" stable\" | \\\n  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null")
		out16, err := client.Run("sudo apt-get update")
		out17, err := client.Run("sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin")
		fmt.Println(string(out9))
		fmt.Println(string(out10))
		fmt.Println(string(out11))
		fmt.Println(string(out12))
		fmt.Println(string(out13))
		fmt.Println(string(out14))
		fmt.Println(string(out15))
		fmt.Println(string(out16))
		fmt.Println(string(out17))
		if err != nil {
			log.Fatal(err)

		}
	} else if strings.Compare("RedHat", res1[0]) == 0 {
		print("this is red hat")
		out9, err := client.Run("sudo yum remove docker \n docker-client \n docker-client-latest \n docker-common \n docker-latest \n docker-latest-logrotate \n docker-logrotate \n docker-engine \n podman \n runc")

		out10, err := client.Run("sudo yum install -y yum-utils")
		out11, err := client.Run("sudo apt-get install ca-certificates curl gnupg")
		out12, err := client.Run("sudo yum-config-manager --add-repo https://download.docker.com/linux/rhel/docker-ce.repo")

		out13, err := client.Run("sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin")
		out14, err := client.Run("sudo systemctl start docker")
		out15, err := client.Run("sudo docker run hello-world")
		fmt.Println(string(out9))
		fmt.Println(string(out10))
		fmt.Println(string(out11))
		fmt.Println(string(out12))
		fmt.Println(string(out13))
		fmt.Println(string(out14))
		fmt.Println(string(out15))

		if err != nil {
			log.Fatal(err)

		}
	} else if strings.Compare("Ubuntu", res1[0]) == 0 {
		print("this is Ubutu")
		out9, err := client.Run("for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done")

		out10, err := client.Run("sudo apt-get update")
		out11, err := client.Run("sudo apt-get install ca-certificates curl gnupg")
		out12, err := client.Run("sudo install -m 0755 -d /etc/apt/keyrings")

		out13, err := client.Run("curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg")
		out14, err := client.Run("sudo chmod a+r /etc/apt/keyrings/docker.gpg")
		out15, err := client.Run("echo \\\n  \"deb [arch=\"$(dpkg --print-architecture)\" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \\\n  \"$(. /etc/os-release && echo \"$VERSION_CODENAME\")\" stable\" | \\\n  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null")
		out16, err := client.Run("sudo apt-get update")
		out17, err := client.Run("sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin")
		fmt.Println(string(out9))
		fmt.Println(string(out10))
		fmt.Println(string(out11))
		fmt.Println(string(out12))
		fmt.Println(string(out13))
		fmt.Println(string(out14))
		fmt.Println(string(out15))
		fmt.Println(string(out16))
		fmt.Println(string(out17))
		if err != nil {
			log.Fatal(err)

		}
	} else {
		out9, err := client.Run("curl https://download.docker.com/linux/static/stable/x86_64/docker-17.03.0-ce.tgz -o /root")
		out10, err := client.Run("tar xzvf /docker-17.03.0-ce.tar.gz")
		out11, err := client.Run(" sudo cp docker/* /root/bin/")
		out12, err := client.Run("sudo dockerd &")
		out13, err := client.Run("sudo docker run hello-world")
		fmt.Println(string(out9))
		fmt.Println(string(out10))
		fmt.Println(string(out11))
		fmt.Println(string(out12))
		fmt.Println(string(out13))
		if err != nil {
			log.Fatal(err)

		}
	}
	client.Close()
	time.Sleep(10 * time.Second)

}

func portIfe(id string, userkey string) {

	getdevisport(id, userkey)
	var url = "https://api.equinix.com/metal/v1/devices/" + id + "/ips"
	req, erro := http.NewRequest("GET", url, nil)
	if erro != nil {
		print(erro.Error())
	}
	var ports iprespoce

	req.Header.Add("X-Auth-Token", userkey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, erro := client.Do(req)

	print(resp.Status + " ")
	erro = json.NewDecoder(resp.Body).Decode(&ports)
	print(" Amount of address")
	print(len(ports.IpAddresses))
	for x := 0; x < len(ports.IpAddresses); x++ {
		print(ports.IpAddresses[x].Address + " ")
		if ports.IpAddresses[x].Public {
			if ports.IpAddresses[x].AddressFamily == 4 {
				//os.Setenv("ip4", ports.IpAddresses[x].Address)
				pingips(ports.IpAddresses[x].Address)
				ip4 = ports.IpAddresses[x].Address
				//str = str[0 : len(str)-3]

			} else if ports.IpAddresses[x].AddressFamily == 6 {
				//os.Setenv("ip6", ports.IpAddresses[x].Address)
				pingips(ports.IpAddresses[x].Address)
				ifer3(ports.IpAddresses[x].Address)
			}

		} else if !ports.IpAddresses[x].Public {
			var urlr = "https://api.equinix.com/metal/v1/ips/" + ports.IpAddresses[x].Id + "/available"
			reqr, erro := http.NewRequest("GET", urlr, nil)
			if erro != nil {
				print(erro.Error())
			}
			var privateip privateips
			reqr.Header.Add("X-Auth-Token", userkey)
			reqr.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			respr, erro := client.Do(reqr)
			fmt.Println(resp.Status + " ")
			fmt.Println("ping private ips")
			erro = json.NewDecoder(respr.Body).Decode(&privateip)
			for x := 0; x < len(privateip.Available); x++ {

				pingips(ip4 + ports.IpAddresses[x].Address + privateip.Available[x])
				print("Amount of possbile private connections")
				print(ports.IpAddresses[x].ParentBlock.Cidr)
			}

			//privatesubnet(ports.IpAddresses[x].Id, string(ports.IpAddresses[x].ParentBlock.Cidr), userkey)
		}
	}
	dockerinstall(ops)
	docksmokeping()
	networktesting()
}

func networktesting() {
	c0 := exec.Command("cmd", "tracert "+ip4)
	if err := c0.Run(); err != nil {

		fmt.Println("Error: ", err)
	}
	print(c0.Output())
	result := ""

	addrs, err := net.LookupIP(ip4)
	if err != nil {
		print(err.Error())
	}

	for _, addr := range addrs {
		result += fmt.Sprintf("%s\n", addr.String())
	}

	print(strings.TrimSpace(result))

}
func docksmokeping() {
	time.Sleep(10 * time.Second)
	auth, err := goph.Key("/Users/esimon/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", ip4, auth)
	if err != nil {
		log.Fatal(err)

	}

	//out18, err := client.Run("docker run hello-world")
	out19, err := client.Run("docker pull dvorak/docker-smokeping")

	out20, err := client.Run("docker run -d \\\n  --name=smokeping \\\n  -e PUID=1000 \\\n  -e PGID=1000 \\\n  -e TZ=Etc/UTC \\\n  -p 80:80 \\\n  -v /path/to/smokeping/config:/config \\\n  -v /path/to/smokeping/data:/data \\\n  --restart unless-stopped \\\n  lscr.io/linuxserver/smokeping:latest\n")
	out21, err := client.Run("docker stop -t 50000 smokeping")
	out22, err := client.Run("docker logs -f smokeping")

	out23, err := client.Run("docker pull networkstatic/iperf3")

	//fmt.Println(string(out18))
	fmt.Println(string(out19))
	fmt.Println(string(out20))
	fmt.Println(string(out21))
	fmt.Println(string(out22))
	fmt.Println(string(out23))
	time.Sleep(10 * time.Second)
	client.Close()
	client3, err := goph.New("root", ip4, auth)
	if err != nil {
		log.Fatal(err)

	}
	//context, cancel := context2.WithTimeout(context2.Background(), time.Minute)
	//defer cancel()

	runContext, err := client3.Run("docker run --name=iperf3-server -p 5000:5000 networkstatic/iperf3 -s")
	out24, err := client3.Run("docker stop -t 50000 iperf3-server")
	out25, err := client3.Run(" docker logs -f iperf3-server")

	if err != nil {
		return
	}
	fmt.Println(string(runContext))

	client3.Close()
	time.Sleep(10 * time.Second)
	client2, err := goph.New("root", ip4, auth)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(string(out24))
	fmt.Println(string(out25))
	iperf3Testing()
	out26, err := client2.Run("docker stop $(docker ps -a -q)")
	out27, err := client2.Run("docker rm $(docker ps -a -q)")
	fmt.Println(string(out26))
	fmt.Println(string(out27))
	client2.Close()

}
func ifer3(address string) {
	//port := 0
	auth, err := goph.Key("/Users/esimon/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", ip4, auth)
	if err != nil {
		log.Fatal(err)

	}

	out, err := client.Run("sudo apt install -y iperf3")
	out1, err := client.Run("sudo ethtool eth0|rg -e 'Speed|Duplex'")
	out2, err := client.Run("lsof -i -P -n | grep LISTEN")
	out3, err := client.Run("ip a")
	out4, err := client.Run("apt-get update")
	out5, err := client.Run("netstat -tulpe")
	out6, err := client.Run("cat ~/.ssh/config")
	out7, err := client.Run("echo 'Host myserver\n HostName " + client.LocalAddr().String() + "\n Port 59519\n User dev\n'")

	holder := string(out3)
	inline := true
	start := false
	building := ""
	for x := 1; x < len(holder); x++ {

		print(" ")
		if []rune(holder)[x] == 10 {

			inline = true
			start = false
		}
		if unicode.IsDigit([]rune(holder)[x]) && !start && inline && []rune(holder)[x-1] == 32 {

			fmt.Printf("%c", []rune(holder)[x])
			print([]rune(holder)[x])
			yo := string([]rune(holder)[x])
			building = building + yo
			start = true
		} else if start && inline && unicode.IsDigit([]rune(holder)[x]) {
			yo := string([]rune(holder)[x])
			building = building + yo
		} else if start && inline && !unicode.IsDigit([]rune(holder)[x]) {
			print(building)
			outr, err := client.Run("lsof -i:" + building)
			print(string(outr))
			building = ""
			inline = false
			start = false
			if err != nil {
				print(":this port is host side")

			}

		}
	}

	fmt.Println(string(out))
	fmt.Println(string(out1))
	fmt.Println(string(out2))
	fmt.Println(string(out3))
	fmt.Println(string(out4))
	fmt.Println(string(out5))
	fmt.Println("ssh info")
	fmt.Println(string(out6))
	fmt.Println(string(out7))

	client.Close()

}

func pingips(addess string) {
	out, _ := exec.Command("ping", addess, "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		fmt.Println("_ " + addess + "is down")
	} else {
		fmt.Println("_" + addess + "is live")
	}
}
func getdevisport(uuid string, userkey string) {
	url := "https://api.equinix.com/metal/v1/devices/" + uuid
	res, erro := http.NewRequest("GET", url, nil)
	if erro != nil {
		print(erro.Error())
	}
	res.Header.Add("X-Auth-Token", userkey)
	res.Header.Set("Content-Type", "application/json")
	var devis devis
	client := &http.Client{}
	resp, erro := client.Do(res)
	print(resp.Status + " ")
	erro = json.NewDecoder(resp.Body).Decode(&devis)
	print(devis.OperatingSystem.Name)
	for x := 0; x < len(devis.NetworkPorts.VirtualNetworks); x++ {
		print(devis.NetworkPorts.VirtualNetworks[x].Href)

	}

	fmt.Println(devis.Metro)
	ops = devis.OperatingSystem.Name
	getvlan(devis.Plan.Id, userkey)
	getbgp(uuid, userkey)
}

func iperf3Testing() {
	fmt.Println("testing")
	var c0 *exec.Cmd
	var c1 *exec.Cmd
	var c2 *exec.Cmd
	var outb, outa, erra, errb, errc, outc bytes.Buffer
	switch runtime.GOOS {
	case "windows":

		c0 = exec.Command("cmd", "docker pull networkstatic/iperf3")
		c1 = exec.Command("cmd", "docker run  -it --rm networkstatic/iperf3 -c "+ip4)
		c2 = exec.Command("cmd", "docker stop -t 50000 networkstatic/iperf3")
	default: //Mac & Linux
		c0 = exec.Command("rm", "docker pull networkstatic/iperf3")
		c1 = exec.Command("rm", "docker run  -it --rm networkstatic/iperf3 -c "+ip4)
		c2 = exec.Command("rm", "docker stop -t 50000 networkstatic/iperf3")

	}
	if err := c0.Run(); err != nil {

		fmt.Println("Error: ", err)
	}
	if err := c1.Run(); err != nil {

		fmt.Println("Error: ", err)
	}
	if err := c2.Run(); err != nil {

		fmt.Println("Error: ", err)
	}
	time.Sleep(100 * time.Millisecond)
	c0.Stdout = &outa
	c1.Stdout = &outb
	c2.Stdout = &outc
	c0.Stderr = &erra
	c1.Stderr = &errb
	c2.Stderr = &errc
	fmt.Println("Outpute " + outa.String() + " Error " + erra.String())
	fmt.Println("Outpute " + outb.String() + " Error " + errb.String())
	fmt.Println("Outpute " + outc.String() + " Error " + errc.String())
}
func getbgp(uid string, userkey string) {
	url := "https://api.equinix.com/metal/v1/devices/" + uid + "/bgp/neighbors"
	req, erro := http.NewRequest("GET", url, nil)
	if erro != nil {
		print(erro.Error())
	}
	req.Header.Add("X-Auth-Token", userkey)
	req.Header.Set("Content-Type", "application/json")
	var port bgp
	client := &http.Client{}
	resp, erro := client.Do(req)
	erro = json.NewDecoder(resp.Body).Decode(&port)
	for x := 0; x < len(port.BgpNeighbors); x++ {
		print(port.BgpNeighbors[x].CustomerIp)
		print(port.BgpNeighbors[x].AddressFamily)
		for y := 0; y < len(port.BgpNeighbors[x].PeerIps); y++ {
			print(port.BgpNeighbors[x].PeerIps[y])
		}
		for z := 0; z < len(port.BgpNeighbors[x].RoutesIn); z++ {

			print(port.BgpNeighbors[x].RoutesIn[z].Route)

		}
		for z1 := 0; z1 < len(port.BgpNeighbors[x].RoutesOut); z1++ {
			print(port.BgpNeighbors[x].RoutesIn[z1].Route)

		}

	}
}
func getvlan(uid string, userkey string) {
	url := "https://api.equinix.com/metal/v1/ports/" + uid + "/vlan-assignments"
	encoding := vlan{include: "virtual_network"}
	encoded, erro := json.Marshal(encoding) // encoder
	os.Stdout.Write(encoded)
	if erro != nil {
		print(erro.Error())
	}
	req, erro := http.NewRequest("GET", url, bytes.NewBuffer(encoded))
	req.Header.Add("X-Auth-Token", userkey)
	req.Header.Set("Content-Type", "application/json")
	var thisport VlanbyPort
	client := &http.Client{}
	resp, erro := client.Do(req)
	print(resp.Status + " ")
	erro = json.NewDecoder(resp.Body).Decode(&thisport)
	for x := 0; x < len(thisport.VlanAssignments); x++ {
		print("" + thisport.VlanAssignments[x].VirtualNetwork.Href)
	}

}

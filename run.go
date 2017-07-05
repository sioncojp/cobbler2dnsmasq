package c2d

// Host ...全てのサーバをsliceで保持する
type Host struct {
	Info []Info
}

// Info ...1つずつのサーバの情報
type Info struct {
	HostName   string
	MACAddress string
	IP         string
}

var ConfigFile string
var config Config

// Run ...dnsmasq.template作成するためにメインタスクを実行する
func Run() error {
	i := []Info{}
	h := Host{i}

	if err := LoadConfig(ConfigFile, &config); err != nil {
		return err
	}

	err := h.GetInfo()
	if err != nil {
		return err
	}

	h.CreateTemplate()

	return nil
}

package c2d

import "github.com/BurntSushi/toml"

// Config ...tomlファイルの情報
type Config struct {
	// TemplateFile ...flagで渡されたtemplateファイルの場所
	TemplateFile string `toml:"template_file"`

	// CobblerDir ...flagで渡されたCobblerディレクトリの場所
	CobblerDir string `toml:"cobbler_dir"`

	// Output ...アウトプット先
	Output string `toml:"output"`

	// CobblerSync ...cobbler syncコマンド
	CobblerSync string `toml:"cobbler_sync"`

	// MACregexp ...MACアドレスの正規表現
	MACregexp string `toml:"mac_regexp"`

	// Hostnameregexp ...Hostnameの正規表現
	Hostnameregexp string `toml:"hostname_regexp"`

	// IPregexp ...IPアドレスの正規表現
	IPregexp string `toml:"ip_regexp"`

	// IgnoreIPs ...DNSやらのIPを除外するため
	IgnoreIPs []string `toml:"ignore_ips"`
}

// LoadConfig ...tomlファイルを読み込む
func LoadConfig(confPath string, config *Config) error {
	if _, err := toml.DecodeFile(confPath, &config); err != nil {
		return err
	}
	return nil
}

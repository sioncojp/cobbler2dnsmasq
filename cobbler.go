package c2d

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

// GetAllKsFilename ...ksファイルの名前を抽出
func GetAllKsFilename() ([]string, error) {
	// Cobblerのprofileがあるディレクトリを開く。なければエラー
	files, err := ioutil.ReadDir(config.CobblerDir)
	if err != nil {
		return nil, errors.New("cannot open profile directory")
	}

	ks := []string{}

	// ファイル名を抽出
	for _, file := range files {
		ks = append(ks, file.Name())
	}

	return ks, nil
}

// MatchStrings ...ksファイルに対して文字列をマッチさせたものをsliceとして取り出す
func MatchStrings(name, reg string) ([]string, error) {
	// profileをOpen。なければエラー
	f, err := os.Open(config.CobblerDir + "/" + name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	macaddress := []string{}

	// 引数の正規表現でマッチした文字列を配列に渡す
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r := regexp.MustCompile(reg)
		if r.MatchString(scanner.Text()) {
			macaddress = append(macaddress, r.FindString(scanner.Text()))
		}
	}

	return macaddress, nil
}

// GetInfo ...MACアドレス, Hostname, IPを抽出
func (h *Host) GetInfo() error {
	i := Info{}
	files, err := GetAllKsFilename()
	if err != nil {
		return err
	}

	for _, f := range files {
		// MACアドレス取得
		MACAddress, err := MatchStrings(f, config.MACregexp)
		if err != nil {
			return err
		}
		// hostname取得
		Hostname, err := MatchStrings(f, config.Hostnameregexp)
		if err != nil {
			return err
		}
		// 指定されたIPを取得
		IP, err := MatchStrings(f, config.IPregexp)
		if err != nil {
			return err
		}

		// profileの一番下に書いてあるMACアドレスを取得。なければ終了
		if len(MACAddress) == 0 {
			break
		} else {
			i.MACAddress = MACAddress[len(MACAddress)-1]
		}

		// profileの最初に書いてあるhostnameを取得。なければ終了
		if len(Hostname) == 0 {
			break
		} else {
			i.HostName = Hostname[0]
		}

		// 指定されたIPを取得し、除外IPのチェックをする。そもそもIPがなければ終了
		if len(IP) == 0 {
			break
		} else {
			i.IP = ExcludeIgnoreIP(IP)

			// DNSやらのIPを除外した後、IPが空の場合は終了
			if i.IP == "" {
				break
			}

		}
		h.AddItem(i)
	}
	return nil
}

// AddItem ...取得したサーバ情報を１つずつ、Hostにいれていく。多次元配列的な処理。
func (h *Host) AddItem(i Info) []Info {
	h.Info = append(h.Info, i)
	return h.Info
}

// ExcludeIgnoreIP ...DNSやらのIPを除外する
func ExcludeIgnoreIP(ips []string) string {
	for i := 0; i < len(ips); i++ {
		ip := ips[i]
		for _, ignoreip := range config.IgnoreIPs {
			if ip == ignoreip {
				ips = append(ips[:i], ips[i+1:]...)
				i--
				break
			}
		}
	}

	// 除外されてなにも残らなければ空文字を返す
	if len(ips) == 0 {
		return ""
	}
	// 残ったIPを返す
	return ips[0]
}

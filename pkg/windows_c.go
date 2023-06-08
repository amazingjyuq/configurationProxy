package pkg

import (
	"golang.org/x/sys/windows/registry"
)

var (
	proxySettingsKey registry.Key
)

func init() {
	var err error
	proxySettingsKey, err = registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		// 处理打开注册表键时的错误
		panic(err)
	}
}

func setProxySettings(proxyServer string) error {
	// 设置代理服务器和端口
	if err := proxySettingsKey.SetStringValue("ProxyServer", proxyServer); err != nil {
		return err
	}

	return nil
}

func enableProxySettings() error {
	// 启用代理
	if err := proxySettingsKey.SetDWordValue("ProxyEnable", 1); err != nil {
		return err
	}

	return nil
}

func disableProxySettings() error {
	// 禁用代理
	if err := proxySettingsKey.SetDWordValue("ProxyEnable", 0); err != nil {
		return err
	}

	return nil
}

func getProxySettings() (string, uint32, uint32, error) {
	// 读取代理设置
	proxyServer, _, err := proxySettingsKey.GetStringValue("ProxyServer")
	if err != nil {
		return "", 0, 0, err
	}

	proxyPort, _, err := proxySettingsKey.GetIntegerValue("ProxyPort")
	if err != nil {
		return "", 0, 0, err
	}

	proxyEnable, _, err := proxySettingsKey.GetIntegerValue("ProxyEnable")
	if err != nil {
		return "", 0, 0, err
	}

	return proxyServer, uint32(proxyPort), uint32(proxyEnable), nil
}
package pkg

import (
	"fmt"
	"os/exec"
)

func MacOSSetProxySettings(proxyServer string, proxyPort int) error {
	cmd := exec.Command("networksetup", "-setwebproxy", "Wi-Fi", proxyServer, fmt.Sprintf("%d", proxyPort))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set web proxy: %v", err)
	}

	cmd = exec.Command("networksetup", "-setsecurewebproxy", "Wi-Fi", proxyServer, fmt.Sprintf("%d", proxyPort))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set secure web proxy: %v", err)
	}

	cmd = exec.Command("networksetup", "-setwebproxystate", "Wi-Fi", "on")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to enable web proxy: %v", err)
	}

	cmd = exec.Command("networksetup", "-setsecurewebproxystate", "Wi-Fi", "on")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to enable secure web proxy: %v", err)
	}

	return nil
}

func MacOSDisableProxySettings() error {
	cmd := exec.Command("networksetup", "-setwebproxystate", "Wi-Fi", "off")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to disable web proxy: %v", err)
	}

	cmd = exec.Command("networksetup", "-setsecurewebproxystate", "Wi-Fi", "off")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to disable secure web proxy: %v", err)
	}

	return nil
}
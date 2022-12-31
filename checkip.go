package ext

// Get Public IP address
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var (
	akamaiCheckIPURI = "http://whatismyip.akamai.com/"
	amazonCheckIPURI = "https://checkip.amazonaws.com/"
	// orayCheckIPURI = "https://ddns.oray.com/checkip"
)

func HttpGetPublicIPAddr(reqURL string) (string, error) {
	// 读取在线通用配置文件, 如果失败，使用代理尝试
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request %s, status %s", reqURL, resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("request %s, read response body failed", reqURL)
	}

	exp := regexp.MustCompile(`((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`)
	addr := exp.FindString(string(b))
	if addr == "" {
		return addr, fmt.Errorf("request %s, ip addr is not found", reqURL)
	}

	return addr, nil
}

func GetPublicIPAddr() (addr string, err error) {
	if addr, err = HttpGetPublicIPAddr(akamaiCheckIPURI); err == nil {
		return addr, err
	}
	return HttpGetPublicIPAddr(amazonCheckIPURI)
}

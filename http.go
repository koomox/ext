package ext

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/proxy"
)

func NewProxyDialer(network, addr string) *http.Client {
	dialer, _ := proxy.SOCKS5(network, addr, nil, proxy.Direct)
	transport := &http.Transport{
		Dial: dialer.Dial,
	}
	return &http.Client{Transport: transport}
}

func GetWithDialer(resource string, dialer *http.Client) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := dialer.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %s, status %s", resource, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func Get(resource string) ([]byte, error) {
	return GetWithDialer(resource, &http.Client{Timeout: 5 * time.Second})
}

func DownloadWithDialer(resource, dst string, dialer *http.Client) (err error) {
	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := dialer.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request %s, status %s", resource, resp.Status)
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

func Download(resource, dst string) (err error) {
	return DownloadWithDialer(resource, dst, &http.Client{Timeout: 5 * time.Second})
}

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

func HttpGetWithDialer(reqURL string, dialer *http.Client) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
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
		return nil, fmt.Errorf("request %s, status %s", reqURL, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func HttpGet(reqURL string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %s, status %s", reqURL, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func HttpGetWithFile(reqURL, dst string) (err error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request %s, status %s", reqURL, resp.Status)
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

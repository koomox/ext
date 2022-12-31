package ext

// Get Public IP address
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

var (
	AkamaiCheckIPURL = "http://whatismyip.akamai.com/"
	AmazonCheckIPURL = "https://checkip.amazonaws.com/"
	OrayCheckIPURL   = "https://ddns.oray.com/checkip"
)

func GetPublicIPAddr(host ...string) (addr string, err error) {
	client := &http.Client{Timeout: 3 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ch := make(chan string)
	for i := range host {
		go func(ctx context.Context, reqURL string) {
			req, err := http.NewRequest(http.MethodGet, reqURL, nil)
			if err != nil {
				return
			}
			resp, err := client.Do(req.WithContext(ctx))
			if err != nil {
				return
			}

			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return
			}
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			exp := regexp.MustCompile(`((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`)
			ipaddr := exp.FindString(string(buf))
			if ipaddr == "" {
				return
			}
			ch <- ipaddr
		}(ctx, host[i])
	}
	select {
	case info := <-ch:
		cancel()
		return info, err
	case <-time.After(5 * time.Second):
		cancel()
		return "", fmt.Errorf("timeout")
	}
}

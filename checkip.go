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
	akamaiCheckIPURL = "http://whatismyip.akamai.com/"
	amazonCheckIPURL = "https://checkip.amazonaws.com/"
	orayCheckIPURL   = "https://ddns.oray.com/checkip"
)

func GetPublicIPAddr(host ...string) (addr string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ch := make(chan string, 1)
	host = append(host, akamaiCheckIPURL, amazonCheckIPURL)
	for i := range host {
		go func(ctx context.Context, resource string, result chan string) {
			req, err := http.NewRequest(http.MethodGet, resource, nil)
			if err != nil {
				return
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:113.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
			resp, err := http.DefaultClient.Do(req.WithContext(ctx))
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
			result <- ipaddr
		}(ctx, host[i], ch)
	}
	select {
	case info := <-ch:
		cancel()
		return info, err
	case <-time.After(3 * time.Second):
		cancel()
		return "", fmt.Errorf("timeout")
	}
}

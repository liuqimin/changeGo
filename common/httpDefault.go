package common

import (
	"net"
	"net/http"
	"time"
)

const (
	MaxIdleConns        = 100
	MaxIdleConnsPerHost = 300
	IdleConnTimeout     = 20
)

func HttpDo() (client *http.Client) {
	client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true, // 表示是否开启http keepalive功能，也即是否重用连接，默认开启(false)
			Proxy:             http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: 6 * time.Second,
	}
	return
	/*
		logs.Info("Report URL:", beego.AppConfig.String("cmdbUrl"), client_ip)
		logs.Debug("Report Data:", string(data))
		req, err := http.NewRequest("POST", beego.AppConfig.String("cmdbUrl"), bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			logs.Error("asset report, err:", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logs.Error("asset report, err:", err)
				return err
			}
			logs.Debug("asset report reslut,", string(body))
		}
		return
	*/
}

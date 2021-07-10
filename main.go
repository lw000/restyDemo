package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"log"
)

const (
	HeaderGameToken = "GameToken"
	HeaderPlayerId  = "PlayerId"
	HeaderFpid      = "Fpid"
)

var (
	GameToken string
	PlayerId  string
	Fpid      string = "levi001"
)

func main() {
	client := resty.New()

	LoginNativeRequest(client)
	ThemeListRequest(client)
	TaskRequest(client)
}

func test(client *resty.Client) {
	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	if err != nil {
		log.Println(err)
		return
	}
	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()

	// Explore trace info
	log.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	log.Println("  DNSLookup     :", ti.DNSLookup)
	log.Println("  ConnTime      :", ti.ConnTime)
	log.Println("  TCPConnTime   :", ti.TCPConnTime)
	log.Println("  TLSHandshake  :", ti.TLSHandshake)
	log.Println("  ServerTime    :", ti.ServerTime)
	log.Println("  ResponseTime  :", ti.ResponseTime)
	log.Println("  TotalTime     :", ti.TotalTime)
	log.Println("  IsConnReused  :", ti.IsConnReused)
	log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	log.Println("  RequestAttempt:", ti.RequestAttempt)
	log.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func LoginNativeRequest(client *resty.Client) {
	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"name": "levi001",
		}).
		SetHeader("Accept", "application/json").
		Post("http://10.0.84.174:9000/login/auto/msg.LoginNativeRequest")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("  Error      :", err)
	log.Println("  resp      :", resp)

	var loginResult map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &loginResult); err == nil {
		GameToken = cast.ToString(loginResult["Token"])
		PlayerId = cast.ToString(loginResult["PlayerID"])
	}
}

func ThemeListRequest(client *resty.Client) {
	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader(HeaderGameToken, GameToken).
		SetHeader(HeaderPlayerId, PlayerId).
		SetHeader(HeaderFpid, Fpid).
		Post("http://10.0.84.174:9000/gateway/auto/msg.ThemeListRequest")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("  Error      :", err)
	log.Println("  resp      :", resp)
}

func TaskRequest(client *resty.Client) {
	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader(HeaderGameToken, GameToken).
		SetHeader(HeaderPlayerId, PlayerId).
		SetHeader(HeaderFpid, Fpid).
		Post("http://10.0.84.174:9000/lobby_batch/auto/msg.TaskRequest")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("  Error      :", err)
	log.Println("  resp      :", resp)
}

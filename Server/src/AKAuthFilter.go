package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var CAs = [...]string{`-----BEGIN CERTIFICATE-----
MIIFozCCA4ugAwIBAgIUKiJ5AUO+JQmpTpg08enKXxn2qdowDQYJKoZIhvcNAQEL
BQAwYDELMAkGA1UEBhMCQ04xEDAOBgNVBAgMB0ppYW5nU3UxDTALBgNVBAcMBFd1
WGkxEzARBgNVBAoMCkFsd2F5c0tpbmcxGzAZBgNVBAMMEkFsd2F5c0tpbmcgUm9v
dCBDQTAgFw0yMDAzMzExMjQ0NDRaGA8yMTIwMDMwNzEyNDQ0NFowYDELMAkGA1UE
BhMCQ04xEDAOBgNVBAgMB0ppYW5nU3UxDTALBgNVBAcMBFd1WGkxEzARBgNVBAoM
CkFsd2F5c0tpbmcxGzAZBgNVBAMMEkFsd2F5c0tpbmcgUm9vdCBDQTCCAiIwDQYJ
KoZIhvcNAQEBBQADggIPADCCAgoCggIBAML75LQSlPkZav/sbi2NjL1skVqJ0Y3Q
drXHnPqXhVsHejxc7zT68WyGEcc2ouAGQvxdtiLQhQvAOfg4FSRoGXilP35ncjYv
Ko34mN6aW9QD9+vt5sdafgv0Ou170Q9H8pjE9mh/MeALcnlWNCxL09EFjiLvKsnS
x4i44d6St16Fusa3PZTL4myDswuxjGsrwQJ8cJsfOBxbSJPPfMbeDOXybGfBPnuJ
y9fWfEBjnQm79xW98tdCojqngYjGXIfPipiCGX4YgdZnT9jfSitqzXJ1afbXXrg2
UF12duJZfqeg8gtTtcHTqVSTVF67itYbQ4croDgHXYUY1kWFjAgMOZxL9D5m2sxw
5yxW7PbMQQgbRe8c3pZY/fCcw2bB9OEvBHLrk8Deq/PMDTdjC5VtxU0P7g77Y9hq
2XLJHb94ySfZKK4MZc9NBY0Mm5YvLvEf0mTvHdo6Y4xjQrK0acFStnjH7MpCw2+q
T9RCGouki/TpxaU5itQAU0l2qBuDEEi22SCwr847cDVVxcm+4s9XaLdFhzDYXCgk
jPcXI4vM2iNjV8d77RHTaAmZkPCpHPlwLDqi36Qg6Bvm+pga8QMm8C4FAP6AF1TT
fO1kxGz5EdHl8D6ibZkMopvTtKb3IIyxaPOaA4OOBQaJ0SZXhGhdoMuEhWb6QHs2
H4cfb4yoivZzAgMBAAGjUzBRMB0GA1UdDgQWBBTIfJ6uUN5NTO9Bc7zw1/Yma2ou
cTAfBgNVHSMEGDAWgBTIfJ6uUN5NTO9Bc7zw1/Yma2oucTAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4ICAQCsNlVlIhhDN75HsXTk19ryJserIduaBwuw
EStGs7lbJ1os37M41KCAvCJYIY3Nlxds7h5SG/HqxX6gElgfGPOL/sHyQrctjW0q
KLfzx82GvwLHD+wOOG36cMeOEm4XLTbAPI+lGTXyz0ZisCGrjQBdsDI8ndrXzKcM
R5DYXK8IedZY+FhfMB8rFW7JmMdxUZvmxKx1i+bNmiVJpeZX3mZ5WrqreOMUcE9w
hEasXWGTvPtKhT3TJB2X9ctF8sYCzD04B4J4lkbOT2D7qwOD6WSTisUxUX6jV0X0
uKVeNVdI75G+AjJ7kdtIlYUsWCJaZUPj0hJCdMGC7d86xC5HCkV8V7WstgmQC/0O
xKyidzF+Do6n0HjTS3sYPxWiIYV5dj8atUZ1I3Z6Sarj6LsmeYfGznzf6fUSAUwf
M+EjlPukk88CeSHF7zu/wVw09wuFiNFBM0QBug2V6T3f57Dyl8Tw345beJ6/Q1xg
HgEClk4H6924Hk4hdMepGx6RHeE05Dw1Z2GHQ9y1CcM4J+WHWUNmGzurvYPdGcWC
2AAHwdKBwR17CQADAL/STCyByFR1bIml1M2t08FasyaVzGo6CvA0/+e6k5m4cbZk
YdMATQhtl4UQzYDmTcmPERWAxeh7VAPWXTxoqIeCFn09JyzzX4F1LGTRoYFSNu9G
qo0AthzZmg==
-----END CERTIFICATE-----
`,
	`-----BEGIN CERTIFICATE-----
MIIEgTCCAmmgAwIBAgIBAzANBgkqhkiG9w0BAQsFADBgMQswCQYDVQQGEwJDTjEQ
MA4GA1UECAwHSmlhbmdTdTENMAsGA1UEBwwEV3VYaTETMBEGA1UECgwKQWx3YXlz
S2luZzEbMBkGA1UEAwwSQWx3YXlzS2luZyBSb290IENBMB4XDTIwMDMzMTEyNDc1
MloXDTI4MDYxNzEyNDc1MlowUzELMAkGA1UEBhMCQ04xEDAOBgNVBAgMB0ppYW5n
U3UxEzARBgNVBAoMCkFsd2F5c0tpbmcxHTAbBgNVBAMMFEFsd2F5c0tpbmcgU2Vj
b25kIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk/9eHgzO1A0o
0W5beCWeL6PUrrntrbF4p9pA9pUei7M6vpxscHc5gvhGiWolB06NR0dbUDhhs6Tf
2gP7836IphipfFp/qYp0R30m7UP9j0yC/bfnG7xoq0RcRAiWW5mtpwV3lC6uA4vq
MEarF8UAVzcG73nU+wng8oraIQKpiZwKw7zFLdQuoIZm8Ov6e8u3xBfovhw3tFZI
y3VCkJmx1QrXPHkmAYuXJA1BJkAif6erg1HWRuNepMUrTlinz1TZdAvg5sTgjQbB
El44KyGsSOXv7CfFgSEUcj7xPplsFfpBAoj+5xfrlvIZ9UkIecDAcM5/JFZsIgqW
hduwrx2WAwIDAQABo1MwUTAdBgNVHQ4EFgQUB1n8PomFglRYNx7BtKHTsSh4A6Uw
HwYDVR0jBBgwFoAUyHyerlDeTUzvQXO88Nf2JmtqLnEwDwYDVR0TAQH/BAUwAwEB
/zANBgkqhkiG9w0BAQsFAAOCAgEAF7oL7N810KSu9YZzzjyI9f5qyuwg4W8xSH1K
ctPCuoFm17H+iykmerQ34SEFALW90G0k4MQ5/Vw/n37Uvja1KAzniqQZDwgCUC8d
chV5zSn/gqF0zWkOC//mEL9sgd8UPQ07h1MIU825WX5G/MyNxlHjsMCCOd0+yfmu
C+cP9/nu0ZdpAWTM0ovnFlEEz69GcpeyL/q6JryqOg0KS8uWz8Ibt1FtaZxOgkGo
e2SXqxmvUXxIimwI0PTdeKI9U97noJYI5nOA1kIBx8jIO6OMNF2t4gM/+j+M7U3x
j5hfc+YziriOUF1OtNUUhZlIEQ6cFbXiTcFKOr5/n9Bi9TGb8ifo0V/7s/y4crlx
JKdKnCzI0RSA0ciom9vVIivFqiYoYdlwf+NqFPVrQN6WHzTRi1KxGaAmIt4+usrC
adjrIcO5iHjrLwitM+Nwdc1IVPqKYcW6hdVYWYeE8LYz4RkNYbSkR//1alJKcetI
TLT8M/leoWZDNGc/t7xb6AL6oKjaYKafRj6m3dl0FqcQXelQtPZ6M9LR/r+DOho9
OxgZ1Z1gJbXzHAz8wNLLWhhVaUe+aAeawMC5Ndaqj5krbBy5wbdmunw5k8L363x2
C+iadC+m7ymNFnh2EKyllQ81pP53CTk7N44/69puP39+2pnP97juSU7vexIBlFMm
IamuV/Q=
-----END CERTIFICATE-----
`}

type FilterHandler func(http.ResponseWriter, *http.Request) bool

type AKAuthFilter struct {
	AppID        string // APPID
	Token        string
	AppTokenName string
	AppHost      string
	AuthGrantCB  string           // 验证回调的方法名
	AuthRevokeCB string           // 许可撤销的地址
	Handler      http.Handler     // 处理真是网页请求的Handler
	ErrPage      string           // 错误页面地址
	AuthServer   string           // 验证服务器的地址
	Trusted      map[string]int64 // 缓存的已经信任的Token
	ErrLog       func(a ...interface{})
}

var AKAuthPost *http.Client

func init() {
	// Load CA cert
	caCertPool := x509.NewCertPool()
	for _, ca := range CAs {
		caCertPool.AppendCertsFromPEM([]byte(ca))
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	AKAuthPost = &http.Client{Transport: transport}
}

func (auth *AKAuthFilter) ErrorOut(a ...interface{}) {
	if auth.ErrLog != nil {
		auth.ErrLog(a...)
	}
}

func (auth *AKAuthFilter) CkeckToken(AppToken string) bool {
	value, ok := auth.Trusted[AppToken]
	if ok {
		// 检测时间
		if time.Now().Unix() > value {
			return false
		} else {
			return true
		}
	} else {
		// 发起网络验证请求
		resp, err := AKAuthPost.Post(auth.AuthServer+"/CheckToken", "application/json", bytes.NewReader([]byte(fmt.Sprintf(`{"AppID":"%s", "Token":"%s","AppToken":"%s"}`, auth.AppID, auth.Token, AppToken))))
		if err != nil {
			auth.ErrorOut("[AKAuth] [CkeckToken] %s", err.Error())
			return false
		}
		RespData, err := io.ReadAll(resp.Body)
		if err != nil {
			auth.ErrorOut("[AKAuth] [CkeckToken] %s", err.Error())
			return false
		}
		if string(RespData) == `{"Status":"success"}` {
			if auth.Trusted == nil {
				auth.Trusted = make(map[string]int64)
			}
			auth.Trusted[AppToken] = time.Now().Add(time.Hour * 24 * 7).Unix()
			return true
		} else {
			auth.ErrorOut("[AKAuth] [CkeckToken] %s", string(RespData))
			return false
		}
	}
}

func (auth *AKAuthFilter) RevokeToken(AppToken string, RevokeAuth bool) bool {
	request := ""
	if RevokeAuth {
		request = fmt.Sprintf(`{"AppID":"%s", "Token":"%s",  "AppToken":"%s", "RevokeAuth":true}`, auth.AppID, auth.Token, AppToken)
	} else {
		request = fmt.Sprintf(`{"AppID":"%s", "Token":"%s",  "AppToken":"%s", "RevokeAuth":false}`, auth.AppID, auth.Token, AppToken)
	}
	resp, err := AKAuthPost.Post(auth.AuthServer+"/RevokeToken", "application/json", bytes.NewReader([]byte(request)))
	if err != nil {
		auth.ErrorOut("[AKAuth] [RevokeToken] %s", err.Error())
		return false
	}
	RespData, err := io.ReadAll(resp.Body)
	if err != nil {
		auth.ErrorOut("[AKAuth] [RevokeToken] %s", err.Error())
		return false
	}
	if string(RespData) == `{"Status":"success"}` {
		return true
	} else {
		auth.ErrorOut("[AKAuth] [RevokeToken] %s", err.Error())
		return false
	}
}

func (auth *AKAuthFilter) RedirectLogin(w http.ResponseWriter, r *http.Request, OriginalURL string) {
	// 获取当前地址
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	if auth.AppHost == "" {
		auth.AppHost = scheme + r.Host
	}

	RedirectURL := base64.StdEncoding.EncodeToString([]byte(auth.AppHost + auth.AuthGrantCB + "?url=" + OriginalURL))
	http.Redirect(w, r, fmt.Sprintf("%s?appid=%s&url=%s", auth.AuthServer, auth.AppID, RedirectURL), http.StatusSeeOther)
	auth.ErrorOut("[AKAuth] [RedirectLogin] %s", OriginalURL)
}

func (auth *AKAuthFilter) RedirectError(w http.ResponseWriter, r *http.Request, Msg string) {
	if len(auth.ErrPage) == 0 {
		w.Write([]byte(Msg))
	} else {
		http.Redirect(w, r, fmt.Sprintf("%s?msg=%s", auth.ErrPage, Msg), http.StatusSeeOther)
	}
	auth.ErrorOut("[AKAuth] [RedirectError] %s", Msg)
}

func (auth *AKAuthFilter) AKAuthFilter(w http.ResponseWriter, r *http.Request) bool {
	// 如果验证页面
	if r.URL.Path == auth.AuthGrantCB {
		r.ParseForm()
		AppTokenList := r.Form["token"]
		OriginalURLList := r.Form["url"]
		if len(AppTokenList) != 1 || len(OriginalURLList) != 1 {
			// 验证有效
			auth.RedirectError(w, r, "无效参数")
			auth.ErrorOut("[AKAuth] [AuthGrantCB] 无效的参数")
			return false
		}

		AppToken := AppTokenList[0]
		OriginalURLBase64 := OriginalURLList[0]

		// Token无效则重新登录
		if !auth.CkeckToken(AppToken) {
			auth.RedirectLogin(w, r, OriginalURLBase64)
			auth.ErrorOut("[AKAuth] [AuthGrantCB] Token 检测失败")
			return false
		}

		// 否则重定向回原始地址
		OriginalURLData, err := base64.StdEncoding.DecodeString(OriginalURLBase64)
		if err != nil {
			// 返回失败
			auth.RedirectError(w, r, "无效原始地址")
			auth.ErrorOut("[AKAuth] [AuthGrantCB] 无效原始地址")
			return false
		}

		// 设置cookie
		http.SetCookie(w, &http.Cookie{Name: auth.AppTokenName, Value: AppToken})
		http.Redirect(w, r, string(OriginalURLData), http.StatusSeeOther)
		return false
	} else if r.URL.Path == auth.AuthRevokeCB {
		r.ParseForm()
		type RevokeRequest struct {
			AppToken string
			Token    string
		}

		RequestData, err := io.ReadAll(r.Body)
		if err != nil {
			auth.ErrorOut("[AKAuth] [AuthRevokeCB] 获取数据失败")
			return false
		}
		Request := &RevokeRequest{}
		if json.Unmarshal(RequestData, Request) != nil {
			auth.ErrorOut("[AKAuth] [AuthRevokeCB] 无效的参数")
			return false
		}

		if Request.AppToken == auth.Token {
			delete(auth.Trusted, Request.Token)
		}

		return false
	}

	// 获取AppToken
	AppToken, err := r.Cookie(auth.AppTokenName)
	if err != nil || !auth.CkeckToken(AppToken.Value) {
		// 重定向
		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		if auth.AppHost == "" {
			auth.AppHost = scheme + r.Host
		}

		OriginalURL := base64.StdEncoding.EncodeToString([]byte(auth.AppHost + r.RequestURI))
		auth.RedirectLogin(w, r, OriginalURL)
		return false
	}

	// 验证有效性
	return true
}

// 模拟业务代码
func (Filter *AKAuthFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 调用前向过滤
	if !Filter.AKAuthFilter(w, r) {
		return
	}

	// 调用原始函数
	if Filter.Handler == nil {
		http.DefaultServeMux.ServeHTTP(w, r)
	} else {
		Filter.Handler.ServeHTTP(w, r)
	}
}

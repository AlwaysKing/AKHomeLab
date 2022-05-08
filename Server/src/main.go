package main

import (
	"log"
	"mime"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
}

var Filter = &AKAuthFilter{
	AppID:        "AKHomeLab",
	Token:        "A_NPQqa,aQ*Yp8b%%iC7",
	AppTokenName: "AKHomeLabAppToken",
	AuthServer:   "https://akauth.alwaysking.cn:16443",
	AuthGrantCB:  "/AKGrantAuth",
	AuthRevokeCB: "/AKRevokeAuth",
}

func main() {

	// 日志初始化
	FileLog := &AKFILELOGOUT{}
	FileLog.Init("./Log/", 30)
	Log.Out = FileLog
	LogOutPM = true
	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/", http.FileServer(http.Dir("./html")))
	err := http.ListenAndServeTLS("0.0.0.0:9886", "./Cert/alwaysking.pem", "./Cert/alwaysking.key", Filter)
	// err := http.ListenAndServe("0.0.0.0:9886", Filter)
	Log.Error(err)
	log.Fatal(err)
}

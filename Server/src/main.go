package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Name string
	Addr string
	File string
}

type Auth struct {
	code string
	time time.Time
	try  int8
}

var USERLIST map[string]*Config
var USERLIST_LOCKER = &sync.RWMutex{}

func cors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	if cors(w, r) {
		return
	}
	// 检查token
	token := r.Header.Get("token")
	vaild, refresh, _ := ParseToken(token)
	if vaild {
		// 没有token, 返回失败
		w.Write([]byte(fmt.Sprintf(`{"vaild":true,"refresh":%t}`, refresh)))
	} else {
		w.Write([]byte(`{"vaild":false}`))
	}
	return
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	if cors(w, r) {
		return
	}
	// 检查token
	token := r.Header.Get("token")
	vaild, _, name := ParseToken(token)
	if vaild {
		token, err := MakeToken(name)
		if err != nil {
			w.Write([]byte(`{"status":false}`))
			return
		}
		// 没有token, 返回失败
		w.Write([]byte(fmt.Sprintf(`{"status":true, "token":"%s"}`, token)))
	} else {
		w.Write([]byte(`{"status":false}`))
	}
	return
}

func SendAuth(w http.ResponseWriter, r *http.Request) {
	if cors(w, r) {
		return
	}
	namevars := r.URL.Query()["name"]
	if len(namevars) == 0 {
		w.Write([]byte(`{"status":false}`))
		return
	}

	name := namevars[0]
	USERLIST_LOCKER.Lock()
	user, ok := USERLIST[name]
	USERLIST_LOCKER.Unlock()
	if ok {
		http.Get(fmt.Sprintf("%s/AkHomeLab 登录验证码/%s", user.Addr, GenAuth(user.Name)))
		w.Write([]byte(`{"status":true}`))
		Log.Info(fmt.Sprintf("发送验证信息 %s", name))
		return
	}

	w.Write([]byte(`{"status":false}`))
	return
}

func GenToken(w http.ResponseWriter, r *http.Request) {
	if cors(w, r) {
		return
	}
	vars := r.URL.Query()
	namevars := vars["name"]
	authvars := vars["auth"]
	if len(namevars) == 0 || len(authvars) == 0 {
		w.Write([]byte(`{"status":false}`))
		return
	}

	if CheckAuth(namevars[0], authvars[0]) {
		token, err := MakeToken(namevars[0])
		if err != nil {
			w.Write([]byte(`{"status":false}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"status":true, "token":"%s"}`, token)))
		USERLIST_LOCKER.Lock()
		user, ok := USERLIST[namevars[0]]
		USERLIST_LOCKER.Unlock()
		if ok {
			http.Get(fmt.Sprintf("%s/AkHomeLab 登录验证码/成功登录", user.Addr))
		}
	} else {
		w.Write([]byte(`{"status":false}`))
	}

	return
}

func Data(w http.ResponseWriter, r *http.Request) {
	if cors(w, r) {
		return
	}
	// 检查token
	token := r.Header.Get("token")
	vaild, _, name := ParseToken(token)
	if !vaild {
		w.Write([]byte(`{"status":false}`))
		return
	}

	USERLIST_LOCKER.Lock()
	user, ok := USERLIST[name]
	USERLIST_LOCKER.Unlock()
	if !ok {
		w.Write([]byte(`{"status":false}`))
		return
	}

	data, err := ioutil.ReadFile(user.File)
	if err != nil {
		Log.Error("读取 config.json 文件失败")
		w.Write([]byte(`{"status":false}`))
		return
	}

	// 返回文件内容
	w.Write([]byte(data))
}

func main() {
	// 日志初始化
	rand.Seed(time.Now().UnixNano())
	FileLog := &AKFILELOGOUT{}
	FileLog.Init("./Log/", 30)
	Log.Out = FileLog
	LogOutPM = true

	// 加载配置
	USERLIST = make(map[string]*Config, 0)
	TMPAUTH = make(map[string]*Auth)

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		Log.Error("读取 config.json 文件失败")
		return
	}

	//读取的数据为json格式，需要进行解码
	tmp := make([]Config, 0)
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		Log.Error("解析 config.json 文件失败")
		return
	}

	for index := range tmp {
		USERLIST[tmp[index].Name] = &tmp[index]
	}

	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.HandleFunc("/checktoken", CheckToken)
	http.HandleFunc("/sendauth", SendAuth)
	http.HandleFunc("/gentoken", GenToken)
	http.HandleFunc("/refreshtoken", RefreshToken)
	http.HandleFunc("/data", Data)
	err = http.ListenAndServe("0.0.0.0:9886", nil)
	Log.Error(err)
	log.Fatal(err)
}

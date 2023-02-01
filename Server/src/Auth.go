package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var TMPAUTH map[string]*Auth
var TMPAUTH_LOCKER = &sync.RWMutex{}

func RandStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenAuth(name string) string {
	auth := &Auth{
		code: RandStr(4),
		time: time.Now(),
		try:  0,
	}

	TMPAUTH_LOCKER.Lock()
	TMPAUTH[name] = auth
	TMPAUTH_LOCKER.Unlock()

	return auth.code
}

func CheckAuth(name string, code string) bool {
	TMPAUTH_LOCKER.Lock()
	defer TMPAUTH_LOCKER.Unlock()

	item, ok := TMPAUTH[name]
	if !ok {
		Log.Error(fmt.Sprintf("check auth 无效用户名 %s", name))
		return false
	}

	if item.try > 5 {
		delete(TMPAUTH, name)
		Log.Error(fmt.Sprintf("check auth 超过5次 %s", name))
		return false
	}

	// 时间过期
	now := time.Now()
	if now.Sub(item.time).Seconds() > 300 {
		delete(TMPAUTH, name)
		Log.Error(fmt.Sprintf("check auth 过期 %s", name))
		return false
	}

	if item.code == code {
		delete(TMPAUTH, name)
		Log.Info(fmt.Sprintf("check auth 成功 %s", name))
		return true
	}

	item.try += 1

	return false
}

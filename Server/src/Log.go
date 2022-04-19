/*
	使用说明
	AKFILELOGOUT := &AKFILELOGOUT{}
	AKFILELOGOUT.Init("文件或路径")
	Log.Out = AKFILELOGOUT

	Log.Info("xxx","xxx")
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var LogOutPM = false

func PerformanceMonitoring(Info string, Start time.Time) {
	if LogOutPM {
		End := time.Now()
		Log.Info(Info, " 用时:", End.UnixMilli()-Start.UnixMilli(), " 毫秒")
	}
}

const (
	SVNT_LOG_LEVEL_DEBUG int = iota
	SVNT_LOG_LEVEL_INFO  int = iota
	SVNT_LOG_LEVEL_ERROR int = iota
)

type AKLOG struct {
	Out   AKLOGOUT
	Level int
}

type AKLOGOUT interface {
	OutPut(string)
}

var Log *AKLOG

func init() {
	// 默认打印到控制台
	Log = &AKLOG{Out: AKFMTLOGOUT{}, Level: SVNT_LOG_LEVEL_INFO}
}

func (Log *AKLOG) Debug(a ...interface{}) {
	if Log.Level > SVNT_LOG_LEVEL_DEBUG {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf("%-25s   %-30s   [DEBUG] %s\r\n", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprint(filepath.Base(file), ":", line), fmt.Sprint(a...))
	Log.Out.OutPut(info)
}

func (Log *AKLOG) Info(a ...interface{}) {
	if Log.Level > SVNT_LOG_LEVEL_INFO {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf("%-25s   %-30s   [Info]  %s\r\n", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprint(filepath.Base(file), ":", line), fmt.Sprint(a...))
	Log.Out.OutPut(info)
}

func (Log *AKLOG) Error(a ...interface{}) {
	if Log.Level > SVNT_LOG_LEVEL_ERROR {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf("%-25s   %-30s   [Error] %s\r\n", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprint(filepath.Base(file), ":", line), fmt.Sprint(a...))
	Log.Out.OutPut(info)
}

// 默认的输出到控制台的Log
type AKEMPTYLOGOUT struct {
}

func (FmtLog AKEMPTYLOGOUT) OutPut(info string) {
}

// 默认的输出到控制台的Log
type AKFMTLOGOUT struct {
}

func (FmtLog AKFMTLOGOUT) OutPut(info string) {
	fmt.Print(info)
}

// 默认的输出到控制台的Log
type AKFILELOGOUT struct {
	FilePath  string
	SavedDays time.Duration
	dir       bool
	pid       string
}

func (FileLog *AKFILELOGOUT) Init(FilePath string, SavedDays time.Duration) error {
	FileLog.FilePath = FilePath
	FileLog.SavedDays = SavedDays

	// 判断是文件夹还是文件
	if strings.HasSuffix(FilePath, `\`) || strings.HasSuffix(FilePath, `/`) {
		FileLog.dir = true
		FileLog.pid = strconv.Itoa(os.Getpid())
		go FileLog.MonitorFile()
	} else {
		FileLog.dir = false
	}

	return nil
}

func (FileLog *AKFILELOGOUT) OutPut(info string) {
	file, err := FileLog.GetFile()
	if err == nil {
		file.Write([]byte(info))
		file.Close()
	}
}

func (FileLog *AKFILELOGOUT) GetFile() (*os.File, error) {
	_File := FileLog.FilePath
	if FileLog.dir && os.MkdirAll(_File, os.ModePerm) == nil {
		// 确保文件夹存在
		_File += FileLog.FileName()
	}

	// 文件模式
	return os.OpenFile(_File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func (FileLog *AKFILELOGOUT) FileName() string {
	return time.Now().Format("20060102") + "." + FileLog.pid + ".log"
}

func (FileLog *AKFILELOGOUT) MonitorFile() {
	CleanCount := 0
	for {
		time.Sleep(time.Second)
		// 清理文件 1小时查一次
		if CleanCount > 3600 {
			CleanCount = 0
			FileLog.Clean()
		}
	}
}

// 保留5天的日志
func (FileLog *AKFILELOGOUT) Clean() {
	Now := time.Now().Add(-1 * 24 * time.Hour * FileLog.SavedDays)
	dir, err := ioutil.ReadDir(FileLog.FilePath)
	if err != nil {
		Log.Error("尝试清理日志文件夹失败")
		return
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".log")
			if ok {
				//files = append(files, dirPth+PthSep+fi.Name())
				if fi.ModTime().Before(Now) {
					os.Remove(FileLog.FilePath + "/" + fi.Name())
				}
			}
		}
	}
}

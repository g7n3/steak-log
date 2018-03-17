package log

import (
	"regexp"
	"log"
	"strings"
	"io"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	regex_o = `\[\d+\]: `
	regex_t = ` \b\w+_\w+\b`
	lg_ext = ".log"
)

func Handle (b string) {
	var sepContent []string

	r, _ := regexp.Compile(regex_o)
	delimiter := r.FindString(b)
	sepContent = strings.Split(b, delimiter)

	if len(sepContent) <= 0 {
		sepContent[0] = b
		sepContent[1] = b
	}

	r, _ = regexp.Compile(regex_t)
	fileName := strings.TrimSpace(r.FindString(sepContent[0]))
	// @todo 每次都要初始化哦，可能要斟酌一下
	prepareLog(fileName)
	printLog(sepContent[1])
}

// 准备输出
// 日志文件格式是 <服务>_<模块>_2010_02_03.log
// Log filename formula <service>_<module>_2010_02_03.log
func prepareLog (lp string) {
	if lp == "" {
		lp = "default"
	}
	tm := time.Now().Format("_2006_01_02")

	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   GetConfig("LogPath") + lp + tm + lg_ext,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}, os.Stdout))
}

func printLog (v ...interface{}) {
	log.Println(v)
}
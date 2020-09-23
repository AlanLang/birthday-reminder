package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nosixtools/solarlunar"
)

var sckey string

func init() {
	flag.StringVar(&sckey, "sckey", "", "")
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile("config")
	config := string(b)
	if err != nil {
		log.Println("无配置文件", err.Error())
		return
	}
	log.Println("今天是：" + getToday(false) + "，农历：" + getToday(true))
	today := getToday(false)
	todayLunar := getToday(true)

	for _, lineStr := range strings.Split(config, "\n") {
		if strings.HasPrefix(lineStr, "#") {
			continue
		}
		birthday := strings.Split(lineStr, " ")
		name := birthday[0]
		day := birthday[1]
		isLunar := birthday[2]
		if (isLunar == "n" && today == day) || (isLunar == "y" && todayLunar == day) {
			sendMessage("今天是" + name + "的生日，请别忘记祝他生日快乐～")
			log.Println("今天是" + name + "的生日，请别忘记祝他生日快乐～")
		}
	}
}

func getToday(isLunar bool) string {
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()
	if isLunar {
		lunarTiem := solarlunar.SolarToSimpleLuanr(strconv.Itoa(year) + "-" + preMonth(month) + "-" + strconv.Itoa(day))
		lunarMonth := string([]rune(lunarTiem)[5:7])
		lunarDay := string([]rune(lunarTiem)[8:10])
		return lunarMonth + lunarDay
	}
	return preMonth(month) + strconv.Itoa(day)
}

func preMonth(month int) string {
	var pre string = ""
	if month < 10 {
		pre = "0"
	}
	return pre + strconv.Itoa(month)
}

func sendMessage(message string) {
	if sckey == "" {
		log.Println("sckey不存在")
		return
	}
	http.Get("https://sc.ftqq.com/" + sckey + ".send?text=" + message)
}

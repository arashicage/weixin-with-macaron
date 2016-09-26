package biz

import (
	"fmt"
	"log"
	"strings"

	"github.com/garyburd/redigo/redis"
)

func CXFP(key string, kprq, kpje string) map[string]string {

	m, err := QueryRedis(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(m) == 0 {

		m["00"] = "查无此票"

	} else if m["10"] != kpje || m["01"] != strings.Replace(kprq, "-", "", -1) {

		m["00"] = "不一致"
		m["01"] = "-"
		m["10"] = "-"
		m["11"] = "-"
		m["12"] = "-"
		m["02"] = "-"
		m["03"] = "-"
		m["04"] = "-"
		m["05"] = "-"
		m["06"] = "-"
		m["07"] = "-"
		m["08"] = "-"
		m["09"] = "-"

	} else {

		m["00"] = "一致"

	}

	return m
}

func QueryRedis(s string) (m map[string]string, err error) {

	rs, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
	}
	defer rs.Close()

	v, err := redis.Values(rs.Do("HGETALL", s)) //s[1]+":"+s[2]+s[3]))
	if err != nil {
		panic(err)
	}

	m, err = redis.StringMap(v, err)

	for k, v := range m {
		log.Println(k, v)
	}

	return m, err

	//	xxx := CYJG{"1", m["06"], m["07"], m["02"], m["03"]}
	//	log.Println(xxx)

	//	output, _ := json.Marshal(&xxx)
	//	log.Println(string(output))
	//	log.Fprintln(res, string(output))

}

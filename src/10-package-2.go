package testlib

import (
	mongo "other/mongo/db"
	mysql "other/mysql/db"
	_ "other/xlib"
)

var pop map[string]string

func init() {

	/**
	init 함수는 패키지가 로드 될 때 별도의 호출 없이도 자동으로 호출되는 함수
	*/

	// 패키지 로드 할 때 map을 초기화 한다
	pop = make(map[string]string)

	/**
	import 할 때 init() 함수만 호출하고자 할 때, '_' 형태의 alias를 지정한다.
	*/

	mongodb := mongo.Get()
	myDb := mysql.Get()
}

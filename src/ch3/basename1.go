package main

import "strings"

// directory 구성 요소와 extends를 제거한다.

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// strings 패키지를 사용한 더 간단한 버전
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // "/"가 없으면 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

/**
path와 path/filepath는 위와 같은 계층적 이름을 조작하는 더 범용적인 함수가 있는데,
path는 슬래시로 구분된 경로에 대해 동작한다. URL의 경로와 같은 다른 도메인에 적합하다.
path/filepath는 POSIX에서는 /foo/bar, 윈도우에서는 c:\foo\bar와 같은 호스트 플랫폼의 규칙에 맞게 동작한다.
*/

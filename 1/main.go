package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-collections/collections/stack"
)

func main() {
	//괄호가 올바르게 쓰였는지 확인하는 로직을 만들려고 합니다.
	//괄호는 대괄호, 중괄호, 소괄호 세가지가 있으며 주어지는 스트링이
	//괄호가 올바르게 사용되었으며 true, 아니면 false를 리턴하는 함수를 작성하세요.
	kbReader := bufio.NewReader(os.Stdin)
	fmt.Print("테스트할 문자열을 입력하세요 : ")
	strtmp, err := kbReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("입력값 : %s\n", strtmp)

	start := time.Now()

	//스택을 사용. 스택은 라이브러리 있는거 사용.
	var stack stack.Stack
	//결과값
	var result bool
	//시작 괄호 체크용 배열 생성.
	bracketStartArr := [3]string{"(", "{", "["}
	//끝 괄호 체크용 배열 생성
	bracketEndArr := [3]string{")", "}", "]"}
	//끝 괄호를 키로 시작괄호와 관계 생성
	bracketMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	slice := strings.Split(strtmp, "")

	hasStrangeCase := false
	for _, str := range slice {
		if strContains(str, bracketStartArr[:]) {
			stack.Push(str)
		} else if strContains(str, bracketEndArr[:]) {
			if stack.Len() == 0 {
				hasStrangeCase = true
				break
			} else if stack.Peek().(string) != bracketMap[str] {
				hasStrangeCase = true
				break
			} else {
				stack.Pop()
			}
		}
	}
	if !hasStrangeCase && stack.Len() == 0 {
		result = true
	}

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("result : ", result)
	fmt.Println("time  : ", elapsed)
}

func strContains(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

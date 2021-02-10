package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 너구리 한 쌍은 한 달 후에 다른 새끼 너구리 한 쌍을 낳습니다.
	// 이 새끼 너구리 한 쌍은 한 달 동안 성체가 되며 성체가 된 너구리 한 쌍은 다시 한 달 후에 다른 새끼 너구리 한 쌍을 낳습니다.
	// 이미 성체가 된 너구리 부부는 달마다 새끼 너구리를 한 쌍씩 낳는다고 가정할 때, n달 후의 너구리 수를 구하는 함수를 작성하세요. (단, 이 너구리들은 죽지 않습니다.)

	//풀이법 대략 3개 정도
	//1. 재귀함수로 풀기 > 시간복잡도 높음, 너무 흔해서 하기 싫음(?)

	//2. 루프로 풀기 > 월, 성채너구리 수, 새끼너구리 수 를 받을 변수들을 만들고 루프를 돌리면 풀릴 것 같다.

	//3. 피보나치 수열에서 온 문제같으니 피보나치 스타일로 풀기.
	// 3번으로만 풀어볼 예정.
	// 달 / 성체 / 새끼 / 총마릿수 순으로 나열해봄 (0달부터 6달까지)
	// 0 2 0 2
	// 1 2 2 4
	// 2 4 2 6
	// 3 6 4 10
	// 4 10 6 16
	// 5 16 10 26
	// 6 26 16 42
	// 보면 새끼수가 피보나치 수열*2의 꼴로 증가하고 있음. 또한 n달에 대한 총마릿수는,
	// n달새끼구하기 함수가 있다면 > n달새끼구하기(n)+n달새끼구하기(n+1) = n달 총마릿수
	//(수식으로 증명 안해도 논리적으로 n달에 성채수만큼 n+1달에 새끼가 생기기 때문에 둘을 더하면 n달 총마릿수가 됨이 보임.)

	//피보나치 수열의 일반항이 있던거 같은데?? 싶어서 구글링해서 찾음.
	// (1/sqrt5) * (pow((1+sqrt5)/2, n) - pow((1-sqrt5)/2, n)); 요런 양식이고, 너구리는 해당 식의 두배로 늘음. 식을 고 문법으로 표현하면 됨.

	kbReader := bufio.NewReader(os.Stdin)
	fmt.Print("n달 후의 너구리 수를 구합니다. n이될 값을 입력하세요 : ")
	strtmp, err := kbReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("입력값 : %s\n", strtmp)
	strtmp = strings.Replace(strtmp, "\r", "", -1)
	strtmp = strings.Replace(strtmp, "\n", "", -1)

	i, err := strconv.Atoi(strtmp)
	if err != nil {
		log.Fatal(err)
	}
	fiboN := fibo(float64(i)) * 2
	fiboNp1 := fibo(float64(i+1)) * 2

	//위에 두 피보나치 일반항*2값을 더한게 n달차 너구리 총합. (부동소수점 문제 안생기게 반올림 처리)
	result := fmt.Sprint(math.Round(fiboN + fiboNp1))
	fmt.Println("result >> ", result)
}

func fibo(n float64) float64 {
	return (1 / math.Sqrt(5)) * (math.Pow((1+math.Sqrt(5))/2, n) - math.Pow((1-math.Sqrt(5))/2, n))
}

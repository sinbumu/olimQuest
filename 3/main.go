package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int
var alreadyVisitCheck []bool

func main() {
	//오프라인 매장이 문을 닫아서 상담원들은 직접 고객을 만나러 각지를 방문해야 하는 상황이 되었습니다.
	// 이들이 최소한의 이동시간으로 고객을 만날 수 있도록 최소 이동시간을 구하는 함수를 작성하세요.
	// 입력은 고객 수만큼의 방문지 리스트가 주어지며, 각 방문지에서는 다른 방문지로 이동할 시에 소요되는 이동시간이 다시 리스트로 주어집니다.
	// (단, 방문지의 수는 10명 미만으로 주어집니다)

	// 이 문제를 풀 수 있는 효율좋은 거리구하기 알고리즘도 있을거 같지만
	// 방문지의 수를 10명 미만으로 제한을 걸은 부분 때문에 전체를 계산하는 케이스라도 10! 번 간단한 연산을 해주면 된다.
	// 그래서 깊이 우선 탐색에 조건만 약간 거는 정도로 풀어도 될것 같다.

	//예 : [ [0, 611, 648], [611, 0, 743], [648, 743, 0] ]
	//[ [0, 326, 503, 290], [326, 0, 225, 395], [503, 225, 0, 620], [290, 395, 620, 0] ]

	// testArr := [][]int{
	// 	{0, 611, 648},
	// 	{611, 0, 743},
	// 	{648, 743, 0},
	// }
	// testArr := [][]int{
	// 	{0, 326, 503, 290},
	// 	{326, 0, 225, 395},
	// 	{503, 225, 0, 620},
	// 	{290, 395, 620, 0},
	// }
	// testArr := [][]int{
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	// }
	testArr := forHeavyTest()
	fmt.Println(fmt.Sprintln(testArr))

	//시간체크
	start := time.Now()

	length := len(testArr) //방문할 고객 수
	alreadyVisitCheck = make([]bool, length)
	depth := 0 //현제 깊이
	score := 0 //점수

	dfsStart(depth, score, length, -1, testArr)
	fmt.Println("result >> ", result)

	fmt.Println("time  : ", time.Since(start))
}

func dfsStart(depth int, score int, length int, currentIndex int, testArr [][]int) {
	//최초 시작으로 아직 방문예정인 고객이 정해지지 않음.
	for i := 0; i < length; i++ {
		// clientItem := testArr[i]//i 번째 고객으로 최초 방문 정해짐. ㄴㄴ 첫 지점은 점수 없음.
		for j := 0; j < length; j++ {
			if i != j { //최초 시작이기 때문에 이미 방문했나는 따지지 않음.
				alreadyVisitCheck[j] = true
				dfs(depth+1, score, length, j, testArr) //댑스는 하나 올라가며, 스코어는 0 그대로, currentIndex는 이제부턴 전달됨.
				alreadyVisitCheck[j] = false
			}
		}
	}
}

func dfs(depth int, score int, length int, currentIndex int, testArr [][]int) {
	//본격적인 깊이탐색 반복.
	if result != 0 && score > result { //result가 0이 아닌데 score가 result보다 크면 해당 심화탐색 종료.
		return
	}
	if depth == length { //이제 모든 방문이 이뤄졌기 때문에 최종 결산
		if result == 0 || result > score {
			result = score
			return
		}
	}
	clientItem := testArr[currentIndex] //이번 함수에서 확인할 배열
	for i := 0; i < length; i++ {
		if i != currentIndex && !alreadyVisitCheck[i] {
			addScore := clientItem[i]
			alreadyVisitCheck[i] = true
			dfs(depth+1, score+addScore, length, i, testArr)
			alreadyVisitCheck[i] = false
		}
	}
}

func forHeavyTest() [][]int {
	//최고로 무거운 경우(10명의 고객방문.)에도 잘 작동하는지 봐야 겠어서 무작위 10명분 배열 만드는 함수.
	rand.Seed(time.Now().Unix())

	var returnV [10][]int

	for i := 0; i < 10; i++ {
		returnV[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			if i != j {
				returnV[i][j] = rand.Intn(1000) + 1
			} else {
				returnV[i][j] = 0
			}
		}
	}
	return returnV[:]
}

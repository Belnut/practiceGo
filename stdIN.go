package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	var f *os.File
	f = os.Stdin		// f를 fileStream변수를 만들고 Stdin 연결
	defer f.Close()	// 함수 종료시 실행되는 문장

	scanner := bufio.NewScanner(f)	//io 버퍼의 Scanner 생성후 f 를 연결

	for scanner.Scan() {
		fmt.Println( ">", scanner.Text() )	//stdin 으로 입력된 값을 Scanner를 통해서 console에 출력함
	}
}


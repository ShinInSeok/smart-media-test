package main

import (
	"shininseok/test2557"
)

func main() {

	test2557.Test2557()
	// 같은 형식의 출력
	//********************************************************************
	//1.
	// var text string
	// text = "ㅋㅋㅋㅋㅋ"
	// fmt.Println(text)
	// 2.
	// text := "ㅋㅋㅋㅋㅋ"
	// fmt.Println(text)

	// 조건문 (if)
	//********************************************************************
	// a := "text1"
	// b := "text2"

	// if a == b {
	// 	fmt.Println("Is Equals!")
	// } else {
	// 	fmt.Println("Is Not Equals!")
	// }

	// 조건문 (case)
	//********************************************************************
	// a := "text"
	// switch a {
	// case "text":
	// 	fmt.Println("1")
	// case "text2":
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }

	// 함수 사용
	//********************************************************************
	// v1 := 10
	// v2 := 20

	// v3 := Add(v1, v2)

	// fmt.Println(v3)

	// func Add(v1 int, v2 int) int {
	// 	return v1 + v2
	// }

	// 예외 처리
	//********************************************************************
	// v1 := 10
	// v2 := 20

	// v3, err := Add(v1, v2)
	// if err != nil {
	// 	fmt.Println("err:", err.Error())
	// 	return
	// }
	// fmt.Println(v3)
	// func Add(v1 int, v2 int) (int, error) {
	// 	result := v1 + v2
	// 	if result >= 30 {
	// 		return 0, errors.New("결과값이 30보다 같거나 큼")
	// 	} else {
	// 		return result, nil
	// 	}
	//}

	// 구조체 , 구조체 함수
	//********************************************************************

	// var sex Sex
	// sex = Sex{
	// 	mail:   "남성",
	// 	femail: "여성",
	// }
	// fmt.Println(sex.mails())
	// fmt.Println(sex)

	// type Sex struct {
	// 	mail   string
	// 	femail string
	// }

	// func (ss Sex) mails() string {
	// 	return ss.mail + "남자야!"
	// }

	// 반복문
	//********************************************************************

	//1.for 형태
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//2. while 형태
	// sum := 0
	// i := 0
	// for true {
	// 	sum += i
	// 	if i >= 100 {
	// 		break
	// 	}
	// 	i++
	// }
	// fmt.Println(sum)

	//3. foreach 형태
	// names := []string{"신인석", "김건국", "윤세영"}

	// for i, v := range names {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }
}

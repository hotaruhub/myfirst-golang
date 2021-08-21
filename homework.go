import "fmt"

type Turtle struct {
	name string
	x    float64
	y    float64
	a    float64
}

func (t *Turtle) addA(num float64) {
	t.a += num
}

func (t *Turtle) changeName(newName string) {
	t.name = newName
}

func (t *Turtle) printMessage(message string) {
	stdout := t.name + ": " + message
	fmt.Println(stdout)
}

func subMain() {
	num1 := 123.456
	num2 := 20.964

	// task 2-1_1(引き算)
	var subtractionResult float64 = minusNums(num1, num2)
	fmt.Println(subtractionResult)

	// task 2-1_2(掛け算)
	var multiplicationResult = multiplicateNums(num1, num2)
	fmt.Println(sumResult)

	// task 2-1_3(割り算)
	division := divideNums(num1, num2)
	fmt.Println(division)

	// task 2-2
	// 指定された符号にそった計算の答えを返却
	calcResult := calculation("/", num1, num2)
	fmt.Println(calcResult)

	// Turtle を宣言
	var kameko Turtle = Turtle{"亀子", 200, 400, 910}
	// task 2-3_1
	// a の値を引数をもとに足し算
	kameko.addA(3102)

	// task 2-3_2
	// nameの値を引数をもとに変更
	kameko.changeNamt("亀山亀実")

	// task2-3_3
	// nameと引数の値からメッセージを出力する
	kameko.printMessage("いってきます")
}

// task 2-1_1(引き算)関数
func minusNums(num1 float64, num2 float64) float64 {
	return (num1 - num2)
}

// task 2-1_2(掛け算)関数
func multiplicateNums(num1 float64, num2 float64) float64 {
	return (num1 + num2)
}

// task 2-1_3(割り算)関数
func divideNums(num1 float64, num2 float64) float64 {
	return (num1 / num2)
}

// task 2-2
// 指定された符号により計算方式を変更
func calculation(sign string, num1 float64, num2 float64) float64 {
	switch sign {
	case "-":
		return minusNums(num1, num2)
	case "*":
		return multiplicateNums(num1, num2)
	case "/":
		return divideNums(num1, num2)
	case "+":
		return num1 + num2
	default:
		fmt.Println("指定された引数で計算できませんでした", sign, num1, num2)
	}
}

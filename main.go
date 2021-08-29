package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// struct
// dbのカラム名指定と、jsonのkeyを指定します
type Turtle struct {
	Id   int     `json:"id" db:"id"`
	Name string  `json:"chome" db:"name"`
	X    float64 `json:"x" db:"x"`
	Y    float64 `json:"y" db:"y"`
	A    float64 `json:"a" db:"a"`
}

type Turtles []Turtle

func main() {
	// connect database (driver, user:passwd@tcp(dst:port)/database)
	db, err := sqlx.Open("mysql", "root:1234@tcp(localhost:3306)/game")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e := echo.New()
	// t := Turtle{"alice", 10, 200, 459831}
	e.GET("/", func(c echo.Context) error {
		var kames Turtles

		// データの取得 & データの加工
		// 宣言は2つ必要。（1つだとエラーになる）
		rows, err := db.Queryx("SELECT * FROM turtles")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var kame Turtle

		for rows.Next() {
			err := rows.StructScan(&kame) // "&" is pointer
			if err != nil {
				fmt.Println(err)
			}
			kames = append(kames, kame)
		}

		fmt.Println(kames)

		// データベースから取得したデータを加工
		if err := c.Bind(kames); err != nil {
			return err
		}
		// message := "TOP PAGE"
		// return c.String(http.StatusOK, message)
		return c.JSON(http.StatusOK, kames)
	})
	e.GET("/about", func(c echo.Context) error {
		// データの取得 & データの加工
		// データベースから取得したデータを加工
		message := c.QueryParam("ms")
		return c.String(http.StatusOK, message)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// package main

// import (
// 	"fmt"
// )

// // struct
// type Turtle struct {
// 	name string
// 	x    float64
// 	y    float64
// 	a    float64
// }

// func (t *Turtle) move(x float64, y float64) {
// 	t.x += x
// 	t.y += y
// }

// func main() {
// 	var t1 Turtle = Turtle{"師匠", 1000, 5, 180.5}
// 	var t2 Turtle = Turtle{"弟子", 10, 250, 270.3}
// 	var t3 Turtle = Turtle{"兄弟子", 750, 50, -2}
// 	fmt.Println(t1)
// 	fmt.Println(t2)
// 	fmt.Println(t3)

// 	t1.move(10, -10)
// 	fmt.Println("after move", t1)

// 	// declare

// 	// string
// 	var message string = "i am string"
// 	var message2 = "I'm string"
// 	message3 := "me"
// 	fmt.Println(message, message2, message3)

// 	// numerics
// 	num1 := 123.44444445555556252435
// 	num2 := 34.1
// 	fmt.Println("sum", num1+num2)
// 	fmt.Println("%", num1/num2)

// 	// constant
// 	const SELECT_LANGUAGE = "ja"
// 	fmt.Println("your select lang is..", SELECT_LANGUAGE)

// 	// array
// 	var nums [3]int = [3]int{}
// 	nums[0] = 103
// 	nums[1] = 109
// 	nums[2] = 17
// 	fmt.Println(nums)
// 	fmt.Println(len(nums))
// 	// slice
// 	var numsSlice []int = []int{}
// 	numsSlice = append(numsSlice, 99)
// 	numsSlice = append(numsSlice, 111)
// 	numsSlice = append(numsSlice, 1234567890123456789)
// 	fmt.Println(numsSlice)

// 	// map
// 	status := map[string]int{
// 		"age":      888,
// 		"height":   198,
// 		"weight":   79,
// 		"eyesight": 2,
// 		"savings":  3772392,
// 	}
// 	status["friends"] = 86
// 	delete(status, "savings")
// 	fmt.Println(status)

// 	switch nums[0] {
// 	case 102:
// 		fmt.Println("case:102")
// 	case 103:
// 		fmt.Println("case:[103],fallthrough")
// 		fallthrough
// 	case 104:
// 		fmt.Println("case:104")
// 	case 105:
// 		fmt.Println("case:105")
// 	default:
// 		fmt.Println("default")
// 	}

// 	// 配列をfor文で回す
// 	for i, num := range nums {
// 		fmt.Println("-----")
// 		if i == 1 {
// 			continue
// 		}
// 		fmt.Println("index", i)
// 		fmt.Println("value", num)
// 	}

// 	// use func
// 	fmt.Println(hot())
// }

// func hot() string {
// 	hotja := "a tsh i yo"
// 	return hotja
// }

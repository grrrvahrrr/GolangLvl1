//Lesson 6 Homework

/*
Задание 1
Хотелось что-то сделать со страктами, методами и еще и поинтер использовать, поэтому переписал задание второго урока,
включив туда сразу все. Работать оно должно в разы эффективнее, чем мой изначальный код, где я принимал число, переводил его в строку,
а потом раскладывал строку на символы.
Для задания третьего урока так же можно использовать стракт для операций с двумя числами и операций с одним числом, а потом
написть методы с поинтерами на каждую операцию
Для задний четвертого и пятого уроков(сортировка вставками и фибоначчи) мне кажется слайс подходит идеально, так как мы можем добавлять в него
значения до бесконечности, а потом выводить их - все в случае с сортировкой, одно последние в случае с фибоначчи.
Для оптимизации в сортировке я думаю возможно создать создать слайс определенной длины с нулевыми значениями(длину мы знаем),
потом используя поинетр на этот слайс присвоить значениям цифры в цикле. Таким образом можно избежать аппенда в цикле,
который каждый раз перезаписывает слайс(аррэй).
*/
/*
Задание 2

package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	c := 2
	d := &c
	e := &d

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)
	fmt.Println(*d)
	fmt.Println(a*c)
	fmt.Println(*b**d)
	fmt.Println(a**d)
	fmt.Println(*b***e)
	fmt.Println(a***e)
}

Немного поиграл в поинтерами и со знаком *. Не читал подробно, но думаю,
что компилятор перед интерпритацией умножения проверяет есть ли 2 значения, которые можно умножить.
Но знаков * мы можем использовать довольно много, так как возможно иметь поинетер на поинтер на поинтер и т.д.
Программа написанная выше работает без ошибки и выдает результаты всех умножений.
Если пойти глубже то наверное можно и fmt.Println(***i****j) написать и это тоже умножиться, но это уже почти не читаемо.
*/
/*
package main

import (
	"fmt"
	"os"
)

type Digits struct {
	h, t, s string
}

func (d *Digits) PrintNums() {
	fmt.Println(`Its hundreds:`, d.h)
	fmt.Println(`Its tens:`, d.t)
	fmt.Println(`Its singles:`, d.s)
}

func main() {

	fmt.Println(`Input a three digit number, please`)

	var number string
	fmt.Scan(&number)

	if len(number) != 3 {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	numStruct := Digits{string(number[0]), string(number[1]), string(number[2])}

	numStruct.PrintNums()
}
*/

package main

import (
	"fmt"
	"math"
	"os"
)

const (
	Addition       = "+"     //adding numbers
	Subtraction    = "-"     //subtracting numbers
	Multiplication = "*"     //multiplying numbers
	Division       = "/"     //divising numbers
	Squared        = "sqr"   //squared numbers
	SqrRoot        = "root"  //square root
	Logarithm      = "log"   //Logarithm of a number
	Prime          = "prime" //Get prime numbers upto this point
)

//Type Declaration
type TwoNums struct {
	num1, num2 float64
}
type OneNum struct {
	num1 float64
}

//Two Number Functions
func (nums *TwoNums) ScanNums() (err error) {
	fmt.Println(`Input two numbers, please (divided by space)`)
	_, err = fmt.Scanf("%f %f", &nums.num1, &nums.num2)
	return
}

func (nums *TwoNums) Addition() float64 {
	return nums.num1 + nums.num2
}

func (nums *TwoNums) Subtraction() float64 {
	return nums.num1 - nums.num2
}

func (nums *TwoNums) Multiplication() float64 {
	return nums.num1 * nums.num2
}

func (nums *TwoNums) Division() float64 {
	return nums.num1 / nums.num2
}

//One number functions
func (num *OneNum) ScanNum() (err error) {
	fmt.Println(`Input a number`)
	_, err = fmt.Scanf("%f", &num.num1)
	return
}

func (num *OneNum) Squared() float64 {
	return num.num1 * num.num1
}

func (num *OneNum) Sqrt() float64 {
	return math.Sqrt(num.num1)
}

func (num *OneNum) Log() float64 {
	return math.Log(num.num1)
}

func main() {
	var operation string

	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "/" for Division, "sqr" for Squared, "root" for Square root, "log" for Logarithm, "prime" to get prime numbers up to input`)

	_, err := fmt.Scan(&operation)
	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	switch operation {
	case Addition:
		twoNumsStruct := TwoNums{}
		err = twoNumsStruct.ScanNums()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", twoNumsStruct.Addition())

	case Subtraction:
		twoNumsStruct := TwoNums{}
		err = twoNumsStruct.ScanNums()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", twoNumsStruct.Subtraction())

	case Multiplication:
		twoNumsStruct := TwoNums{}
		err = twoNumsStruct.ScanNums()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", twoNumsStruct.Multiplication())

	case Division:
		twoNumsStruct := TwoNums{}
		err = twoNumsStruct.ScanNums()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", twoNumsStruct.Division())

	case Squared:
		oneNumStruct := OneNum{}
		err = oneNumStruct.ScanNum()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", oneNumStruct.Squared())

	case SqrRoot:
		oneNumStruct := OneNum{}
		err = oneNumStruct.ScanNum()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", oneNumStruct.Sqrt())

	case Logarithm:
		oneNumStruct := OneNum{}
		err = oneNumStruct.ScanNum()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f\n", oneNumStruct.Log())

	default:
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}
}

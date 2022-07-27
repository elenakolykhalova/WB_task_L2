package main

import (
	"fmt"
	"strconv"
	"unicode"
	"errors"
)

var errCustome = errors.New("incorrect input")

func UnpakingStr(str string) (string, error) {
	
	rStr := []rune(str) //Создаем слайс рун из полученной строки
	res := make([]rune, 0) // Создаем слайс для записи результата
	flag := false //флаг для escape-последовательностей

	if len(rStr) == 0 { //Если подадут пустую строку
		return "", nil
	}
	for i := 0; i < len(rStr);{ //в цикле начинаем перебирать строку, пока не дойдем до последнего символа
		if flag || unicode.IsLetter(rStr[i]) { //проверяем на существование флага или является ли сивол буквой
			if i < len(rStr) - 1 && unicode.IsDigit(rStr[i + 1]) { // если это буква и она не последняя и следующий элемпнт слайса цифра то 
				n, _ := strconv.Atoi(string(rStr[i + 1])) // берем этот элемент в строке и делаем ее цифрой
				for k := 0; k < n; k++ { // далее добавляем в итоговый слайс эту букву столько раз, чему равна эта цифра
					res = append(res, rStr[i])
				}
				i++ // передвигаемся далее по слайсу рун
			} else {
				res = append(res, rStr[i]) // если после буквы нет цифры, то просто добавляем букву
			}
			flag = false //обнуляем значение флага на false
		} else if !flag { //если символ не является буквой и флаг имеет значение false 
			if string(rStr[i]) == "\\" { //если символ является escape-последовательностью, делаем флаг true для того, чтобы напечатать символ после него 
				flag = true

			} else { //иначе возвращаем пустую строку и записывваем в поток ошибок, что ввод не корректен
				return "", errCustome
			}
		}
		i++
	}
	return string(res), nil //возврат результирующего слайса и ошибки в виде nil
}

func main() {
	res1, _ := UnpakingStr("")
	res2, _ := UnpakingStr("a4bc2d5e")
	res3, _ := UnpakingStr("abcd")
	res4, err := UnpakingStr("45")
	res5, _ := UnpakingStr("qwe\\4\\5")
	res6, _ := UnpakingStr("qwe\\45")
	res7, _ := UnpakingStr("qwe\\\\5")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4, err)
	fmt.Println(res5)
	fmt.Println(res6)
	fmt.Println(res7)
}
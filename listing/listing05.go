package listing

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func Test() *customError {
	{
		// do something
	}
	return nil
}

/*
Функция test() фактически не возвращает nil, а возвращает значение типа *customError.
В данном случае, ошибка не создается явно, но nil указатель на *customError присваивается переменной err.
Из-за этого условие if err != nil выполняется, и программа выводит "error".
*/
func main() {
	var err error
	err = Test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

//Вывод программы: error

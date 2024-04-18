## Функции
`func getFirst() int {`
`    return 1`
`}`

`func log(message string) {}`

`func add(a int, b int) int {}`
При одинаковом типе можно юзать короткий синтаксис
`func add(a, b int) int {}`


`func power(name string) (int, bool) {}`
`value, exists := power("text")`
если одно из возвращаемых значений не нужно, его присваивают переменной "_"
`_, exists := power("text")`

"_" - пустой идентификатор, можно использовать много раз с любыми типами


## Условия
`if exists == false {}`
`if (exists == false) {}`

## Циклы
есть только `for`
`arr := []string{"A", "B", "C"}`
- стандартный
`for i :=0; i< len(arr); i++ {`
`   val = arr[i]`
`}`
- через `range` перебирается массив (аналог **foreach**)
`for index, letter := range arr {`
`   ...`
`}`
- через состояние (аналог **while**)
`for contition {`
`   ...`
`}`
- бесконечный
`for {`
`   ...`
`}`

выход из цикла - `break` (так же при помощи `return`)

**Важно!** При итерации map порядок НЕ гарантируется! Слайс - гарантирует.
`m := map[string]int{`
`    "a": 1,`
`    "b": 2,`
`    "c": 3,`
`    "d": 4,`
`}`
`for k, v := range m {`
`   fmt.Printf("%s/%d\n", k, v)`
`}`
Результат МОЖЕТ быть таким!
`d/4`, `a/1`, `b/2`, `c/3`
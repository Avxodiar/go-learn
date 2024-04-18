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
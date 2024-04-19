# Структуры

В GO нет объектов, нет наследования и многих других понятий, свойственных ОО-языкам, полиморфизма или перегрузки.
В Go есть структуры, которые могут быть связаны с методами.

## Объявление
*без запятых!*
`type User struct {`
`  Race string`
`  Age int`
`}`

## Инициализация
*с запятыми!*
`Frodo := User {`
`  Race: "Hobbit",`
`  Age: 130,`
`}`

все свойства не обязательные!
`gollum := User{}`

`Elrond := User{Race: "Elf"}`
`Elrond.Age = 3000`

Неинициализированные свойства == `nil`

Можно указывать значения по порядку их следования
`Gandalf := User{"Mayar", 9000}`

## Передача структуры как параметр функции
`upAge(Gandalf)`

`func upAge(s User) {`
`  s.Age += 10000`
`}`

однако при таком использовании все изменения будут локальными - в копии структуры
для изменения оригинала необходимо использовать

## Указатели
`&` - оператор получения адреса

`Gandalf := &User{"Mayar", 9000}`

`*` - использование указателя при передаче в ф-цию, указатель на значение типа
`func upAge(s *User) {`
`  s.Age += 10000`
`}`

## Функции в структурах (аля методы класса)
`type User struct {`
`  Name string`
`  Age int`
`}`
`func (s *User) upAge() {`
`  s.Age += 1`
`}`

вызов
`Gandalf := &User{"Mayar", 9000}`
`Gandalf.upAge()`

## Конструкторы
Структуры не имеют конструкторов. Вместо этого, вы создаёте функцию, которая возвращает экземпляр нужного типа аля фабрика.
`func NewUser(race string, age int) *User {`
`  return &User{`
`    Race: race,`
`    Age: age,`
`  }`
`}`

Но это неверный путь? Фабрика не должна возвращать указатель.
`func NewUser(race string, age int) User {`
`  return User {`
`    Race: race,`
`    Age: age,`
`  }`
`}`

Несмотря на отсутствие конструкторов, есть встроенная функция `new`

## Выделение памяти для типа данных

goku := new(Saiyan)
`Gandalf := new(User)`
то же самое
`Gandalf := &User{}`

Свойсва в этом случае инициализируются так
`goku := new(Saiyan)`
`goku.name = "goku"`
`goku.power = 9001`

## Поля структур

`type User struct {`
`  Race string`
`  Age int`
`  Father *User`
`}`

инициализация
`Gandalf := &User{`
`  Race: "Mayar",`
`  Age: 9000,`
`  Father: &User {`
`    Race: "God",`
`    Age: 9001,`
`    Father: nil,`
`  },`
`}`

## Композиция (аля трейты)

`type Person struct {`
`  Name string`
`}`
` `
`func (p *Person) hello() {`
`  fmt.Printf("Hi, I'm %s\n", p.Name)`
`}`
` `
`type User struct {`
`  *Person`
`  Age int`
`}`

использование
`Gandalf := &User{`
`  Person: &Person{"Gray"},`
`  Age: 9001,`
`}`
`Gandalf.hello()`

Если композиция использована без имени (как в примере), то получаем косвенный доступ, но компилятор автоматически задаст ему имя равное наименованию композиции
`fmt.Println(Gandalf.Person.Name)`
что равнозначно
`fmt.Println(Gandalf.Name)`

Композитная функция всегда доступна через s.Person.hello()
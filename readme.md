Устные вопросы
1. Какой самый эффективный способ конкатенации строк?
Strings.Builder{}
2. Что такое интерфейсы, как они применяются в Go?
Это абстракция, которая содержит в себе методы, которые должны быть у структуры, чтобы ее можно было использовать
в качестве возвращаемого значения или принимаего аргумента функции.
3. Чем отличаются RWMutex от Mutex?
мьютекс блокирует запись для всех возможных операций, а рв блокирует отдельно запись и чтение. Соответственно если заблокирована запись, то чтение и паралельная запись не возможны. А если заблокировано чтение, то запись не возмножна, зато возможно паралельно другое чтение.
4. Чем отличаются буферизированные и не буферизированные каналы?
Небуферизированный канал всегда блокирует горутину, до чтения\записи
Буферизированный канал не всегда блокирует горутину.
В буферизированный канал можно писать до конца буфера, без блокировки.
А блокировка будет вызвана только записью сверх буфера.
5. Какой размер у структуры struct{}{}?
0 byte.
6. Есть ли в Go перегрузка методов или операторов?
У методов и операторов нет перегрузки.
7. В какой последовательности будут выведены элементы map[int]int?
Пример:
m[0]=1
m[1]=124
m[2]=281
в случайном порядке, если выводить через for range
по порядку ключей если через fmt.print
8. В чем разница make и new?
make выделяет память и возвращает саму сущность создаваемую make'ом
и работает только с картой, срезом и каналами, хотя все эти сущности и так Являются указателями
new выделяет память и возвращает указатель на сущность.
в случае со срезом и картой len и cap будут 0
9. Сколько существует способов задать переменную типа slice или map?
два: make и инициировать в ручную
10. Что выведет данная программа и почему?
func update(p *int) { // копирование адреса
b := 2
p = &b // здесь поменяется значение адреса который мы скопировали, а не того, который мы передали
}
func main() {
var (
a = 1
p = &a
)
fmt.Println(*p)
update(p)
fmt.Println(*p)
}
1
1
11. Что выведет данная программа и почему?
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) { // ошибка sync.WaitGroup надо передавать по указателю (*)
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}
1
надо передавать указатель на wg
12. Что выведет данная программа и почему?
func main() {
n := 0
if true {
n := 1 // лишние :
n++
}
fmt.Println(n)
}
0 повторная инициализация в {}
13. Что выведет данная программа и почему?
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b) вот тут мы меняем базовый массив на который ссылается слайс, поэтому все изменения далее и включительно этого аппенда не затронут слайс
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
100 2 3 4 5
14. Что выведет данная программа и почему?
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")  аналогично предыдущему примеру, тут мы меняем базовый массив среза, и работаем с другим массивом, который никак не касается нашего слайса
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
[b b a]
[a a]

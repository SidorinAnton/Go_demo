# Ветвление

## Операторы сравнения
- `>` -- больше
- `<`-- меньше
- `>=` -- больше или равно
- `<=` -- меньше или равно
- `==` -- равно
- `!=` -- не равно

## Логические операторы
- `&&` -- логическое И
- `||` -- логическое ИЛИ
- `!` -- логическое НЕ

## if-else
```go
if a == 1 {
    // сценарий, если условие if выполнено
} else if a == 2 {
    // сценарий, если условие else if выполнено
} else {
    // сценарий, если условие else if не выполнено
}
```


## switch-case
```go
var a int

switch a {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
case 3, 4:
    fmt.Println("3 or 4")
default:
    fmt.Println("Default case")
}
```

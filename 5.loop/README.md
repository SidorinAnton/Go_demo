# Циклы

Для всех циклов в Go используется `for`

## Бесконечный цикл

```{go}
for {
    // что-то
}
```

## Трёхкомпонентный цикл

```{go}
for i := 1; i < 10; i++ {
    // что-то
}
```

## Цикл while

```{go}
for i < 5 {
    // что-то
}
```

## Цикл range

```{go}
array := [3]int{1, 2, 3}

for arrayIndex, arrayValue := range array {
    fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
}
```

## Есть break и continue
## Есть метки и goto
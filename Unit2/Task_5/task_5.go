package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var N, T int
	// 1. Считайте два числа: N (количество чисел) и T (таймаут в миллисекундах)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	input := scanner.Text()
	strValues := strings.Fields(input)
	values := make([]int, len(strValues))
	for i, v := range strValues {
		num, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		values[i] = num
	}
	N = values[0]
	T = values[1]
	// 2. Создайте буферизированный канал для результатов (размер = N)
	results := make(chan string, N)

	// 3. Для каждого числа от 1 до N запустите горутину
	for i := 1; i <= N; i++ {
		go func(num int) {
			// 4. Вычислите время обработки для этого числа: num * 30
			processTime := time.Duration(num*30) * time.Millisecond
			// 5. Создайте буферизированный канал (размер 1) для результата обработки
			resultCh := make(chan int, 1)
			// 6. Запустите вложенную горутину для обработки:
			//    - Используйте time.Sleep(processTime)
			//    - Отправьте результат (num * 2) в канал
			go func() {
				time.Sleep(processTime)
				resultCh <- num * 2
			}()
			// 7. Реализуйте таймаут с помощью select:
			select {
			case res := <-resultCh:
				results <- fmt.Sprintf("Число %d: %d", num, res)
			case <-time.After(time.Duration(T)*time.Millisecond + time.Millisecond):
				results <- fmt.Sprintf("Число %d: timeout", num)
			}
		}(i)
	}
	// 8. Получите N результатов из канала и выведите их
	for i := 0; i < N; i++ {
		fmt.Println(<-results)
	}
}

package main
import (
	"fmt"
)

var table = make([]int, 10)

func main(){
	initialiseTable()
	fmt.Println("Hello!")
	insertItem(3)
	insertItem(13)
	printTable()
}
func printTable(){
	for i:=0; i<10; i++{
		if table[i] != 0{
			fmt.Println("index : ", i, "value : ", table[i])
		}
	}
}

func insertItem(key int){
	location := key % 10
	count :=0
	for ; table[location] != 0 && count < 10; {
		count++
		if location == 9{
			location = 0
		}else{
			location++
		}
			
	}
	if count == 10{
		fmt.Println("Table is full!")
	}else{
		table[location] = key
	}
}

func initialiseTable(){
	for i :=0; i<10; i++{
		table[i] = 0
	}
}

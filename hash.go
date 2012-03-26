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
	fmt.Println("key : ", 13, searchTable(13))
	fmt.Println("key : ", 13, "isDelete?",  deleteItem(13))
	fmt.Println("key : ", 13, searchTable(13))
	fmt.Println("Print all Table")
	printAllTable()
}
func printTable(){
	for i:=0; i<10; i++{
		if table[i] != 0{
			fmt.Println("index : ", i, "value : ", table[i])
		}
	}
}
func printAllTable(){
	for i:=0; i<10; i++{
		fmt.Println(i, " : ", table[i])
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

func searchTable(key int) int{
	location := key % 10
	count := 0
	for ; table[location] != 0 && table[location] != key && count < 10 ;{
		count++
		if location == 9{
			location = 0
		}else{
			location++
		}
	}
	if table[location] == key{
		return location
	}
	return -1
}

func deleteItem(key int) bool{
	location := key % 10
	count := 0
	for ; table[location] != 0 && table[location] != key && count <10 ; {
		count++
		if location == 9{
			location = 0
		}else{
			location++
		}
	}
	if table[location] == key{
		table[location] = 0
		return true
	}
	return false
}


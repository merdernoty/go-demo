package main

import "fmt"

func pointerToPointer() {
    x := 42
    ptr := &x       // указатель на x
    ptrPtr := &ptr  // указатель на указатель
    
    fmt.Printf("x = %d\n", x)
    fmt.Printf("*ptr = %d\n", *ptr)
    fmt.Printf("**ptrPtr = %d\n", **ptrPtr) // двойное разыменование
    
    // Изменение через указатель на указатель
    **ptrPtr = 100
    fmt.Printf("После изменения x = %d\n", x) // 100
}

func slicePointers() {
    slice := []int{1, 2, 3, 4, 5}
    
    // Указатель на срез (на заголовок среза)
    slicePtr := &slice
    
    fmt.Printf("Оригинальный срез: %v\n", slice)
    fmt.Print(&slicePtr)
    
    // Изменение среза через указатель
    *slicePtr = append(*slicePtr, 6)
    fmt.Printf("После append: %v\n", slice) // [1 2 3 4 5 6]
    
    // Указатели на элементы среза
    firstElement := &slice[0]
    *firstElement = 999
    fmt.Printf("После изменения первого элемента: %v\n", slice) // [999 2 3 4 5 6]
}


// func main() {
//     // pointerToPointer()
// 	 slicePointers() 
// }


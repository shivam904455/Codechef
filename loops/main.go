package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("value of i =%d\n", i)
	}

	for i:=0; i<=100 ; i+=10 {
		fmt.Printf("value of i =%d\n", i)
	}
	for i:=0; i<=100;i++{
		if i==23{
continue
		}
		fmt.Println("value of i =", i)
	}
	for i:=0; i<=100;i++{
		if i==23{
break
		}
		fmt.Println("value of i =", i)
	}


	Size:=[] string{"big","testy","small"}
	Fruit:=[]string{"Apple","banana","Greps","orange","pomrgrenate","gwawa"}

	for x:=0; x< len (Size) ;x++{
		for Y:=0 ; Y<len (Fruit); Y++{
			fmt.Println(Size[x],Fruit[Y])
 
		}
	}

for i:=0 ;i<26 ;i++{
	if i %2 ==0 {
		fmt.Printf("%c ",'A' +i)
	}else{
		fmt.Printf("%c " , 'a' +i)
	}
}

}

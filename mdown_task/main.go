package main

import "fmt"

//http://www.oshioki.net/sample/750/images/photo/05.jpg

func main() {

	var url="wget http://www.oshioki.net/sample/%d/images/photo/0%d.jpg -O %d-%d.jpg"

	for i:=400;i<750;i++{
		for j:=1;j<6;j++{
			fmt.Println(fmt.Sprintf(url,i,j,i,j))
		}
	}
}

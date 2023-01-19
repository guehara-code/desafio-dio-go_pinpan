package main //pacore principal//

import "fmt"

//se numero divisivel por 3 imprime Pin
//se numero divisivel por 5 imprime Pan
//se numero divisivel por 3 e 5 imprime PinPan
//se não for divisivel por 3 e nem por 5 imprime o numero

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else {
			if i%3 == 0 {
				fmt.Println("Pin")
			} else {
				if i%5 == 0 {
					fmt.Println("Pan")
				} else {
					fmt.Println(i)
				}
			}

		}

	}
}

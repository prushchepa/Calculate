package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculate")
	var oper string
	fmt.Scan(&oper)
	components := strings.Split(oper, " ")
	n = components[1]string 

	if len(components) != 3 {
		fmt.Println("Паника")
	    } else {		
		oper1, err := strconv.Atoi(components[0])	
		oper2, err := strconv.Atoi(components[2])
		if n=="+"{
			fmt.Println(oper1 + oper2)
		}
		if n=="-"{
			fmt.Println(oper1 - oper2)
		}
		if n=="*"{
			fmt.Println(oper1 * oper2)
		}
		if n=="/"&& oper2 != 0 {
			fmt.Println(oper1/oper2)
			} else {
				fmt.Println("Паника")
			}
		


        romoper1, err = convertRomanToArabic1(components[0] string)
		romoper2, err = convertRomanToArabic1(components[2] string)
		if n=="+"{
			a=(romoper1 + romoper2)
		}
		if n=="-"{
			fmt.Println("Паника")
		}
		if n=="*"{
			a=(romoper1 * romoper2)
		}
		if n=="/" {
			a=(romoper1/romoper2)
		    } else {
				fmt.Println("Паника")
			} 


			var rim strings.Builder
			for a > 0 {
				if a / 100 >= 1 {
					rim.WriteRune(RomanNumbers.C);
					a -= 100;
				}
	
				if a / 50 >= 1 {
					rim.WriteRune(RomanNumbers.L);
					a -= 50;
				}
	
				if a / 10 >= 1 {
					int count = a / 10;
	
					for int i = 0; i < count; i++ {
						rim.WriteRune(RomanNumbers.X);
						a -= 10;
					}
				
	
				if (y >0) {
					rim.WriteRune(RomanNumbers.values()[y - 1]);
					y -= y;
				}
				fmt.Println(rim)
			}
		}	
	}	
		
}


func convertRomanToArabic1(components[0], components[2] string) int {
	switch   {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		fmt.Println("Паника")
	}
}

func RomanNumbers(a int) srting {
	switch   {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 10:	
		return "X"
	case 50:
		return "L"
	case 100:	
		return "C"	
	default:
		fmt.Println("Паника")
	}
}








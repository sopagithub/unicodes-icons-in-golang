package main
import (
    "fmt"
    "math/rand"
    "time"
)


func main() {
    
    args := [5]string{"\U0001F334", "\u2F4A", "\U0001F332", "\U0001F333", "\U0001F38B"}
    
    
     var randomNum int
     for i := 0; i < len(args); i++ {
    	var vector [5]string
    	var temp [5]string
    	for j := 0; j < len(args); j++ {
    		rand.Seed(time.Now().UTC().UnixNano())
		    randomNum = rand.Intn(5 - 0) + 0
		    if(j==0) {
			    vector[j] = args[randomNum] 
			    temp[j] = args[randomNum]
                //fmt.Println(args[randomNum])
		    } else {
		        
               for {
                    rand.Seed(time.Now().UTC().UnixNano())
                    randomNum = rand.Intn(5 - 0) + 0
                    b  := contains(vector, args[randomNum])
                    if b == false {
            			vector[j] = args[randomNum] 
            			temp[j] = args[randomNum]
            			//fmt.Println(args[randomNum])
                        break
                    }
                }
                
		    }
		}

    	for k := 0; k < len(temp); k++ {
    		fmt.Print(temp[k])
    	}
	    fmt.Println("\n")
    }

}

func contains(s [5]string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


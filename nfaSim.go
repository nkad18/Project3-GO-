package nfa

import (
    "sync"
)
// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// An symbol in the NFA is a single rune, i.e. a character.
type symbol rune

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym symbol) []state


var mutex = sync.Mutex{}// should this be inside or not


// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(transitions TransitionFunction, start, final state, input []symbol) bool {
       
	   var found bool =false
       var w sync.WaitGroup
	   
       m := make(chan int,1)
       w.Add(1)
       go tran(transitions, start, final, input, m, &w)
       w.Wait()
       
       close(m)
	   for num := range m {
		if num==2{
		
        found = true
		}
		//fmt.Println("channel range vales")
        }
		if found==true{
		    //fmt.Println(input, "true")
			return true
		}
	
	return false
}

func tran(transitions TransitionFunction, start, final state, input []symbol, c chan <- int, waitG *sync.WaitGroup){
    
	defer waitG.Done()
    var i int =0
	var curr []state
	var fork1 int = 0
	var last state
	//var tstate state
	//c2 := make(chan int)
	//fmt.Println(start)
	curr = transitions(start, input[i])
	//fmt.Println(start," ",curr," ",input[i], input, i)
	i = i+1
	if len(curr)==0{ //will have loop elsewhere to
	   //do we fail when we are in a state that doesn't have next states based on our inputs?
	   return
	   }
	if len(curr)>1{
	       fork1=1
	}
	//last = start
	
	if((len(curr)==1)&&(i==len(input))){
	   if curr[0]==final{ //fmt.Println("first")
		   mutex.Lock()
		   select {
           case c <- 2: // Put 2 in the channel unless it is full
           default:
		   }
		   mutex.Unlock()
		return
	   }
	}
	
	
	if fork1==1&&i==len(input){
		//create a new go routine for each new path you can fork to
		for z:=0; z<len(curr); z++{
	       if curr[z]==final{ //fmt.Println("sec")
		     mutex.Lock()
		     select { 
              case c <- 2: // Put 2 in the channel unless it is full
             default:
		   }
		   mutex.Unlock()
			}
		}
	    return
	}
	
	
	   
	   
    for (i<len(input)&&fork1==0){
        curr = transitions(curr[0], input[i])
        i = i+1
	   if len(curr)==0{ //will have loop elsewhere to
	   return
	   }
	   if len(curr)>1{ // fork: have multiple next states based on input
	   /*if i==0{
	       i=i+1
	   }*/
	   fork1 = 1
	   break
	   }else {
	   last = curr[0]
	   if i==len(input){
		   //fmt.Println("first")
	   if last==final{ 
		   mutex.Lock()
		   select {
           case c <- 2: // Put 2 in the channel unless it is full
           default:
		   }
		   mutex.Unlock()
		return
	   }
	   return
	   }
	   }
	   //we're not at the end of string and there was only 1 state
	   //fmt.Println("first created", len(curr)," ",curr[0]," ", i) causes index out of bounds becasue I dont
	   //know at this point if array is empty (if it has an index 0 or not )
		
	}//there is more than one option to transition to 
    if fork1==1&&i<=len(input)&& i>0{ 
		if i ==len(input){
		    //create a new go routine for each new path you can fork to
		for z:=0; z<len(curr); z++{
	       if curr[z]==final{ //fmt.Println("sec")
		     mutex.Lock()
		     select { 
              case c <- 2: // Put 2 in the channel unless it is full
             default:
		   }
		   mutex.Unlock()
			}
		}
	    return
		    
		}
         //fmt.Println("create", len(curr)," ",curr[0]," ", i)
         g:=resize(i,input[:])
		//create a new go routine for each new path you can fork to
		for y:=0; y<len(curr); y++{ //fmt.Println(curr," ", g)
			waitG.Add(1)
			//don't need to use input[i+1:] because already on char after
			//the char that cuased the fork
			
			go tran(transitions, curr[y], final, g, c, waitG)
		}
		
	    return
		
		
	}
	//for when you only have 1 symbol and it goes in mult states
	/*if fork1==1&&i==len(input)-1{
         
		//create a new go routine for each new path you can fork to
		for z:=0; z<len(curr); z++{
		    tstate = transitions(curr[z], input[i])
	       if tstate==final{ fmt.Println("third")
		     mutex.Lock()
		     select { 
              case c <- 2: // Put 2 in the channel unless it is full
             default:
		   }
		   mutex.Unlock()
			}
			
		}
		
	    return
		
		
	}*/

}
func resize(i int, x []symbol) []symbol{

if i< len(x){
 g := make([]symbol,len(x)-i)
       for y:=0; y<len(g); y++{
      	   g[y]= x[i+y]
       }
    return g
}
return nil
}


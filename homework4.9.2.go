

package main 



import ("fmt"// format library
			"reflect"// used for the interface object
			"os"// string parsing input
			"bufio"// string parsing input
			"strconv" // string converter
			"math"// Math 
			"stack"// stack 
			"errors"// errors 
			)

	var parenthesisCounter int = 0 // global parenthesis counter used to determine where in the recursion we are.....
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////	
	//	Description: This function is designed to do the calculations on a valid arithmatic expression 
	//  and, if nessisary, use recursion when there is a sub expression. This function walks through 
	//  a string containing the expression and reads it in. if there is an error we return the error to the user. 
	//
	//	Input: this function takes in the current index, thewn last index, and the string of characters 
	//  in this respective order
	//
	//	Output: This function returns an interface object that is either of float64 or int type.
	//
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	
	func calculate(index* int,last int, buffer string)(interface{}, error) {
	
		 operandStack := stack.New() // operand stack used to hold all of the operands in the current stack frame. 
		 operatorStack := stack.New() // operator stack used to hold all of the operators on the current stack frame. 
		
		var leftInteger interface{}	// variable of interface type used to hold the incomming or outgoing variablees to other functions
		var rightInteger interface{}		// variable of interface type used to hold the incomming or outgoing variablees to other functions
		character := " "							// temp character
		var operator interface{}// temp interfaces for calculatios
		var err error// temp variable for error 
		
		  operatorEncountered:= true			// flags to do error checking 
		  endParenthesis:= false
		  doNextCalc:= false
	
	
			if err != nil{				// bubble up the error if present after recursion.... otherwise continue
			return 0, err
			}
	
	
	
		for *index <= last{							//while we havent seen the end of the string
					
													//if not an applicable character then nope error
				if string(buffer[*index]) !="." && string(buffer[*index]) !=" " && string(buffer[*index]) !="(" && string(buffer[*index]) !=")"&&string(buffer[*index]) !="+" &&string(buffer[*index]) !="-" &&string(buffer[*index]) != "*"&&string(buffer[*index]) !="/" && (int(buffer [*index]) < 48 || int(buffer [*index]) > 57){				
						return 0, errors.New("Not an expression asfdas")
				}
					
				if string(buffer [*index])== " "{							//if we encounter a space then skip it
						*index++
				}
				
				
				if string(buffer [*index]) ==  "("{							//if we encounter a open parenthesis then we start a recursion and push a new frame onto the stack
					//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
					if operatorEncountered == false {							//if we encounter no operator.. ie some variation of the error "4(" 
						return 0, errors.New("Not an expression")
					}
					 if endParenthesis == true {    									//--> are two parenthesis back to back )(
            			return 0, errors.New("Not an expression")
					}
					
					*index++							//move past the parenthesis 
					parenthesisCounter++							//mark the parenthesis on the variable
					
					leftInteger, err = calculate(index,last, buffer)							//call ourselves
					if err != nil {							//check for an error
					return 0, err
					}
					operandStack.Push(leftInteger)							//push onto the stack the new operand
					endParenthesis = true;							//flag the parenthesis
					operatorEncountered = false // removes the possabliity of "operator"    ")"
					*index++							//increment the index
				
				if *index > last {									//if we are at the end of the string then exit the main loop and clean up the stacks
					break
				}
				}
				//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
				
				
				
				
				
				
				
				if  string(buffer[*index])== ")"{							//if we encounter an ending parenthesis then we move to clean the stacks and return the current expression result
				
				if *index == last {					// resolves mutiple recursions at ending case
					parenthesisCounter--;
					break
				}
						//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
						parenthesisCounter--;							//decrement the global variable
						
									
						
						if parenthesisCounter <0{							//there is a logical error or we have  encountered to many ")" -- in either case error out
						return 0, errors.New("Not an expression")

						}

						if operatorEncountered == true && parenthesisCounter%2 != 0{	// removes the possabliity of "operator"
							return 0, errors.New("Not an expression")
						}						
						
						 if (operatorStack.IsEmpty() == false) {							//begin cleaning the current stacks
								for (operatorStack.IsEmpty() == false){
															   
								   
										if !operandStack.IsEmpty(){
											rightInteger, err = operandStack.Pop()
											errorHandle(err)
										}
										if !operatorStack.IsEmpty() {
											operator,err = operatorStack.Pop()
											errorHandle(err)
										}
										if !operandStack.IsEmpty(){
											leftInteger,err = operandStack.Pop()
											errorHandle(err)
										}
			

										//check for a division by zero error and prevent it if nessisary-- otherwise continue with the calculations
										if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {
							
											if rightInteger.(float64) == 0.0 && operator.(string) == "/" {
											return 0, errors.New("Division by zero encountered")

											} else {
												operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
											}
										} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
										
											if (operator.(string) == "/" && rightInteger.(int) == 0){
												return 0, errors.New("Division by zero encountered")

											} else {
												operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
											}
										}
								}// end for loop 
								leftInteger, err = operandStack.Pop()// return the calculation	
								if err != nil{panic("ErroR")}
								return leftInteger, nil
						} else  if parenthesisCounter > 0{// if we are inside a recursion and there is nothing opn the operand stack
						
						leftInteger, err = operandStack.Pop()
											if err != nil{panic("ErroR")}
									return leftInteger, nil
						
						} 
			}
			
			
			
			
										//if we encounter a number process it as such and stick it on the stack
				if int(buffer [*index]) >= 48 && int(buffer [*index]) <= 57 || string(buffer [*index]) == "."{
					
						
					if endParenthesis == true {										//if there is a number directly next to a parenthesis
						
						return 0, errors.New("Not an expression")
					}
						//fmt.Println(*index,"  ",string(buffer [*index]))
					
					
					if *index >= last && operatorEncountered == false {										//if we are at the end of the string 
						break 
					}
					
					
						if  operatorEncountered == true{										//if there has been an operator encountered
							operatorEncountered = false
							leftInteger = getNumber(index,last , buffer )
							operandStack.Push(leftInteger );
						
					} else {
							return 0, errors.New("Not an expression")										//otherwise not an expression 
						
						}
					
				}
							//if an operator character is encountered
				if string(buffer [*index]) ==  "*"||string(buffer [*index]) =="/"||string(buffer [*index]) =="+"||string(buffer [*index]) =="-"{
						endParenthesis = false 
						operatorEncountered = true
								//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
				
						if doNextCalc == true{										//this segment is used when we encounter an operator and already have something on the stack with a greater or equal precdence
								doNextCalc= false
							
								if !operandStack.IsEmpty() {										//here we do the next calculation and then push its answer to the stack along with the operator
								   rightInteger, err = operandStack.Pop()
										errorHandle(err)
								}
								
								operator, err = operatorStack.Pop()
								errorHandle(err)
								if !operandStack.IsEmpty() {
								   leftInteger, err = operandStack.Pop()
									errorHandle(err)
								}

									
								if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {										//check for divide by zero
					
									if rightInteger.(float64) == 0.0 && operator.(string) == "/" {
									return 0, errors.New("Division by zero encountered")

									} else {
										operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
									}
								} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
								
									if (operator.(string) == "/" && rightInteger.(int) == 0){
										return 0, errors.New("Division by zero encountered")

									} else {
										operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
									}
								}
						}
						
						
						//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
																	// since we encountered an operator we must deal with it... there are treee posablities
						if operatorStack.IsEmpty() == true{										//if the stack is empty then place it on the stack
						
							operatorStack.Push(string(buffer[*index]))
									*index++
						}else  {										//										//otherwise we must findout the precednece of the operator in relation to the current character in the string buffer
								// there is already something on the operand stack										
								leftInteger = operatorStack.Top()
								character = leftInteger.(string)
								
							if (string(buffer[*index]) == "*" || string(buffer[*index] ) == "/")&& (character == "+"|| character == "-" ){										//if it has lower precedence then place it on the 
								
								operatorStack.Push(string(buffer[*index]))
								doNextCalc = true
								*index++
								
							} else if (string(buffer[*index]) == "+" || string(buffer[*index])  == "-")&& (character == "*" || character == "/" ){										//if it has higher precedence then use the operator currently on the stack with its operancd and then place this new operato onto the stack
									if !operandStack.IsEmpty() {
									   rightInteger, err = operandStack.Pop()
											errorHandle(err)
									}
									
									if !operatorStack.IsEmpty() {
									operator, err = operatorStack.Pop()
									errorHandle(err)
									}
									if !operandStack.IsEmpty() {
									   leftInteger, err = operandStack.Pop()
										errorHandle(err)
									}
							
							
									if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {										//watch for divide by zero 
								
									if rightInteger.(float64) == 0.0 && operator.(string) == "/" {
									return 0, errors.New("Division by zero encountered")

									} else {
										operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
									}
									} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
								
									if (operator.(string) == "/" && rightInteger.(int) == 0){
										return 0, errors.New("Division by zero encountered")

									} else {
										operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
									}
									}
									operatorStack.Push(string(buffer[*index]))
									*index++
								
							} else {										//otherwise both operands have same precedence so do one calculation and then place the new operand onto the stack										//
							
									if !operandStack.IsEmpty() {
									   rightInteger, err = operandStack.Pop()
											errorHandle(err)
									}
									
									if !operatorStack.IsEmpty() {
									operator, err = operatorStack.Pop()
									errorHandle(err)
									}
									if !operandStack.IsEmpty() {
									   leftInteger, err = operandStack.Pop()
										errorHandle(err)
									}
								
									if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {										//watch for divide by zero
								
											if rightInteger.(float64) == 0.0 && operator.(string) == "/" {
											return 0, errors.New("Division by zero encountered")

											} else {
												operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
											}
									} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
							
										if (operator.(string) == "/" && rightInteger.(int) == 0){
											return 0, errors.New("Division by zero encountered")

										} else {
											operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
										}
									}
									operatorStack.Push(string(buffer[*index]))
									*index++
								}
							
							}				
					}
				}// end of loop 										//if here then we probably have stuff on the stack that need cleaning
	  
	  
			  // clean up the recursion stacks
			  // if we dont have a parenthesis error and there is something on the operator stack--> resolve the
			  // remaining expressions or just return the answer
			
			
			
			if parenthesisCounter == 0 && operatorStack.IsEmpty() == false{									//if there is no parenthesis error and we have an operator on the stack 
				
					for !operatorStack.IsEmpty() {					// cleaning off stack for the end of the calculatios 
						if (operatorStack.IsEmpty() == true){
						break 
						}
					
						if !operandStack.IsEmpty() {
						   rightInteger, err = operandStack.Pop()
								errorHandle(err)
						}
						
						if !operatorStack.IsEmpty() {
						operator, err = operatorStack.Pop()
						errorHandle(err)
						}
						if !operandStack.IsEmpty() {
						   leftInteger, err = operandStack.Pop()
							errorHandle(err)
						}
								
				
						if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {									//watch for divide by zero
						
							if rightInteger.(float64) == 0.0 && operator.(string) == "/" {
							return 0, errors.New("Division by zero encountered")

							} else {
								operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
							}
						} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
						
							if (operator.(string) == "/" && rightInteger.(int) == 0){
								return 0, errors.New("Division by zero encountered")

							} else {
								operandStack.Push(answerInterface(leftInteger, rightInteger, operator));
							}
						}
											
				}// end clean off loop 
			
					if reflect.TypeOf(rightInteger) == reflect.TypeOf(1.2) {
				
					if rightInteger.(float64) == 0.0 && operator.(string) == "/" {									//watch for divide by zero
				return 0, errors.New("Division by zero encountered")

					} else {
					return answerInterface(leftInteger, rightInteger, operator), nil
					}
				} else if reflect.TypeOf(rightInteger) == reflect.TypeOf(1){
				
					if (operator.(string) == "/" && rightInteger.(int) == 0){
						return 0, errors.New("Division by zero encountered")

					} else {
					return answerInterface(leftInteger, rightInteger, operator), nil
					}
				}
	}else if  parenthesisCounter == 0 && operatorStack.IsEmpty() && ! operandStack.IsEmpty() {									//otherwise if we have an operator on the stack and no errors

            if !operandStack.IsEmpty() {
						   rightInteger, err = operandStack.Pop()
								errorHandle(err)
						}
						
			return rightInteger, nil
	} else if parenthesisCounter !=0 {									//otherwise we have an error
	
	
	return  0, errors.New("Not an expression")
	
	}
		return 0, nil
	}
	
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////	
	//	Description: Simple error handle function for when we return from a function
	//
	//	Input: the error
	//
	//	Output: hopefully nothing
	//
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	func errorHandle(err error){
	
		if err != nil{
		
			panic(err);
		}
	
	}
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////	
	//	Description: This function takes in 3 iterface types, in a specific order and returns the calculation based on the interface types
	//	there is no divide by zero error in this function. in retrospect there should have been but I built it outside of the function
	//
	//	Input: 2 interface types of either in or float in any combination and one iterface type of string type for the operator
	//
	//	Output: one interface type of either int or float
	//
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	func answerInterface(leftInterface interface{}, rightInterface interface{}, operatorInterface interface{})(interface{}){
		
		// find situation that we are in and then calculate and return an interface
		rightI:= 0									//temporary variables
		leftI:= 0
		rightF:= 0.0
		leftF:= 0.0
		floatBoolL := false									//flags for if we encounter an integer
		floatBoolR := false
		operator := "S"											//operator variable
		
		////////////////////////////////////////////////////////////////////////////////////////////////////////////
		t:= reflect.TypeOf(leftInterface)																						// determine what the variables are and extract there values to proper variables
			if t == reflect.TypeOf(leftI){		// if value is int type then use int type
				leftI= leftInterface.(int)
			}else if t == reflect.TypeOf(leftF){	
			leftF= leftInterface.(float64)
				floatBoolL = true
			}
		//-----------------------------------------------	
			t= reflect.TypeOf(rightInterface)
			if t == reflect.TypeOf(leftI){		// if value is int type then use int type
				rightI= rightInterface.(int)
			}else if t == reflect.TypeOf(leftF){	
			rightF= rightInterface.(float64)
				floatBoolR = true
	
			}
		//-----------------------------------------------
		t= reflect.TypeOf(operatorInterface)
			if t == reflect.TypeOf(operator){		// if value is int type then use int type
				operator= operatorInterface.(string)
			}
		//////////////////////////////////////////////////////////////////////////////////////////////////////////////			
			
			if floatBoolL == true &&  floatBoolR == false {									//									//based on the situation use the appropreate method the do the coersions and or calculations
				switch operator{

					case "+":
						return interface{}(leftF + float64(rightI)); // make this for all of the cases --> I want to just push the inerface onto the stack---- i may not need to do this because its a template type 
					case "-":
						return leftF - float64(rightI); 
					case "*":
						return leftF * float64(rightI); 
					case "/":
						return leftF / float64(rightI); 
					default :
						fmt.Println("OPERATOR: ",operator);
			
						panic("\n\n\nerror on Line 488. NO OPERATOR\n\n\n")
				
				}
			}else if floatBoolL == false && floatBoolR == true{
				
				switch operator{		
					case "+":
						return float64(leftI)+ rightF; 
					case "-":
						return float64(leftI) - rightF; 
					case "*":
						return float64(leftI) * rightF; 
					case "/":
					
							 if rightF != 0{
						return float64(leftI) / rightF; 
							} //else 
							
						default :
						panic("\n\n\nerror on Line 506. NO OPERATOR\n\n\n")
				}
			
			} else if floatBoolL == true && floatBoolR == true  {
		
				switch operator{	
					case "+":
						return leftF + rightF; 
					case "-":
						return leftF - rightF; 
					case "*":
						return leftF * rightF; 
					case "/":
					
						if rightF != 0{
						return leftF / rightF; }
						default :			
						panic("\n\n\nerror. NO OPERATOR\n\n\n")
				}
			}else if  floatBoolL == false && floatBoolR == false{
				switch operator{
				
					case "+":
						return leftI + rightI; 
					case "-":
						return leftI - rightI; 
					case "*":
						return leftI * rightI; 
					case "/":
					
						if rightI != 0{
						return float64(leftI) / float64(rightI); 
						}
					default :
						panic("\n\n\nerror on Line 540. NO OPERATOR\n\n\n")
						
				}		
			}
	return nil
	
	}
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////	
	//	Description: gathers the appropreate numbers from the string at the current position
	//
	//	Input: The current index position, the last position in the sting, and the string
	//
	//	Output: an appropreate interface type with the value 
	//
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	
	func getNumber(index* int,last int, buffer string) (interface{}) {/// returning a 'vector' of ordered variables
			
			decimalPoint := false									//flag for the decimal point if applicable
			intergeNumber := 0
			float64Number := 0.0									//temp variables
			precision := 1
			exit := false									//exit situation 
			//--------------------------------------------------
		for *index <= last{													//while we havent seen the last of the string
			
	
			switch string(buffer[*index]){									//switch for the current character
			
				case "0","1","2","3","4","5","6","7","8","9" :									//if its a number
					
					if decimalPoint == false{									//if its not a decimal 

						n,err := strconv.Atoi(string(buffer[*index]))
						
						if err == nil {
							intergeNumber= intergeNumber* 10 + (int)(n)										// append the number to the current number
							
							if *index !=  last{
							*index++
							} else { exit = true}
							
				
						}
					} else {																						// decimal point present
					
						n,err := strconv.Atoi(string(buffer[*index]))
						if err == nil {
							intergeNumber= intergeNumber* 10 + (int)(n)
							
							if *index !=  last{
							*index++}else { exit = true}
						}
					value :=math.Pow(0.1 ,float64(precision))										//append value
					
					float64Number= float64Number +value* (float64)(n) 										//move value to temp float number
					precision++
					}
					
				case  "." :										//										//	alert decimal flag
				float64Number = float64(intergeNumber)
				decimalPoint = true
				*index++	
			
				 
				default : 										//exit situation 
				exit = true
		
		
			 }
			 
			 if exit == true{
				break
			 }
			 
		 }
		 
		 if decimalPoint == true{										//return appropreate value
		 
			var empty interface{} = float64Number
			return empty
		} else {
		
			var empty interface{} = intergeNumber
			return empty
		}
	}
	
	
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////	
	//	Description: This function is the main function for the program and is used to process the testExpressions.txt file 
	//	and then to print out the result of the expressions contained therein
	//
	//	Input: nothing
	//
	//	Output: nothing
	//
	//
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	
func main() {
	var i int =0 

  file, err := os.Open("C:\\Users\\John\\Desktop\\Go Programs\\New folder\\src\\testExpressions.txt")					//open the file
  defer file.Close()					//defer closing of the file
  
  if err != nil {
        fmt.Println(err)
    }
 scanner := bufio.NewScanner(file)					//set up the scanner for the strings

	for scanner.Scan(){					//while there is something to read

		buffer := scanner.Text()					//grab a line of text
		i++					//increment the index for a display counter
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~  ", i)
		
		last := len(buffer)-1					//set up index
	
// the zero represents the statring index
		index := 0					//reset the index
		parenthesisCounter =0						//refresh the global variable
		 intNumber, err := calculate(& index,last, buffer)					//start the calculations
		
	  
		if err == nil{										//if we get a real result
			if reflect.TypeOf(intNumber) == reflect.TypeOf(1.1){					//if it is a floating point number print it as such
			
			fmt.Printf("%v = \n%9.5E",buffer,intNumber)
			fmt.Println("")
			} else {					//if it is an integer point number print it as such
			
			
			fmt.Printf("%v = \n%d",buffer,intNumber)
			fmt.Println("")
			
			
			
			}
		} else {					//if it is an error  print it as such
			fmt.Println(buffer)
			fmt.Println("error: ",err)
		}
	}
	
}
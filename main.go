package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

    var movieName = "Kalki mivie"
	const ticketsCount = 50
	var remaningTickes uint = 50
	var bookings = make([]userData, 0)

	type userData struct{

		firstName string 
		lastName string
		email string
		numberofTickets uint
	}

	

func main() {
	
	router := gin.Default()
	router.GET("/greet", greetUsers)
	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Printf("Failed to run server: %v\n", err)
		}
	 }()

	ticketBookingLoop()
	}
	
	//greetUsers(&gin.Context{})
func ticketBookingLoop(){
	for {

		firstName,lastName,email,userTickets:=getUserInput()

		isvanlideName,isvalideEmail,isValideTicketNumber:=validateUsers(firstName,lastName,email,userTickets)
	
		if isvanlideName && isvalideEmail && isValideTicketNumber{
		bookTicket(userTickets,firstName,lastName,email)
		go sendTickets(userTickets,firstName,lastName,email)

	  

		//Colling the print First Name Function
		firstNames :=getFirstName()
		fmt.Printf("First names of our bookings are : %v\n",firstNames)


	   if remaningTickes==0{
		println("House full")
		break
	   }
		}else{
			if !isvanlideName{
				println("please check first name and last name")
			}
			if !isvalideEmail{
				println("plesae enter validce email")
			}
			if !isValideTicketNumber{
				println("plesae enter vlaide number of tickets")
			}
		}
		
	}
}

func greetUsers(context *gin.Context){
	response := gin.H{

	"WelcomeMessage" : fmt.Sprintf("Welcome to our %v booking applicaton\n", movieName),
	"TicketIno" : fmt.Sprintf("We have total of %v tickets and %v are remaning\n", ticketsCount, remaningTickes),
	 "bookingPrompt" : fmt.Sprintf("Place book your tickets here\n"),

	}
	context.IndentedJSON(http.StatusOK, response)

}

func validateUsers(firstName string,lastName string,email string,userTickets uint)(bool,bool,bool){
	isvanlideName :=len(firstName)>=2 && len(lastName)>=2
	isvalideEmail := strings.Contains(email, "@")
	isValideTicketNumber := userTickets>0 && userTickets<=remaningTickes

	return isvanlideName,isvalideEmail,isValideTicketNumber


}
func getFirstName() [] string{
	var firstNames[] string
		for _,booking := range bookings{
		
			firstNames=append(firstNames, booking.firstName)
		} 
	    // fmt.Printf("Whole Slice:  %v\n",bookings)
	   // fmt.Printf("First element of Slice:  %v\n",bookings)
	   // fmt.Printf("Type of Slice: %T\n",bookings)
	   // fmt.Printf("Size of Slice: %v\n",len(bookings))
	   //fmt.Printf("These are all our bookings: %v\n",bookings)
	   //fmt.Printf("First names of our bookings are : %v\n",firstNames)
	   return firstNames

}



func getUserInput()(string,string,string,uint){

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	println("Enter your First Name")
	fmt.Scan(&firstName)
	println("Enter your Last Name")
	fmt.Scan(&lastName)
	println("Enter your email")
	fmt.Scan(&email)
	println("Enter No Of Tickets")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets

}

func bookTicket(userTickets uint, firstName string , lastName string,email string) {
	remaningTickes = remaningTickes - userTickets
	//Creatikng map for user
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberofTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"]    = email
	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets),10)
	bookings = append(bookings, userData)

	fmt.Printf("Tank you %v %v for booking %v tickts, you will recive a conformation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets are remaning for %v\n",remaningTickes,movieName)
	fmt.Printf("User Details %v\n",bookings)
   
}

func sendTickets(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets,firstName,lastName)
	fmt.Println("##############")
	fmt.Printf("%v has been sent to %v\n",ticket,email )
	fmt.Println("##############")
}

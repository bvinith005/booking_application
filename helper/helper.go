package helper

import "strings"

func ValidateUsers(firstName string,lastName string,email string,userTickets uint,remaningTickes uint)(bool,bool,bool){
	isvanlideName :=len(firstName)>=2 && len(lastName)>=2
	isvalideEmail := strings.Contains(email, "@")
	isValideTicketNumber := userTickets>0 && userTickets<=remaningTickes

	return isvanlideName,isvalideEmail,isValideTicketNumber


}
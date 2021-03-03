package main

import (
	"./services"

	"flag"
)

func main() {

	CreateTNXFlag, VerfiyPaymentFlag,InquiryPaymentFlag,TransactionsPaymentFlag := MasterFlag()
	ChackFlag(CreateTNXFlag, VerfiyPaymentFlag,InquiryPaymentFlag,TransactionsPaymentFlag)
}


func ChackFlag(CreateTNXFlag *bool, VerfiyPaymentFlag *bool,InquiryPaymentFlag *bool, TransactionsPaymentFlag *bool) {
	if *CreateTNXFlag {
		services.CreateTNX()
	}
	if *VerfiyPaymentFlag {
		services.VerfiyPayment()
	}

	if *InquiryPaymentFlag {
		services.InquiryPayment()
	}
	if *TransactionsPaymentFlag {
		services.TransactionsPayment()
	}
}
func MasterFlag() (*bool, *bool,*bool, *bool) {
	CreateTNXFlag := flag.Bool("create", false, "create tnx")
	VerfiyPaymentFlag := flag.Bool("verfiy", false, "verfiy tnx")
	InquiryPaymentFlag := flag.Bool("inquiry", false, "inquiry tnx")
	TransactionsPaymentFlag := flag.Bool("tnx", false, "List all tnx")
	flag.Parse()
	return CreateTNXFlag, VerfiyPaymentFlag,InquiryPaymentFlag,TransactionsPaymentFlag
}

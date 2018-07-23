package main

//EmailMessage ...used to store messages
type EmailMessage []struct {
	From    string
	To      string
	Subject string
	Date    string
}

package main

type (
	TicketMessage struct {
		ID             int       `json:"id"`
		Message        string    `json:"message"`
		SenderType     string    `json:"senderType"`
		SenderID       int       `json:"senderId"`
		AccountSender  *Account  `json:"sender,omitempty" gorm:"polymorphic:Sender;polymorphicValue:account"`
		CustomerSender *Customer `json:"sender,omitempty" gorm:"polymorphic:Sender;polymorphicValue:customer"`
	}

	Account struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	Customer struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

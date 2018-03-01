package message

import "github.com/dpb587/go-schemaorg"

var (
	// The date/time at which the message has been read by the recipient if a single
	// recipient exists.
	DateRead = schemaorg.NewProperty("dateRead")

	// A sub property of recipient. The recipient blind copied on a message.
	BccRecipient = schemaorg.NewProperty("bccRecipient")

	// The date/time at which the message was sent.
	DateSent = schemaorg.NewProperty("dateSent")

	// A sub property of recipient. The recipient copied on a message.
	CcRecipient = schemaorg.NewProperty("ccRecipient")

	// A sub property of participant. The participant who is at the receiving end of
	// the action.
	Recipient = schemaorg.NewProperty("recipient")

	// A sub property of participant. The participant who is at the sending end of
	// the action.
	Sender = schemaorg.NewProperty("sender")

	// A CreativeWork attached to the message.
	MessageAttachment = schemaorg.NewProperty("messageAttachment")

	// The date/time the message was received if a single recipient exists.
	DateReceived = schemaorg.NewProperty("dateReceived")

	// A sub property of recipient. The recipient who was directly sent the message.
	ToRecipient = schemaorg.NewProperty("toRecipient")
)

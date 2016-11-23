package main

import (
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/dragonfax/evernote-sdk-go/notestore"
	"github.com/dragonfax/evernote-sdk-go/types"
)

func strP(s string) *string {
	return &s
}

func main() {
	notestoreUrl := os.Getenv("EVERNOTE_NOTESTORE")

	trans, err := thrift.NewTHttpPostClient(notestoreUrl)
	if err != nil {
		panic(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := notestore.NewNoteStoreClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		panic(err)
	}

	authenticationToken := os.Getenv("EVERNOTE_TOKEN")

	note := types.NewNote()
	note.Title = strP("Test Note")
	note.Content = strP(
		`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE en-note SYSTEM "http://xml.evernote.com/pub/enml2.dtd">
<en-note>  
	<h1>Hello, world</h1>  
</en-note>
`)

	updatedNote, err := client.CreateNote(authenticationToken, note)
	if err != nil {
		panic(err)
	}

	fmt.Println(updatedNote)

}
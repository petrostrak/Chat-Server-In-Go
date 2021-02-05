# Simple Chat Server

A simple chat server written in Golang, with very basic features:


- There is only a single chat room for now
- User can connect to the chat server
- User can set their name
- User can send message to the chat room

### Protocol

For this excersie , a simple text-based message over TCP will be used :

 - All messages are terminated with `\n`
 - To send a chat message, client will send:
    *  `SEND chat message`
    *  For now, chat message can not contain new line.
 - To set client name, client will send:
    *  `NAME username`
    *  For now, username can not contain space
 - Server will send the following command to all clients when there are new message:
    *  `MESSAGE username the actual message`

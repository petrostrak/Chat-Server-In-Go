package tui

import "Chat-Server-In-Go/client"

func StartUi(c client.ChatClient) {
	loginView := NewLoginView()
	chatView := NewChatView()
}

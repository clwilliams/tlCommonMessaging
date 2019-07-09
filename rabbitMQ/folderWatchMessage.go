package rabbitMQ

import (
	"encoding/json"
	"fmt"
)

// CreateAction - represents a folder creation action
const CreateAction = "CREATE"

// DeleteAction - represents a folder deletion action
const DeleteAction = "DELETE"

// RenameAction - represents a folder renaming action
const RenameAction = "RENAME"

// MoveAction - represents a folder move action
const MoveAction = "MOVE"

// FolderWatchMessage - encapsulates a change to a single file / folder
type FolderWatchMessage struct {
	WatchFolder string `json:"watchFolder"` // the folder being watched
	Action      string `json:"action"`      // the type of change being reported
	Path        string `json:"path"`        // the path of the file / folder
	IsDir       string `json:"isDir"`       // whether this is a folder or not
}

// PostToQueue - posts the JSON message
func (msg *FolderWatchMessage) PostToQueue(exchange, routingKey *string,
	messageClient *MessageClient) error {
	jsonMsg, err := msg.convertToJSON()
	if err != nil {
		return fmt.Errorf("Problem marshalling message into JSON %v", err)
	}
	if err := messageClient.SendMessage(*exchange, *routingKey, jsonMsg); err != nil {
		return fmt.Errorf("Problem sending message to message broker %v", err)
	}
	return nil
}

// converts the struct into the JSON body for the message
func (msg *FolderWatchMessage) convertToJSON() (string, error) {
	jsonMsg, err := json.MarshalIndent(msg, "  ", "    ")
	if err != nil {
		return "", fmt.Errorf("Error marshalling JSON %v", err)
	}
	return string(jsonMsg), nil
}

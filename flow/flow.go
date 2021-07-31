package flow

import "github.com/gofrs/uuid"

type Import struct {
	Version   string     `json:"version"`
	Site      string     `json:"site"`
	Flows     []Flow     `json:"flows"`
	Campaigns []struct{} `json:"campaigns"`
	Triggers  []struct{} `json:"triggers"`
	Fields    []struct{} `json:"fields"`
	Groups    []struct{} `json:"groups"`
}

type Flow struct {
	Name               string `json:"name"`
	UUID               string `json:"uuid"`
	SpecVersion        string `json:"spec_version"`
	Language           string `json:"language"`
	Type               string `json:"type"`
	Nodes              []Node `json:"nodes"`
	UI                 UI     `json:"_ui"`
	Revision           int    `json:"revision"`
	ExpireAfterMinutes int    `json:"expire_after_minutes"`
	// Metadata           struct {
	// 	Expires int `json:"expires"`
	// } `json:"metadata"`
}

type Node struct {
	UUID    string   `json:"uuid"`
	Actions []Action `json:"actions"`
	Exits   []Exit   `json:"exits"`
}

type Action struct {
	UUID         string   `json:"uuid"`
	Attachments  []string `json:"attachments"`
	Text         string   `json:"text"`
	Type         string   `json:"type"`
	QuickReplies []string `json:"quick_replies"`
}

type Exit struct {
	UUID            string  `json:"uuid"`
	DestinationUUID *string `json:"destination_uuid"`
}

type UI struct {
	Nodes map[string]UINode `json:"nodes"`
}

type UINode struct {
	Position Position `json:"position"`
	Type     string   `json:"type"`
}

type Position struct {
	Left int `json:"left"`
	Top  int `json:"top"`
}

func NewImport() Import {
	return Import{
		Version:   "13",
		Site:      "https://new.push.al",
		Flows:     make([]Flow, 0),
		Campaigns: make([]struct{}, 0),
		Triggers:  make([]struct{}, 0),
		Fields:    make([]struct{}, 0),
		Groups:    make([]struct{}, 0),
	}
}

func NewFlow() Flow {
	flowUUID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return Flow{
		Name:        "",
		UUID:        flowUUID.String(),
		SpecVersion: "13.1.0",
		Language:    "eng",
		Type:        "messaging",
		Nodes:       make([]Node, 0),
		UI: UI{
			Nodes: make(map[string]UINode),
		},
		Revision:           1,
		ExpireAfterMinutes: 5,
	}
}

func MessagesToFlow(msgs []string) Flow {

	newFlow := NewFlow()

	newFlow.Nodes = NodesFromMessages(msgs)
	newFlow.UI = UIFromNodes(newFlow.Nodes)

	return newFlow
}

func NodesFromMessages(msgs []string) []Node {
	newNodes := make([]Node, 0)
	for i, message := range msgs {
		newNodeUUID, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		newActionUUID, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		newExitsUUID, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}

		newNode := Node{
			UUID: newNodeUUID.String(),
			Actions: []Action{
				{
					UUID:         newActionUUID.String(),
					Attachments:  make([]string, 0),
					Text:         message,
					Type:         "send_msg",
					QuickReplies: make([]string, 0),
				},
			},
			Exits: []Exit{
				{
					UUID: newExitsUUID.String(),
				},
			},
		}
		if i > 0 {
			newNodes[i-1].Exits[0].DestinationUUID = &newNode.UUID
		}
		newNodes = append(newNodes, newNode)
	}
	return newNodes
}

func UIFromNodes(nodes []Node) UI {

	uiNodes := make(map[string]UINode)

	for i, node := range nodes {
		uiNodes[node.UUID] = UINode{
			Type: "execute_actions",
			Position: Position{
				Left: 0,
				Top:  i * 140,
			},
		}
	}

	return UI{
		Nodes: uiNodes,
	}
}

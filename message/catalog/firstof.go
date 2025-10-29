package catalog

import "golang.org/x/text/internal/catmsg"

func FirstOf(messages ...Message) Message {
	return catmsg.FirstOf(messages)
}

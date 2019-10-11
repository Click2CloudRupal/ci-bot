package cibot

import (
	"gitee.com/openeuler/go-gitee/gitee"
	"github.com/golang/glog"
)

// HandleNoteEvent handles note event
func (s *Server) HandleNoteEvent(event *gitee.NoteEvent) {
	if event == nil {
		return
	}
	// just handle create comment event
	if *event.Action != "comment" {
		return
	}

	// add label
	if RegAddLabel.MatchString(event.Comment.Body) {
		err := s.AddLabel(event)
		if err != nil {
			glog.Errorf("failed to add label: %v", err)
		}
	}

	// remove label
	if RegRemoveLabel.MatchString(event.Comment.Body) {
		err := s.RemoveLabel(event)
		if err != nil {
			glog.Errorf("failed to remove label: %v", err)
		}
	}

	// check cla by note event
	if RegCheckCLA.MatchString(event.Comment.Body) {
		err := s.CheckCLAByNoteEvent(event)
		if err != nil {
			glog.Errorf("failed to check cla by note event: %v", err)
		}
	}
}

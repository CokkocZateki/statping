package core

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/plugin"
	"upper.io/db.v3/lib/sqlbuilder"
)

func OnLoad(db sqlbuilder.Database) {
	for _, p := range AllPlugins {
		p.OnLoad(db)
	}
}

func OnSuccess(s *Service) {
	for _, p := range AllPlugins {
		p.OnSuccess(structs.Map(s))
	}
}

func OnFailure(s *Service, f FailureData) {
	for _, p := range AllPlugins {
		p.OnFailure(structs.Map(s))
	}
	slack := SelectCommunication(2)
	if slack == nil {
		return
	}
	if slack.Enabled {
		msg := fmt.Sprintf("Service %v is currently offline! Issue: %v", s.Name, f.Issue)
		SendSlackMessage(msg)
	}
}

func OnSettingsSaved(c *Core) {
	for _, p := range AllPlugins {
		p.OnSettingsSaved(structs.Map(c))
	}
}

func OnNewUser(u *User) {
	for _, p := range AllPlugins {
		p.OnNewUser(structs.Map(u))
	}
}

func OnNewService(s *Service) {
	for _, p := range AllPlugins {
		p.OnNewService(structs.Map(s))
	}
}

func OnDeletedService(s *Service) {
	for _, p := range AllPlugins {
		p.OnDeletedService(structs.Map(s))
	}
}

func OnUpdateService(s *Service) {
	for _, p := range AllPlugins {
		p.OnUpdatedService(structs.Map(s))
	}
}

func SelectPlugin(name string) plugin.PluginActions {
	for _, p := range AllPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return plugin.PluginInfo{}
}
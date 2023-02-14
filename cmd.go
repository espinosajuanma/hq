package hq

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/espinosajuanma/hq/types"
	S "github.com/espinosajuanma/slingr-go"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"github.com/rwxrob/vars"
)

var app = S.NewApp("hq", S.EnvProd)

func init() {
	Z.Vars.SoftInit()
	Z.Conf.SoftInit()
	token := Z.Vars.Get(".token")
	if token != "" {
		app.Token = token
	}
}

var Cmd = &Z.Cmd{
	Name: `hq`,
	Commands: []*Z.Cmd{
		help.Cmd,
		LoginCmd, LogoutCmd,
		aliveCmd, currentCmd,
		TimeTrackingCmd,
		/*
			PlatformTicketCmd,
			PlatformReleaseCmd,
			FeedbackCmd,
		*/
		vars.Cmd, conf.Cmd,
	},
	Shortcuts: Z.ArgMap{
		"email": {"var", "set", "email"},
	},
	Version:     `v0.0.1`,
	Source:      `https://github.com/espinosajuanma/hq`,
	Issues:      `https://github.com/espinosajuanma/hq/issues`,
	Summary:     `Use HQ slingr app from the console`,
	Description: `CLI tool to handle some popular features of HQ slingr app`,
}

var LoginCmd = &Z.Cmd{
	Name:        `login`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     `Login to the HQ app`,
	Description: ``,
	Usage:       `<password>`,
	Call: func(x *Z.Cmd, args ...string) error {
		email, _ := x.Caller.Get("email")
		if email == "" {
			email := term.Prompt("Email: ")
			if email == "" {
				return fmt.Errorf("email can't be empty")
			}
		}
		pass := ""
		if len(args) > 0 {
			pass = args[0]
		}
		if pass == "" {
			pass = term.PromptHidden("Password: ")
		}
		if pass == "" {
			return fmt.Errorf("password can't be empty")
		}

		r, err := app.Login(email, pass)
		if err != nil {
			return err
		}

		x.Caller.Set("email", r.UserEmail)
		x.Caller.Set("token", r.Token)
		x.Caller.Set("userName", r.UserName)
		x.Caller.Set("userId", r.UserID)

		term.Printf("Logged in [%s-%s]", app.Name, string(app.Env))
		return nil
	},
}

var LogoutCmd = &Z.Cmd{
	Name:        `logout`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     `Logs out the HQ app`,
	Description: ``,
	Call: func(x *Z.Cmd, args ...string) error {
		_, err := app.Logout()
		if err != nil {
			return err
		}
		x.Caller.Set("token", "")
		x.Caller.Set("userName", "")
		x.Caller.Set("userId", "")
		term.Printf("Logged out [%s-%s]", app.Name, string(app.Env))
		return nil
	},
}

var currentCmd = &Z.Cmd{
	Name:        `current`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     `Get current user information`,
	Description: ``,
	Call: func(x *Z.Cmd, args ...string) error {
		r, err := app.Get("/users/current", nil)
		if err != nil {
			return err
		}
		var current types.CurrentUser
		err = json.Unmarshal(r, &current)
		if err != nil {
			return err
		}
		term.Printf("Email: %s\nName: %s\nDeveloper: %t\n", current.Email, current.FullName, current.Permissions.Developer)
		return nil
	},
}

var aliveCmd = &Z.Cmd{
	Name:        `alive`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     `Keeps token alive making request every 30 minutes`,
	Description: ``,
	Call: func(x *Z.Cmd, args ...string) error {
		for {
			term.Printf("Making request to [%s-%s] to keep token alive", app.Name, string(app.Env))
			_, err := app.Get("/users/current", nil)
			if err != nil {
				return err
			}
			time.Sleep(1800 * time.Second)
		}
	},
}

func checkToken() bool {
	_, err := app.Get("/users/current", nil)
	return err == nil
}

package hq

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/espinosajuanma/hq/types"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var TimeTrackingCmd = &Z.Cmd{
	Name:    `timetracking`,
	Aliases: []string{"time"},
	Commands: []*Z.Cmd{
		listTimeTrackingCmd, addTimeTrackingCmd,
		setTimeTrackingDefaultsCmd,
		help.Cmd,
	},
	Summary:     `manage time tracking entries`,
	Description: ``,
}

var addTimeTrackingCmd = &Z.Cmd{
	Name:        `add`,
	Commands:    []*Z.Cmd{help.Cmd},
	Aliases:     []string{"create"},
	Summary:     ``,
	Description: ``,
	Usage:       `<note> <duration> <project> <date>`,
	Call: func(x *Z.Cmd, args ...string) error {
		var err error

		userId, err := x.Root().Get("userId")
		if err != nil {
			return err
		}

		// Argument 0 - Note
		note := ""
		if len(args) >= 1 {
			note = args[0]
		}

		// Argument 1 - Duration
		timeSpent, _ := time.ParseDuration("8h")
		if len(args) >= 2 {
			timeSpent, err = time.ParseDuration(args[1])
			if err != nil {
				return err
			}
		}

		// Argument 3 - Project
		projectId := ""
		var selectedProject types.FrontendProject
		if len(args) >= 3 {
			projectLabel := args[2]
			if projectLabel != "" {
				query := map[string]string{
					"people.user": userId,
					"name":        projectLabel,
					"status":      "development,maintenance",
					"_fields":     "people",
				}
				r, err := app.GetRecords(types.FRONTEND_PROJECTS_ENTITY, query)
				if err != nil {
					return err
				}
				var projects types.ManyFrontendProjects
				err = json.Unmarshal(r, &projects)
				if err != nil {
					return err
				}
				if projects.Total > 0 {
					selectedProject = projects.Items[0]
					projectId = selectedProject.Id
				}
			}
		}

		// Argument 4 - Date
		date := time.Now()
		if len(args) >= 4 {
			date, err = time.Parse("2006-01-02", args[3])
			if err != nil {
				return err
			}
		}
		isWorkingDay := date.Weekday() != 0 && date.Weekday() != 6
		if !isWorkingDay {
			return fmt.Errorf("%s, is not a business day", date.Format("2006-01-02"))
		}

		// Check if date is a holiday
		country, err := x.Root().Get("country")
		if err != nil {
			return err
		}
		if country == "" {
			country = "Argentina"
		}
		query := map[string]string{
			"day":     date.Format("2006-01-02"),
			"country": country,
		}
		r, err := app.GetRecords(types.HOLIDAY_ENTITY, query)
		if err != nil {
			return err
		}
		var holidays types.ManyHolidays
		err = json.Unmarshal(r, &holidays)
		if err != nil {
			return err
		}
		if holidays.Total > 0 {
			return fmt.Errorf("date [%s] is a holiday in [%s]", date.Format("2006-01-02"), country)
		}

		// Check if date logged more than 8 hours
		var sum int64 = timeSpent.Milliseconds()
		aggregate := types.Aggregate{
			{
				Match: struct {
					Day    string `json:"day"`
					Person string `json:"person"`
				}{
					Day:    date.Format("2006-01-02"),
					Person: userId,
				},
			},
			{
				Group: struct {
					By        string `json:"by"`
					TimeSpent string `json:"timeSpent"`
				}{
					By:        "day",
					TimeSpent: "sum(timeSpent)",
				},
			},
		}
		r, err = app.Aggregate(types.TIME_TRACKING_ENTITY, aggregate)
		if err != nil {
			return err
		}
		var tts types.ManyTimeTracking
		err = json.Unmarshal(r, &tts)
		if err != nil {
			return err
		}
		if tts.Total >= 1 {
			sum += tts.Items[0].TimeSpent
		}
		d, _ := time.ParseDuration("8h")
		if sum > d.Milliseconds() {
			a, _ := time.ParseDuration(fmt.Sprint(sum) + "ms")
			return fmt.Errorf("you can't log %v in a day", a.String())
		}

		// Select project
		if projectId == "" {
			query = map[string]string{
				"people.user": userId,
				"status":      "development,maintenance",
				"_fields":     "people",
			}
			r, err = app.GetRecords(types.FRONTEND_PROJECTS_ENTITY, query)
			if err != nil {
				return err
			}
			var projects types.ManyFrontendProjects
			err = json.Unmarshal(r, &projects)
			if err != nil {
				return err
			}

			if projects.Total == 0 {
				return fmt.Errorf("couldn't find any active project for you")
			} else if projects.Total == 1 {
				selectedProject = projects.Items[0]
			} else {
				term.Print("Select project: ")
				for i, p := range projects.Items {
					term.Printf("%v. %s", i+1, p.Label)
				}
				i := term.Prompt("Select project: ")
				cI, err := strconv.Atoi(i)
				if err != nil {
					return err
				}
				selectedProject = projects.Items[cI-1]
			}
		}
		projectId = selectedProject.Id

		// Select sow
		sowId := ""
		for _, p := range selectedProject.People {
			if p.User.Id == userId {
				sowId = p.Sow.Id
				break
			}
		}

		var sow types.FrontendSOW
		defaultServiceId := ""
		if sowId != "" {
			r, err := app.GetRecord(types.FRONTEND_SOW_ENTITY, sowId, map[string]string{})
			if err != nil {
				return err
			}
			err = json.Unmarshal(r, &sow)
			if err != nil {
				return err
			}
			for _, v := range sow.People {
				if v.User.Id == userId {
					defaultServiceId = v.DefaultService.Id
					break
				}
			}
		}

		ms := timeSpent.Milliseconds()
		payload := &types.TimeTrackingPayload{
			TimeSpent: ms,
			Notes:     note,
			Day:       date.Format("2006-01-02"),
			Project:   projectId,
			Person:    userId,
			Billable:  true,
		}

		if sowId != "" {
			payload.Sow = sowId
		}
		if defaultServiceId != "" {
			payload.Service = defaultServiceId
		}
		e, err := app.CreateRecord(types.TIME_TRACKING_ENTITY, payload)
		if err != nil {
			return err
		}
		var tt types.TimeTracking
		err = json.Unmarshal(e, &tt)
		if err != nil {
			return err
		}
		term.Print(printTimeTracking(&tt))
		return nil
	},
}

var listTimeTrackingCmd = &Z.Cmd{
	Name:        `list`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     `show list of time tracking entries`,
	Description: `prints a list of last time tracking entries`,
	Call: func(x *Z.Cmd, args ...string) error {
		query := map[string]string{
			"_sortField": "day",
			"_sortType":  "desc",
		}
		r, err := app.GetRecords(types.TIME_TRACKING_ENTITY, query)
		if err != nil {
			return err
		}
		var tts types.ManyTimeTracking
		err = json.Unmarshal(r, &tts)
		if err != nil {
			return err
		}
		for _, tt := range tts.Items {
			term.Print(printTimeTracking(&tt))
		}
		return nil
	},
}

var setTimeTrackingDefaultsCmd = &Z.Cmd{
	Name:        `setup`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     ``,
	Description: ``,
	Call: func(x *Z.Cmd, args ...string) error {
		project, _ := x.Caller.Get("project")
		customer, _ := x.Caller.Get("customer")
		if project == "" || customer == "" {
			// Find projects. Select one of the list
			var manyProjects types.ManyFrontendProjects
			r, err := app.GetRecords("frontend.projects", nil)
			if err != nil {
				return err
			}
			err = json.Unmarshal(r, &manyProjects)
			if err != nil {
				return err
			}
			for i, p := range manyProjects.Items {
				term.Printf("[%d] ", i+1, p.Label)
			}
			selected, err := strconv.Atoi(term.Prompt("Select default project: "))
			if err != nil {
				return err
			}
			if len(manyProjects.Items) < selected || selected < 1 {
				return fmt.Errorf("there is no project with index [%d]", selected)
			}
			project = manyProjects.Items[selected-1].Id
			x.Caller.Set("project", project)
			customer = manyProjects.Items[selected-1].Id
			x.Caller.Set("customer", customer)
			userId, err := x.Root().Get("userId")
			if err != nil {
				return err
			}
			for _, p := range manyProjects.Items[selected-1].People {
				if p.Id == userId {
					sow := p.Sow.Id
					// Get frontendBilling.sows
					r, err = app.GetRecord("frontendBilling.sows", sow, nil)
					if err != nil {
						return err
					}
					var sowRecord types.FrontendSOW
					err = json.Unmarshal(r, &sowRecord)
					if err != nil {
						return err
					}
					for _, s := range sowRecord.People {
						if s.Id == userId {
							x.Caller.Set("defaultService", s.DefaultService.Id)
							break
						}
					}
					break
				}
			}
		}
		return nil
	},
}

func printTimeTracking(t *types.TimeTracking) string {
	date, _ := time.Parse("2006-01-02", t.Day)
	c := term.HGreen
	if !t.Billable {
		c = term.HYellow
	}
	return fmt.Sprintf("[%s] %s[%s]%s [%s] %s", t.Project.Label, c, t.Day, term.Reset, date.Weekday().String()[0:3], t.Notes)
}

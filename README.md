# Go `hq` Command Line

Simple CLI tool to manage some functionalities of slingr HQ app.

## Install

You can just grab the latest binary [release](https://github.com/espinosajuanma/hq/releases).

This command can be installed as a standalone program or composed into a Bonzai command tree.

Standalone

```bash
go install github.com/espinosajuanma/hq/cmd/hq@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/espinosajuanma/hq"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{hq.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C hq hq
```

## Usage

### Initial configuration

- `hq email <email>` - Set your email

### Time Tracking

- `hq time list` - List latest time tracking entries
	- Aliases: `ls`
- `hq time add` - Add a time tracking entry

## Road map - To do list

### Minimum viable product (MVP)

- [x] Login command: `hq login`
	- Prompt password secret
	- Login to HQ app
	- Handle token cache storage
- [x] Logout command: `hq logout`
	- Logout request and clear cached token
- [x] Keep alive command: `hq alive`
	- Make calls every `~30m` to keep token alive
- [x] Time tracking list command: `hq time list`
- [x] Time tracking add command: `hq time add`
	- [x] Be able to create a time tracking entry
	- [x] Avoid logging on weekends
	- [x] Get projects
		- [x] Error if there is none
		- [x] Prompt to select one if there is `> 1`
		- [x] Select one if there is only one
	- [x] Usage: `hq time add <note> <time> <date>`
		- `<note>` default: empty
		- `<time>` default: `8h`
		- `<date>` default: `today`

### Time tracking improvements and Platform Ticket

- [x] Time Tracking avoid log hours on limit
	- Already hours logged + current shouldn't be more than 8
- [ ] Time tracking select default project
	- If there is more than one project be able to set as default
	- [ ] Command `hq time defaults`
- [ ] Platform Ticket list command: `hq ticket list`
- [ ] Platform Ticket add command: `hq ticket add`
	- Usage `hq ticket add <title> <priority> <type>`
	- Prompt title if is not provided
	- Prompt priority if is not provided
		- Default `p3`
	- Prompt type if is not provided
		- Default `bug`
	- Open `$EDITOR`
	- Prompt confirm and submit ticket
	- Transform markdown to html and POST ticket

### Platform Release command

- [ ] Platform release list command: `hq release list`
- [ ] Platform release get command: `hq release get <number>`
	- Should parse HTML to markdown
	- Open `$EDITOR`

### Profile command

This command would be just to get current user information
	- Available holidays
	- Manager
	- Project Managers `[]`
	- Days until birthday
	- Days until slingr birthday
	- Seniority
	- Skills `[]`
	- Projects `[]`
	- Office - Manager
	- Etc, check `/sys.users/{id}`

### Feedback

- [ ] Feedback list command
- [ ] Feedback give command
	- Would be nice to do alias for persons and office
	- Open `$EDITOR`
	- Prompt confirm and submit feedback
	- Transform markdown to html and POST feedback

## Skills

- `hq skill list`
- `hq skill add`
	- Prompt search skill
	- Prompt select skill
	- Prompt select level
	- Add skill to profile

### Time tracking automatic notes and stories

If you have an external command to get your current stories, it may be
worth it to set that to the note.

Additionally for projects that use stories from HQ that could be doable
with this same tool. So we could have both options:

- `hq stories current`
	- Would be worth it to add `hq stories list` as well

### Time tracking import hours from tempo

Not a huge priority. But people that use tempo and has `CSV` files, it
could be cool if they can use the file to import them. This would need
an update to the `slingr` sdk to be able to upload files.

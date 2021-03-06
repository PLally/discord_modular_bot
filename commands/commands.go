package commands

import (
	"github.com/plally/FoxBot/commands/middleware"
	"github.com/plally/FoxBot/commands/permissions"
	"github.com/plally/FoxBot/commands/random"
	"github.com/plally/FoxBot/commands/subscriptions"
	"github.com/plally/dgcommand"
)

func CommandGroup() *dgcommand.CommandGroup {
	var CommandGroup = dgcommand.Group()

	CommandGroup.AddHandler("random", random.CommandGroup())

	CommandGroup.Command("ping", ping)
	CommandGroup.Command("e621 [tags...]", e621Func).Use(middleware.RequireNSFW()).
		Desc("A random picture from e621")
	CommandGroup.Command("e926 [tags...]", e926Func).
		Desc("A random picture from e926")
	CommandGroup.Command("info <object>", objInfoFunc).
		Desc("Gets info about the given discord object")
	CommandGroup.Command("help [command...]", helpCommand).
		Desc("shows the help message")

	CommandGroup.Command("findsource", findSourceCommand).
		Desc("Finds the source of an image")
	CommandGroup.AddHandler("sub", subscriptions.CommandGroup())
	CommandGroup.AddHandler("perms", permissions.CommandGroup())
	return CommandGroup
}

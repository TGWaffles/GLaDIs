---
sidebar_position: 2
---

# Quirks

Dapper-go has been designed with ease of use but there are some strange quirks either due to the Discord API or Go itself. These choices have been made to make the library more lightweight and easier to use.

## Inline Literal Pointers

When using one of the many structures provided in the `discord` package, you may want to use a literal value where there may be a pointer. This is due to the reuse of structures when reading from the Discord API and so we can have nil values.

The package `helper` contains a utility function which takes a value of any type and returns a pointer to that value.

```go
helpers.Ptr(“My String”) // Returns a *string
helpers.Ptr(false) // Returns a *bool
```

The following example shows how you would define an emoji using a literal value.

```go
emoji := discord.Emoji{
	Name: helpers.Ptr(“🐊”),
}
```

## Action Rows

As action rows are technically classed as a message component, they must be added to a message component slice. This is so that they can be parsed and sent to the Discord API correctly.

You may see the following code:

```go
func Handler(itx *discord.Interaction) {
	err := itx.EditResponse(discord.EditResponse{
		Components: []discord.Component{
			&discord.ActionRow{
				Components: []discord.Component{
					&discord.Button{
						Label: “Click Me”,
						Style: discord.Primary,
						CustomID: “click_me”,
					},
				},
			},
		},
	})
}
```

Another helper is included to shorten the creation of action rows. This is the `helpers.CreateActionRow()` function.

```go
func Handler(itx *discord.Interaction) {
	err := itx.EditResponse(discord.EditResponse{
		Components: helpers.CreateActionRow(
			&discord.Button{
				Label: “Click Me”,
				Style: discord.Primary,
				CustomID: “click_me”,
			}
		),
	})
}
```

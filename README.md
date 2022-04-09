# Drie op een rij

Drie op een rij, of boter, kaas en eieren.

Kunnen we hier een AI voor maken? Op een computer moet dit zeker kunnen omdat we de computer
gewoon alle mogelijke zetten kunnen laten doen om te bepalen wat een "beste" zet is.

Een voorbeeld, waarbij X gewonnen heeft:

X..
.O.
XXX

# Development

The development is supported by a linter which you can execute with:

`golangci-lint run`

When running the above on the command line it will check the code for certain rules and either auto-fix or suggest where to fix issues.

On my development system I've also installed `pre-commit` which will automatically run the linter before a git commit. Thereby preventing commits when the linter fails. This just ensures a some minimal code quality is maintained on the project.


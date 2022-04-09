# Drie op een rij

Drie op een rij, of boter, kaas en eieren.

Can we create an AI, or computerplayer for this game? A computer should be able
to simply try all moves till the game ends to determine what move is the best
to play out in the game.

In this example we represent the board with the 9 spaces in memory as follows

//     |     |
//  1  |  2  |  3
//_____|_____|_____
//     |     |
//  4  |  5  |  6
//_____|_____|_____
//     |     |
//  7  |  8  |  9
//     |     |

# Development

We have chosen go as our programming language.

## GoLang liner

The development in `go` is supported by a linter which you can execute with:

`golangci-lint run`

When running the above on the command line the linter will check the code for certain rules and either auto-fix or suggest where to fix issues. The linter is an additional check on style, readability above the default go compiler checks.

## pre-commit

On my development system I've also installed `pre-commit` which will automatically run the linter before a git commit. Thereby preventing commits when the linter fails. This just ensures a some minimal code quality is maintained on the project.

On a windows system, the pre-commit can be enabled as a git-hook. See more info on the net how to enable the `pre-commit` by enabling it as a githook below your `.git/hooks` folder when cloning the repository. Having the `pre-commit` enabled will prevent pushing code that does not abide by the repository checks.

## Model package

The /internal/model represents the board that consists of 9 spaces. See the comments in code for documentation.

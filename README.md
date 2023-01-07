# TipTap

Typing Test TUI 

This is a simple terminal-based typing test.

It's also mainly just an excuse to write something using 
[Bubble Tea](https://github.com/charmbracelet/bubbletea), 
since it's been catching my eye for a bit. 

## Controls 

Menu navigation is done using either `arrow keys` or `j,k` to move the cursor. 

The `enter` key or `spacebar`, can be used for selecting the currently highlighted option. 

The `escape` key (or also `q` in most cases) will return you to the previous screen, 
or exit the program if you are on the main menu.

`Ctrl+C` will exit the program immediately

## Settings 

Values that can be customized are saved in `$XDG_CONFIG_HOME/tiptap/tiptap-settings.json`. 
If that location is not defined in the runtime, settings are saved to `~/.config/tiptap/tiptap-settings.json`. 

Available settings are: 

- `wordcount` - The number of words to include in the typing test 
- `style.background` - Color of the default app background
- `style.errorbackground` - Color of the background of an incorrectly typed character 
- `style.cursor` - Color of the cursor when typing, and the selected item in menus
- `style.text` - Color of general text, and untyped characters in the typing test 
- `style.correct` - Color of text that was entered correctly in the test 
- `style.error` - Color of text that was entered incorrectly in the test 

If no settings file exists, one is created at startup. 

## WPM Calculations

- Elapsed Time - The number of seconds that passed between the typing screen being shown to the user, and the user typing the last character on screen. 
- Accuracy - `Correct characters / Total characters`
- Raw WPM - `Total characters / 5 / Elapsed Time in Minutes`
- Adjusted WPM - `Raw WPM * Accuracy`


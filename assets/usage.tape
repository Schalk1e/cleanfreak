# Set output
Output usage.gif

# Set up a 1200x600 terminal with 46px font.
Set TypingSpeed 250ms
Set FontSize 22
Set Width 1200
Set Height 950

# Use zsh
Set Shell zsh
Sleep 1s
Type "cf"
Enter
Sleep 500ms

# Diagnose system
Type "cf diagnose"
Enter
Sleep 1s

# Action clean on downloads
Type "cf clean downloads"
Enter
Sleep 1.5s
Down@500ms 1
Enter
Sleep 1s

# Diagnose system again
Type "cf diagnose"
Enter
Sleep 1s


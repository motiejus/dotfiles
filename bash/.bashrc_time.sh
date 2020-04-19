# from https://redandblack.io/blog/2020/bash-prompt-with-updating-time/
DIRECTORY="\w"
DOUBLE_SPACE="  "
NEWLINE="\n"
NO_COLOUR="\e[00m"
PRINTING_OFF="\["
PRINTING_ON="\]"
PROMPT_COLOUR="\e[0;33m"
PS1_PROMPT="\$"
PS2_PROMPT=">"
RESTORE_CURSOR_POSITION="\e[u"
SAVE_CURSOR_POSITION="\e[s"
SINGLE_SPACE=" "
TIMESTAMP="\A"
TIMESTAMP_PLACEHOLDER="--:--"

move_cursor_to_start_of_ps1() {
    command_rows=$(history 1 | wc -l)
    if [ "$command_rows" -gt 1 ]; then
        let vertical_movement=$command_rows+1
    else
        command=$(history 1 | sed 's/^\s*[0-9]*\s*//')
        command_length=${#command}
        ps1_prompt_length=${#PS1_PROMPT}
        let total_length=$command_length+$ps1_prompt_length
        let lines=$total_length/${COLUMNS}+1
        let vertical_movement=$lines+1
    fi
    tput cuu $vertical_movement
}

PS0_ELEMENTS=(
    "$SAVE_CURSOR_POSITION" "\$(move_cursor_to_start_of_ps1)"
    "$PROMPT_COLOUR" "$TIMESTAMP" "$NO_COLOUR" "$RESTORE_CURSOR_POSITION"
)
PS0=$(IFS=; echo "${PS0_ELEMENTS[*]}")

PS1_ELEMENTS=(
    # Empty line after last command.
    #"$NEWLINE"
    # First line of prompt.
    "$PRINTING_OFF" "$PROMPT_COLOUR" "$PRINTING_ON"
    "$TIMESTAMP_PLACEHOLDER" "$DOUBLE_SPACE" "$DIRECTORY" "$PRINTING_OFF"
    "$NO_COLOUR" "$PRINTING_ON" "$NEWLINE"
    # Second line of prompt.
    "$PRINTING_OFF" "$PROMPT_COLOUR" "$PRINTING_ON" "$PS1_PROMPT"
    "$SINGLE_SPACE" "$PRINTING_OFF" "$NO_COLOUR" "$PRINTING_ON"
)
PS1=$(IFS=; echo "${PS1_ELEMENTS[*]}")

PS2_ELEMENTS=(
    "$PRINTING_OFF" "$PROMPT_COLOUR" "$PRINTING_ON" "$PS2_PROMPT"
    "$SINGLE_SPACE" "$PRINTING_OFF" "$NO_COLOUR" "$PRINTING_ON"
)
PS2=$(IFS=; echo "${PS2_ELEMENTS[*]}")

shopt -s histverify

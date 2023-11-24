#! /bin/sh

width="$1"
icon_color="$2"

format="+ %I:%M %p • %a %d"

[ "$width" -lt 100 ] && format="+ %I:%M %p"

time=$(date "$format")

printf "#[fg=terminal,bg=terminal]%s #[fg=black,bg=%s] 󰥔 #[fg=white,bg=terminal]" "$time" "$icon_color"

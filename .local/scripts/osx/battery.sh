#!/bin/sh

percentage=$(system_profiler SPPowerDataType | grep -i "state of charge" | awk '{print $(NF)}')

low_battery=$([ $percentage -le 20 ] && echo yes || echo no)
plugged=$(system_profiler SPPowerDataType | grep -i "connected: no" && echo no || echo yes)

if [ "$plugged" == "yes" ]; then
	echo "#[fg=green,bg=terminal]  $percentage%"
elif [ "$low_battery" == "no" ]; then

	if [ "$percentage" -ge "90" ]; then
		echo "#[fg=green,bg=terminal] $percentage%"
	elif [ "$percentage" -ge "70" ]; then
		echo "#[fg=green,bg=terminal] $percentage%"
	elif [ "$percentage" -ge "50" ]; then
		echo "#[fg=yellow,bg=terminal] $percentage%"
	elif [ "$percentage" -ge "30" ]; then
		echo "#[fg=yellow,bg=terminal] $percentage%"
	elif [ "$percentage" -ge "20" ]; then
		echo "#[fg=orange,bg=terminal] $percentage%"
	fi

else
	echo "#[fg=red,bg=terminal] $percentage%"
fi

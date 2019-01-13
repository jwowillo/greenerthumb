# deactivate.sh removes function aliases to greenerthumb programs.

ALIASES=`declare -F`

while read -r line
do
	if [[ $line == *"greenerthumb"* ]]
	then
		alias="$(echo $line | cut -d' ' -f3)"
		unset -f "$alias"
	fi
done <<< "$ALIASES"

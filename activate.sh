# activate.sh makes function aliases to greenerthumb programs.

ALIASES=$(./find-aliases)

while read -r line
do
	name="$(echo $line | cut -d',' -f1)"
	path="$(echo $line | cut -d',' -f2)"
	eval "$name() { $path \"\$@\"; }"
	export -f "$name"
done <<< "$ALIASES"


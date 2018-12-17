# `log` Design

`log` messages from STDIN to a file (4).

This just copies each line of STDIN to a file. Logs have the time appended to
the file-name and are rotated each day.

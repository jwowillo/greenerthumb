# `log` Design

`log` messages from STDIN to a file.

This just copies each line of STDIN to a file. Logs have the time appended to
the file-name and are rotated each day.

all: design.pdf

design.pdf: design.md fan.png
	pandoc design.md --latex-engine xelatex -o design.pdf

fan.png: fan.xml

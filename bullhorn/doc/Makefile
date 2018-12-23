all: design.pdf

design.pdf: design.md sequence.png
	pandoc design.md --latex-engine xelatex -o design.pdf

sequence.png: sequence.uml
	java -jar plantuml.jar sequence.uml

all:
	java -jar plantuml.jar bullhorn.txt
	java -jar plantuml.jar log.txt
	java -jar plantuml.jar sense.txt
	java -jar plantuml.jar plot.txt
	pandoc requirements.md --latex-engine xelatex -o requirements.pdf
	pandoc design.md --latex-engine xelatex -o design.pdf
	pandoc icd.md --latex-engine xelatex -o icd.pdf

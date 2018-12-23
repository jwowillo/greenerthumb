all: design.pdf

design.pdf: design.md air.png soil.png
	pandoc design.md --latex-engine xelatex -o design.pdf

air.png:
	air.fzz

soil.png:
	soil.fzz

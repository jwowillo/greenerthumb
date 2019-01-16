all: design.pdf

design.pdf: design.md air-sensor.png soil-sensor.png waterer.png
	pandoc -f markdown-implicit_figures design.md --latex-engine xelatex \
		-o design.pdf

air-sensor.png:
	air-sensor.fzz

soil-sensor.png:
	soil-sensor.fzz

waterer.png:
	waterer.fzz

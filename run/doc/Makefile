all: design.pdf

design.pdf: design.md air-sensor.xml soil-sensor.xml \
	logger-sensor.xml plotter-sensor.xml
	pandoc -f markdown-implicit_figures design.md --latex-engine xelatex \
		-o design.pdf

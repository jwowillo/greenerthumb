all: design.pdf

design.pdf: design.md air.xml soil.xml logger.xml plotter.xml
	pandoc -f markdown-implicit_figures design.md --latex-engine xelatex \
		-o design.pdf

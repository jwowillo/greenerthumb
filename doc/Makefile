all: requirements.pdf design.pdf icd.pdf
	pdfunite design.pdf \
		../bullhorn/doc/design.pdf \
		../fan/doc/design.pdf \
		../log/doc/design.pdf \
		../message/doc/design.pdf \
		../plot/doc/design.pdf \
		../process/doc/design.pdf \
		../sense/doc/design.pdf \
		../store/doc/design.pdf \
		../disclosure/doc/design.pdf \
		../app/doc/design.pdf \
		merged.pdf
	mv merged.pdf design.pdf

requirements.pdf: requirements.md
	pandoc requirements.md --latex-engine xelatex -o requirements.pdf

design.pdf: design.md bullhorn fan log message plot process sense store \
	disclosure app
	pandoc -f markdown-implicit_figures design.md --latex-engine xelatex \
		-o design.pdf

icd.pdf: icd.md
	pandoc icd.md --latex-engine xelatex -o icd.pdf

%:
	$(MAKE) -C ../$*/doc

all: bullhorn fan log message plot process sense
	pandoc requirements.md --latex-engine xelatex -o requirements.pdf
	pandoc -f markdown-implicit_figures design.md --latex-engine xelatex \
		-o design.pdf
	pandoc icd.md --latex-engine xelatex -o icd.pdf
	pdfunite design.pdf \
		../bullhorn/doc/design.pdf \
		../fan/doc/design.pdf \
		../log/doc/design.pdf \
		../message/doc/design.pdf \
		../plot/doc/design.pdf \
		../process/doc/design.pdf \
		../sense/doc/design.pdf \
		merged.pdf
	mv merged.pdf design.pdf

%:
	$(MAKE) -C ../$*/doc

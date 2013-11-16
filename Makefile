all: cpp go
cpp:
	cat enginemodels.cy editormodels.cy > models.cy
	cyborgbear -i models.cy -o models -lc
	rm -f models.cy
go:
	cat enginemodels.cy editormodels.cy > models.cy
	cyborgbear -i models.cy -o models.go -t go
	rm -f models.cy
clean:
	rm -f models.*

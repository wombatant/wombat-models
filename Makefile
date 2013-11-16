all: cpp editorcpp go
editorcpp:
	cat enginemodels.cy editormodels.cy > models.cy
	cyborgbear -i models.cy -o models -lc
cpp:
	cyborgbear -i enginemodels.cy -o enginemodels -lc
	rm -f models.cy
go:
	cyborgbear -i enginemodels.cy -o models.go -t go
	rm -f models.cy
clean:
	rm -f models.*

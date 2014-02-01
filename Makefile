all: enginemodels editormodels go liccor
editormodels:
	cat enginemodels.cy editormodels.cy > models.cy
	cyborgbear -i models.cy -o editormodels -t cpp-qt
	rm -f models.cy
enginemodels:
	cyborgbear -i enginemodels.cy -o enginemodels -lc
go:
	cyborgbear -i enginemodels.cy -o models.go -t go
	go build *.go
liccor:
	liccor
clean:
	rm -f models.cy enginemodels.hpp enginemodels.cpp editormodels.hpp editormodels.cpp

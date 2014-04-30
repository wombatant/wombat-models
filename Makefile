all: enginemodels editormodels go liccor
editormodels:
	cat enginemodels.cy editormodels.cy > models.cy
	cyborgbear -i models.cy -o editormodels -t cpp-qt
	rm -f models.cy
enginemodels:
	cat enginemodels.cy enginesettings.cy > models.cy
	cyborgbear -i models.cy -o enginemodels
	rm -f models.cy
go:
	cyborgbear -i enginemodels.cy -o models.go -t go
	go build *.go
liccor: enginemodels editormodels go
	liccor
clean:
	rm -f models.cy enginemodels.hpp enginemodels.cpp editormodels.hpp editormodels.cpp

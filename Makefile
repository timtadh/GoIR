build: clean
	gobuild -a -I "../goast"

test: build
	-gobuild -t -run -I "../goast"
	rm _testmain _testmain.6 _testmain.go

fmt: clean
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w .

.PHONY : clean
clean :
	-rm -r main
	-rm _testmain _testmain.6 _testmain.go
	-find . -name "*.6" | xargs -I"%s" rm %s
	-find . -name "*.a" | xargs -I"%s" rm %s


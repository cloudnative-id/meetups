update:
	git submodule foreach git pull origin master

add:
	git submodule add $(url) $(module)

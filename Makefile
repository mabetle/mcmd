
install-mcmds: install-mfind install-mrepl

install-mfind:
	cd $(GOCODES)/src/github.com/mabetle/mcmd/mfind_task/ && \
		go install 

install-mrepl:
	cd $(GOCODES)/src/github.com/mabetle/mcmd/mrepl_task/ && \
		go install 


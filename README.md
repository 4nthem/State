# State
*A web app that manages teacher/student interaction through grades.*


### Installation and launching App
	1. Make sure you have gin installed (go get github.com/vendor/codegangsta/gin)
	2. Run the command 'gin'
	3. Livereload will now happen on port 3000

### Testing webserver:
	1. Run command 'curl http://localhost:3000/users'

### To inject a new dependency:
	1. Run 'go get PACKAGENAME'
	2. cd into 'State' directory
	3. Run 'godep save ./...'

### Logging
	- seelog.xml holds the configuration for our logging patterns
	- All messages are printed to the command line and written to all.log
	- Errors are filtered as well to errors.log


### Testing
	*Make sure you run the command "go get github.com/smartystreets/goconvey" first*
	We are using the GoConvey continuous testing framework
	1. In a different terminal tab run "goconvey" (assuming you have go/bin in your PATH variable)
	2. It will auto open a tab showing the state of the tests. This will rerun the tests every time a file is saved.
	
	- There is an option for it to send desktop notifications on every save



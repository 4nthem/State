# State
*A web app that manages teacher/student interaction through grades.*


### Installation and launching App
	1. Make sure you have gin installed (go get github.com/codegangsta/gin)
	2. Run the command 'gin'
	3. Livereload will now happen on port 3000

### Testing webserver:
	1. Run command 'curl http://localhost:3000/users'
	2. Or run in your webbrowser/postman

### To inject a new dependency:
	1. Run 'go get PACKAGENAME'
	2. cd into 'State' directory
	3. Run 'godep save ./...'

### Logging
	- seelog.xml holds the configuration for our logging patterns
	- All messages are printed to the command line and written to all.log
	- Errors are filtered as well to errors.log


### Testing: [Intro to GoConvey](https://www.youtube.com/watch?v=wlUKRxWEELU)
	*Make sure you run the command "go get github.com/smartystreets/goconvey" first*
	We are using the GoConvey continuous testing framework
	1. In a different terminal tab run "goconvey" (assuming you have go/bin in your PATH variable)
	2. It will auto open a tab showing the state of the tests. This will rerun the tests every time a file is saved.
	
	- There is an option for it to send desktop notifications on every save


### Mongo DB Testing:
	Run the following:
	1. gin
	*[gin] listening on port 3000*

		To GET a user in the db, open new terminal
		1. curl http://localhost:3000/users/563f87375872ef08010eeddc
			You will get the following error in the gin terminal:
				*2015/11/08 12:04:15 http: proxy error: dial tcp [::1]:3001: getsockopt: connection refused*
				*[martini] listening on :3001 (development)*
			Please refer to: https://github.com/gin-gonic/gin/issues/159 
		Run the command again!
		2. curl http://localhost:3000/users/563f87375872ef08010eeddc
			*{"data":{"id":"563f87375872ef08010eeddc","name":"Erik Muro","email":"erik@email.com"},"message":"Successfully found user"}*

		In the gin terminal:
			*[martini] Started GET /users/563f87375872ef08010eeddc for ::1*
			*[martini] Completed 200 OK in 51.726269ms*

		To POST a user (Create)
		1. curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Erik Muro", "email": "erik@email.com"}' http://localhost:3000/users
			*{"data":{"id":"563f8fd45872ef0933ead2a2","name":"test","email":"test@email.com"},"message":"Successfully created user"}*

		In the gin terminal:
			*[martini] Started POST /users for ::1*
			*[martini] Completed 200 OK in 44.525796ms*


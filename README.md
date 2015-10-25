# State
A web app that manages teacher/student interaction through grades.


Installation:
	1. Make sure you are in the State directory
	2. Run the command 'go install'

Launching Application:
	1. Run the command 'State'

Testing webserver:
	1. Run command 'curl http://localhost:3000/users'

To inject a new dependency:
	1. Run 'go get PACKAGENAME'
	2. cd into 'State' directory
	3. Run 'godep save --r ./...'
	4. It is now safe to remove the package from the src and pkg directories in your GOPATH


NOTE: Server Configuration will be placed in env/config.json
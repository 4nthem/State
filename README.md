# State
A web app that manages teacher/student interaction through grades.



cd into the 'models' folder to install models package:
	'go install'

cd into the 'controllers' folder to install controllers package:
	'go install'

Install packages:

	'github.com/julienschmidt/httprouter'
	'github.com/4nthem/State/controllers'
	'github.com/4nthem/State/models'
	'gopkg.in/mgo.v2/bson'
	'gopkg.in/mgo.v2'


cd into the 'State' folder to compile app:
	'go install'


run:
	'State'

open new terminal and run:
	'curl http://localhost:3000/users'
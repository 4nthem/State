package start


import(
	log "github.com/cihub/seelog"
	DEATH "github.com/vrecan/death"
	"github.com/4nthem/State/rest"
	SYS "syscall"
	"io"
)

func Run() {
	death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)

	goRoutines := make([]io.Closer, 0)

	webserver := rest.NewRestServer()
	go webserver.StartServer()

	death.WaitForDeath(goRoutines...)
	log.Info("Finished shutting down webserver")
}
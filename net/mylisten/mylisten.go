// This package will help you bind to systemd managed FDs
//
// Big ups to @kr for writing this sick code
// 🙏

package mylisten

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"syscall"
)

// SystemdOr returns listeners for the inherited systemd fds:
// serve for serving responses (on port 443)
// and redir for redirecting to https (on port 80).
// If there are no inherited fds, it listens on addr
// and returns serve; in that case, redir is nil.
func SystemdOr(addr string) (serve, redir net.Listener, err error) {
	serve, redir, err = systemdListeners()
	if err != nil {
		return nil, nil, err
	}
	if serve != nil {
		return serve, redir, nil
	}
	serve, err = net.Listen("tcp", addr)
	return serve, nil, err
}

func systemdListeners() (serve, redir net.Listener, err error) {
	// Env vars LISTEN_FDS and LISTEN_PID are how systemd
	// tells us the number of inherited fds we have and that
	// they're meant for this process (as opposed to an
	// ancestor process that neglected to mark them as
	// close-on-exec). See also
	// https://www.freedesktop.org/software/systemd/man/sd_listen_fds.html
	pid, err := strconv.Atoi(os.Getenv("LISTEN_PID"))
	if err != nil || pid != os.Getpid() {
		return nil, nil, nil
	}
	n, err := strconv.Atoi(os.Getenv("LISTEN_FDS"))
	if err != nil {
		return nil, nil, nil
	}
	if n != 2 {
		return nil, nil, fmt.Errorf("got %d inherited fds, need 2 (port 80 and 443, in that order)", n)
	}
	os.Unsetenv("LISTEN_PID")
	os.Unsetenv("LISTEN_FDS")
	const fd = 3 // systemd always uses fd 3
	syscall.CloseOnExec(fd)
	syscall.CloseOnExec(fd + 1)
	fredir := os.NewFile(fd, fmt.Sprintf("fd%d", fd))     // port 80
	fserve := os.NewFile(fd+1, fmt.Sprintf("fd%d", fd+1)) // port 443
	redir, err = net.FileListener(fredir)
	if err != nil {
		return nil, nil, err
	}
	serve, err = net.FileListener(fserve)
	return serve, redir, err
}

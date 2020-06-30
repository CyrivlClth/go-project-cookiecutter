// package internal
package internal

func Run() error {
	s, err := setup()
	if err != nil {
		return err
	}
	return s.engine.Run(s.addrs...)
}

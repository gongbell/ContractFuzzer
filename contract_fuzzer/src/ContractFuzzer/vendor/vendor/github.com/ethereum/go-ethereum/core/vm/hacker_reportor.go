package vm
/**
Hacker_reportor.go: report the current state of the test case finishing executing just now.
              way: send http packet to the listening server
 */
type(
	make_packet func()[]byte
	make_conn   func() error
	do_report   func(data []byte)
)
func Hacker_packet() []byte {
	return  nil
}
func Hacker_conn() error{
	return nil
}
func Hacker_write(data []byte){

}
func Hacker_http_handler(f1 make_conn,f2 do_report,f3 make_packet){

}
func Hacker_report(){
	s := make(chan int,3)
	for true{
		if ok:= <-s;ok>0 {
			go Hacker_http_handler(Hacker_conn, Hacker_write, Hacker_packet)
		}
	}
}

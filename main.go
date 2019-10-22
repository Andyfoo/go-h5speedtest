package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/GeertJohan/go.rice"
)

var addr string
var wwwroot string
var isDaemon bool
var help bool
var mime map[string]string = make(map[string]string)
var isWin = os.IsPathSeparator('\\')
var randomdat []byte = make([]byte, 1024*1024)

func init() {

	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	wwwroot = p + string(os.PathSeparator) + "speedtest"
	io.ReadFull(rand.Reader, randomdat)

	flag.StringVar(&addr, "l", ":80", "binding address")
	flag.StringVar(&wwwroot, "r", wwwroot, "web root folder")
	flag.BoolVar(&isDaemon, "d", false, "daemon start")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()

}

func main() {
	if help {
		flag.Usage()
		return
	}
	if isDaemon {
		daemon()
		return
	}
	StartHttp()

}
func daemon() {

	args := os.Args[1:]
	i := 0
	for ; i < len(args); i++ {
		if args[i] == "-d=true" || args[i] == "-d" {
			args[i] = "-d=false"
			break
		}
	}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Start()
	fmt.Println("[PID]", cmd.Process.Pid, cmd.Args)
	os.Exit(0)
}
func StartHttp() {
	http.HandleFunc("/empty.php", ulHandler)
	http.HandleFunc("/garbage.php", dlHandler)
	http.HandleFunc("/getIP.php", ipHandler)
	if IsDir(wwwroot) {
		log.Println("webroot at", wwwroot)
		http.Handle("/", http.FileServer(http.Dir(wwwroot)))
	} else {
		log.Println("webroot at", "inline ./speedtest")
		http.Handle("/", http.FileServer(rice.MustFindBox("./speedtest").HTTPBox()))
	}
	log.Println("listen", addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal(http.Serve(ln, nil))
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if os.IsNotExist(err) {
		fmt.Println(err)
		return false
	}
	if f.IsDir() {
		return true
	}
	return false
}
func ulHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Add("Cache-Control", "post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	io.Copy(ioutil.Discard, r.Body)
	w.WriteHeader(200)
}

func dlHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Add("Cache-Control", "post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Content-Description", "File Transfer")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=random.dat")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.WriteHeader(200)

	var e error
	szCkSize := r.URL.Query().Get("ckSize")
	ckSize := 4
	if szCkSize != "" {
		if ckSize, e = strconv.Atoi(szCkSize); e != nil {
			ckSize = 4
		}
		if ckSize > 100 {
			ckSize = 100
		}
	}
	for ; ckSize > 0; ckSize-- {
		w.Write(randomdat)
	}
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	sip := r.Header.Get("X-FORWARDED-FOR")
	if sip == "" {
		sip = r.RemoteAddr
	}
	if !strings.Contains(sip, ":") {
		sip += ":0"
	}
	if ip, e := net.ResolveTCPAddr("tcp", sip); e == nil {
		w.Write([]byte("{\"processedString\": \"" + ip.IP.String() + "\", \"rawIspInfo\":\"\"}"))
	} else {
		w.Write([]byte(e.Error()))
	}
}

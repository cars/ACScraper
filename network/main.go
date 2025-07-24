package network

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func getDefaultLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Failed to get IP addresses: %s", err)
	}
	for _, addr := range addrs {
		fmt.Printf("Address Info: %+V", addr)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getIPByInterfaceName(interfaceName string) string {
	adapter, err := net.InterfaceByName(interfaceName)
	if err != nil {
		log.Fatalf("Failed to find interface with name %s\nErr: %s", interfaceName, err)
	}
	addresses, err2 := adapter.Addrs()
	if err2 != nil {
		log.Fatalf("Failed to get addresses for interface[%s] nErr: %s", interfaceName, err)
	}
	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getInterfaces() []string {
	ifaces, if_err := net.Interfaces()
	if if_err != nil {
		//nop
	}
	var iface_names []string
	for _, iface := range ifaces {
		fmt.Printf("Interface info: %+v\n", iface)
		iface_names = append(iface_names, iface.Name)
	}
	return iface_names
}

func runProxy(ipAddress string) {
	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.String(), "http://www.acinfinityserver.com/api/") {
			token := r.Header.Get("token")
			userAgent := r.Header.Get("User-Agent")

			if token != "" {
				fmt.Printf("Captured Token: %s\n", token)
			}
			if userAgent != "" {
				fmt.Printf("User-Agent: %s\n", userAgent)
			}
		}

		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			http.Error(w, "Error making request to target", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		copyHeader(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	//ip := getLocalIP()
	fmt.Println("Starting proxy server on IP:", ipAddress, "and port: 8080")
	fmt.Println("On your phone, set your proxy server to:", ipAddress, "with port: 8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

# go-h5speedtest
An easy to deploy backend for HTML5 speedtest



No need to deploy an nginx with php backend. With low memory cost and low disk IO usage, go-h5speedtest accelerates the deployment of speedtest to a single execuble!

This is the most speedy speedtest backend of all. Less of **10% at 1Gbps**!

Download file from release (or compile by yourself), put it anywhere and run and tada! You're ready to test the sincerity of your ISP.

**Usage:**
```
./go-h5speedtest 
  [-l  listening address(:80 by default)] 
  [-r  webroot path(speedtest by default)]
  [-d  daemon start]
```

![screenshot](https://raw.githubusercontent.com/Andyfoo/go-h5speedtest/master/go-h5speedtest.png)

--------

**Credit**

speedtest for html5 is a project by [github.com/librespeed/speedtest](https://github.com/librespeed/speedtest)

speedtest backend for golang is a project by [github.com/snowie2000/h5speedtest-go](https://github.com/snowie2000/h5speedtest-go)
Usecase 1
----------

hariharasudhan@mac clibuilder % go run main.go pwd -path=/bin/

✅ Output:
/Users/hariharasudhan/Documents/clibuilder/mycli/clibuilder

hariharasudhan@mac clibuilder % go run main.go ls -path=/bin/

✅ Output:
commands.yaml
go.mod
go.sum
main1.go
main2.go
var

-------------------------

Usecase 2
----------

hariharasudhan@mac clibuilder % go run main.go -config=commands.yaml list

total 32
drwxr-xr-x  7 hariharasudhan  staff   224 Apr 11 16:34 .
drwxr-xr-x  4 hariharasudhan  staff   128 Apr 11 14:18 ..
-rw-r--r--  1 hariharasudhan  staff    58 Apr 11 16:33 commands.yaml
-rw-r--r--  1 hariharasudhan  staff    62 Apr 11 16:34 go.mod
-rw-r--r--  1 hariharasudhan  staff   360 Apr 11 16:34 go.sum
-rw-r--r--  1 hariharasudhan  staff  1348 Apr 11 16:33 main.go
drwxr-xr-x  3 hariharasudhan  staff    96 Apr 11 10:33 var

hariharasudhan@mac clibuilder % go run main.go -config=commands.yaml show
package main
main(){

}

-------------------------

Usecase 3
----------

To install/ upgrade / view remote urls (Public git repos)

go run main.go install [repo_url]
go run main.go upgrade [repo_url]
go run main.go view

-------------------------

Usecase 4
----------

In this case plugin as a input . and to perform various operations like
* view all plugins
* To print args from the plugins
 (eg) go run main.go  cli-plugin hi hello 
* To execute plugin

I have created a plugin to check the health of the url

go run main.go url-health-checker https://www.google.com/
✅ https://www.google.com/ is healthy


export NAVI_PATH="/Users/hariharasudhan/Documents/clibuilder/mycli/navi"


# PiGit :grapes:	
## A simple git server in RaspberryPi without frontend 

---

### How to use

1. execute install.sh on your raspberry, you need a ssh server
2. compile main.go using: 

```bash
env GOOS=linux GOARCH=arm GOARM=5 go build
```
3. Send the output file and run in your raspberry
4. Open a browser or a http client and go to http://ip/createRepo/{yourRepoName}
5. Your repo is created

---

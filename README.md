This is the web service in Go that powers http://sshpot.com/ - it runs on Heroku.

# Local development
Use gin to get automatic reload while developing

```
gin --port 4040 -a '9090'
```

Use `make` to build the binary.

Use `make deploy` to push it to Heroku.

# TODO

* [ ] init.d script
* [ ] api usage graphs/page
* [ ] systemd script
* [ ] Tests
* [ ] Log commands executed without executing them
* [ ] Restrict access to sharing via API keys / auth management


# License

See LICENSE file

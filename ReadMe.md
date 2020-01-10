# An Util to Sort APNIC Address

Basicly, this program can deal with the list downloaded from http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest

## TODO

- [x] Download list from url & load it into struct.
- [x] Process IP range, convert it to CIDR schema.
- [ ] Return a list of data struct with given conditions.
- [ ] Appending data slice to a predefine template.
- [ ] Periodically download & process list with a daemon process in background.
- [ ] Ability to execute command after process job has done.
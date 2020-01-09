# A Util to Sort APNIC Address

Basicly, this program can deal with a list downloaded from http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest

## TODO

- [ ] Download list from url & load it into struct.
- [ ] Process IP range, convert it to CIDR schema.
- [ ] Return a slice of data struct with given conditions.
- [ ] Appending data slice to a predefine template.
- [ ] Periodically download & process list with a daemon process in background.
- [ ] Ability to execute command after process job has done.
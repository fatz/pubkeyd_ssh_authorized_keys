# pubkeyd_ssh_authorized_keys [![Build Status](https://travis-ci.org/fatz/pubkeyd_ssh_authorized_keys.svg?branch=master)](https://travis-ci.org/fatz/pubkeyd_ssh_authorized_keys)
A tiny tool requesting a users authorized_keys via HTTP to be used in `sshd_config`. It could be used with [pugkeyd](/lloesche/pubkeyd) or directly against github.com

## SSHD config
add this to your sshd config

```
AuthorizedKeysCommand /path/to/pubkeyd_ssh_authorized_keys
AuthorizedKeysCommandUser nobody
```

## pubkeyd_ssh_authorized_keys config
this client needs to know where pubkeyd is running or if github.com is used.

The client is searching for configuration at `/etc`, `$HOME` and `.` for a `pubkeyd.yaml`


### use your pubkeyd
just change baseurl to your pubkeyd hostname

```yaml
---
baseurl: https://yourpubkeyd.example.com
```

### use github
rewrite enables you to statically map local users against github users

```yaml
---
baseurl: https://github.com
pathf: "/%s.keys"
rewrite:
  jdoe: fatz
  foobar: fatz
```

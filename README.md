# MQuery-SRU

MQuery-SRU is an easy to setup endpoint for Clarin FCS 2.0 (Federated Content Search) based on
Manatee-open corpus search engine.

## Features

* Full support for the FCQ-QL query language
    * definable layer &#8596; pos. attribute mapping
* Level 1 support for basic search via CQL (Context Query
Language)
* search in multiple defined corpora


## Requirements

* a working Linux server with installed [Manatee-open](https://nlp.fi.muni.cz/trac/noske) library
* [Redis](https://redis.io/) database
* [Go](https://go.dev/)  language compiler and tools
* (optional) an HTTP proxy server (Nginx, Apache, ...)


## How to install

1. Install `Go` language environment, either via a package manager or manually from Go [download page](https://go.dev/dl/)
   1. make sure `/usr/local/go/bin` and `~/go/bin` are in your `$PATH` so you can run any installed Go tools without specifying a full path
2. Install Manatee-open from the [download page](https://nlp.fi.muni.cz/trac/noske). No specific language bindings are required.
   1. `configure --with-pcre && make && sudo make install && sudo ldconfig`
3. Get MQuery-SRU sources (`git clone --depth 1 github.com:czcorpus/mquery-sru.git`)
4. Run `make tools`
5. Run `make`


## Configuration

To run the endpoint, you need at least

1. to configure listening address and port
2. defined path to your Manatee corpora registry (= configuration) files
2. defined corpora along with:
    * positional attributes to be exposed
    * mapping of FCS-QL's `within` structures (`s`, `sentence`, `p` etc.) to your specific corpora structures
3. address and port of your Redis service plus a number of database to be used for passing queries and results around

See `conf.sample.json` for detailed info.

## OS integration (systemd)

(Here we assume the service will run with user `www-data`)

Create a directory for logging (e.g. `/var/log/mquery-sru`) and set proper permissions for `www-data` to be able to write there.

You can use predefined systemd files from `/scripts/systemd/*`. Copy (or link) them to `/etc/systemd/system` and then run:

```
systemctl enable mquery-sru-server.service
systemctl enable mquery-sru-worker-all.target
```

Now you can try to run the service:

```
systemctl start mquery-sru-server
systemctl start mquery-sru-worker-all
```


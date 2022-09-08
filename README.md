# DNS Experiments

A DNS server built using [miekg/dns](https://github.com/miekg/dns) in Golang that provides utilities related to the IndieWeb and my personal website.

## Using the Service

To use this DNS server, make a DNS query to the following resource:

    jamesg.blog

One way to query the resource is to use the `dig` command from the command line:

    dig @jamesg.blog resource.name

In this query:

- @jamesg.blog tells `dig` that you want to query the `jamesg.blog` DNS server
- resource.name is the name of the resource you want to query

## Reference Manual

Here are the commands presently available:

### IndieWeb

    dig @jamesg.blog is.it.hwc.day - check if today is Homebrew Website Club day
    dig @jamesg.blog next.indieweb.event - get the next IndieWeb event from events.indieweb.org

### Personal

    dig @jamesg.blog resume - get my resume
    dig @jamesg.blog recent.blog - get the title and URL of my most recent blog post

The server only returns TXT records. Querying for other record types will not return any responses.

## Installation

You can install this DNS server for yourself.

To do so, first clone this repository and install the required dependencies:

    git clone https://github.com/capjamesg/dns-experiments
    cd dns-experiments
    go install

Next, run the program:

    go run .

This command opens up a DNS server on port `5003`.

To make queries to the server on your local computer, you will need to query `localhost` and the port `5003`. Here is the dig equivalent for this query:

    dig @localhost -p 5003 resource.name

This command makes a request for the `resource.name` resource from the `localhost` DNS server on port `5003`.

## Acknowledgements

Thank you to [miekg](https://github.com/miekg) for building the Go miekg/dns library upon which this project depends. The miekg/dns library runs the actual DNS server and my code writes the requisite responses for queries.

## License

This project is licensed under an [MIT license](LICENSE).

## Contributors

- capjamesg

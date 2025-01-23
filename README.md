# DigitalOcean-to-Name.com DNS Migrator

Copy DNS records from DigitalOcean to Name.com.

This is a very simple tool intended for one-time migrations with a human involved in the process; it is not sufficiently robust to work without supervision or to be run repeatedly.

In particular, this tool is _not_ idempotent:

- It will attempt to copy every record, every time it is run, without regard for whether a similar record already exists in the destination.
- In the event of a failure partway through a migration, this tool will not resume from where it left off; it will start over from the beginning. 

This tool will not edit or remove existing records in Name.com under any circumstances.

## Build

Because this is a one-off tool for occasional use, I'm not providing a CI pipeline or prebuilt binaries/Docker images. To get a binary, check out the repo and build it:

```
git clone https://github.com/cdzombak/dns-do-to-namecom.git
cd dns-do-to-namecom
go build -o out .
```

The only requirement is a working Go installation.

## Usage

```
DO_API_TOKEN=dop_v1_foobarbaz \
NC_USERNAME=myusername \
NC_API_TOKEN=abcd0000 \
./out/dns-do-to-namecom -domain MYDOMAIN.COM [-dry-run=false] 
```

- Flag `-domain`: The domain name for DNS migration.
- Flag `-dry-run`: If true, the tool will not make any changes to Name.com; it will only print what it would do. `True` by default; you must pass `-dry-run=false` to make any changes to Name.com.
- Environment variable `DO_API_TOKEN`: DigitalOcean API token with read access.
- Environment variable `NC_USERNAME`: Name.com username.
- Environment variable `NC_API_TOKEN`: Name.com API token.

## Configuration

Environment variables may be placed in the special file `.env`. This file will be read automatically from the working directory if it exists. See [`env.sample`](env.sample) for an example.

## API Tokens/Management

- DigitalOcean: https://cloud.digitalocean.com/account/api/tokens
- Name.com: https://www.name.com/account/settings/api

## License

LGPL 3.0; see [LICENSE](LICENSE) in this repository.

## Author

Chris Dzombak ([dzombak.com](https://www.dzombak.com); [GitHub @cdzombak](https://github.com/cdzombak))

## See Also

- [DNS Auditor](https://github.com/cdzombak/dns-auditor)
- [DigitalOcean to Porkbun DNS Migrator](https://github.com/cdzombak/dns-do-to-porkbun)
- [Name.com to DigitalOcean DNS Migrator](https://github.com/cdzombak/dns-migrator)
- [DigitalOcean Dynamic DNS tool](https://github.com/cdzombak/do-ddns)

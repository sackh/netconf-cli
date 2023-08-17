# NETCONF CLI

A simple CLI tool to interact with devices using  NETCONF protocol.

## Installation

```bash
go install github.com/sackh/netconf-cli@latest
```

```bash
go get -u github.com/sackh/netconf-cli
```

## Usage

For netconf-cli tool config file is mandatory. Config file must be in dotenv format.

Example config file must have below key value pairs:

```dotenv
host=a.b.c.d
username=<username>
password=<password>
```

Value of the host should be the ip address of the device and username and password of the device.

## Commands

Currently below commands are supported:

- Get Capabilities

```bash
netconf-cli get capabilities --config <netconf-cli.env>
```

- Get Candidate Config

```bash
netconf-cli get config --candidate --config <netconf-cli.env>
```

- Get Running Config

```bash
netconf-cli get config --running --config <netconf-cli.env>
```

- Get Reply

```bash
netconf-cli get reply -r <rpc-request-xml-file> --config <netconf-cli.env>
```

> :warning: **This CLI tool is currently in the experimental stage and APIs and commands might change in future.**

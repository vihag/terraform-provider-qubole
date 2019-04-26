Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-qubole`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; 
$ cd $GOPATH/src/github.com/terraform-providers
$ git clone https://github.com/vihag/terraform-provider-qubole.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-qubole
$ export TF_LOG=DEBUG
$ export TF_LOG_PATH=/tmp/tflog.log

$ clear
$ >/tmp/tflog.log
$ go build -o terraform-provider-qubole
```

Using the provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it. Documentation about the provider specific configuration options can be found on the [provider's website]

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `go build -o terraform-provider-qubole`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ go build -o terraform-provider-qubole
...
$ $GOPATH/bin/terraform-provider-qubole
...
```

Test suites are coming


# ZEROSSL CLI

ZeroSSL CLI is a tool to interact with the ZEROSSL API using CLI.

## Must know

The CLI generates a folder in `$HOME/.zerossl-cli/`, where a sqlite database will be stored (only for validations) and folders will also be created there with the certificates you download.

## Features

- List certificates by status and expiration date.
- Create certificates from scratch
- Cancel certificates
- Revoke certificates
- Validate certificates
- Download certificates

## How to use?

First, it is necessary to have the APIKEY that Zero SSL provides you. You can find this in your panel, entering the "Developers" area

Once you have it, you must save it in an environment variable on your system.

On linux, you can run the following

```bash
export ZEROSSL_API_KEY=YOUR_API_KEY_HERE
```

And ready! Now you can start using the CLI!

## List certificates

```bash
zerossl-cli certificates list
```

By executing the previous command, the tool will search for certificates with status **"issued"**, if you want to filter by another status, add the **--status** flag to the command.

> Available statuses: `draft`, `pending_validation`, `issued`, `cancelled`, `revoked` and `expired`

Likewise, if you want to see the certificates that are soon to expire, you can add the --expiring-days flag as follows:

```bash
zerossl-cli certificates list --expiring-days=15 --status=issued
```

The previous command looks for the issued certificates that expire in the next 15 days.

## Create certificate

To create a certificate simply run the following command

```bash
zerossl-cli certificates create
```

The cli will ask you questions to finish completing the creation.

## Validate certificate

Once you finish generating the certificate and have chosen the verification method, when you complete the verification method, to validate the certificate you must execute the following command

```bash
zerossl-cli certificates validate ${CERTIFICATE_ID}
```

> You must replace **${CERTIFICATE_ID}** with the ID of the certificate you want to validate

## Cancel certificate

To cancel a certificate you must execute the following command

```bash
zerossl-cli certificates cancel ${CERTIFICATE_ID}
```

> You must replace **${CERTIFICATE_ID}** with the ID of the certificate you want to cancel.<br>
> Only **non issued** certificates can be cancelled

## Revoke certificate

To revoke a certificate you must execute the following command.

```bash
zerossl-cli certificates revoke ${CERTIFICATE_ID}
```

> You must replace **${CERTIFICATE_ID}** with the ID of the certificate you want to revoke.<br>
> Only certificates with **issued** status can be revoked

## Download certificates

Once the certificate is in issued state, you can download it using the following command

```bash
zerossl-cli certificates download ${CERTIFICATE_ID}
```

> You must replace **${CERTIFICATE_ID}** with the ID of the certificate you want to download.<br>
> Only certificates with **issued** status can be downloaded

The certificates are stored in a folder at the following path:

```bash
$HOME/.zerossl-cli/
```

## Contributions

To contribute, fork the repository and then submit the pull request for analysis.<br>
This is an open source project and is open to modifications by the community.

## Contact

This tool was developed by Federico Juretich <fedejuret@gmail.com>.<br>
Any suggestion or query can be made via email without any inconvenience.

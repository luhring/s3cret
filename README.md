# s3cret

Copy files to/from AWS S3 with automatic client-side encryption/decryption.

s3cret is **pronounced "secret"**. If you pronounce it "sthreecret", you will have bad luck for the rest of the day.

## Installation

For now, the only supported installation method is cloning the source and running `go install`. More installation methods to come.

## Usage

s3cret is designed to feel familiar for people that have used the AWS CLI commands for S3, and in particular, the `aws s3 cp` subcommand.

### From local to S3

Encrypt a local file and then copy the encrypted file to S3.

```shell-output
$ s3cret ./special-file.txt s3://my-bucket/
encrypting "special-file.txt" using KMS key A1B2C3D4
copying encrypted file to S3 bucket "my-bucket"
```

### From S3 to local

Copy an encrypted file from S3 to the local system and then decrypt it.

```shell-output
$ s3cret s3://my-bucket/special-file.txt ./
copying encrypted file "special-file.txt" from S3 bucket "my-bucket"
decrypting "special-file.txt" using KMS key A1B2C3D4
```

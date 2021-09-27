# How to run build

1. Create an admin user in SAST
2. Run export
```
    ./cxsast_exporter --user username --pass password --url http://localhost
```

This will generate an export package in the same folder where the command is run.

add --debug parameter to bypass the zip and encryption process. 

# How to build

1. Make sure you have access to AWS KMS
2. Fetch KMS public key
3. Build

Having configured AWS CLI access, you can fetch the public key using the following command:
```
aws kms get-public-key --key-id "cb3052be-1e3a-4a9c-b3f0-84d963c53a06" | jq -r .PublicKey > public.key
```

Note: 
 * the key id in this example is for development environment
 * this command depends on having jq command (https://stedolan.github.io/jq/) installed
 
Then you can build with: `make build`

# How to run repo

In order to run during development we'll need to `make public_key`, like for build.
Once the public.key file exists, we can run and test using the file reference:
```
go run -ldflags "-X sast-export/internal.buildTimeRSAPublicKey=$(cat .\public.key)" .
```

```
go test -ldflags "-X sast-export/internal.buildTimeRSAPublicKey=$(cat .\public.key)" .\...
```

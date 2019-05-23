# standard-mapper

Used for mapping a document (e.g. <https://github.com/andersju/gdpr-xml/blob/master/consolidated/gdpr-consolidated-en.xml>) to standard format

## Steps

### Step 1

Clone repo

```bash
git clone git@github.com:slalom/standard-mapper.git
cd standard-mapper
```

### Step 2

First convert document to `json`. For example, to convert gdpr-consolidated-en.xml to json use `xq` (`yq` <https://github.com/kislyuk/yq> comes with `xq`):

```bash
xq '.' gdpr-consolidated-en.xml > gdpr.json
```

A version has been saved to [assets/gdpr-en.json](assets/gdpr-en.json)

### Step 3

Use `jsonfiddle` <https://github.com/go-jsonfile/jsonfiddle> to generate `go` types from json data

```bash
jsonfiddle j2s -i assets/gdpr-en.json
```

Create a file in `types` package and add generated types. For `gdpr` see [pkg/types/gdpr.go](pkg/types/gdpr.go)

### Step 4

Create a new mapper in `mappers` package. It will map from the types created in step #2 to standard types at [pkg/types/standards.go](pkg/types/standards.go). For `gdpr` see [pkg/mappers/gdpr-mapper.go](pkg/mappers/gdpr-mapper.go).

### Step 5

Install go

```bash
brew install go
```

### Step 6

Build and run standard mapper

```bash
export GO111MODULE=on
go build && ./standard-mapper --inputFile assets/gdpr-en.json --outputFile standard.json
```

This will map gdpr-en.json to standard format and write to file `standard.json`. It also creates the file `unmarshal-debug.json` that is result of marshal'ing gdpr-en.json

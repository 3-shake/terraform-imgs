# terraform-imgs

## What is terraform-imgs
A utility to generate MERMAID code from terraform code using OpenAI API.

## Installlation
TBD

## Usage

```bash
$ terraform-imgs --help
A utility to generate MERMAID code from terraform code using OpenAI API.

Usage:
  terraform-imgs [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  mermaid     Generate mermaid diagrams from terraform files
  version     Print the version number of terraform-imgs

Flags:
  -h, --help                 help for terraform-imgs
  -o, --output-file string   output file path(e.g. README.md)

Use "terraform-imgs [command] --help" for more information about a command.
```

### Preparation
Set the environment variable `OPENAI_API_KEY` to the API key of the OpenAI API.

```
export OPENAI_API_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx >> ~/.bash_profile
```


### Output to stdout

```bash
$ terraform-imgs mermaid ./path/to/terraform-root
```

### Writing to md files

```bash
$ terraform-imgs mermaid ./path/to/terraform-root -o ./path/to/README.md
```
The mermaid code block is inserted between the following comment-outs. If it is not mentioned, it will be added to the end of the file.
```md
<!-- BEGIN_TF_IMGS -->
mermaid code block
<!-- END_TF_IMGS -->
```


## Notes

- The OpenAI API is used to generate the MERMAID code; API keys are the user's own, so API usage fees apply(https://openai.com/pricing).
  - Also, the API has a token limit, so it may not work for large terraforms.
- Please be sure to check the output results and modify them accordingly.

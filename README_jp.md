# terraform-imgs

## 概要
OpenAI APIを利用してterraformコードからMERMAIDコードを生成します。

## インストール
TBD

## 使用方法

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

### 事前準備
環境変数`OPENAI_API_KEY`にOpenAI APIのAPIキーを設定してください。

```
export OPENAI_API_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx >> ~/.bash_profile
```


### 標準出力

```bash
$ terraform-imgs mermaid ./path/to/terraform-root
```

### ファイルへの書き込み

```bash
$ terraform-imgs mermaid ./path/to/terraform-root -o ./path/to/README.md
```
mermaidコードブロックは以下のコメントアウトの間に挿入されます。記載がない場合はファイル末尾へ追記されます。

```md
<!-- BEGIN_TF_IMGS -->
mermaid code block
<!-- END_TF_IMGS -->
```


## Notes

- MERMAIDコードの生成にはOpenAI APIを利用しています。APIキーは利用者自身のものを利用するためAPI利用料金が発生します(https://openai.com/pricing)。
  - また、token制限がAPIには存在するため巨大なterraformの場合動作しない場合があります。
- あくまでもAIによる作図補助ツールであるため利用の際は必ず出力結果を確認し、適宜修正を行ってください。

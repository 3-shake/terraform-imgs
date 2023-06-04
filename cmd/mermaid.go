package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/3-shake/terraform-imgs/internal/openai"

	"github.com/spf13/cobra"
)

const (
	beginTag           = "<!-- BEGIN_TF_IMGS -->"
	endTag             = "<!-- END_TF_IMGS -->"
	codeBlockEndLength = 3
	basePrompt         = `
## GOAL ###
Create a highly visible system configuration diagram from terraform.

### CONTEXT ###
- You are an AI supporting an infrastructure engineer.
- The user does not have time to create a system configuration diagram to explain to his team members and would like you to do it for him.
- The only information about the current system configuration is what is described in terraform

### RULES ###.
- "### TEXT ###" will be followed by the directory structure of the terraform project. After that, "==<file_name>==" will be followed by the contents of the file in question.
- You will create a system configuration diagram using mermaid.
- Since it is a system configuration diagram, the detailed IAM relationships and information about the contents of Secrets are not important.
- Think in the order described in "### STEP ###".

### STEP ###
1. get an overview of the system from the information entered.
2. draw a diagram of the system overview using mermaid. 
3. Check mermaid for syntax errors.

### TEXT ###
`
)

// mermaidCmd represents the mermaid command
var mermaidCmd = &cobra.Command{
	Use:   "mermaid",
	Short: `Generate mermaid diagrams from terraform files`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := cmd.Help()
			if err != nil {
				return
			}
		}
		path := args[0]

		tfContext, err := scanFiles(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		apiResponse, err := openai.Chat(basePrompt + tfContext)
		if err != nil {
			fmt.Println(err)
			return
		}

		mermaidCode, err := formatMermaid(apiResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		if outputFile != "" {
			writeMermaid(mermaidCode, outputFile)
		} else {
			fmt.Println(mermaidCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(mermaidCmd)
}

// scanFiles scans directory and returns tree-like string and content of Terraform files.
func scanFiles(dirPath string) (string, error) {
	var tree strings.Builder
	var tfFiles strings.Builder

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore hidden files and directories.
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Add directory and file structure to output.
		relPath, _ := filepath.Rel(dirPath, path)
		indent := strings.Repeat("  ", strings.Count(relPath, string(os.PathSeparator)))
		tree.WriteString(indent + info.Name() + "\n")

		// If file is a Terraform file, add its content to output.
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".tf") || strings.HasSuffix(info.Name(), ".tfvars")) {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			tfFiles.WriteString("== " + info.Name() + " ==\n")
			tfFiles.WriteString(string(content) + "\n")
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return tree.String() + tfFiles.String(), nil
}

func formatMermaid(apiResponse string) (string, error) {
	// mermaid記載部分を抜き取り文字列を返却する
	begin := strings.Index(apiResponse, "```mermaid")
	end := strings.LastIndex(apiResponse, "```") + codeBlockEndLength

	if begin == -1 || end == -1 {
		return "", fmt.Errorf("mermaid code block not found")
	}
	return apiResponse[begin:end], nil
}

func writeMermaid(mermaidCode string, outputFile string) {
	// ファイル内で<TF_IMG_BEGIN>と<TF_IMG_END>で囲まれた部分をmermaidCodeで置換する。存在しない場合は末尾に追記する。

	content, err := os.ReadFile(outputFile)
	if err != nil {
		panic(err)
	}

	text := string(content)

	beginIndex := strings.Index(text, beginTag)
	endIndex := strings.Index(text, endTag)

	var tmplStr string

	if beginIndex == -1 || endIndex == -1 {
		tmplStr = text + beginTag + "\n{{.Result}}\n" + endTag + "\n"
	} else {
		tmplStr = text[:beginIndex+len(beginTag)] + "\n{{.Result}}\n" + text[endIndex:] + "\n"
	}
	tmpl, err := template.New("mermaid").Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct{ Result string }{Result: mermaidCode})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outputFile, buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

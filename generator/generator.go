package generator

import (
	"bytes"
	goformat "go/format"
	"text/template"
	
	"github.com/zeromicro/go-zero/tools/goctl/api/util"
)

type (
	fileGenConfig struct {
		dir             string
		subDir          string
		filename        string
		templateName    string
		builtinTemplate string
		data            interface{}
	}
)

func genFile(c fileGenConfig) error {
	fp, created, err := util.MaybeCreateFile(c.dir, c.subDir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	// 暂时不支持自定义模板
	text := c.builtinTemplate

	t := template.Must(template.New(c.templateName).Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := formatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func formatCode(code string) string {
	ret, err := goformat.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"gopkg.in/yaml.v2"
)

func ExecuteTemplate(in io.Reader, out io.Writer, templateFiles ...string) error {

	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return fmt.Errorf("[ERROR] Parsing template(s): %v", err)
	}

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, in)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read stdin: %v", err)
	}

	var values map[string]interface{}

	err = yaml.Unmarshal(buffer.Bytes(), &values)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse stdin: %v", err)
	}

	err = tmpl.Execute(out, &values)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to parse stdin: %v", err)
	}

	return nil
}

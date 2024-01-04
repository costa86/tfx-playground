package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	tfe "github.com/hashicorp/go-tfe"
	"github.com/rodaine/table"
)

var print = fmt.Println

func createVariable(client tfe.Client) {
	ctx := context.Background()

	a, err := client.Variables.Create(ctx, "ws-WD18c57pErqiDzp2", tfe.VariableCreateOptions{
		Type:        "vars",
		Key:         tfe.String("personName"),
		Value:       tfe.String("ana"),
		Description: tfe.String("wef"), Category: tfe.Category("env"), HCL: tfe.Bool(true), Sensitive: tfe.Bool(false)})
	if err != nil {
		log.Fatal(err)
	}
	print(a)
}

func listWorkspaces(client tfe.Client, name string) *tfe.WorkspaceList {
	ctx := context.Background()

	a, err := client.Workspaces.List(ctx, name, nil)
	if err != nil {
		log.Fatal(err)
	}
	println(a)
	return a

}

func main() {
	token := flag.String("token", "", "string")
	workspace := flag.String("workspace", "", "string")
	address := flag.String("address", "", "string")

	flag.Parse()

	config := &tfe.Config{
		Token:             *token,
		RetryServerErrors: true,
		Address:           *address,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	displayWorkspacesTable(client, *workspace)
	// createVariable(*client)

}

func displayWorkspacesTable(client *tfe.Client, name string) {
	list := listWorkspaces(*client, name)
	headerFmt := color.New(color.FgMagenta, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "NAME", "TERRAFORM VERSION")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, v := range list.Items {
		tbl.AddRow(v.ID, v.Name, v.TerraformVersion)
	}

	tbl.Print()
}

package cmd

import(
  "fmt"
  "github.com/spf13/cobra"
) 

var RootCmd = &cobra.Command{
  Use: "task",
  Short: "Hugo lalalalalalal",
  Run: func(cmd *cobra.Command, args []string){
    // initial actions on open 
    fmt.Println("hello2")
    
  },
}

package main

import (
    "os"
    "fmt"
	"bufio"
    "regexp"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/cobra"
)
import "hackstore/db"

func isValidRootDomain(domain string) bool {
    // Define a regular expression pattern to match domain names
    pattern := `^(?:[-A-Za-z0-9]+\.)*[A-Za-z0-9][-A-Za-z0-9]*\.[A-Za-z]{2,6}$`
    match, _ := regexp.MatchString(pattern, domain)
    return match
}

func readLinesFromFile(file string) ([]string, error) {
    var lines []string

    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

type Program struct {
    ProgramName string
    CreatedAt   string
}

type RootDomain struct {
    RootDomain  string
    ProgramName string
    CreatedAt   string
}

type Subdomain struct {
    Subdomain   string
    RootDomain  string
    CreatedAt   string
}

type AliveDomain struct {
    AliveDomain string
    RootDomain  string
    CreatedAt   string
}

var rootCmd = &cobra.Command{
    Use:   "hackstore",
    Short: "CLI tool for Hackstore",
    Long:  `A command-line interface tool for managing programs, root domains, subdomains, and alive domains in Hackstore.`,
    Run: func(cmd *cobra.Command, args []string) {
        // If no subcommand is provided
        cmd.Help()
    },
}

var programsCmd = &cobra.Command{
    Use:   "programs",
    Short: "Manage programs",
    Long:  "Manage programs in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // If no subcommand is provided for 'programs'
        cmd.Help()
    },
}

var rootDomainsCmd = &cobra.Command{
    Use:   "root-domains",
    Short: "Manage root domains",
    Long:  "Manage root domains associated with programs in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // If no subcommand is provided for 'root-domains'
        cmd.Help()
    },
}

var subdomainsCmd = &cobra.Command{
    Use:   "subdomains",
    Short: "Manage subdomains",
    Long:  "Manage subdomains associated with root domains in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // If no subcommand is provided for 'subdomains'
        cmd.Help()
    },
}

var alivedomainsCmd = &cobra.Command{
    Use:   "alivedomains",
    Short: "Manage alivedomains",
    Long:  "Manage alivedomains associated with root domains in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // If no subcommand is provided for 'subdomains'
        cmd.Help()
    },
}
//listing commands-------------------

var listallSubdomainsCmd = &cobra.Command{
    Use:   "listall",
    Short: "List all subdomains",
    Long:  "List all subdomains in the database",
    Run: func(cmd *cobra.Command, args []string) {
        // Call your function to fetch all subdomains
        subdomains, err := db.GetAllSubdomains()
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched subdomains
        for _, subdomain := range subdomains {
            fmt.Println(subdomain) // Replace this with your logic to display subdomains
        }
    },
}


var listSubdomainsCmd = &cobra.Command{
    Use:   "list",
    Short: "List subdomains associated with a root domain",
    Long:  "List subdomains associated with a root domain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        rootDomain, _ := cmd.Flags().GetString("root-domain")

        // Call your db function to fetch subdomains associated with the specified root domain
        subdomains, err := db.GetSubdomainsByRootDomain(rootDomain)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched subdomains
        for _, subdomain := range subdomains {
            fmt.Println(subdomain) // Replace this with your fmtic to display subdomains
        }
    },
}

var listProgramsCmd = &cobra.Command{
    Use:   "list",
    Short: "List all programs",
    Long:  "List all programs in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // Call your db function to fetch all programs
        programs, err := db.GetAllPrograms()
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched programs
        for _, program := range programs {
            fmt.Println(program.ProgramName) // Replace this with your fmtic to display programs
        }
    },
}

var listallRootDomainsCmd = &cobra.Command{
    Use:   "listall",
    Short: "List all root domains",
    Long:  "List all root domains in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        // Call your db function to fetch all root domains
        rootDomains, err := db.GetAllRootDomains()
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched root domains
        for _, rd := range rootDomains {
            fmt.Println(rd.RootDomain) // Replace this with your logic to display root domains
        }
    },
}

var listRootDomainsCmd = &cobra.Command{
    Use:   "list",
    Short: "List root domains by program",
    Long:  "List root domains by program in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        program, _ := cmd.Flags().GetString("program")
        // Call your db function to fetch root domains by program
        rootDomains, err := db.GetRootDomainsByProgram(program)
        
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched root domains
        for _, rootDomain := range rootDomains {
            fmt.Println(rootDomain.RootDomain) // Replace this with your fmtic to display root domains
        }
    },
}

var listallAliveDomainsCmd = &cobra.Command{
    Use:   "listall",
    Short: "List all alive domains",
    Long:  "List all alive domains in the database",
    Run: func(cmd *cobra.Command, args []string) {
        // Call your db function to fetch all alive domains
        aliveDomains, err := db.GetAllAliveDomains()
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched alive domains
        for _, domain := range aliveDomains {
            fmt.Println(domain)
        }
    },
}


var listAliveDomainsCmd = &cobra.Command{
    Use:   "list",
    Short: "List alive domains by root domain",
    Long:  "List alive domains by root domain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        rootDomain, _ := cmd.Flags().GetString("root-domain")
        // Call your db function to fetch alive domains by root domain
        aliveDomains, err := db.GetAliveDomainsByRootDomain(rootDomain)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Display the fetched alive domains
        for _, aliveDomain := range aliveDomains {
            fmt.Println(aliveDomain) // Replace this with your fmtic to display alive domains
        }
    },
}


//importing from file data
var importProgramsCmd = &cobra.Command{
    Use:   "import",
    Short: "Import programs from a file",
    Long:  "Import programs from a text file",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        programName, _ := cmd.Flags().GetString("program")

        lines, err := readLinesFromFile(file)
        if err != nil {
            fmt.Println(err)
            return
        }

        if programName != "" {
            lines = append(lines, programName)
        }

        // Call your db function to add programs to the database
        for _, line := range lines {
            // Add each program to the database
            err := db.CreateProgram(line)
            if err != nil {
                fmt.Println(err)
            }
        }
    },
}

var importRootDomainsCmd = &cobra.Command{
    Use:   "import",
    Short: "Import root domains from a file",
    Long:  "Import root domains from a text file and associate them with a program",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        programName, _ := cmd.Flags().GetString("program")

        lines, err := readLinesFromFile(file)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Filter out invalid root domain lines and eliminate duplicates
        uniqueRootDomains := make(map[string]bool)
        for _, rootDomain := range lines {
            rootDomain = strings.TrimSpace(rootDomain)
            if isValidRootDomain(rootDomain) {
                uniqueRootDomains[rootDomain] = true
            } else {
                fmt.Printf("Invalid root domain: %s\n", rootDomain)
            }
        }

        // Store unique valid root domains in a slice
        validRootDomains := make([]string, 0)
        for rootDomain := range uniqueRootDomains {
            validRootDomains = append(validRootDomains, rootDomain)
        }

        successfulImports := 0 // Counter for successful imports

        // Call your db function to add valid root domains by program to the database
        for _, rootDomain := range validRootDomains {
            err := db.AddRootDomain(programName, rootDomain)
            if err != nil {
                fmt.Println(err)
            } else {
                successfulImports++
            }
        }

        // Display success message after the loop completes
        fmt.Printf("%d root domains successfully imported.\n", successfulImports)
    },
}


var importSubdomainsCmd = &cobra.Command{
    Use:   "import",
    Short: "Import subdomains from a file",
    Long:  "Import subdomains from a text file and associate them with a root domain",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        rootDomain, _ := cmd.Flags().GetString("rootdomain")

        lines, err := readLinesFromFile(file)
        if err != nil {
            fmt.Println(err)
            return
        }

        uniqueSubdomains := make(map[string]struct{}) // Track unique subdomains
        successfulImports := 0                        // Counter for successful imports

        // Call your db function to add subdomains by root domain to the database
        for _, subdomain := range lines {
            // Check if the subdomain is unique
            _, exists := uniqueSubdomains[subdomain]
            if exists {
                continue // Skip duplicate subdomains
            }

            // Add the subdomain to the uniqueSubdomains map
            uniqueSubdomains[subdomain] = struct{}{}

            // Add each unique subdomain by root domain to the database
            err := db.AddSubdomain(rootDomain, subdomain)
            if err != nil {
                fmt.Println(err)
            } else {
                successfulImports++
            }
        }

        // Display success message after the loop completes
        fmt.Printf("%d subdomains successfully imported.\n", successfulImports)
    },
}


var importAliveDomainsCmd = &cobra.Command{
    Use:   "import",
    Short: "Import alive domains from a file",
    Long:  "Import alive domains from a text file and associate them with a root domain",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        rootDomain, _ := cmd.Flags().GetString("rootdomain")

        lines, err := readLinesFromFile(file)
        if err != nil {
            fmt.Println(err)
            return
        }

        uniqueAliveDomains := make(map[string]bool)
        successfulImports := 0 // Counter for successful imports

        // Eliminate duplicates and add unique alive domains to the map
        for _, aliveDomain := range lines {
            if _, exists := uniqueAliveDomains[aliveDomain]; !exists {
                uniqueAliveDomains[aliveDomain] = true
            }
        }

        // Call your db function to add alive domains by root domain to the database
        for aliveDomain := range uniqueAliveDomains {
            // Add each unique alive domain by root domain to the database
            err := db.AddAliveDomain(rootDomain, aliveDomain)
            if err != nil {
                fmt.Println(err)
            } else {
                successfulImports++
            }
        }

        // Display success message after the loop completes
        fmt.Printf("%d domains successfully imported.\n", successfulImports)
    },
}

//-------------Adding single values---------------

var addProgramCmd = &cobra.Command{
    Use:   "add-program",
    Short: "Add a new program",
    Long:  "Add a new program to Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        programName, _ := cmd.Flags().GetString("program")
        // Call your db function to add a new program to the database
        err := db.CreateProgram(programName)
        if err != nil {
            fmt.Println(err)
        }
    },
}

var addRootDomainsCmd = &cobra.Command{
    Use:   "add-root-domains",
    Short: "Add root domains by program",
    Long:  "Add root domains by program to Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        programName, _ := cmd.Flags().GetString("program")
        rootDomains, _ := cmd.Flags().GetStringSlice("rootdomain")
        successfulImports := 0 // Counter for successful imports

        for _, rootDomain := range rootDomains {
            if !isValidRootDomain(rootDomain) {
                fmt.Printf("Invalid root domain: %s\n", rootDomain)
                continue
            }

            // Call your db function to add root domain by program to the database
            err := db.AddRootDomain(programName, rootDomain)
            if err != nil {
                fmt.Println(err)
            } else {
                successfulImports++
            }
        }

        // Display success message after the loop completes
        fmt.Printf("%d root domains successfully added.\n", successfulImports)
    },
}


var addSubdomainCmd = &cobra.Command{
    Use:   "add-subdomain",
    Short: "Add a subdomain to a root domain",
    Long:  "Add a subdomain to a root domain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        rootDomain, _ := cmd.Flags().GetString("rootdomain")
        subdomain, _ := cmd.Flags().GetString("subdomain")
        // Call your db function to add a subdomain to the root domain in the database
        err := db.AddSubdomain(rootDomain, subdomain)
        if err != nil {
            fmt.Println(err)
        }
    },
}

var addAliveDomainCmd = &cobra.Command{
    Use:   "add-alive-domain",
    Short: "Add an alive domain to a root domain",
    Long:  "Add an alive domain to a root domain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        rootDomain, _ := cmd.Flags().GetString("rootdomain")
        aliveDomain, _ := cmd.Flags().GetString("alivedomain")
        // Call your db function to add an alive domain to the root domain in the database
        err := db.AddAliveDomain(rootDomain, aliveDomain)
        if err != nil {
            fmt.Println(err)
        }
    },
}

//-----deleting data
var deleteProgramCmd = &cobra.Command{
    Use:   "delete-program",
    Short: "Delete a program and associated data",
    Long:  "Delete a program and its associated data in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        programName, _ := cmd.Flags().GetString("program")
        // Call your db function to delete the program and associated data from the database
        err := db.DeleteProgram(programName)
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("Program '%s' and associated data successfully deleted.\n", programName)
    },
}

var deleteRootDomainCmd = &cobra.Command{
    Use:   "delete-root-domain",
    Short: "Delete a root domain and associated data",
    Long:  "Delete a root domain and its associated data in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        rootDomain, _ := cmd.Flags().GetString("rootdomain")
        programName, _ := cmd.Flags().GetString("program")
        // Call your db function to delete the root domain and associated data from the database
        err := db.DeleteRootDomain(rootDomain, programName)
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("Root domain '%s' and associated data for program '%s' successfully deleted.\n", rootDomain, programName)
    },
}


var deleteSubdomainCmd = &cobra.Command{
    Use:   "delete-subdomain",
    Short: "Delete a subdomain",
    Long:  "Delete a subdomain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        subdomain, _ := cmd.Flags().GetString("subdomain")
        rootDomain, _ := cmd.Flags().GetString("rootdomain")
        // Call your db function to delete the subdomain from the root domain in the database
        err := db.DeleteSubdomain(subdomain, rootDomain)
        if err != nil {
            fmt.Println(err)
        }
    },
}

var deleteAliveDomainCmd = &cobra.Command{
    Use:   "delete-alive-domain",
    Short: "Delete an alive domain",
    Long:  "Delete an alive domain in Hackstore",
    Run: func(cmd *cobra.Command, args []string) {
        aliveDomain, _ := cmd.Flags().GetString("alivedomain")
        rootDomain, _ := cmd.Flags().GetString("rootdomain")
        // Call your db function to delete the alive domain from the root domain in the database
        err := db.DeleteAliveDomain(aliveDomain, rootDomain)
        if err != nil {
            fmt.Println(err)
        }
    },
}
//-----



func init() {
    rootCmd.AddCommand(programsCmd, rootDomainsCmd, subdomainsCmd, alivedomainsCmd)
    programsCmd.AddCommand(listProgramsCmd, importProgramsCmd, addProgramCmd, deleteProgramCmd)
    rootDomainsCmd.AddCommand(listallRootDomainsCmd, listRootDomainsCmd, importRootDomainsCmd, addRootDomainsCmd, deleteRootDomainCmd)
    subdomainsCmd.AddCommand(listallSubdomainsCmd, listSubdomainsCmd, importSubdomainsCmd, addSubdomainCmd, deleteSubdomainCmd)
    alivedomainsCmd.AddCommand(listallAliveDomainsCmd, listAliveDomainsCmd, importAliveDomainsCmd, addAliveDomainCmd, deleteAliveDomainCmd)

    listRootDomainsCmd.Flags().StringP("program", "p", "", "Program name: test")
    listRootDomainsCmd.MarkFlagRequired("program")

    listSubdomainsCmd.Flags().StringP("root-domain", "r", "", "Root domain ex.com")
    listSubdomainsCmd.MarkFlagRequired("root-domain")

    listAliveDomainsCmd.Flags().StringP("root-domain", "r", "", "Root domain ex.com")
    listSubdomainsCmd.MarkFlagRequired("root-domain")

    importProgramsCmd.Flags().StringP("file", "f", "", "File containing programs: programs.txt")
    importProgramsCmd.MarkFlagRequired("file")

	importRootDomainsCmd.Flags().StringP("program", "p", "", "Program name: test")
    importRootDomainsCmd.Flags().StringP("file", "f", "", "File containing root domains: root-domains.txt")
    importRootDomainsCmd.MarkFlagRequired("program")
    importRootDomainsCmd.MarkFlagRequired("file")

	importSubdomainsCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    importSubdomainsCmd.Flags().StringP("file", "f", "", "File containing subdomains: subdomains.txt")
    importSubdomainsCmd.MarkFlagRequired("rootdomain")
    importSubdomainsCmd.MarkFlagRequired("file")

	importAliveDomainsCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test")
    importAliveDomainsCmd.Flags().StringP("file", "f", "", "File containing alive domains: alive_domains.txt")
    importAliveDomainsCmd.MarkFlagRequired("rootdomain")
    importAliveDomainsCmd.MarkFlagRequired("file")

	addProgramCmd.Flags().StringP("program", "p", "", "Program name: test")
    addProgramCmd.MarkFlagRequired("program")

	addRootDomainsCmd.Flags().StringP("program", "p", "", "Program name: test")
    addRootDomainsCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    addRootDomainsCmd.MarkFlagRequired("program")
    addRootDomainsCmd.MarkFlagRequired("rootdomain")

    addSubdomainCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    addSubdomainCmd.Flags().StringP("subdomain", "s", "", "Subdomain: sub.test.com")
    addSubdomainCmd.MarkFlagRequired("rootdomain")
    addSubdomainCmd.MarkFlagRequired("subdomain")

	addAliveDomainCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    addAliveDomainCmd.Flags().StringP("alivedomain", "a", "", "Alive domain: http(s)://sub.test.com")
    addAliveDomainCmd.MarkFlagRequired("rootdomain")
    addAliveDomainCmd.MarkFlagRequired("alivedomain")

	deleteProgramCmd.Flags().StringP("program", "p", "", "Program name: test")
    deleteProgramCmd.MarkFlagRequired("program")

	deleteRootDomainCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    deleteRootDomainCmd.Flags().StringP("program", "p", "", "Program name: test")
    deleteRootDomainCmd.MarkFlagRequired("rootdomain")
    deleteRootDomainCmd.MarkFlagRequired("program")

	deleteSubdomainCmd.Flags().StringP("subdomain", "s", "", "Subdomain: sub.test.com")
    deleteSubdomainCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    deleteSubdomainCmd.MarkFlagRequired("subdomain")
    deleteSubdomainCmd.MarkFlagRequired("rootdomain")

	deleteAliveDomainCmd.Flags().StringP("alivedomain", "a", "", "Alive domain: http(s)://sub.test.com")
    deleteAliveDomainCmd.Flags().StringP("rootdomain", "r", "", "Root domain: test.com")
    deleteAliveDomainCmd.MarkFlagRequired("alivedomain")
    deleteAliveDomainCmd.MarkFlagRequired("rootdomain")

}

func main() {
    db.ConnectDB()
	if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

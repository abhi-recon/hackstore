// db/root_domains.go

package db

import (
    "database/sql"
    "log"
)

type RootDomain struct {
    RootDomain  string
    ProgramName string
    CreatedAt   sql.NullTime
}

func GetAllRootDomains() ([]RootDomain, error) {
    var rootDomains []RootDomain

    // Assuming `db` is your database connection
    rows, err := db.Query("SELECT root_domain FROM root_domains")
    if err != nil {
        log.Println(err)
        return []RootDomain{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var rd RootDomain
        if err := rows.Scan(&rd.RootDomain); err != nil {
            log.Println(err)
            return []RootDomain{}, err
        }
        rootDomains = append(rootDomains, rd)
    }
    return rootDomains, nil
}


func GetRootDomainsByProgram(programName string) ([]RootDomain, error) {
    var rootDomains []RootDomain
    rows, err := db.Query("SELECT root_domain FROM root_domains WHERE program_name = ?", programName)
    if err != nil {
        log.Println(err)
        return []RootDomain{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var rd RootDomain
        if err := rows.Scan(&rd.RootDomain); err != nil {
            log.Println(err)
            return []RootDomain{}, err
        }
        rootDomains = append(rootDomains, rd)
    }
    return rootDomains, nil
}



// Add other CRUD operations for Root Domains table (Create, Update, Delete)
func AddRootDomain(programName, rootDomain string) error {
    _, err := db.Exec("INSERT INTO root_domains (root_domain, program_name) VALUES (?, ?)", rootDomain, programName)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func DeleteRootDomain(rootDomain, programName string) error {
    tx, err := db.Begin()
    if err != nil {
        log.Println(err)
        return err
    }

    // Deleting entries from 'subdomains' table associated with the root domain and program
    _, err = tx.Exec("DELETE FROM subdomains WHERE root_domain = ? AND root_domain IN (SELECT root_domain FROM root_domains WHERE program_name = ?)", rootDomain, programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    // Deleting entries from 'alive_domains' table associated with the root domain and program
    _, err = tx.Exec("DELETE FROM alive_domains WHERE root_domain = ? AND root_domain IN (SELECT root_domain FROM root_domains WHERE program_name = ?)", rootDomain, programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    // Deleting the root domain entry from the 'root_domains' table
    _, err = tx.Exec("DELETE FROM root_domains WHERE root_domain = ? AND program_name = ?", rootDomain, programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    if err := tx.Commit(); err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    return nil
}
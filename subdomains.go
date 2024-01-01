// db/subdomains.go

package db

import (
    "database/sql"
    "log"
)

type Subdomain struct {
    Subdomain   string
    RootDomain  string
    CreatedAt   sql.NullTime
}


// db/subdomains.go
// db/subdomains.go

func GetSubdomainsByRootDomain(rootDomain string) ([]string, error) {
    var subdomains []string

    rows, err := db.Query("SELECT subdomain FROM subdomains WHERE root_domain = ?", rootDomain)
    if err != nil {
        log.Println(err)
        return []string{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var subdomain string
        if err := rows.Scan(&subdomain); err != nil {
            log.Println(err)
            return []string{}, err
        }
        subdomains = append(subdomains, subdomain)
    }

    return subdomains, nil
}


func AddSubdomain(rootDomain, subdomain string) error {
    _, err := db.Exec("INSERT INTO subdomains (subdomain, root_domain) VALUES (?, ?)", subdomain, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}


// db/subdomains.go

func DeleteSubdomain(subdomain, rootDomain string) error {
    _, err := db.Exec("DELETE FROM subdomains WHERE subdomain = ? AND root_domain = ?", subdomain, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

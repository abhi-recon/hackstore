// db/alive_domains.go

package db

import (
    "database/sql"
    "log"
)

type AliveDomain struct {
    AliveDomain string
    RootDomain  string
    CreatedAt   sql.NullTime
}

// db/alive_domains.go
func GetAllAliveDomains() ([]string, error) {
    var aliveDomains []string

    rows, err := db.Query("SELECT alive_domain FROM alive_domains")
    if err != nil {
        log.Println(err)
        return []string{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var aliveDomain string
        if err := rows.Scan(&aliveDomain); err != nil {
            log.Println(err)
            return []string{}, err
        }
        aliveDomains = append(aliveDomains, aliveDomain)
    }

    return aliveDomains, nil
}

func GetAliveDomainsByRootDomain(rootDomain string) ([]string, error) {
    var aliveDomains []string

    rows, err := db.Query("SELECT alive_domain FROM alive_domains WHERE root_domain = ?", rootDomain)
    if err != nil {
        log.Println(err)
        return []string{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var aliveDomain string
        if err := rows.Scan(&aliveDomain); err != nil {
            log.Println(err)
            return []string{}, err
        }
        aliveDomains = append(aliveDomains, aliveDomain)
    }

    return aliveDomains, nil
}


func AddAliveDomain(rootDomain, aliveDomain string) error {
    _, err := db.Exec("INSERT INTO alive_domains (alive_domain, root_domain) VALUES (?, ?)", aliveDomain, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}


// db/alive_domains.go

func DeleteAliveDomain(aliveDomain, rootDomain string) error {
    _, err := db.Exec("DELETE FROM alive_domains WHERE alive_domain = ? AND root_domain = ?", aliveDomain, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}



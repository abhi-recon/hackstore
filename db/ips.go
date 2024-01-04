package db

import (
	"log"
)

// db/ips.go

func GetAllIPs() ([]string, error) {
    var ips []string

    rows, err := db.Query("SELECT ip FROM ips")
    if err != nil {
        log.Println(err)
        return []string{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var ip string
        if err := rows.Scan(&ip); err != nil {
            log.Println(err)
            return []string{}, err
        }
        ips = append(ips, ip)
    }

    return ips, nil
}

func GetIPsByRootDomain(rootDomain string) ([]string, error) {
    var ips []string

    rows, err := db.Query("SELECT ip FROM ips WHERE root_domain = ?", rootDomain)
    if err != nil {
        log.Println(err)
        return []string{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var ip string
        if err := rows.Scan(&ip); err != nil {
            log.Println(err)
            return []string{}, err
        }
        ips = append(ips, ip)
    }

    return ips, nil
}

func AddIP(ip, rootDomain string) error {
    _, err := db.Exec("INSERT INTO ips (ip, root_domain) VALUES (?, ?)", ip, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func DeleteIP(ip, rootDomain string) error {
    _, err := db.Exec("DELETE FROM ips WHERE ip = ? AND root_domain = ?", ip, rootDomain)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

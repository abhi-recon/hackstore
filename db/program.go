// db/program.go

package db

import (
    "database/sql"
    "log"
)

type Program struct {
    ProgramName string
    CreatedAt   sql.NullTime
}

func GetAllPrograms() ([]Program, error) {
    var programs []Program

    rows, err := db.Query("SELECT program_name FROM program")
    if err != nil {
        log.Println(err)
        return []Program{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var program Program
        if err := rows.Scan(&program.ProgramName); err != nil {
            log.Println(err)
            return []Program{}, err
        }
        programs = append(programs, program)
    }

    if err := rows.Err(); err != nil {
        log.Println(err)
        return []Program{}, err
    }

    return programs, nil
}


// Add other CRUD operations for Program table (Create, Update, Delete)
func CreateProgram(programName string) error {
    _, err := db.Exec("INSERT INTO program (program_name) VALUES (?)", programName)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

// db/program.go

func DeleteProgram(programName string) error {
    tx, err := db.Begin()
    if err != nil {
        log.Println(err)
        return err
    }

    // Deleting entries from 'subdomains' table where 'root_domain' is associated with the program
    _, err = tx.Exec("DELETE FROM subdomains WHERE root_domain IN (SELECT root_domain FROM root_domains WHERE program_name = ?)", programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    // Deleting entries from 'alive_domains' table where 'root_domain' is associated with the program
    _, err = tx.Exec("DELETE FROM alive_domains WHERE root_domain IN (SELECT root_domain FROM root_domains WHERE program_name = ?)", programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    // Deleting entries from 'root_domains' table associated with the program
    _, err = tx.Exec("DELETE FROM root_domains WHERE program_name = ?", programName)
    if err != nil {
        tx.Rollback()
        log.Println(err)
        return err
    }

    // Finally, delete the program itself from the 'program' table
    _, err = tx.Exec("DELETE FROM program WHERE program_name = ?", programName)
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

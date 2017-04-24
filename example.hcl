# This contains the connection information to the database.
connection {
    # The supported databases currently are:
    #  - mysql, sqlite3, postgres, mssql
    driver = "mysql"

    # The data source name used to connect to the database. When using the
    # mysql driver, you MUST include `parseTime=True` if you wish for time to
    # be parsed correctly.
    dsn = "user:pass@tcp(localhost:3306)/database_name?charset=utf8&parseTime=True"
}

# All of these tables will be truncated
truncate = [
    "table_a",
    "table_b",
    "table_c",
]

# We're making changes to the table `table_d`
table "table_d" {
    where {
        column_name_a = "xyz"
        column_name_b = "abc"
    }

    update {
        column_name_c = "new value for column"
    }
}

# We're making changes to the table `table_d`
table "table_d" {
    where {
        column_name_a = "fff"
        column_name_b = "ccc"
    }

    update {
        # The column `column_name_e` will have a NULL value instead of this
        # "NULL" string.
        column_name_e = "NULL"
    }
}

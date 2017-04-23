connection {
    driver = "mysql"
    dsn = "user:pass@tcp(localhost:3306)/database_name?charset=utf8&parseTime=True"
}

truncate = [
    "table_a",
    "table_b",
    "table_c",
]

table "table_d" {
    where {
        column_name_a = "xyz"
        column_name_b = "abc"
    }

    update {
        column_name_c = "new value for column"
    }
}

table "table_d" {
    where {
        column_name_a = "fff"
        column_name_b = "ccc"
    }

    update {
        column_name_e = "new new new!!"
    }
}

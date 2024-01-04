CREATE TABLE IF NOT EXISTS mahasiswa (
    nim         VARCHAR(100) PRIMARY KEY,
    email       VARCHAR(150) NOT NULL UNIQUE,
    password    VARCHAR(250) NOT NULL,
    nama        VARCHAR(150) NOT NULL,
    jurusan     VARCHAR(100) NOT NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB
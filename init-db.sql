-- init-db.sql


-- Предоставление пользователю gorm прав на базу данных mydb
GRANT ALL PRIVILEGES ON DATABASE mydb TO gorm;

-- Предоставление пользователю gorm права на использование схемы public
GRANT USAGE ON SCHEMA public TO gorm;

-- Предоставление пользователю gorm права на все таблицы в схеме public
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO gorm;

-- Предоставление пользователю gorm права на все последовательности в схеме public
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO gorm;

-- Настройка прав по умолчанию для новых объектов в схеме public
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO gorm;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO gorm;
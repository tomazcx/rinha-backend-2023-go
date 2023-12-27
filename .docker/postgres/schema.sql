CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE OR REPLACE FUNCTION generate_searchable(nome VARCHAR, apelido VARCHAR, stack TEXT) RETURNS TEXT AS $$ BEGIN
	RETURN nome || apelido || stack;
END; $$ LANGUAGE plpgsql IMMUTABLE;		

CREATE TABLE 
	IF NOT EXISTS pessoa (
		id UUID PRIMARY KEY,
		apelido VARCHAR(32) UNIQUE NOT NULL,
		nome VARCHAR(100) NOT NULL,
		nascimento DATE NOT NULL, 
		stack TEXT,
		searchable TEXT GENERATED ALWAYS AS (generate_searchable(nome, apelido, stack)) STORED
	);

CREATE INDEX 
	CONCURRENTLY IF NOT EXISTS idx_pessoa_trigram ON public.pessoa USING gist (
		searchable gist_trgm_ops(siglen = 64)
	);

CREATE UNIQUE INDEX IF NOT EXISTS idx_pessoa_apelido ON public.pessoa (apelido);

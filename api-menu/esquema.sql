-- Role: titan

-- DROP ROLE titan;

CREATE ROLE titan LOGIN
  ENCRYPTED PASSWORD 'md5e836cea17be4bcb609f7005def7c435e'
  NOSUPERUSER INHERIT CREATEDB NOCREATEROLE NOREPLICATION;
COMMENT ON ROLE titan IS 'Usuario para acceder a la base de datos de TITAN';

-- Database: titan_db

-- DROP DATABASE titan_db;

CREATE DATABASE titan_db
  WITH OWNER = titan
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'es_CO.UTF-8'
       LC_CTYPE = 'es_CO.UTF-8'
       CONNECTION LIMIT = -1;

COMMENT ON DATABASE titan_db
  IS 'Base de datos para la aplicación TITAN';


-- Table: perfil

-- DROP TABLE perfil;

CREATE TABLE perfil
(
  id integer NOT NULL,
  nombre character varying(50) NOT NULL,
  dominio character varying NOT NULL, -- Dominio al cual pertenece el perfil de usuario
  CONSTRAINT perfil_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE perfil
  OWNER TO titan;
COMMENT ON COLUMN perfil.dominio IS 'Dominio al cual pertenece el perfil de usuario';


-- Table: menu_opcion

-- DROP TABLE menu_opcion;

CREATE TABLE menu_opcion
(
  id integer NOT NULL,
  nombre character varying(60) NOT NULL DEFAULT ''::character varying,
  variable character varying(50) NOT NULL DEFAULT ''::character varying,
  url character varying(250) NOT NULL DEFAULT ''::character varying,
  padre integer NOT NULL DEFAULT 0,
  orden integer NOT NULL DEFAULT 0,
  layout character varying(60) DEFAULT ''::character varying,
  dominio character varying NOT NULL, -- Dominio al cual pertenecen los menús
  CONSTRAINT "PF_menu_opcion" PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE menu_opcion
  OWNER TO titan;
COMMENT ON COLUMN menu_opcion.dominio IS 'Dominio al cual pertenecen los menús';



-- Table: perfil_x_menu_opcion

-- DROP TABLE perfil_x_menu_opcion;

CREATE TABLE perfil_x_menu_opcion
(
  perfil integer NOT NULL,
  opcion integer NOT NULL,
  pxo_nivel integer NOT NULL,
  CONSTRAINT perfil_x_menu_opcion_pkey PRIMARY KEY (perfil, opcion),
  CONSTRAINT "FK_perfil_x_menu_opcion_menu" FOREIGN KEY (opcion)
      REFERENCES menu_opcion (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "FK_perfil_x_menu_opcion_perfil" FOREIGN KEY (perfil)
      REFERENCES perfil (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE RESTRICT
)
WITH (
  OIDS=FALSE
);
ALTER TABLE perfil_x_menu_opcion
  OWNER TO titan;



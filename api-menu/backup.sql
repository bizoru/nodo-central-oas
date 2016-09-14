--
-- PostgreSQL database dump
--

-- Dumped from database version 9.2.14
-- Dumped by pg_dump version 9.2.15
-- Started on 2016-06-21 11:05:58 COT

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2857 (class 0 OID 26664)
-- Dependencies: 169
-- Data for Name: menu_opcion; Type: TABLE DATA; Schema: public; Owner: titan
--

COPY menu_opcion (id, nombre, variable, url, padre, orden, layout, dominio) FROM stdin;
1	Parámetros	 	 	0	9		titan.com
2	Gestionar Personas	 	 	0	2		titan.com
3	Solicitud de Registro	 	 	2	1		titan.com
4	Hoja de Vida	 	 	2	2		titan.com
5	Vinculación	 	 	2	3		titan.com
6	Usuario	 	 	0	1		titan.com
7	Cerrar Sesión	 	 	6	1		titan.com
\.


--
-- TOC entry 2858 (class 0 OID 26681)
-- Dependencies: 170
-- Data for Name: perfil; Type: TABLE DATA; Schema: public; Owner: titan
--

COPY perfil (id, nombre, dominio) FROM stdin;
1	administrador	titan.com
4	administrador	titan.com
3	asistente_rrhh	titan.com
2	asistente_contabilidad	titan.com
\.


--
-- TOC entry 2859 (class 0 OID 26686)
-- Dependencies: 171
-- Data for Name: perfil_x_menu_opcion; Type: TABLE DATA; Schema: public; Owner: titan
--

COPY perfil_x_menu_opcion (perfil, opcion, pxo_nivel) FROM stdin;
1	1	1
1	2	1
1	3	1
1	4	1
1	5	1
1	6	1
1	7	1
2	1	1
2	6	1
2	7	1
3	2	1
3	3	1
3	4	1
3	5	1
3	6	1
3	7	1
\.


-- Completed on 2016-06-21 11:05:58 COT

--
-- PostgreSQL database dump complete
--


--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2
-- Dumped by pg_dump version 11.2

-- Started on 2019-06-21 07:37:22

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 213 (class 1259 OID 67239)
-- Name: users; Type: TABLE; Schema: public; Owner: ylbadmin
--

CREATE TABLE public.users (
    id BIGINT PRIMARY KEY,
    province_agent_id bigint NOT NULL,
    city_agent_id bigint NOT NULL,
    county_agent_id bigint NOT NULL,
    name character varying(80),
    phone character varying(22),
    points integer DEFAULT 0,
    biz_level smallint DEFAULT 1,
    role smallint DEFAULT 100,
    grade smallint DEFAULT 5,
    ref_type smallint DEFAULT 1,
    ref_id bigint DEFAULT 0,
    geog public.geography(Point,4326)
);


ALTER TABLE public.users OWNER TO ylbadmin;

--
-- TOC entry 212 (class 1259 OID 67237)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: ylbadmin
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO ylbadmin;

--
-- TOC entry 4259 (class 0 OID 0)
-- Dependencies: 212
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ylbadmin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4120 (class 2604 OID 67254)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: ylbadmin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4253 (class 0 OID 67239)
-- Dependencies: 213
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ylbadmin
--

COPY public.users (id, province_agent_id, city_agent_id, county_agent_id, name, phone, points, biz_level, role, grade, ref_type, ref_id, geog) FROM stdin;
1	41	371	2	测试地点	\N	0	1	100	5	1	0	0101000020E6100000618C48145A6F5C40984EEB36A82D4140
2	41	371	2	新郑市公安局	\N	0	1	100	5	1	0	0101000020E61000001A8A3BDEE4705C40B935E9B644324140
3	41	371	2	新郑市人民法院	\N	0	1	100	5	1	0	0101000020E610000013F3ACA415705C40DD9733DB15324140
4	41	371	2	新郑市人民检察院	\N	0	1	100	5	1	0	0101000020E61000001361C3D32B6E5C40B98B3045B9344140
5	41	371	3	新郑机场	\N	0	1	100	5	1	0	0101000020E61000007D410B0918775C4059A2B3CC22444140
6	41	371	1	郑州大学	\N	0	1	100	5	1	0	0101000020E61000008198840B79625C40B58AFED0CC694140
7	41	372	1	安阳市	\N	0	1	100	5	1	0	0101000020E61000007EC7F0D8CF995C404D2F3196E90D4240
8	42	270	1	武汉市	\N	0	1	100	5	1	0	0101000020E610000060014C1938945C4076340EF5BB983E40
1101	41	371	2	bbb	13800138000	0	2	100	0	1	0	0101000020E61000006956B60F79625C40B58AFED0CC694140
1102	41	371	2	ccc	13800138000	0	2	100	0	1	0	0101000020E61000006956B60F79625C40B58AFED0CC694140
\.


--
-- TOC entry 4260 (class 0 OID 0)
-- Dependencies: 212
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ylbadmin
--

SELECT pg_catalog.setval('public.users_id_seq', 999, true);


--
-- TOC entry 4123 (class 2606 OID 67253)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ylbadmin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4121 (class 1259 OID 67255)
-- Name: users_geog_idx; Type: INDEX; Schema: public; Owner: ylbadmin
--

CREATE INDEX users_geog_idx ON public.users USING gist (geog);


-- Completed on 2019-06-21 07:37:23

--
-- PostgreSQL database dump complete
--


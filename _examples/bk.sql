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
-- Name: user; Type: TABLE; Schema: public; Owner: geopoint_test_user
--

CREATE TABLE public.user (
    id BIGINT PRIMARY KEY,
    label character varying(64),
    geog public.geography(Point,4326)
);


ALTER TABLE public.user OWNER TO geopoint_test_user;

--
-- TOC entry 212 (class 1259 OID 67237)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: geopoint_test_user
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO geopoint_test_user;

--
-- TOC entry 4259 (class 0 OID 0)
-- Dependencies: 212
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: geopoint_test_user
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.user.id;


--
-- TOC entry 4120 (class 2604 OID 67254)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: geopoint_test_user
--

ALTER TABLE ONLY public.user ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 4253 (class 0 OID 67239)
-- Dependencies: 213
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: geopoint_test_user
--

COPY public.user (id,  label, geog) FROM stdin;
1	测试地点	0101000020E6100000618C48145A6F5C40984EEB36A82D4140
2	新郑市公安局	0101000020E61000001A8A3BDEE4705C40B935E9B644324140
3	新郑市人民法院	0101000020E610000013F3ACA415705C40DD9733DB15324140
4	新郑市人民检察院	0101000020E61000001361C3D32B6E5C40B98B3045B9344140
5	新郑机场	0101000020E61000007D410B0918775C4059A2B3CC22444140
6	郑州大学	0101000020E61000008198840B79625C40B58AFED0CC694140
7	安阳市	0101000020E61000007EC7F0D8CF995C404D2F3196E90D4240
8	武汉市	0101000020E610000060014C1938945C4076340EF5BB983E40
1101	bbb	0101000020E61000006956B60F79625C40B58AFED0CC694140
1102	ccc	0101000020E61000006956B60F79625C40B58AFED0CC694140
\.


--
-- TOC entry 4260 (class 0 OID 0)
-- Dependencies: 212
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: geopoint_test_user
--

SELECT pg_catalog.setval('public.user_id_seq', 999, true);


--
-- TOC entry 4123 (class 2606 OID 67253)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: geopoint_test_user
--

ALTER TABLE ONLY public.user
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 4121 (class 1259 OID 67255)
-- Name: user_geog_idx; Type: INDEX; Schema: public; Owner: geopoint_test_user
--

CREATE INDEX user_geog_idx ON public.user USING gist (geog);


-- Completed on 2019-06-21 07:37:23

--
-- PostgreSQL database dump complete
--


--
-- PostgreSQL database dump
--

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
-- TOC entry 197 (class 1259 OID 70691)
-- Name: accounts; Type: TABLE; Schema: public; Owner: artur0us
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    username character varying(96) NOT NULL,
    password character varying(96) NOT NULL,
    created_at bigint NOT NULL
);


ALTER TABLE public.accounts OWNER TO artur0us;

--
-- TOC entry 203 (class 1259 OID 70721)
-- Name: accounts_groups; Type: TABLE; Schema: public; Owner: artur0us
--

CREATE TABLE public.accounts_groups (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.accounts_groups OWNER TO artur0us;

--
-- TOC entry 202 (class 1259 OID 70719)
-- Name: accounts_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: artur0us
--

CREATE SEQUENCE public.accounts_groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_groups_id_seq OWNER TO artur0us;

--
-- TOC entry 2965 (class 0 OID 0)
-- Dependencies: 202
-- Name: accounts_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artur0us
--

ALTER SEQUENCE public.accounts_groups_id_seq OWNED BY public.accounts_groups.id;


--
-- TOC entry 196 (class 1259 OID 70689)
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: artur0us
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO artur0us;

--
-- TOC entry 2966 (class 0 OID 0)
-- Dependencies: 196
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artur0us
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- TOC entry 201 (class 1259 OID 70710)
-- Name: accounts_info; Type: TABLE; Schema: public; Owner: artur0us
--

CREATE TABLE public.accounts_info (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    last_name text NOT NULL,
    first_name text NOT NULL,
    middle_name text NOT NULL,
    account_group_id integer NOT NULL
);


ALTER TABLE public.accounts_info OWNER TO artur0us;

--
-- TOC entry 200 (class 1259 OID 70708)
-- Name: accounts_info_id_seq; Type: SEQUENCE; Schema: public; Owner: artur0us
--

CREATE SEQUENCE public.accounts_info_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_info_id_seq OWNER TO artur0us;

--
-- TOC entry 2967 (class 0 OID 0)
-- Dependencies: 200
-- Name: accounts_info_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artur0us
--

ALTER SEQUENCE public.accounts_info_id_seq OWNED BY public.accounts_info.id;


--
-- TOC entry 199 (class 1259 OID 70699)
-- Name: accounts_sessions; Type: TABLE; Schema: public; Owner: artur0us
--

CREATE TABLE public.accounts_sessions (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    token text NOT NULL,
    created_at bigint NOT NULL,
    expires_at bigint NOT NULL,
    user_agent_info text NOT NULL
);


ALTER TABLE public.accounts_sessions OWNER TO artur0us;

--
-- TOC entry 198 (class 1259 OID 70697)
-- Name: accounts_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: artur0us
--

CREATE SEQUENCE public.accounts_sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_sessions_id_seq OWNER TO artur0us;

--
-- TOC entry 2968 (class 0 OID 0)
-- Dependencies: 198
-- Name: accounts_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artur0us
--

ALTER SEQUENCE public.accounts_sessions_id_seq OWNED BY public.accounts_sessions.id;


--
-- TOC entry 205 (class 1259 OID 70732)
-- Name: notes; Type: TABLE; Schema: public; Owner: artur0us
--

CREATE TABLE public.notes (
    id bigint NOT NULL,
    author_account_id bigint NOT NULL,
    created_at bigint NOT NULL,
    title text NOT NULL,
    body text NOT NULL
);


ALTER TABLE public.notes OWNER TO artur0us;

--
-- TOC entry 204 (class 1259 OID 70730)
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: artur0us
--

CREATE SEQUENCE public.notes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notes_id_seq OWNER TO artur0us;

--
-- TOC entry 2969 (class 0 OID 0)
-- Dependencies: 204
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artur0us
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- TOC entry 2814 (class 2604 OID 70694)
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- TOC entry 2817 (class 2604 OID 70724)
-- Name: accounts_groups id; Type: DEFAULT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_groups ALTER COLUMN id SET DEFAULT nextval('public.accounts_groups_id_seq'::regclass);


--
-- TOC entry 2816 (class 2604 OID 70713)
-- Name: accounts_info id; Type: DEFAULT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_info ALTER COLUMN id SET DEFAULT nextval('public.accounts_info_id_seq'::regclass);


--
-- TOC entry 2815 (class 2604 OID 70702)
-- Name: accounts_sessions id; Type: DEFAULT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_sessions ALTER COLUMN id SET DEFAULT nextval('public.accounts_sessions_id_seq'::regclass);


--
-- TOC entry 2818 (class 2604 OID 70735)
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- TOC entry 2951 (class 0 OID 70691)
-- Dependencies: 197
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: artur0us
--

COPY public.accounts (id, username, password, created_at) FROM stdin;
1	artur0us	65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5	15555
2	user	65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5	15556
\.


--
-- TOC entry 2957 (class 0 OID 70721)
-- Dependencies: 203
-- Data for Name: accounts_groups; Type: TABLE DATA; Schema: public; Owner: artur0us
--

COPY public.accounts_groups (id, name) FROM stdin;
1	Admin
2	User
\.


--
-- TOC entry 2955 (class 0 OID 70710)
-- Dependencies: 201
-- Data for Name: accounts_info; Type: TABLE DATA; Schema: public; Owner: artur0us
--

COPY public.accounts_info (id, account_id, last_name, first_name, middle_name, account_group_id) FROM stdin;
1	1	Arturov	Artur	Arturovich	1
2	2	User	User	User	2
\.


--
-- TOC entry 2953 (class 0 OID 70699)
-- Dependencies: 199
-- Data for Name: accounts_sessions; Type: TABLE DATA; Schema: public; Owner: artur0us
--

COPY public.accounts_sessions (id, account_id, token, created_at, expires_at, user_agent_info) FROM stdin;
1	1	7b508ce6-f00c-4a08-8e24-5c5e039b2335	1572519971	1572919972	some info about client
10	2	1781b4bf-b347-4f90-a0a9-aba83503fc72	1572528266	1573737866	Some info about client
\.


--
-- TOC entry 2959 (class 0 OID 70732)
-- Dependencies: 205
-- Data for Name: notes; Type: TABLE DATA; Schema: public; Owner: artur0us
--

COPY public.notes (id, author_account_id, created_at, title, body) FROM stdin;
2	1	1572527493	First note	Interesting fact about GoLang...
4	1	1572527497	First note	Interesting fact about GoLang...
5	2	1572528274	First note	Interesting fact about GoLang...
\.


--
-- TOC entry 2970 (class 0 OID 0)
-- Dependencies: 202
-- Name: accounts_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: artur0us
--

SELECT pg_catalog.setval('public.accounts_groups_id_seq', 1, false);


--
-- TOC entry 2971 (class 0 OID 0)
-- Dependencies: 196
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: artur0us
--

SELECT pg_catalog.setval('public.accounts_id_seq', 1, false);


--
-- TOC entry 2972 (class 0 OID 0)
-- Dependencies: 200
-- Name: accounts_info_id_seq; Type: SEQUENCE SET; Schema: public; Owner: artur0us
--

SELECT pg_catalog.setval('public.accounts_info_id_seq', 1, false);


--
-- TOC entry 2973 (class 0 OID 0)
-- Dependencies: 198
-- Name: accounts_sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: artur0us
--

SELECT pg_catalog.setval('public.accounts_sessions_id_seq', 10, true);


--
-- TOC entry 2974 (class 0 OID 0)
-- Dependencies: 204
-- Name: notes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: artur0us
--

SELECT pg_catalog.setval('public.notes_id_seq', 5, true);


--
-- TOC entry 2826 (class 2606 OID 70729)
-- Name: accounts_groups accounts_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_groups
    ADD CONSTRAINT accounts_groups_pkey PRIMARY KEY (id);


--
-- TOC entry 2824 (class 2606 OID 70718)
-- Name: accounts_info accounts_info_pkey; Type: CONSTRAINT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_info
    ADD CONSTRAINT accounts_info_pkey PRIMARY KEY (id);


--
-- TOC entry 2820 (class 2606 OID 70696)
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- TOC entry 2822 (class 2606 OID 70707)
-- Name: accounts_sessions accounts_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.accounts_sessions
    ADD CONSTRAINT accounts_sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2828 (class 2606 OID 70740)
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: artur0us
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


-- Completed on 2019-10-31 16:25:32

--
-- PostgreSQL database dump complete
--


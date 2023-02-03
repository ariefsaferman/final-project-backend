--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2023-02-03 09:58:41 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 242 (class 1259 OID 33550)
-- Name: admin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.admin (
    id integer NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    role_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.admin OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 33549)
-- Name: admin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.admin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admin_id_seq OWNER TO postgres;

--
-- TOC entry 3554 (class 0 OID 0)
-- Dependencies: 241
-- Name: admin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.admin_id_seq OWNED BY public.admin.id;


--
-- TOC entry 229 (class 1259 OID 33353)
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    house_id integer NOT NULL,
    reservation_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 33352)
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bookings_id_seq OWNER TO postgres;

--
-- TOC entry 3555 (class 0 OID 0)
-- Dependencies: 228
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- TOC entry 221 (class 1259 OID 33317)
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 33316)
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cities_id_seq OWNER TO postgres;

--
-- TOC entry 3556 (class 0 OID 0)
-- Dependencies: 220
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- TOC entry 217 (class 1259 OID 33299)
-- Name: games_chances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.games_chances (
    id integer NOT NULL,
    user_id integer NOT NULL,
    chances integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    number_of_play integer,
    CONSTRAINT chances_nonnegative CHECK ((chances >= 0))
);


ALTER TABLE public.games_chances OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 33298)
-- Name: games_chances_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.games_chances_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.games_chances_id_seq OWNER TO postgres;

--
-- TOC entry 3557 (class 0 OID 0)
-- Dependencies: 216
-- Name: games_chances_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.games_chances_id_seq OWNED BY public.games_chances.id;


--
-- TOC entry 227 (class 1259 OID 33344)
-- Name: houses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.houses (
    id integer NOT NULL,
    name character varying NOT NULL,
    user_id integer NOT NULL,
    price_per_night integer NOT NULL,
    description character varying NOT NULL,
    city_id integer NOT NULL,
    max_guest integer,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.houses OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 33343)
-- Name: houses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.houses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_id_seq OWNER TO postgres;

--
-- TOC entry 3558 (class 0 OID 0)
-- Dependencies: 226
-- Name: houses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.houses_id_seq OWNED BY public.houses.id;


--
-- TOC entry 225 (class 1259 OID 33335)
-- Name: houses_photos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.houses_photos (
    id integer NOT NULL,
    house_id integer NOT NULL,
    photo_url character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.houses_photos OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 33334)
-- Name: houses_photos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.houses_photos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_photos_id_seq OWNER TO postgres;

--
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 224
-- Name: houses_photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.houses_photos_id_seq OWNED BY public.houses_photos.id;


--
-- TOC entry 235 (class 1259 OID 33378)
-- Name: pickups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    reservation_id integer NOT NULL,
    pickup_status_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.pickups OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 33377)
-- Name: pickup_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickup_id_seq OWNER TO postgres;

--
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 234
-- Name: pickup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickup_id_seq OWNED BY public.pickups.id;


--
-- TOC entry 237 (class 1259 OID 33387)
-- Name: pickups_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickups_statuses (
    id integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.pickups_statuses OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 33386)
-- Name: pickup_status_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickup_status_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickup_status_id_seq OWNER TO postgres;

--
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 236
-- Name: pickup_status_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickup_status_id_seq OWNED BY public.pickups_statuses.id;


--
-- TOC entry 233 (class 1259 OID 33369)
-- Name: reservations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservations (
    id integer NOT NULL,
    house_id integer NOT NULL,
    user_id integer NOT NULL,
    check_in_date date NOT NULL,
    check_out_date date NOT NULL,
    total_price numeric NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.reservations OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 33368)
-- Name: reservations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reservations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservations_id_seq OWNER TO postgres;

--
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 232
-- Name: reservations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reservations_id_seq OWNED BY public.reservations.id;


--
-- TOC entry 219 (class 1259 OID 33306)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 33305)
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO postgres;

--
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 218
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- TOC entry 239 (class 1259 OID 33396)
-- Name: source_of_funds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.source_of_funds (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.source_of_funds OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 33395)
-- Name: source_of_funds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.source_of_funds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.source_of_funds_id_seq OWNER TO postgres;

--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 238
-- Name: source_of_funds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.source_of_funds_id_seq OWNED BY public.source_of_funds.id;


--
-- TOC entry 231 (class 1259 OID 33362)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    source_of_funds_id integer NOT NULL,
    user_id integer,
    balance numeric NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 33361)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 230
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 215 (class 1259 OID 33288)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    full_name character varying NOT NULL,
    address character varying NOT NULL,
    city_id integer NOT NULL,
    role_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 33287)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 240 (class 1259 OID 33507)
-- Name: walletsequence; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.walletsequence
    START WITH 700001
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.walletsequence OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 33328)
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id integer DEFAULT nextval('public.walletsequence'::regclass) NOT NULL,
    user_id integer NOT NULL,
    balance numeric NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 33327)
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_id_seq OWNER TO postgres;

--
-- TOC entry 3567 (class 0 OID 0)
-- Dependencies: 222
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- TOC entry 3321 (class 2604 OID 33553)
-- Name: admin id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admin ALTER COLUMN id SET DEFAULT nextval('public.admin_id_seq'::regclass);


--
-- TOC entry 3303 (class 2604 OID 33356)
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- TOC entry 3291 (class 2604 OID 33320)
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- TOC entry 3285 (class 2604 OID 33302)
-- Name: games_chances id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chances ALTER COLUMN id SET DEFAULT nextval('public.games_chances_id_seq'::regclass);


--
-- TOC entry 3300 (class 2604 OID 33347)
-- Name: houses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses ALTER COLUMN id SET DEFAULT nextval('public.houses_id_seq'::regclass);


--
-- TOC entry 3297 (class 2604 OID 33338)
-- Name: houses_photos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos ALTER COLUMN id SET DEFAULT nextval('public.houses_photos_id_seq'::regclass);


--
-- TOC entry 3312 (class 2604 OID 33381)
-- Name: pickups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups ALTER COLUMN id SET DEFAULT nextval('public.pickup_id_seq'::regclass);


--
-- TOC entry 3315 (class 2604 OID 33390)
-- Name: pickups_statuses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_statuses ALTER COLUMN id SET DEFAULT nextval('public.pickup_status_id_seq'::regclass);


--
-- TOC entry 3309 (class 2604 OID 33372)
-- Name: reservations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations ALTER COLUMN id SET DEFAULT nextval('public.reservations_id_seq'::regclass);


--
-- TOC entry 3288 (class 2604 OID 33309)
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- TOC entry 3318 (class 2604 OID 33399)
-- Name: source_of_funds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds ALTER COLUMN id SET DEFAULT nextval('public.source_of_funds_id_seq'::regclass);


--
-- TOC entry 3306 (class 2604 OID 33365)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3282 (class 2604 OID 33291)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3548 (class 0 OID 33550)
-- Dependencies: 242
-- Data for Name: admin; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.admin (id, email, password, role_id, created_at, updated_at, deleted_at) VALUES (1, 'admin@gmail.com', '$2a$12$uBb5fuv6joAcewSi9bhqk.q8wafzmMqwUy3FylzBgOudx0edHCoJu', 4, '2023-01-20 16:02:56.775114', '2023-01-20 16:02:56.775114', NULL);
INSERT INTO public.admin (id, email, password, role_id, created_at, updated_at, deleted_at) VALUES (2, 'adminBaru@gmail.com', '$2a$04$Imq6d5LZoBlqE3cCwMkWeOflG9z6ILsjrQ1uoQkbuRxFXnoOm0PlC', 1, '2023-01-20 17:08:19.193626', '2023-01-20 17:08:19.193626', NULL);


--
-- TOC entry 3535 (class 0 OID 33353)
-- Dependencies: 229
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3527 (class 0 OID 33317)
-- Dependencies: 221
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cities (id, name, created_at, updated_at, deleted_at) VALUES (1, 'jakarta', '2023-01-18 09:28:29.99152', '2023-01-18 09:28:29.99152', NULL);
INSERT INTO public.cities (id, name, created_at, updated_at, deleted_at) VALUES (2, 'bogor', '2023-01-18 09:28:29.99152', '2023-01-18 09:28:29.99152', NULL);
INSERT INTO public.cities (id, name, created_at, updated_at, deleted_at) VALUES (3, 'depok', '2023-01-18 09:28:29.99152', '2023-01-18 09:28:29.99152', NULL);
INSERT INTO public.cities (id, name, created_at, updated_at, deleted_at) VALUES (4, 'tanggerang', '2023-01-18 09:28:29.99152', '2023-01-18 09:28:29.99152', NULL);
INSERT INTO public.cities (id, name, created_at, updated_at, deleted_at) VALUES (5, 'bekasi', '2023-01-18 09:28:29.99152', '2023-01-18 09:28:29.99152', NULL);


--
-- TOC entry 3523 (class 0 OID 33299)
-- Dependencies: 217
-- Data for Name: games_chances; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (1, 19, 0, '2023-01-20 09:14:44.552408', '2023-01-20 09:14:44.552408', NULL, NULL);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (2, 24, 0, '2023-01-23 15:37:12.268168', '2023-01-23 15:37:12.268168', NULL, NULL);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (4, 26, 0, '2023-01-23 15:41:27.234529', '2023-01-23 15:41:27.234529', NULL, NULL);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (6, 28, 2, '2023-01-28 15:09:33.966704', '2023-01-28 15:09:33.966704', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (7, 29, 2, '2023-01-28 16:08:32.433247', '2023-01-28 16:08:32.433247', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (8, 30, 2, '2023-01-28 20:03:24.882092', '2023-01-28 20:03:24.882092', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (9, 31, 2, '2023-01-28 20:04:27.161253', '2023-01-28 20:04:27.161253', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (10, 34, 2, '2023-01-28 20:09:16.01925', '2023-01-28 20:09:16.01925', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (11, 35, 2, '2023-01-28 20:09:44.184031', '2023-01-28 20:09:44.184031', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (12, 36, 2, '2023-01-28 20:10:56.885194', '2023-01-28 20:10:56.885194', NULL, 0);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (5, 27, 8, '2023-01-23 15:43:14.176672', '2023-01-23 15:43:14.176672', NULL, 11);
INSERT INTO public.games_chances (id, user_id, chances, created_at, updated_at, deleted_at, number_of_play) VALUES (3, 25, 0, '2023-01-23 15:38:45.253093', '2023-01-23 15:38:45.253093', NULL, NULL);


--
-- TOC entry 3533 (class 0 OID 33344)
-- Dependencies: 227
-- Data for Name: houses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (5, 'house', 19, 300000, 'house of love', 1, 3, '2023-01-20 23:25:00.939536', '2023-01-20 23:25:00.939536', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (6, 'house', 19, 300000, 'house of love', 1, 3, '2023-01-20 23:28:52.574883', '2023-01-20 23:28:52.574883', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (7, 'house', 19, 300000, 'house of love', 1, 3, '2023-01-20 23:43:19.823196', '2023-01-20 23:43:19.823196', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (8, 'house', 19, 300000, 'house of love', 1, 3, '2023-01-20 23:46:44.141591', '2023-01-20 23:46:44.141591', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (11, 'house', 19, 300000, 'house of love', 1, 3, '2023-01-22 12:53:59.751069', '2023-01-22 12:53:59.751069', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (2, 'house', 19, 100000, 'house of love', 1, 3, '2023-01-20 23:05:44.241496', '2023-01-20 23:05:44.241496', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (12, 'house', 19, 300000, 'house of love', 4, 3, '2023-01-22 12:55:47.007033', '2023-01-22 12:55:47.007033', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (1, 'bali house', 1, 150000, 'house of love', 2, 3, '2023-01-20 18:22:04.804921', '2023-01-20 18:22:04.804921', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (14, 'house kecak', 2, 50000, 'house of comedy', 5, 3, '2023-01-23 12:41:00.917981', '2023-01-23 12:41:00.917981', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (15, 'house padang', 4, 500000, 'house of comedy', 4, 3, '2023-01-23 12:48:16.469612', '2023-01-23 12:48:16.469612', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (16, 'house padang', 27, 500000, 'house of comedy', 4, 3, '2023-01-23 12:49:20.576497', '2023-01-23 12:49:20.576497', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (17, 'house jakarta', 27, 1000000, 'ondel-ondel', 4, 4, '2023-01-23 12:49:30.672588', '2023-01-23 14:51:09.741704', '2023-01-23 15:24:46.71989');
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (13, 'house', 24, 300000, 'house of love', 3, 3, '2023-01-22 14:08:49.449576', '2023-01-22 14:08:49.449576', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (18, 'house singaraja', 27, 500000, 'house of happiness', 1, 2, '2023-01-31 09:55:28.259207', '2023-01-31 09:55:28.259207', NULL);
INSERT INTO public.houses (id, name, user_id, price_per_night, description, city_id, max_guest, created_at, updated_at, deleted_at) VALUES (19, 'house kebaya', 25, 100000, 'house of betawi', 3, 5, '2023-02-01 14:31:33.773179', '2023-02-01 14:31:33.773179', NULL);


--
-- TOC entry 3531 (class 0 OID 33335)
-- Dependencies: 225
-- Data for Name: houses_photos; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (2, 5, '', '2023-01-20 23:25:00.941196', '2023-01-20 23:25:00.941196', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (3, 6, '', '2023-01-20 23:28:52.577666', '2023-01-20 23:28:52.577666', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (4, 7, 'blabla.jpg', '2023-01-20 23:43:19.825228', '2023-01-20 23:43:19.825228', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (5, 8, 'blabla.jpg', '2023-01-20 23:46:44.142204', '2023-01-20 23:46:44.142204', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (8, 11, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674366840/house-cloudinary/rxe3h7250ow0zhfutr3m.png', '2023-01-22 12:53:59.753023', '2023-01-22 12:53:59.753023', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (9, 12, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674366945/house-cloudinary/wh6cjf180h7q5ubm3qwr.png', '2023-01-22 12:55:47.007818', '2023-01-22 12:55:47.007818', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (10, 12, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674366947/house-cloudinary/jajzwjo1akjvuhpevvqk.png', '2023-01-22 12:55:47.007818', '2023-01-22 12:55:47.007818', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (11, 13, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674371323/house-cloudinary/v1b1uu9mmqum2nhwouik.png', '2023-01-22 14:08:49.451611', '2023-01-22 14:08:49.451611', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (12, 13, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674371328/house-cloudinary/ghzzhk4bwjwjwhlmdm6n.png', '2023-01-22 14:08:49.451611', '2023-01-22 14:08:49.451611', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (13, 14, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674452460/house-cloudinary/ga8jwfdzl8auascujiza.png', '2023-01-23 12:41:00.91945', '2023-01-23 12:41:00.91945', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (14, 15, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674452896/house-cloudinary/dewccg4r5drcsmkbxudk.png', '2023-01-23 12:48:16.470421', '2023-01-23 12:48:16.470421', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (15, 16, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674452960/house-cloudinary/x9zxxhgnzpdxbviwh9fe.png', '2023-01-23 12:49:20.577437', '2023-01-23 12:49:20.577437', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (17, 17, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674458789/house-cloudinary/wdmwiclvaktdr4enxixg.png', '2023-01-23 14:26:29.899159', '2023-01-23 14:26:29.899159', '2023-01-23 14:40:59.369049');
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (16, 17, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674452970/house-cloudinary/buvlgzspnhfallr6pvrt.png', '2023-01-23 12:49:30.67385', '2023-01-23 12:49:30.67385', '2023-01-23 14:40:59.369049');
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (18, 17, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674459658/house-cloudinary/fclccv0lor9hzgdfjb6i.png', '2023-01-23 14:40:59.36771', '2023-01-23 14:40:59.36771', '2023-01-23 14:40:59.369049');
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (19, 17, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674460268/house-cloudinary/mzigapvixslqeljnixkk.png', '2023-01-23 14:51:09.742539', '2023-01-23 14:51:09.742539', '2023-01-23 15:24:46.718488');
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (20, 17, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1674460269/house-cloudinary/xcvwy9w5gsug8hrfn8tr.png', '2023-01-23 14:51:09.742539', '2023-01-23 14:51:09.742539', '2023-01-23 15:24:46.718488');
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (21, 18, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1675133727/house-cloudinary/z28tgvsdgphwurjkklrs.png', '2023-01-31 09:55:28.26735', '2023-01-31 09:55:28.26735', NULL);
INSERT INTO public.houses_photos (id, house_id, photo_url, created_at, updated_at, deleted_at) VALUES (22, 19, 'https://res.cloudinary.com/dfgtlhook/image/upload/v1675236693/house-cloudinary/ifzydnqlpmogr8ogtopi.jpg', '2023-02-01 14:31:33.77511', '2023-02-01 14:31:33.77511', NULL);


--
-- TOC entry 3541 (class 0 OID 33378)
-- Dependencies: 235
-- Data for Name: pickups; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.pickups (id, user_id, reservation_id, pickup_status_id, created_at, updated_at, deleted_at) VALUES (3, 27, 23, 1, '2023-01-25 12:09:54.232829', '2023-01-25 12:09:54.232829', NULL);
INSERT INTO public.pickups (id, user_id, reservation_id, pickup_status_id, created_at, updated_at, deleted_at) VALUES (10, 27, 41, 1, '2023-01-27 16:33:28.497481', '2023-01-27 16:33:28.497481', NULL);


--
-- TOC entry 3543 (class 0 OID 33387)
-- Dependencies: 237
-- Data for Name: pickups_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.pickups_statuses (id, status, created_at, updated_at, deleted_at) VALUES (1, 'Pending', '2023-01-25 11:20:37.718695', '2023-01-25 11:20:37.718695', NULL);
INSERT INTO public.pickups_statuses (id, status, created_at, updated_at, deleted_at) VALUES (2, 'Awaiting Check-in Date', '2023-01-25 11:20:56.308618', '2023-01-25 11:20:56.308618', NULL);
INSERT INTO public.pickups_statuses (id, status, created_at, updated_at, deleted_at) VALUES (3, 'On the way to pickup', '2023-01-25 11:21:22.369726', '2023-01-25 11:21:22.369726', NULL);
INSERT INTO public.pickups_statuses (id, status, created_at, updated_at, deleted_at) VALUES (4, 'On the way to reservation', '2023-01-25 11:21:40.00762', '2023-01-25 11:21:40.00762', NULL);
INSERT INTO public.pickups_statuses (id, status, created_at, updated_at, deleted_at) VALUES (5, 'Completed', '2023-01-25 11:21:51.791836', '2023-01-25 11:21:51.791836', NULL);


--
-- TOC entry 3539 (class 0 OID 33369)
-- Dependencies: 233
-- Data for Name: reservations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (13, 15, 27, '2023-01-26', '2023-01-28', 1000000, '2023-01-25 09:37:52.64658', '2023-01-25 09:37:52.64658', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (14, 15, 27, '2023-01-27', '2023-01-28', 500000, '2023-01-25 09:38:19.890401', '2023-01-25 09:38:19.890401', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (15, 15, 27, '2023-02-26', '2023-02-28', 1000000, '2023-01-25 09:53:52.606913', '2023-01-25 09:53:52.606913', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (16, 15, 27, '2023-02-27', '2023-02-28', 500000, '2023-01-25 09:54:06.210331', '2023-01-25 09:54:06.210331', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (17, 15, 27, '2023-04-25', '2023-04-27', 1000000, '2023-01-25 10:10:59.360751', '2023-01-25 10:10:59.360751', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (18, 15, 27, '2023-04-30', '2023-05-01', 500000, '2023-01-25 10:11:19.364631', '2023-01-25 10:11:19.364631', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (19, 15, 27, '2023-04-01', '2023-04-05', 2000000, '2023-01-25 10:28:58.433476', '2023-01-25 10:28:58.433476', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (20, 15, 27, '2023-04-06', '2023-04-07', 500000, '2023-01-25 10:30:54.733395', '2023-01-25 10:30:54.733395', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (23, 15, 27, '2023-06-06', '2023-06-08', 1000000, '2023-01-25 12:09:54.216703', '2023-01-25 12:09:54.216703', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (26, 15, 27, '2023-06-09', '2023-06-10', 500000, '2023-01-25 12:20:03.422747', '2023-01-25 12:20:03.422747', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (29, 13, 27, '2023-01-25', '2023-01-27', 579375, '2023-01-25 13:38:38.342278', '2023-01-25 13:38:38.342278', NULL);
INSERT INTO public.reservations (id, house_id, user_id, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) VALUES (41, 7, 27, '2023-01-27', '2023-01-28', 541666.6666666666, '2023-01-27 16:33:28.496183', '2023-01-27 16:33:28.496183', NULL);


--
-- TOC entry 3525 (class 0 OID 33306)
-- Dependencies: 219
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (1, 'admin', '2023-01-18 09:28:13.776385', '2023-01-18 09:28:13.776385', NULL);
INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (2, 'user', '2023-01-18 09:28:13.776385', '2023-01-18 09:28:13.776385', NULL);
INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (3, 'host', '2023-01-18 09:28:13.776385', '2023-01-18 09:28:13.776385', NULL);
INSERT INTO public.roles (id, name, created_at, updated_at, deleted_at) VALUES (4, 'superadmin', '2023-01-20 15:43:45.617248', '2023-01-20 15:43:45.617248', NULL);


--
-- TOC entry 3545 (class 0 OID 33396)
-- Dependencies: 239
-- Data for Name: source_of_funds; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.source_of_funds (id, name, created_at, updated_at, deleted_at) VALUES (1, 'TOP UP', '2023-01-19 15:14:22.153239', '2023-01-19 15:14:22.153239', NULL);
INSERT INTO public.source_of_funds (id, name, created_at, updated_at, deleted_at) VALUES (2, 'RENT', '2023-01-19 15:14:35.301155', '2023-01-19 15:14:35.301155', NULL);
INSERT INTO public.source_of_funds (id, name, created_at, updated_at, deleted_at) VALUES (3, 'GAMES', '2023-01-19 15:14:52.618777', '2023-01-19 15:14:52.618777', NULL);
INSERT INTO public.source_of_funds (id, name, created_at, updated_at, deleted_at) VALUES (4, 'DEBIT', '2023-01-24 18:16:11.138364', '2023-01-24 18:16:11.138364', NULL);


--
-- TOC entry 3537 (class 0 OID 33362)
-- Dependencies: 231
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (1, 1, 10, 50000, '2023-01-19 16:05:16.596382', '2023-01-19 16:05:16.596382', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (2, 3, 27, 727887, '2023-01-23 23:36:29.752202', '2023-01-23 23:36:29.752202', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (3, 3, 27, 498081, '2023-01-24 02:46:07.674014', '2023-01-24 02:46:07.674014', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (4, 3, 27, 498081, '2023-01-24 02:47:21.620761', '2023-01-24 02:47:21.620761', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (5, 3, 27, 498081, '2023-01-24 09:07:33.135861', '2023-01-24 09:07:33.135861', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (6, 3, 27, 440371, '2023-01-24 09:14:41.756712', '2023-01-24 09:14:41.756712', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (7, 3, 27, 724367, '2023-01-24 09:15:12.901343', '2023-01-24 09:15:12.901343', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (8, 3, 27, 169754, '2023-01-24 09:15:15.356022', '2023-01-24 09:15:15.356022', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (14, 3, 27, 971164, '2023-01-24 09:56:58.04924', '2023-01-24 09:56:58.04924', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (17, 3, 27, 144322, '2023-01-24 10:21:45.188399', '2023-01-24 10:21:45.188399', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (18, 3, 27, 106781, '2023-01-24 10:22:17.360462', '2023-01-24 10:22:17.360462', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (19, 3, 27, 957396, '2023-01-24 10:22:18.474084', '2023-01-24 10:22:18.474084', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (20, 3, 27, 239596, '2023-01-24 10:22:19.554885', '2023-01-24 10:22:19.554885', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (21, 3, 27, 293250, '2023-01-24 10:22:20.523846', '2023-01-24 10:22:20.523846', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (22, 3, 27, 629237, '2023-01-24 10:22:21.434973', '2023-01-24 10:22:21.434973', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (23, 3, 27, 728480, '2023-01-24 10:22:22.675962', '2023-01-24 10:22:22.675962', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (24, 3, 27, 3843, '2023-01-24 10:22:23.637495', '2023-01-24 10:22:23.637495', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (25, 3, 27, 2460, '2023-01-24 10:22:24.582752', '2023-01-24 10:22:24.582752', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (26, 3, 27, 792335, '2023-01-24 10:22:25.452926', '2023-01-24 10:22:25.452926', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (27, 1, 27, 100000, '2023-01-24 11:57:31.162308', '2023-01-24 11:57:31.162308', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (28, 1, 27, 500000, '2023-01-24 11:57:43.880762', '2023-01-24 11:57:43.880762', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (29, 1, 27, 500000, '2023-01-24 12:03:28.26687', '2023-01-24 12:03:28.26687', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (30, 1, 27, 500000, '2023-01-24 12:04:46.812906', '2023-01-24 12:04:46.812906', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (31, 1, 27, 500000, '2023-01-24 12:15:17.405383', '2023-01-24 12:15:17.405383', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (32, 1, 27, 50000, '2023-01-24 12:17:05.878017', '2023-01-24 12:17:05.878017', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (33, 1, 27, 500000, '2023-01-24 12:17:10.62574', '2023-01-24 12:17:10.62574', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (34, 1, 27, 1000000, '2023-01-24 12:17:39.284365', '2023-01-24 12:17:39.284365', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (40, 4, 27, 0, '2023-01-24 19:53:58.751701', '2023-01-24 19:53:58.751701', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (41, 4, 27, 0, '2023-01-24 19:59:25.085151', '2023-01-24 19:59:25.085151', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (42, 4, 27, 1500000, '2023-01-25 08:42:40.396698', '2023-01-25 08:42:40.396698', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (43, 4, 27, 1500000, '2023-01-25 08:44:47.607772', '2023-01-25 08:44:47.607772', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (44, 4, 27, 1500000, '2023-01-25 08:49:22.863213', '2023-01-25 08:49:22.863213', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (45, 4, 27, 1500000, '2023-01-25 08:52:35.354465', '2023-01-25 08:52:35.354465', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (46, 4, 27, 2000000, '2023-01-25 09:14:30.186035', '2023-01-25 09:14:30.186035', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (47, 4, 27, 1000000, '2023-01-25 09:37:52.647963', '2023-01-25 09:37:52.647963', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (48, 4, 27, 500000, '2023-01-25 09:38:19.890917', '2023-01-25 09:38:19.890917', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (49, 4, 27, 1000000, '2023-01-25 09:53:52.608198', '2023-01-25 09:53:52.608198', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (50, 4, 27, 500000, '2023-01-25 09:54:06.210671', '2023-01-25 09:54:06.210671', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (51, 4, 27, 1000000, '2023-01-25 10:10:59.362312', '2023-01-25 10:10:59.362312', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (52, 4, 27, 500000, '2023-01-25 10:11:19.36585', '2023-01-25 10:11:19.36585', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (53, 4, 27, 2000000, '2023-01-25 10:28:58.434639', '2023-01-25 10:28:58.434639', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (54, 4, 27, 500000, '2023-01-25 10:30:54.733811', '2023-01-25 10:30:54.733811', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (57, 4, 27, 1000000, '2023-01-25 12:09:54.218425', '2023-01-25 12:09:54.218425', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (60, 4, 27, 500000, '2023-01-25 12:20:03.425171', '2023-01-25 12:20:03.425171', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (62, 4, 27, 575000, '2023-01-25 13:35:32.951808', '2023-01-25 13:35:32.951808', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (63, 4, 27, 579375, '2023-01-25 13:38:38.344129', '2023-01-25 13:38:38.344129', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (64, 1, 27, 1000000, '2023-01-27 10:44:20.764921', '2023-01-27 10:44:20.764921', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (65, 4, 27, 251666.6666666667, '2023-01-27 15:53:59.362022', '2023-01-27 15:53:59.362022', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (66, 4, 27, 251666.6666666667, '2023-01-27 15:59:04.177011', '2023-01-27 15:59:04.177011', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (67, 4, 27, 251666.6666666667, '2023-01-27 16:01:30.547552', '2023-01-27 16:01:30.547552', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (68, 4, 27, 250000, '2023-01-27 16:08:36.618431', '2023-01-27 16:08:36.618431', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (69, 4, 27, 250000, '2023-01-27 16:10:26.092689', '2023-01-27 16:10:26.092689', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (70, 4, 27, 250000, '2023-01-27 16:13:07.792032', '2023-01-27 16:13:07.792032', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (71, 4, 27, 243750, '2023-01-27 16:16:59.220647', '2023-01-27 16:16:59.220647', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (72, 4, 27, 543750, '2023-01-27 16:27:30.871331', '2023-01-27 16:27:30.871331', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (73, 4, 27, 543333.3333333333, '2023-01-27 16:31:08.532279', '2023-01-27 16:31:08.532279', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (74, 4, 27, 643125, '2023-01-27 16:32:01.140546', '2023-01-27 16:32:01.140546', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (75, 4, 27, 343125, '2023-01-27 16:32:53.71432', '2023-01-27 16:32:53.71432', NULL);
INSERT INTO public.transactions (id, source_of_funds_id, user_id, balance, created_at, updated_at, deleted_at) VALUES (76, 4, 27, 541666.6666666666, '2023-01-27 16:33:28.49646', '2023-01-27 16:33:28.49646', NULL);


--
-- TOC entry 3521 (class 0 OID 33288)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (1, 'acong@gmail.com', '$2a$12$uBb5fuv6joAcewSi9bhqk.q8wafzmMqwUy3FylzBgOudx0edHCoJu', 'acong suherman', 'Jakarta Timur', 1, 1, '2023-01-18 09:27:27.045805', '2023-01-18 09:27:27.045805', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (2, 'arief@gmail.com', '$2a$12$uBb5fuv6joAcewSi9bhqk.q8wafzmMqwUy3FylzBgOudx0edHCoJu', 'arief saferman', 'Jakarta Selatan', 1, 2, '2023-01-18 14:43:13.086084', '2023-01-18 14:43:13.086084', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (4, 'jeki@gmail.com', '$2a$04$X6MDxdKfyUTEiTcc2VimdeZOtQyM/tTKRl4Y4p0gluWFnfP/8cpEu', 'jeki jeki', 'cipali', 1, 2, '2023-01-18 16:58:22.286748', '2023-01-18 16:58:22.286748', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (5, 'ali@gmail.com', '$2a$04$tQqfu.JVxp7Dr7CJjNonH.ogAbcwRF8PHPKxUAdpw7hMFjj7.lr0q', 'ali ali', 'cipali', 1, 2, '2023-01-18 17:08:31.218083', '2023-01-18 17:08:31.218083', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (6, 'peter@gmail.com', '$2a$04$DD9zvUsbkg8NqG5VrlGy.uBzQPKKdsGq0qzhD1EFMaZw6GfkzUl02', 'peter', 'cipali', 1, 2, '2023-01-18 17:49:16.265909', '2023-01-18 17:49:16.265909', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (8, 'kola@gmail.com', '$2a$04$KOIZRU9YORi/wUd088c1DeAdD1WfqbpT82zX50Xn5FgR5AAKnAVVi', 'kola', 'cipali', 1, 2, '2023-01-19 08:59:02.854576', '2023-01-19 08:59:02.854576', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (9, 'arief.saferman@gmail.com', '$2a$04$.UlKkStCg93kwE.k4usRqe5f2132/xMJHXKVUySBMPTEyf26PKOYm', 'ARIEF SAFERMAN', 'cipayung', 1, 3, '2023-01-19 09:26:20.928579', '2023-01-19 10:11:06.300484', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (10, 'melisa@gmail.com', '$2a$04$kZf5ghDEYM9MI8IlG1HNBer0nwNPCKiwCh0J8w1cyXWNAdBuU9/Pe', 'melisa saferman', 'pasar rebo', 1, 2, '2023-01-19 11:55:08.646642', '2023-01-19 13:15:02.87245', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (11, 'saman@gmail.com', '$2a$04$YnlTYvPWtICkiCvpLkLDWu13irHcZzKIWLZaIhLnsih/TfwcNUhi.', 'saman', 'pasar rebo', 1, 2, '2023-01-19 16:18:38.078807', '2023-01-19 16:18:38.078807', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (12, 'fera@gmail.com', '$2a$04$RjopxjhZe5dnmHrjOV.IUefQnaRb1q1peEuM2qPR5GZlH8Sffq8Xm', 'fera', 'pasar rebo', 1, 2, '2023-01-19 17:03:59.881385', '2023-01-19 17:03:59.881385', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (13, 'yudi@gmail.com', '$2a$04$OKJ6QTcdHQKV7BqrzygCveQNdf.v4uubDPwXyaHzdmdPpAS8SMUmK', 'yudi', 'pasar rebo', 1, 2, '2023-01-19 17:04:53.737646', '2023-01-19 17:04:53.737646', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (14, 'admin2@gmail.com', '$2a$04$Vw7bm1jdw1rT0a/EpctvQ..EKmzJqwXg7gOin5nBzHK8s8fpVPH6a', 'admin2', 'pasar rebo', 1, 2, '2023-01-19 17:30:38.084579', '2023-01-19 17:30:38.084579', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (15, 'admin3@gmail.com', '$2a$04$rCwY0RB5f4sGiR/5Nug2Je1hJismCjDk71Agt.tr91O0RKLVA1c1e', 'admin3', 'pasar rebo', 1, 1, '2023-01-19 17:34:32.571495', '2023-01-19 17:34:32.571495', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (16, 'renol@gmail.com', '$2a$04$g.u.FVp8gas1xbb1PRtjSu5BJ4Dyn3WRQZ2loXVUUDHURoFTnTEf6', 'renol', 'pasar rebo', 1, 2, '2023-01-20 09:03:13.235399', '2023-01-20 09:03:13.235399', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (18, 'arul@gmail.com', '$2a$04$9RG3tw5adEjT4xEfiOdlcukcsCqccqaTxjQ4EGBpdcJpd8ZhNsG1W', 'arul', 'pasar rebo', 1, 2, '2023-01-20 09:12:13.291915', '2023-01-20 09:12:13.291915', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (20, 'test4@gmail.com', '$2a$04$hzA3T0eG24m0JsvOoReHc.y9zbev4l6K8hy.uvoX9EE1RLxaWOzGq', 'test4', 'pasar rebo', 1, 2, '2023-01-20 11:32:32.953332', '2023-01-20 11:32:32.953332', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (21, 'test5@gmail.com', '$2a$04$Bvsu.h/krXMlj0Hs09um.O9RZHStqAx3lCf9hu5bQpJk2qAxbODaq', 'test5', 'pasar rebo', 1, 2, '2023-01-20 11:36:22.347828', '2023-01-20 11:36:22.347828', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (22, 'test6@gmail.com', '$2a$04$YPZYWIDE4voLEddAK99Qqub/TPyXNUUpvDDs4enF..ZuGV1s8FqdC', 'test6', 'pasar rebo', 1, 2, '2023-01-20 11:38:10.488917', '2023-01-20 11:38:10.488917', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (19, 'merisa@gmail.com', '$2a$04$4Jz5owD/qPgVYSe9cQcYo.UbGvh2hRE2L5GPruyla7Ja4jP27AJpe', 'merisa', 'pasar rebo', 1, 3, '2023-01-20 09:14:44.546823', '2023-01-20 15:21:51.624141', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (24, 'javin@gmail.com', '$2a$04$ADuShDFQW6HLGyao0ZH9aOhAaQhyHx2NdwR/ilu4Ah54y8gOQIMri', 'javin djapri', 'kalideres', 5, 2, '2023-01-23 15:37:12.268283', '2023-01-23 15:37:12.268283', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (26, 'yan@gmail.com', '$2a$04$A3HKnNSXBPBhOxoIsFYkXOhH8FRF4cd5SAOhGa0y9mr2D1lVVzIMS', 'yan djapri', 'kalideres', 5, 2, '2023-01-23 15:41:27.235064', '2023-01-23 15:41:27.235064', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (28, 'test99@gmail.com', '$2a$04$zuILkjXCPq6f.Q9v1RdTQ.J4LMkN0uQ3uMSr6jYI1JM3/G3eVPbpu', 'angga djapri', 'kalideres', 1, 2, '2023-01-28 15:09:33.96736', '2023-01-28 15:09:33.96736', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (29, 'joko@gmail.com', '$2a$04$gqYw3EAJKjp860SzBwApuuukBGUfCVioqWjQV8AFLCJ3JqkOpoMgK', 'jokok', 'kalisari', 1, 2, '2023-01-28 16:08:32.440615', '2023-01-28 16:08:32.440615', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (30, 'devi@gmail.com', '$2a$04$9LR2VWC9Qyiy5odZ4fElbe4eigWRWDHNqnhNgSi0HFgCp/M.Ky.7q', 'devi monica', 'pekapuran', 3, 2, '2023-01-28 20:03:24.883723', '2023-01-28 20:03:24.883723', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (31, 'fruit@gmail.com', '$2a$04$4gAKciOg.KNnIMZbPnfEW.VGPa/TL6KeAWKf0AkJ8gxfyPPKV9tY6', 'test', 'asdfasdf', 1, 2, '2023-01-28 20:04:27.161416', '2023-01-28 20:04:27.161416', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (34, 'blabla@gmail.com', '$2a$04$qvUao6ZW4.7ZBFvPXs.bwuFskG7u.EdgWSZ2FJ0qZepcf9p2uOlXa', 'blaem', 'asdfasdf', 1, 2, '2023-01-28 20:09:16.019321', '2023-01-28 20:09:16.019321', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (35, 'kito@gmail.com', '$2a$04$3sprHN8p4hTRc.0MnVRrSeBK.tORC.7iwIujeSbfvmFt1ZZmE8OmG', 'kito', 'kalisari', 1, 2, '2023-01-28 20:09:44.184183', '2023-01-28 20:09:44.184183', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (36, 'pop@gmail.com', '$2a$04$aUOLds0xx0Y.noIst8NO3uAF5it7XABVi.63JKsi/brcILNF2v7Um', 'kasjdfksdjf', 'asfdasdf', 5, 2, '2023-01-28 20:10:56.885519', '2023-01-28 20:10:56.885519', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (27, 'angga@gmail.com', '$2a$04$SI68JCgDrETkz9PsYLQrB.MT4wQC6etZNsmHybmDb2z3dQ4xwaPoi', 'angga djapri', 'kalideres', 5, 3, '2023-01-23 15:43:14.177101', '2023-01-31 09:48:15.045066', NULL);
INSERT INTO public.users (id, email, password, full_name, address, city_id, role_id, created_at, updated_at, deleted_at) VALUES (25, 'dapur@gmail.com', '$2a$04$tDZ7Yggg/o7h.UtecHZ/D.hlflEEquwNhYEvjgOXEvcayQ252cAHa', 'dapur djapri', 'kalideres', 5, 3, '2023-01-23 15:38:45.253459', '2023-02-01 14:28:55.161759', NULL);


--
-- TOC entry 3529 (class 0 OID 33328)
-- Dependencies: 223
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700001, 1, 500000000, '2023-01-18 09:28:47.192081', '2023-01-18 09:28:47.192081', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700003, 5, 0, '2023-01-18 17:08:31.226875', '2023-01-18 17:08:31.226875', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700004, 6, 0, '2023-01-18 17:49:16.268565', '2023-01-18 17:49:16.268565', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700005, 8, 0, '2023-01-19 08:59:02.862127', '2023-01-19 08:59:02.862127', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700025, 28, 0, '2023-01-28 15:09:33.966704', '2023-01-28 15:09:33.966704', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700026, 29, 0, '2023-01-28 16:08:32.433247', '2023-01-28 16:08:32.433247', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700006, 9, 0, '2023-01-19 09:26:20.93593', '2023-01-19 09:26:20.93593', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700027, 30, 0, '2023-01-28 20:03:24.882092', '2023-01-28 20:03:24.882092', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700028, 31, 0, '2023-01-28 20:04:27.161253', '2023-01-28 20:04:27.161253', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700029, 34, 0, '2023-01-28 20:09:16.01925', '2023-01-28 20:09:16.01925', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700030, 35, 0, '2023-01-28 20:09:44.184031', '2023-01-28 20:09:44.184031', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700031, 36, 0, '2023-01-28 20:10:56.885194', '2023-01-28 20:10:56.885194', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700007, 10, 200000, '2023-01-19 11:55:08.657282', '2023-01-19 11:55:08.657282', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700009, 11, 0, '2023-01-19 16:18:38.087295', '2023-01-19 16:18:38.087295', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700010, 12, 0, '2023-01-19 17:03:59.889461', '2023-01-19 17:03:59.889461', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700011, 13, 0, '2023-01-19 17:04:53.739557', '2023-01-19 17:04:53.739557', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700012, 14, 0, '2023-01-19 17:30:38.092821', '2023-01-19 17:30:38.092821', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700013, 15, 0, '2023-01-19 17:34:32.579426', '2023-01-19 17:34:32.579426', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700014, 16, 0, '2023-01-20 09:03:13.237054', '2023-01-20 09:03:13.237054', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700015, 18, 0, '2023-01-20 09:12:13.30225', '2023-01-20 09:12:13.30225', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700017, 20, 0, '2023-01-20 11:32:32.962813', '2023-01-20 11:32:32.962813', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700018, 21, 0, '2023-01-20 11:36:22.357908', '2023-01-20 11:36:22.357908', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700019, 22, 0, '2023-01-20 11:38:10.497776', '2023-01-20 11:38:10.497776', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700024, 27, 736735.0000000000, '2023-01-23 15:43:14.176672', '2023-01-23 15:43:14.176672', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700022, 25, 0, '2023-01-23 15:38:45.253093', '2023-01-23 15:38:45.253093', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700023, 26, 0, '2023-01-23 15:41:27.234529', '2023-01-23 15:41:27.234529', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700002, 4, 4000000, '2023-01-18 16:58:22.295518', '2023-01-18 16:58:22.295518', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700021, 24, 1154375, '2023-01-23 15:37:12.268168', '2023-01-23 15:37:12.268168', NULL);
INSERT INTO public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) VALUES (700016, 19, 4363750.0000000000, '2023-01-20 09:14:44.550571', '2023-01-20 09:14:44.550571', NULL);


--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 241
-- Name: admin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.admin_id_seq', 2, true);


--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 228
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, false);


--
-- TOC entry 3570 (class 0 OID 0)
-- Dependencies: 220
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 5, true);


--
-- TOC entry 3571 (class 0 OID 0)
-- Dependencies: 216
-- Name: games_chances_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.games_chances_id_seq', 12, true);


--
-- TOC entry 3572 (class 0 OID 0)
-- Dependencies: 226
-- Name: houses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.houses_id_seq', 19, true);


--
-- TOC entry 3573 (class 0 OID 0)
-- Dependencies: 224
-- Name: houses_photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.houses_photos_id_seq', 22, true);


--
-- TOC entry 3574 (class 0 OID 0)
-- Dependencies: 234
-- Name: pickup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickup_id_seq', 10, true);


--
-- TOC entry 3575 (class 0 OID 0)
-- Dependencies: 236
-- Name: pickup_status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickup_status_id_seq', 5, true);


--
-- TOC entry 3576 (class 0 OID 0)
-- Dependencies: 232
-- Name: reservations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reservations_id_seq', 41, true);


--
-- TOC entry 3577 (class 0 OID 0)
-- Dependencies: 218
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 6, true);


--
-- TOC entry 3578 (class 0 OID 0)
-- Dependencies: 238
-- Name: source_of_funds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.source_of_funds_id_seq', 4, true);


--
-- TOC entry 3579 (class 0 OID 0)
-- Dependencies: 230
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 76, true);


--
-- TOC entry 3580 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 36, true);


--
-- TOC entry 3581 (class 0 OID 0)
-- Dependencies: 222
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_id_seq', 1, false);


--
-- TOC entry 3582 (class 0 OID 0)
-- Dependencies: 240
-- Name: walletsequence; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.walletsequence', 700031, true);


--
-- TOC entry 3358 (class 2606 OID 33561)
-- Name: admin admin_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT admin_email_key UNIQUE (email);


--
-- TOC entry 3360 (class 2606 OID 33559)
-- Name: admin admin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (id);


--
-- TOC entry 3346 (class 2606 OID 33360)
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- TOC entry 3336 (class 2606 OID 33326)
-- Name: cities cities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_name_key UNIQUE (name);


--
-- TOC entry 3338 (class 2606 OID 33324)
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3330 (class 2606 OID 33304)
-- Name: games_chances games_chances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chances
    ADD CONSTRAINT games_chances_pkey PRIMARY KEY (id);


--
-- TOC entry 3342 (class 2606 OID 33342)
-- Name: houses_photos houses_photos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos
    ADD CONSTRAINT houses_photos_pkey PRIMARY KEY (id);


--
-- TOC entry 3344 (class 2606 OID 33351)
-- Name: houses houses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT houses_pkey PRIMARY KEY (id);


--
-- TOC entry 3352 (class 2606 OID 33385)
-- Name: pickups pickup_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT pickup_pkey PRIMARY KEY (id);


--
-- TOC entry 3354 (class 2606 OID 33394)
-- Name: pickups_statuses pickup_status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups_statuses
    ADD CONSTRAINT pickup_status_pkey PRIMARY KEY (id);


--
-- TOC entry 3350 (class 2606 OID 33376)
-- Name: reservations reservations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pkey PRIMARY KEY (id);


--
-- TOC entry 3332 (class 2606 OID 33315)
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- TOC entry 3334 (class 2606 OID 33313)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 3356 (class 2606 OID 33403)
-- Name: source_of_funds source_of_funds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds
    ADD CONSTRAINT source_of_funds_pkey PRIMARY KEY (id);


--
-- TOC entry 3348 (class 2606 OID 33367)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3326 (class 2606 OID 33297)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3328 (class 2606 OID 33295)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3340 (class 2606 OID 33333)
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- TOC entry 3368 (class 2606 OID 33444)
-- Name: bookings bookings_house_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_house_id_fkey FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- TOC entry 3369 (class 2606 OID 33449)
-- Name: bookings bookings_reservation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_reservation_id_fkey FOREIGN KEY (reservation_id) REFERENCES public.reservations(id);


--
-- TOC entry 3370 (class 2606 OID 33439)
-- Name: bookings bookings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3363 (class 2606 OID 33414)
-- Name: games_chances games_chances_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.games_chances
    ADD CONSTRAINT games_chances_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3366 (class 2606 OID 33434)
-- Name: houses houses_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT houses_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- TOC entry 3365 (class 2606 OID 33424)
-- Name: houses_photos houses_photos_house_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses_photos
    ADD CONSTRAINT houses_photos_house_id_fkey FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- TOC entry 3367 (class 2606 OID 33429)
-- Name: houses houses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT houses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3375 (class 2606 OID 33484)
-- Name: pickups pickup_pickup_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT pickup_pickup_status_id_fkey FOREIGN KEY (pickup_status_id) REFERENCES public.pickups_statuses(id);


--
-- TOC entry 3376 (class 2606 OID 33479)
-- Name: pickups pickup_reservation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT pickup_reservation_id_fkey FOREIGN KEY (reservation_id) REFERENCES public.reservations(id);


--
-- TOC entry 3377 (class 2606 OID 33474)
-- Name: pickups pickup_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT pickup_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3373 (class 2606 OID 33464)
-- Name: reservations reservations_house_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_house_id_fkey FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- TOC entry 3374 (class 2606 OID 33469)
-- Name: reservations reservations_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3371 (class 2606 OID 33454)
-- Name: transactions transactions_source_of_funds_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_source_of_funds_id_fkey FOREIGN KEY (source_of_funds_id) REFERENCES public.source_of_funds(id);


--
-- TOC entry 3372 (class 2606 OID 33459)
-- Name: transactions transactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3361 (class 2606 OID 33404)
-- Name: users users_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- TOC entry 3362 (class 2606 OID 33409)
-- Name: users users_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 3364 (class 2606 OID 33419)
-- Name: wallets wallets_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2023-02-03 09:58:41 WIB

--
-- PostgreSQL database dump complete
--


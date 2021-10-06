--
-- PostgreSQL database dump
--

-- Dumped from database version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)

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
-- Name: tbl_product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tbl_product (
    id integer NOT NULL,
    name text NOT NULL,
    price integer NOT NULL,
    qty integer NOT NULL,
    path text
);


ALTER TABLE public.tbl_product OWNER TO postgres;

--
-- Name: tbl_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tbl_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_product_id_seq OWNER TO postgres;

--
-- Name: tbl_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tbl_product_id_seq OWNED BY public.tbl_product.id;


--
-- Name: tbl_transaksi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tbl_transaksi (
    id integer NOT NULL,
    user_id text NOT NULL,
    product_id integer NOT NULL,
    order_id text NOT NULL,
    amount integer NOT NULL,
    status boolean DEFAULT false
);


ALTER TABLE public.tbl_transaksi OWNER TO postgres;

--
-- Name: tbl_transaksi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tbl_transaksi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_transaksi_id_seq OWNER TO postgres;

--
-- Name: tbl_transaksi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tbl_transaksi_id_seq OWNED BY public.tbl_transaksi.id;


--
-- Name: tbl_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tbl_user (
    name text NOT NULL,
    email text NOT NULL,
    password text,
    status text NOT NULL
);


ALTER TABLE public.tbl_user OWNER TO postgres;

--
-- Name: tbl_product id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_product ALTER COLUMN id SET DEFAULT nextval('public.tbl_product_id_seq'::regclass);


--
-- Name: tbl_transaksi id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_transaksi ALTER COLUMN id SET DEFAULT nextval('public.tbl_transaksi_id_seq'::regclass);


--
-- Data for Name: tbl_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tbl_product (id, name, price, qty, path) FROM stdin;
2	contoh2	1000	3	fae4abfb0981f113d7a35d9dd33685a12106c2973246a97f74dd07178dbf2425.png
\.


--
-- Data for Name: tbl_transaksi; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tbl_transaksi (id, user_id, product_id, order_id, amount, status) FROM stdin;
\.


--
-- Data for Name: tbl_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tbl_user (name, email, password, status) FROM stdin;
admin	admin@email.com	5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8	admin
ridwan	ridwan@email.com	5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8	admin
\.


--
-- Name: tbl_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tbl_product_id_seq', 2, true);


--
-- Name: tbl_transaksi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tbl_transaksi_id_seq', 1, false);


--
-- Name: tbl_product tbl_product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_product
    ADD CONSTRAINT tbl_product_pkey PRIMARY KEY (id);


--
-- Name: tbl_transaksi tbl_transaksi_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_transaksi
    ADD CONSTRAINT tbl_transaksi_pk PRIMARY KEY (id);


--
-- Name: tbl_transaksi_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tbl_transaksi_id_uindex ON public.tbl_transaksi USING btree (id);


--
-- Name: tbl_transaksi_order_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tbl_transaksi_order_id_uindex ON public.tbl_transaksi USING btree (order_id);


--
-- Name: tbl_user_email_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tbl_user_email_uindex ON public.tbl_user USING btree (email);


--
-- Name: tbl_transaksi tbl_transaksi_tbl_product_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_transaksi
    ADD CONSTRAINT tbl_transaksi_tbl_product_id_fk FOREIGN KEY (product_id) REFERENCES public.tbl_product(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tbl_transaksi tbl_transaksi_tbl_user_email_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tbl_transaksi
    ADD CONSTRAINT tbl_transaksi_tbl_user_email_fk FOREIGN KEY (user_id) REFERENCES public.tbl_user(email);


--
-- PostgreSQL database dump complete
--


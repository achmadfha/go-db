--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9
-- Dumped by pg_dump version 16.0

-- Started on 2024-01-24 05:57:59

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

ALTER TABLE IF EXISTS ONLY public.transaksi_laundry DROP CONSTRAINT IF EXISTS transaksi_laundry_customer_id_fkey;
ALTER TABLE IF EXISTS ONLY public.transaksi_detail DROP CONSTRAINT IF EXISTS transaksi_detail_transaksi_id_fkey;
ALTER TABLE IF EXISTS ONLY public.transaksi_detail DROP CONSTRAINT IF EXISTS transaksi_detail_pelayanan_id_fkey;
ALTER TABLE IF EXISTS ONLY public.transaksi_detail DROP CONSTRAINT IF EXISTS fk_transaksi;
ALTER TABLE IF EXISTS ONLY public.transaksi_detail DROP CONSTRAINT IF EXISTS fk_pelayanan;
ALTER TABLE IF EXISTS ONLY public.transaksi_laundry DROP CONSTRAINT IF EXISTS transaksi_laundry_pkey;
ALTER TABLE IF EXISTS ONLY public.transaksi_detail DROP CONSTRAINT IF EXISTS transaksi_detail_pkey;
ALTER TABLE IF EXISTS ONLY public.pelayanan DROP CONSTRAINT IF EXISTS pelayanan_pkey;
ALTER TABLE IF EXISTS ONLY public.customer DROP CONSTRAINT IF EXISTS customer_pkey;
ALTER TABLE IF EXISTS public.transaksi_laundry ALTER COLUMN transaksi_id DROP DEFAULT;
ALTER TABLE IF EXISTS public.transaksi_detail ALTER COLUMN transaksi_detail_id DROP DEFAULT;
ALTER TABLE IF EXISTS public.pelayanan ALTER COLUMN pelayanan_id DROP DEFAULT;
ALTER TABLE IF EXISTS public.customer ALTER COLUMN customer_id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.transaksi_laundry_transaksi_id_seq;
DROP TABLE IF EXISTS public.transaksi_laundry;
DROP SEQUENCE IF EXISTS public.transaksi_detail_transaksi_detail_id_seq;
DROP TABLE IF EXISTS public.transaksi_detail;
DROP SEQUENCE IF EXISTS public.pelayanan_pelayanan_id_seq;
DROP TABLE IF EXISTS public.pelayanan;
DROP SEQUENCE IF EXISTS public.customer_customer_id_seq;
DROP TABLE IF EXISTS public.customer;
SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 230573)
-- Name: customer; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.customer (
    customer_id integer NOT NULL,
    nama_customer character varying(255) NOT NULL,
    no_hp character varying(15) NOT NULL
);


--
-- TOC entry 202 (class 1259 OID 230571)
-- Name: customer_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.customer_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2849 (class 0 OID 0)
-- Dependencies: 202
-- Name: customer_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.customer_customer_id_seq OWNED BY public.customer.customer_id;


--
-- TOC entry 205 (class 1259 OID 230581)
-- Name: pelayanan; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.pelayanan (
    pelayanan_id integer NOT NULL,
    jenis_pelayanan character varying(50) NOT NULL,
    satuan character varying(10) NOT NULL,
    harga integer NOT NULL
);


--
-- TOC entry 204 (class 1259 OID 230579)
-- Name: pelayanan_pelayanan_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.pelayanan_pelayanan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2850 (class 0 OID 0)
-- Dependencies: 204
-- Name: pelayanan_pelayanan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.pelayanan_pelayanan_id_seq OWNED BY public.pelayanan.pelayanan_id;


--
-- TOC entry 209 (class 1259 OID 230602)
-- Name: transaksi_detail; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transaksi_detail (
    transaksi_detail_id integer NOT NULL,
    transaksi_id integer,
    pelayanan_id integer,
    jumlah integer NOT NULL
);


--
-- TOC entry 208 (class 1259 OID 230600)
-- Name: transaksi_detail_transaksi_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.transaksi_detail_transaksi_detail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2851 (class 0 OID 0)
-- Dependencies: 208
-- Name: transaksi_detail_transaksi_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.transaksi_detail_transaksi_detail_id_seq OWNED BY public.transaksi_detail.transaksi_detail_id;


--
-- TOC entry 207 (class 1259 OID 230589)
-- Name: transaksi_laundry; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transaksi_laundry (
    transaksi_id integer NOT NULL,
    no_nota character varying(10) NOT NULL,
    tanggal_masuk date NOT NULL,
    tanggal_selesai date NOT NULL,
    diterima_oleh character varying(50) NOT NULL,
    customer_id integer
);


--
-- TOC entry 206 (class 1259 OID 230587)
-- Name: transaksi_laundry_transaksi_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.transaksi_laundry_transaksi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2852 (class 0 OID 0)
-- Dependencies: 206
-- Name: transaksi_laundry_transaksi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.transaksi_laundry_transaksi_id_seq OWNED BY public.transaksi_laundry.transaksi_id;


--
-- TOC entry 2701 (class 2604 OID 230576)
-- Name: customer customer_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_customer_id_seq'::regclass);


--
-- TOC entry 2702 (class 2604 OID 230584)
-- Name: pelayanan pelayanan_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pelayanan ALTER COLUMN pelayanan_id SET DEFAULT nextval('public.pelayanan_pelayanan_id_seq'::regclass);


--
-- TOC entry 2704 (class 2604 OID 230605)
-- Name: transaksi_detail transaksi_detail_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail ALTER COLUMN transaksi_detail_id SET DEFAULT nextval('public.transaksi_detail_transaksi_detail_id_seq'::regclass);


--
-- TOC entry 2703 (class 2604 OID 230592)
-- Name: transaksi_laundry transaksi_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_laundry ALTER COLUMN transaksi_id SET DEFAULT nextval('public.transaksi_laundry_transaksi_id_seq'::regclass);


--
-- TOC entry 2706 (class 2606 OID 230578)
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);


--
-- TOC entry 2708 (class 2606 OID 230586)
-- Name: pelayanan pelayanan_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pelayanan
    ADD CONSTRAINT pelayanan_pkey PRIMARY KEY (pelayanan_id);


--
-- TOC entry 2712 (class 2606 OID 230607)
-- Name: transaksi_detail transaksi_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail
    ADD CONSTRAINT transaksi_detail_pkey PRIMARY KEY (transaksi_detail_id);


--
-- TOC entry 2710 (class 2606 OID 230594)
-- Name: transaksi_laundry transaksi_laundry_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_laundry
    ADD CONSTRAINT transaksi_laundry_pkey PRIMARY KEY (transaksi_id);


--
-- TOC entry 2714 (class 2606 OID 230623)
-- Name: transaksi_detail fk_pelayanan; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail
    ADD CONSTRAINT fk_pelayanan FOREIGN KEY (pelayanan_id) REFERENCES public.pelayanan(pelayanan_id);


--
-- TOC entry 2715 (class 2606 OID 230618)
-- Name: transaksi_detail fk_transaksi; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail
    ADD CONSTRAINT fk_transaksi FOREIGN KEY (transaksi_id) REFERENCES public.transaksi_laundry(transaksi_id);


--
-- TOC entry 2716 (class 2606 OID 230613)
-- Name: transaksi_detail transaksi_detail_pelayanan_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail
    ADD CONSTRAINT transaksi_detail_pelayanan_id_fkey FOREIGN KEY (pelayanan_id) REFERENCES public.pelayanan(pelayanan_id);


--
-- TOC entry 2717 (class 2606 OID 230608)
-- Name: transaksi_detail transaksi_detail_transaksi_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_detail
    ADD CONSTRAINT transaksi_detail_transaksi_id_fkey FOREIGN KEY (transaksi_id) REFERENCES public.transaksi_laundry(transaksi_id);


--
-- TOC entry 2713 (class 2606 OID 230595)
-- Name: transaksi_laundry transaksi_laundry_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksi_laundry
    ADD CONSTRAINT transaksi_laundry_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id);


-- Completed on 2024-01-24 05:57:59

--
-- PostgreSQL database dump complete
--


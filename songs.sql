--
-- PostgreSQL database dump
--

-- Dumped from database version 10.23
-- Dumped by pg_dump version 15.1

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

--
-- Name: songs; Type: TABLE; Schema: public; Owner: sesa
--

CREATE TABLE public.songs (
    song_id integer NOT NULL,
    title text NOT NULL,
    artist text NOT NULL,
    album text NOT NULL,
    year integer NOT NULL,
    genre text,
    duration interval,
    audio_file_path character varying(255) NOT NULL
);


ALTER TABLE public.songs OWNER TO sesa;

--
-- Name: songs_id_seq; Type: SEQUENCE; Schema: public; Owner: sesa
--

CREATE SEQUENCE public.songs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.songs_id_seq OWNER TO sesa;

--
-- Name: songs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sesa
--

ALTER SEQUENCE public.songs_id_seq OWNED BY public.songs.song_id;


--
-- Name: songs song_id; Type: DEFAULT; Schema: public; Owner: sesa
--

ALTER TABLE ONLY public.songs ALTER COLUMN song_id SET DEFAULT nextval('public.songs_id_seq'::regclass);


--
-- Data for Name: songs; Type: TABLE DATA; Schema: public; Owner: sesa
--

COPY public.songs (song_id, title, artist, album, year, genre, duration, audio_file_path) FROM stdin;
1	Title	Artist	Album	2023	Punk	03:00:00	song.wav
\.


--
-- Name: songs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sesa
--

SELECT pg_catalog.setval('public.songs_id_seq', 1, true);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: sesa
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (song_id);


--
-- PostgreSQL database dump complete
--


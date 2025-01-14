PGDMP     ,    -    
            {            progflow    15.3    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    24591    progflow    DATABASE     {   CREATE DATABASE progflow WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_India.1252';
    DROP DATABASE progflow;
                postgres    false            �            1259    41074    coding_activities    TABLE     :  CREATE TABLE public.coding_activities (
    activity_id integer NOT NULL,
    username character varying(30) NOT NULL,
    project_name character varying(30) NOT NULL,
    start_time character varying(10),
    end_time character varying(10),
    duration integer DEFAULT 0,
    created_at character varying(20)
);
 %   DROP TABLE public.coding_activities;
       public         heap    postgres    false            �            1259    41073 !   coding_activities_activity_id_seq    SEQUENCE     �   CREATE SEQUENCE public.coding_activities_activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 8   DROP SEQUENCE public.coding_activities_activity_id_seq;
       public          postgres    false    219                       0    0 !   coding_activities_activity_id_seq    SEQUENCE OWNED BY     g   ALTER SEQUENCE public.coding_activities_activity_id_seq OWNED BY public.coding_activities.activity_id;
          public          postgres    false    218            �            1259    65578    language_activities    TABLE     v  CREATE TABLE public.language_activities (
    activity_id integer NOT NULL,
    username character varying(20) NOT NULL,
    project_name character varying(30) NOT NULL,
    language_name character varying(20) NOT NULL,
    start_time character varying(10) NOT NULL,
    end_time character varying(10),
    duration integer,
    created_at character varying(20) NOT NULL
);
 '   DROP TABLE public.language_activities;
       public         heap    postgres    false            �            1259    65577 #   language_activities_activity_id_seq    SEQUENCE     �   CREATE SEQUENCE public.language_activities_activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 :   DROP SEQUENCE public.language_activities_activity_id_seq;
       public          postgres    false    221                       0    0 #   language_activities_activity_id_seq    SEQUENCE OWNED BY     k   ALTER SEQUENCE public.language_activities_activity_id_seq OWNED BY public.language_activities.activity_id;
          public          postgres    false    220            �            1259    32805    projects    TABLE     �   CREATE TABLE public.projects (
    project_id integer NOT NULL,
    project_name character varying(100) NOT NULL,
    username character varying(30) NOT NULL,
    created_at character varying(20) DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.projects;
       public         heap    postgres    false            �            1259    32804    projects_project_id_seq    SEQUENCE     �   CREATE SEQUENCE public.projects_project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.projects_project_id_seq;
       public          postgres    false    217                       0    0    projects_project_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.projects_project_id_seq OWNED BY public.projects.project_id;
          public          postgres    false    216            �            1259    32793    users    TABLE       CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    username text NOT NULL,
    email text NOT NULL,
    join_date date DEFAULT CURRENT_DATE NOT NULL,
    is_premium boolean NOT NULL,
    api_key text NOT NULL,
    photo character varying NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    32792    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    215                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    214            x           2604    41077    coding_activities activity_id    DEFAULT     �   ALTER TABLE ONLY public.coding_activities ALTER COLUMN activity_id SET DEFAULT nextval('public.coding_activities_activity_id_seq'::regclass);
 L   ALTER TABLE public.coding_activities ALTER COLUMN activity_id DROP DEFAULT;
       public          postgres    false    219    218    219            z           2604    65581    language_activities activity_id    DEFAULT     �   ALTER TABLE ONLY public.language_activities ALTER COLUMN activity_id SET DEFAULT nextval('public.language_activities_activity_id_seq'::regclass);
 N   ALTER TABLE public.language_activities ALTER COLUMN activity_id DROP DEFAULT;
       public          postgres    false    221    220    221            v           2604    32808    projects project_id    DEFAULT     z   ALTER TABLE ONLY public.projects ALTER COLUMN project_id SET DEFAULT nextval('public.projects_project_id_seq'::regclass);
 B   ALTER TABLE public.projects ALTER COLUMN project_id DROP DEFAULT;
       public          postgres    false    217    216    217            t           2604    32796    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            �           2606    41079 (   coding_activities coding_activities_pkey 
   CONSTRAINT     o   ALTER TABLE ONLY public.coding_activities
    ADD CONSTRAINT coding_activities_pkey PRIMARY KEY (activity_id);
 R   ALTER TABLE ONLY public.coding_activities DROP CONSTRAINT coding_activities_pkey;
       public            postgres    false    219            �           2606    65583 ,   language_activities language_activities_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY public.language_activities
    ADD CONSTRAINT language_activities_pkey PRIMARY KEY (activity_id);
 V   ALTER TABLE ONLY public.language_activities DROP CONSTRAINT language_activities_pkey;
       public            postgres    false    221            ~           2606    32811    projects projects_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (project_id);
 @   ALTER TABLE ONLY public.projects DROP CONSTRAINT projects_pkey;
       public            postgres    false    217            |           2606    32803    users users_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (username);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215            �           2606    41080 1   coding_activities coding_activities_username_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.coding_activities
    ADD CONSTRAINT coding_activities_username_fkey FOREIGN KEY (username) REFERENCES public.users(username);
 [   ALTER TABLE ONLY public.coding_activities DROP CONSTRAINT coding_activities_username_fkey;
       public          postgres    false    219    3196    215            �           2606    32812    projects projects_username_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_username_fkey FOREIGN KEY (username) REFERENCES public.users(username);
 I   ALTER TABLE ONLY public.projects DROP CONSTRAINT projects_username_fkey;
       public          postgres    false    215    217    3196           
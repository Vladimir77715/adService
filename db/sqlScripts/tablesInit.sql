CREATE TABLE public.ad (
                          id int4 PRIMARY KEY,
                          name text NOT NULL,
                          description text,
                          price int4,
                          create_ts timestamp
);


CREATE TABLE public.ad_images_links (
                           id int4,
                           link text
);

CREATE INDEX ON public.ad_images_links (id);
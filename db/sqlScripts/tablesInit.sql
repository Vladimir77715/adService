CREATE TABLE public.ad (
                          id SERIAL PRIMARY KEY,
                          name text NOT NULL,
                          description text,
                          price int4,
                          create_ts timestamp DEFAULT now()
);


CREATE TABLE public.ad_images_links (
                           id SERIAL  PRIMARY KEY,
                           link text,
                           ad_id int4,
                           create_ts timestamp DEFAULT now(),
                           CONSTRAINT fk_ad_images_links  FOREIGN KEY(ad_id) REFERENCES public.ad(id)
);

CREATE INDEX ON public.ad_images_links (id,ad_id);
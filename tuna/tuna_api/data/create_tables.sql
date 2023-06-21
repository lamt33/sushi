CREATE TABLE public.projects (
	project_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name varchar NOT null DEFAULT '',
	project_type varchar  NOT null DEFAULT '',
	description text  NOT null DEFAULT '',
	image_url varchar  NOT null DEFAULT '',
	github_urls varchar  NOT null DEFAULT '',
	external_url varchar  NOT null DEFAULT '',
	discord_url varchar NOT null DEFAULT '',
	twitter_username varchar NOT null DEFAULT '',
	telegram_url varchar  NOT null DEFAULT '',
	subreddit_url varchar  NOT null DEFAULT '',
	instagram_url varchar  NOT null DEFAULT '',
	source varchar  NOT null DEFAULT '',
	source_url varchar  NOT null DEFAULT '',
	created_time timestamp(0) NULL,
	source_id varchar  NOT null DEFAULT '',
	is_active boolean NOT null DEFAULT false,
	is_new boolean NOT null DEFAULT false
);

ALTER TABLE projects
    ADD CONSTRAINT  projects_constraints_name
    UNIQUE (name);

CREATE TABLE public.nft_names (
	name_id UUID NOT NULL DEFAULT gen_random_uuid(),
	name varchar NOT null DEFAULT '',
	source varchar NOT null DEFAULT '',
	is_new boolean NOT null DEFAULT true
);

ALTER TABLE nft_names
    ADD CONSTRAINT nft_name_constraint
    UNIQUE (name, source);

CREATE TABLE public.tweets (
	tweet_id varchar NOT NULL,
	text varchar not null default '',
	likes int NOT NULL,
	quotes int NOT NULL,
	retweets int NOT NULL,
	replies int NOT NULL default 0,
	author_id varchar NOT NULL,
	twitter_username varchar NOT NULL,
	project_id UUID,
	created_time timestamp(0) NULL,
	sink_time timestamp(0) NULL
);


CREATE TABLE public.agg_tweets (
	project_id UUID NOT NULL,
	author_id varchar NOT NULL,
	twitter_username varchar NOT NULL,
	posts int NOT NULL default 0,
	likes int NOT NULL default 0,
	quotes int NOT NULL default 0,
	retweets int NOT NULL default 0,
	replies int NOT NULL default 0,
	created_time timestamp(0) NULL
);

ALTER TABLE tweets
    ADD CONSTRAINT tweet_id_constraint
        UNIQUE (tweet_id);

CREATE TABLE public.nft_stats (
	nft_stats_id UUID NOT NULL DEFAULT gen_random_uuid(),
	project_id UUID NULL,
	one_day_volume float8 NULL,
	one_day_sales varchar NULL,
	created_time timestamp(0) NULL,
	seven_day_volume float8 NULL,
	seven_day_sales float8 NULL,
	one_day_average_price float8 NULL,
	seven_day_average_price float8 NULL,
	thirty_day_volume float8 NULL,
	thirty_day_sales float8 NULL,
	thirty_day_average_price float8 NULL,
	total_volume float8 NULL,
	total_sales float8 NULL,
	total_supply float8 NULL,
	count float8 NULL,
	num_owners int NULL,
	avg_price float8 NULL,
	market_cap float8 NULL,
	floor_price float8 NULL,
	twitter_followers int NULL,
	tweet_posts_count int NOT NULL DEFAULT 0,
	tweet_quotes_count int NOT NULL DEFAULT 0,
	tweet_likes_count int NOT NULL DEFAULT 0,
	tweet_replies_count int NOT NULL DEFAULT 0,
	tweet_retweets_count int NOT NULL DEFAULT 0,
	denomination varchar NOT NULL DEFAULT 'eth'
);


CREATE TABLE public.coin_stats (
    coin_stats_id UUID NOT NULL DEFAULT gen_random_uuid(),
    project_id UUID NULL,
    current_price float8 NULL,
    market_cap float8 NULL,
    volume float8 NULL,
    high_24h float8 NULL,
    low_24h float8 NULL,
    circulating_supply float8 NULL,
    twitter_followers int NULL,
	tweet_posts_count int NOT NULL DEFAULT 0,
	tweet_quotes_count int NOT NULL DEFAULT 0,
	tweet_likes_count int NOT NULL DEFAULT 0,
	tweet_replies_count int NOT NULL DEFAULT 0,
	tweet_retweets_count int NOT NULL DEFAULT 0,
    reddit_subscribers int NULL,
    reddit_average_posts_48h float8 NULL,
    reddit_average_comments_48h float8 NULL,
    git_forks int NULL,
    git_stars int NULL,
    git_subscribers int NULL,
    git_pull_requests_merged int NULL,
    git_contributors int NULL,
    git_commit_count_4w int NULL,
    created_time timestamp(0) NULL
);


CREATE TABLE public.chains (
	chain_id UUID NOT NULL DEFAULT gen_random_uuid(),
	project_id UUID NULL,
	primary_address varchar NULL,
	chain_name varchar null,
	token_id varchar null
);

ALTER TABLE chains
    ADD CONSTRAINT chains_project_id_constraint
    UNIQUE (project_id, chain_name);

CREATE TABLE public.feed (
	feed_event_id UUID NOT NULL DEFAULT gen_random_uuid(),
	project_id UUID NOT NULL DEFAULT gen_random_uuid(),
	project_name varchar NULL,
	project_type  varchar NOT NULL DEFAULT '',
	image_url varchar NULL,
	event_type varchar NULL,
	created_time timestamp(0) NULL,
	"source" varchar NULL,
	source_url varchar NULL,
	event_current_stat float8 NULL,
	event_prev_stat float8 NULL,
	event_delta float8 NULL,
	event_magnitude float8 NULL,
	constraint_value varchar NOT NULL DEFAULT ''
);

ALTER TABLE feed
    ADD CONSTRAINT feed_constraint
    UNIQUE (project_id, constraint_value);

CREATE TABLE public.pipeline_metadata (
	sink_id UUID NOT NULL DEFAULT gen_random_uuid(),
	pipeline varchar NULL,
	sink_time timestamp(0) NULL
);
--seed initial sink time with previous date 
insert into pipeline_metadata (pipeline, sink_time) values ('nft_stats', NOW()-interval '1 day');
insert into pipeline_metadata (pipeline, sink_time) values ('coin_stats', NOW()-interval '1 day');
insert into pipeline_metadata (pipeline, sink_time) values ('twitter_project', NOW()-interval '1 day');
insert into pipeline_metadata (pipeline, sink_time) values ('twitter_list', NOW()-interval '1 day');


CREATE TABLE public.user_auth (
	public_key varchar,
	nonce int, 
	created_time timestamp(0) NULL
);

ALTER TABLE user_auth
    ADD CONSTRAINT user_auth_public_key_constraint
    UNIQUE (public_key);

CREATE TABLE public.users (
	user_id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_name  varchar NOT NULL DEFAULT '',
	display_name varchar NOT NULL DEFAULT '', 
	image_url  varchar NOT NULL DEFAULT '', 
	bio varchar NOT NULL DEFAULT '',
	is_public boolean NOT NULL DEFAULT true,
	karma_points int NOT NULL DEFAULT 0,
	show_karma boolean NOT NULL DEFAULT true,
	location varchar NOT NULL DEFAULT '',
	birthday timestamp without time zone default (now() at time zone 'utc'),
	created_time timestamp without time zone default (now() at time zone 'utc')
);

ALTER TABLE users
    ADD CONSTRAINT user_name_constraint
    UNIQUE (user_name);

CREATE TABLE public.user_projects (
    user_project_id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    project_id UUID NOT NULL,
    created_time  timestamp(0) NULL
);

ALTER TABLE user_projects
    ADD CONSTRAINT user_projects_constraint
    UNIQUE (user_id, project_id);


CREATE TABLE public.user_nfts (
  user_nft_id UUID NOT NULL DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  nft_address varchar NOT NULL,
  collection_name varchar NOT NULL,
  wallet_address varchar NOT NULL,
  wallet_type varchar NOT NULL,
  metadata_uri varchar NOT NULL,
  created_time timestamp without time zone default (now() at time zone 'utc')
);

ALTER TABLE user_nfts
    ADD CONSTRAINT user_nfts_constraint
        UNIQUE (wallet_address, nft_address);


CREATE TABLE public.wallets (
	wallet_id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_id  UUID NOT NULL DEFAULT gen_random_uuid(),
	public_key  varchar NOT NULL DEFAULT '',
	wallet_type  varchar NOT NULL DEFAULT '',
	wallet_name  varchar NOT NULL DEFAULT ''
);


CREATE TABLE public.guides (
    guide_id UUID NOT NULL DEFAULT gen_random_uuid(),
    slug varchar NOT NULL,
    title varchar NOT NULL,
    thumbnail_url varchar NOT NULL,
    video_id varchar NOT NULL,
    content  text NOT NULL DEFAULT '',
    content_type  varchar NOT NULL DEFAULT '',
    meta  varchar NOT NULL DEFAULT '',
    course varchar,
    course_description varchar,
    course_image varchar,
    subject varchar,
    subject_description varchar,
    subject_image varchar,
    course_order int,
    subject_order int,
    topic_order int,
    created_time  timestamp(0) NULL
);


ALTER TABLE public.guides
    ADD CONSTRAINT guides_slug_constraint
        UNIQUE (slug);


CREATE TABLE public.guide_comments (
	comment_id UUID NOT NULL DEFAULT gen_random_uuid(),
	parent_comment_id  UUID NULL,
	user_id UUID NOT NULL,
	comment  varchar NOT NULL DEFAULT '',
	created_time  timestamp(0) NULL
);

ALTER TABLE public.guide_comments add slug varchar NOT NULL;

CREATE TABLE public.guide_interactions (
	interaction_id UUID NOT NULL DEFAULT gen_random_uuid(),
	slug  varchar NOT NULL  DEFAULT '',
	user_id UUID NOT NULL,
	is_like boolean NOT NULL DEFAULT false
);

ALTER TABLE guide_interactions
    ADD CONSTRAINT guide_interactions_constraint
    UNIQUE (slug, user_id);


CREATE TABLE public.user_completed_guides (
	completed_guide_id UUID NOT NULL DEFAULT gen_random_uuid(),
	slug  varchar NOT NULL  DEFAULT '',
	user_id UUID NOT NULL,
	guide_state varchar NOT NULL  DEFAULT '',
    created_time timestamp without time zone default (now() at time zone 'utc')
);

ALTER TABLE user_completed_guides
    ADD CONSTRAINT user_completed_guides_id_slug
        UNIQUE (user_id, slug);

CREATE TABLE public.user_social_accounts (
	social_account_id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL,
	social_type  varchar NOT NULL  DEFAULT '',
	social_id varchar NOT NULL  DEFAULT '',
	social_url  varchar NOT NULL  DEFAULT '',
	access_token  varchar NOT NULL  DEFAULT '', 
	refresh_token  varchar NOT NULL  DEFAULT ''
);

ALTER TABLE user_social_accounts
    ADD CONSTRAINT user_social_accounts_constraint
        UNIQUE (user_id, social_id);

CREATE TABLE public.post_groups (
      post_group_id UUID NOT NULL DEFAULT gen_random_uuid(),
      name  varchar NOT NULL DEFAULT '',
      image_url varchar NOT NULL,
      post_ids UUID[],
      created_time timestamp without time zone default (now() at time zone 'utc')
);

CREATE TABLE public.posts (
    post_id UUID NOT NULL DEFAULT gen_random_uuid(),
    post_group_id  varchar NOT NULL DEFAULT '',
    sub_group varchar NOT NULL DEFAULT '',
    title varchar NOT NULL,
    content varchar NOT NULL,
	raw_content varchar not null default '',
    user_id UUID NOT NULL,
    created_time timestamp without time zone default (now() at time zone 'utc'),
    updated_time timestamp without time zone default (now() at time zone 'utc'),
    PRIMARY KEY(post_id)
);

ALTER TABLE posts ADD COLUMN search_tsv tsvector 
    GENERATED ALWAYS AS (setweight(to_tsvector('english', coalesce(title, '')), 'A') || 
         setweight(to_tsvector('english', coalesce(raw_content, '')), 'B')) STORED;

CREATE EXTENSION pg_trgm;
--create indexes on search_tsv and raw_body col for faster search
CREATE INDEX ts_idx ON posts USING GIN (search_tsv); 
CREATE INDEX CONCURRENTLY index_raw_body_trigram ON posts USING gin (raw_content gin_trgm_ops);       


CREATE TABLE public.post_comments (
      comment_id UUID NOT NULL DEFAULT gen_random_uuid(),
      parent_comment_id UUID,
      post_id UUID NOT NULL,
      user_id UUID NOT NULL,
      content varchar NOT NULL,
      created_time timestamp without time zone default (now() at time zone 'utc'),
      updated_time timestamp without time zone default (now() at time zone 'utc'),
      PRIMARY KEY(comment_id),
      CONSTRAINT post_parent_id
          FOREIGN KEY(post_id)
              REFERENCES posts(post_id)
			   on delete cascade
);

CREATE TABLE public.interactions (
      interaction_id UUID NOT NULL DEFAULT gen_random_uuid(),
      post_id UUID,
      comment_id UUID,
      user_id UUID NOT NULL,
      vote int NOT NULL default 0,
      shares int,
      created_time timestamp without time zone default (now() at time zone 'utc'),
      updated_time timestamp without time zone default (now() at time zone 'utc'),
      PRIMARY KEY(interaction_id),
      CONSTRAINT i_post_parent_id
          FOREIGN KEY(post_id)
              REFERENCES posts(post_id)
				on delete cascade,
      CONSTRAINT i_comment_parent_id
        FOREIGN KEY(comment_id)
        REFERENCES post_comments(comment_id)
		 on delete cascade
);

ALTER TABLE interactions
    ADD CONSTRAINT user_interactions
        UNIQUE (user_id, post_id, comment_id);

CREATE UNIQUE INDEX upidx
    ON interactions (user_id, COALESCE(comment_id, '00000000-0000-0000-0000-000000000000'), post_id);

CREATE UNIQUE INDEX upidx2
    ON interactions (user_id, COALESCE(post_id, '00000000-0000-0000-0000-000000000000'), comment_id);


CREATE TABLE public.articles (
                                 slug varchar PRIMARY KEY,
                                 topic_type varchar default 'article',
                                 title varchar NOT NULL CHECK (coalesce(TRIM(title), '') <> ''),
                                 article_ref varchar NOT NULL CHECK (coalesce(TRIM(article_ref), '') <> ''),
                                 content varchar,
                                 author varchar,
                                 author_image varchar,
                                 sources varchar,
                                 created_time timestamp without time zone default (now() at time zone 'utc')
);



CREATE TABLE public.videos (
                               slug varchar PRIMARY KEY,
                               topic_type varchar default 'video',
                               title varchar NOT NULL CHECK (coalesce(TRIM(title), '') <> ''),
                               author varchar,
                               author_image varchar,
                               video_id varchar NOT NULL CHECK (coalesce(TRIM(video_id), '') <> ''),
                               content varchar,
                               sources varchar,
                               created_time timestamp without time zone default (now() at time zone 'utc')
);


alter table user_projects add is_trust bool;
alter table user_projects add quantity float8 NOT NULL default 0;

create table  persons (
     person_id UUID primary KEY  NOT NULL DEFAULT gen_random_uuid(),
     name varchar not null,
     twitter_username varchar
);

create table  project_wiki (
      project_wiki_id uuid primary KEY NOT NULL DEFAULT gen_random_uuid(),
      project_id uuid not null,
      main_content varchar,
      person_ids uuid[],
      token_allocations json[],
      is_active bool,
      created_time timestamp without time zone default (now() at time zone 'utc')
);

-- user project inputs

ALTER TABLE users add  is_editor bool;
ALTER TABLE users ADD  PRIMARY KEY (user_id);

ALTER TABLE projects ADD PRIMARY KEY (project_id);


create table  user_project_trusts (
                                     user_id uuid REFERENCES users (user_id) not null,
                                     project_id uuid REFERENCES projects (project_id) not null,
                                     trust bool not null,
                                     PRIMARY KEY (user_id, project_id)
);


create table  person_refs (
    person_ref_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
    person_id UUID REFERENCES persons (person_id) not null,
    external_id_type VARCHAR not null,
    external_id VARCHAR not null,
    unique (person_id, external_id_type, external_id)
    );


create table  project_wikis (
                               project_id uuid primary KEY NOT NULL,
                               main_content varchar,
                               token_allocations json,
                               created_time timestamp without time zone default (now() at time zone 'utc')
);


create table  project_persons (
                                 project_id  UUID REFERENCES projects (project_id) not null,
                                 person_id UUID REFERENCES persons (person_id) not null,
                                 PRIMARY KEY (project_id, person_id)
);


-- playlists

create table  wikis (
                       wiki_id UUID primary key not null default gen_random_uuid(),
                       slug varchar not null unique,
                       name varchar not null,
                       content varchar not null
);


create table  topics (
                        topic_id UUID primary key not null default gen_random_uuid(),
                        title varchar not null,
                        description  varchar,
                        image_url varchar,
                        topic_type  varchar not null,
                        topic_ref_url  varchar not null,
                        slug varchar,
                        wikis varchar,
                        created_time timestamp without time zone default (now() at time zone 'utc'),
                        updated_time timestamp without time zone default (now() at time zone 'utc')
);


create table  badges (
                        badge_id UUID primary key not null default gen_random_uuid(),
                        badge_name varchar not null,
                        image_url varchar not null,
                        contract_address varchar
);
ALTER table badges add  price float8 not null;
ALTER table badges add  description varchar not null;


create table  playlists (
                           playlist_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                           title varchar,
                           slug varchar not null unique,
                           description varchar,
                           image_url varchar,
                           creator_user_id UUID REFERENCES users(user_id),
                           creator_wallet_address varchar,
                           badge_id uuid  REFERENCES badges (badge_id),
                           created_time timestamp without time zone default (now() at time zone 'utc'),
                           updated_time timestamp without time zone default (now() at time zone 'utc')
);

create table  playlist_views (
                                playlist_view_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                                session_id varchar,
                                playlist_id uuid references playlists (playlist_id) not null,
                                user_id uuid references users (user_id),
                                created_time timestamp without time zone default (now() at time zone 'utc'),
                                unique (session_id, playlist_id)
);


create table  user_playlist_likes (
                                     user_id uuid references users (user_id),
                                     playlist_id uuid references playlists (playlist_id),
                                     created_time timestamp without time zone default (now() at time zone 'utc'),
                                     PRIMARY KEY (user_id, playlist_id)
);


create table  user_playlist_shares (
                                      user_playlist_share_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                                      user_id uuid references users (user_id),
                                      playlist_id uuid references playlists (playlist_id),
                                      shared_to varchar not null,
                                      created_time timestamp without time zone default (now() at time zone 'utc')
);


create table  playlist_topics (
                                 playlist_topic_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                                 playlist_id uuid not null references playlists (playlist_id),
                                 topic_id uuid not null references topics (topic_id) on delete cascade,
                                 topic_order int not null,
                                 unique (playlist_id, topic_id, topic_order)
);

create table  user_badges (
                         user_id uuid references users (user_id),
                         badge_id uuid references badges (badge_id),
                         created_time timestamp without time zone default (now() at time zone 'utc'),
                         PRIMARY KEY (user_id, badge_id)
);

create table  user_followers (
                         user_id uuid references users (user_id),
                         follower_user_id uuid references users (user_id),
                         PRIMARY KEY (user_id, follower_user_id)
);

create table  project_tweet_mentions (
                         project_id uuid references projects (project_id),
                         tweet_id VARCHAR,
                         created_time timestamp not null,
                         unique (project_id, tweet_id)
);

create EXTENSION btree_gist;
create table  project_sentiments (
                        project_id uuid references projects (project_id),
                        user_id uuid references users (user_id),
                        sentiment varchar(50),
                        created_time timestamp not null,
                        end_time timestamp not null,
                        exclude using gist (project_id WITH =, user_id WITH =, tsrange(created_time, end_time) with && )
);

create table reddit_posts (
                        subreddit_post_id varchar NOT NULL,
                        subreddit_id varchar NOT NULL,
                        score int,
                        upvote_ratio float8,
                        num_comments int,
                        created_time timestamp not null,
                        unique (subreddit_post_id)
);
create table  project_reddit_mentions (
                         project_id uuid references projects (project_id),
                         subreddit_post_id VARCHAR,
                         created_time timestamp not null,
                         unique (project_id, subreddit_post_id)
);

create table project_analysis (
                         project_id uuid references projects (project_id),
                         keywords varchar
);

ALTER TABLE playlists
    ADD  categories varchar;
   

ALTER table reddit_posts 
            add subreddit_name varchar,
            add title varchar,
            add body varchar,
            add author varchar,
            add project_id uuid;

-- ===================================================================================
-- ==================== CRM TABLES ===============================================================
-- ===================================================================================

create table if not exists crm_users (
                                         crm_user_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                                         user_name  varchar NOT NULL,
                                         picture  varchar NOT NULL DEFAULT '',
                                         email  varchar NOT NULL DEFAULT '',
                                         oauth_provider  varchar NOT NULL,
                                         created_time timestamp without time zone default (now() at time zone 'utc'),
                                         unique(oauth_provider, email)
);


create table if not exists crm_projects (
                                            crm_project_id UUID primary KEY NOT NULL DEFAULT gen_random_uuid(),
                                            long_name  varchar NOT NULL,
                                            symbol  varchar NOT NULL,
                                            picture  varchar NOT NULL,
                                            data_source varchar not NULL,
                                            created_time timestamp without time zone default (now() at time zone 'utc'),
                                            unique(long_name, symbol)
);
alter table crm_projects add column updated_time timestamp without time zone default (now() at time zone 'utc');
alter table crm_projects add column labels varchar;


create table if not exists crm_project_socials (
                                                   crm_project_id uuid references crm_projects(crm_project_id) ON DELETE CASCADE not NULL,
                                                   social_type varchar NOT NULL,
                                                   social_id varchar NOT NULL,
                                                   backfill_status varchar NOT NULL DEFAULT 'active',
                                                   created_time timestamp without time zone default (now() at time zone 'utc'),
                                                   unique(crm_project_id, social_type)
);


create table if not exists crm_project_users (
                                                 crm_user_id uuid references crm_users(crm_user_id) not NULL,
                                                 crm_project_id uuid references crm_projects(crm_project_id) ON DELETE CASCADE not NULL,
                                                 created_time timestamp without time zone default (now() at time zone 'utc'),
                                                 unique(crm_user_id, crm_project_id)
);

create table if not exists crm_tweets (
    tweet_id varchar not null PRIMARY KEY,
    crm_project_id UUID not null, 
    twitter_member_id varchar not null,
    conversation_id varchar not null,
    referenced_tweet_id varchar not null default '',
    retweet_count int not null default 0,
    reply_count int not null default 0,
    like_count int not null default 0,
    quote_count int not null default 0,
    text varchar not null default '', 
    created_time timestamp without time zone default (now() at time zone 'utc'),
    unique(tweet_id)
);

create table if not exists crm_mentions_tweets (
    tweet_id varchar not null PRIMARY KEY,
    crm_project_id UUID not null,
    twitter_member_id varchar not null, 
    conversation_id varchar not null,
    referenced_tweet_id varchar not null default '',
    retweet_count int not null default 0,
    reply_count int not null default 0,
    like_count int not null default 0,
    quote_count int not null default 0,
    text varchar not null default '', -- Random Parent Post , --@Tokoin
    created_time timestamp without time zone default (now() at time zone 'utc'),
    unique(tweet_id)
);

create table if not exists crm_twitter_members (
    twitter_member_id varchar not null PRIMARY KEY,
    username varchar not null,
    name varchar not null,
    location varchar not null default '',
    picture varchar not null default '',
    followers_count int not null default 0,
    following_count int not null default 0,
    tweet_count int not null default 0,
    listed_count int not null default 0,
    is_entity bool not null default false,
    is_verified bool not null default false,
    created_time timestamp without time zone default (now() at time zone 'utc'),
    unique(twitter_member_id)
);

create table if not exists crm_tweet_interactions (
    crm_project_id uuid not null,
    twitter_member_id varchar not null,
    username varchar not null,
    tweet_id varchar not null,
    interaction_type varchar not null,
    created_time timestamp without time zone default (now() at time zone 'utc')
);



create table if not exists crm_tweet_contexts (
    tweet_id varchar not null PRIMARY KEY,
    crm_project_id uuid not null,
    domain_id varchar not null,
    domain_name varchar not null,
    domain_description varchar not null, 
    entity_id varchar not null,
    entity_name varchar not null
);

create table if not exists crm_tweet_tags (
    tweet_id varchar not null PRIMARY KEY,
    crm_project_id uuid not null,
    tag varchar not null,
    type varchar not null
);





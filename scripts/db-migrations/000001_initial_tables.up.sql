CREATE TABLE tags (
    tag_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX tags_name_idx ON tags(name);

CREATE TABLE users (
    user_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(200) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX users_name_idx ON users(name);
CREATE INDEX users_email_idx ON users(email);

CREATE TABLE posts (
    post_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    text TEXT NOT NULL,
    user_id UUID NOT NULL CONSTRAINT posts_users_fk REFERENCES users(user_id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX posts_user_id_idx ON posts(user_id);

CREATE TABLE posts_tags (
    post_id UUID NOT NULL CONSTRAINT posts_tags_post_fk REFERENCES posts(post_id),
    tag_id UUID NOT NULL CONSTRAINT posts_tags_tag_fk REFERENCES tags(tag_id),
    PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE comments (
    comment_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    post_id UUID NOT NULL CONSTRAINT comments_posts_fk REFERENCES posts(post_id),
    user_id UUID NOT NULL CONSTRAINT comments_users_fk REFERENCES users(user_id),
    text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX comments_post_id_idx ON comments(post_id);
CREATE INDEX comments_user_id_idx ON comments(user_id);

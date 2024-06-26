namespace go base

struct Status {
    1: required i64 code;
    2: required string msg;
}

struct User {
    1: string uid;
    2: string username;
    3: string avatar_url;
    4: string created_at;
    5: string updated_at;
    6: string deleted_at;
}

struct UserLite {
    1: string uid;
    2: string username;
    3: string avatar_url;
}

struct Video {
    1:  string id;
    2:  string user_id;
    3:  string video_url;
    4:  string cover_url;
    5:  string title;
    6:  string description;
    7:  i64 visit_count;
    8:  i64 like_count;
    9:  i64 comment_count;
    10: string created_at;
    11: string updated_at;
    12: string deleted_at;
}

struct Comment {
    1:  string id;
    2:  string user_id;
    3:  string video_id;
    4:  string parent_id;
    5:  i64 like_count;
    6:  i64 child_count;
    7:  string content;
    8:  string created_at;
    9:  string updated_at;
    10: string deleted_at;
}
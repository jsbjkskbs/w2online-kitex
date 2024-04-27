namespace go interact

include "base.thrift"

struct LikeActionRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string video_id;
    4: string comment_id;
    5: string action_type;
}

struct LikeActionResponse {
    1: base.Status base;
}

struct LikeListRequest {
    1: string user_id (api.query="user_id");
    2: i64 page_size (api.query="page_size");
    3: i64 page_num (api.query="page_num");
    4: string access_token (api.header="Access-Token");
    5: string refresh_token (api.header="Refresh-Token");
}


struct LikeListResponseData {
    required list<base.Video> items;
}
struct LikeListResponse {
    1: base.Status base;
    2: LikeListResponseData data;
}

struct CommentPublishRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string video_id;
    4: string comment_id;
    5: required string content;
}

struct CommentPublishResponse {
    1: base.Status base;
}

struct CommentListRequest {
    1: string video_id (api.query="video_id");
    2: string comment_id (api.query="comment_id");
    3: i64 page_size (api.query="page_size");
    4: i64 page_num (api.query="page_num");
    5: string access_token (api.header="Access-Token");
    6: string refresh_token (api.header="Refresh-Token");
}

struct CommentListResponseData {
    required list<base.Comment> items;
}
struct CommentListResponse {
    1: base.Status base;
    2: CommentListResponseData data;
}

struct CommentDeleteRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string video_id;
    4: string comment_id;
}

struct CommentDeleteResponse {
    1: base.Status base;
}

service InteractService {
    LikeActionResponse LikeAction(1: LikeActionRequest request);
    LikeListResponse LikeList(1: LikeListRequest request);
    CommentPublishResponse CommentPublish(1: CommentPublishRequest request);
    CommentListResponse CommentList(1: CommentListRequest request);
    CommentDeleteResponse CommentDelete(1: CommentDeleteRequest request);
}
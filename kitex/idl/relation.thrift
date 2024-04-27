namespace go relation

include "base.thrift"

struct RelationActionRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: required string to_user_id;
    4: required i64 action_type;
}

struct RelationActionResponse {
    1: base.Status base;
}

struct FollowingListRequest {
    1: string user_id (api.query="user_id");
    2: i64 page_num (api.query="page_num");
    3: i64 page_size (api.query="page_size");
    4: string access_token (api.header="Access-Token");
    5: string refresh_token (api.header="Refresh-Token");
}

struct FollowingListResponseData {
    1: required list<base.UserLite> items;
    2: required i64 total;
}
struct FollowingListResponse {
    1: base.Status base;
    2: FollowingListResponseData data;
}

struct FollowerListRequest {
    1: string user_id (api.query="user_id");
    2: i64 page_num (api.query="page_num");
    3: i64 page_size (api.query="page_size");
    4: string access_token (api.header="Access-Token");
    5: string refresh_token (api.header="Refresh-Token");
}

struct FollowerListResponseData {
    1: required list<base.UserLite> items;
    2: required i64 total;
}
struct FollowerListResponse {
    1: base.Status base;
    2: FollowerListResponseData data;
}

struct FriendListRequest {
    1: i64 page_num (api.query="page_num");
    2: i64 page_size (api.query="page_size");
    3: string access_token (api.header="Access-Token");
    4: string refresh_token (api.header="Refresh-Token");
}

struct FriendListResponseData {
    1: required list<base.UserLite> items;
    2: required i64 total;
}
struct FriendListResponse {
    1: base.Status base;
    2: FriendListResponseData data;
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest request);
    FollowingListResponse FollowingList(1: FollowingListRequest request);
    FollowerListResponse FollowerList(1: FollowerListRequest request);
    FriendListResponse FriendList(1: FriendListRequest request);
}
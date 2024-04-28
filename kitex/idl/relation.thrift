namespace go relation

include "base.thrift"

struct RelationActionRequest {
    1: required string from_user_id;
    3: required string to_user_id;
    4: required i64 action_type;
}

struct RelationActionResponse {
    1: base.Status base;
}

struct FollowingListRequest {
    1: required string user_id;
    2: i64 page_num;
    3: i64 page_size;
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
    1: required string user_id;
    2: i64 page_num;
    3: i64 page_size;
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
    1: required string user_id;
    2: i64 page_num;
    3: i64 page_size;
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
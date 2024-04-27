namespace go video

include "base.thrift"

struct VideoFeedRequest {
    1: string latest_time (api.query="latest_time");
}

struct VideoFeedResponseData {
    list<base.Video> items;
}
struct VideoFeedResponse {
    1: base.Status base;
    2: VideoFeedResponseData data;
}

struct VideoPublishStartRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string title (vt.max_size = "1");
    4: string description;
    5: i64 chunk_total_number (vt.gt = "0");
}

struct VideoPublishStartResponse {
    1: base.Status base;
    2: string uuid;
}

struct VideoPublishUploadingRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: required string uuid;
    4: binary data;
    5: required string md5;
    6: required bool is_m3u8;
    7: required string filename;
    8: required i64 chunk_number;
}

struct VideoPublishUploadingResponse {
    1: base.Status base;
}

struct VideoPublishCompleteRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: required string uuid;
}

struct VideoPublishCompleteResponse {
    1: base.Status base;
}

struct VideoPublishCancleRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: required string uuid;
}

struct VideoPublishCancleResponse {
    1: base.Status base;
}

struct VideoListRequest {
    1: string user_id (api.query="user_id");
    2: i64 page_num (api.query="page_num");
    3: i64 page_size (api.query="page_size");
    4: string access_token (api.header="Access-Token");
    5: string refresh_token (api.header="Refresh-Token");
}

struct VideoListResponseData {
    1: list<base.Video> data;
    2: i64 total;
}
struct VideoListResponse {
    1: base.Status base;
    2: VideoListResponseData data;
}

struct VideoPopularRequest {
    1: i64 page_num (api.query="page_num");
    2: i64 page_size (api.query="page_size");
    3: string access_token (api.header="Access-Token");
    4: string refresh_token (api.header="Refresh-Token");
}

struct VideoPopularResponseData {
    list<base.Video> items;
}
struct VideoPopularResponse {
    1: base.Status base;
    2: VideoPopularResponseData data;
}

struct VideoSearchRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string keywords;
    4: required i64 page_num;
    5: required i64 page_size;
    6: i64 from_date;
    7: i64 to_date;
    8: string username;
}

struct VideoSearchResponseData {
    1: list<base.Video> items;
    2: i64 total;
}
struct VideoSearchResponse {
    1: base.Status base;
    2: VideoSearchResponseData data;
}

struct VideoVisitRequest {
}

struct VideoVisitResponse {
    1: base.Status base;
    2: base.Video item;
}

service VideoService {
    VideoFeedResponse Feed(1: VideoFeedRequest request);
    VideoPublishStartResponse VideoPublishStart(1: VideoPublishStartRequest request);
    VideoPublishUploadingResponse VideoPublishUploading(1: VideoPublishUploadingRequest request);
    VideoPublishCompleteResponse VideoPublishComplete(1: VideoPublishCompleteRequest request);
    VideoPublishCancleResponse VideoPublishCancle(1: VideoPublishCancleRequest request);
    VideoListResponse List(1: VideoListRequest request);
    VideoPopularResponse Popular(1: VideoPopularRequest request);
    VideoSearchResponse Search(1: VideoSearchRequest request);
    VideoVisitResponse Visit(1: VideoVisitRequest request);
}
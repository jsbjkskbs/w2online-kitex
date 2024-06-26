namespace go video

include "base.thrift"

struct VideoFeedRequest {
    1: string latest_time;
}

struct VideoFeedResponseData {
    list<base.Video> items;
}
struct VideoFeedResponse {
    1: base.Status base;
    2: VideoFeedResponseData data;
}

struct VideoPublishStartRequest {
    1: required string user_id;
    3: required string title (vt.min_size = "1");
    4: required string description;
    5: required i64 chunk_total_number (vt.gt = "0");
}

struct VideoPublishStartResponse {
    1: base.Status base;
    2: string uuid;
}

struct VideoPublishUploadingRequest {
    1: required string user_id;
    2: required string uuid;
    3: required binary data;
    4: required string md5;
    5: required bool is_m3u8;
    6: required string filename;
    7: required i64 chunk_number;
}

struct VideoPublishUploadingResponse {
    1: base.Status base;
}

struct VideoPublishCompleteRequest {
    1: required string user_id;
    2: required string uuid;
}

struct VideoPublishCompleteResponse {
    1: base.Status base;
}

struct VideoPublishCancleRequest {
    1: required string user_id;
    2: required string uuid;
}

struct VideoPublishCancleResponse {
    1: base.Status base;
}

struct VideoListRequest {
    1: required string user_id;
    2: i64 page_num;
    3: i64 page_size;
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
    1: i64 page_num;
    2: i64 page_size;
}

struct VideoPopularResponseData {
    list<base.Video> items;
}
struct VideoPopularResponse {
    1: base.Status base;
    2: VideoPopularResponseData data;
}

struct VideoSearchRequest {
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

struct VideoInfoRequest {
    1: required string video_id;
}

struct VideoInfoResponseData {
    1: base.Video item;
}
struct VideoInfoResponse {
    1: base.Status base;
    2: VideoInfoResponseData data;
}

struct VideoDeleteRequest {
    1: required string video_id;
}

struct VideoDeleteResponse {
    1: base.Status base;
}

struct VideoVisitRequest {
    1: required string from_ip
    2: required string video_id
}

struct VideoVisitResponse {
    1: base.Status base;
    2: base.Video item;
}

struct VideoIdListRequest {
    1: i64 page_num;
    2: i64 page_size;
}

struct VideoIdListResponse {
    1: base.Status base;
    2: bool is_end;
    3: list<string> list;
}

struct UpdateVisitCountRequest {
    1: string video_id;
    2: i64 visit_count
}

struct UpdateVisitCountResponse {
    1: base.Status base;
}

struct GetVideoVisitCountRequest {
    1: string video_id
}

struct GetVideoVisitCountResponse {
    1: base.Status base;
    2: i64 visit_count;
}

struct GetVideoVisitCountInRedisRequest {
    1: string video_id;
}

struct GetVideoVisitCountInRedisResponse {
    1: i64 visit_count;
    2: base.Status base;
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
    VideoInfoResponse Info(1: VideoInfoRequest request);
    VideoDeleteResponse Delete(1: VideoDeleteRequest request);
    VideoIdListResponse IdList(1: VideoIdListRequest request);
    VideoVisitResponse VideoVisit(1: VideoVisitRequest request);
    UpdateVisitCountResponse UpdateVisitCount(1: UpdateVisitCountRequest request);
    GetVideoVisitCountResponse GetVideoVisitCount(1: GetVideoVisitCountRequest request);
    GetVideoVisitCountInRedisResponse GetVideoVisitCountInRedis(1: GetVideoVisitCountInRedisRequest request);
}
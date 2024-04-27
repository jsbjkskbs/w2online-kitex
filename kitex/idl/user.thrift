namespace go user

include "base.thrift"

struct UserRegisterRequest {
    1: required string username;
    2: required string password (vt.min_size = "5", vt.max_size = "18");
}

struct UserRegisterResponse {
    1: base.Status base;
    //2: string access_token;
    //3: string refresh_token;
}

struct UserLoginRequest {
    1: required string username;
    2: required string password (vt.min_size = "5", vt.max_size = "18");
    3: string code;
}

struct UserLoginResponse {
    1: base.Status base;
    2: base.User data;
    //3: string access_token;
    //4: string refresh_token;
}

struct UserInfoRequest {
    1: string user_id (api.query="user_id");
    2: string token (api.query="token");
    3: string access_token (api.header="Access-Token");
    4: string refresh_token (api.header="Refresh-Token");
}

struct UserInfoResponse {
    1: base.Status base;
    2: base.User data;
}

struct UserAvatarUploadRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: binary data;
}

struct UserAvatarUploadResponse {
    1: base.Status base;
    2: base.User data;
}

struct AuthMfaQrcodeRequest { 
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
}

struct Qrcode {
    1: string secret;
    2: string qrcode;
}
struct AuthMfaQrcodeResponse {
    1: base.Status base;
    2: Qrcode data;
}

struct AuthMfaBindRequest {
    1: string access_token (api.header="Access-Token");
    2: string refresh_token (api.header="Refresh-Token");
    3: string code;
    4: string secret;
}

struct AuthMfaBindResponse {
    1: base.Status base;
}

service UserService {
    UserRegisterResponse Register(1: UserRegisterRequest request);
    UserLoginResponse Login(1: UserLoginRequest request);
    UserInfoResponse Info(1: UserInfoRequest request);
    UserAvatarUploadResponse AvatarUpload(1: UserAvatarUploadRequest request);
    AuthMfaQrcodeResponse AuthMfaQrcode(1: AuthMfaQrcodeRequest request);
    AuthMfaBindResponse AuthMfaBind(1: AuthMfaBindRequest request);
}